package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ProjectConfig Config

type Config struct {
	Port        string
	DatabaseUrl string
	DatabaseKey string
}

func init() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not defined")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable not defined")
	}

	databaseKey := os.Getenv("DATABASE_KEY")
	if databaseKey == "" {
		log.Fatal("DATABASE_KEY environment variable not defined")
	}

	ProjectConfig = Config{
		Port:        port,
		DatabaseUrl: databaseUrl,
		DatabaseKey: databaseKey,
	}
}
