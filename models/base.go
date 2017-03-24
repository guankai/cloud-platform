package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(ClPlugin), new(ClService), new(ClRelation),new(ClType))
}

// CodeInfo definiton.
type CodeInfo struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

// NewErrorInfo return a CodeInfo represents error.
func NewErrorInfo(info string) *CodeInfo {
	return &CodeInfo{-1, info}
}

// NewNormalInfo return a CodeInfo represents OK.
func NewNormalInfo(info string) *CodeInfo {
	return &CodeInfo{0, info}
}
