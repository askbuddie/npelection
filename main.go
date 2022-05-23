package main

// curl --silent 'https://election.ekantipur.com/pradesh-3/district-kathmandu/kathmandu?lng=eng' | htmlq div.candidate-meta-wrapper | grep 'Balendra' | htmlq div.vote-numbers --text

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
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
	selection.Each(func(i int, groupSelection *goquery.Selection) {
		if i == 0 {
			fmt.Println("---Mayor Results:---")
		} else {
			fmt.Println("---Deputy Mayor Results:---")
		}
		groupSelection.Find("div.candidate-meta-wrapper").EachWithBreak(func(j int, candidateSelection *goquery.Selection) bool {
			if j >= 3 {
				return false
			}
			name := candidateSelection.Find("div.candidate-name").Text()
			vote := candidateSelection.Find("div.vote-numbers").Text()

			fmt.Printf("%s\t%s\n", name, vote)
			return true
		})
	})
}
