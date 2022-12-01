package tg

import (
	"deadlines/internal/models/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	PingCmd         = "ping"
	NewReminderCmd  = "new"
	GetRemindersCmd = "get_all"
	DeleteCmd       = "delete"
	UpdateCmd       = "update"
	DescriptionCmd  = "description"
)

func (r *ReminderBot) ping(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "pong")
	return r.bot.Send(msg)
}

func (r *ReminderBot) getDescription(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, botDescription)
	msg.ParseMode = "html"

	return r.bot.Send(msg)
}

func (r *ReminderBot) newReminder(u *tgbotapi.Update) (tgbotapi.Message, error) {
	reminder, err := dto.ReminderFromTgArgs(u.Message.CommandArguments(), u.Message.From.ID)
	msg := tgbotapi.NewMessage(
		u.Message.Chat.ID,
		"Sorry, I can't parse your reminder. Please, try again.",
	)

	if err == nil {
		err = r.repo.SaveReminder(reminder)
		if err == nil {
			msg = tgbotapi.NewMessage(u.Message.Chat.ID, reminder.ToString())
		} else {
			msg.Text = "Sorry, I can't save your reminder. Please, try again."
		}
	}

	return r.bot.Send(msg)
}

func (r *ReminderBot) getReminders(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Here's your reminders:\n\n")
	reminders, err := r.repo.GetReminders(u.Message.From.ID)
	if err != nil {
		log.Printf("error while getting reminders: %v", err)
		msg.Text = "Sorry, I can't get your reminders. Please, try again."
		return r.bot.Send(msg)
	}

	for _, reminder := range reminders {
		msg.Text += reminder.ToString() + "\n"
	}

	return r.bot.Send(msg)
}
