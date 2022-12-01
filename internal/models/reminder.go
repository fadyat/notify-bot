package models

import "time"

type Reminder struct {
	ID        int64
	Content   string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time

	// Deadline is the time when the task will be marked as overdue, if it is not completed.
	Deadline time.Time

	// RemindFrequency is a string that represents the frequency of the reminder.
	RemindFrequency time.Duration
}

type ReminderTg struct {
	Reminder
	ChatID int64
}
