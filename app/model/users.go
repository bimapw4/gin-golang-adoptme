package model

import "time"

type User struct {
	Id                uint
	Fullname          string
	Email             string `gorm:"unique"`
	Password          string
	Phone             string `gorm:"size:15"`
	Verification_code int    `gorm:"size:4"`
	Address           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
