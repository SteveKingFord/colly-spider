package main

import (
	"colly-spider/pubmed"
	"colly-spider/utils"
	"fmt"
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
	//book := &Book{Id: 1, Title: "Go Web 编程", Summary: "Go Web 编程入门指南", Author: "学院君"}
	//
	//data, _ := json.Marshal(book)
	//testWrite(data)
	//print(string(data))


	pubmed.SpiderPubmed()

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


func HtmlToPdf(){
	pdf :=utils.NewPdf()
	//url,err:=pdf.OutFile("http://www.baidu.com","./test.pdf")
	// ./pubmed/articles
	url,err:=pdf.OutFile("https://pubmed.ncbi.nlm.nih.gov/26460381","./26460381.pdf")
	if err != nil{
		log.Println(err)
	}
	fmt.Println(url)
}