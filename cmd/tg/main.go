package main

import (
	"deadlines/internal/tg/commands"
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func main() {
	token := ""
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	cron := gocron.NewScheduler(time.UTC)
	_, err = cron.Every(1).Second().Do(func() {
		log.Println("tick")
	})
	cron.StartAsync()
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if isCommand(bot, &update) {
			continue
		}

		if isCallback(bot, &update) {
			continue
		}

		if isSticker(bot, &update) {
			continue
		}
	}
}

func isCommand(bot *tgbotapi.BotAPI, u *tgbotapi.Update) bool {
	if u.Message == nil || !u.Message.IsCommand() {
		return false
	}

	msg := commands.Resolve(u)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}

	return true
}

func isCallback(bot *tgbotapi.BotAPI, u *tgbotapi.Update) bool {
	if u.CallbackQuery == nil {
		return false
	}

	msg := tgbotapi.NewMessage(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Data)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}

	return true
}

func isSticker(bot *tgbotapi.BotAPI, u *tgbotapi.Update) bool {
	if u.Message == nil || u.Message.Sticker == nil {
		return false
	}

	sticker := tgbotapi.FileID(u.Message.Sticker.FileID)
	msg := tgbotapi.NewSticker(u.Message.Chat.ID, sticker)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}

	return true
}
