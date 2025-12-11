package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`    // JSON response’ta gizle
	Role      string    `json:"role"` // admin/user
	CreatedAt time.Time `json:"created_at"`
}
