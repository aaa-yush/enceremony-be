package events

import (
	"enceremony-be/pkg/product"
	"time"
)

type EventList struct {
	Events []EventListItem `json:"events"`
}

type EventDetailCommons struct {
	Id        string    `json:"id"`
	Creator   *Creator  `json:"creator"`
	CAt       time.Time `json:"c_at"`
	UAt       time.Time `json:"u_at"`
	EventDate time.Time `json:"event_date"`
	Name      string    `json:"name"`
	ShareUrl  string    `json:"share_url,omitempty"`
}

type EventListItem struct {
	EventDetailCommons
}

type Creator struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type EventDetails struct {
	EventDetailCommons
	PhotoUrl         string            `json:"photo_url"`
	Products         []product.Product `json:"products"`
	EventDescription string            `json:"event_description"`
}
