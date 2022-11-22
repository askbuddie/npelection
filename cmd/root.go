package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var districtLinkMap map[string]string

var rootCmd = &cobra.Command{
	Use:   "npelection",
	Short: "Get the election vote counts right from your terminal.",
	Long: `NP Election commands helps you to get the vote count of
	current year election happening in Nepal.`,
}

func Execute(data string) {
	json.Unmarshal([]byte(data), &districtLinkMap)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
