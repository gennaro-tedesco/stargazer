package cmd

import (
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Args:  cobra.ExactArgs(1),
	Short: "stats counts per repository",
	Long: "list user's github repositories showing number of stars and forks",
	Run: func(cmd *cobra.Command, args []string) {
		repos := getRepoList(args[0])
		sort, _ := cmd.Flags().GetBool("sort")
		getStats(repos, sort)
	},
}

func init() {
	statsCmd.Flags().BoolP("sort", "s", true, "sort repositories by feature counts")
	statsCmd.SetHelpTemplate(GetStatsHelp())
	rootCmd.AddCommand(statsCmd)
}

func GetStatsHelp() string {
	return `
list user's github repositories showing number of stars and forks

Usage:
  stargazer stats [username] [flags]

Flags:
  -h, --help   help for stats
  -s, --sort   sort repositories by feature counts (default true)
`
}
