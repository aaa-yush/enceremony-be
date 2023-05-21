package events

import "enceremony-be/pkg/product"

type EventList struct {
	Events []EventListItem `json:"events"`
}

type EventDetailCommons struct {
	Id       string   `json:"id"`
	Creator  *Creator `json:"creator"`
	CAt      string   `json:"c_at"`
	UAt      string   `json:"u_at"`
	Name     string   `json:"name"`
	ShareUrl string   `json:"share_url"`
}

type EventListItem struct {
	EventDetailCommons
}

type Creator struct {
	Fn string `json:"fn"`
	Mn string `json:"mn"`
	Ln string `json:"ln"`
	Id string `json:"id"`
}

type EventDetails struct {
	EventDetailCommons
	PhotoUrl string            `json:"photo_url"`
	Products []product.Product `json:"products"`
}
