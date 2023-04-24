package router

import (
	"fmt"
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

	urlPrefix := Config.Hostname

	if Config.Hostname == "localhost" || len(Config.Hostname) == 0 {
		urlPrefix = fmt.Sprintf("http://localhost:%s/", Config.Port)
	} else if !strings.HasSuffix(urlPrefix, "/") {
		urlPrefix += "/"
	}

	body := gin.H{
		"file_url": urlPrefix + file.Filename,
	}

	c.JSON(http.StatusOK, body)
}

func downloadEndpoint(c *gin.Context) {
}