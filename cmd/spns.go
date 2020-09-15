package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ldsview "github.com/kgoins/ldsview/pkg"
)

var spnsCmd = &cobra.Command{
	Use:   "spns",
	Short: "Display entities with service principal names set",
	Run: func(cmd *cobra.Command, args []string) {
		dumpFile, _ := cmd.Flags().GetString("file")
		builder := ldsview.NewLdifParser(dumpFile)

		filterParts := []string{"servicePrincipalName~:/"}

		getUsers, _ := cmd.Flags().GetBool("users")
		if getUsers {
			filterParts = []string{
				"objectClass=:user",
				"objectClass!=:computer",
				"servicePrincipalName~:/",
			}
		}

		filter, _ := ldsview.BuildEntityFilter(filterParts)
		builder.SetEntityFilter(filter)

		includeFilterparts, _ := cmd.Flags().GetStringSlice("include")
		includeFilter := ldsview.BuildAttributeFilter(includeFilterparts)
		builder.SetAttributeFilter(includeFilter)

		entities, err := builder.BuildEntities()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, entity := range entities {
			PrintEntity(entity)
		}
	},
}

func init() {
	rootCmd.AddCommand(spnsCmd)

	spnsCmd.PersistentFlags().BoolP(
		"users",
		"u",
		false,
		"Find all users with a service principal set",
	)

	spnsCmd.PersistentFlags().StringSliceP(
		"include",
		"i",
		[]string{},
		"Select which attributes are displayed from the returned entities",
	)
}
