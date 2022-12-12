package app

import (
	"net/http"
	"log"
	"fmt"

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

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println(c.Request.ParseForm())
	if err != nil {		
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "No file is received",
        })

	}

	err = c.SaveUploadedFile(file, "storage/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
