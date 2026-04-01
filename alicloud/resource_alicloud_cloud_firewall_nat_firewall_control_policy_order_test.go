// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test CloudFirewall NatFirewallControlPolicyOrder. >>> Resource test cases, automatically generated.
// Case test 12703
func TestAccAliCloudCloudFirewallNatFirewallControlPolicyOrder_basic12703(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_nat_firewall_control_policy_order.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallNatFirewallControlPolicyOrderMap12703)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallNatFirewallControlPolicyOrder")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallNatFirewallControlPolicyOrderBasicDependence12703)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"acl_uuid":       "a3b5e8f3-6d2c-4e26-b078-87cee0790106",
					"nat_gateway_id": "ngw-2ze8hqi59w9wrm30zwgnq",
					"direction":      "out",
					"order":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"acl_uuid":       "a3b5e8f3-6d2c-4e26-b078-87cee0790106",
						"nat_gateway_id": "ngw-2ze8hqi59w9wrm30zwgnq",
						"direction":      "out",
						"order":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"order": "2",
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

var AlicloudCloudFirewallNatFirewallControlPolicyOrderMap12703 = map[string]string{}

func AlicloudCloudFirewallNatFirewallControlPolicyOrderBasicDependence12703(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test CloudFirewall NatFirewallControlPolicyOrder. <<< Resource test cases, automatically generated.
