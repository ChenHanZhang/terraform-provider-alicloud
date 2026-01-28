package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test CloudFirewall VpcFirewallAclEngineMode. >>> Resource test cases, automatically generated.
// Case ACL引擎管理-VPC边界防火墙 12344
func TestAccAliCloudCloudFirewallVpcFirewallAclEngineMode_basic12344(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_acl_engine_mode.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallAclEngineModeMap12344)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallAclEngineMode")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallAclEngineModeBasicDependence12344)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"strict_mode":     "0",
					"vpc_firewall_id": "${alicloud_cen_instance.cen.id}",
					"member_uid":      "1511928242963727",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strict_mode":     "0",
						"vpc_firewall_id": CHECKSET,
						"member_uid":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strict_mode": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strict_mode": "1",
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

var AlicloudCloudFirewallVpcFirewallAclEngineModeMap12344 = map[string]string{}

func AlicloudCloudFirewallVpcFirewallAclEngineModeBasicDependence12344(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  description       = "yqc-test001"
  cen_instance_name = "yqc-test-CenInstance001"
}

resource "alicloud_cen_transit_router" "TR" {
  cen_id = alicloud_cen_instance.cen.id
}

resource "alicloud_vpc" "vpc1" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "yqc-vpc-test-001"
}

resource "alicloud_vswitch" "vpc1vsw1" {
  vpc_id     = alicloud_vpc.vpc1.id
  zone_id    = "cn-hangzhou-h"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vpc1vsw2" {
  vpc_id     = alicloud_vpc.vpc1.id
  zone_id    = "cn-hangzhou-i"
  cidr_block = "172.16.2.0/24"
}

resource "alicloud_cen_transit_router_vpc_attachment" "tr-vpc1" {
  vpc_id = alicloud_vpc.vpc1.id
  cen_id = alicloud_cen_instance.cen.id
  zone_mappings {
    vswitch_id = alicloud_vswitch.vpc1vsw1.id
    zone_id    = alicloud_vswitch.vpc1vsw1.zone_id
  }
  zone_mappings {
    vswitch_id = alicloud_vswitch.vpc1vsw2.id
    zone_id    = alicloud_vswitch.vpc1vsw2.zone_id
  }
  transit_router_vpc_attachment_name    = "test"
  transit_router_attachment_description = "111"
  auto_publish_route_enabled            = true
  transit_router_id                     = alicloud_cen_transit_router.TR.transit_router_id
}


`, name)
}

// Test CloudFirewall VpcFirewallAclEngineMode. <<< Resource test cases, automatically generated.
