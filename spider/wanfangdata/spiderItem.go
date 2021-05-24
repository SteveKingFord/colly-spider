package wanfangdata

import (
	"github.com/gocolly/colly"
)

func SpiderItem(c *colly.Collector,url string){
	iC:= c.Clone()
	iC.Visit(url)

	iC.OnHTML(`a[name]`, func(e *colly.HTMLElement) {

	})
}
