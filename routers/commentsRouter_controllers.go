package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["service-cloud/controllers:PluginController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:PluginController"],
		beego.ControllerComments{
			Method: "GetList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service-cloud/controllers:PluginController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:PluginController"],
		beego.ControllerComments{
			Method: "AddPlugin",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
