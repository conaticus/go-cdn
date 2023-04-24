package router

import (
	"net/http"
	"strings"

	. "cdn/api/util"

	"github.com/gin-gonic/gin"
)

func uploadEndpoint(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusForbidden, "Invalid file upload, parameter must be named 'file'.")
		return
	}

	fileType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(fileType, "image/") {
		c.String(http.StatusForbidden, "Invalid file upload, file type must be an image. Found '%s'.", fileType)
	}

	c.SaveUploadedFile(file, "./uploads/images/" + file.Filename)

	body := gin.H{
		"file_url": Config.HostUrl + file.Filename,
	}

	c.JSON(http.StatusOK, body)
}

func downloadEndpoint(c *gin.Context) {
}