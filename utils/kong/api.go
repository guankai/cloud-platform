package kong

import "service-cloud/utils/kong/models"

//AddAPI Add API
func AddAPI(api *models.API) (*models.API, error) {
	// POST /apis/

	return nil, nil
}

//GetAPI Retrieve API
func GetAPI(nameOrID string) (*models.API, error) {
	//GET /apis/{name or id}

	return nil, nil
}

//ListAPIs List APIs
func ListAPIs() ([]models.API, error) {
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
