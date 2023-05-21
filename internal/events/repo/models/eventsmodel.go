package models

type EventDetails struct {
	Id        string `json:"event_id" dynamodbav:"event_id"`
	Name      string `json:"name" dynamodbav:"name"`
	CreatorId string `json:"creator_id" dynamodbav:"creator_id"`
	CreatedAt string `json:"c_at" dynamodbav:"c_at"`
}

type EventCommons struct {
	Id      string `json:"e_id" dynamodbav:"e_id"`
	Creator struct {
		Fn string `json:"fn"`
		Mn string `json:"mn"`
		Ln string `json:"ln"`
		Id string `json:"id"`
	} `json:"creator"`
	CAt      string `json:"c_at"`
	UAt      string `json:"u_at"`
	Name     string `json:"name"`
	ShareUrl string `json:"share_url"`
}
