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

// 安全组相关操作

// 创建安全组请求
type CreateSecurityGroupRequest struct {
	Action            string
	RegionId          string
	SecurityGroupName string
	Description       string
	VpcId             string
}

func (r *CreateSecurityGroupRequest) AddToParams(params *util.OrderedParams) error {
	r.Action = "CreateSecurityGroup"
	params.Add("Action", r.Action)
	if r.RegionId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "RegionId")
	}
	params.Add("RegionId", r.RegionId)
	AddNotNullFieldToParams(params, r.SecurityGroupName, "SecurityGroupName")
	AddNotNullFieldToParams(params, r.Description, "Description")
	AddNotNullFieldToParams(params, r.VpcId, "VpcId")
	return nil
}

// 删除安全组请求
type DeleteSecurityGroupRequest struct {
	Action          string
	RegionId        string
	SecurityGroupId string
}

func (r *DeleteSecurityGroupRequest) AddToParams(params *util.OrderedParams) error {
	r.Action = "DeleteSecurityGroup"
	params.Add("Action", r.Action)
	if r.RegionId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "RegionId")
	}
	if r.SecurityGroupId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "SecurityGroupId")
	}
	params.Add("RegionId", r.RegionId)
	params.Add("SecurityGroupId", r.SecurityGroupId)
	return nil
}

// 授权安全组In方向的访问权限
type AuthorizeSecurityGroupRequest struct {
	Action          string
	SecurityGroupId string
	RegionId        string
	// IP协议，取值：tcp | udp | icmp | gre | all；all表示同时支持四种协议
	IpProtocol string
	// IP协议相关的端口号范围
	// 协议为tcp、udp时默认端口号，取值范围为1~65535；例如“1/200”意思是端口号范围为1~200，若输入值为：“200/1”接口调用将报错。
	// 协议为icmp时端口号范围值为-1/-1；
	// gre协议时端口号范围值为-1/-1；
	// 协议为all时端口号范围值为-1/-1
	PortRange               string
	SourceGroupId           string
	SourceGroupOwnerAccount string
	// 源IP地址范围（采用CIDR格式来指定IP地址范围），默认值为0.0.0.0/0（表示不受限制），其他支持的格式如10.159.6.18/12或10.159.6.186。仅支持IPV4。
	SourceCidrIp string
	// 授权策略，参数值可为：accept（接受访问），drop (拒绝访问) 默认值为：accept
	Policy string
	// 授权策略优先级，参数值可为：1-100 默认值为：1
	Priority string
	// 网络类型，取值：
	// internet
	// intranet；
	// 默认值为internet
	// 当对安全组进行相互授权时（即指定了SourceGroupId且没有指定SourceCidrIp），必须指定NicType为intranet
	NicType string
}

func (r *AuthorizeSecurityGroupRequest) Validate() error {
	if r.SecurityGroupId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "SecurityGroupId")
	}
	if r.RegionId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "RegionId")
	}
	if r.IpProtocol == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "IpProtocol")
	}
	if r.PortRange == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "PortRange")
	}
	return nil
}

func (r *AuthorizeSecurityGroupRequest) AddToParams(params *util.OrderedParams) error {
	r.Action = "AuthorizeSecurityGroup"
	if err := r.Validate(); err != nil {
		return err
	}
	params.Add("Action", r.Action)
	params.Add("RegionId", r.RegionId)
	params.Add("SecurityGroupId", r.SecurityGroupId)
	params.Add("IpProtocol", r.IpProtocol)
	params.Add("PortRange", r.PortRange)
	AddNotNullFieldToParams(params, r.SourceGroupId, "SourceGroupId")
	AddNotNullFieldToParams(params, r.SourceGroupOwnerAccount, "SourceGroupOwnerAccount")
	AddNotNullFieldToParams(params, r.SourceCidrIp, "SourceCidrIp")
	AddNotNullFieldToParams(params, r.Policy, "Policy")
	AddNotNullFieldToParams(params, r.Priority, "Priority")
	AddNotNullFieldToParams(params, r.NicType, "NicType")
	return nil
}

// 撤销安全组授权规则请求
type RevokeSecurityGroupRequest AuthorizeSecurityGroupRequest

func (r *RevokeSecurityGroupRequest) Validate() error {
	if r.SecurityGroupId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "SecurityGroupId")
	}
	if r.RegionId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "RegionId")
	}
	if r.IpProtocol == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "IpProtocol")
	}
	if r.PortRange == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "PortRange")
	}
	return nil
}

