package wanfangdata

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

const (
	StartUrl = "https://s.wanfangdata.com.cn/paper?q=%E8%84%82%E8%82%AA%E8%82%9D%2F%E8%B6%85%E5%A3%B0%20%E6%85%A2%E4%B9%99%E8%82%9D%E5%90%88%E5%B9%B6%E8%84%82%E8%82%AA%E8%82%9D%28%E8%84%82%E8%82%AA%E8%82%9D%E5%90%88%E5%B9%B6%E6%85%A2%E4%B9%99%E8%82%9D%29"
)



func StartSpider(){

	c := colly.NewCollector(
		colly.AllowedDomains("s.wanfangdata.com.cn"),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)


	// On every a element which has href attribute call callback
	c.OnHTML(".right-list", func(e *colly.HTMLElement) {

		html := e.Response.Body
		fmt.Println(string(html))
		//link := e.Attr("href")
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		//c.Visit(e.Request.AbsoluteURL(link))
	})


	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(response *colly.Response) {
		html := string(response.Body)
		fmt.Println("html:",html)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(StartUrl)
}