package initialize

import (
	"log"

	"github.com/skingford/colly-spider/api"
	"github.com/skingford/colly-spider/middleware"

	"github.com/gin-gonic/gin"
)

func RunApp() {
	// 设置 release模式
	// gin.SetMode(gin.ReleaseMode)

	// 或者 设置debug模式
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.Use(middleware.Cors2())

	v1 := r.Group("/v1")

	{
		v1.GET("/pubmed/articles", api.GetPubmedList)
	}

	log.Fatal(r.Run("0.0.0.0:80"))
}
