// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//由于阿里云不提供API沙箱测试环境，本示例仅说明简单的单元测试方法
package ecs

import (
	"gopkg.in/check.v1"
	"os"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type ECSTestSuite struct {
	client *Client
}

var _ = check.Suite(&ECSTestSuite{NewClient(
	os.Getenv("ECS_ACCESS_KEY_ID"),
	os.Getenv("ECS_ACCESS_KEY_SECRET"),
)})

//测试查询实例列表
func (s *ECSTestSuite) TestDescribeInstances(c *check.C) {
	_, err := s.client.DescribeInstances("cn-beijing")
	c.Assert(err, check.IsNil)
}

func (s *ECSTestSuite) TestDescribeInstanceAttribute(c *check.C) {
	_, err := s.client.DescribeInstanceAttribute("cn-beijing", "i-25c26cnig")
	c.Assert(err, check.IsNil)

}
