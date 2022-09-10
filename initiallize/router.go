package initiallize

import (
	"colly-spider/api"
	"colly-spider/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() {
	r := gin.Default()

	r.Use(middleware.Cors2())

	v1 := r.Group("/api")

	{
		v1.GET("/pubmed/fatty-liver", api.GetPubmedList)
		v1.GET("/pubmed/fibrosis", api.GetFibrosisList)
	}

	_ = r.Run(":8080")

}
