package main

import (
	"log"
	"os"

	"github.com/SHIKHAR-Bajpai/bot/bot"
	"github.com/SHIKHAR-Bajpai/bot/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.ReadConfig()
	bot.Start()

	<-make(chan struct{})

}
