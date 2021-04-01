package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print current stargazer version",
	Long: "print current stargazer version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.3")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
