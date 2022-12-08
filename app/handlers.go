package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gujjsPet/image-transform-service/file"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func DownloadHandler(ctx *gin.Context) {
	fileName := ctx.Param("filename")
	targetPath, err := file.GetFilePath(ctx.Param("filename"))
	if err != nil {
		ctx.String(403, err.Error())
		return
	}
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(targetPath)
}
