package util

import (
	"log"
	"os"
	. "strconv"
)

type EnvConfig struct {
	Port int
	ImageUploadLimit int
}

var Config EnvConfig

func EnvGetNumber(key string, required bool) int {
	valueRaw := os.Getenv(key)
	if required && len(valueRaw) == 0 {
		log.Fatal("must provide 'PORT'")
	}

	result, err := Atoi(valueRaw)
	if err != nil {
		log.Fatal("port must be a number")
	}

	return result
}