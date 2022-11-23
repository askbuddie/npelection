/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"npelection/app"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	district      string
	numParties    int
	numCandidates int
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get election data for certain district",
	RunE: func(cmd *cobra.Command, _ []string) error {
		// Get the link for the district to scrape
		link, ok := districtLinkMap[district]
		// If the district is not found, return an error
		if !ok {
			closestDistrict := app.GetClosestDistrict(district, districtLinkMap)
			fmt.Printf("Incorrect District Name. Did you mean %s?\n", closestDistrict)
			fmt.Printf("Use `npelection list` to get a list of all districts.\n\n")
			cmd.Help()
			return nil
		}
		// Get the HTML document
		response, err := http.Get(link)
		if err != nil {
			e := fmt.Errorf("error getting response: %v", err)
			return e
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			e := fmt.Errorf("status code error- Got: %v Wanted: %v", response.StatusCode, http.StatusOK)
			return e
		}
		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			e := fmt.Errorf("error parsing response body: %v", err)
			return e
		}

		// Find the table with constituency wise data
		constituencyTableHTML := doc.Find("body > section > section:nth-child(5) > div > div > div.gy-4 > div.election-2079 > div > div > div")
		constituencyTable := tablewriter.NewWriter(os.Stdout)
		constituencyTable.SetHeader([]string{"Constituency", "Candidate", "Party", "Vote", "Won"})
		constituencyTable.SetAutoMergeCellsByColumnIndex([]int{0})
		constituencyTable.SetAlignment(tablewriter.ALIGN_LEFT)
		constituencyTable.SetRowLine(true)

		// For each constituency, scrape the candidate name, party name, vote count and if their win condition.
		constituencyTableHTML.Each(func(_ int, constituencySelection *goquery.Selection) {
			constituencyName := constituencySelection.Find(".card-header > .card-title").Text()
			candidatesListHTML := constituencySelection.Find(".candidate-list__item")
			candidatesListHTML.EachWithBreak(func(i int, candidateSelection *goquery.Selection) bool {
				if i >= numCandidates {
					return false
				}
				candidateName := candidateSelection.Find(".nominee-name > a").Text()
				candidatePartyName := candidateSelection.Find(".candidate-party-name > a").Text()
				voteCount := candidateSelection.Find(".vote-count").Text()
				won := ""
				if candidateSelection.HasClass("elected") {
					won = "X"
				}
				constituencyTable.Append([]string{
					strings.TrimSpace(constituencyName),
					strings.TrimSpace(candidateName),
					strings.TrimSpace(candidatePartyName),
					strings.TrimSpace(voteCount),
					won,
				})
				return true
			})
		})
		constituencyTable.Render()

		// Find the table with overall party wise data
		partyTableHTML := doc.Find("body > section > section:nth-child(5) > div > div > div.row > div.col-xl-8 > div > div.col-xl-6.parties.mb-xl-0.mb-4 > div")
		partiesRow := partyTableHTML.Find(".card-body > .row")
		// Prepare table output in display
		partyTable := tablewriter.NewWriter(os.Stdout)
		partyTable.SetHeader([]string{"Party", "Win", "Lead"})
		partyTable.SetRowLine(true)
		// For each party scrape the partyName, wins and leads in the given district
		partiesRow.EachWithBreak(func(i int, partySelection *goquery.Selection) bool {
			if i >= numParties {
				return false
			}
			partyName := partySelection.Find(".party-name > a").Text()
			numbers := partySelection.Find(".number-display")
			if numbers.Length() != 2 {
				log.Fatal(fmt.Errorf("unexpected html for wins and leads: %v", numbers.Nodes))
				return false
			}
			wins := numbers.First().Text()
			leads := numbers.Last().Text()
			partyTable.Append([]string{
				strings.TrimSpace(partyName),
				strings.TrimSpace(wins),
				strings.TrimSpace(leads),
			})
			return true
		})
		partyTable.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&district, "district", "d", "", "District for which the election data is to be shown")
	getCmd.Flags().IntVarP(&numParties, "parties", "p", 3, "Number of parties to show in the overall result")
	getCmd.Flags().IntVarP(&numCandidates, "candidates", "c", 3, "Number of candidates to show for each constituency")
	getCmd.MarkFlagRequired("district")
}
