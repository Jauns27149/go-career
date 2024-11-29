package model

import (
	"time"
)

// 虚机块存储设备
type BlockDevice struct {
	DeleteOnTermination bool       `bson:"delete_on_termination" json:"delete_on_termination"` //在虚机删除时是否删除该块设备
	Id                  string     `bson:"id" json:"id"`                                       //块设备的Id 根据类型来判断具体的id
	VolumeSize          int        `bson:"volume_size" json:"volume_size" `                    //块设备对应卷的大小，单位Gb
	SourceType          string     `bson:"source_type" json:"source_type" `                    //块设备源类型，如image、volume,snapshot等
	VolumeType          string     `bson:"volume_type" json:"volume_type"`
	DestinationType     string     `bson:"destination_type" json:"destination_type" ` //块设备目标类型，如local、volume、blank等
	BootIndex           int        `bson:"boot_index" json:"boot_index" `             //块设备启动顺序
	Deleted             bool       `bson:"deleted" json:"deleted" `                   //块设备是否被标记删除
	CreatedAt           int64      `bson:"created_at" json:"created_at"`              //块设备创建时间
	UpdatedAt           int64      `bson:"updated_at" json:"updated_at" `             //块设备更新时间
	DeletedAt           int64      `bson:"deleted_at" json:"deleted_at"`              //块设备删除时间
	GuestFormat         string     `bson:"guest_format" json:"guest_format"`
	DiskBus             string     `bson:"disk_bus" json:"disk_bus"`
	CacheMode           string     `bson:"cache_mode" json:"cache_mode"`
	DiskIOTune          DiskIOTune `bson:"disk_io_tune" json:"disk_io_tune"`
	Shareable           bool       `bson:"shareable" json:"shareable"` //如果设备支持多挂载，Shareable为true
}

type DiskIOTune struct {
	TotalBytesSec uint64 `bson:"total_bytes_sec" json:"total_bytes_sec"`
	ReadBytesSec  uint64 `bson:"read_bytes_sec" json:"read_bytes_sec"`
	WriteBytesSec uint64 `bson:"write_bytes_sec" json:"write_bytes_sec"`
	TotalIopsSec  uint64 `bson:"total_iops_sec" json:"total_iops_sec"`
	ReadIopsSec   uint64 `bson:"read_iops_sec" json:"read_iops_sec"`
	WriteIopsSec  uint64 `bson:"write_iops_sec" json:"write_iops_sec"`
}

type BlockDeviceMapping struct {
	BlockDevice    `bson:",squash" json:",squash"`
	VolumeId       string           `bson:"volume_id" json:"volume_id"`
	DeviceType     string           `bson:"device_type" json:"device_type"`         //设备类型
	ConnectionInfo ConnectionInfo   `bson:"connection_info" json:"connection_info"` //块设备连接信息
	InstanceId     string           `bson:"instance_id" json:"instance_id"`         //块设备所关联的虚机Id
	DeviceName     string           `bson:"device_name" json:"device_name"`         //设备挂载的名称
	MountDevice    string           `bson:"mount_device" json:"mount_device"`
	AttachmentId   string           `bson:"attachment_id" json:"attachment_id"`
	AddressUnit    uint             `bson:"address_unit" json:"address_unit"` //挂载scsi类型卷需要自定义地址
	SGIO           string           `bson:"sgio" json:"sgio"`
	Encryption     Encryption       `bson:"encryption" json:"encryption"`
	Boot           DomainDeviceBoot `bson:"boot" json:"boot"`
	EnableDpu      bool             `bson:"enable_dpu" json:"enable_dpu"`
	DPUDiskQueue   int              `bson:"dpu_disk_queue" json:"dpu_disk_queue"` //dpu 单盘队列数
}

type DomainDeviceBoot struct {
	Order    uint   `bson:"order" json:"order"`
	LoadParm string `bson:"loadparm" json:"loadparm"`
}

