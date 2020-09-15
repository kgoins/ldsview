package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	ldsview "github.com/kgoins/ldsview/pkg"
)

// structureCmd represents the structure command
var structureCmd = &cobra.Command{
	Use:   "structure",
	Short: "Extracts the OU path structure of the ldif file",
	Run: func(cmd *cobra.Command, args []string) {
		dumpFile, _ := cmd.Flags().GetString("file")
		source := ldsview.NewLdifParser(dumpFile)

		structure, err := ldsview.GetStructure(&source)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		sort.Strings(structure)

		for _, path := range structure {
			fmt.Println(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(structureCmd)
}
