package cmd

import (
	"fmt"
	"log"
	"os"

	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Counts the number of entities in an ldif file",
	Run: func(cmd *cobra.Command, args []string) {
		inputFilePath, _ := cmd.Flags().GetString("file")

		inputFile, err := os.Open(inputFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer inputFile.Close()

		searcher := searcher.NewLdifSearcher(inputFile)
		count, err := ldsview.CountEntities(searcher)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Entities: ", count)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().Bool("count", true, "")
	countCmd.Flags().MarkHidden("count")
}
