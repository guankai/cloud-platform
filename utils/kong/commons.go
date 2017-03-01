package kong

import "cloud-platform/utils/environment"

var kongAdminURL string

func init() {
	kongAdminURL = environment.GetEnv("KONG_ADMIN_URL", "http://13.76.42.81:8001")
}
