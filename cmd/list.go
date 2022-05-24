/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the election data of specific city ( KTM by default )",
	Long: `Displays the list of votes with candidates name of specific city specified.
by default the capital city data is displayed if no city is defined.`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("https://election.ekantipur.com/pradesh-3/district-kathmandu/kathmandu?lng=eng")
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
	
		if response.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
		}
	
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		selection := doc.Find("div.nominee-list-group").Children()
	
		type CandidatesStruct struct {
			Name string
			VoteCount string
			Title string
		}
	
		candidates := [][]string{}
	
		
	
		selection.Each(func(i int, groupSelection *goquery.Selection) {
			
			title := groupSelection.Find("h6.card-title").Text()
			
			groupSelection.Find("div.candidate-meta-wrapper").EachWithBreak(func(j int, candidateSelection *goquery.Selection) bool {
				if j >= 3 {
					return false
				}
	
				name := candidateSelection.Find("div.candidate-name").Text()
				vote := candidateSelection.Find("div.vote-numbers").Text()
	
				candidate := []string{
					name,
					vote,
					title,
				}
	
				candidates = append(
					candidates,
					candidate,
				)
				return true
			})
		})
	
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"NAME", "VOTE COUNT", "TITLE"})
		table.AppendBulk(candidates)
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}