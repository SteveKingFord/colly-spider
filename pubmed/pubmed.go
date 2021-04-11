package pubmed

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
)

const DOMAIN = "https://pubmed.ncbi.nlm.nih.gov"
const START_URL =  DOMAIN + "/?term=intima-media+thickness%E3%80%81Atherosclerotic+Cardiovascular+Disease"


func SpiderPubmed(){
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: pubmed.ncbi.nlm.nih.gov
		colly.AllowedDomains("pubmed.ncbi.nlm.nih.gov"),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)



	// On every a element which has href attribute call callback
	c.OnHTML(".full-docsum", func(e *colly.HTMLElement) {
		//e.ForEach(".docsum-content", func(i int, item *colly.HTMLElement) {
		//	href := item.ChildAttr("a[class='docsum-title']", "href")
		//	title := item.ChildText("a[class='docsum-title']")
		//	author := item.ChildText("span.docsum-authors")
		//
		//	fmt.Println("c1数据：",href,title,author)
		//	//通过Context上下文对象将采集器1采集到的数据传递到采集器2
		//	//url := DOMAIN + href
		//})

		e.DOM.Find(".docsum-content").Each(func(i int, selection *goquery.Selection) {
			href, found := selection.Find("a.docsum-title").Attr("href")
			print("href,found",href,found)
			// 如果找到了详情页，则继续下一步的处理
			url := DOMAIN + href
			if found {
				SpiderArticle(c, url)
				log.Println(href)
			}
		})
	})


	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("c 请求地址", r.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(START_URL)
}