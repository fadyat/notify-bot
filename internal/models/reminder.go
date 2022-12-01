package models

import (
	"encoding/json"
	"time"
)

type Reminder struct {
	ID        int64     `json:"id,omitempty"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Deadline is the time when the task will be marked as overdue, if it is not completed.
	Deadline time.Time `json:"deadline"`

	// RemindFrequency is a string that represents the frequency of the reminder.
	RemindFrequency time.Duration `json:"remind_frequency"`
}

func (r *Reminder) ToString() string {
	s, _ := SerializeReminder(r)
	return s
}

func DeserializeReminder(data string) (*Reminder, error) {
	var reminder Reminder
	if err := json.Unmarshal([]byte(data), &reminder); err != nil {
		return nil, err
	}

	return &reminder, nil
}

func SerializeReminder(reminder *Reminder) (string, error) {
	data, err := json.Marshal(reminder)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
