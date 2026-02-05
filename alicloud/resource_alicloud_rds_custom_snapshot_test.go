// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Rds CustomSnapshot. >>> Resource test cases, automatically generated.
// Case CustomSnapshot_20250909 11434
func TestAccAliCloudRdsCustomSnapshot_basic11434(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomSnapshotMap11434)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomSnapshotBasicDependence11434)
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
					"description":                   "快照tag-带customdisk版本",
					"disk_id":                       "${alicloud_rds_custom_disk.customdiskItem.id}",
					"zone_id":                       "${var.test_zone_id}",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"retention_days":                "3",
					"instant_access_retention_days": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":                   "快照tag-带customdisk版本",
						"disk_id":                       CHECKSET,
						"zone_id":                       CHECKSET,
						"resource_group_id":             CHECKSET,
						"retention_days":                "3",
						"instant_access_retention_days": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"instant_access_retention_days", "retention_days", "zone_id"},
			},
		},
	})
}

var AlicloudRdsCustomSnapshotMap11434 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomSnapshotBasicDependence11434(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id = alicloud_vpc.vpcId.id
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}

resource "alicloud_rds_custom" "customItem" {
  amount        = "1"
  vswitch_id    = alicloud_vswitch.vSwitchId.id
  auto_renew    = false
  period        = "1"
  auto_pay      = true
  instance_type = "mysql.x2.xlarge.6cm"
  data_disk {
    category          = "cloud_essd"
    size              = "50"
    performance_level = "PL1"
  }
  status                        = "Running"
  security_group_ids            = ["${alicloud_security_group.securityGroupId.id}"]
  io_optimized                  = "optimized"
  description                   = "disk测试依赖custom"
  key_pair_name                 = alicloud_ecs_key_pair.KeyPairName.id
  zone_id                       = var.test_zone_id
  instance_charge_type          = "Prepaid"
  internet_max_bandwidth_out    = "0"
  image_id                      = "aliyun_2_1903_x64_20G_alibase_20240628.vhd"
  security_enhancement_strategy = "Active"
  period_unit                   = "Month"
  password                      = "jingyiTEST@123"
  system_disk {
    size     = "40"
    category = "cloud_essd"
  }
  host_name     = "1770258625"
  create_mode   = "0"
  spot_strategy = "NoSpot"
}

resource "alicloud_rds_custom_disk" "customdiskItem" {
  description          = "包年disk快照用"
  zone_id              = var.test_zone_id
  size                 = "40"
  instance_charge_type = "Prepaid"
  disk_category        = "cloud_ssd"
  disk_name            = "custom_disk_tosnapshot"
  auto_renew           = false
  period               = "1"
  auto_pay             = true
  period_unit          = "Month"
}


`, name)
}

// Case 测试快照three带CustomDisk版本 10940
func TestAccAliCloudRdsCustomSnapshot_basic10940(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomSnapshotMap10940)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomSnapshotBasicDependence10940)
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
					"description":                   "快照tag-带customdisk版本",
					"disk_id":                       "${alicloud_rds_custom_disk.customdiskItem.id}",
					"zone_id":                       "${var.test_zone_id}",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"retention_days":                "3",
					"instant_access_retention_days": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":                   "快照tag-带customdisk版本",
						"disk_id":                       CHECKSET,
						"zone_id":                       CHECKSET,
						"resource_group_id":             CHECKSET,
						"retention_days":                "3",
						"instant_access_retention_days": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"instant_access_retention_days", "retention_days", "zone_id"},
			},
		},
	})
}

var AlicloudRdsCustomSnapshotMap10940 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomSnapshotBasicDependence10940(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id = alicloud_vpc.vpcId.id
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}

resource "alicloud_rds_custom" "customItem" {
  amount        = "1"
  vswitch_id    = alicloud_vswitch.vSwitchId.id
  auto_renew    = false
  period        = "1"
  auto_pay      = true
  instance_type = "mysql.x2.xlarge.6cm"
  data_disk {
    category          = "cloud_essd"
    size              = "50"
    performance_level = "PL1"
  }
  status                        = "Running"
  security_group_ids            = ["${alicloud_security_group.securityGroupId.id}"]
  io_optimized                  = "optimized"
  description                   = "disk测试依赖custom"
  key_pair_name                 = alicloud_ecs_key_pair.KeyPairName.id
  zone_id                       = var.test_zone_id
  instance_charge_type          = "Prepaid"
  internet_max_bandwidth_out    = "0"
  image_id                      = "aliyun_2_1903_x64_20G_alibase_20240628.vhd"
  security_enhancement_strategy = "Active"
  period_unit                   = "Month"
  password                      = "jingyiTEST@123"
  system_disk {
    size     = "40"
    category = "cloud_essd"
  }
  host_name     = "1770258639"
  create_mode   = "0"
  spot_strategy = "NoSpot"
}

resource "alicloud_rds_custom_disk" "customdiskItem" {
  description          = "包年disk快照用"
  zone_id              = var.test_zone_id
  size                 = "40"
  instance_charge_type = "Prepaid"
  disk_category        = "cloud_ssd"
  disk_name            = "custom_disk_tosnapshot"
  auto_renew           = false
  period               = "1"
  auto_pay             = true
  period_unit          = "Month"
}


`, name)
}

