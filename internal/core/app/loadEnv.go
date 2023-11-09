package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	base, _ := os.Getwd()
	fmt.Println(base)

	localPath := filepath.Join(base, "/../../", "config/environment/.env")

	if strings.HasSuffix(base, "egommerce") {
		localPath = filepath.Join(base, "config/environment/.env")
	}

	err := godotenv.Load(localPath)
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}
