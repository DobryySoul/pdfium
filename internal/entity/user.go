package entity

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email"`
	PassHash  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
