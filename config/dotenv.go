package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetDotEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("not found env value")
	}

	return os.Getenv(key)
}
