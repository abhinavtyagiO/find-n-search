package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type data struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No file provided")
		os.Exit(-1)
	}
	fileName := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "No search term provided")
		os.Exit(-1)
	}

	var (
		items []data
		terms []string
		input io.ReadCloser
		count int
		err   error
	)

	if input, err = os.Open(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Bad File: %s\n", err)
		os.Exit(-1)
	}

	// decode json file to go objects
	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "Bad JSON: %s\n", err)
		os.Exit(-1)
	}

	// get search terms
	for _, term := range os.Args[2:] {
		terms = append(terms, strings.ToLower(term))
	}

	// find the comic/s using the search terms
outer:
	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer
			}
		}
		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s %s\n", item.Num, item.Day, item.Month, item.Year, item.Title)
		count++
	}
	fmt.Fprintf(os.Stderr, "Found %d comics\n", count)
}
