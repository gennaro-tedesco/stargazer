package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.ExactArgs(1),
	Short: "list user's repositories names and descriptions",
	Long:  "list user's repositories names and descriptions",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepoList(args[0], true)
		getList(repos)
	},
}

func init() {
	listCmd.SetHelpTemplate(getListHelp())
	rootCmd.AddCommand(listCmd)
}

func getListHelp() string {
	return `
list user's repositories names and descriptions

Usage:
  stargazer list username

Flags:
  -h, --help   open this help page
`
}
