// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test DataWorks Route. >>> Resource test cases, automatically generated.
// Case 资源组网络资源路由管理_创建依赖资源 10329
func TestAccAliCloudDataWorksRoute_basic10329(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_route.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksRouteMap10329)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksRoute")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksRouteBasicDependence10329)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"destination_cidr":     "198.162.0.0/24",
					"network_id":           "${alicloud_data_works_network.defaultwEWYyK.id}",
					"dw_resource_group_id": "${alicloud_data_works_dw_resource_group.defaultVJvKvl.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"destination_cidr":     "198.162.0.0/24",
						"network_id":           CHECKSET,
						"dw_resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"destination_cidr": "198.162.0.1/32",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"destination_cidr": "198.162.0.1/32",
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

var AlicloudDataWorksRouteMap10329 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudDataWorksRouteBasicDependence10329(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "default5Bia4h" {
  description = "network_default_resgv2061"
  vpc_name    = "network_default_resgv2061"
  cidr_block  = "10.0.0.0/8"
}

resource "alicloud_vswitch" "defaultss7s7F" {
  description  = "network_default_resg102"
  vpc_id       = alicloud_vpc.default5Bia4h.id
  zone_id      = "cn-beijing-g"
  vswitch_name = "network_default_resg102"
  cidr_block   = "10.0.0.0/24"
}

resource "alicloud_data_works_dw_resource_group" "defaultVJvKvl" {
  default_vswitch_id  = alicloud_vswitch.defaultss7s7F.id
  default_vpc_id      = alicloud_vpc.default5Bia4h.id
  remark              = "route_test"
  payment_type        = "PostPaid"
  resource_group_name = "route_test0008"
}

resource "alicloud_vpc" "defaulte4zhaL" {
  description = "network_default_resgv2010"
  vpc_name    = "network_default_resgv2010"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default675v38" {
  description  = "network_default_resg010"
  vpc_id       = alicloud_vpc.defaulte4zhaL.id
  zone_id      = "cn-beijing-g"
  vswitch_name = "network_default_resg011"
  cidr_block   = "172.16.0.0/24"
}

resource "alicloud_data_works_network" "defaultwEWYyK" {
  vpc_id               = alicloud_vpc.defaulte4zhaL.id
  vswitch_id           = alicloud_vswitch.default675v38.id
  dw_resource_group_id = alicloud_data_works_dw_resource_group.defaultVJvKvl.id
}


`, name)
}

// Test DataWorks Route. <<< Resource test cases, automatically generated.
