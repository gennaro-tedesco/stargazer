package cmd

import (
	"github.com/spf13/cobra"
)

// forksCmd represents the forks command
var forksCmd = &cobra.Command{
	Use:   "forks",
	Short: "A brief description of your command",
	Long: "A longer description that spans multiple lines and likely contains examples",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepoList(args[0])
		sort, _ := cmd.Flags().GetBool("sort")
		getForks(repos, sort)
	},
}

func init() {
	rootCmd.AddCommand(forksCmd)
}
