package util

import (
	"fmt"
	"log"
	"os"
	. "strconv"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
)

// Modes
const (
	ModeDebug = "debug"
	ModeRelease = "release"
)

type EnvConfig struct {
	Port string
	FileUploadLimit int
	Mode string
	HostUrl string
}

var Config EnvConfig

func GetUrlFromHostname() (string) {
	hostname := EnvGetString("HOSTNAME", false)
	var hostUrl string

	if hostname == "localhost" || len(hostname) == 0 {
		hostUrl = fmt.Sprintf("http://localhost:%s/", Config.Port)
	} else {
		hostUrl = "https://" + hostname + "/"
	}

	return hostUrl
}

func init() {
	env.Load()
	Config = EnvConfig{
		Port: EnvGetString("PORT", true),
		FileUploadLimit: EnvGetNumber("FILE_UPLOAD_LIMIT_MB", true),

		Mode: EnvGetString("MODE", false),
	}

	Config.HostUrl = GetUrlFromHostname()

	gin.SetMode(Config.Mode)
}

// Errors if does not exist
func checkExists(key string, value string) {
	if len(value) == 0 {
		log.Fatalf("must provide '%s'", key)
	}
}

func EnvGetNumber(key string, required bool) int {
	valueRaw := os.Getenv(key)
	if required { checkExists(key, valueRaw) }

	result, err := Atoi(valueRaw)
	if err != nil {
		log.Fatal("port must be a number")
	}

	return result
}

func EnvGetString(key string, required bool) string {
	value := os.Getenv(key)
	if required { checkExists(key, value) }

	return value
}