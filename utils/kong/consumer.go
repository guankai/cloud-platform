package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"service-cloud/utils/kong/models"

	"github.com/astaxie/beego/httplib"
)

// AddConsumer Create Consumer
// consumer -- You must send either username or custom_id in the consumer.
func AddConsumer(consumer *models.Consumer) (*models.Consumer, error) {
	// POST /consumers/
	req := httplib.Post(kongAdminURL + `/consumers/`)
	req.Param("username", consumer.Username)
	req.Param("custom_id", consumer.CustomID)

	var retConsumer models.Consumer
	err := req.ToJSON(&retConsumer)
	if err != nil {
		return nil, err
	}
	return &retConsumer, nil
}

//GetConsumer Retrieve Consumer
//nameOrID -- The unique identifier or the username of the consumer to retrieve
func GetConsumer(nameOrID string) (*models.Consumer, error) {
	//GET /consumers/{username or id}
	if len(nameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the Consumer can not be null")
	}
	req := httplib.Get(kongAdminURL + `/consumers/` + nameOrID)

	var retConsumer models.Consumer
	err := req.ToJSON(&retConsumer)
	if err != nil {
		return nil, err
	}
	return &retConsumer, nil
}

//ListConsumers List Consumers
func ListConsumers(size int, offset string) (*models.ConsumerList, error) {
	// GET /consumers/
	req := httplib.Get(kongAdminURL + `/consumers/`)
	if size > 0 {
		req.Param("size", string(size))
	}
	if len(offset) > 0 {
		req.Param("offset", offset)
	}
	var retConsumerList models.ConsumerList
	err := req.ToJSON(&retConsumerList)
	if err != nil {
		return nil, err
	}
	return &retConsumerList, nil
}

// UpdateConsumer Update Consumer
// usernameOrID -- (required)The unique identifier or the username of the consumer to update
// consumer -- new information
func UpdateConsumer(usernameOrID string, consumer *models.Consumer) (*models.Consumer, error) {
	//PATCH /consumers/{name or id}
	if len(usernameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the Consumer can not be null")
	}
	jsonStr, err := json.Marshal(consumer)
	url := kongAdminURL + `/consumers/` + usernameOrID
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var retConsumer models.Consumer
	err = json.NewDecoder(resp.Body).Decode(&retConsumer)
	if err != nil {
		return nil, err
	}

	return &retConsumer, nil
}

//DeleteConsumer Delete Consumer
func DeleteConsumer(nameOrID string) error {
	//DELETE /consumers/{username or id}
	if len(nameOrID) == 0 {
		return errors.New("The unique identifier or the name of the Consumer can not be null")
	}
	req := httplib.Delete(kongAdminURL + `/consumers/` + nameOrID)

	_, err := req.Response()
	if err != nil {
		return err
	}
	return nil
}
