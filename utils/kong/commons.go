package kong

import "cloud-platform/utils/environment"

var kongAdminURL string

func init() {
	kongAdminURL = environment.GetEnv("KONG_ADMIN_URL", "http://192.168.100.94:8001")
}
