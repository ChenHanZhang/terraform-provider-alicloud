// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ens HaVip. >>> Resource test cases, automatically generated.
// Case 高可用VIP基础信息_20250110 9956
func TestAccAliCloudEnsHaVip_basic9956(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_ha_vip.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsHaVipMap9956)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsHaVip")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsHaVipBasicDependence9956)
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
					"description": "desc1",
					"vswitch_id":  "${alicloud_ens_vswitch.defaultcW3Eib.id}",
					"amount":      "1",
					"ip_address":  "10.0.9.5",
					"ha_vip_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "desc1",
						"vswitch_id":  CHECKSET,
						"amount":      "1",
						"ip_address":  "10.0.9.5",
						"ha_vip_name": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ha_vip_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ha_vip_name": name + "_update",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"amount"},
			},
		},
	})
}

var AlicloudEnsHaVipMap9956 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudEnsHaVipBasicDependence9956(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "ens_region_id" {
  default = "cn-hangzhou-58"
}

resource "alicloud_ens_network" "default4wYgcV" {
  network_name  = "镇元-测试HaVip"
  cidr_block    = "10.0.0.0/8"
  ens_region_id = var.ens_region_id
}

resource "alicloud_ens_vswitch" "defaultcW3Eib" {
  cidr_block    = "10.0.9.0/24"
  vswitch_name  = "镇元-测试HaVip"
  ens_region_id = var.ens_region_id
  network_id    = alicloud_ens_network.default4wYgcV.id
}


`, name)
}

// Test Ens HaVip. <<< Resource test cases, automatically generated.
