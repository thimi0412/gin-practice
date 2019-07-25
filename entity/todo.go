package entity

import "time"

// Todo todoプロパティ
type Todo struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Context   string     `json:"context"`
	LimitDate *time.Time `json:"limit_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
