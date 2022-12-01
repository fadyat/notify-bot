package cache

import "deadlines/internal/models"

type Cache interface {
	// GetConfiguredReminder returns a reminder that is configured by the user, in current moment.
	GetConfiguredReminder(userID string) (*models.Reminder, error)

	// SetConfiguredReminder sets a reminder that is configured by the user, in current moment.
	SetConfiguredReminder(userID string, reminder *models.Reminder) error

	// ContainsConfiguredReminder checks if a reminder that is configured by the user, in current moment, exists.
	ContainsConfiguredReminder(userID string) (bool, error)

	// DeleteConfiguredReminder deletes a reminder that is configured by the user, in current moment.
	DeleteConfiguredReminder(userID string) error
}
