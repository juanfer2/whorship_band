package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

func LoadEnv() {
	env := utilities.GetRootFolder() + "/.env"
	err := godotenv.Load(env)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
