package display

import (
	"fmt"
	"strings"

	"github.com/yigitdemirko/currency-rates/scraper"
)

type Formatter struct{}

func New() *Formatter {
	return &Formatter{}
}

func (f *Formatter) DisplayRates(rates []scraper.Rate, filter string) error {
	// Print header
	fmt.Println("| Ticker       | Price        | Change    |")
	fmt.Println("|--------------|--------------|-----------|")

	// Print rates
	for _, rate := range rates {
		if !matchesFilter(rate.Ticker, filter) {
			continue
		}
		
		price := formatPrice(rate)
		fmt.Printf("| %-12s | %-12s | %-9s |\n", rate.Ticker, price, rate.Change)
	}

	return nil
}

func matchesFilter(ticker, filter string) bool {
	if filter == "" {
		return true
	}
	return strings.Contains(strings.ToLower(ticker), strings.ToLower(filter))
}

func formatPrice(rate scraper.Rate) string {
	if rate.Price == "" {
		return ""
	}
	
	price := strings.TrimPrefix(rate.Price, "$") // Remove any existing $ sign
	
	switch rate.RateType {
	case scraper.USD:
		return "$" + price
	case scraper.TRY:
		return "₺" + price
	default:
		return price
	}
} 