// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cen RouteMapRule. >>> Resource test cases, automatically generated.
// Case RouteMapRule测试 12673
func TestAccAliCloudCenRouteMapRule_basic12673(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_route_map_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCenRouteMapRuleMap12673)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenRouteMapRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenRouteMapRuleBasicDependence12673)
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
					"description": "test",
					"match_conditions": []map[string]interface{}{
						{
							"address_prefixes_include": []string{
								"192.168.1.0/24", "192.168.2.0/24", "192.168.3.0/24"},
							"as_paths_include": []string{
								"65500", "65501", "65502"},
							"community_set_include": []string{
								"400:1", "400:2", "400:3"},
							"destination_instance_ids": []string{
								"vpc-1", "vpc-2", "vpc-3"},
							"destination_instance_ids_reverse_match": "false",
							"destination_instance_types": []string{
								"VPC", "VPN", "VBR"},
							"destination_region_ids": []string{
								"cn-shanghai", "cn-beijing", "cn-qingdao"},
							"route_types": []string{
								"BGP"},
							"source_instance_ids": []string{
								"vpc-1", "vpc-2", "vpc-3"},
							"source_instance_ids_reverse_match": "false",
							"source_instance_types": []string{
								"VPC", "VPN", "VBR"},
							"source_region_ids": []string{
								"cn-beijing", "cn-qingdao", "cn-shanghai"},
							"destination_route_table_ids": []string{
								"vtb-1", "vtb-2", "vtb-3"},
							"source_route_table_ids": []string{
								"vtb-1", "vtb-2", "vtb-3"},
							"match_address_type": "IPv4",
						},
					},
					"priority": "1",
					"dry_run":  "false",
					"set_actions": []map[string]interface{}{
						{
							"route_action": "Permit",
							"as_path_prepend": []string{
								"65500", "65501", "65502"},
							"community_add": []string{
								"400:1", "400:2", "400:3"},
						},
					},
					"route_map_id": "${alicloud_cen_route_map.defaultxtaOyM.id}",
					"name":         name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  "test",
						"priority":     "1",
						"dry_run":      "false",
						"route_map_id": CHECKSET,
						"name":         name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test2",
					"match_conditions": []map[string]interface{}{
						{
							"destination_instance_ids_reverse_match": "true",
							"source_instance_ids_reverse_match":      "true",
							"match_address_type":                     "IPv6",
							"address_prefixes_include":               []string{},
							"as_paths_include":                       []string{},
							"community_set_include":                  []string{},
							"destination_instance_ids":               []string{},
							"destination_instance_types":             []string{},
							"destination_region_ids":                 []string{},
							"route_types":                            []string{},
							"source_instance_ids":                    []string{},
							"source_instance_types":                  []string{},
							"source_region_ids":                      []string{},
							"destination_route_table_ids":            []string{},
							"source_route_table_ids":                 []string{},
						},
					},
					"priority": "2",
					"set_actions": []map[string]interface{}{
						{
							"route_action":    "Deny",
							"as_path_prepend": []string{},
							"community_add":   []string{},
						},
					},
					"name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test2",
						"priority":    "2",
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

var AlicloudCenRouteMapRuleMap12673 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCenRouteMapRuleBasicDependence12673(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_route_map" "defaultxtaOyM" {
  priority = "94"
}


`, name)
}

// Test Cen RouteMapRule. <<< Resource test cases, automatically generated.
