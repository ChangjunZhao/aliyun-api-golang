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

type SecurityGroupSetType struct {
	SecurityGroup SecurityGroupItemType //安全组SecurityGroupItemType
}

type SecurityGroupItemType struct {
	SecurityGroupId   string //安全组ID
	SecurityGroupName string //安全组名称
	Description       string //描述信息
	VpcId             string //安全组所属的专有网络
	CreationTime      string //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

//IP段集合的类型
type IpRangeSetType struct {
	IpAddress string //采用CIDR格式来指定IP地址范围。
	NicType   string //网络类型internet | intranet的一种
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

//包含Vpc信息的集合
type VpcSetType struct {
	VpcId        string //	VpcId
	RegionId     string //	VPC所在的地域
	Status       string //	VPC状态，包括Pending和Available两种
	VpcName      string //	VPC名称，不填则为空，默认值为空，[2,128]英文或中文字符，必须以大小字母或中文开头，可包含数字，”_”或”-”，这个值会展示在控制台。不能以http:// 和https:// 开头。
	VSwitchIds   string //	VSwitchId列表
	CidrBlock    string //	VPC的网段地址
	VRouterId    string //	VRouter的Id
	Description  string //	描述，不填则为空，默认值为空，[2,256]英文或中文字符，不能以http:// 和https:// 开头。
	CreationTime string //	创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

//包含虚拟路由器信息的集合
type VRouterSetType struct {
	VRouterId     string //	虚拟路由器的Id
	RegionId      string //	地域Id
	VpcId         string //	专有网络Id
	RouteTableIds string //	虚拟路由表Id列表
	VRouterName   string //	虚拟路由器名称
	Description   string //	虚拟路由器的描述，不填则为空，默认值为空，[2,256]英文或中文字符，不能以http:// 和https:// 开头。
	CreationTime  string //	创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

//包含路由表数据的集合
type RouteTableSetType struct {
	VRouterId      string              //		虚拟路由器的ID
	RouteTableId   string              //		路由表的ID
	RouteEntrys    []RouteEntrySetType //路由条目详情RouteEntrySetType组成的集合
	RouteTableType string              //		路由表类型，System | Custom 二者选其一
	CreationTime   string              //		创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

//包含路由条目数据的集合
type RouteEntrySetType struct {
	RouteTableId         string //路由条目所在的虚拟路由器
	DestinationCidrBlock string //目标网段地址
	Type                 string //路由类型(System | Custom)
	NextHopId            string //下一跳实例ID
	Status               string //路由条目状态Pending | Available | Modifying
}

//包含虚拟交换机信息的集合
type VSwitchSetType struct {
	VSwitchId               string //	虚拟交换机ID
	VpcId                   string //	虚拟交换机所在的专有网络
	Status                  string //	虚拟交换机状态，包括Pending和Available两种
	CidrBlock               string //	虚拟交换机的地址
	ZoneId                  string //	虚拟交换机所在的可用区
	AvailableIpAddressCount int    //	虚拟交换机当前可用的IP地址数量
	Description             string //	描述，不填则为空，默认值为空，[2,256]英文或中文字符，不能以http:// 和https:// 开头。
	VSwitchName             string //	虚拟交换机名字，不填则为空，默认值为空，[2,128]英文或中文字符，必须以大小字母或中文开头，可包含数字，”_”或”-”，这个值会展示在控制台。不能以http:// 和https:// 开头。
	CreationTime            string //	创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

type EipAddressAssociateType struct {
	AllocationId       string `json:"AllocationId"`       //弹性公网IP实例Id
	IpAddress          string `json:"IpAddress"`          //弹性公网IP
	Bandwidth          int    `json:"Bandwidth"`          //弹性公网IP的公网带宽限速，默认是5Mbps
	InternetChargeType string `json:"InternetChargeType"` //弹性公网IP的计费方式
}

//包含弹性公网IP信息的集合
type EipAddressSetType struct {
	RegionId           string             //	弹性公网IP所在的地域
	IpAddress          string             //	弹性公网IP
	AllocationId       string             //	弹性公网IP实例Id
	Status             string             //	弹性公网IP当前的状态，包括Associating、Unassociating、InUse和Available
	InstanceId         string             //	弹性公网IP当前绑定的实例，如果未绑定则值为空。
	Bandwidth          int                //	弹性公网IP的公网带宽限速，默认是5Mbps
	InternetChargeType string             //	弹性公网IP的计费方式。
	OperationLocks     OperationLocksType //LockReason组成的字符串数组，如果没有被锁定则其子节点不出现。
	AllocationTime     string             //	分配时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

//包含弹性公网IP监控数据的类型
type EipMonitorDataType struct {
	EipRX        int    //	一段时间（Period）内，EIP接收到的数据流量，单位：kbytes。
	EipTX        int    //	一段时间（Period）内，EIP接发送的数据流量，单位：kbytes。
	EipFlow      int    //	一段时间（Period）内，EIP网络流量，单位Kbytes。
	EipBandwidth int    //	弹性公网IP的带宽（单位时间内的网络流量），单位为kbytes/s。
	EipPackets   int    //	一段时间（Period）内，EIP接受和发送的报文总数。
	TimeStamp    string //	查询流量的时间点，按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

//自动快照策略类型，自动快照策略的详细设置信息。
type AutoSnapshotPolicyType struct {
	SystemDiskPolicyEnabled string //系统盘自动快照策略开关：true：该磁盘属性的磁盘打自动快照 false：不打自动快照
	//系统盘自动快照策略的时间段
	//4个时间段：
	//1：1:00-7:00
	//2：7:00-13:00
	//3：13:00-19:00
	//4：19:00-1:00
	SystemDiskPolicyTimePeriod int

	SystemDiskPolicyRetentionDays     int    //系统盘自动快照策略的保留天数 可选值， 1，2，3
	SystemDiskPolicyRetentionLastWeek string //系统盘自动快照策略的保留上周日选项： true：代表保留上周日的快照 false：不保留
	DataDiskPolicyEnabled             string //数据盘自动快照策略开关： true：该磁盘属性的磁盘打自动快照 false：不打自动快照
	//数据盘自动快照策略的时间段
	//4个时间段：
	//1：1:00-7:00
	//2：7:00-13:00
	//3：13:00-19:00
	//4：19:00-1:00
	DataDiskPolicyTimePeriod        int
	DataDiskPolicyRetentionDays     int    //数据盘自动快照策略的保留天数 可选值， 1，2，3
	DataDiskPolicyRetentionLastWeek string //数据盘自动快照策略的保留上周日选项：true：代表保留上周日的快照 false：不保留
}

//自动快照执行状态类型，返回上一次的执行结果。
type AutoSnapshotExecutionStatusType struct {
	/*
	   返回最近一次执行的状态：Standby|Executed|Failed
	   Standby：刚设置完成还未开始执行或者系统盘的策略被关闭
	   Executed：执行成功
	   Failed：执行失败
	*/
	SystemDiskExecutionStatus string
	/*
	   返回最近一次执行的状态：Standby|Executed|Failed
	   Standby：刚设置完成还未开始执行或者数据盘的策略被关闭
	   Executed：执行成功
	   Failed：执行失败
	*/
	DataDiskExecutionStatus string
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

type DiskItemType struct {
	DiskId             string             // 磁盘ID
	RegionId           string             // 	磁盘所属的地域 ID
	ZoneId             string             // 磁盘所属的可用区ID
	DiskName           string             // 磁盘名
	Description        string             //磁盘描述
	Type               string             //磁盘类型 可选值：system: 系统盘 data: 数据盘
	Category           string             //磁盘种类 可选值： cloud: 普通云盘 ephemeral: 临时磁盘 ephemeral_ssd: 临时SSD盘
	Size               int                //磁盘大小，单位GB
	ImageId            string             //创建磁盘的镜像ID，只有通过镜像创建的磁盘才有值，否则为空。这个值在磁盘的生命周期内始终不变
	SourceSnapshotId   string             //创建磁盘使用的快照，如果创建磁盘时，没有指定快照，则为空。这个值在磁盘的生命周期内始终不变。
	ProductCode        string             //镜像市场的商品标识
	Portable           string             //磁盘是否可卸载 true代表是独立普通云盘，可以独立存在且可以自由在可用区内挂载和下载 false代表非独立普通云盘，只能和实例同生同灭。用户如果需要做attach和detach操作，必须先查询一下这个属性为true的磁盘才能操作。临时磁盘，临时SSD盘，普通云盘的系统盘和包月的普通云盘，该属性都为false。这个属性用户不能更改。
	Status             string             //磁盘状态 In_use | Available | Attaching | Detaching | Creating | ReIniting
	OperationLocks     OperationLocksType //磁盘锁定原因类型
	InstanceId         string             //所属Instance ID 只有在Status为In_use时才有值，其他状态为空。
	Device             string             //所属Instance的Device信息：比如/dev/xvdb 只有在Status为In_use是才有值，其他状态为空。
	DeleteWithInstance string             //True表示Instance释放时，这块磁盘随Instance一起释放；false表示Instance释放时，这块磁盘保留不释放。
	DeleteAutoSnapshot string             //是否同时删除自动快照，true | false。通过CreateSnapshot或者在控制台创建的快照，不受这个参数的影响，始终会被保留。
	EnableAutoSnapshot string             //磁盘是否执行自动快照策略，true | false。true表示这块磁盘执行自动快照策略，false表示这块磁盘不执行自动快照策略
	CreationTime       string             //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
	AttachedTime       string             //挂载时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ 只有在Status为In_use时才有意义
	DetachedTime       string             //卸载时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ只有在Status为Available时才有意义
}

type DiskSetType struct {
	Disk DiskItemType
}

type ImageType struct {
	ImageId            string            //镜像编码
	ImageVersion       string            //镜像版本
	Architecture       string            //镜像系统类型：i386 | x86_64
	ImageName          string            //镜像的名称
	Description        string            //描述信息
	Size               int               //镜像大小
	ImageOwnerAlias    string            //镜像所有者别名 有效值：system – 系统公共镜像 self – 用户的自定义镜像 others – 其他用户的公开镜像 marketplace -镜像市场镜像
	OSName             string            //操作系统的显示名称
	DiskDeviceMappings DiskDeviceMapping //镜像下包含磁盘和快照的系统描述
	ProductCode        string            //镜像市场的镜像商品标示
	IsSubscribed       string            //用户是否订阅了该镜像的ProductCode对应的镜像商品的服务条款. true：表示已经订阅 false：表示未订阅
	Progress           string            //镜像完成的进度，单位为百分比
	Status             string            //镜像的状态，可能的值有：UnAvailable 不可用 Available 可用 Creating 创建中 CreateFailed 创建失败
	CreationTime       string            //创建时间。按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

type DiskDeviceMapping struct {
	SnapshotId string //快照ID
	Size       string //生成磁盘的大小
	Device     string //生成磁盘的Device信息：比如/dev/xvdb
}

//包含磁盘监控数据的类型
type DiskMonitorDataType struct {
	DiskId    string //	磁盘编号
	IOPSRead  int    //	磁盘IO读操作，单位：次/s
	IOPSWrite int    //	磁盘IO写操作，单位：次/s
	IOPSTotal int    //	磁盘IO读写总操作，单位：次/s
	BPSRead   int    //	磁盘读带宽，单位：byte/s
	BPSWrite  int    //	磁盘写带宽，单位：byte/s
	BPSTotal  int    //	磁盘读写总带宽，单位：byte/s
	TimeStamp string //	查询的时间点，按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

type InstanceMonitorDataType struct {
	InstanceId        string //实例ID
	CPU               int    //CPU的使用比例，单位：百分比（%）
	IntranetRX        int    //云服务器实例接收到的数据流量，单位：kbits
	IntranetTX        int    //云服务器实例接发送的数据流量，单位：kbits
	IntranetBandwidth int    //云服务器实例的带宽（单位时间内的网络流量），单位为kbits/s
	InternetRX        int    //云服务器实例接收到的数据流量，单位：kbits
	InternetTX        int    //云服务器实例接发送的数据流量，单位：kbits
	InternetBandwidth int    //云服务器实例的带宽（单位时间内的网络流量），单位为kbits/s
	IOPSRead          int    //系统盘IO读操作，单位：次/s
	IOPSWrite         int    //系统盘IO写操作，单位：次/s
	BPSRead           int    //系统盘磁盘读带宽，单位：Byte/s
	BPSWrite          int    //系统盘磁盘写带宽，单位：Byte/s
	TimeStamp         string //查询流量的时间点，按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

//镜像信息的类型
type AccountType struct {
	AliyunId string //阿里云账号Id
}

//共享组类型
type ShareGroupType struct {
	Group string //	共享分组
}

//实例状态的类型
type InstanceStatusItemType struct {
	InstanceId string //实例ID
	Status     string //实例状态
}

//实例状态的项的集合
type InstanceStatusSetType struct {
	InstanceStatus InstanceStatusItemType //由InstanceStatusItemType组成的集合
}

//安全组规则组类型集合
type PermissionSetType struct {
	Permision PermissionType //安全组规则PermissionType
}

//安全组规则类型
type PermissionType struct {
	IpProtocol              string //授权指定的IP协议
	PortRange               string //授权指定的端口范围
	SourceCidrIp            string //授权给指定IP地址段
	SourceGroupId           string //源安全组编码
	SourceGroupOwnerAccount string //源安全组所属阿里云账户
	DestCidrIp              string //授权访问指定IP地址段
	DestGroupId             string //目标安全组编码
	DestGroupOwnerAccount   string //目标安全组所属阿里云账户
	Policy                  string //授权策略
	NicType                 string //网络类型
	Priority                string //规则优先级
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
