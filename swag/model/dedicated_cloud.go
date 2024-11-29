package model

type DDCHostsShow struct {
	Name      string `bson:"name" json:"name"`
	DDCHostId string `bson:"ddc_host_id" json:"ddc_host_id"`
}
