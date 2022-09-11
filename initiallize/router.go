package initiallize

import (
	"colly-spider/api"
	"colly-spider/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() {
	r := gin.Default()

	r.Use(middleware.Cors2())

	v1 := r.Group("/v1")

	{
		v1.GET("/pubmed/articles", api.GetPubmedList)
	}

	log.Fatal(r.Run("0.0.0.0:8080"))
}
