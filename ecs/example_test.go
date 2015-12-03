// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//示例代码
package ecs_test

import (
	"fmt"
	"github.com/ChangjunZhao/aliyun-api-golang/ecs"
	"os"
)

func ExampleClient_CreateInstance() {
	c := ecs.NewClient(
		os.Getenv("ECS_ACCESS_KEY_ID"),
		os.Getenv("ECS_ACCESS_KEY_SECRET"),
	)
	c.Debug(true)
	//创建实例
	request := &ecs.CreateInstanceRequest{
		RegionId:                "cn-beijing",
		ImageId:                 "m-25mtsy38b",
		InstanceType:            "ecs.t1.small",
		SecurityGroupId:         "securitygroup",
		Password:                "rootpassword",
		InternetChargeType:      "PayByTraffic",
		InternetMaxBandwidthIn:  "10",
		InternetMaxBandwidthOut: "10",
	}
	if response, err := c.CreateInstanceByRequest(request); err == nil {
		fmt.Println(response.InstanceId)
	} else {
		fmt.Println("error:", err)
	}
	//查询实例
	if instance, err := c.DescribeInstanceAttribute("cn-beijing", "instanceId"); err == nil {
		fmt.Println("instance:", instance)

	} else {
		fmt.Println("error:", err)
	}

}
