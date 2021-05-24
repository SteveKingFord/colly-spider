package api

import (
	"colly-spider/global"
	"colly-spider/model/Fibrosis"
	"colly-spider/repository"
	"fmt"
	"github.com/gin-gonic/gin"
)


type FibrosisQuery struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func (q *FibrosisQuery) GetList()([]Fibrosis.FibrosisArticle, error){
	r:=repository.FibrosisArticleRepository{
		DB: global.DB,
	}
	return r.GetList(q.PageIndex,q.PageSize)
}

func (q *FibrosisQuery) GetTotal()(*int64,error){
	r:=repository.FibrosisArticleRepository{
		DB: global.DB,
	}
	return r.Total()
}


func GetFibrosisList(c *gin.Context)  {
	q := FibrosisQuery{}
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