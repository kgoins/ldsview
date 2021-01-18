package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kgoins/ldsview/internal"
	"github.com/spf13/cobra"
)

// cmdbuilderCmd represents the uac command
var cmdbuilderCmd = &cobra.Command{
	Use:   "cmdbuilder",
	Short: "Builds the ldapsearch command needed to extract an ldif",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Host: ")
		host := getString()

		fmt.Print("Domain DN: ")
		domainDN := getString()

		fmt.Print(`User (domain\username): `)
		user := getString()

		fmt.Print("Password: ")
		password := getString()

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

func getString() (output string) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		output = scanner.Text()
	}
	return
}
