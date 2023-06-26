package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id   string `gorm:"primarykey;not null"`
	Name string `gorm:"not null;size:50"`
	//User        User
	UserId      string `gorm:"size:45"`
	Description string `gorm:"size:200"`
	Link        string `gorm:"size:100"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (pr Product) TableName() string {
	return "product"
}

type EventProductDetails struct {
	Product
	IsPurchased bool
}

type EventProducts struct {
	EventId     string `gorm:"size:45"`
	ProductId   string `gorm:"size:45"`
	IsPurchased bool
}

func (epr EventProducts) TableName() string {
	return "event_products"
}

// select p.* from product p left join event_products ep on p.id = ep.product_id where ep.event_id = ?
