package demo

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
)



func  SpiderCoinmarketcap()  {
	fName := "cryptocoinmarketcap.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Change (1h)"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML(".cmc-table__table-wrapper-outer tbody tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".cmc-table__column-name"),
			e.ChildText(".cmc-table__cell--sort-by__symbol div"),
			e.ChildText(".cmc-table__cell--sort-by__price div a.cmc-link"),
			e.ChildText(".cmc-table__cell--sort-by__volume-24-h a.cmc-link"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h div.cmc--change-positive"),
			//e.ChildAttr(".market-cap", "data-usd"),
			//e.ChildText(".percent-1h"),
			//e.ChildText(".percent-24h"),
			//e.ChildText(".percent-7d"),
		})
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}