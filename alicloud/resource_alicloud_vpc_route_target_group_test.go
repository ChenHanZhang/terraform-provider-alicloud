// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Vpc RouteTargetGroup. >>> Resource test cases, automatically generated.
// Case 路由目标组生命周期测试-预发 12108
func TestAccAliCloudVpcRouteTargetGroup_basic12108(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_vpc_route_target_group.default"
	ra := resourceAttrInit(resourceId, AlicloudVpcRouteTargetGroupMap12108)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &VpcServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeVpcRouteTargetGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccvpc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudVpcRouteTargetGroupBasicDependence12108)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"eu-central-1"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"route_target_group_description": "预发-路由目标组-资源测试",
					"vpc_id":                         "${alicloud_vpc.defaultVpc.id}",
					"route_target_member_list": []map[string]interface{}{
						{
							"member_id":           "${alicloud_privatelink_vpc_endpoint.getVpcEndpointA.id}",
							"member_type":         "GatewayLoadBalancerEndpoint",
							"weight":              "100",
							"enable_status":       "Enable",
							"health_check_status": "Normal",
						},
						{
							"member_id":           "${alicloud_privatelink_vpc_endpoint.getVpcEndpointB.id}",
							"member_type":         "GatewayLoadBalancerEndpoint",
							"weight":              "0",
							"enable_status":       "Disable",
							"health_check_status": "Normal",
						},
					},
					"config_mode":             "Active-Standby",
					"route_target_group_name": name,
					"resource_group_id":       "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"route_target_group_description": "预发-路由目标组-资源测试",
						"vpc_id":                         CHECKSET,
						"route_target_member_list.#":     "2",
						"config_mode":                    "Active-Standby",
						"route_target_group_name":        name,
						"resource_group_id":              CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"route_target_group_description": "预发-路由目标组-资源测试-更新后",
					"route_target_member_list": []map[string]interface{}{
						{
							"member_id":           "${alicloud_privatelink_vpc_endpoint.getVpcEndpointB2.id}",
							"enable_status":       "Disable",
							"member_type":         "GatewayLoadBalancerEndpoint",
							"health_check_status": "Normal",
							"weight":              "0",
						},
						{
							"member_id":           "${alicloud_privatelink_vpc_endpoint.getVpcEndpointA.id}",
							"enable_status":       "Enable",
							"member_type":         "GatewayLoadBalancerEndpoint",
							"health_check_status": "Normal",
							"weight":              "100",
						},
					},
					"route_target_group_name": name + "_update",
					"resource_group_id":       "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"route_target_group_description": "预发-路由目标组-资源测试-更新后",
						"route_target_member_list.#":     "2",
						"route_target_group_name":        name + "_update",
						"resource_group_id":              CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudVpcRouteTargetGroupMap12108 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudVpcRouteTargetGroupBasicDependence12108(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "vpc_name" {
  default = "预发-GWLB-业务VPC-勿删"
}

variable "gwlbe_name_a1" {
  default = "预发-GWLBE-可用区A-勿删"
}

variable "gwlbe_name_b2" {
  default = "预发-GWLBE-可用区B-2-勿删"
}

variable "eps_name_a" {
  default = "预发-GWLB-GES-可用区A-勿删"
}

variable "eps_name_b" {
  default = "预发-GWLB-GES-可用区B-勿删"
}

variable "gwlbe_name_b1" {
  default = "预发-GWLBE-可用区B-勿删"
}

variable "region" {
  default = "eu-central-1"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultVpc" {
  vpc_name    = var.vpc_name
  cidr_block  = "192.168.0.0/16"
  enable_ipv6 = false
}

resource "alicloud_privatelink_vpc_endpoint_service" "getVpcEndpointServiceA" {
  auto_accept_connection = true
  service_description    = var.eps_name_a
}

resource "alicloud_privatelink_vpc_endpoint" "getVpcEndpointA" {
  service_id        = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceA.id
  vpc_endpoint_name = var.gwlbe_name_a1
  vpc_id            = alicloud_vpc.defaultVpc.id
  service_name      = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceA.vpc_endpoint_service_name
  endpoint_type     = "GatewayLoadBalancer"
}

resource "alicloud_privatelink_vpc_endpoint_service" "getVpcEndpointServiceB" {
  auto_accept_connection = true
  service_description    = var.eps_name_b
}

resource "alicloud_privatelink_vpc_endpoint" "getVpcEndpointB" {
  service_id        = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceB.id
  vpc_endpoint_name = var.gwlbe_name_b1
  vpc_id            = alicloud_vpc.defaultVpc.id
  service_name      = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceB.vpc_endpoint_service_name
  endpoint_type     = "GatewayLoadBalancer"
}

resource "alicloud_privatelink_vpc_endpoint" "getVpcEndpointB2" {
  service_id        = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceB.id
  vpc_endpoint_name = var.gwlbe_name_b2
  vpc_id            = alicloud_vpc.defaultVpc.id
  service_name      = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceB.vpc_endpoint_service_name
  endpoint_type     = "GatewayLoadBalancer"
}


`, name)
}

// Test Vpc RouteTargetGroup. <<< Resource test cases, automatically generated.
