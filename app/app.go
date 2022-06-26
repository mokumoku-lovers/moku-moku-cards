package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	router.Use(cors.Default())
	mapUrls()
	err := router.Run(":9002")
	if err != nil {
		return
	}
}
