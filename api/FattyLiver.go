package api

import (
	"colly-spider/global"
	"colly-spider/model/FattyLiver"
	"colly-spider/repository"
	"fmt"
	"github.com/gin-gonic/gin"
)


type Query struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func (q *Query) GetList()([]FattyLiver.FattyLiverArticle, error){
	r:=repository.FattyLiverArticleRepository{
		DB: global.DB,
	}
	return r.GetList(q.PageIndex,q.PageSize)
}

func (q *Query) GetTotal()(*int64,error){
	r:=repository.FattyLiverArticleRepository{
		DB: global.DB,
	}
	return r.Total()
}


func GetFattyLiverList(c *gin.Context)  {
	fmt.Println(c.Query("pageIndex"))
	fmt.Println(c.Query("pageSize"))
	 q := Query{}
	if err :=c.ShouldBind(&q);err ==nil {
		value ,err := q.GetList()
		total,_:=q.GetTotal()

		if err != nil {
			_ = fmt.Errorf("failed a something %s", err)
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200,gin.H{
			"data":value,
			"total":total,
		})
	}else {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	}


}