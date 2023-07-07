package model

import "time"

type User struct {
	ID        int    `gorm:"primary_key;"`
	UserId    string `json:"user_id" db:"user_id"`
	Username  string
	Password  string `json:"-"`
	Email     string
	Gender    int       `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (User) TableName() string {
	return "user"
}
