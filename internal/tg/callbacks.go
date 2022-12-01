package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	UpdateCallback          = "update_content"
	UpdateDeadlineCallback  = "update_deadline"
	UpdateFrequencyCallback = "update_frequency"
	SaveCallback            = "save_reminder"
	InfoCallback            = "info_reminder"
	CancelCallback          = "cancel_reminder"
)

func (r *ReminderBot) save(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.CallbackQuery.Message.Chat.ID,
		"OK, I'll remind you.",
	)
	return r.bot.Send(msg)
}

func (r *ReminderBot) info(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.CallbackQuery.Message.Chat.ID,
		"Here's some info about current reminder.",
	)
	return r.bot.Send(msg)
}

func (r *ReminderBot) cancel(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.CallbackQuery.Message.Chat.ID,
		"OK, I won't remind you.",
	)
	return r.bot.Send(msg)
}

func (r *ReminderBot) updateContent(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.CallbackQuery.Message.Chat.ID,
		"OK, what should I remind you about?",
	)
	return r.bot.Send(msg)
}

func (r *ReminderBot) updateDeadline(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.CallbackQuery.Message.Chat.ID,
		"OK, when should I remind you?",
	)
	return r.bot.Send(msg)
}

func (r *ReminderBot) updateFrequency(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.CallbackQuery.Message.Chat.ID,
		"OK, how often should I remind you?",
	)
	return r.bot.Send(msg)
}
