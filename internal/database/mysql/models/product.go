package models

import "time"

type Product struct {
	Id          string `gorm:"primarykey;not null"`
	Name        string `gorm:"not null;size:50"`
	User        User
	Description string `gorm:"size:200"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	EventDate   time.Time
}

func (pr Product) TableName() string {
	return "product"
}
