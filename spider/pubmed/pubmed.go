package pubmed

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
)

const (
	DOMAIN   = "https://pubmed.ncbi.nlm.nih.gov"
	StartUrl = DOMAIN + "/?term=Ultrasound+in+fatty+liver+with+chronic+hepatitis+B&filter=datesearch.y_5"
)

var page = 1


func SpiderPubmed() {

	c := colly.NewCollector(
		// 开启本机debug
		// colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains("pubmed.ncbi.nlm.nih.gov"),
		// 防止页面重复下载
		//colly.CacheDir("./pubmed/coursera_cache"),
		//colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)


	c.OnHTML(".search-results", func(e *colly.HTMLElement) {
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
			print("href,found", href, found)
			// 如果找到了详情页，则继续下一步的处理

			if found {
				SpiderArticle(c, href)
				log.Println(href)
			}

		})
	})


	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("c 请求地址:%s,---当前页：%d\n", r.URL.String(), page)
		//r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	})

	c.OnResponse(func(response *colly.Response) {
		//html := string(response.Body)
		//fmt.Println(html)
	})

	c.OnScraped(func(r *colly.Response) {
		page = page + 1
		if page*10 <= 61 {
			nextPage := fmt.Sprintf("%s&page=%d", StartUrl, page)
			fmt.Println("next page is", nextPage)
			err := c.Visit(nextPage)
			if err != nil {
				return
			}
		}

		fmt.Println("Finished", r.Request.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping on https://hackerspaces.org
	//startPage := fmt.Sprintf("%s&page=%d", StartUrl, page)
	err := c.Visit(StartUrl)
	if err != nil {
		return
	}
}
