package models

import (
	"service-cloud/utils/db"
	"time"
)

type ClRelation struct {
	RelationId string `orm:"pk;column(relation_id)" json:"relationId"`
	UserName   string `orm:"column(user_name)" json:"userName"`
	Service    *ClService `orm:"rel(fk)"`
	ApiKeyId   string `orm:"column(api_key_id)" json:"apiKeyId"`
	ApiKey     string `orm:"column(api_key)" json:"apiKey"`
	ConsumerId string `orm:"column(consumer_id)" json:"consumerId"`
	Status     string `orm:"column(status)" json:"status"`
}

type UserService struct {
	Id          string `json:"serviceId"`
	ServiceName string `json:"serviceName"`
	ServicePic  string `json:"servicePic"`
	RequestPath string `json:"requestPath"`
	ApiId       string `json:"apiId"`
	Provider    string `json:"provider"`
	CreateTime  time.Time `json:"createTime"`
	Version     string `json:"version"`
	ServiceDesc string `json:"serviceDesc"`
	UpstreamUrl string `json:"upstreamUrl"`
	Status      string `json:"status"`
	RelationId  string `json:"relationId"`
	ConsumerId  string `json:"consumerId"`
	ApiKeyId    string `json:"apiKeyId"`
	ApiKey      string `json:"apikey"`
}

type RelJson struct {
	UserName string `json:"userame"`
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
}

// 添加一个用户与服务的关系
func InsertRelation(relation *ClRelation) error {
	o := db.GetOrmer()
	_, err := o.Insert(relation)
	if err != nil {
		return err
	}
	return nil
}
// 获取用户所有的服务(已开启和未开启)
func GetAllRelations(relJson *RelJson) ([]UserService, int, error) {
	o := db.GetOrmer()
	var userService []UserService
	var count int
	_, err := o.Raw(`select a.service_id serviceId,a.service_name serviceName,a.service_pic servicePic,a.request_path requestPath,a.api_id apiId,a.provider provider,a.create_time createTime,a.version version,a.service_desc serviceDesc,a.upstream_url upstreamUrl,ifnull(b.status,'0') status,ifnull(b.relation_id,'') relationId,ifnull(b.consumer_id,'') consumerId,ifnull(b.api_key_id,'') apiKeyId,ifnull(b.api_key,'') apiKey from cl_service a left join cl_relation b on a.service_id = b.service_id and b.user_name = ? order by status limit ?,?`, relJson.UserName, relJson.Offset, relJson.Limit).QueryRows(&userService)
	if err != nil {
		return nil, -1, err
	}
	errCount := o.Raw(`select count(1) from cl_service a left join cl_relation b on a.service_id = b.service_id and b.user_name = ? order by status limit ?,?`, relJson.UserName, relJson.Offset, relJson.Limit).QueryRow(count)
	if errCount != nil {
		return nil, -1, errCount
	}
	return userService, count, nil
}
// 设置用户服务状态
func SetStatus(relationId string, status string, apiKey string, apikeyId string) (error) {
	o := db.GetOrmer()
	relation := ClRelation{RelationId:relationId}
	errRead := o.Read(&relation)
	if errRead != nil {
		return errRead
	}
	relation.Status = status
	relation.ApiKey = apiKey
	relation.ApiKeyId = apikeyId
	_, err := o.Update(&relation, "status", "api_key", "api_key_id")
	if err != nil {
		return err
	}
	return nil
}
// 通过serviceId和用户名查询relationId
func GetServiceByUser(serviceId string, userName string) (*ClRelation, error) {
	o := db.GetOrmer()
	relation := new(ClRelation)
	err := o.QueryTable("cl_relation").Filter("service_id", serviceId).Filter("userName", userName).RelatedSel().One(relation)
	if err != nil {
		return nil, err
	}
	return relation, nil
}