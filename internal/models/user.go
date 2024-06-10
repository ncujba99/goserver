package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // exclude password from json response
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
