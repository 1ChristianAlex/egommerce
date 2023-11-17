package app

import (
	"log"

	pathresolver "khrix/egommerce/internal/libs/path_resolver"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	enfFile := pathresolver.GetCurrentPath("config/environment/.env")

	err := godotenv.Load(enfFile)
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}
