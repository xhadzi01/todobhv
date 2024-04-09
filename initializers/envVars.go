package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()

	if err != nil {
		errMsg := "Error loading .env file"
		log.Fatal(errMsg)
		panic(errMsg)
	}
}
