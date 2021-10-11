package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/audibleblink/msldapuac"
	uac "github.com/audibleblink/msldapuac"
	"github.com/spf13/cobra"
)

// uacCmd represents the uac command
var uacCmd = &cobra.Command{
	Use:   "uac <int>",
	Short: "Parses a useraccountcontrol attribute value as an int64 into its flag components",
	Long:  `Example: ldsview uac 512`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetOut(os.Stdout)

		shouldList, _ := cmd.Flags().GetBool("list")
		if shouldList {
			printUACVals(cmd.OutOrStdout())
			return
		}

		if len(args) < 1 {
			cmd.Help()
			return
		}

		uacInt, err := strconv.ParseInt(args[0], 0, 64)
		if err != nil {
			log.Fatal(err)
		}

		uacFlags, err := uac.ParseUAC(uacInt)
		if err != nil {
			cmd.PrintErr(err)
			cmd.Help()
		}

		for _, flag := range uacFlags {
			cmd.Println(flag)
		}

		return
	},
}

// prints all available UAC values
func printUACVals(dest io.Writer) {
	w := new(tabwriter.Writer)
	w.Init(dest, 8, 8, 0, '\t', 0)
	defer w.Flush()

	template := "%s\t%d\n"
	var sorted []string
	for k, v := range msldapuac.PropertyMap {
		sorted = append(sorted, fmt.Sprintf(template, v, k))
	}

	sort.Strings(sorted)
	fmt.Fprintf(w, "Property\tValue\n")
	fmt.Fprintf(w, "---\t---\n")
	for _, line := range sorted {
		fmt.Fprintf(w, line)
	}
}

func init() {
	rootCmd.AddCommand(uacCmd)

	uacCmd.PersistentFlags().Bool(
		"list",
		false,
		"Lists the available UAC properties by which to search",
	)
}
