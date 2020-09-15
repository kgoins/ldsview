package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/kgoins/ldsview/internal"
	ldsview "github.com/kgoins/ldsview/pkg"
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
	Run: func(cmd *cobra.Command, args []string) {
		dumpFile, _ := cmd.Flags().GetString("file")

		builder := ldsview.NewLdifParser(dumpFile)

		filter, err := buildEntityFilter(cmd, args)
		if err != nil {
			fmt.Printf("Unable to parse filter\n")
			return
		}
		builder.SetEntityFilter(filter)

		attrFilter := buildAttrFilter(cmd)
		builder.SetAttributeFilter(attrFilter)

		count, _ := cmd.Flags().GetBool("count")
		if count {
			numEntities, _ := builder.CountEntities()
			fmt.Printf("Entities: %d\n", numEntities)
			return
		}

		entities, err := builder.BuildEntities()
		if err != nil {
			fmt.Printf("Unable to parse file: %s\n", err.Error())
			return
		}

		printLimit, intParseErr := cmd.Flags().GetInt("first")
		if intParseErr != nil {
			fmt.Printf("Unable to parse value: %s\n", intParseErr.Error())
		}

		if printLimit != 0 {
			entities = entities[:printLimit]
		}

		tdc, _ := cmd.Flags().GetBool("tdc")

		for _, entity := range entities {
			if tdc {
				entity.DeocdeTimestamps()
			}

			PrintEntity(entity)
		}
	},
}

func buildEntityFilter(cmd *cobra.Command, args []string) ([]ldsview.IEntityFilter, error) {
	hailmary, _ := cmd.Flags().GetBool("hailmary")

	var err error
	var filter []ldsview.IEntityFilter

	if len(args) < 1 {
		return filter, nil
	}
	filterStr := args[0]

	if hailmary {
		filter = []ldsview.IEntityFilter{
			ldsview.NewHailmaryFilter(filterStr),
		}
	} else {
		filterParts := strings.Split(filterStr, ",")

		filter, err = ldsview.BuildEntityFilter(filterParts)
		if err != nil {
			return nil, err
		}
	}

	return filter, nil
}

func buildAttrFilter(cmd *cobra.Command) internal.HashSetStr {
	filterParts, _ := cmd.Flags().GetStringSlice("include")
	return ldsview.BuildAttributeFilter(filterParts)
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

	searchCmd.PersistentFlags().BoolP(
		"hailmary",
		"m",
		false,
		"Filter term will be used in a broad search of multiple common attributes",
	)

	searchCmd.PersistentFlags().StringSliceP(
		"include",
		"i",
		[]string{},
		"Select which attributes are displayed from the returned entities",
	)
}
