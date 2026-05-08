// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cen TransitRouterPolicyTable. >>> Resource test cases, automatically generated.
// Case PolicyTable生命周期 12652
func TestAccAliCloudCenTransitRouterPolicyTable_basic12652(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_policy_table.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPolicyTableMap12652)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPolicyTable")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPolicyTableBasicDependence12652)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-wulanchabu-test-6"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"name":              name,
					"transit_router_id": "${alicloud_cen_transit_router.defaultu3Rgvn.transit_router_id}",
					"description":       "自动化资源用例描述",
					"dry_run":           "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"name":              name,
						"transit_router_id": CHECKSET,
						"description":       "自动化资源用例描述",
						"dry_run":           "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"name":        name + "_update",
					"description": "自动化资源用例-描述更新",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"name":        name + "_update",
						"description": "自动化资源用例-描述更新",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "Test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF-update",
						"tags.For":     "Test-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "0",
						"tags.Created": REMOVEKEY,
						"tags.For":     REMOVEKEY,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dry_run"},
			},
		},
	})
}

var AlicloudCenTransitRouterPolicyTableMap12652 = map[string]string{
	"status":    CHECKSET,
	"region_id": CHECKSET,
}

func AlicloudCenTransitRouterPolicyTableBasicDependence12652(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-wulanchabu-test-6"
}

resource "alicloud_cen_instance" "defaultiEuUyd" {
  cen_instance_name = "镇元自动测试用例"
}

resource "alicloud_cen_transit_router" "defaultu3Rgvn" {
  cen_id = alicloud_cen_instance.defaultiEuUyd.id
}


`, name)
}

// Test Cen TransitRouterPolicyTable. <<< Resource test cases, automatically generated.
