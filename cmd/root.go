package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stargazer",
	Args:  cobra.ExactArgs(1),
	Short: "stargazer: retrieve github statistics",
	Long:  "stargazer: CLI to retrieve statistics of users' github repositories",
	Run: func(cmd *cobra.Command, args []string) {
		var repos = getRepoList(args[0])
		getUrl(repos)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().BoolP("sort", "s", false, "sort repositories by feature counts")
}
