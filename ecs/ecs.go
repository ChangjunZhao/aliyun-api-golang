// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 阿里云ECS API go语言版本
package ecs

import (
	"github.com/ChangjunZhao/aliyun-api-golang/signer"
	"github.com/ChangjunZhao/aliyun-api-golang/util"
	"math/rand"
	"strconv"
	"time"
)

//定义常量
const (
	API_SERVER                 = "http://ecs.aliyuncs.com/"
	VERSION                    = "2014-05-26"  //API版本
	SIGNATURE_VERSION          = "1.0"         //签名版本
	SIGNATURE_METHOD_HMAC_SHA1 = "HMAC-SHA1"   //HMAC-SHA1签名
	ACCESS_KEY_ID_PARAM        = "AccessKeyId" //access key id
	SIGNATURE_VERSION_PARAM    = "SignatureVersion"
	NONCE_PARAM                = "SignatureNonce"
	SIGNATURE_METHOD_PARAM     = "SignatureMethod"
	SIGNATURE_PARAM            = "Signature"
	TIMESTAMP_PARAM            = "Timestamp"
	VERSION_PARAM              = "Version"
)

//调用API的Client
type Client struct {
	accessKeyId    string
	debug          bool
	nonceGenerator nonceGenerator
	signer         *signer.SHA1Signer //签名类
}

