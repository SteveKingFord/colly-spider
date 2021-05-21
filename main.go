package main

import (
	"colly-spider/wanfangdata"
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

	wanfangdata.StartSpider()
	//global.DB = initiallize.InitialGORM()
    //if global.DB  != nil {
	//	common.InitMigrate(global.DB)
	//	wanfangdata.StartSpider()
	//	initiallize.RegisterRouter()
	//}

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


