package repo

import (
	"deadlines/internal/models"
	"deadlines/internal/models/dto"
)

type IReminderRepo interface {
	GetReminders(userID int64) ([]*models.Reminder, error)
	SaveReminder(rem *dto.Reminder) error
	DeleteReminder(id int64) error
	UpdateReminder(rem *dto.Reminder) error
}
