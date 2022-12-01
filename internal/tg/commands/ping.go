package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func ping(chatID int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, "pong")
}
