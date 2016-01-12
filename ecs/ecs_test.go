// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//由于阿里云不提供API沙箱测试环境，本示例仅说明简单的单元测试方法
package ecs

import (
	"gopkg.in/check.v1"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type ECSTestSuite struct {
	client   *Client
	RegionId string
}

var _ = check.Suite(&ECSTestSuite{client: NewClient(
	os.Getenv("ECS_ACCESS_KEY_ID"),
	os.Getenv("ECS_ACCESS_KEY_SECRET"),
)})

func (s *ECSTestSuite) SetUpSuite(c *check.C) {
	s.RegionId = "cn-beijing"
}

func (s *ECSTestSuite) TearDownSuite(c *check.C) {
	log.Println("TearDownSuite")
}

//测试查询实例列表
func (s *ECSTestSuite) TestDescribeInstances(c *check.C) {
	_, err := s.client.DescribeInstances(s.RegionId)
	c.Assert(err, check.IsNil)
}

func (s *ECSTestSuite) TestDescribeInstanceAttribute(c *check.C) {
	_, err := s.client.DescribeInstanceAttribute(s.RegionId, "i-25c26cnig")
	c.Assert(err, check.IsNil)

}

// 测试安全组相关操作
func (s *ECSTestSuite) TestSecurityGroup(c *check.C) {

	// 测试创建安全组
	response, err := s.client.CreateSecurityGroup(&CreateSecurityGroupRequest{RegionId: s.RegionId})
	securityGroupId := response.SecurityGroupId
	c.Assert(err, check.IsNil)

	// 测试授权安全组In方向的访问权限
	authorizeSecurityGroupRequest := &AuthorizeSecurityGroupRequest{
		SecurityGroupId: securityGroupId,
		RegionId:        s.RegionId,
		IpProtocol:      "tcp",
		PortRange:       "80/80",
		SourceCidrIp:    "0.0.0.0/0",
	}
	_, err = s.client.AuthorizeSecurityGroup(authorizeSecurityGroupRequest)
	c.Assert(err, check.IsNil)

	// 测试撤销安全组授权规则
	revokeSecurityGroupRequest := &RevokeSecurityGroupRequest{
		SecurityGroupId: securityGroupId,
		RegionId:        s.RegionId,
		IpProtocol:      "tcp",
		PortRange:       "80/80",
		SourceCidrIp:    "0.0.0.0/0",
	}
	_, err = s.client.RevokeSecurityGroup(revokeSecurityGroupRequest)
	c.Assert(err, check.IsNil)

	// 测试删除安全组
	_, err = s.client.DeleteSecurityGroup(&DeleteSecurityGroupRequest{RegionId: s.RegionId, SecurityGroupId: securityGroupId})
	c.Assert(err, check.IsNil)
}
