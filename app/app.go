package app

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/ping", PingHandler)
	router.GET("/file/:filename", DownloadHandler)
	router.POST("/file/up", UploadHandler)

	router.Run("localhost:8080")
}
