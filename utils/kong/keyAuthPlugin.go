package kong

import (
	"errors"
	"net/url"
	"service-cloud/utils/kong/models"
	"strconv"

	"github.com/astaxie/beego/httplib"
	urlQuery "github.com/google/go-querystring/query"
)

// AddKeyAuthPlugin2Api Add Plugin to  a specific API
func AddKeyAuthPlugin2Api(apiNameOrId string) (*models.KeyAuthPlugin, error) {
	//curl -X POST http://kong:8001/apis/{api}/plugins  --data "name=key-auth"
	if len(apiNameOrId) == 0 {
		return nil, errors.New("The unique identifier or the name of the api can not be null")
	}
	req := httplib.Post(kongAdminURL + `/apis/` + apiNameOrId + `}/plugins`)
	req.Param("name", "key-auth")
	var retKeyAuthPlugin models.KeyAuthPlugin
	err := req.ToJSON(&retKeyAuthPlugin)
	if err != nil {
		return nil, err
	}
	return &retKeyAuthPlugin, nil
}

// GetKeyAuthPlugin Retrieve Plugin
// id required	The unique identifier of the plugin to retrieve
func GetKeyAuthPlugin(id string) (*models.KeyAuthPlugin, error) {
	//Get /plugins/{id}
	if len(id) == 0 {
		return nil, errors.New("The unique identifier of the KeyAuthPlugin can not be null")
	}
	req := httplib.Get(kongAdminURL + `/plugins/` + id)

	var retKeyAuthPlugin models.KeyAuthPlugin
	err := req.ToJSON(&retKeyAuthPlugin)
	if err != nil {
		return nil, err
	}
	return &retKeyAuthPlugin, nil
}

// ListKeyAuthPlugin List All Plugins
// plugin.id -- optional	A filter on the list based on the id field.
// plugin.name -- optional	A filter on the list based on the name field.
// plugin.ApiId -- optional	A filter on the list based on the api_id field.
// plugin.consumer_id -- optional	A filter on the list based on the consumer_id field.
// plugin.size -- optional, default is 100	A limit on the number of objects to be returned.
// plugin.offset -- optional	A cursor used for pagination. offset is an object identifier that defines a place in the list.
func ListKeyAuthPlugin(plugin models.KeyAuthPlugin, size int, offset string) (*models.KeyAuthPluginList, error) {
	//GET /plugins/

	//do get
	u, err := url.Parse(kongAdminURL + `/plugins/`)
	if err != nil {
		return nil, err
	}
	urlValues, _ := urlQuery.Values(plugin)
	if size > 0 {
		urlValues.Add("size", strconv.Itoa(size))
	}
	if len(offset) > 0 {
		urlValues.Add("offset", offset)
	}

	u.RawQuery = urlValues.Encode()
	req := httplib.Get(u.String())

	//to json
	var retKeyAuthPluginList models.KeyAuthPluginList
	err = req.ToJSON(&retKeyAuthPluginList)
	if err != nil {
		return nil, err
	}
	return &retKeyAuthPluginList, nil
}

// ListKeyAuthPluginPerApi List All Plugins for specific api
// plugin.id -- optional	A filter on the list based on the id field.
// plugin.name -- optional	A filter on the list based on the name field.
// plugin.api_id -- optional	A filter on the list based on the api_id field.
// plugin.consumer_id -- optional	A filter on the list based on the consumer_id field.
// plugin.size -- optional, default is 100	A limit on the number of objects to be returned.
// plugin.offset -- optional	A cursor used for pagination. offset is an object identifier that defines a place in the list.
// apiNameOrId --
func ListKeyAuthPluginPerApi(plugin models.KeyAuthPlugin, size int, offset string, apiNameOrId string) (*models.KeyAuthPluginList, error) {
	//GET /apis/{api name or id}/plugins/
	if len(apiNameOrId) == 0 {
		return nil, errors.New("The unique identifier or the name of the api can not be null")
	}
	//do get
	u, err := url.Parse(kongAdminURL + `/apis/` + apiNameOrId + `/plugins/`)
	if err != nil {
		return nil, err
	}
	urlValues, _ := urlQuery.Values(plugin)
	if size > 0 {
		urlValues.Add("size", strconv.Itoa(size))
	}
	if len(offset) > 0 {
		urlValues.Add("offset", offset)
	}

	u.RawQuery = urlValues.Encode()
	req := httplib.Get(u.String())

	//to json
	var retKeyAuthPluginList models.KeyAuthPluginList
	err = req.ToJSON(&retKeyAuthPluginList)
	if err != nil {
		return nil, err
	}
	return &retKeyAuthPluginList, nil
}

