package cnki

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
)

const (
	DOMAIN   = "https://kns.cnki.net"
	StartUrl = DOMAIN + "/KNS8/Brief/GetGridTableHtml"
)

var page = 1



func SpiderCnki() {

	c := colly.NewCollector(
		// 开启本机debug
		// colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains("kns.cnki.net"),
		// 防止页面重复下载
		//colly.CacheDir("./pubmed/coursera_cache"),
		//colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)


	c.OnHTML(".result-table-list tbody", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(i int, item *colly.HTMLElement){
			eubDate :=item.DOM.Find(".date").Text()

			url,_ := item.DOM.Find(".name a").Attr("href")

			fmt.Printf("发布时间：%s,请求地址：%s\n",eubDate,url)

			if url != "" {
				//SpiderCnkiItem(c,url)
			}
		})


	})


	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("c 请求地址:%s,---当前页：%d\n", r.URL.String(), page)
		//r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
		//r.Headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

		for k,v := range HeadMap {
			fmt.Printf("set request header -->%s:%s\n",k,v)
			r.Headers.Set(k,v)
		}
	})

	c.OnResponse(func(response *colly.Response) {
		html := string(response.Body)
		fmt.Println(html)
	})

	c.OnScraped(func(r *colly.Response) {
		//page = page + 1
		//if page*10 <= 283 {
		//	nextPage := fmt.Sprintf("%s&page=%d", StartUrl, page)
		//	fmt.Println("next page is", nextPage)
		//	err := c.Visit(nextPage)
		//	if err != nil {
		//		return
		//	}
		//}
		//
		//fmt.Println("Finished", r.Request.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Post(StartUrl,QueryData)

	if err != nil {
		log.Fatal("post failed with response:",err)
	}
}
