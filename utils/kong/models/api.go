package models

//API is api model of kong
type API struct {
	ID               string `json:"ip"`
	Name             string `json:"name"`               //"Mockbin",
	RequestHost      string `json:"request_host"`       //"mockbin.com",
	RequestPath      string `json:"request_path"`       //"/someservice",
	StripRequestPath string `json:"strip_request_path"` //false,
	PreserveHost     string `json:"preserve_host"`      //false,
	UpstreamURL      string `json:"upstream_url"`       //"https://mockbin.com"
}
