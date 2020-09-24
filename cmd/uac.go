package cmd

import (
	"fmt"
	"strconv"

	ldsview "github.com/kgoins/ldsview/pkg"
	"github.com/spf13/cobra"
)

// uacCmd represents the uac command
var uacCmd = &cobra.Command{
	Use:   "uac",
	Short: "Parses a useraccountcontrol attribute value as an int64 into its flag components",
	Long:  `Example: ldsview uac 512`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		uacInt, parseErr := strconv.ParseInt(args[0], 10, 64)
		if parseErr != nil {
			fmt.Println("Unable to parse input value into a 64-bit int")
			return
		}

		uacFlags, err := ldsview.GetFlagsFromUAC(uacInt)
		if err != nil {
			fmt.Println("Unable to parse UAC: ", err)
			return
		}

		for _, flag := range uacFlags {
			fmt.Println(flag)
		}
	},
}

func init() {
	rootCmd.AddCommand(uacCmd)
}
