package main

import (
	"deadlines/internal/db/cache"
	"deadlines/internal/tg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	token := ""
	bot, err := tgbotapi.NewBotAPI(token)
	bot.Debug = true
	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	redis := cache.NewRedisClient("localhost:6379")
	tgService := tg.NewReminderBot(bot, redis)
	for update := range updates {
		tgService.HandleUpdates(&update)
	}
}
