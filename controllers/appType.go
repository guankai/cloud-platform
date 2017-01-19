package controllers

import (
	"github.com/astaxie/beego"
	"service-cloud/models"
	"github.com/astaxie/beego/logs"
	"net/http"
)

type AppTypeController struct {
	beego.Controller
}
// @Description 获取所有的AppType
// @router / [get]
func (this *AppTypeController) GetAllAppType() {
	appTypes, err := models.GetAppTypes()
	if err != nil {
		logs.Error("获取所有的应用类型失败%v", err)
		this.CustomAbort(http.StatusInternalServerError, "获取所有的应用类型失败")
	}
	this.Data["json"] = appTypes
	this.ServeJSON()
}
