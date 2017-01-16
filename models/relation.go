package models

type ClRelation struct {
	RelationId string `orm:"pk;column(relation_id)" json:"relationId"`
	UserName string `orm:"column(user_name)" json:"userName"`
	Service *ClService `orm:"rel(fk)"`
	ApiKey string `orm:"column(api_key)" json:"apiKey"`
	ConsumerId string `orm:"column(consumer_id)" json:"consumerId"`
}