package cache

import (
	"context"
	"deadlines/internal/models"
	"github.com/go-redis/redis/v9"
	"time"
)

type Redis struct {
	c *redis.Client
}

func NewRedisClient(addr string) *Redis {
	return &Redis{
		c: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
		}),
	}
}

func (r *Redis) GetConfiguredReminder(userID string) (*models.Reminder, error) {
	reminderJSON, err := r.c.Get(context.Background(), userID).Result()
	if err != nil {
		return nil, err
	}

	reminder, err := models.DeserializeReminder(reminderJSON)
	if err != nil {
		return nil, err
	}

	return reminder, nil
}

func (r *Redis) SetConfiguredReminder(userID string, reminder *models.Reminder) error {
	reminderJSON, err := models.SerializeReminder(reminder)
	if err != nil {
		return err
	}

	return r.c.Set(
		context.Background(),
		userID,
		reminderJSON,
		24*30*time.Hour,
	).Err()
}

func (r *Redis) ContainsConfiguredReminder(userID string) (bool, error) {
	v := r.c.Exists(context.Background(), userID).Val()
	return v == 1, nil
}

func (r *Redis) DeleteConfiguredReminder(userID string) error {
	return r.c.Del(context.Background(), userID).Err()
}
