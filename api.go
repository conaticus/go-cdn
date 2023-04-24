package main

import (
	"cdn/api/router"
	. "cdn/api/util"
)

func init() {
	Config = EnvConfig{Port: EnvGetNumber("PORT", true), ImageUploadLimit: EnvGetNumber("IMAGE_UPLOAD_LIMIT_MB", true)}
}

func main() {
	router.InitRoutes()
}