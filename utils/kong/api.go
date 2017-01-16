package kong

import "service-cloud/utils/kong/models"

// AddAPI Add API
// name -- optional	The API name. If none is specified, will default to the request_host or request_path.
// request_host -- semi-optional	The public DNS address that points to your API. For example, mockbin.com. At least request_host or request_path or both should be specified.
// request_path -- emi-optional	The public path that points to your API. For example, /someservice. At least request_host or request_path or both should be specified.
// strip_request_path -- optional	Strip the request_path value before proxying the request to the final API. For example a request made to /someservice/hello will be resolved to upstream_url/hello. By default is false.
// preserve_host -- optional	Preserves the original Host header sent by the client, instead of replacing it with the hostname of the upstream_url. By default is false.
// upstream_url	The base target URL that points to your API server, this URL will be used for proxying requests. For example, https://mockbin.com.
func AddAPI(api *models.API) (*models.API, error) {
	// POST /apis/

	return nil, nil
}

// GetAPI Retrieve API
// nameOrID -- (required)The unique identifier or the name of the API to retrieve.
func GetAPI(nameOrID string) (*models.API, error) {
	//GET /apis/{name or id}

	return nil, nil
}

//ListAPIs List APIs
func ListAPIs(size int, offset string) (*models.APIList, error) {
	//GET /apis/
	return nil, nil
}

//UpdateAPI Update API
func UpdateAPI(nameOrID string) (*models.API, error) {
	//PATCH /apis/{name or id}

	return nil, nil
}

//DeleteAPI Delete API
func DeleteAPI(nameOrID string) error {
	//DELETE /apis/{name or id}

	return nil
}
