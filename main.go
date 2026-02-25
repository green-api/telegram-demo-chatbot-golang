package main

import (
	"log"
	"os"

	greenapi "github.com/green-api/telegram-api-client-golang"

	chatbot "github.com/green-api/telegram-chatbot-golang"
	"github.com/green-api/telegram-demo-chatbot-golang/scenes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	idInstance := os.Getenv("ID_INSTANCE")
	authToken := os.Getenv("AUTH_TOKEN")

	if idInstance == "" || authToken == "" {
		log.Fatal("ID_INSTANCE and AUTH_TOKEN must be set in .env file or environment variables.")
	}

	baseBot := chatbot.NewBot(idInstance, authToken)

	go func() {
		for err := range baseBot.ErrorChannel {
			if err != nil {
				log.Printf("ERROR: %v", err)
			}
		}
	}()

	_, err = baseBot.Account().SetSettings(
		greenapi.OptionalIncomingWebhook(true),
		greenapi.OptionalOutgoingWebhook(false),
		greenapi.OptionalStateWebhook(false),
		greenapi.OptionalOutgoingAPIMessageWebhook(false),
		greenapi.OptionalOutgoingMessageWebhook(false),
		greenapi.OptionalMarkIncomingMessagesRead(true),
	)
	if err != nil {
		log.Fatalf("Failed to set Green API settings: %v", err)
	}

	baseBot.SetStartScene(scenes.StartScene{})

	log.Println("Starting Green API Demo Bot...")
	baseBot.StartReceivingNotifications()

	log.Println("Bot stopped.")
}
