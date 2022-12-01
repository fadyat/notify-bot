package services

import (
	"deadlines/internal/models/dto"
	"deadlines/internal/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type ReminderBot struct {
	bot *tgbotapi.BotAPI
	tgHandler
	repo repo.IReminderRepo
}

func NewReminderBot(b *tgbotapi.BotAPI, r repo.IReminderRepo) *ReminderBot {
	return &ReminderBot{bot: b, repo: r}
}

const (
	PingCmd         = "ping"
	NewReminderCmd  = "new"
	GetRemindersCmd = "get_all"
	DeleteCmd       = "delete"
	UpdateCmd       = "update"
	DescriptionCmd  = "description"
)

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

func (r *ReminderBot) ping(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "pong")
	return r.bot.Send(msg)
}

func (r *ReminderBot) getDescription(u *tgbotapi.Update) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(
		u.Message.Chat.ID,
		`
<b>Hello, I'm a reminder bot!</b>
I can remind you about something you want to do.

<b>Message for creating a new reminder:</b>
<code>/new "name" date time frequency</code>

<b>Message for updating a reminder:</b>
<code>/update reminder_id "content" date time frequency</code>

<b>Message for getting all your reminders:</b>
<code>/get_all</code>

<b>Message for deleting a reminder:</b>
<code>/delete reminder_id</code>

<b>Date can be provided in the following formats:</b>
- <i>${number_of_days}</i>

<b>Frequency can be provided in the following formats:</b>
- <i>once</i>
- <i>daily</i>
- <i>weekly</i>
- <i>monthly</i>

<b>Example:</b>
<code>/new "Solve some Leetcode problems" 01-01 12:00 every_day</code>
`)
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
