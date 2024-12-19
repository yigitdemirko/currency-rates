package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yigitdemirko/currency-rates/display"
	"github.com/yigitdemirko/currency-rates/scraper"
)

const helpText = `Currency Rates CLI 💰

A simple command-line tool to fetch and display current currency rates for Turkish Lira from [doviz.com](https://www.doviz.com/).

Usage:
  currency-rates [options]

Options:
  -filter string   Filter by currency name (e.g., DOLAR, EURO)
  -h, --help      Show this help message

Examples:
  currency-rates              # Show all rates
  currency-rates -filter DOLAR # Show only DOLAR rates

Data source: https://doviz.com/
`

func main() {
	// Check for help flag before parsing other flags
	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" {
			fmt.Print(helpText)
			os.Exit(0)
		}
	}

	url := flag.String("url", "https://www.doviz.com/", "URL to scrape rates from")
	filter := flag.String("filter", "", "Filter by currency name")

	flag.Parse()

	s := scraper.New(*url)
	rates, err := s.FetchRates()
	if err != nil {
		log.Fatal(err)
	}

	f := display.New()
	if err := f.DisplayRates(rates, *filter); err != nil {
		log.Fatal(err)
	}
} 