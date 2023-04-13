package domain

import "time"

type ActivityParams struct {
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TodoParams struct {
	Title           *string `json:"title,omitempty"`
	ActivityGroupId *int64  `json:"activity_group_id,omitempty"`
	IsActive        *bool   `json:"is_active,omitempty"`
	Priority        string  `json:"priority,omitempty"`
}
