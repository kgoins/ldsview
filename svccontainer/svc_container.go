package svccontainer

import (
	"os"

	"github.com/kgoins/ldifparser"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
)

// BuildContainer constructs the di container holding all of
// the app's services from a toml stream, or dies trying
func BuildContainer(confReader *viper.Viper) (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	containerDefs := []di.Def{
		{
			Name:  "ldiffile",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				filePath := confReader.GetString("file")
				return os.Open(filePath)
			},
			Close: func(obj interface{}) error {
				return obj.(*os.File).Close()
			},
		},
		{
			Name:  "ldapsearcher",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				ldifFile := ctn.Get("ldiffile").(ldifparser.ReadSeekerAt)

				s := searcher.NewLdifSearcher(ldifFile)
				return s, nil
			},
		},
	}

	err = builder.Add(containerDefs...)
	if err != nil {
		return nil, err
	}

	return builder.Build(), nil
}
