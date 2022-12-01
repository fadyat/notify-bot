package tg

import (
	"deadlines/internal/db/cache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ReminderBot struct {
	bot   *tgbotapi.BotAPI
	cache cache.Cache
	tgHandler
}

func NewReminderBot(b *tgbotapi.BotAPI, c cache.Cache) *ReminderBot {
	return &ReminderBot{bot: b, cache: c}
}

type tgHandler interface {
	commands
	callbacks
	HandleUpdates(u *tgbotapi.Update)
}

type commands interface {
	getReminders(u *tgbotapi.Update) (tgbotapi.Message, error)
	newReminder(u *tgbotapi.Update) (tgbotapi.Message, error)
	ping(u *tgbotapi.Update) (tgbotapi.Message, error)
	isCommand(u *tgbotapi.Update) bool
}

type callbacks interface {
	updateContent(u *tgbotapi.Update) (tgbotapi.Message, error)
	updateDeadline(u *tgbotapi.Update) (tgbotapi.Message, error)
	updateFrequency(u *tgbotapi.Update) (tgbotapi.Message, error)
	save(u *tgbotapi.Update) (tgbotapi.Message, error)
	info(u *tgbotapi.Update) (tgbotapi.Message, error)
	cancel(u *tgbotapi.Update) (tgbotapi.Message, error)
	isCallback(u *tgbotapi.Update) bool
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
		}
	}

	if r.isCallback(u) {
		switch u.CallbackQuery.Data {
		case UpdateCallback:
			_, _ = r.updateContent(u)
		case UpdateDeadlineCallback:
			_, _ = r.updateDeadline(u)
		case UpdateFrequencyCallback:
			_, _ = r.updateFrequency(u)
		case SaveCallback:
			_, _ = r.save(u)
		case InfoCallback:
			_, _ = r.info(u)
		case CancelCallback:
			_, _ = r.cancel(u)
		}
		return
	}

}

func (r *ReminderBot) isCommand(u *tgbotapi.Update) bool {
	return u.Message != nil && u.Message.IsCommand()
}

func (r *ReminderBot) isCallback(u *tgbotapi.Update) bool {
	return u.CallbackQuery != nil
}
