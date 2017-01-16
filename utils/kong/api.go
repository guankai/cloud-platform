package kong

import "service-cloud/utils/kong/models"

//Add Add API
func Add(api *models.API) (*models.API, error) {
	// POST /apis/

	return nil, nil
}

//Get Retrieve API
func Get(nameOrID string) (*models.API, error) {
	//GET /apis/{name or id}

	return nil, nil
}

//All List APIs
func All() ([]models.API, error) {
	//GET /apis/
	return nil, nil
}

//Update Update API
func Update(nameOrID string) (*models.API, error) {
	//PATCH /apis/{name or id}

	return nil, nil
}

//Delete Delete API
func Delete(nameOrID string) error {
	//DELETE /apis/{name or id}

	return nil
}
