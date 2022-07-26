package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	corsMiddlewareConfig := cors.DefaultConfig()
	corsMiddlewareConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "access_token"}
	corsMiddlewareConfig.AllowAllOrigins = true
	router.Use(cors.New(corsMiddlewareConfig))

	mapUrls()
	err := router.Run(":9002")
	if err != nil {
		return
	}
}
