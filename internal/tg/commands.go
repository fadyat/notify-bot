package tg

import (
	"deadlines/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

const (
	PingCmd         = "ping"
	NewReminderCmd  = "new"
	GetRemindersCmd = "get_all"
)

func (r *ReminderBot) ping(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.Message.Chat.ID,
		"pong",
	)
	return r.bot.Send(msg)
}

func (r *ReminderBot) newReminder(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.Message.Chat.ID,
		"This is a new form of reminder.\n\nClick on the buttons below to fill it in.",
	)
	msg.ReplyMarkup = NewReminderMarkup

	userID := strconv.FormatInt(u.Message.From.ID, 10)
	have, err := r.cache.ContainsConfiguredReminder(userID)
	if err != nil {
		return tgbotapi.Message{}, err
	}

	if !have {
		err = r.cache.SetConfiguredReminder(userID, &models.Reminder{})
		if err != nil {
			return tgbotapi.Message{}, err
		}
	}

	return r.bot.Send(msg)
}

func (r *ReminderBot) getReminders(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.Message.Chat.ID,
		"Here's your reminders:",
	)
	return r.bot.Send(msg)
}
