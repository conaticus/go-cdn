package main

import (
	"cdn/api/router"
	. "cdn/api/util"

	"github.com/gin-gonic/gin"
)

func init() {
	Config = EnvConfig{
		Port: EnvGetNumber("PORT", true),
		ImageUploadLimit: EnvGetNumber("IMAGE_UPLOAD_LIMIT_MB", true),
		Mode: EnvGetString("MODE", false),
	}
	
	gin.SetMode(Config.Mode)
}

func main() {
	router.InitRoutes()
}