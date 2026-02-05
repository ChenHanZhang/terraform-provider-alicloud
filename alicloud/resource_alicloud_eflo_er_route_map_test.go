// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Eflo ErRouteMap. >>> Resource test cases, automatically generated.
// Case er_route_map_vpd2vcc 12479
func TestAccAliCloudEfloErRouteMap_basic12479(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_er_route_map.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloErRouteMapMap12479)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloErRouteMap")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloErRouteMapBasicDependence12479)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-wulanchabu"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"transmission_instance_type":  "VPD",
					"action":                      "permit",
					"reception_instance_type":     "VCC",
					"description":                 "录入策略VPD-VCC",
					"reception_instance_id":       "${alicloud_eflo_vcc.VCC.id}",
					"er_id":                       "${alicloud_eflo_er.ER.id}",
					"reception_instance_owner":    "1511928242963727",
					"transmission_instance_owner": "1511928242963727",
					"transmission_instance_id":    "${alicloud_eflo_vpd.VPD.id}",
					"er_route_map_num":            "1001",
					"destination_cidr_block":      "0.0.0.0/0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"transmission_instance_type":  "VPD",
						"action":                      "permit",
						"reception_instance_type":     "VCC",
						"description":                 "录入策略VPD-VCC",
						"reception_instance_id":       CHECKSET,
						"er_id":                       CHECKSET,
						"reception_instance_owner":    CHECKSET,
						"transmission_instance_owner": CHECKSET,
						"transmission_instance_id":    CHECKSET,
						"er_route_map_num":            "1001",
						"destination_cidr_block":      "0.0.0.0/0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "update-test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "update-test",
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

var AlicloudEfloErRouteMapMap12479 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"er_route_map_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudEfloErRouteMapBasicDependence12479(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}

resource "alicloud_vpc" "VPC" {
  cidr_block = "192.168.0.0/16"
}

resource "alicloud_vswitch" "VSW" {
  vpc_id     = alicloud_vpc.VPC.id
  zone_id    = var.zone_id
  cidr_block = "192.168.0.0/24"
}

resource "alicloud_eflo_vcc" "VCC" {
  connection_type = "VPC"
  zone_id         = var.zone_id
  vswitch_id      = alicloud_vswitch.VSW.id
  vpc_id          = alicloud_vpc.VPC.id
  bandwidth       = "1000"
}

resource "alicloud_eflo_vpd" "VPD" {
  cidr     = "10.0.0.0/8"
  vpd_name = "test-route-map"
}

resource "alicloud_eflo_er" "ER" {
  er_name        = "er-test-routemap"
  master_zone_id = var.zone_id
}

resource "alicloud_eflo_er_attachment" "ER_AT_VPD" {
  resource_tenant_id     = "1511928242963727"
  instance_id            = alicloud_eflo_vpd.VPD.id
  auto_receive_all_route = false
  er_id                  = alicloud_eflo_er.ER.id
  instance_type          = "VPD"
  er_attachment_name     = "test-route-map-tf"
}

resource "alicloud_eflo_er_attachment" "ER_AT_VCC" {
  resource_tenant_id     = "1511928242963727"
  instance_id            = alicloud_eflo_vcc.VCC.id
  er_id                  = alicloud_eflo_er.ER.id
  instance_type          = "VCC"
  er_attachment_name     = "test-route-map-tf"
  auto_receive_all_route = false
}


`, name)
}

// Case er_route_map_vcc2vpd 12484
func TestAccAliCloudEfloErRouteMap_basic12484(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_er_route_map.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloErRouteMapMap12484)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloErRouteMap")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloErRouteMapBasicDependence12484)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-wulanchabu"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"transmission_instance_type":  "VCC",
					"action":                      "deny",
					"reception_instance_type":     "VPD",
					"description":                 "录入策略VCC-VPD",
					"reception_instance_id":       "${alicloud_eflo_vpd.VPD.id}",
					"er_id":                       "${alicloud_eflo_er.ER.id}",
					"reception_instance_owner":    "1511928242963727",
					"transmission_instance_owner": "1511928242963727",
					"transmission_instance_id":    "${alicloud_eflo_vcc.VCC.id}",
					"er_route_map_num":            "1002",
					"destination_cidr_block":      "0.0.0.0/0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"transmission_instance_type":  "VCC",
						"action":                      "deny",
						"reception_instance_type":     "VPD",
						"description":                 "录入策略VCC-VPD",
						"reception_instance_id":       CHECKSET,
						"er_id":                       CHECKSET,
						"reception_instance_owner":    CHECKSET,
						"transmission_instance_owner": CHECKSET,
						"transmission_instance_id":    CHECKSET,
						"er_route_map_num":            "1002",
						"destination_cidr_block":      "0.0.0.0/0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "update-test-2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "update-test-2",
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

var AlicloudEfloErRouteMapMap12484 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"er_route_map_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudEfloErRouteMapBasicDependence12484(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}

resource "alicloud_vpc" "VPC" {
  cidr_block = "192.168.0.0/16"
}

resource "alicloud_vswitch" "VSW" {
  vpc_id     = alicloud_vpc.VPC.id
  zone_id    = var.zone_id
  cidr_block = "192.168.0.0/24"
}

resource "alicloud_eflo_vcc" "VCC" {
  connection_type = "VPC"
  zone_id         = var.zone_id
  vswitch_id      = alicloud_vswitch.VSW.id
  vpc_id          = alicloud_vpc.VPC.id
  bandwidth       = "1000"
}

resource "alicloud_eflo_vpd" "VPD" {
  cidr     = "10.0.0.0/8"
  vpd_name = "test-route-map"
}

resource "alicloud_eflo_er" "ER" {
  er_name        = "er-test-routemap"
  master_zone_id = var.zone_id
}

resource "alicloud_eflo_er_attachment" "ER_AT_VPD" {
  resource_tenant_id     = "1511928242963727"
  instance_id            = alicloud_eflo_vpd.VPD.id
  auto_receive_all_route = false
  er_id                  = alicloud_eflo_er.ER.id
  instance_type          = "VPD"
  er_attachment_name     = "test-route-map-tf"
}

resource "alicloud_eflo_er_attachment" "ER_AT_VCC" {
  resource_tenant_id     = "1511928242963727"
  instance_id            = alicloud_eflo_vcc.VCC.id
  er_id                  = alicloud_eflo_er.ER.id
  instance_type          = "VCC"
  er_attachment_name     = "test-route-map-tf"
  auto_receive_all_route = false
}


`, name)
}

// Test Eflo ErRouteMap. <<< Resource test cases, automatically generated.
