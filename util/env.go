package util

import (
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
	Hostname string
}

var Config EnvConfig

func init() {
	env.Load()
	Config = EnvConfig{
		Port: EnvGetString("PORT", true),
		FileUploadLimit: EnvGetNumber("FILE_UPLOAD_LIMIT_MB", true),

		Mode: EnvGetString("MODE", false),
		Hostname: EnvGetString("HOSTNAME", false),
	}

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