// ConnectionInfo 获取volume的连接参数。根据driver_volume_type的类型选择Data类型
type ConnectionInfo struct {
	Connector        Connector   `bson:"connector" json:"connector"`
	DriverVolumeType string      `bson:"driver_volume_type" json:"driver_volume_type"` //volume类型
	Data             interface{} `bson:"data" json:"data"`
	Serial           string      `bson:"serial" json:"serial"`
	Multiattach      bool        `bson:"multiattach" json:"multiattach"` //卷是否支持多连接
}

// Connector 连接器的相关信息
type Connector struct {
	Platform      string `bson:"platform" json:"platform"`
	Host          string `bson:"host" json:"host"`
	SystemUuid    string `bson:"system uuid" json:"system uuid"` // 暂时不用，使用nvme类型后端存储时需要配置
	DoLocalAttach bool   `bson:"do_local_attach" json:"do_local_attach"`
	OsType        string `bson:"os_type" json:"os_type"`
	MultiPath     bool   `bson:"multi_path" json:"multi_path"`
	Initiator     string `bson:"initiator" json:"initiator"`
}

type ConnectionInfoRbdData struct {
	SecretType     string           `bson:"secret_type" json:"secret_type"`
	Name           string           `bson:"name" json:"name"`
	Encrypted      bool             `bson:"encrypted" json:"encrypted"`
	Keyring        string           `bson:"keyring" json:"keyring"`
	ClusterName    string           `bson:"cluster_name" json:"cluster_name"`
	SecretUUID     string           `bson:"secret_uuid" json:"secret_uuid"`
	QosSpecs       map[string]int64 `bson:"qos_specs" json:"qos_specs"`
	AuthEnabled    bool             `bson:"auth_enabled" json:"auth_enabled"`
	Hosts          []string         `bson:"hosts" json:"hosts"`
	VolumeId       string           `bson:"volume_id" json:"volume_id"`
	Discard        bool             `bson:"discard" json:"discard"`
	AccessMode     string           `bson:"access_mode" json:"access_mode"`
	AuthUserName   string           `bson:"auth_user_name" json:"auth_user_name"`
	Ports          []string         `bson:"ports" json:"ports"`
	EnableIOThread bool             `bson:"enable_iothread" json:"enable_iothread"`
	DevicePath     string           `bson:"device_path" json:"device_path"`
	ConfigFile     string           `bson:"config_file" json:"config_file"`
}

type ConnectionInfoLavaData struct {
	EnableIOThread bool             `bson:"enable_iothread" json:"enable_iothread"`
	Encrypted      bool             `bson:"encrypted" json:"encrypted"`
	Name           string           `bson:"name" json:"name"`
	QosSpecs       map[string]int64 `bson:"qos_specs" json:"qos_specs"`
	VolumeId       string           `bson:"volume_id" json:"volume_id"`
	AccessMode     string           `bson:"access_mode" json:"access_mode"`
	SecretUUID     string           `bson:"secret_uuid" json:"secret_uuid"`
}

type ConnectionInfoIr3Data struct {
	DiskPath string `bson:"disk_path" json:"disk_path"`
}

