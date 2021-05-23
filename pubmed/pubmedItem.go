package pubmed

import (
	"colly-spider/global"
	"colly-spider/model"
	"colly-spider/repository"
	"colly-spider/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"time"
)


func SpiderArticle(collector *colly.Collector, href string) {
	url := DOMAIN + href

	collector = collector.Clone()

	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 2 * time.Second,
		Parallelism: 10,
	})

	if err != nil {
	log.Fatal("limit error:",err)
	}

	collector.OnRequest(func(request *colly.Request) {
		log.Println("articles visit: ", request.URL.String())
	})

	collector.OnResponse(func(response *colly.Response) {

	})

	// 解析详情页数据
	collector.OnHTML(".article-details", func(e *colly.HTMLElement) {
		var (
			authors []model.FattyLiverAuthor
			abstracts []model.FattyLiverAbstract
		)

		title := e.ChildText("h1.heading-title")

		e.DOM.Find(".full-view .authors .authors-list .authors-list-item").Each(func(i int, item *goquery.Selection){
			author := item.Find("a.full-name").Text()
			authors = append(authors, model.FattyLiverAuthor{Name: author})
		})

		fmt.Println("authors:",authors)

		e.DOM.Find(".abstract-content p").Each(func(i int, item *goquery.Selection) {

			strong := item.Find("strong.sub-title").Text()
			strong = utils.CleanSpaceLine(strong)
			text := item.Text()
			text = utils.DeleteExtraSpace(text)
			text = utils.CleanLine(text)

			abstracts = append(abstracts, model.FattyLiverAbstract{
				Content: text,
				Strong:  strong,
			})
		})

		article := &model.FattyLiverArticle{
			Type:      "pubmed",
			Href:      href,
			Title:     title,
			FattyLiverAuthors:   authors,
			FattyLiverAbstracts: abstracts,
		}

		fmt.Println(article)
		r := repository.FattyLiverArticleRepository{DB: global.DB}
		err := r.Create(article)
		if err != nil {
			fmt.Println("create article failed!")
		}
	})

	err = collector.Visit(url)
	if err != nil {
		log.Fatal("Visit error:",err)
	}
}
