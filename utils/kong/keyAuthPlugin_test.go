package kong

import (
	"service-cloud/utils/kong/models"
	"testing"
)

func TestAddKeyAuthPlugin2Api(t *testing.T) {
	//Creat an api for using first.
	var api models.API
	api.Name = "testkeyauth"
	api.PreserveHost = true
	api.RequestHost = "reqhost"
	api.RequestPath = "/test"
	api.StripRequestPath = false
	api.UpstreamURL = "http://www.baidu.com/"
	retApi, _ := AddAPI(&api)
	//Creat a consumer for using next.
	var consumer models.Consumer
	consumer.Username = "testkeyauth"
	retConsumer, _ := AddConsumer(&consumer)
	defer func() {
		//test DeleteConsumer
		err := DeleteConsumer(retConsumer.ID)
		if err != nil {
			t.Error("DeleteConsumer Error, ", err)
		}
		//DeleteAPI
		err = DeleteAPI(retApi.ID)
		if err != nil {
			t.Error("DeleteAPI Error, ", err)
		}
	}()

	//----------------------------key-auth test below-----------------------------------

	//test add plugin
	var plugin models.KeyAuthPlugin
	apiNameOrId := retApi.ID
	consumerId := retConsumer.ID

	retKeyAuthPlugin, err := AddKeyAuthPlugin2Api(apiNameOrId)
	if err != nil {
		t.Error("AddKeyAuthPlugin Error, ", err)
	}
	if len(retKeyAuthPlugin.ID) < 1 {
		t.Error("AddKeyAuthPlugin2Api Error, ", *retKeyAuthPlugin)
	}
	// test ListKeyAuthPlugin
	pluginlist, err := ListKeyAuthPlugin(plugin, 5, "")
	if err != nil {
		t.Error("ListKeyAuthPlugin Error, ", err)
	}
	if pluginlist.Total < 1 {
		t.Error("ListKeyAuthPlugin Error, ", *pluginlist)
	}
	//test GetKeyAuthPlugin
	retKeyAuthPlugin, err = GetKeyAuthPlugin(retKeyAuthPlugin.ID)
	if err != nil {
		t.Error("GetKeyAuthPlugin Error, ", err)
	}
	if len(retKeyAuthPlugin.ID) < 1 {
		t.Error("GetKeyAuthPlugin Error, ", *retKeyAuthPlugin)
	}

	apikey, err := CreateAPIKey(consumerId, "")
	if err != nil {
		t.Error("CreateAPIKey Error, ", consumerId, err)
	}
	if len(apikey.ID) < 1 {
		t.Error("CreateAPIKey Error, ", *retKeyAuthPlugin)
	}
	// test ListKeyAuthPlugin
	apikeylist, err := ListAPIKey(consumerId, 5, "")
	if err != nil {
		t.Error("ListAPIKey Error, ", err)
	}
	if apikeylist.Total < 1 {
		t.Error("ListAPIKey Error, ", *apikeylist)
	}
	//test GetKeyAuthPlugin
	apikey, err = GetAPIKey(consumerId, apikey.ID)
	if err != nil {
		t.Error("GetAPIKey Error, ", err)
	}
	if len(apikey.ID) < 1 {
		t.Error("GetAPIKey Error, ", *apikey)
	}

	//test DeleteAPIKey
	err = DeleteAPIKey(consumerId, apikey.ID)
	if err != nil {
		t.Error("DeleteAPIKey Error, ", err)
	}
	apikey, err = GetAPIKey(consumerId, apikey.ID)
	if err == nil {
		t.Error("DeleteAPIKey Error. process failure!")
	}

	//test DeleteKeyAuthPlugin
	err = DeleteKeyAuthPluginPerApi(retKeyAuthPlugin.ID, apiNameOrId)
	if err != nil {
		t.Error("DeleteKeyAuthPlugin Error, ", err)
	}
	retKeyAuthPlugin, err = GetKeyAuthPlugin(retKeyAuthPlugin.ID)
	if err == nil {
		t.Error("DeleteKeyAuthPlugin Error. process failure!")
	}

}
