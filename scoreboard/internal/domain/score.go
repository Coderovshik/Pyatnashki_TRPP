package domain

import "time"

type Score struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	Score     int       `json:"score"`
}
