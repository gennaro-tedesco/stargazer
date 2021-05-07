package cmd

import (
	"github.com/spf13/cobra"
)

var starredCmd = &cobra.Command{
	Use:   "starred",
	Args:  cobra.ExactArgs(1),
	Short: "list user's starred repositories",
	Long:  "list user's starred repositories",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepoList(args[0], false)
		getUrl(repos)
	},
}

func init() {
	starredCmd.SetHelpTemplate(getStarredHelp())
	rootCmd.AddCommand(starredCmd)
}

func getStarredHelp() string {
	return `
list user's starred repositories

Usage:
  stargazer starred username

Flags:
  -h, --help   help for stats
`
}
