// 阿里云API go语言版本
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
	c.signer.Debug(enabled)
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
