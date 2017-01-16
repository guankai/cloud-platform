package models

import "time"

// The Consumer object represents a consumer - or a user - of an API.
type Consumer struct {
	Username string    `json:"username"`   //: "guan",
	CustomID string    `json:"custom_id"`  //: "abc123",
	CreateAt time.Time `json:"created_at"` //: 1484277177000,
	ID       string    `json:"id"`         //: "5878ca3a-13a8-4cee-8ac9-de2cdb588381"
}

//ConsumerList find all Consumers by page
type ConsumerList struct {
	Total  int        `json:"total"` // total count of apis
	Data   []Consumer `json:"data"`  // apis
	Next   string     `json:"next"`  // next page url
	Offset string     `json:"offset"`
}
