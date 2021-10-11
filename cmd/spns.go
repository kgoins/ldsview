package cmd

import (
	"log"
	"strings"

	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/kgoins/ldifparser/entitybuilder"
	"github.com/kgoins/ldsview/internal"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/spf13/cobra"
)

var spnsCmd = &cobra.Command{
	Use:   "spns",
	Short: "Display entities with service principal names set",
	Run: func(cmd *cobra.Command, args []string) {
		svcs, err := internal.BulidContainerFromFlags(cmd)
		if err != nil {
			log.Fatal(err)
		}
		searcher := svcs.Get("ldapsearcher").(searcher.LdapSearcher)

		filterStr := "servicePrincipalName:~/"

		getUsers, _ := cmd.Flags().GetBool("users")
		if getUsers {
			filterParts := []string{
				"objectClass:=user",
				"objectClass:!=computer",
				"servicePrincipalName:~/",
			}

			filterStr = strings.Join(filterParts, ",")
		}

		entityFilter, err := filter.ParseFilterStr(filterStr)
		if err != nil {
			log.Fatal(err)
		}

		includeFilterparts, _ := cmd.Flags().GetStringSlice("include")
		includeFilter := entitybuilder.NewAttributeFilter(includeFilterparts...)

		done := make(chan bool)
		defer close(done)

		entities := searcher.SearchEntities(done, includeFilter, entityFilter)

		err = ChannelPrinter(entities, cmd)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(spnsCmd)

	spnsCmd.PersistentFlags().BoolP("count", "c", false, "")
	spnsCmd.PersistentFlags().Int("first", 0, "Print only the first <n> entries")

	spnsCmd.PersistentFlags().Bool(
		"tdc",
		false,
		"Decodes timestamps to a human readable format",
	)

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
