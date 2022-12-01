package models

type Frequency string

const (
	Once    Frequency = "once"
	Daily   Frequency = "daily"
	Weekly  Frequency = "weekly"
	Monthly Frequency = "monthly"
)

func ValidFrequency(f Frequency) bool {
	return f == Once || f == Daily || f == Weekly || f == Monthly
}
