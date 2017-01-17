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

	beego.GlobalControllerRouter["service-cloud/controllers:PluginController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:PluginController"],
		beego.ControllerComments{
			Method: "DeletePlugin",
			Router: `/del/:pluginId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service-cloud/controllers:RelationController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:RelationController"],
		beego.ControllerComments{
			Method: "GetServiceList",
			Router: `/query`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service-cloud/controllers:RelationController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:RelationController"],
		beego.ControllerComments{
			Method: "EnableService",
			Router: `/startup`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service-cloud/controllers:RelationController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:RelationController"],
		beego.ControllerComments{
			Method: "ShutdownService",
			Router: `/shutdown`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service-cloud/controllers:ServiceController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "AddService",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service-cloud/controllers:ServiceController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "QueryServiceList",
			Router: `/query`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service-cloud/controllers:ServiceController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "UpdateService",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service-cloud/controllers:ServiceController"] = append(beego.GlobalControllerRouter["service-cloud/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "DeleteService",
			Router: `/del/:serviceId`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
