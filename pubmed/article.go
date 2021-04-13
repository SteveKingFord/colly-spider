package pubmed

import (
	"bufio"
	"colly-spider/global"
	"colly-spider/model"
	"colly-spider/repository"
	"colly-spider/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"os"
	"time"
)

type abstractContent struct {
	Content string `json:"content"`
	Strong  string `json:"strong"`
}

type Article struct {
	Title     string            `json:"title"`
	Authors   []string          `json:"authors"`
	Abstracts []abstractContent `json:"abstracts"`
}

/**
 * 处理详情页
 */
func SpiderArticle(collector *colly.Collector, href string) {
	url := DOMAIN + href

	collector = collector.Clone()

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
	})

	collector.OnRequest(func(request *colly.Request) {
		log.Println("articles visit: ", request.URL.String())
	})

	collector.OnResponse(func(response *colly.Response) {
		//html := string(response.Body)

		//dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
		//
		//if err!=nil{
		//	log.Fatalln(err)
		//}
		//
		//title:=dom.Find("title").Text()
		//filePath := "./pubmed/articles/" + title + ".html"
		//
		//fmt.Printf("文件将保存到：%s",filePath)

		//writeToFile(filePath,html)
	})

	// 解析详情页数据
	collector.OnHTML(".article-details", func(e *colly.HTMLElement) {
		var authors []model.Author
		var abstracts []model.Content

		title := e.ChildText("h1.heading-title")

		e.DOM.Find(".authors-list .authors-list-item ").Each(func(i int, item *goquery.Selection) {
			name := item.Find(".full-name").Text()
			authors = append(authors, model.Author{Name: name})
		})

		e.DOM.Find(".abstract-content p").Each(func(i int, item *goquery.Selection) {

			strong := item.Find("strong.sub-title").Text()
			strong = utils.CleanSpaceLine(strong)
			text := item.Text()
			text = utils.DeleteExtraSpace(text)
			text = utils.CleanLine(text)

			abstracts = append(abstracts, model.Content{
				Content: text,
				Strong:  strong,
			})
		})

		article := &model.Article{
			Type:      "pubmed",
			Href:      href,
			Title:     title,
			Authors:   authors,
			Abstracts: abstracts,
		}

		r := repository.ArticleRespository{DB: global.DB}
		err := r.Create(article)
		if err != nil {
			fmt.Println("create article failed!")
		}

		//article := &Article{
		//	Title:     utils.DeleteExtraSpace(title),
		//	Authors:   authors,
		//	Abstracts: abstracts,
		//}
		//
		//data, err := json.Marshal(article)
		//if err != nil {
		//	fmt.Println(err.Error())
		//}
		//
		//fmt.Println(string(data))
		//
		//dir := href[0 : len(href)-1]
		//
		//filePath := "./pubmed/articles" + dir + ".json"
		//
		//fmt.Printf("now file path is:%s", filePath)
		//writeToFile(filePath,[]byte(data))
		//utils.WriteFileWithBytes(filePath,data)

	})

	collector.Visit(url)
}

func writeToFile(filePath string, content string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString(content)

}

func testWrite(filePath string, data []byte) {

	fp, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	defer func() { fp.Close() }()

	if err != nil && os.IsExist(err) {
		fp, _ = os.Create(filePath)

	}

	_, err = fp.Write(data)

	if err != nil {
		log.Fatal(err)
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
