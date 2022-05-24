package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "npelection",
	Short: "Get the election vote counts right from your terminal.",
	Long: `NP Election commands helps you to get the vote count of
	current year election happening in Nepal.`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    os.Exit(1)
  }
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


