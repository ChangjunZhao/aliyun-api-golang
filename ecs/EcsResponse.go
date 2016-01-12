package ecs

type CreateSecurityGroupResponse struct {
	EcsBaseResponse
	SecurityGroupId string `json:"SecurityGroupId"` //安全组ID
}
