package models

import "time"

type ClService struct {
	Id          string `orm:"pk;column(service_id)" json:"serviceId"`
	ServiceName string `orm:"column(service_name)" json:"serviceName"`
	ServicePic  string `orm:"column(service_pic)" json:"servicePic"`
	RequestPath string `orm:"column(request_path)" json:"requestPath"`
	ApiId       string `orm:"column(api_id)" json:"apiId"`
	Plugin      *ClPlugin `orm:"rel(fk)"`
	Provider    string `orm:"column(provider)" json:"provider"`
	CreateTime  time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"createTime"`
	Version     string `orm:"column(version)" json:"version"`
	ServiceDesc string `orm:"column(service_desc)" json:"serviceDesc"`
}
