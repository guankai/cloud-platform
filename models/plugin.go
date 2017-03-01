package models

import "cloud-platform/utils/db"
//插件
type ClPlugin struct {
	Id   string `orm:"pk;column(plugin_id)" json:"pluginId"`
	Name string `orm:"column(plugin_name)" json:"pluginName"`
	Desc string `orm:"column(plugin_desc)" json:"pluginDesc"`
}
//获取所有插件的列表
func GetAllPlugins() ([]ClPlugin, error) {
	o := db.GetOrmer()
	var pluginList []ClPlugin
	_, err := o.QueryTable("cl_plugin").All(&pluginList)
	if err != nil {
		return nil, err
	}
	return pluginList, nil
}
//获取插件by插件id
func GetPlugin(pluginId string) (*ClPlugin, error) {
	o := db.GetOrmer()
	var plugin ClPlugin
	err := o.QueryTable("cl_plugin").Filter("plugin_id", pluginId).One(&plugin)
	if err != nil {
		return nil, err
	}
	return &plugin, nil
}
//获取插件by插件名称
func GetPluginByName(pluginName string) (*ClPlugin, error) {
	o := db.GetOrmer()
	var plugin ClPlugin
	err := o.QueryTable("cl_plugin").Filter("plugin_name", pluginName).One(&plugin)
	if err != nil {
		return nil, err
	}
	return &plugin, nil
}
//insert插件
func InsertPlugin(plugin *ClPlugin) error {
	o := db.GetOrmer()
	_, err := o.Insert(plugin)
	if err != nil {
		return err
	}
	return nil
}

func DeletePlugin(pluginId string) error {
	o := db.GetOrmer()
	_, err := o.Delete(&ClPlugin{Id:pluginId})
	if err != nil {
		return err
	}
	return nil
}