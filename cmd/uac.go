package cmd

import (
	"os"

	ldsview "github.com/kgoins/ldsview/pkg"
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
			ldsview.UACPrint(os.Stdout)
			return
		}

		shouldSearch, _ := cmd.Flags().GetInt("search")
		if shouldSearch != 0 {
			file, _ := cmd.Flags().GetString("file")
			ldifFile := ldsview.NewLdifParser(file)
			entities, err := ldifFile.BuildEntities()
			if err != nil {
				cmd.PrintErr("Error while parsing entities: ", err)
				return
			}

			matches := ldsview.UACSearch(&entities, shouldSearch)
			for _, match := range matches {
				PrintEntity(match)
			}
			return
		}

		if len(args) > 0 {
			uacFlags, err := ldsview.UACParse(args[0])
			if err != nil {
				cmd.PrintErr(err)
				cmd.Help()
			}

			for _, flag := range uacFlags {
				cmd.Println(flag)
			}
			return
		}
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(uacCmd)

	uacCmd.PersistentFlags().Bool(
		"list",
		false,
		"Lists the available UAC properties by which to search",
	)

	uacCmd.PersistentFlags().Int(
		"search",
		0,
		"UAC property by which to search",
	)
}
