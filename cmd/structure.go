package cmd

import (
	"fmt"
	"log"
	"sort"

	"github.com/spf13/cobra"

	"github.com/kgoins/ldsview/internal"
	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/kgoins/ldsview/pkg/searcher"
)

// structureCmd represents the structure command
var structureCmd = &cobra.Command{
	Use:   "structure",
	Short: "Extracts the OU path structure of the ldif file",
	Run: func(cmd *cobra.Command, args []string) {
		svcs, err := internal.BulidContainerFromFlags(cmd)
		if err != nil {
			log.Fatal(err)
		}

		searcher := svcs.Get("ldapsearcher").(searcher.LdapSearcher)

		structure, err := ldsview.GetStructure(searcher)
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
