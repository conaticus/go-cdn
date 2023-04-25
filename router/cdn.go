package router

import (
	"crypto/md5"
	"io/ioutil"
	"net/http"
	"strings"

	. "cdn/api/util"

	"cdn/api/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func uploadEndpoint(c *gin.Context) {
	const paramName = "file"

	fileHeader, err := c.FormFile(paramName)
	if err != nil {
		c.String(http.StatusForbidden, "Invalid file upload, parameter must be named '%s'.", paramName)
		return
	}

	fileType := fileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(fileType, "image/") {
		c.String(http.StatusForbidden, "Invalid file upload, file type must be an image. Found '%s'.", fileType)
	}

	file, _ := fileHeader.Open()
	defer file.Close()

	fileBuffer, _ := ioutil.ReadAll(file)
	fileHashBuffer := md5.Sum(fileBuffer)

	fileName := uuid.NewString() + "." + GetFileExtension(fileHeader.Filename)

	savedFileName, alreadyExists := database.AddImage(fileName, fileHashBuffer[:])
	if !alreadyExists {
		c.SaveUploadedFile(fileHeader, "./uploads/images/" + fileName)
	}

	body := gin.H{
		"file_url": Config.HostUrl + "download/" + savedFileName,
	}

	c.JSON(http.StatusOK, body)
}