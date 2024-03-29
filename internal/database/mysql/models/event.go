package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	Id   string `gorm:"index;primarykey;not null"`
	Name string `gorm:"not null;size:50"`
	//User        User
	UserId      string `gorm:"size:45"`
	Description string `gorm:"size:200"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	EventDate   time.Time
	DeletedAt   gorm.DeletedAt
}

func (ev Event) TableName() string {
	return "events"
}

type EventDetails struct {
	Event
	Products []EventProductDetails
}
