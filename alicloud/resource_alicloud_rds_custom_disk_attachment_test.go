// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Rds CustomDiskAttachment. >>> Resource test cases, automatically generated.
// Case resourceCase_20260407_MCNQ7Y 12735
func TestAccAliCloudRdsCustomDiskAttachment_basic12735(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskAttachmentMap12735)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDiskAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskAttachmentBasicDependence12735)
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
					"instance_id":          "rc-dx4wkoh5zi094o66wpq1",
					"delete_with_instance": "true",
					"disk_id":              "rcd-rx50e2118681h9lbs0fpo",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id":          "rc-dx4wkoh5zi094o66wpq1",
						"delete_with_instance": "true",
						"disk_id":              "rcd-rx50e2118681h9lbs0fpo",
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

var AlicloudRdsCustomDiskAttachmentMap12735 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomDiskAttachmentBasicDependence12735(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260407_03_clone_0 12736
func TestAccAliCloudRdsCustomDiskAttachment_basic12736(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskAttachmentMap12736)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDiskAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskAttachmentBasicDependence12736)
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
					"instance_id":          "${alicloud_rds_custom.custom_clone_0.id}",
					"delete_with_instance": "true",
					"disk_id":              "${alicloud_rds_custom_disk.customdisk_clone_0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id":          CHECKSET,
						"delete_with_instance": "true",
						"disk_id":              CHECKSET,
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

var AlicloudRdsCustomDiskAttachmentMap12736 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomDiskAttachmentBasicDependence12736(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc_clone_0" {
  is_default = false
  cidr_block = "172.16.0.0/16"
  vpc_name   = "ran_vpc1"
}

resource "alicloud_security_group" "sec_clone_0" {
  description         = "sgAttach"
  security_group_name = "ran_sg"
  vpc_id              = alicloud_vpc.vpc_clone_0.id
}

resource "alicloud_vswitch" "vsw_clone_0" {
  is_default   = false
  description  = "rante_vsw1"
  vpc_id       = alicloud_vpc.vpc_clone_0.id
  cidr_block   = "172.16.0.0/23"
  vswitch_name = "rante_vsw1"
}

resource "alicloud_rds_custom" "custom_clone_0" {
  description          = "test资源用例0407A"
  zone_id              = alicloud_vswitch.vsw_clone_0.zone_id
  instance_charge_type = "Prepaid"
  auto_renew           = true
  vswitch_id           = alicloud_vswitch.vsw_clone_0.id
  amount               = "1"
  dry_run              = false
  period               = "1"
  security_group_ids   = ["${alicloud_security_group.sec_clone_0.id}"]
  system_disk {
    size = "40"
  }
  instance_type = "mysql.xa2.xlarge.8cm"
  spot_strategy = "NoSpot"
  period_unit   = "Month"
}

resource "alicloud_rds_custom_disk" "customdisk_clone_0" {
  zone_id       = alicloud_vswitch.vsw_clone_0.zone_id
  size          = "40"
  disk_category = "cloud_ssd"
  auto_pay      = true
  disk_name     = "ran_disk_attach"
}


`, name)
}

// Case resourceCase_20260407_03 12738
func TestAccAliCloudRdsCustomDiskAttachment_basic12738(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskAttachmentMap12738)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDiskAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskAttachmentBasicDependence12738)
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
					"instance_id":          "${alicloud_rds_custom.custom.id}",
					"delete_with_instance": "true",
					"disk_id":              "${alicloud_rds_custom_disk.customdisk.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id":          CHECKSET,
						"delete_with_instance": "true",
						"disk_id":              CHECKSET,
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

var AlicloudRdsCustomDiskAttachmentMap12738 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomDiskAttachmentBasicDependence12738(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  is_default = false
  cidr_block = "172.16.0.0/16"
  vpc_name   = "ran_vpc1"
}

resource "alicloud_security_group" "sec" {
  description         = "sgAttach"
  security_group_name = "ran_sg"
  vpc_id              = alicloud_vpc.vpc.id
}

resource "alicloud_vswitch" "vsw" {
  is_default   = false
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "cn-beijing-i"
  cidr_block   = "172.16.0.0/23"
  vswitch_name = "rante_vsw1"
}

resource "alicloud_rds_custom" "custom" {
  description          = "test资源用例0408A"
  zone_id              = alicloud_vswitch.vsw.zone_id
  instance_charge_type = "Prepaid"
  auto_renew           = true
  vswitch_id           = alicloud_vswitch.vsw.id
  amount               = "1"
  dry_run              = false
  period               = "1"
  security_group_ids   = ["${alicloud_security_group.sec.id}"]
  system_disk {
    size = "40"
  }
  instance_type = "mysql.xa2.xlarge.8cm"
  spot_strategy = "NoSpot"
  period_unit   = "Month"
}

resource "alicloud_rds_custom_disk" "customdisk" {
  zone_id       = alicloud_vswitch.vsw.zone_id
  size          = "40"
  disk_category = "cloud_essd"
  auto_pay      = true
  disk_name     = "ran_disk_attach"
}


`, name)
}

// Test Rds CustomDiskAttachment. <<< Resource test cases, automatically generated.
