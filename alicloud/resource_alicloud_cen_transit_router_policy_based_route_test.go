// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cen TransitRouterPolicyBasedRoute. >>> Resource test cases, automatically generated.
// Case TransitRouterPolicyBasedRoute生命周期 12658
func TestAccAliCloudCenTransitRouterPolicyBasedRoute_basic12658(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_policy_based_route.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPolicyBasedRouteMap12658)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPolicyBasedRoute")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPolicyBasedRouteBasicDependence12658)
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
					"description":           "路由描述",
					"priority":              "50",
					"name":                  name,
					"policy_table_id":       "${alicloud_cen_transit_router_policy_table.defaultEN72S0.id}",
					"dry_run":               "false",
					"source_port_range":     "1/10",
					"source_cidr":           "192.168.1.0/24",
					"target_route_table_id": "${alicloud_cen_transit_router_route_table.default9xsUKi.transit_router_route_table_id}",
					"dest_port_range":       "200/2000",
					"address_family":        "IPv4",
					"protocol":              "ALL",
					"dscp":                  "23/24",
					"destination_cidr":      "192.168.1.0/24",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":           "路由描述",
						"priority":              "50",
						"name":                  name,
						"policy_table_id":       CHECKSET,
						"dry_run":               "false",
						"source_port_range":     "1/10",
						"source_cidr":           "192.168.1.0/24",
						"target_route_table_id": CHECKSET,
						"dest_port_range":       "200/2000",
						"address_family":        "IPv4",
						"protocol":              "ALL",
						"dscp":                  "23/24",
						"destination_cidr":      "192.168.1.0/24",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "策略路由描述-更新",
					"name":        name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "策略路由描述-更新",
						"name":        name + "_update",
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

var AlicloudCenTransitRouterPolicyBasedRouteMap12658 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCenTransitRouterPolicyBasedRouteBasicDependence12658(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "defaultO0XU5Y" {
}

resource "alicloud_cen_transit_router" "defaultF4twp9" {
  cen_id = alicloud_cen_instance.defaultO0XU5Y.id
}

resource "alicloud_cen_transit_router_policy_table" "defaultEN72S0" {
  transit_router_id = alicloud_cen_transit_router.defaultF4twp9.transit_router_id
}

resource "alicloud_cen_transit_router_route_table" "default9xsUKi" {
  transit_router_id = alicloud_cen_transit_router.defaultF4twp9.transit_router_id
}


`, name)
}

// Case cj_test 12692
func TestAccAliCloudCenTransitRouterPolicyBasedRoute_basic12692(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_policy_based_route.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPolicyBasedRouteMap12692)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPolicyBasedRoute")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPolicyBasedRouteBasicDependence12692)
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
					"description":           "路由描述",
					"priority":              "50",
					"name":                  name,
					"policy_table_id":       "ptbl-udne3ruc9u4wsy09t0",
					"dry_run":               "false",
					"source_port_range":     "1/10",
					"source_cidr":           "192.168.1.0/24",
					"target_route_table_id": "${alicloud_cen_transit_router_route_table.default9xsUKi.transit_router_route_table_id}",
					"dest_port_range":       "200/2000",
					"address_family":        "IPv4",
					"protocol":              "ALL",
					"dscp":                  "23/24",
					"destination_cidr":      "192.168.1.0/24",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":           "路由描述",
						"priority":              "50",
						"name":                  name,
						"policy_table_id":       "ptbl-udne3ruc9u4wsy09t0",
						"dry_run":               "false",
						"source_port_range":     "1/10",
						"source_cidr":           "192.168.1.0/24",
						"target_route_table_id": CHECKSET,
						"dest_port_range":       "200/2000",
						"address_family":        "IPv4",
						"protocol":              "ALL",
						"dscp":                  "23/24",
						"destination_cidr":      "192.168.1.0/24",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "策略路由描述-更新",
					"name":        name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "策略路由描述-更新",
						"name":        name + "_update",
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

var AlicloudCenTransitRouterPolicyBasedRouteMap12692 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCenTransitRouterPolicyBasedRouteBasicDependence12692(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_transit_router_policy_table" "defaultEN72S0" {
  transit_router_id = "cc"
}


`, name)
}

// Test Cen TransitRouterPolicyBasedRoute. <<< Resource test cases, automatically generated.
