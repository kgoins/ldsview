package cmd

import (
	"log"

	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/kgoins/ldifparser/entitybuilder"
	"github.com/kgoins/ldsview/internal"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [filter]",
	Short: "Searches all entities for results matching the specified filter",
	Long: `The search command will return all entities in the input file matching the user speicified filter. 
Filters are a comma separated list of attribute:value pairs.
Pairs are separated by a filter condition and attributes are case insensitive.

Available conditions:
	FILTER_EQUALS       ":="
	FILTER_NOT_EQUALS   ":!="
	FILTER_CONTAINS     ":~"
	FILTER_NOT_CONTAINS ":!~"

Example: 
    ldsview search "objectclass:=computer,dnshostname:~net"
	--> matches all computer objects with a hostname containing "net"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hailmary, _ := cmd.Flags().GetBool("hailmary")
		if hailmary {
			log.Fatal("hailmary searches are currently disabled")
		}

		svcs, err := internal.BulidContainerFromFlags(cmd)
		if err != nil {
			log.Fatal(err)
		}
		searcher := svcs.Get("ldapsearcher").(searcher.LdapSearcher)

		filterStr := args[0]
		entityFilter, err := filter.ParseFilterStr(filterStr)
		if err != nil {
			log.Fatal("Unable to parse entity filter: " + err.Error())
		}

		attrFilterParts, _ := cmd.Flags().GetStringSlice("include")
		attrFilter := entitybuilder.NewAttributeFilter(attrFilterParts...)

		done := make(chan bool) // ownership passes to SearchEntities
		defer close(done)

		entities := searcher.SearchEntities(done, attrFilter, entityFilter)

		err = ChannelPrinter(entities, cmd)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.PersistentFlags().BoolP("count", "c", false, "")
	searchCmd.PersistentFlags().Int("first", 0, "Print only the first <n> entries")

	searchCmd.PersistentFlags().Bool(
		"tdc",
		false,
		"Decodes timestamps to a human readable format",
	)

	// Disabled
	searchCmd.PersistentFlags().BoolP(
		"hailmary",
		"m",
		false,
		"(DISABLED) Filter term will be used in a broad search of multiple common attributes",
	)

	searchCmd.PersistentFlags().StringSliceP(
		"include",
		"i",
		[]string{},
		"Select which attributes are displayed from the returned entities",
	)
}
