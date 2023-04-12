package domain

import "time"

type ActivityParams struct {
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