type ConnectionInfoIscsiData struct {
	AuthPassword          string           `bson:"auth_password" json:"auth_password"`
	TargetDiscovered      bool             `bson:"target_discovered" json:"target_discovered"`
	Encrypted             bool             `bson:"encrypted" json:"encrypted"`
	QosSpecs              map[string]int64 `bson:"qos_specs" json:"qos_specs"`
	TargetIqn             string           `bson:"target_iqn" json:"target_iqn"`
	TargetPortal          string           `bson:"target_portal" json:"target_portal"`
	VolumeId              string           `bson:"volume_id" json:"volume_id"`
	TargetLun             int              `bson:"target_lun" json:"target_lun"`
	AccessMode            string           `bson:"access_mode" json:"access_mode"`
	AuthUsername          string           `bson:"auth_username" json:"auth_username"`
	AuthMethod            string           `bson:"auth_method" json:"auth_method"`
	DevicePath            string           `bson:"device_path" json:"device_path"`
	MultipathId           string           `bson:"multipath_id" json:"multipath_id"`
	TargetPortals         []string         `bson:"target_portals" json:"target_portals"`
	TargetIqns            []string         `bson:"target_iqns" json:"target_iqns"`
	TargetLuns            []int            `bson:"target_luns" json:"target_luns"`
	DiscoveryAuthMethod   string           `bson:"discovery_auth_method" json:"discovery_auth_method"`
	DiscoveryAuthUsername string           `bson:"discovery_auth_username" json:"discovery_auth_username"`
	DiscoveryPassword     string           `bson:"discovery_password" json:"discovery_password"`
}

// Volume 卷设备相关信息.
type Volume struct {
	ID                  string            `bson:"id" json:"id"`         //卷设备id
	Status              string            `bson:"status" json:"status"` //卷设备当前的状态
	Size                int               `bson:"size" json:"size"`     //卷设备的大小（GB）
	AvailabilityZone    string            `bson:"availability_zone" json:"availability_zone"`
	CreatedAt           time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time         `bson:"updated_at" json:"updated_at"`
	Attachments         []Attachment      `bson:"attachments" json:"attachments"` //卷挂载相关参数
	Name                string            `bson:"name" json:"name"`
	Description         string            `bson:"description" json:"description"`
	VolumeType          string            `bson:"volume_type" json:"volume_type"`   //卷设备类型SATA或SSD
	SnapshotID          string            `bson:"snapshot_id" json:"snapshot_id"`   //卷来源的快照id
	SourceVolID         string            `bson:"source_volid" json:"source_volid"` //创建当前卷的另一个块存储卷的ID
	BackupID            *string           `bson:"backup_id" json:"backup_id"`       //从其中恢复卷的备份ID
	Metadata            map[string]string `bson:"metadata" json:"metadata"`         //由用户定义的任意键值对
	UserID              string            `bson:"user_id" json:"user_id"`
	Bootable            string            `bson:"bootable" json:"bootable"`   //代表是不是可引导的卷
	Encrypted           bool              `bson:"encrypted" json:"encrypted"` //表示卷是否加密
	ReplicationStatus   string            `bson:"replication_status" json:"replication_status"`
	ConsistencyGroupID  string            `bson:"consistencygroup_id" json:"consistencygroup_id"`
	Multiattach         bool              `bson:"multiattach" json:"multiattach"`                     //卷是否支持多连接
	VolumeImageMetadata map[string]string `bson:"volume_image_metadata" json:"volume_image_metadata"` //映像元数据项，仅用于从映像创建的卷，或最初从映像创建的卷的快照
	AttachedServers     []string          `bson:"attached_servers" json:"attached_servers"`
	//ConnectInfoMap      map[string]interface{} `bson:"connection_info_map"` //卷的连接信息,未进行具体盘的对应类型解析,仅用来接收存放 cinder 返回结果
	ConnectionInfo ConnectionInfo `bson:"connection_info"` //卷的连接信息，根据不同的盘类型解析 ConnectInfoMap 后存入
}

// Attachment 卷连接记录
type Attachment struct {
	AttachedAt   time.Time `bson:"-" json:"-"`
	AttachmentID string    `bson:"attachment_id" json:"attachment_id"`
	Device       string    `bson:"device" json:"device"`
	HostName     string    `bson:"host_name" json:"host_name"`
	ID           string    `bson:"id" json:"id"`
	ServerID     string    `bson:"server_id" json:"server_id"`
	VolumeID     string    `bson:"volume_id" json:"volume_id"`
}

