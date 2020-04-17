package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// get dot env variable
func gDEV(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
