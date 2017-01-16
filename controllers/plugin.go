package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"github.com/astaxie/beego/logs"
	"service-cloud/models"
)

type PluginController struct {
	beego.Controller
}
// @Description 获取KONG所有插件列表
// @router / [get]
func (this *PluginController) GetList() {
	plugins, err := models.GetAllPlugins()
	if err != nil {
		logs.Error("获取所有的插件失败 %v", err)
		this.CustomAbort(http.StatusInternalServerError, "获取所有插件失败")
	}
	this.Data["json"] = plugins
	this.ServeJSON()
}
// @Description 新增一个插件
// @Param pluginId formData string true "插件ID"
// @Param pluginName formData string true "插件名称"
// @Param pluginDesc formData string true "插件描述"
// @router /add [post]
func (this *PluginController) AddPlugin() {
	var plugin models.ClPlugin
	plugin.Name = this.GetString("pluginName")
	plugin.Desc = this.GetString("pluginDesc")
	plugin.Id = this.GetString("pluginId")
	err := models.InsertPlugin(&plugin)
	if err != nil {
		logs.Error("新增插件失败%v", err)
		this.CustomAbort(http.StatusInternalServerError, "新增插件失败")
	}
	this.Data["json"] = "新增插件成功"
	this.ServeJSON()
}
// @Description 删除一个插件
// @Param pluginId path string true "待删除的pluginId"
// @router /del/:pluginId [delete]
func (this *PluginController) DeletePlugin() {
	pluginId := this.GetString(":pluginId")
	err := models.DeletePlugin(pluginId)
	if err != nil {
		logs.Error("删除插件失败%v", err)
		this.CustomAbort(http.StatusInternalServerError, "删除插件失败")
	}
	this.Data["json"] = "删除插件成功"
	this.ServeJSON()
}