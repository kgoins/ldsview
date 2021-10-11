package cmd

import (
	"fmt"
	"os"

	"github.com/kgoins/ldsview/internal"
	"github.com/kgoins/snakecharmer/snakecharmer"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ldsview",
	Short: "CLI application to parse offline dumps from ldapsearch queries",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		snakeCharmer := snakecharmer.NewSnakeCharmer("LDSVIEW", ".ldsview")
		confPath, _ := cmd.Flags().GetString("config")
		return snakeCharmer.InitConfig(cmd, confPath)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if version, _ := cmd.Flags().GetBool("version"); version == true {
			internal.PrintVersionInfo()
			os.Exit(0)
		}

		cmd.Usage()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String(
		"config",
		"",
		"config file (default is $HOME/.ldsview.yaml)",
	)

	rootCmd.PersistentFlags().StringP(
		"file",
		"f",
		"",
		"path to ldapsearch dump file",
	)
	rootCmd.MarkFlagRequired("file")

	rootCmd.PersistentFlags().BoolP(
		"verbose",
		"v",
		false,
		"enables info level logging",
	)

	rootCmd.PersistentFlags().Bool(
		"debug",
		false,
		"enables debug level logging",
	)

	rootCmd.PersistentFlags().Bool(
		"version",
		false,
		"Print version information",
	)
}
