package kong

import "service-cloud/utils/kong/models"

// AddKeyAuthPlugin2Api Add Plugin to  a specific API
func AddKeyAuthPlugin2Api(apiNameOrId string) (*models.KeyAuthPlugin, error) {
	//POST /apis/{name or id}/plugins/

	return nil, nil
}

// GetKeyAuthPlugin Retrieve Plugin
// id required	The unique identifier of the plugin to retrieve
func GetKeyAuthPlugin(id string) (*models.KeyAuthPlugin, error) {
	//Get /plugins/{id}

	return nil, nil
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
	return nil, nil
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

	return nil, nil
}

//TODO update

// DeleteKeyAuthPluginPerApi Delete Plugin
// apiNameOrId -- required	The unique identifier or the name of the API for which to delete the plugin configuration
// id -- required	The unique identifier of the plugin configuration to delete on this API
func DeleteKeyAuthPluginPerApi(id string, apiNameOrId string) error {
	//GET /apis/{api name or id}/plugins/

	return nil
}

// CreateAPIKey Create an API Key
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
// apikey -- (optional)You can optionally set your own unique key to authenticate the client. If missing, the plugin will generate one.
func CreateAPIKey(consumerNameOrId string, apikey string) (*models.ApiKey, error) {
	//POST http://kong:8001/consumers/{consumerNameOrId}/key-auth -d '' HTTP/1.1 201 Created
	return nil, nil
}

// CreateAPIKey Create an API Key
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
func GetAPIKey(consumerNameOrId string) (*models.ApiKey, error) {
	//Get http://kong:8001/consumers/{consumerNameOrId}/key-auth -d '' HTTP/1.1 201 Created
	return nil, nil
}

// DeleteAPIKey Create an API Key
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
// apikeyid -- api_key id
func DeleteAPIKey(consumerNameOrId string, apikeyid string) error {
	//Get http://kong:8001/consumers/{consumerNameOrId}/key-auth -d '' HTTP/1.1 201 Created
	return nil
}
