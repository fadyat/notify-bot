package repo

import (
	"deadlines/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"time"
)

type ReminderRepoCron struct {
	c *sqlx.DB
}

func NewReminderRepoCron(c *sqlx.DB) *ReminderRepoCron {
	return &ReminderRepoCron{c: c}
}

func (r *ReminderRepoCron) GetRemindersToRun() ([]*models.Reminder, error) {
	tx, err := r.c.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	q := `SELECT * FROM reminders
          WHERE (
              		frequency NOT IN ('once', 'never') AND
              		calc_next_cron_time(updated_at, frequency) <= $1
              	)
             OR 
              	(
              	    frequency = 'once' AND
              	 	deadline - $1 <= 30 * interval '1 minute' AND
              	 	deadline - $1 >= 0 * interval '1 minute'
              	)`

	rows, err := tx.Query(q, time.Now().UTC())
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

func (r *ReminderRepoCron) ChangeUpdatedAtMultiple(ids []int64) error {
	tx, err := r.c.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	q := `UPDATE reminders
		  SET updated_at = $1,
		      frequency = CASE WHEN frequency = 'once' THEN 'never' ELSE frequency END
		  WHERE id = ANY($2)`
	if _, err = tx.Exec(q, time.Now().UTC(), pq.Array(ids)); err != nil {
		return err
	}

	return tx.Commit()
}
