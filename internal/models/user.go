package models

import "time"

type User struct {
	Name        string
	Email       string
	Password    string
	IsBanned    bool
	IsConfirmed bool
	CreatedAt   time.Time
}
