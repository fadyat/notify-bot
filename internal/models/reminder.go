package models

import (
	"fmt"
	"time"
)

type Reminder struct {
	ID        int64     `json:"id,omitempty"`
	UserID    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Deadline is the time when the task will be marked as overdue, if it is not completed.
	Deadline time.Time `json:"deadline"`

	// RemindFrequency is a string that represents the frequency of the reminder.
	RemindFrequency Frequency `json:"remind_frequency"`
}

func (r *Reminder) ToString() string {
	return fmt.Sprintf(
		"Name: %s\nDeadline: %s\nFrequency: %s\n",
		r.Name,
		r.Deadline.Format("2006-01-02 15:04"),
		r.RemindFrequency,
	)
}
