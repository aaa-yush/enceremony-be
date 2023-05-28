package models

import "time"

type User struct {
	Id          string `gorm:"primarykey"`
	Name        string `gorm:"size:100"`
	Email       string `gorm:"size:110"`
	DateOfBirth time.Time
	Gender      string `gorm:"size:10"`
	PhoneNumber string `gorm:"size:12"`
	CountryCode string `gorm:"size:5"`
	Address     string `gorm:"size:200"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (us User) TableName() string {
	return "user"
}
