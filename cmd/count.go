package cmd

import (
	"fmt"
	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Counts the number of entities in an ldif file",
	Run: func(cmd *cobra.Command, args []string) {
		dumpFile, _ := cmd.Flags().GetString("file")
		builder := ldsview.NewLdifParser(dumpFile)

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
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().Bool( "count", true, "" )
	countCmd.Flags().MarkHidden("count")
}
