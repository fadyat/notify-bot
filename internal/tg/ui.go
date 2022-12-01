package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	UpdateContentData   = "update_content"
	UpdateDeadlineDate  = "update_deadline"
	UpdateFrequencyData = "update_frequency"
	CancelData          = "cancel"
)

var NewReminderMarkup = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Content", UpdateContentData),
		tgbotapi.NewInlineKeyboardButtonData("Deadline", UpdateDeadlineDate),
		tgbotapi.NewInlineKeyboardButtonData("Frequency", UpdateFrequencyData),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", CancelData),
	),
)
