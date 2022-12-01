package tg

import (
	"deadlines/internal/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ReminderBot struct {
	bot *tgbotapi.BotAPI
	tgHandler
	repo repo.IReminderRepo
}

func NewReminderBot(b *tgbotapi.BotAPI, r repo.IReminderRepo) *ReminderBot {
	return &ReminderBot{bot: b, repo: r}
}

type tgHandler interface {
	commands
	HandleUpdates(u *tgbotapi.Update)
}

type commands interface {
	getReminders(u *tgbotapi.Update) (tgbotapi.Message, error)
	newReminder(u *tgbotapi.Update) (tgbotapi.Message, error)
	getDescription(u *tgbotapi.Update) (tgbotapi.Message, error)
	ping(u *tgbotapi.Update) (tgbotapi.Message, error)
	isCommand(u *tgbotapi.Update) bool
}

func (r *ReminderBot) HandleUpdates(u *tgbotapi.Update) {
	if r.isCommand(u) {
		switch u.Message.Command() {
		case PingCmd:
			_, _ = r.ping(u)
		case NewReminderCmd:
			_, _ = r.newReminder(u)
		case GetRemindersCmd:
			_, _ = r.getReminders(u)
		case DescriptionCmd:
			_, _ = r.getDescription(u)
		}
	}

}

func (r *ReminderBot) isCommand(u *tgbotapi.Update) bool {
	return u.Message != nil && u.Message.IsCommand()
}
