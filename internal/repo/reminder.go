package repo

import (
	"deadlines/internal/models"
	"deadlines/internal/models/dto"
	"github.com/jmoiron/sqlx"
	"time"
)

type ReminderRepo struct {
	c *sqlx.DB
}

func NewReminderRepo(c *sqlx.DB) *ReminderRepo {
	return &ReminderRepo{c: c}
}

func (r *ReminderRepo) GetReminders(userID int64) ([]*models.Reminder, error) {
	tx, err := r.c.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	q := `SELECT * FROM reminders WHERE user_id = $1`
	rows, err := tx.Query(q, userID)
	if err != nil {
		return nil, err
	}

	var reminders []*models.Reminder
	for rows.Next() {
		rem := &models.Reminder{}
		if err := rows.Scan(
			&rem.ID, &rem.UserID, &rem.Name, &rem.RemindFrequency,
			&rem.Completed, &rem.CreatedAt, &rem.UpdatedAt, &rem.Deadline,
		); err != nil {
			return nil, err
		}
		reminders = append(reminders, rem)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return reminders, nil
}

func (r *ReminderRepo) SaveReminder(rem *dto.Reminder) error {
	tx, err := r.c.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	q := `INSERT INTO reminders (name, frequency, deadline, user_id) VALUES ($1, $2, $3, $4)`
	if _, err = tx.Exec(q, rem.Name, rem.RemindFrequency, rem.Deadline, rem.UserID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ReminderRepo) DeleteReminder(id int64) error {
	tx, err := r.c.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	q := `DELETE FROM reminders WHERE id = $1`
	if _, err = tx.Exec(q, id); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ReminderRepo) UpdateReminder(rem *dto.Reminder) error {
	tx, err := r.c.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	q := `UPDATE reminders
		  SET name = $1, frequency = $2, deadline = $3, updated_at = $4
		  WHERE id = $5`
	if _, err = tx.Exec(q, rem.Name, rem.RemindFrequency, rem.Deadline, rem.ID, time.Now().UTC()); err != nil {
		return err
	}

	return tx.Commit()
}
