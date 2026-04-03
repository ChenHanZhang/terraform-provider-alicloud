package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Rds CustomDisk. >>> Resource test cases, automatically generated.
// Case CustomDisk_tmk_11_26 11913
func TestAccAliCloudRdsCustomDisk_basic11913(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap11913)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence11913)
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
					"description":          "ran测试用例09-11",
					"zone_id":              "${var.test_zone_id}",
					"size":                 "40",
					"performance_level":    "PL0",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_pop2rata",
					"auto_renew":           "false",
					"delete_with_instance": "true",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
					"instance_id":          "${alicloud_rds_custom.RCInstance.id}",
					"status":               "In_use",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"snapshot_id":          "${alicloud_rds_custom_snapshot.RCSnapshot.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例09-11",
						"zone_id":              CHECKSET,
						"size":                 "40",
						"performance_level":    "PL0",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_pop2rata",
						"auto_renew":           "false",
						"delete_with_instance": "true",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
						"instance_id":          CHECKSET,
						"status":               "In_use",
						"resource_group_id":    CHECKSET,
						"snapshot_id":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":                 "50",
					"performance_level":    "PL1",
					"delete_with_instance": "false",
					"status":               "Available",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"type":                 "offline",
					"dry_run":              "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":                 "50",
						"performance_level":    "PL1",
						"delete_with_instance": "false",
						"status":               "Available",
						"resource_group_id":    CHECKSET,
						"type":                 "offline",
						"dry_run":              "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "In_use",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "In_use",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Available",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Available",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap11913 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence11913(name string) string {
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
  description         = "custom用例"
  security_group_name = "test_r_sg"
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_type = "normal"
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}

resource "alicloud_rds_custom" "RCInstance" {
  description          = "快照依赖实例"
  instance_charge_type = "Prepaid"
  auto_renew           = false
  system_disk {
    category = "cloud_essd"
    size     = "20"
  }
  image_id           = "aliyun_2_1903_x64_20G_alibase_20240628.vhd"
  instance_type      = "mysql.x2.xlarge.6cm"
  host_name          = "1775196355"
  spot_strategy      = "NoSpot"
  password           = "jingyiTEST@123"
  status             = "Running"
  key_pair_name      = alicloud_ecs_key_pair.KeyPairName.id
  io_optimized       = "optimized"
  zone_id            = var.test_zone_id
  amount             = "1"
  vswitch_id         = alicloud_vswitch.vSwitchId.id
  period             = "1"
  auto_pay           = true
  security_group_ids = ["${alicloud_security_group.securityGroupId.id}"]
  data_disk {
    category          = "cloud_essd"
    performance_level = "PL1"
    size              = "50"
  }
  internet_max_bandwidth_out    = "0"
  create_mode                   = "0"
  security_enhancement_strategy = "Active"
  period_unit                   = "Month"
}

resource "alicloud_rds_custom_disk" "RCDisk" {
  description          = "包年disk快照用"
  zone_id              = var.test_zone_id
  size                 = "40"
  instance_charge_type = "Prepaid"
  auto_renew           = false
  disk_category        = "cloud_essd"
  period               = "1"
  auto_pay             = true
  disk_name            = "custom_disk_tosnapshot"
  period_unit          = "Month"
}

resource "alicloud_rds_custom_snapshot" "RCSnapshot" {
  description                   = "创建实例使用快照"
  zone_id                       = var.test_zone_id
  retention_days                = "3"
  instant_access_retention_days = "1"
  disk_id                       = alicloud_rds_custom_disk.RCDisk.id
}


`, name)
}

// Case CustomDisk_ran_091_rg_one 11447
func TestAccAliCloudRdsCustomDisk_basic11447(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap11447)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence11447)
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
					"description":          "ran测试用例09-11",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_pop2rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例09-11",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_pop2rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap11447 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence11447(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0606-nodetach-three资源组api 10943
func TestAccAliCloudRdsCustomDisk_basic10943(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10943)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10943)
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
					"description":          "ran测试用加入写死实例tag06-13",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom-preTest-tocreateDisk",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
					"zone_id":              "${var.test_zone_id}",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用加入写死实例tag06-13",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom-preTest-tocreateDisk",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
						"zone_id":              CHECKSET,
						"resource_group_id":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":              "50",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"type":              "offline",
					"dry_run":           "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":              "50",
						"resource_group_id": CHECKSET,
						"type":              "offline",
						"dry_run":           "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10943 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10943(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "instance_id" {
  default = "rc-f77m06d2s5wsb93qo0m1"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Case CustomDisk_ran-pro-one 10853
func TestAccAliCloudRdsCustomDisk_basic10853(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10853)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10853)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-chengdu"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description":          "ran测试用例05-29",
					"zone_id":              "cn-chengdu-b",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_pop2rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例05-29",
						"zone_id":              "cn-chengdu-b",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_pop2rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10853 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10853(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-chengdu"
}


`, name)
}