func (r *RevokeSecurityGroupRequest) AddToParams(params *util.OrderedParams) error {
	r.Action = "RevokeSecurityGroup"
	if err := r.Validate(); err != nil {
		return err
	}
	params.Add("Action", r.Action)
	params.Add("RegionId", r.RegionId)
	params.Add("SecurityGroupId", r.SecurityGroupId)
	params.Add("IpProtocol", r.IpProtocol)
	params.Add("PortRange", r.PortRange)
	AddNotNullFieldToParams(params, r.SourceGroupId, "SourceGroupId")
	AddNotNullFieldToParams(params, r.SourceGroupOwnerAccount, "SourceGroupOwnerAccount")
	AddNotNullFieldToParams(params, r.SourceCidrIp, "SourceCidrIp")
	AddNotNullFieldToParams(params, r.Policy, "Policy")
	AddNotNullFieldToParams(params, r.Priority, "Priority")
	AddNotNullFieldToParams(params, r.NicType, "NicType")
	return nil
}

// 授权安全组In方向的访问权限
type AuthorizeSecurityGroupEgressRequest struct {
	Action          string
	SecurityGroupId string
	RegionId        string
	// IP协议，取值：tcp | udp | icmp | gre | all；all表示同时支持四种协议
	IpProtocol string
	// IP协议相关的端口号范围
	// 协议为tcp、udp时默认端口号，取值范围为1~65535；例如“1/200”意思是端口号范围为1~200，若输入值为：“200/1”接口调用将报错。
	// 协议为icmp时端口号范围值为-1/-1；
	// gre协议时端口号范围值为-1/-1；
	// 协议为all时端口号范围值为-1/-1
	PortRange             string
	DestGroupId           string
	DestGroupOwnerAccount string
	// 目标IP地址范围（采用CIDR格式来指定IP地址范围），默认值为0.0.0.0/0（表示不受限制），其他支持的格式如10.159.6.18/12或10.159.6.186。仅支持IPV4。
	DestCidrIp string
	// 授权策略，参数值可为：accept（接受访问），drop (拒绝访问) 默认值为：accept
	Policy string
	// 授权策略优先级，参数值可为：1-100 默认值为：1
	Priority string
	// 网络类型，取值：
	// internet
	// intranet；
	// 默认值为internet
	// 当对安全组进行相互授权时（即指定了SourceGroupId且没有指定SourceCidrIp），必须指定NicType为intranet
	NicType string
}

func (r *AuthorizeSecurityGroupEgressRequest) Validate() error {
	if r.SecurityGroupId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "SecurityGroupId")
	}
	if r.RegionId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "RegionId")
	}
	if r.IpProtocol == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "IpProtocol")
	}
	if r.PortRange == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "PortRange")
	}
	return nil
}

func (r *AuthorizeSecurityGroupEgressRequest) AddToParams(params *util.OrderedParams) error {
	r.Action = "AuthorizeSecurityGroupEgress"
	if err := r.Validate(); err != nil {
		return err
	}
	params.Add("Action", r.Action)
	params.Add("RegionId", r.RegionId)
	params.Add("SecurityGroupId", r.SecurityGroupId)
	params.Add("IpProtocol", r.IpProtocol)
	params.Add("PortRange", r.PortRange)
	AddNotNullFieldToParams(params, r.DestGroupId, "DestGroupId")
	AddNotNullFieldToParams(params, r.DestGroupOwnerAccount, "DestGroupOwnerAccount")
	AddNotNullFieldToParams(params, r.DestCidrIp, "DestCidrIp")
	AddNotNullFieldToParams(params, r.Policy, "Policy")
	AddNotNullFieldToParams(params, r.Priority, "Priority")
	AddNotNullFieldToParams(params, r.NicType, "NicType")
	return nil
}

// 撤销安全组Out方向的访问规则
type RevokeSecurityGroupEgressRequest AuthorizeSecurityGroupEgressRequest

func (r *RevokeSecurityGroupEgressRequest) Validate() error {
	if r.SecurityGroupId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "SecurityGroupId")
	}
	if r.RegionId == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "RegionId")
	}
	if r.IpProtocol == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "IpProtocol")
	}
	if r.PortRange == "" {
		return fmt.Errorf(fieldCanNotNullErrMsg, "PortRange")
	}
	return nil
}

func (r *RevokeSecurityGroupEgressRequest) AddToParams(params *util.OrderedParams) error {
	r.Action = "RevokeSecurityGroupEgress"
	if err := r.Validate(); err != nil {
		return err
	}
	params.Add("Action", r.Action)
	params.Add("RegionId", r.RegionId)
	params.Add("SecurityGroupId", r.SecurityGroupId)
	params.Add("IpProtocol", r.IpProtocol)
	params.Add("PortRange", r.PortRange)
	AddNotNullFieldToParams(params, r.DestGroupId, "DestGroupId")
	AddNotNullFieldToParams(params, r.DestGroupOwnerAccount, "DestGroupOwnerAccount")
	AddNotNullFieldToParams(params, r.DestCidrIp, "DestCidrIp")
	AddNotNullFieldToParams(params, r.Policy, "Policy")
	AddNotNullFieldToParams(params, r.Priority, "Priority")
	AddNotNullFieldToParams(params, r.NicType, "NicType")
	return nil
}
