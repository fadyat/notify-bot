package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var NewReminderMarkup = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Content", UpdateCallback),
		tgbotapi.NewInlineKeyboardButtonData("Deadline", UpdateDeadlineCallback),
		tgbotapi.NewInlineKeyboardButtonData("Frequency", UpdateFrequencyCallback),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("✅", SaveCallback),
		tgbotapi.NewInlineKeyboardButtonData("ℹ️", InfoCallback),
		tgbotapi.NewInlineKeyboardButtonData("❌", CancelCallback),
	),
)
