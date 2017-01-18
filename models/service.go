package models

import (
	"time"
	"service-cloud/utils/db"
	"github.com/astaxie/beego/orm"
)

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
	UpstreamUrl string `orm:"column(upstream_url)" json:"upstreamUrl"`
	Type     *ClType `orm:"rel(fk)"`
}

type QueryCon struct {
	Keyword string `json:"keyword"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
}

// 新增一个service
func InsertService(service *ClService) error {
	o := db.GetOrmer()
	_, err := o.Insert(service)
	if err != nil {
		return err
	}
	return nil
}
// 获取所有的service列表
func GetServiceList(queryJson *QueryCon) ([]ClService, int, error) {
	o := db.GetOrmer()
	var serviceList []ClService
	cond := orm.NewCondition()
	if len(queryJson.Keyword) == 0 {
		cond = nil
	} else {
		cond = cond.And("service_name__contains", queryJson.Keyword)
	}
	qs := o.QueryTable("cl_service")
	if cond != nil {
		qs = qs.SetCond(cond)
	}
	qs = qs.OrderBy("-create_time")

	_, err := qs.Limit(queryJson.Limit, queryJson.Offset).All(&serviceList)
	if err != nil {
		return nil, -1, err
	}
	count, errCount := qs.Count()
	if errCount != nil {
		return nil, -1, errCount
	}
	return serviceList, int(count), nil
}
//更新服务
func UpdService(clService *ClService) error {
	o := db.GetOrmer()
	_, err := o.Update(clService, "service_name", "service_pic", "request_path", "provider", "version", "service_desc", "upstream_url")
	if err != nil {
		return err
	}
	return nil
}
//获取服务
func GetService(serviceId string) (*ClService, error) {
	o := db.GetOrmer()
	var clService ClService
	err := o.QueryTable("cl_service").Filter("service_id", serviceId).One(&clService)
	if err != nil {
		return nil, err
	}
	return &clService, nil
}
//删除服务
func DelService(serviceId string) error {
	o := db.GetOrmer()
	_, err := o.Delete(&ClService{Id:serviceId})
	if err != nil {
		return err
	}
	return nil
}
