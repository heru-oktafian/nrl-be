package entity

import (
	"time"
)

type AdminUser struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	Password  string         `json:"-"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}