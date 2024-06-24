package structs

import "time"

type Vehicle struct {
	ID        uint      `json:"id"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
