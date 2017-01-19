package models

import "service-cloud/utils/db"

// 应用类型
type ClType struct {
	Id   int `orm:"column(type_id);pk;auto" json:"typeId"`
	Name string `orm:"column(type_name)" json:"typeName"`
	Desc string `orm:"column(type_desc)" json:"typeDesc"`
}

func InsertType(appType *ClType) (error) {
	o := db.GetOrmer()
	_, err := o.Insert(appType)
	if err != nil {
		return err
	}
	return nil
}

func GetAppTypes() ([]ClType, error) {
	o := db.GetOrmer()
	var appTypes []ClType
	_, err := o.QueryTable("cl_type").All(&appTypes)
	if err != nil {
		return nil, err
	}
	return appTypes, nil
}

func GetAppType(typeId int) (*ClType, error) {
	o := db.GetOrmer()
	var appType ClType
	err := o.QueryTable("cl_type").Filter("type_id", typeId).One(&appType)
	if err != nil {
		return nil, err
	}
	return &appType, nil
}