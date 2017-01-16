package kong

import "service-cloud/utils/kong/models"

// AddConsumer Create Consumer
// consumer -- You must send either username or custom_id with the request.
func AddConsumer(consumer *models.Consumer) (*models.Consumer, error) {
	// POST /consumers/

	return nil, nil
}

//GetConsumer Retrieve Consumer
//nameOrID -- The unique identifier or the username of the consumer to retrieve
func GetConsumer(nameOrID string) (*models.Consumer, error) {
	//GET /consumers/{username or id}

	return nil, nil
}

//ListConsumers List Consumers
func ListConsumers(size int, offset string) (*models.ConsumerList, error) {
	//GET /consumers/
	return nil, nil
}

//UpdateConsumer Update Consumer
func UpdateConsumer(nameOrID string) (*models.Consumer, error) {
	//PATCH /Consumers/{name or id}

	return nil, nil
}

//DeleteConsumer Delete Consumer
func DeleteConsumer(nameOrID string) error {
	//DELETE /consumers/{username or id}

	return nil
}
