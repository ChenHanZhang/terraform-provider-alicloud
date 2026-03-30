// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ExpressConnectRouter RouteMapAssociation. >>> Resource test cases, automatically generated.
// Case RouteMap预发 12668
func TestAccAliCloudExpressConnectRouterRouteMapAssociation_basic12668(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_express_connect_router_route_map_association.default"
	ra := resourceAttrInit(resourceId, AlicloudExpressConnectRouterRouteMapAssociationMap12668)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ExpressConnectRouterServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeExpressConnectRouterRouteMapAssociation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccexpressconnectrouter%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudExpressConnectRouterRouteMapAssociationBasicDependence12668)
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
					"ecr_id": "${var.ecr_id}",
					"region_id_list": []string{
						"${var.region_id_list}"},
					"route_map_id": "${var.route_map_id}",
					"dry_run":      "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ecr_id":           CHECKSET,
						"region_id_list.#": "1",
						"route_map_id":     CHECKSET,
						"dry_run":          "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"region_id_list": []string{
						"cn-wulanchabu-test-5", "cn-wulanchabu-test-6"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"region_id_list.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"region_id_list": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"region_id_list.#": "0",
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

var AlicloudExpressConnectRouterRouteMapAssociationMap12668 = map[string]string{
	"status":         CHECKSET,
	"association_id": CHECKSET,
}

func AlicloudExpressConnectRouterRouteMapAssociationBasicDependence12668(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "ecr_id" {
  default = "ecr-jl48tkimaa8g51o7nd"
}

variable "region_id_list" {
  default = "cn-wulanchabu-test-5,cn-wulanchabu-test-6"
}

variable "route_map_id" {
  default = "rm-yfs0jpvp487aj2dzrg"
}


`, name)
}

// Test ExpressConnectRouter RouteMapAssociation. <<< Resource test cases, automatically generated.
