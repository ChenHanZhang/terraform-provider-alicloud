// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cen TransitRouterPolicyTableAssociation. >>> Resource test cases, automatically generated.
// Case TransitRouterPolicyTableAssociation生命周期_线上_菲律宾 12718
func TestAccAliCloudCenTransitRouterPolicyTableAssociation_basic12718(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_policy_table_association.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPolicyTableAssociationMap12718)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPolicyTableAssociation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12718)
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
					"policy_table_id": "${alicloud_cen_transit_router_policy_table.default0LZN3E.id}",
					"dry_run":         "false",
					"attachment_id":   "${alicloud_cen_transit_router_vpc_attachment.defaultb3bDZ6.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"policy_table_id": CHECKSET,
						"dry_run":         "false",
						"attachment_id":   CHECKSET,
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

var AlicloudCenTransitRouterPolicyTableAssociationMap12718 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12718(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "instance_name" {
  default = "镇元模型测试PolicyTableAssociation"
}

variable "instance_zone_id_1" {
  default = "ap-southeast-6a"
}

variable "instance_region" {
  default = "ap-southeast-6"
}

resource "alicloud_vpc" "defaultXWE6ox" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = var.instance_name
}

resource "alicloud_vswitch" "defaultAJ7Hxd" {
  vpc_id       = alicloud_vpc.defaultXWE6ox.id
  cidr_block   = "192.168.1.0/24"
  zone_id      = var.instance_zone_id_1
  vswitch_name = var.instance_name
}

resource "alicloud_cen_instance" "defaultZN8rN1" {
  cen_instance_name = var.instance_name
}

resource "alicloud_cen_transit_router" "defaultlMj6lY" {
  cen_id              = alicloud_cen_instance.defaultZN8rN1.id
  transit_router_name = var.instance_name
}

resource "alicloud_cen_transit_router_vpc_attachment" "defaultb3bDZ6" {
  vpc_id = alicloud_vpc.defaultXWE6ox.id
  cen_id = alicloud_cen_instance.defaultZN8rN1.id
  zone_mappings {
    vswitch_id = alicloud_vswitch.defaultAJ7Hxd.id
    zone_id    = alicloud_vswitch.defaultAJ7Hxd.zone_id
  }
  transit_router_id = alicloud_cen_transit_router.defaultlMj6lY.transit_router_id
}

resource "alicloud_cen_transit_router_policy_table" "default0LZN3E" {
  transit_router_id = alicloud_cen_transit_router.defaultlMj6lY.transit_router_id
  name              = var.instance_name
}


`, name)
}

// Case cs 12740
func TestAccAliCloudCenTransitRouterPolicyTableAssociation_basic12740(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_policy_table_association.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPolicyTableAssociationMap12740)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPolicyTableAssociation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12740)
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
					"policy_table_id": "${alicloud_cen_transit_router_policy_table.default0LZN3E.id}",
					"dry_run":         "false",
					"attachment_id":   "${alicloud_cen_transit_router_vpc_attachment.defaultb3bDZ6.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"policy_table_id": CHECKSET,
						"dry_run":         "false",
						"attachment_id":   CHECKSET,
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

var AlicloudCenTransitRouterPolicyTableAssociationMap12740 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12740(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "instance_name" {
  default = "镇元模型测试PolicyTableAssociation"
}

variable "instance_zone_id_1" {
  default = "ap-southeast-6a"
}

variable "instance_region" {
  default = "ap-southeast-6"
}

resource "alicloud_vpc" "defaultXWE6ox" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = var.instance_name
}

resource "alicloud_vswitch" "defaultAJ7Hxd" {
  vpc_id       = alicloud_vpc.defaultXWE6ox.id
  cidr_block   = "192.168.1.0/24"
  zone_id      = var.instance_zone_id_1
  vswitch_name = var.instance_name
}

resource "alicloud_cen_transit_router_policy_table" "default0LZN3E" {
  transit_router_id = "cs"
  name              = var.instance_name
}


`, name)
}

// Case TransitRouterPolicyTableAssociation生命周期_写死实例 12694
func TestAccAliCloudCenTransitRouterPolicyTableAssociation_basic12694(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_policy_table_association.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPolicyTableAssociationMap12694)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPolicyTableAssociation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12694)
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
					"policy_table_id": "ptbl-6ks0j63q2f5nrfees1",
					"attachment_id":   "tr-attach-pljo1ollxcbin6mbgm",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"policy_table_id": "ptbl-6ks0j63q2f5nrfees1",
						"attachment_id":   "tr-attach-pljo1ollxcbin6mbgm",
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

var AlicloudCenTransitRouterPolicyTableAssociationMap12694 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12694(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case TransitRouterPolicyTableAssociation生命周期 12669
func TestAccAliCloudCenTransitRouterPolicyTableAssociation_basic12669(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_policy_table_association.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPolicyTableAssociationMap12669)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPolicyTableAssociation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12669)
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
					"policy_table_id": "${alicloud_cen_transit_router_policy_table.default0LZN3E.id}",
					"dry_run":         "false",
					"attachment_id":   "${alicloud_cen_transit_router_vpc_attachment.defaultb3bDZ6.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"policy_table_id": CHECKSET,
						"dry_run":         "false",
						"attachment_id":   CHECKSET,
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

var AlicloudCenTransitRouterPolicyTableAssociationMap12669 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12669(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultXWE6ox" {
  cidr_block = "192.168.0.0/16"
}

resource "alicloud_vswitch" "defaultAJ7Hxd" {
  vpc_id       = alicloud_vpc.defaultXWE6ox.id
  cidr_block   = "192.168.1.0/24"
  zone_id      = "cn-wulanchabu-test-6a"
  vswitch_name = "镇元资源用例自动化测试20260330"
}

resource "alicloud_cen_instance" "defaultZN8rN1" {
}

resource "alicloud_cen_transit_router" "defaultlMj6lY" {
  cen_id = alicloud_cen_instance.defaultZN8rN1.id
}

resource "alicloud_cen_transit_router_vpc_attachment" "defaultb3bDZ6" {
  zone_mappings {
    vswitch_id = alicloud_vswitch.defaultAJ7Hxd.id
    zone_id    = alicloud_vswitch.defaultAJ7Hxd.zone_id
  }
  vpc_id            = alicloud_vpc.defaultXWE6ox.id
  cen_id            = alicloud_cen_instance.defaultZN8rN1.id
  transit_router_id = alicloud_cen_transit_router.defaultlMj6lY.transit_router_id
}

resource "alicloud_cen_transit_router_policy_table" "default0LZN3E" {
  transit_router_id = alicloud_cen_transit_router.defaultlMj6lY.transit_router_id
}


`, name)
}

// Case cjTmp_test 12693
func TestAccAliCloudCenTransitRouterPolicyTableAssociation_basic12693(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_policy_table_association.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterPolicyTableAssociationMap12693)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterPolicyTableAssociation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12693)
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
					"policy_table_id": "${alicloud_cen_transit_router_policy_table.default0LZN3E.id}",
					"dry_run":         "false",
					"attachment_id":   "${alicloud_cen_transit_router_vpc_attachment.defaultb3bDZ6.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"policy_table_id": CHECKSET,
						"dry_run":         "false",
						"attachment_id":   CHECKSET,
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

var AlicloudCenTransitRouterPolicyTableAssociationMap12693 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCenTransitRouterPolicyTableAssociationBasicDependence12693(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_transit_router_policy_table" "default0LZN3E" {
  transit_router_id = "cs"
}


`, name)
}

// Test Cen TransitRouterPolicyTableAssociation. <<< Resource test cases, automatically generated.
