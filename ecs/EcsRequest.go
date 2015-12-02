package ecs

import (
	"fmt"
	"github.com/ChangjunZhao/aliyun-api-golang/util"
)

type DescribeInstancesRequest struct {
	Action              string
	RegionId            string
	VpcId               string
	VSwitchId           string
	ZoneId              string
	InstanceIds         string
	InstanceNetworkType string
	PrivateIpAddresses  string
	InnerIpAddresses    string
	PublicIpAddresses   string
	SecurityGroupIds    string
	InstanceChargeType  string
	InternetChargeType  string
	InstanceName        string
	ImageId             string
	Status              string
	DeviceAvailable     string
	IoOptimized         string
	PageNumber          int
	PageSize            int
}

func (r *DescribeInstancesRequest) AddToParams(params *util.OrderedParams) error {
	r.Action = "DescribeInstances"
	params.Add("Action", r.Action)
	if r.RegionId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "RegionId")
	} else {
		params.Add("RegionId", r.RegionId)
	}
	AddNotNullFieldToParams(params, r.VpcId, "VpcId")
	AddNotNullFieldToParams(params, r.InstanceIds, "InstanceIds")
	return nil
}

// Create instance http request object
type CreateInstanceRequest struct {
	Action                  string
	RegionId                string
	ZoneId                  string
	ImageId                 string
	InstanceType            string
	SecurityGroupId         string
	InstanceName            string
	Description             string
	InternetChargeType      string
	InternetMaxBandwidthIn  string
	InternetMaxBandwidthOut string
	HostName                string
	Password                string
	IoOptimized             string
	SystemDiskCategory      string
	SystemDiskDiskName      string
	SystemDiskDescription   string
	VSwitchId               string
	PrivateIpAddress        string
}

func (r *CreateInstanceRequest) Validate() error {
	if r.RegionId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "RegionId")
	}
	if r.ImageId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "ImageId")
	}
	if r.InstanceType == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "InstanceType")
	}
	if r.SecurityGroupId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "SecurityGroupId")
	}
	if r.Password == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "Password")
	}
	return nil
}

const (
	fieldCanNotNullErrMsg = "%s can not null."
)

func (r *CreateInstanceRequest) AddToParams(params *util.OrderedParams) error {
	r.Action = "CreateInstance"
	if err := r.Validate(); err != nil {
		return err
	}
	params.Add("Action", r.Action)
	params.Add("RegionId", r.RegionId)
	params.Add("ImageId", r.ImageId)
	params.Add("InstanceType", r.InstanceType)
	params.Add("SecurityGroupId", r.SecurityGroupId)
	params.Add("Password", r.Password)
	AddNotNullFieldToParams(params, r.ZoneId, "ZoneId")
	AddNotNullFieldToParams(params, r.InstanceName, "InstanceName")
	AddNotNullFieldToParams(params, r.Description, "Description")
	AddNotNullFieldToParams(params, r.InternetChargeType, "InternetChargeType")
	if r.InternetChargeType == "PayByBandwidth" {
		AddNotNullFieldToParams(params, r.InternetMaxBandwidthIn, "InternetMaxBandwidthIn")
		AddNotNullFieldToParams(params, r.InternetMaxBandwidthOut, "InternetMaxBandwidthOut")
	} else {
		AddNotNullFieldToParams(params, r.InternetMaxBandwidthOut, "InternetMaxBandwidthOut")
	}
	AddNotNullFieldToParams(params, r.HostName, "HostName")
	AddNotNullFieldToParams(params, r.IoOptimized, "IoOptimized")
	AddNotNullFieldToParams(params, r.SystemDiskCategory, "SystemDisk.Category")
	AddNotNullFieldToParams(params, r.SystemDiskDiskName, "SystemDisk.DiskName")
	AddNotNullFieldToParams(params, r.SystemDiskDescription, "SystemDisk.Description")
	AddNotNullFieldToParams(params, r.VSwitchId, "VSwitchId")
	AddNotNullFieldToParams(params, r.PrivateIpAddress, "PrivateIpAddress")
	return nil
}

func AddNotNullFieldToParams(params *util.OrderedParams, value string, fieldName string) {
	if value != "" {
		params.Add(fieldName, value)
	}
}
