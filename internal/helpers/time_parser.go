package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func SetHoursAndMinutes(t time.Time, hours, minutes int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hours, minutes, 0, 0, t.Location())
}

func ParseHoursAndMinutes(timeStr string) (h, m int, err error) {
	_, err = fmt.Sscanf(timeStr, "%d:%d", &h, &m)
	if err != nil {
		return 0, 0, err
	}
	return h, m, nil
}

func ParseDays(daysStr string) (int, error) {
	days, err := strconv.Atoi(daysStr)
	if err != nil {
		return 0, err
	}
	return days, nil
}
