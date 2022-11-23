/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all districts of Nepal",
	Run: func(_ *cobra.Command, _ []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"District"})
		for district := range districtLinkMap {
			table.Append([]string{strings.Title(district)})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
