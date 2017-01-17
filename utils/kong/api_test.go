package kong

import (
	"service-cloud/utils/kong/models"
	"testing"
)

func TestAddApi(t *testing.T) {
	var api models.API
	api.Name = "test"
	api.PreserveHost = true
	api.RequestHost = "reqhost"
	api.RequestPath = "/test"
	api.StripRequestPath = false
	api.UpstreamURL = "http://www.baidu.com/"
	retApi, err := AddAPI(&api)
	if err != nil {
		t.Error("Error, ", err)
	}
	if retApi.Name != api.Name {
		t.Errorf("name error: %s.", retApi.Name)
		t.Error("Error, ", *retApi)
	}
}
