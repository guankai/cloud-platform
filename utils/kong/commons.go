package kong

import "cloud-platform/utils/environment"

var kongAdminURL string

func init() {
	kongAdminURL = environment.GetEnv("KONG_ADMIN_URL", "http://223.202.32.56:8056")
}
