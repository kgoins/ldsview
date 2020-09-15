package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ldsview "github.com/kgoins/ldsview/pkg"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Counts the number of entities in an ldif file",
	Run: func(cmd *cobra.Command, args []string) {
		dumpFile, _ := cmd.Flags().GetString("file")

		builder := ldsview.NewLdifParser(dumpFile)
		count, err := builder.CountEntities()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		fmt.Printf("Entities: %d\n", count)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}
