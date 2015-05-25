package ecs

//API基础返回对象
type EcsBaseResponse struct {
	RequestId string `json:"RequestId"` //请求ID
}

//API调用错误返回对象
type ErrorResponse struct {
	EcsBaseResponse
	HostId  string `json:"HostId"`  //服务器HOSTID
	Code    string `json:"Code"`    //错误代码
	Message string `json:"Message"` //错误信息
}

// Region信息的类型
type RegionType struct {
	RegionId  string `json:"RegionId"`  //Region ID
	LocalName string `json:"LocalName"` //Region名称
}

//区域数组对象
type Regions struct {
	Regions []RegionType `json:"Region"`
}

//查询可用区域返回对象
type DescribeRegionsResponse struct {
	EcsBaseResponse
	Response Regions `json:"Regions"`
}

//允许创建的资源类型
type AvailableResourceCreationType struct {
	//资源类型，如下：
	//Instance：支持实例创建
	//Disk：支持磁盘创建
	//VSwitch：支持专有网络创建
	ResourceTypes string `json:"ResourceTypes"`
}

//支持的磁盘种类
type AvailableDiskCategoriesType struct {
	//磁盘种类
	//cloud：支持创建普通云盘和独立普通云盘
	//ephemeral：支持创建本地磁盘
	//ephemeral_ssd：支持创建本地 SSD 盘
	DiskCategories string `json:"DiskCategories"`
}

//可用区信息的类型
type ZoneType struct {
	ZoneId                    string                        `json:"ZoneId"`    //可用区ID
	LocalName                 string                        `json:"LocalName"` //可用区本地语言名
	AvailableResourceCreation AvailableResourceCreationType `json:"AvailableResourceCreation"`
	AvailableDiskCategories   AvailableDiskCategoriesType   `json:"AvailableDiskCategories"`
}

//集群信息的类型
type ClusterType struct {
	ClusterId string `json:"ClusterId"` //集群ID
}

//快照详情的数据类型
type SnapshotType struct {
	SnapshotId     string `json:"SnapshotId"`     //快照ID
	SnapshotName   string `json:"SnapshotName"`   //快照显示名称。如果创建时指定了快照显示名称，则返回
	Description    string `json:"Description"`    //描述信息
	Progress       string `json:"Progress"`       //快照创建进度，单位为百分比
	SourceDiskId   string `json:"SourceDiskId"`   //源磁盘ID，如果快照的源磁盘已经被删除，该字段任旧保留
	SourceDiskSize int    `json:"SourceDiskSize"` //源磁盘容量，GB
	SourceDiskType string `json:"SourceDiskType"` //源磁盘属性，System | Data
	ProductCode    string `json:"ProductCode"`    //从镜像市场继承的产品编号
	CreationTime   string `json:"CreationTime"`   //创建时间。按照ISO8601标准表示，并需要使用UTC时间。
}

type SecurityGroupIdSetType struct {
	SecurityGroupId []string `json:"SecurityGroupId"` //安全组ID
}

type IpAddressSetType struct {
	IpAddress []string `json:"IpAddress"` //IP地址
}

type VpcAttributesType struct {
	VpcId            string           `json:"VpcId"`     //虚拟专有网络Id
	VSwitchId        string           `json:"VSwitchId"` //虚拟交换机Id
	PrivateIpAddress IpAddressSetType `json:"PrivateIpAddress"`
	NatIpAddress     string           `json:"NatIpAddress"` //云产品Ip，用于云产品之间的网络互通
}

type EipAddressAssociateType struct {
	AllocationId       string `json:"AllocationId"`       //弹性公网IP实例Id
	IpAddress          string `json:"IpAddress"`          //弹性公网IP
	Bandwidth          int    `json:"Bandwidth"`          //弹性公网IP的公网带宽限速，默认是5Mbps
	InternetChargeType string `json:"InternetChargeType"` //弹性公网IP的计费方式
}

type OperationLocksType struct {
	//锁定类型
	//financial：因欠费被锁定
	//security：因安全原因被锁定
	LockReason []string `json:"LockReason"`
}

//实例资源规格项的类型
type InstanceTypeItemType struct {
	InstanceTypeId string `json:"InstanceTypeId"` //实例规格的ID
	CpuCoreCount   int    `json:"CpuCoreCount"`   //CPU的内核数目
	MemorySize     int    `json:"MemorySize"`     //内存大小，单位GB
}

//云服务器实例属性
type InstanceAttributesType struct {
	InstanceId              string                  `json:"InstanceId"`              //实例ID
	InstanceName            string                  `json:"InstanceName"`            //实例的显示名称
	Description             string                  `json:"Description"`             //实例的描述
	ImageId                 string                  `json:"ImageId"`                 //镜像ID
	RegionId                string                  `json:"RegionId"`                //实例所属地域ID
	ZoneId                  string                  `json:"ZoneId"`                  //实例所属可用区
	InstanceType            string                  `json:"InstanceType"`            //实例资源规格
	HostName                string                  `json:"HostName"`                //实例机器名称
	Status                  string                  `json:"Status"`                  //实例状态
	SecurityGroupIds        SecurityGroupIdSetType  `json:"SecurityGroupIds"`        //实例所属安全组的集合SecurityGroupIdSetType
	InnerIpAddress          IpAddressSetType        `json:"InnerIpAddress"`          //实例的内网IP地址
	PublicIpAddress         IpAddressSetType        `json:"PublicIpAddress"`         //实例的公网IP地址
	InternetMaxBandwidthIn  int                     `json:"InternetMaxBandwidthIn"`  //公网入带宽最大值
	InternetMaxBandwidthOut int                     `json:"InternetMaxBandwidthOut"` //公网出带宽最大值
	InternetChargeType      string                  `json:"InternetChargeType"`      //网络计费类型，PayByBandwidth | PayByTraffic两个值中的一个。预付费实例显示PayByBandwidth（按带宽计费）
	CreationTime            string                  `json:"CreationTime"`            //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
	VpcAttributes           VpcAttributesType       `json:"VpcAttributes"`           //VpcAttributesType类型
	EipAddress              EipAddressAssociateType `json:"EipAddress"`              //EipAddressAssociateType类型
	InstanceNetworkType     string                  `json:"InstanceNetworkType"`     //实例网络类型，可选值Classic | Vpc
	OperationLocks          OperationLocksType      `json:"OperationLocks"`          //锁定列表
}

//实例数组
type Instances struct {
	Instance []InstanceAttributesType `json:"Instance"`
}

//查询实例列表返回对象
type DescribeInstancesResponse struct {
	EcsBaseResponse
	TotalCount int       `json:"TotalCount"` //总实例数
	PageNumber int       `json:"PageNumber"` //页码
	PageSize   int       `json:"PageSize"`   //单页实例数量
	Instances  Instances `json:"Instances"`  //实例列表
}

//创建实例返回对象
type CreateInstanceResponse struct {
	EcsBaseResponse
	InstanceId string `json:"InstanceId"` //实例ID
}

//分配公网IP返回对象
type AllocatePublicIpAddressResponse struct {
	EcsBaseResponse
	IpAddress string `json:"IpAddress"` //公网IP地址
}
