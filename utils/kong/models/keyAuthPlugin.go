package models

import "time"

type KeyAuthPlugin struct {
	ID         string        `json:"id" url:"id,omitempty"`                   //: "4d924084-1adb-40a5-c042-63b19db421d1",
	ApiId      string        `json:"api_id" url:"api_id,omitempty"`           //: "5fd1z584-1adb-40a5-c042-63b19db49x21",
	ConsumerId string        `json:"consumer_id" url:"consumer_id,omitempty"` //: "a3dX2dh2-1adb-40a5-c042-63b19dbx83hF4",
	Name       string        `json:"name" url:"name,omitempty"`               //: "rate-limiting",
	Config     KeyAuthConfig `json:"config" url:"-"`                          //:///
	Enabled    bool          `json:"enabled" url:"-"`                         //: true,
	CreatedAt  time.Time     `json:"created_at" url:"-"`                      //: 1422386534
}

type KeyAuthConfig struct {
	KeyNames        []string `json:"key_names"`        //: ["apikey"],
	HideCredentials bool     `json:"hide_credentials"` //:false
}

type ApiKey struct {
	Key        string    `json:"key"`         //: "1234567890",
	ConsumerId string    `json:"consumer_id"` //: "5878ca3a-13a8-4cee-8ac9-de2cdb588381",
	CreatedAt  time.Time `json:"created_at"`  //: 1484278640000,
	ID         string    `json:"id"`          //: "88ca1d4a-cc29-4856-98e2-36c08afbf698"
}

//KeyAuthPluginList find all apis by page
type KeyAuthPluginList struct {
	Total  int             `json:"total"` // total count of apis
	Data   []KeyAuthPlugin `json:"data"`  // apis
	Next   string          `json:"next"`  // next page url
	Offset string          `json:"offset"`
}

type ApiKeyList struct {
	Total  int      `json:"total"` // total count of apis
	Data   []ApiKey `json:"data"`  // apis
	Next   string   `json:"next"`  // next page url
	Offset string   `json:"offset"`
}
