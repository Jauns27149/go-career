package response

import "go-career/swag/model"

type DetachVolumeResponse struct {
	BlockDevicesMapping []model.BlockDeviceMapping `json:"block_devices_mapping"`
}
