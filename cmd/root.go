// this file contains the command line interface for the application

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stargazer",
	Short: "list github repositories with star count",
	Long:  "list github repositories with star count",
	Run: func(cmd *cobra.Command, args []string) {
		var repos = getRepoList("gennaro-tedesco")
		getStars(repos)
		getForks(repos)
		getUrl(repos)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