// Snapshot 包含与快照关联的所有信息(单个volume)
type Snapshot struct {
	Status           string            `bson:"status" json:"status"` //快照当前状态
	Name             string            `bson:"display_name" json:"display_name"`
	Attachments      []string          `bson:"attachments" json:"attachments"` //快照挂载参数
	AvailabilityZone string            `bson:"availability_zone" json:"availability_zone"`
	Bootable         string            `bson:"bootable" json:"bootable"` //表示卷是不是可引导
	CreatedAt        time.Time         `bson:"-" json:"-"`
	Description      string            `bson:"display_description" json:"display_description"`
	VolumeType       string            `bson:"volume_type" json:"volume_type"`
	SnapshotID       string            `bson:"snapshot_id" json:"snapshot_id"` //创建该快照的快照ID。
	VolumeID         string            `bson:"volume_id" json:"volume_id"`     //创建快照的卷ID。
	Metadata         map[string]string `bson:"metadata" json:"metadata"`       //用户自定义键值对
	ID               string            `bson:"id" json:"id"`
	Size             int               `bson:"size" json:"size"` //快照大小（GB）
}

// ImageMember represents sub image for the whole image.
type ImageMember struct {
	ImageId    string `bson:"image_id" json:"image_id"`
	VolumeId   string `bson:"volume_id" json:"volume_id"`
	Mount      string `bson:"mount" json:"mount"`
	Size       int64  `bson:"size" json:"size"`
	DiskFormat string `bson:"disk_format" json:"disk_format"`
}

// Image 镜像相关信息.
type Image struct {
	ID               string                 `bson:"id" json:"id"` //镜像id
	Name             string                 `bson:"name" json:"name"`
	Status           string                 `bson:"status" json:"status"`                     //镜像当前状态queued或active
	Tags             []string               `bson:"tags" json:"tags"`                         //挂载镜像自定义标签
	ContainerFormat  string                 `bson:"container_format" json:"container_format"` //容器格式，有效格式为ami, ari, aki, bare, ovf
	DiskFormat       string                 `bson:"disk_format" json:"disk_format"`           //磁盘格式，有效格式为ami, ari, aki, vhd, vmdk, raw, qcow2, vdi, iso
	MinDiskGigabytes int                    `bson:"min_disk" json:"min_disk"`                 //引导磁盘所需的磁盘空间
	MinRAMMegabytes  int                    `bson:"min_ram" json:"min_ram"`
	Owner            string                 `bson:"owner" json:"owner"`
	Protected        bool                   `bson:"protected" json:"protected"`   //镜像是否可删除
	Visibility       string                 `bson:"visibility" json:"visibility"` //定义哪些用户可以看到此镜像
	Checksum         string                 `bson:"checksum" json:"checksum"`     //校验和镜像相关的参数
	SizeBytes        int64                  `bson:"-" json:"-"`                   //和镜像相关参数大小
	Metadata         map[string]string      `bson:"metadata" json:"metadata"`     //用户自定义键值对
	Properties       map[string]interface{} //和镜像相关的键值对参数
	CreatedAt        string                 `bson:"created_at" json:"created_at"`
	UpdatedAt        string                 `bson:"updated_at" json:"updated_at"`
	File             string                 `bson:"file" json:"file"`                 //镜像的来源路径或获取路径
	Schema           string                 `bson:"schema" json:"schema"`             //表示镜像或镜像实体的json路径
	VirtualSize      int64                  `bson:"virtual_size" json:"virtual_size"` //镜像的虚拟尺寸

	ImageType string        `bson:"image_type" json:"image_type"` //镜像类型
	Members   []ImageMember `bson:"members" json:"members"`       //整机镜像的子镜像信息

}

