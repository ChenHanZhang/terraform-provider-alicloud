// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test CloudFirewall VpcFirewallControlPolicyOrder. >>> Resource test cases, automatically generated.
// Case test1 12717
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicyOrder_basic12717(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy_order.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyOrderMap12717)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicyOrder")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyOrderBasicDependence12717)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"order":           "1",
					"vpc_firewall_id": "cen-38mhpjiqwbkfullqdj",
					"lang":            "zh",
					"acl_uuid":        "b71137c7-23f0-411d-b6a0-8a2f1977fe6f",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":           CHECKSET,
						"vpc_firewall_id": "cen-38mhpjiqwbkfullqdj",
						"lang":            "zh",
						"acl_uuid":        "b71137c7-23f0-411d-b6a0-8a2f1977fe6f",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"order": "4",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order": CHECKSET,
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

var AlicloudCloudFirewallVpcFirewallControlPolicyOrderMap12717 = map[string]string{}

func AlicloudCloudFirewallVpcFirewallControlPolicyOrderBasicDependence12717(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test CloudFirewall VpcFirewallControlPolicyOrder. <<< Resource test cases, automatically generated.
