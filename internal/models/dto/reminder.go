package dto

import (
	"deadlines/internal/helpers"
	"deadlines/internal/models"
	"fmt"
	"regexp"
	"time"
)

type Reminder struct {
	ID              int64            `json:"id,omitempty"`
	UserID          int64            `json:"user_id"`
	Name            string           `json:"name"`
	Deadline        time.Time        `json:"deadline"`
	RemindFrequency models.Frequency `json:"remind_frequency"`
}

func (r *Reminder) ToString() string {
	return fmt.Sprintf(
		"Name: %s\nDeadline: %s\nFrequency: %s",
		r.Name,
		r.Deadline.Format("2006-01-02 15:04"),
		r.RemindFrequency,
	)
}

func ReminderFromTgArgs(args string, userID int64) (*Reminder, error) {
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
	parsed := r.FindAllString(args, -1)
	if len(parsed) < 2 || len(parsed) > 4 {
		return &Reminder{}, models.ErrInvalidArgs
	}

	curTime := time.Now().UTC()
	name, deadlineDate, frequency := parsed[0], curTime.AddDate(0, 0, 1), models.Once
	if name[0] == '"' && name[len(name)-1] == '"' {
		name = name[1 : len(name)-1]
	}

	days, err := helpers.ParseDays(parsed[1])
	if err == nil {
		deadlineDate = deadlineDate.AddDate(0, 0, days-1)
	}

	hours, minutes := deadlineDate.Hour(), deadlineDate.Minute()
	if len(parsed) > 2 {
		h, m, err := helpers.ParseHoursAndMinutes(parsed[2])
		if err != nil {
			hours, minutes = h, m
		}
	}

	deadlineDate = helpers.SetHoursAndMinutes(deadlineDate, hours, minutes)
	if len(parsed) > 3 && models.ValidFrequency(models.Frequency(parsed[3])) {
		frequency = models.Frequency(parsed[3])
	}

	return &Reminder{
		Name:            name,
		Deadline:        deadlineDate,
		RemindFrequency: frequency,
		UserID:          userID,
	}, nil
}
