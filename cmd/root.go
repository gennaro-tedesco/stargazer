package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stargazer",
	Args:  cobra.ExactArgs(1),
	Short: "stargazer: retrieve github statistics",
	Long:  "stargazer: CLI to retrieve statistics of github repositories",
	Run: func(cmd *cobra.Command, args []string) {
		var repos = getRepoList(args[0])
		getUrl(repos)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
stargazer: retrieve statistics of github repositories

Arguments:
  username    github username to retrieve statistics of

Usage:
  stargazer [command] username [flag]

Available Commands:
  stats       stats counts per repository
  url         repositories urls

Help commands:
  help        help about any command
  version     prints current version

Flags:
  -h, --help   help for stargazer

Use "stargazer [command] --help" for more information about a command.
`
}
