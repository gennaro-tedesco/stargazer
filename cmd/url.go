package cmd

import (
	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Args:  cobra.ExactArgs(1),
	Short: "repositories urls",
	Long: "list user's GitHub repositories with corresponding complete urls",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepoList(args[0])
		getUrl(repos)
	},
}

func init() {
	urlCmd.SetHelpTemplate(GetHelpText("cmd/url_help.txt"))
	rootCmd.AddCommand(urlCmd)
}
