package app

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gujjsPet/image-transform-service/file"
	"github.com/gujjsPet/image-transform-service/image"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func setFileUploadHeaders(c *gin.Context, fileName string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
}

func DownloadHandler(c *gin.Context) {
	fileName := c.Param("filename")
	targetPath, err := file.GetFilePath(c.Param("filename"))
	if err != nil {
		c.String(403, err.Error())
		return
	}

	setFileUploadHeaders(c, fileName)
	c.File(targetPath)
}

func UploadHandler(c *gin.Context) {
	f, err := c.FormFile("file")

	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "error receiving file",
		})

	}

	newFileName := file.GenerateFilename(f.Filename)

	err = c.SaveUploadedFile(f, "storage/"+newFileName)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error saving file",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    fmt.Sprintf("'%s' uploaded", newFileName),
		"accessName": newFileName,
	})
}

func decodeCoordinates(v url.Values) (int, int, int, int) {
	x0, _ := strconv.Atoi(v.Get("x0"))
	y0, _ := strconv.Atoi(v.Get("y0"))
	x1, _ := strconv.Atoi(v.Get("x1"))
	y1, _ := strconv.Atoi(v.Get("y1"))
	return x0, y0, x1, y1
}

func CropHandler(c *gin.Context) {
	fileName := c.Param("filename")
	srcFilePath, err := file.GetFilePath(fileName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	src, err := os.Open(srcFilePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("couldn't open %s source file", fileName),
		})
		return
	}
	defer src.Close()

	base, err := image.DecodeImage(src)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error decoding the source file",
		})
		return
	}

	var b bytes.Buffer
	x0, y0, x1, y1 := decodeCoordinates(c.Request.URL.Query())
	i := image.CropImage(base, x0, y0, x1, y1)
	err = image.EncodeImage(&b, i)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error encoding the result image",
		})
		return
	}

	setFileUploadHeaders(c, fileName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}