//TODO update

// DeleteKeyAuthPluginPerApi Delete Plugin
// apiNameOrId -- required	The unique identifier or the name of the API for which to delete the plugin configuration
// id -- required	The unique identifier of the plugin configuration to delete on this API
func DeleteKeyAuthPluginPerApi(id string, apiNameOrId string) error {
	//DELETE /apis/{api name or id}/plugins/{id}
	if len(apiNameOrId) == 0 {
		return errors.New("The unique identifier or the name of the API can not be null")
	}
	if len(id) == 0 {
		return errors.New("The unique identifier of the plugin can not be null")
	}

	req := httplib.Delete(kongAdminURL + `/apis/` + apiNameOrId + `/plugins/` + id)

	_, err := req.Response()
	if err != nil {
		return err
	}
	return nil
}

// CreateAPIKey Create an API Key
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
// apikey -- (optional)You can optionally set your own unique key to authenticate the client. If missing, the plugin will generate one.
func CreateAPIKey(consumerNameOrId string, apikey string) (*models.ApiKey, error) {
	//POST http://kong:8001/consumers/{consumerNameOrId}/key-auth -d '' HTTP/1.1 201 Created
	if len(consumerNameOrId) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	req := httplib.Post(kongAdminURL + `/consumers/` + consumerNameOrId + `/key-auth`)
	req.Param("key", apikey)
	var retApiKey models.ApiKey
	err := req.ToJSON(&retApiKey)
	if err != nil {
		return nil, err
	}
	return &retApiKey, nil
}

func GetAPIKey(consumerNameOrId string, apikeyid string) (*models.ApiKey, error) {
	// GET http://13.76.42.81:8001/consumers/guan/key-auth/a0fdba77-fc6d-4632-845c-cc1a623cf59d
	if len(consumerNameOrId) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	if len(apikeyid) == 0 {
		return nil, errors.New("The unique identifier of the apikey can not be null")
	}
	req := httplib.Get(kongAdminURL + `/consumers/` + consumerNameOrId + `/key-auth/` + apikeyid)
	var retApiKey models.ApiKey
	err := req.ToJSON(&retApiKey)
	if err != nil {
		return nil, err
	}
	return &retApiKey, nil
}

// ListAPIKey Create an API Key
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
func ListAPIKey(consumerNameOrId string, size int, offset string) (*models.ApiKeyList, error) {
	//Get http://kong:8001/consumers/{consumerNameOrId}/key-auth?size=1 -d '' HTTP/1.1 201 Created
	if len(consumerNameOrId) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	//do get
	u, err := url.Parse(kongAdminURL + `/consumers/` + consumerNameOrId + `/key-auth`)
	if err != nil {
		return nil, err
	}
	urlValues := u.Query()
	if size > 0 {
		urlValues.Add("size", strconv.Itoa(size))
	}
	if len(offset) > 0 {
		urlValues.Add("offset", offset)
	}
	u.RawQuery = urlValues.Encode()
	req := httplib.Get(u.String())
	//to json
	var retApiKeyList models.ApiKeyList
	err = req.ToJSON(&retApiKeyList)
	if err != nil {
		return nil, err
	}
	return &retApiKeyList, nil
}

// DeleteAPIKey Create an API Key
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
// apikeyid -- api_key id
func DeleteAPIKey(consumerNameOrId string, apikeyid string) error {
	//Delete http://kong:8001/consumers/{consumerNameOrId}/key-auth/{apikeyid} -d '' HTTP/1.1 201 Created
	if len(consumerNameOrId) == 0 {
		return errors.New("The unique identifier or the name of the consumer can not be null")
	}
	if len(apikeyid) == 0 {
		return errors.New("The unique identifier of the apikey can not be null")
	}
	req := httplib.Delete(kongAdminURL + `/consumers/` + consumerNameOrId + `/key-auth/` + apikeyid)
	_, err := req.Response()
	if err != nil {
		return err
	}
	return nil
}