// Case CustomDisk_ran_0526-rmc-three 10834
func TestAccAliCloudRdsCustomDisk_basic10834(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10834)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10834)
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
					"description":          "ran测试用例05-26",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_pop2rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例05-26",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_pop2rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10834 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10834(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_preone 10667
func TestAccAliCloudRdsCustomDisk_basic10667(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10667)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10667)
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
					"description":          "ran测试用例",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_ra",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_ra",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10667 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10667(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_pretwo 10692
func TestAccAliCloudRdsCustomDisk_basic10692(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10692)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10692)
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
					"description":          "ran测试用例",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_ra",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_ra",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10692 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10692(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_test1 10599
func TestAccAliCloudRdsCustomDisk_basic10599(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10599)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10599)
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
					"description":          "zcc测试用例",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "zcc测试用例",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10599 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10599(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0422-one 10720
func TestAccAliCloudRdsCustomDisk_basic10720(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10720)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10720)
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
					"description":          "ran测试用例",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10720 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10720(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_preseven 10714
func TestAccAliCloudRdsCustomDisk_basic10714(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10714)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10714)
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
					"description":          "ran测试用例",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10714 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10714(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_prethree 10704
func TestAccAliCloudRdsCustomDisk_basic10704(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10704)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10704)
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
					"description":          "ran测试用例",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10704 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10704(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0423-one 10728
func TestAccAliCloudRdsCustomDisk_basic10728(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10728)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10728)
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
					"description":          "ran测试用例04-22",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例04-22",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10728 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10728(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0423-two 10729
func TestAccAliCloudRdsCustomDisk_basic10729(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10729)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10729)
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
					"description":          "ran测试用例04-22",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例04-22",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10729 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10729(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0422-four 10727
func TestAccAliCloudRdsCustomDisk_basic10727(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10727)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10727)
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
					"description":          "ran测试用例04-22",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例04-22",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10727 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10727(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0422-three 10725
func TestAccAliCloudRdsCustomDisk_basic10725(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10725)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10725)
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
					"description":          "ran测试用例04-22",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例04-22",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10725 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10725(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0425-one 10741
func TestAccAliCloudRdsCustomDisk_basic10741(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10741)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10741)
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
					"description":          "ran测试用例04-25",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例04-25",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
						"resource_group_id":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10741 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10741(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0425-two 10752
func TestAccAliCloudRdsCustomDisk_basic10752(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10752)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10752)
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
					"description":          "ran测试用例04-25",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例04-25",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
						"resource_group_id":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10752 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10752(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0519-rmc-one 10798
func TestAccAliCloudRdsCustomDisk_basic10798(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10798)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10798)
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
					"description":          "ran测试用例05-19",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例05-19",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10798 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10798(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Case CustomDisk_ran_0519-rmc-two 10802
func TestAccAliCloudRdsCustomDisk_basic10802(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomDiskMap10802)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustomDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomDiskBasicDependence10802)
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
					"description":          "ran测试用例05-19",
					"zone_id":              "cn-beijing-i",
					"size":                 "40",
					"performance_level":    "PL1",
					"instance_charge_type": "Postpaid",
					"disk_category":        "cloud_essd",
					"disk_name":            "custom_disk_001_rata",
					"auto_renew":           "false",
					"period":               "1",
					"auto_pay":             "true",
					"period_unit":          "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran测试用例05-19",
						"zone_id":              "cn-beijing-i",
						"size":                 "40",
						"performance_level":    "PL1",
						"instance_charge_type": "Postpaid",
						"disk_category":        "cloud_essd",
						"disk_name":            "custom_disk_001_rata",
						"auto_renew":           "false",
						"period":               "1",
						"auto_pay":             "true",
						"period_unit":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"size":    "50",
					"type":    "offline",
					"dry_run": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"size":    "50",
						"type":    "offline",
						"dry_run": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "delete_with_instance", "disk_category", "dry_run", "instance_charge_type", "instance_id", "period", "period_unit", "snapshot_id", "type"},
			},
		},
	})
}

var AlicloudRdsCustomDiskMap10802 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudRdsCustomDiskBasicDependence10802(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}


`, name)
}

// Test Rds CustomDisk. <<< Resource test cases, automatically generated.
