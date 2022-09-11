package main

import (
	"colly-spider/global"
	"colly-spider/initiallize"
	"colly-spider/model/common"
	"log"
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
