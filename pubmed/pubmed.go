package pubmed

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
)

const (
	DOMAIN   = "https://pubmed.ncbi.nlm.nih.gov"
	StartUrl = DOMAIN + "/?term=Fatty+liver+with+chronic+hepatitis+B"
)

var page int = 1

type Data struct {
	Term                string `json:"term" form:"term"`
	Page                int    `json:"page" form:"page"`
	NoCache             string `json:"no-cache" form:"no-cache"`
	Csrfmiddlewaretoken string `json:"csrfmiddlewaretoken" form:"csrfmiddlewaretoken"`
}

func SpiderPubmed() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: pubmed.ncbi.nlm.nih.gov
		colly.AllowedDomains("pubmed.ncbi.nlm.nih.gov"),
		colly.CacheDir("./pubmed/coursera_cache"),
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
		fmt.Println("c 请求地址", r.URL.String())
		//r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	})

	c.OnResponse(func(response *colly.Response) {
		//html := string(response.Body)
	})

	c.OnScraped(func(r *colly.Response) {
		page = page + 1
		if page*10 <= 12186 {
			nextPage := fmt.Sprintf("%s&page=%d", StartUrl, page)
			fmt.Println("next page is", nextPage)
			c.Visit(nextPage)
		}

		fmt.Println("Finished", r.Request.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(StartUrl)
}

func nextPage(c *colly.Collector) {
	nextUrl := "https://pubmed.ncbi.nlm.nih.gov/more/"

	param := &Data{
		Term:                "intima-media thickness、Atherosclerotic Cardiovascular Disease",
		Page:                2,
		NoCache:             "1618148661073",
		Csrfmiddlewaretoken: "ShCyEgOye3MxKB8SCSTBb0EKnNs3sZQsrOQozXOr0PdZnmLVCX5h6BvTAOqKEcP9",
	}

	formData, _ := json.Marshal(param)
	c.PostRaw(nextUrl, formData)
	c.Visit(nextUrl)
}
