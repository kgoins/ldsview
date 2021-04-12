package cmd

import (
	"fmt"
	"log"

	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/spf13/cobra"
)

var valuesCmd = &cobra.Command{
	Use:   "values attributeName",
	Short: "Extract an ldap object with a given attribute value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dumpFile, _ := cmd.Flags().GetString("file")
		attrName := args[0]

		parser := ldsview.NewLdifParser(dumpFile)
		vals, err := ldsview.GetValues(&parser, attrName)
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
