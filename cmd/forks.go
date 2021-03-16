package cmd

import (
	"github.com/spf13/cobra"
)

var forksCmd = &cobra.Command{
	Use:   "forks",
	Args:  cobra.ExactArgs(1),
	Short: "forks count per repository",
	Long:  "list user's GitHub repositories with corresponding forks count",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepoList(args[0])
		sort, _ := cmd.Flags().GetBool("sort")
		getForks(repos, sort)
	},
}

func init() {
	rootCmd.AddCommand(forksCmd)
}