type Encryption struct {
	ControlLocation      string    `bson:"control_location" json:"control_location"`
	KeySize              int       `bson:"key_size" json:"key_size"`
	Provider             string    `bson:"provider" json:"provider"`
	EncryptionCiphertext string    `bson:"encryption_ciphertext" json:"encryption_ciphertext"`
	EncryptionKeyId      string    `bson:"encryption_key_id" json:"encryption_key_id"`
	Cipher               string    `bson:"cipher" json:"cipher"`
	Cmkuuid              string    `bson:"cmkuuid" json:"cmkuuid"`
	PlainText            string    `bson:"plain_text" json:"plain_text"`
	SecretUuid           string    `bson:"secret_uuid" json:"secret_uuid"`
	Format               string    `bson:"format" json:"format"`
	Secret               Secret    `bson:"secret" json:"secret"`
	CreatedAt            time.Time `bson:"-" json:"-"`
	UpdatedAt            time.Time `bson:"-" json:"-"`
}

type Secret struct {
	Type string `bson:"type" json:"type"`
	Uuid string `bson:"uuid" json:"uuid"`
}

func (bdm BlockDeviceMapping) IsVolume() bool {
	return bdm.BlockDevice.DestinationType == "volume"
}

func (bdm BlockDeviceMapping) IsISCSI() bool {
	return bdm.ConnectionInfo.DriverVolumeType == "iscsi"
}

// Volume Type contains all the information associated with an OpenStack Volume Type.
type VolumeType struct {
	ID          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	// Arbitrary key-value pairs defined by the user.
	ExtraSpecs map[string]string `bson:"extra_specs" json:"extra_specs"`
	// Whether the volume type is publicly visible.
	IsPublic bool `bson:"is_public" json:"is_public"`
	// Qos Spec ID
	QosSpecID string `bson:"qos_specs_id" json:"qos_specs_id"`
	// Volume Type access public attribute
	PublicAccess bool `bson:"os-volume-type-access:is_public" json:"os-volume-type-access:is_public"`
}

// Capabilities represents the information of an individual StoragePool.
type Capabilities struct {
	// The following fields should be present in all storage drivers.
	DriverVersion     string  `bson:"driver_version" json:"driver_version"`
	FreeCapacityGB    float64 `bson:"-" json:"-"`
	StorageProtocol   string  `bson:"storage_protocol" json:"storage_protocol"`
	TotalCapacityGB   float64 `bson:"-" json:"-"`
	VendorName        string  `bson:"vendor_name" json:"vendor_name"`
	VolumeBackendName string  `bson:"volume_backend_name" json:"volume_backend_name"`

	// The following fields are optional and may have empty values depending
	// on the storage driver in use.
	ReservedPercentage       int64   `bson:"reserved_percentage" json:"reserved_percentage"`
	LocationInfo             string  `bson:"location_info" json:"location_info"`
	QoSSupport               bool    `bson:"QoS_support" json:"QoS_support"`
	ProvisionedCapacityGB    float64 `bson:"provisioned_capacity_gb" json:"provisioned_capacity_gb"`
	MaxOverSubscriptionRatio string  `bson:"-" json:"-"`
	ThinProvisioningSupport  bool    `bson:"thin_provisioning_support" json:"thin_provisioning_support"`
	ThickProvisioningSupport bool    `bson:"thick_provisioning_support" json:"thick_provisioning_support"`
	TotalVolumes             int64   `bson:"total_volumes" json:"total_volumes"`
	FilterFunction           string  `bson:"filter_function" json:"filter_function"`
	GoodnessFuction          string  `bson:"goodness_function" json:"goodness_function"`
	Multiattach              bool    `bson:"multiattach" json:"multiattach"`
	SparseCopyVolume         bool    `bson:"sparse_copy_volume" json:"sparse_copy_volume"`
}

// StoragePool represents an individual StoragePool retrieved from the
// schedulerstats API.
type StoragePool struct {
	Name         string       `bson:"name" json:"name"`
	Capabilities Capabilities `bson:"capabilities" json:"capabilities"`
}

type SysLocalInfo struct {
	Size    int64  `bson:"size" json:"size"`
	ImageId string `bson:"image_id" json:"image_id"`
}
