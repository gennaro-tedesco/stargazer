package cmd

import (
	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Args:  cobra.ExactArgs(1),
	Short: "repositories urls",
	Long:  "list user's GitHub repositories with corresponding complete urls",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepoList(args[0], true)
		getUrl(repos)
	},
}

func init() {
	urlCmd.SetHelpTemplate(getUrlHelp())
	rootCmd.AddCommand(urlCmd)
}

func getUrlHelp() string {
	return `
list user's github repositories with corresponding complete urls

Usage:
  stargazer url username

Flags:
  -h, --help   help for url
`
}
