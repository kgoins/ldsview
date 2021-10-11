package cmd

import (
	"fmt"
	"log"

	"github.com/kgoins/ldsview/internal"
	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/kgoins/ldsview/pkg/searcher"
	"github.com/spf13/cobra"
)

var valuesCmd = &cobra.Command{
	Use:   "values attributeName",
	Short: "Extract unique values for input attribute across all entities",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		svcs, err := internal.BulidContainerFromFlags(cmd)
		if err != nil {
			log.Fatal(err)
		}

		attrName := args[0]
		searcher := svcs.Get("ldapsearcher").(searcher.LdapSearcher)

		vals, err := ldsview.GetUniqueValues(searcher, attrName)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if len(vals) == 0 {
			fmt.Println("Value not found")
		}

		for _, val := range vals {
			fmt.Println(val)
		}
	},
}

func init() {
	rootCmd.AddCommand(valuesCmd)
}
