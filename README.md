# aliyun-api-golang
目前仅实现了ECS部分API，加入进来，一起完善吧

## ECS API

### ECS目前封装了以下API

* DescribeInstances :查询实例列表
* DescribeInstanceAttribute :查询实例信息
* CreateInstance :创建实例
* AllocatePublicIpAddress :分配公网 IP 地址
* StartInstance :启动一个指定的实例
* RebootInstance :重启指定的实例
* StopInstance :停止一个指定的实例
* DeleteInstance :删除实例
* CreateSecurityGroup :创建安全组
* DeleteSecurityGroup :删除安全组
* AuthorizeSecurityGroup :授权安全组In方向的访问权限
* RevokeSecurityGroup :撤销安全组授权规则
* AuthorizeSecurityGroupEgress :添加安全组Out方向的访问规则
* RevokeSecurityGroupEgress :撤销安全组Out方向的访问规则
* 更多API正在完善中, 您也可以参考现有API自己完成

### 安装
 `go get github.com/ChangjunZhao/aliyun-api-golang/ecs`

### 使用方法

```
package main

import (
        "fmt"
        "github.com/ChangjunZhao/aliyun-api-golang/ecs"
       )

func main() {
	c := ecs.NewClient(
                "Access Key ID",
                "Access Key Secret", 
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
```

## RDS API

与ECS API差不多，可参考ECS API完成

