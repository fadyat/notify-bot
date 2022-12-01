package services

import (
	"deadlines/internal/repo"
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

const (
	checkInterval = 5 * time.Second
)

type ReminderCron struct {
	bot  *tgbotapi.BotAPI
	s    *gocron.Scheduler
	repo repo.IReminderRepoCron
}

func NewReminderCron(b *tgbotapi.BotAPI, r repo.IReminderRepoCron) *ReminderCron {
	return &ReminderCron{s: gocron.NewScheduler(time.UTC), repo: r, bot: b}
}

func (r *ReminderCron) Start() {
	r.s.StartAsync()
	r.initJobs()
}

func (r *ReminderCron) Stop() {
	r.s.Stop()
}

func (r *ReminderCron) initJobs() {
	j, err := r.s.Every(checkInterval).Do(r.notify)
	if err != nil {
		panic(err)
	}

	j.SingletonMode()
}

func (r *ReminderCron) notify() {
	log.Printf("Checking for reminders")
	reminders, err := r.repo.GetRemindersToRun()
	if err != nil {
		log.Printf("failed to get reminders to run: %v", err)
	}

	var sent = make([]int64, 0)
	for _, reminder := range reminders {
		msg := tgbotapi.NewMessage(999848817, reminder.NotifyString())
		_, err = r.bot.Send(msg)
		if err != nil {
			log.Printf("failed to send reminder: %v", err)
			continue
		}

		sent = append(sent, reminder.ID)
	}
	if len(sent) > 0 {
		err = r.repo.ChangeUpdatedAtMultiple(sent)
		if err != nil {
			log.Printf("failed to change updated_at: %v", err)
			return
		}

	}

	log.Printf("Sent %d reminders", len(sent))
}
