package app

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/ping", PingHandler)
	router.GET("/file/:filename", DownloadHandler)

	router.Run("localhost:8080")
}
