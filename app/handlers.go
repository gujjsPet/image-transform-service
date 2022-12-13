package app

import (
	"net/http"
	"log"
	"fmt"
	"strings"

	"github.com/google/uuid"
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

	if err != nil {		
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "error receiving file",
        })

	}

	fileNameSlice := strings.Split(file.Filename, ".")
	fileExtension := fileNameSlice[len(fileNameSlice) - 1]	
	newFileName := uuid.New().String()

	err = c.SaveUploadedFile(file, "storage/" + newFileName + "." + fileExtension)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "message": "error saving file",
        })
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded", file.Filename),
		"accessName": newFileName + "." + fileExtension,
	})
}
