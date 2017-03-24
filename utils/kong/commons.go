package kong

import (
	"cloud-platform/utils/environment"
	"github.com/astaxie/beego"
)

var kongAdminURL string

func init() {
	kongAdminURL = environment.GetEnv("KONG_ADMIN_URL", beego.AppConfig.String("kong_admin_url"))
}
