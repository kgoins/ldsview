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
		parser := ldsview.NewLdifParser(dumpFile)

		count, err := parser.CountEntities()
		if err != nil {
			fmt.Printf("Unable to parse file: %s\n", err.Error())
			return
		}

		fmt.Println("Entities: ", count)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().Bool("count", true, "")
	countCmd.Flags().MarkHidden("count")
}