//创建新的客户端
//
//使用方法：
//
//c = NewClient("Access Key ID","Access Key Secret")
func NewClient(accessKeyId string, accessKeySecret string) *Client {
	return &Client{
		accessKeyId:    accessKeyId,
		signer:         signer.NewSigner(accessKeySecret),
		nonceGenerator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *Client) Debug(enabled bool) {
	c.debug = enabled
}

// 查询实例列表
//
// regionId 地域ID,如cn-beijing
//
// 返回值InstanceAttributesType数组及错误信息
func (c *Client) DescribeInstances(regionId string) ([]InstanceAttributesType, error) {
	params := c.baseParams(c.accessKeyId, nil)
	params.Add("Format", "JSON")
	params.Add("Action", "DescribeInstances")
	params.Add("RegionId", regionId)
	var describeInstancesResponse DescribeInstancesResponse
	err := util.CallApiServer(API_SERVER, c.signer, params, &describeInstancesResponse)
	if err == nil {
		return describeInstancesResponse.Instances.Instance, nil
	} else {
		return nil, err
	}
}

//查询实例信息
//
//instanceId :实例ID
//
//返回值：InstanceAttributesType 实例对象
func (c *Client) DescribeInstanceAttribute(instanceId string) (*InstanceAttributesType, error) {
	params := c.baseParams(c.accessKeyId, nil)
	params.Add("Format", "JSON")
	params.Add("Action", "DescribeInstanceAttribute")
	params.Add("InstanceId", instanceId)
	var instanceAttributesType InstanceAttributesType
	err := util.CallApiServer(API_SERVER, c.signer, params, &instanceAttributesType)
	if err == nil {
		return &instanceAttributesType, nil
	} else {
		return nil, err
	}
}

/*
给一个特定实例分配一个可用公网IP地址。

实例的状态必须为 Running 或 Stopped 状态，才可以调用此接口。

分配的 IP 必须在实例启动或重启后才能生效。

分配的时候只能是 IP，不能是 IP 段。

目前，一个实例只能分配一个 IP。当调用此接口时，如果实例已经拥有一个公网 IP，将直接返回原 IP 地址。

被安全控制在实例的 OperationLocks 中标记了 "LockReason" : "security" 的锁定状态时，不能分配公网 IP。
*/
func (c *Client) AllocatePublicIpAddress(instanceId string) (string, error) {
	params := c.baseParams(c.accessKeyId, nil)
	params.Add("Format", "JSON")
	params.Add("Action", "AllocatePublicIpAddress")
	params.Add("InstanceId", instanceId)
	var allocatePublicIpAddress AllocatePublicIpAddressResponse
	err := util.CallApiServer(API_SERVER, c.signer, params, &allocatePublicIpAddress)
	if err == nil {
		return allocatePublicIpAddress.IpAddress, nil
	} else {
		return "", err
	}
}

//启动一个指定的实例
//
//接口调用成功后实例进入 Starting 状态。
//
//实例状态必须为 Stopped，才可以调用该接口。
//
//被安全控制在实例的 OperationLocks 中标记了 "LockReason" : "security" 的锁定状态时，不能启动实例。
func (c *Client) StartInstance(instanceId string) error {
	params := c.baseParams(c.accessKeyId, nil)
	params.Add("Format", "JSON")
	params.Add("Action", "StartInstance")
	params.Add("InstanceId", instanceId)
	var response EcsBaseResponse
	err := util.CallApiServer(API_SERVER, c.signer, params, &response)
	if err == nil {
		return nil
	} else {
		return err
	}
}

/*
重启指定的实例

只有状态为 Running 的实例才可以进行此操作。

接口调用成功后实例进入 Starting 状态。

支持强制重启，强制重启等同于传统服务器的断电重启，可能丢失实例操作系统中未写入磁盘的数据。

被安全控制在实例的 OperationLocks 中标记了 "LockReason" : "security" 的锁定状态时，不能重启实例。
*/
func (c *Client) RebootInstance(instanceId string, forceStop string) error {
	params := c.baseParams(c.accessKeyId, nil)
	params.Add("Format", "JSON")
	params.Add("Action", "RebootInstance")
	params.Add("InstanceId", instanceId)
	params.Add("ForceStop", forceStop)
	var response EcsBaseResponse
	err := util.CallApiServer(API_SERVER, c.signer, params, &response)
	if err == nil {
		return nil
	} else {
		return err
	}
}

/*
停止一个指定的实例。

只有状态为 Running 的实例才可以进行此操作。
接口调用成功后实例进入 Stopping 状态。系统后台会在实例实际 Stop 成功后进入 Stopped 状态。
实例支持强制停止，强制停止等同于断电处理，可能丢失实例操作系统中未写入磁盘的数据。
被安全控制在实例的 OperationLocks 中标记了 "LockReason" : "security" 的锁定状态时，不能停止实例。
*/
func (c *Client) StopInstance(instanceId string, forceStop string) error {
	params := c.baseParams(c.accessKeyId, nil)
	params.Add("Format", "JSON")
	params.Add("Action", "StopInstance")
	params.Add("InstanceId", instanceId)
	params.Add("ForceStop", forceStop)
	var response EcsBaseResponse
	err := util.CallApiServer(API_SERVER, c.signer, params, &response)
	if err == nil {
		return nil
	} else {
		return err
	}
}

/*
删除实例

根据传入实例的名称来释放实例资源。释放后实例所使用的物理资源都被回收，包括磁盘及快照，相关数据全部丢失且永久不可恢复。
实例状态必须为 Stopped，才可以进行删除操作。删除后，实例的状态为 Deleted，表示资源已释放，删除完成。
实例被删除时，挂载在实例上的 DeleteWithInstance的属性为 True 的磁盘会相应被删除，这些磁盘的快照任旧保留，
自动快照根据磁盘的 DeleteAutoSnapshot 属性，如果为 false 的，保留自动快照，如果为 true 的，则删除自动快照。
实例被删除后，相关数据全部丢失且永久不可恢复。
如果删除实例时，实例被安全控制在实例的 OperationLocks 中标记了 "LockReason" : "security" 的锁定状态时，
即使独立普通云盘的 DeleteWithInstnace 的属性为 False，系统会忽略这个属性而释放挂载在实例上面的普通云盘。
*/
func (c *Client) DeleteInstance(instanceId string) error {
	params := c.baseParams(c.accessKeyId, nil)
	params.Add("Format", "JSON")
	params.Add("Action", "DeleteInstance")
	params.Add("InstanceId", instanceId)
	var response EcsBaseResponse
	err := util.CallApiServer(API_SERVER, c.signer, params, &response)
	if err == nil {
		return nil
	} else {
		return err
	}
}

//创建实例
func (c *Client) CreateInstance(instance InstanceAttributesType, password string, securityGroupId string) (string, error) {
	params := c.baseParams(c.accessKeyId, nil)
	params.Add("Format", "JSON")
	params.Add("Action", "CreateInstance")
	params.Add("RegionId", instance.RegionId)
	params.Add("ZoneId", instance.ZoneId)
	params.Add("ImageId", instance.ImageId)
	params.Add("InstanceType", instance.InstanceType)
	params.Add("SecurityGroupId", securityGroupId)
	params.Add("InstanceName", instance.InstanceName)
	params.Add("Description", instance.Description)
	params.Add("InternetChargeType", instance.InternetChargeType)
	params.Add("InternetMaxBandwidthIn", strconv.Itoa(instance.InternetMaxBandwidthIn))
	params.Add("InternetMaxBandwidthOut", strconv.Itoa(instance.InternetMaxBandwidthOut))
	params.Add("HostName", instance.HostName)
	params.Add("Password", password)
	var createInstanceResponse CreateInstanceResponse
	err := util.CallApiServer(API_SERVER, c.signer, params, &createInstanceResponse)
	if err == nil {
		return createInstanceResponse.InstanceId, nil
	} else {
		return "", nil
	}
}

type nonceGenerator interface {
	Int63() int64
}

// 构造公共参数
func (c *Client) baseParams(accessKeyId string, additionalParams map[string]string) *util.OrderedParams {
	params := util.NewOrderedParams()
	params.Add(VERSION_PARAM, VERSION)
	params.Add(SIGNATURE_VERSION_PARAM, SIGNATURE_VERSION)
	params.Add(SIGNATURE_METHOD_PARAM, SIGNATURE_METHOD_HMAC_SHA1)
	params.Add(TIMESTAMP_PARAM, time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	params.Add(NONCE_PARAM, strconv.FormatInt(c.nonceGenerator.Int63(), 10))
	params.Add(ACCESS_KEY_ID_PARAM, accessKeyId)
	for key, value := range additionalParams {
		params.Add(key, value)
	}
	return params
}
