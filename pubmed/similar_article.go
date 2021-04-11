package pubmed

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"time"
)

func SpiderSimilarArticle(collector *colly.Collector, url string){
	collector = collector.Clone()

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})

	collector.OnRequest(func(request *colly.Request) {
		log.Println("articles visit: ", request.URL.String())
	})

	collector.OnResponse(func(response *colly.Response) {
		html := string(response.Body)
		fmt.Printf("artilce html：%s",html)
	})

	// 解析详情页数据
	collector.OnHTML(".article-details", func(e *colly.HTMLElement) {
		var affiliations []string

		e.DOM.Find(".expanded-authors > .affiliations > .item-list li").Each(func(i int, item *goquery.Selection) {
			affiliations =  append(affiliations, 	item.Text())
		})






	})

	collector.Visit(url)
}