package models

import (
	"time"
)

//API is api model of kong
type API struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`               //"Mockbin",
	RequestHost      string    `json:"request_host"`       //"mockbin.com",
	RequestPath      string    `json:"request_path"`       //"/someservice",
	StripRequestPath bool      `json:"strip_request_path"` //false,
	PreserveHost     bool      `json:"preserve_host"`      //false,
	UpstreamURL      string    `json:"upstream_url"`       //"https://mockbin.com"
	CreateAt         time.Time `json:"created_at"`
}

//APIList find all apis by page
type APIList struct {
	Total int    `json:"total"` // total count of apis
	Data  []API  `json:"data"`  // apis
	Next  string `json:"next"`  // next page url
}
