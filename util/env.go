package util

import (
	"log"
	"os"
	. "strconv"

	env "github.com/joho/godotenv"
)

// Modes
const (
	ModeDebug = "debug"
	ModeRelease = "release"
)

type EnvConfig struct {
	Port int
	ImageUploadLimit int
	Mode string
}

var Config EnvConfig

func init() {
	env.Load()
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