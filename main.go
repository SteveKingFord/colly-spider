package main

import (
	"colly-spider/global"
	"colly-spider/initiallize"
	"colly-spider/model/common"
	"colly-spider/pubmed"
	"log"
	"os"
)


type Book struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Author string `json:"author"`
}

func main()  {
	global.DB = initiallize.InitialGORM()
    if global.DB  != nil {
		err := common.InitMigrate(global.DB)
		if err != nil {
			log.Fatal("database init err:",err)
		}
		pubmed.SpiderPubmed()
		//initiallize.RegisterRouter()
	}

}




func testWrite(data []byte) {
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}


