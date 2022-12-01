package commands

import (
	"deadlines/internal/tg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func newReminder(u *tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Setup reminder")
	msg.ReplyMarkup = tg.NewReminderMarkup
	return msg
}
