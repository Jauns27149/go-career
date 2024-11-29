package response

import "go-career/swag/model"

type DedicatedCloud struct {
	Id        string   `json:"id"`         //专属云ID
	ProjectId string   `json:"project_id"` //租户id
	Name      string   `json:"name"`       //专属云名称
	CreatedAt int64    `json:"created_at"` //创建时间
	UpdatedAt int64    `json:"updated_at"` //更新时间
	UpdatedBy string   `json:"updated_by"` //更新时间
	Hosts     []string `json:"hosts"`
}

type DedicatedCloudShow struct {
	Id        string               `json:"id"`         //专属云ID
	ProjectId string               `json:"project_id"` //租户id
	Name      string               `json:"name"`       //专属云名称
	CreatedAt int64                `json:"created_at"` //创建时间
	UpdatedAt int64                `json:"updated_at"` //更新时间
	UpdatedBy string               `json:"updated_by"` //更新时间
	Hosts     []model.DDCHostsShow `json:"hosts"`      //展示宿主机name和ddc_host_id
}
