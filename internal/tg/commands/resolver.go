package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Resolve(u *tgbotapi.Update) tgbotapi.MessageConfig {
	switch u.Message.Command() {
	case "ping":
		return ping(u.Message.Chat.ID)
	case "new":
		return newReminder(u)
	default:
		return tgbotapi.NewMessage(u.Message.Chat.ID, "outdated command")
	}
}
