package initiallize

import (
	"colly-spider/api"
	"colly-spider/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter()  {
	r := gin.Default()

	r.Use(middleware.Cors2())

	v1 := r.Group("/api")

	{
		v1.GET("/article", api.GetList)
	}

	r.Run(":8080")

}