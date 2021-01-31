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
			dumpFile, _ := cmd.Flags().GetString("file")
			builder := ldsview.NewLdifParser(dumpFile)
			filter := []ldsview.IEntityFilter {
				ldsview.NewUACFilter(shouldSearch),
			}
			builder.SetEntityFilter(filter)
			attrFilter := buildAttrFilter(cmd)

			usedI := cmd.Flags().Changed("include")
			if usedI {
				// must use AttrFilter if intent is to Match with
				// EntityFilter
				attrFilter.Add("useraccountcontrol")
			}
			builder.SetAttributeFilter(attrFilter)

			entities := make(chan ldsview.Entity)
			done := make(chan bool)

			// Start the printing goroutine
			go ChannelPrinter(entities, done, cmd)
			err := builder.BuildEntities(entities, done)
			if err != nil {
				cmd.PrintErr("Error while parsing entities: ", err)
				return
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

	uacCmd.PersistentFlags().BoolP("count", "c", false, "")
	uacCmd.PersistentFlags().Int("first", 0, "Print only the first <n> entries")

	uacCmd.PersistentFlags().Bool(
		"tdc",
		false,
		"Decodes timestamps to a human readable format",
	)

	uacCmd.PersistentFlags().StringSliceP(
		"include",
		"i",
		[]string{},
		"Select which attributes are displayed from the returned entities",
	)
}
