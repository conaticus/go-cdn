package util

import (
	"log"
	"os"
	. "strconv"

	env "github.com/joho/godotenv"
)

type EnvConfig struct {
	Port int
	ImageUploadLimit int
}

var Config EnvConfig

func init() {
	env.Load()
}

func EnvGetNumber(key string, required bool) int {
	valueRaw := os.Getenv(key)
	if required && len(valueRaw) == 0 {
		log.Fatalf("must provide '%s'", key)
	}

	result, err := Atoi(valueRaw)
	if err != nil {
		log.Fatal("port must be a number")
	}

	return result
}