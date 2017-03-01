package kong

import (
	"cloud-platform/utils/kong/models"
	"testing"
)

func TestAddApi(t *testing.T) {
	//test add api
	var api models.API
	api.Name = "test"
	api.PreserveHost = true
	api.RequestHost = "reqhost"
	api.RequestPath = "/test"
	api.StripRequestPath = false
	api.UpstreamURL = "http://www.baidu.com/"

	retApi, err := AddAPI(&api)
	if err != nil {
		t.Error("AddAPI Error, ", err)
	}
	if retApi.Name != api.Name {
		t.Errorf("AddAPI error: %s.", retApi.Name)
		t.Error("AddAPI Error, ", *retApi)
	}
	// test ListAPIs
	apilist, err := ListAPIs(5, "")
	if err != nil {
		t.Error("ListAPIs Error, ", err)
	}
	if apilist.Total < 1 {
		t.Error("ListAPIs Error, ", *apilist)
	}
	//test GetAPI
	retApi, err = GetAPI(api.Name)
	if err != nil {
		t.Error("GetAPI Error, ", err)
	}
	if retApi.Name != api.Name {
		t.Error("GetAPI Error, ", *retApi)
	}
	//test UpdateAPI
	api.Name = "test1"
	retApi, err = UpdateAPI(retApi.Name, &api)
	if err != nil {
		t.Error("UpdateAPI Error, ", err)
	}
	if retApi.Name != "test1" {
		t.Error("UpdateAPI Error, ", *retApi)
	}

	//test DeleteAPI
	err = DeleteAPI(retApi.ID)
	if err != nil {
		t.Error("DeleteAPI Error, ", err)
	}
	retApi, err = GetAPI(retApi.ID)
	if err == nil {
		t.Error("DeleteAPI Error. process failure!")
	}
}
