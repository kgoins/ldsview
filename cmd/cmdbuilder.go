package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/kgoins/ldsview/internal"
)

// cmdbuilderCmd represents the uac command
var cmdbuilderCmd = &cobra.Command{
	Use:   "cmdbuilder",
	Short: "Builds the ldapsearch command needed to extract an ldif",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Host: ")
		var host string
		fmt.Scanln(&host)

		fmt.Print("Domain DN: ")
		var domainDN string
		fmt.Scanln(&domainDN)

		fmt.Print(`User (domain\username): `)
		var user string
		fmt.Scanln(&user)

		fmt.Print("Password: ")
		var password string
		fmt.Scanln(&password)

		options := internal.NewLdapsearchCmdOptions(
			host,
			domainDN,
			user,
			password,
		)

		cmdStr := internal.BuildLdapsearchCmd(options)
		fmt.Println(cmdStr)
	},
}

func init() {
	rootCmd.AddCommand(cmdbuilderCmd)
}
