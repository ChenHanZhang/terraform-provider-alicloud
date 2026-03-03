// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test EventBridge Agent. >>> Resource test cases, automatically generated.
// Case AgentTest 12569
func TestAccAliCloudEventBridgeAgent_basic12569(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_event_bridge_agent.default"
	ra := resourceAttrInit(resourceId, AlicloudEventBridgeAgentMap12569)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeAgent")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceventbridge%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEventBridgeAgentBasicDependence12569)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "You are a helpful assistant",
					"prompt":      "You are a helpful assistant",
					"name":        name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "You are a helpful assistant",
						"prompt":      "You are a helpful assistant",
						"name":        name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "You are a helpful assistant for nl2sql",
					"prompt":      "You are a helpful assistant for nl2sql",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "You are a helpful assistant for nl2sql",
						"prompt":      "You are a helpful assistant for nl2sql",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudEventBridgeAgentMap12569 = map[string]string{}

func AlicloudEventBridgeAgentBasicDependence12569(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "agent_name" {
  default = "AgentForPopTest"
}


`, name)
}

// Test EventBridge Agent. <<< Resource test cases, automatically generated.
