package main

import (
	"log"

	"github.com/skingford/colly-spider/global"
	"github.com/skingford/colly-spider/initiallize"
	"github.com/skingford/colly-spider/model/common"
)

type Book struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  string `json:"author"`
}

func main() {
	global.DB = initiallize.InitialGORM()
	if global.DB != nil {
		err := common.InitMigrate(global.DB)
		if err != nil {
			log.Fatal("database init err:", err)
		}
		//pubmed.SpiderPubmed()
		// pubmed.SpiderPubmed()
		initiallize.RegisterRouter()
	}
}
