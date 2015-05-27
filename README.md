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
* 更多API正在完善中, 您也可以参考现有API自己完成

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
```

## RDS API

与ECS API差不多，可参考ECS API完成

