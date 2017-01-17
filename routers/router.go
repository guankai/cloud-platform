// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"service-cloud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("serviceFind/v1",
		beego.NSNamespace("/plugin",
			beego.NSInclude(
				&controllers.PluginController{},
			),
		),
		beego.NSNamespace("/service",
			beego.NSInclude(
				&controllers.ServiceController{},
			),
		),
		beego.NSNamespace("/relation",
			beego.NSInclude(
				&controllers.RelationController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
