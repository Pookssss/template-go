package models

import "time"

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
