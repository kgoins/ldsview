package cmd

import (
	"fmt"
	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/spf13/cobra"
)

var spnsCmd = &cobra.Command{
	Use:   "spns",
	Short: "Display entities with service principal names set",
	Run: func(cmd *cobra.Command, args []string) {
		dumpFile, _ := cmd.Flags().GetString("file")
		builder := ldsview.NewLdifParser(dumpFile)

		filterParts := []string{"servicePrincipalName:~/"}

		getUsers, _ := cmd.Flags().GetBool("users")
		if getUsers {
			filterParts = []string{
				"objectClass:=user",
				"objectClass:!=computer",
				"servicePrincipalName:~/",
			}
		}

		filter, _ := ldsview.BuildEntityFilter(filterParts)
		builder.SetEntityFilter(filter)

		includeFilterparts, _ := cmd.Flags().GetStringSlice("include")
		includeFilter := ldsview.BuildAttributeFilter(includeFilterparts)
		builder.SetAttributeFilter(includeFilter)

		entities := make(chan ldsview.Entity)
		done := make(chan bool)

		// Start the printing goroutine
		go ChannelPrinter(entities, done, cmd)

		err := builder.BuildEntities(entities, done)
		if err != nil {
			fmt.Printf("Unable to parse file: %s\n", err.Error())
			return
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
