package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RateType int

const (
	TRY RateType = iota // Turkish Lira
	USD                 // US Dollar
	NONE                // No currency symbol
)

type Rate struct {
	Ticker     string
	Price      string
	Change     string
	RateType   RateType
}

type Scraper struct {
	URL string
}

func New(url string) *Scraper {
	return &Scraper{URL: url}
}

func (s *Scraper) FetchRates() ([]Rate, error) {
	res, err := http.Get(s.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed with status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	var rates []Rate
	doc.Find("div.market-data .item").Each(func(_ int, element *goquery.Selection) {
		ticker := element.Find("span.name").Text()
		rate := Rate{
			Ticker: ticker,
			Price:  element.Find("span.value").Text(),
			Change: strings.TrimSpace(element.Find("div.change-rate").Text()),
			RateType: getRateType(ticker),
		}
		rates = append(rates, rate)
	})

	return rates, nil
}

func getRateType(ticker string) RateType {
	ticker = strings.ToUpper(ticker)
	switch {
	case strings.Contains(ticker, "BITCOIN"), strings.Contains(ticker, "BRENT"):
		return USD
	case strings.Contains(ticker, "BIST"):
		return NONE
	default:
		return TRY
	}
} 