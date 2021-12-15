package main

import (
	"fmt"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

const endpoint = "lighthouse.tencentcloudapi.com"

type Client struct {
	SecretId  string
	SecretKey string
	InstaceId string
	Region    string
	Endpoint  string
}

func NewClient() Client {
	client := Client{
		SecretId:  os.Getenv("TENCENTCLOUD_SECRET_ID"),
		SecretKey: os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		InstaceId: os.Getenv("TENCENTCLOUD_INSTANCE_ID"),
		Region:    os.Getenv("TENCENTCLOUD_REGION"),
		Endpoint:  endpoint,
	}
	if client.SecretId == "" || client.SecretKey == "" || client.InstaceId == "" || client.Region == "" {
		panic("Please set TENCENTCLOUD_SECRET_ID, TENCENTCLOUD_SECRET_KEY, TENCENTCLOUD_INSTANCE_ID, TENCENTCLOUD_REGION")
	}
	return client
}

func (c Client) AddRules(firewallRules []*lighthouse.FirewallRule) {
	credential := common.NewCredential(c.SecretId, c.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	client, _ := lighthouse.NewClient(credential, c.Region, cpf)

	request := lighthouse.NewCreateFirewallRulesRequest()
	request.InstanceId = common.StringPtr(c.InstaceId)
	request.FirewallRules = firewallRules

	response, err := client.CreateFirewallRules(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

func (c Client) GetRules() string {
	credential := common.NewCredential(c.SecretId, c.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	client, _ := lighthouse.NewClient(credential, c.Region, cpf)

	request := lighthouse.NewDescribeFirewallRulesRequest()

	request.InstanceId = common.StringPtr(c.InstaceId)

	response, err := client.DescribeFirewallRules(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return ""
	}
	if err != nil {
		panic(err)
	}
	return response.ToJsonString()
}
