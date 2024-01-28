package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("no .env file")
	}
	value, err1 := os.LookupEnv(key)
	if err1 == false {
		log.Fatalf("No var %s in .env file", key)
	}
	return value
}