// Case 新克隆要去除update-测试快照three带CustomDisk版本-06-06 10915
func TestAccAliCloudRdsCustomSnapshot_basic10915(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomSnapshotMap10915)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomSnapshotBasicDependence10915)
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
					"description": "快照two-0606",
					"disk_id":     "${alicloud_rds_custom_disk.customdiskItem.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "快照two-0606",
						"disk_id":     CHECKSET,
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
				ImportStateVerifyIgnore: []string{"instant_access_retention_days", "retention_days", "zone_id"},
			},
		},
	})
}

var AlicloudRdsCustomSnapshotMap10915 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomSnapshotBasicDependence10915(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id = alicloud_vpc.vpcId.id
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}

resource "alicloud_rds_custom" "customItem" {
  amount        = "1"
  vswitch_id    = alicloud_vswitch.vSwitchId.id
  auto_renew    = false
  period        = "1"
  auto_pay      = true
  instance_type = "mysql.x2.xlarge.6cm"
  data_disk {
    category          = "cloud_essd"
    size              = "50"
    performance_level = "PL1"
  }
  status                        = "Running"
  security_group_ids            = ["${alicloud_security_group.securityGroupId.id}"]
  io_optimized                  = "optimized"
  description                   = "disk测试依赖custom"
  key_pair_name                 = alicloud_ecs_key_pair.KeyPairName.id
  zone_id                       = var.test_zone_id
  instance_charge_type          = "Prepaid"
  internet_max_bandwidth_out    = "0"
  image_id                      = "aliyun_2_1903_x64_20G_alibase_20240628.vhd"
  security_enhancement_strategy = "Active"
  period_unit                   = "Month"
  password                      = "jingyiTEST@123"
  system_disk {
    size     = "40"
    category = "cloud_essd"
  }
  host_name     = "1770258653"
  create_mode   = "0"
  spot_strategy = "NoSpot"
}

resource "alicloud_rds_custom_disk" "customdiskItem" {
  description          = "包年disk快照用"
  zone_id              = var.test_zone_id
  size                 = "40"
  instance_charge_type = "Prepaid"
  disk_category        = "cloud_ssd"
  disk_name            = "custom_disk_tosnapshot"
  auto_renew           = false
  period               = "1"
  auto_pay             = true
  period_unit          = "Month"
}


`, name)
}

// Case 测试快照two带CustomDisk版本-06-06_依赖 10914
func TestAccAliCloudRdsCustomSnapshot_basic10914(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomSnapshotMap10914)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomSnapshotBasicDependence10914)
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
					"description": "快照one-0606",
					"disk_id":     "${alicloud_rds_custom_disk.customdiskItem.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "快照one-0606",
						"disk_id":     CHECKSET,
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
				ImportStateVerifyIgnore: []string{"instant_access_retention_days", "retention_days", "zone_id"},
			},
		},
	})
}

var AlicloudRdsCustomSnapshotMap10914 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomSnapshotBasicDependence10914(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id = alicloud_vpc.vpcId.id
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}

resource "alicloud_rds_custom" "customItem" {
  amount        = "1"
  vswitch_id    = alicloud_vswitch.vSwitchId.id
  auto_renew    = false
  period        = "1"
  auto_pay      = true
  instance_type = "mysql.x2.xlarge.6cm"
  data_disk {
    category          = "cloud_essd"
    size              = "50"
    performance_level = "PL1"
  }
  status                        = "Running"
  security_group_ids            = ["${alicloud_security_group.securityGroupId.id}"]
  io_optimized                  = "optimized"
  description                   = "disk测试依赖custom"
  key_pair_name                 = alicloud_ecs_key_pair.KeyPairName.id
  zone_id                       = var.test_zone_id
  instance_charge_type          = "Prepaid"
  internet_max_bandwidth_out    = "0"
  image_id                      = "aliyun_2_1903_x64_20G_alibase_20240628.vhd"
  security_enhancement_strategy = "Active"
  period_unit                   = "Month"
  password                      = "jingyiTEST@123"
  system_disk {
    size     = "40"
    category = "cloud_essd"
  }
  host_name     = "1770258667"
  create_mode   = "0"
  spot_strategy = "NoSpot"
}

resource "alicloud_rds_custom_disk" "customdiskItem" {
  description          = "包年disk快照用"
  zone_id              = var.test_zone_id
  size                 = "40"
  instance_charge_type = "Prepaid"
  disk_category        = "cloud_ssd"
  disk_name            = "custom_disk_tosnapshot"
  auto_renew           = false
  period               = "1"
  auto_pay             = true
  period_unit          = "Month"
}


`, name)
}

// Test Rds CustomSnapshot. <<< Resource test cases, automatically generated.
