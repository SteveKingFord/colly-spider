package pubmed

import (
	"colly-spider/global"
	"colly-spider/model/FattyLiver"
	"colly-spider/repository"
	"colly-spider/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"strings"
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
			authors []FattyLiver.FattyLiverAuthor
			abstracts []FattyLiver.FattyLiverAbstract
		)

		eubDate :=  e.ChildText(".full-view .secondary-date")

		if eubDate == "" {
			cit := e.ChildText(".full-view .cit")
			cits := strings.Split(cit,";")
			eubDate = cits[0]
		}


		title := e.ChildText(".full-view h1.heading-title")

		e.DOM.Find(".full-view .authors .authors-list .authors-list-item").Each(func(i int, item *goquery.Selection){
			author := item.Find("a.full-name").Text()
			authors = append(authors, FattyLiver.FattyLiverAuthor{Name: author})
		})

		fmt.Println("authors:",authors)

		e.DOM.Find(".abstract-content p").Each(func(i int, item *goquery.Selection) {

			strong := item.Find("strong.sub-title").Text()
			strong = utils.CleanSpaceLine(strong)
			text := item.Text()
			text = utils.DeleteExtraSpace(text)
			text = utils.CleanLine(text)

			abstracts = append(abstracts, FattyLiver.FattyLiverAbstract{
				Content: text,
				Strong:  strong,
			})
		})

		article := &FattyLiver.FattyLiverArticle{
			Type:      "超声-肝合并慢乙肝",
			Href:      href,
			Title:     title,
			EupDate:  eubDate,
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
