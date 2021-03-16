package cmd

import (
	"github.com/spf13/cobra"
)

var starsCmd = &cobra.Command{
	Use:   "stars",
	Args:  cobra.ExactArgs(1),
	Short: "stars count per repository",
	Long:  "list user's GitHub repositories with corresponding stars count",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepoList(args[0])
		sort, _ := cmd.Flags().GetBool("sort")
		getStars(repos, sort)
	},
}

func init() {
	rootCmd.AddCommand(starsCmd)
}
