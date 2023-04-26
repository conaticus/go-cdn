package router

import (
	"cdn/api/database"
	. "cdn/api/util"
	"crypto/md5"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func uploadEndpoint(c *gin.Context) {
	const paramName = "file"

	fileHeader, err := c.FormFile(paramName)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid file upload: %s", err.Error())
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to open uploaded file: %s", err.Error())
		return
	}
	defer file.Close()

	// Validate file type based on file content
	fileBuffer := make([]byte, 512) // Read first 512 bytes of file content
	_, err = file.Read(fileBuffer)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read file: %s", err.Error())
		return
	}
	fileType := http.DetectContentType(fileBuffer)

	// Check if file type is allowed (image MIME types whitelist)
	allowedMIMETypes := map[string]bool{
		"image/jpg":  true,
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	if !allowedMIMETypes[fileType] {
		c.String(http.StatusBadRequest, "Invalid file upload: file type must be an image. Found '%s'.", fileType)
		return
	}

	// Upload the file if has not already been uploaded
	fileHashBuffer := md5.Sum(fileBuffer)
	fileName := uuid.NewString() + filepath.Ext(fileHeader.Filename)
	savedFileName, alreadyExists := database.AddImage(fileName, fileHashBuffer[:])
	if !alreadyExists {
		err = c.SaveUploadedFile(fileHeader, "./uploads/images/"+fileName)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to save uploaded file: %s", err.Error())
			return
		}
	}

	body := gin.H{
		"file_url": Config.HostUrl + "download/" + savedFileName,
	}
	c.JSON(http.StatusOK, body)
}