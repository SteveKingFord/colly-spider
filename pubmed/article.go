package pubmed

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
	"time"
)

type Article struct {
	headTitle string
	title string
	affiliations []string
}

/**
 * 处理详情页
 */
func SpiderArticle(collector *colly.Collector, url string) {
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

		dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))

		if err!=nil{
			log.Fatalln(err)
		}

		title:=dom.Find("title").Text()
		filePath := "./pubmed/articles/" + title + ".html"

		fmt.Println(title)
		fmt.Println(filePath)
		fmt.Println(html)


		writeToHtmlFile(filePath,html)
	})

	// 解析详情页数据
	collector.OnHTML(".article-details", func(e *colly.HTMLElement) {
      headTitle := e.ChildText("h1.heading-title")
      title := e.ChildText(".expanded-authors > .affiliations > .title")
		fmt.Println(headTitle,title)
      var affiliations []string

     e.DOM.Find(".expanded-authors > .affiliations > .item-list li").Each(func(i int, item *goquery.Selection) {
		 affiliations =  append(affiliations, 	item.Text())
	  })

      article := Article{
		 headTitle:headTitle,
		 title:title,
		 affiliations: affiliations,
	  }

		jsonData, err := json.Marshal(article)

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(jsonData))

	})

	collector.Visit(url)
}

func writeToHtmlFile(filePath string, str string )  {

	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString(str)

}