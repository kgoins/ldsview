package internal

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/kgoins/ldsview/logger"
	"github.com/pelletier/go-toml"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
)

func BulidContainerFromCmd(cmd *cobra.Command) (container di.Container, err error) {
	config, _ := cmd.Flags().GetString("config")

	configFile, err := os.Open(config)
	if err != nil {
		log.Fatalf("Unable to open config file: %s", config)
	}
	defer configFile.Close()

	container, err = BuildContainer(configFile)
	if err != nil {
		log.Fatalf("Unable to bootstrap app services: %s", err.Error())
	}

	return
}

// BuildContainer constructs the di container holding all of
// the app's services from a toml stream, or dies trying
func BuildContainer(confReader io.Reader) (di.Container, error) {
	conf, err := toml.LoadReader(confReader)
	if err != nil {
		return nil, err
	}

	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	containerDefs := []di.Def{
		{
			Name:  "logger",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				loggerSection, ok := conf.Get("logger").(*toml.Tree)
				if !ok {
					return nil, errors.New("Unable to load log config section")
				}

				logConf := logger.LoggerConfig{}
				err := loggerSection.Unmarshal(&logConf)
				if err != nil {
					return nil, err
				}

				return logger.NewZapLogger(logConf), nil
			},
		},
		{
			Name:  "sonar",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				sonarSection, ok := conf.Get("sonar").(*toml.Tree)
				if !ok {
					return nil, errors.New("Unable to load sonar config section")
				}

				unparsedConf := sonar.UnparsedConfig{}
				err := sonarSection.Unmarshal(&unparsedConf)
				if err != nil {
					return nil, err
				}

				sonarConf, err := unparsedConf.BuildConfig()
				if err != nil {
					return nil, err
				}

				db := ctn.Get("sonardb").(sonar.SonarDb)
				log := ctn.Get("logger").(logger.ILogger)

				return sonar.NewDataSource(db, sonarConf, log), nil
			},
		},
		{
			Name:  "datasourcectl",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				static := []datasources.StaticDataSource{
					ctn.Get("sonar").(datasources.StaticDataSource),
				}

				log := ctn.Get("logger").(logger.ILogger)
				return services.NewDataSourceCtrService(log, static...), nil
			},
		},
		{
			Name:  "allsources",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				sources := []datasources.DataSource{
					ctn.Get("sonar").(datasources.StaticDataSource),
				}

				return sources, nil
			},
		},
		{
			Name:  "hostbuilder",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				sources := ctn.Get("allsources").([]datasources.DataSource)
				repo := ctn.Get("hostrepo").(hostrepo.HostRepo)
				log := ctn.Get("logger").(logger.ILogger)

				return services.NewHostBuilder(log, repo, sources...), nil
			},
		},
		{
			Name:  "servicemapper",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				repo := ctn.Get("hostrepo").(hostrepo.HostRepo)
				log := ctn.Get("logger").(logger.ILogger)

				return services.NewServiceMapper(log, repo), nil
			},
		},
	}

	err = builder.Add(containerDefs...)
	if err != nil {
		return nil, err
	}

	return builder.Build(), nil
}
