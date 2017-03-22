package models

import (
	"cloud-platform/utils/db"
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type ClRelation struct {
	RelationId string `orm:"pk;column(relation_id)" json:"relationId"`
	UserName   string `orm:"column(user_name)" json:"userName"`
	Service    *ClService `orm:"rel(fk)"`
	ApiKeyId   string `orm:"column(api_key_id)" json:"apiKeyId"`
	ApiKey     string `orm:"column(api_key)" json:"apiKey"`
	ConsumerId string `orm:"column(consumer_id)" json:"consumerId"`
	Status     string `orm:"column(status)" json:"status"`
	AccessNum  int    `orm:"column(access_num)" json:"accessNum"`
}

type UserService struct {
	ServiceId   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	ServicePic  string `json:"service_pic"`
	RequestPath string `json:"request_path"`
	ApiId       string `json:"api_id"`
	Provider    string `json:"provider"`
	CreateTime  time.Time `json:"create_time"`
	Version     string `json:"version"`
	ServiceDesc string `json:"service_desc"`
	UpstreamUrl string `json:"upstream_url"`
	Status      string `json:"status"`
	RelationId  string `json:"relation_id"`
	ConsumerId  string `json:"consumer_id"`
	ApiKeyId    string `json:"api_key_id"`
	ApiKey      string `json:"api_key"`
	TypeName    string `json:"type_name"`
}

type RelJson struct {
	UserName string `json:"userName"`
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
}

type TypeJson struct {
	UserName string `json:"userName"`
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
	TypeId   int `json:"typeId"`
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
	var userServiceList []UserService
	var count int
	_, err := o.Raw(`select a.service_id,a.service_name,a.service_pic,a.request_path,a.api_id,a.provider,a.create_time,a.version,a.service_desc,a.upstream_url,ifnull(b.status,'0') status,ifnull(b.relation_id,'') relation_id,ifnull(b.consumer_id,'') consumer_id,ifnull(b.api_key_id,'') api_key_id,ifnull(b.api_key,'') api_key,c.type_name from cl_service a left join cl_relation b on a.service_id = b.service_id and b.user_name = ? left join cl_type c on a.type_id = c.type_id order by status desc limit ?,?`, relJson.UserName, relJson.Offset, relJson.Limit).QueryRows(&userServiceList)
	if err != nil {
		return nil, -1, err
	}
	errCount := o.Raw(`select count(1) from cl_service a left join cl_relation b on a.service_id = b.service_id and b.user_name = ? left join cl_type c on a.type_id = c.type_id`, relJson.UserName).QueryRow(&count)
	if errCount != nil {
		return nil, -1, errCount
	}
	return userServiceList, count, nil
}

func GetAllRelationsByType(typeJson *TypeJson) ([]UserService, int, error) {
	o := db.GetOrmer()
	var userServiceList []UserService
	var count int
	_, err := o.Raw(`select a.service_id,a.service_name,a.service_pic,a.request_path,a.api_id,a.provider,a.create_time,a.version,a.service_desc,a.upstream_url,ifnull(b.status,'0') status,ifnull(b.relation_id,'') relation_id,ifnull(b.consumer_id,'') consumer_id,ifnull(b.api_key_id,'') api_key_id,ifnull(b.api_key,'') api_key,c.type_name from cl_service a left join cl_relation b on a.service_id = b.service_id and b.user_name = ? left join cl_type c on a.type_id = c.type_id where c.type_id = ? order by status desc limit ?,?`, typeJson.UserName, typeJson.TypeId, typeJson.Offset, typeJson.Limit).QueryRows(&userServiceList)
	if err != nil {
		fmt.Println("enter err")
		return nil, -1, err
	}
	errCount := o.Raw(`select count(1) from cl_service a left join cl_relation b on a.service_id = b.service_id and b.user_name = ? left join cl_type c on a.type_id = c.type_id where c.type_id = ?`, typeJson.UserName, typeJson.TypeId).QueryRow(&count)
	if errCount != nil {
		fmt.Println("enter errCount")
		return nil, -1, errCount
	}
	return userServiceList, count, nil
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
	err := o.QueryTable("cl_relation").Filter("service_id", serviceId).Filter("user_name", userName).RelatedSel().One(relation)
	if err != nil {
		return nil, err
	}
	return relation, nil
}

func GetUserServiceCount(userName string) (int, []ClRelation, error) {
	o := db.GetOrmer()
	var relation []ClRelation
	count, err := o.QueryTable("cl_relation").Filter("user_name", userName).RelatedSel().All(&relation)
	if err != nil {
		if err == orm.ErrNoRows {
			return 0, nil, nil
		} else {
			return -1, nil, err
		}
	}
	return int(count), relation, nil
}
