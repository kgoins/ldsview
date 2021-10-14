package internal

import (
	"log"

	"github.com/kgoins/ldsview/svccontainer"
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func BulidContainerFromFlags(cmd *cobra.Command) (container di.Container, err error) {
	conf := viper.New()
	conf.BindPFlags(cmd.Flags())

	container, err = svccontainer.BuildContainer(conf)
	if err != nil {
		log.Fatalf("Unable to bootstrap app services: %s", err.Error())
	}

	return
}
