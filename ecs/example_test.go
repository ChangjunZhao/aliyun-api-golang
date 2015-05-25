package ecs_test

import (
	"fmt"
	"github.com/ChangjunZhao/aliyun-api-golang/ecs"
)

func ExampleClient_CreateInstance() {
	c := ecs.NewClient(
		"Access Key ID",
		"Access Key Secret",
	)
	c.Debug(true)
	//创建实例
	var instance ecs.InstanceAttributesType
	instance.RegionId = "cn-beijing"
	instance.ImageId = "m-25mtsy38b"
	instance.InstanceType = "ecs.t1.small"
	instance.InternetChargeType = "PayByTraffic"
	instance.InternetMaxBandwidthIn = 1
	instance.InternetMaxBandwidthOut = 1
	instanceId, _ := c.CreateInstance(instance, "rootpassword", "securitygroup")
	//查询实例
	instancenew, err := c.DescribeInstanceAttribute("instanceId")
	if err == nil {
		fmt.Println("instance:", instances)

	} else {
		fmt.Println("error:", err)

	}

}
