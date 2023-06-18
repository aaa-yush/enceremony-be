package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	Id   string `gorm:"index;primarykey;not null"`
	Name string `gorm:"not null;size:50"`
	//User        User
	Description string `gorm:"size:200"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	EventDate   time.Time
	DeletedAt   gorm.DeletedAt
}

func (ev Event) TableName() string {
	return "events"
}

type EventProducts struct {
	Event   Event
	Product Product
}

func (epr EventProducts) TableName() string {
	return "event_products"
}
