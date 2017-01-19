package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(ClPlugin), new(ClService), new(ClRelation),new(ClType))
}
