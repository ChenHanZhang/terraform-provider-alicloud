package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/PaesslerAG/jsonpath"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func init() {
	resource.AddTestSweepers("alicloud_ecs_snapshot", &resource.Sweeper{
		Name: "alicloud_ecs_snapshot",
		F:    testSweepEcsSnapshots,
	})
}

func testSweepEcsSnapshots(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return WrapError(err)
	}
	client := rawClient.(*connectivity.AliyunClient)
	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}
	action := "DescribeSnapshots"

	request := map[string]interface{}{
		"PageSize":   PageSizeLarge,
		"PageNumber": 1,
		"RegionId":   client.RegionId,
	}

	var response map[string]interface{}

	for {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_ecs_snapshot", action, AlibabaCloudSdkGoERROR)
		}
		resp, err := jsonpath.Get("$.Snapshots.Snapshot", response)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, action, "$.Snapshots.Snapshot", response)
		}

		result, _ := resp.([]interface{})

		for _, v := range result {
			item := v.(map[string]interface{})

			name := item["SnapshotName"]
			id := item["SnapshotId"]
			skip := true
			if !sweepAll() {
				for _, prefix := range prefixes {
					if strings.HasPrefix(strings.ToLower(name.(string)), strings.ToLower(prefix)) {
						skip = false
						break
					}
				}
				if skip {
					log.Printf("[INFO] Skipping snapshot: %s (%s)", name, id)
					continue
				}
			}
			log.Printf("[INFO] Deleting snapshot: %s (%s)", name, id)
			action = "DeleteSnapshot"
			request := map[string]interface{}{
				"SnapshotId": item["SnapshotId"],
			}

			_, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)

			if err != nil {
				log.Printf("[ERROR] Failed to delete snapshot(%s (%s)): %s", name, id, err)
			}

			log.Printf("[INFO] Delete snapshot success: %s ", item["SnapshotId"].(string))
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}

	return nil
}

// Test Ecs Snapshot. >>> Resource test cases, automatically generated.
// Case resourceCase_20260204_X4KCgM 12562
func TestAccAliCloudEcsSnapshot_basic12562(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsSnapshotMap12562)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsSnapshotBasicDependence12562)
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
					"description": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"lock_mode"},
			},
		},
	})
}

var AlicloudEcsSnapshotMap12562 = map[string]string{
	"status":      " ",
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEcsSnapshotBasicDependence12562(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 基本创建+修改+标签+锁定_换账号测试 12162
func TestAccAliCloudEcsSnapshot_basic12162(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsSnapshotMap12162)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsSnapshotBasicDependence12162)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description":      "model_test",
					"disk_id":          "${alicloud_ecs_disk.查询磁盘.id}",
					"snapshot_name":    name,
					"retention_days":   "3",
					"source_region_id": "${var.region_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":      "model_test",
						"disk_id":          CHECKSET,
						"snapshot_name":    name,
						"retention_days":   "3",
						"source_region_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":    "修改描述",
					"snapshot_name":  name + "_update",
					"retention_days": "4",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":    "修改描述",
						"snapshot_name":  name + "_update",
						"retention_days": "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cool_off_period": "1",
					"lock_duration":   "1",
					"lock_mode":       "compliance",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cool_off_period": "1",
						"lock_duration":   "1",
						"lock_mode":       "compliance",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lock_status": "unlocked",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lock_status": "unlocked",
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
				ImportStateVerifyIgnore: []string{"lock_mode"},
			},
		},
	})
}

var AlicloudEcsSnapshotMap12162 = map[string]string{
	"status":      " ",
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEcsSnapshotBasicDependence12162(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "zone_id" {
  default = "cn-hangzhou-h"
}

variable "destination_region_id" {
  default = "cn-qingdao"
}

variable "region_id" {
  default = "cn-hangzhou"
}

resource "alicloud_vpc" "创建VPC" {
  is_default = false
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "创建交换机" {
  vpc_id     = alicloud_vpc.创建VPC.id
  zone_id    = var.zone_id
  cidr_block = "172.16.0.0/24"
}

resource "alicloud_security_group" "创建安全组" {
  vpc_id              = alicloud_vpc.创建VPC.id
  security_group_type = "normal"
  description         = "sg"
  security_group_name = "sg_name"
}

resource "alicloud_ecs_instance" "创建实例" {
  system_disk {
    size      = "20"
    category  = "cloud_essd"
    disk_name = "快照模型测试"
  }
  vpc_attributes {
    vpc_id     = alicloud_vpc.创建VPC.id
    vswitch_id = alicloud_vswitch.创建交换机.id
  }
  instance_name     = "快照模型测试"
  instance_type     = "ecs.u1-c1m1.large"
  image_id          = "aliyun_3_x64_20G_alibase_20251030.vhd"
  payment_type      = "PayAsYouGo"
  security_group_id = alicloud_security_group.创建安全组.id
  status            = "Running"
}

resource "alicloud_ecs_disk" "查询磁盘" {
  disk_name = "快照模型测试"
  category  = "cloud_essd"
}


`, name)
}

// Case 加密复制_换账号测试 12496
func TestAccAliCloudEcsSnapshot_basic12496(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsSnapshotMap12496)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsSnapshotBasicDependence12496)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description":        "copy_test",
					"snapshot_name":      name,
					"retention_days":     "1",
					"source_snapshot_id": "${alicloud_ecs_snapshot.创建快照.id}",
					"encrypted":          "true",
					"source_region_id":   "${var.source_region_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":        "copy_test",
						"snapshot_name":      name,
						"retention_days":     "1",
						"source_snapshot_id": CHECKSET,
						"encrypted":          "true",
						"source_region_id":   CHECKSET,
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
				ImportStateVerifyIgnore: []string{"lock_mode"},
			},
		},
	})
}

var AlicloudEcsSnapshotMap12496 = map[string]string{
	"status":      " ",
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEcsSnapshotBasicDependence12496(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "source_region_id" {
  default = "cn-hangzhou"
}

variable "zone_id" {
  default = "cn-hangzhou-h"
}

variable "region_id" {
  default = "cn-hangzhou"
}

resource "alicloud_vpc" "创建VPC" {
  is_default = false
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "创建交换机" {
  vpc_id     = alicloud_vpc.创建VPC.id
  zone_id    = var.zone_id
  cidr_block = "172.16.0.0/24"
}

resource "alicloud_security_group" "创建安全组" {
  vpc_id              = alicloud_vpc.创建VPC.id
  security_group_type = "normal"
  description         = "sg"
  security_group_name = "sg_name"
}

resource "alicloud_ecs_instance" "创建实例" {
  system_disk {
    size      = "20"
    category  = "cloud_essd"
    disk_name = "复制测试"
  }
  vpc_attributes {
    vpc_id     = alicloud_vpc.创建VPC.id
    vswitch_id = alicloud_vswitch.创建交换机.id
  }
  instance_name     = "复制测试"
  instance_type     = "ecs.u1-c1m1.large"
  image_id          = "aliyun_3_x64_20G_alibase_20251030.vhd"
  payment_type      = "PayAsYouGo"
  security_group_id = alicloud_security_group.创建安全组.id
  status            = "Running"
}

resource "alicloud_ecs_disk" "查询磁盘" {
  disk_name = "复制测试"
  category  = "cloud_essd"
}

resource "alicloud_ecs_snapshot" "创建快照" {
  retention_days = "1"
  snapshot_name  = "复制测试"
  disk_id        = alicloud_ecs_disk.查询磁盘.id
}


`, name)
}

// Case 类型_换账号测试 12516
func TestAccAliCloudEcsSnapshot_basic12516(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsSnapshotMap12516)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsSnapshotBasicDependence12516)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description":       "model_test",
					"disk_id":           "${alicloud_ecs_disk.查询磁盘.id}",
					"snapshot_name":     name,
					"retention_days":    "3",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "model_test",
						"disk_id":           CHECKSET,
						"snapshot_name":     name,
						"retention_days":    "3",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":   "修改描述",
					"snapshot_name": name + "_update",
					"category":      "standard",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "修改描述",
						"snapshot_name": name + "_update",
						"category":      "standard",
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
				ImportStateVerifyIgnore: []string{"lock_mode"},
			},
		},
	})
}

var AlicloudEcsSnapshotMap12516 = map[string]string{
	"status":      " ",
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEcsSnapshotBasicDependence12516(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "zone_id" {
  default = "cn-hangzhou-h"
}

variable "destination_region_id" {
  default = "cn-qingdao"
}

variable "region_id" {
  default = "cn-hangzhou"
}

resource "alicloud_vpc" "创建VPC" {
  is_default = false
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "创建交换机" {
  vpc_id     = alicloud_vpc.创建VPC.id
  zone_id    = var.zone_id
  cidr_block = "172.16.0.0/24"
}

resource "alicloud_security_group" "创建安全组" {
  vpc_id              = alicloud_vpc.创建VPC.id
  security_group_type = "normal"
  description         = "sg"
  security_group_name = "sg_name"
}

resource "alicloud_ecs_instance" "创建实例" {
  system_disk {
    size      = "20"
    category  = "cloud_essd"
    disk_name = "系统盘"
  }
  vpc_attributes {
    vpc_id     = alicloud_vpc.创建VPC.id
    vswitch_id = alicloud_vswitch.创建交换机.id
  }
  instance_name     = "加密测试"
  instance_type     = "ecs.u1-c1m1.large"
  image_id          = "aliyun_3_x64_20G_alibase_20251030.vhd"
  payment_type      = "PayAsYouGo"
  security_group_id = alicloud_security_group.创建安全组.id
  status            = "Running"
  data_disk {
    size                 = "20"
    disk_name            = "数据盘"
    category             = "cloud_essd"
    delete_with_instance = true
    encrypted            = true
  }
}

resource "alicloud_ecs_disk" "查询磁盘" {
  disk_name = "数据盘"
  category  = "cloud_essd"
}


`, name)
}

// Case 加密_换账号测试 12509
func TestAccAliCloudEcsSnapshot_basic12509(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsSnapshotMap12509)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsSnapshotBasicDependence12509)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description":       "model_test",
					"disk_id":           "${alicloud_ecs_disk.查询磁盘.id}",
					"snapshot_name":     name,
					"retention_days":    "3",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"encrypted":         "true",
					"kms_key_id":        "${alicloud_ecs_disk.查询磁盘.kms_key_id}",
					"category":          "standard",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "model_test",
						"disk_id":           CHECKSET,
						"snapshot_name":     name,
						"retention_days":    "3",
						"resource_group_id": CHECKSET,
						"encrypted":         "true",
						"kms_key_id":        CHECKSET,
						"category":          "standard",
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
				ImportStateVerifyIgnore: []string{"lock_mode"},
			},
		},
	})
}

var AlicloudEcsSnapshotMap12509 = map[string]string{
	"status":      " ",
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEcsSnapshotBasicDependence12509(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "zone_id" {
  default = "cn-hangzhou-h"
}

variable "destination_region_id" {
  default = "cn-qingdao"
}

variable "region_id" {
  default = "cn-hangzhou"
}

resource "alicloud_vpc" "创建VPC" {
  is_default = false
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "创建交换机" {
  vpc_id     = alicloud_vpc.创建VPC.id
  zone_id    = var.zone_id
  cidr_block = "172.16.0.0/24"
}

resource "alicloud_security_group" "创建安全组" {
  vpc_id              = alicloud_vpc.创建VPC.id
  security_group_type = "normal"
  description         = "sg"
  security_group_name = "sg_name"
}

resource "alicloud_ecs_instance" "创建实例" {
  system_disk {
    size      = "20"
    category  = "cloud_essd"
    disk_name = "系统盘"
  }
  vpc_attributes {
    vpc_id     = alicloud_vpc.创建VPC.id
    vswitch_id = alicloud_vswitch.创建交换机.id
  }
  instance_name     = "加密测试"
  instance_type     = "ecs.u1-c1m1.large"
  image_id          = "aliyun_3_x64_20G_alibase_20251030.vhd"
  payment_type      = "PayAsYouGo"
  security_group_id = alicloud_security_group.创建安全组.id
  status            = "Running"
  data_disk {
    size                 = "20"
    disk_name            = "数据盘"
    category             = "cloud_essd"
    delete_with_instance = true
    encrypted            = true
  }
}

resource "alicloud_ecs_disk" "查询磁盘" {
  disk_name = "数据盘"
  category  = "cloud_essd"
}


`, name)
}

// Case testaaa 12243
func TestAccAliCloudEcsSnapshot_basic12243(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsSnapshotMap12243)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsSnapshotBasicDependence12243)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description":       "xinglv_test",
					"disk_id":           "${alicloud_ecs_disk.查询磁盘.id}",
					"snapshot_name":     name,
					"retention_days":    "3",
					"source_region_id":  "${var.region_id}",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "xinglv_test",
						"disk_id":           CHECKSET,
						"snapshot_name":     name,
						"retention_days":    "3",
						"source_region_id":  CHECKSET,
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":     "修改描述",
					"snapshot_name":   name + "_update",
					"retention_days":  "4",
					"lock_mode":       "compliance",
					"cool_off_period": "1",
					"lock_duration":   "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":     "修改描述",
						"snapshot_name":   name + "_update",
						"retention_days":  "4",
						"lock_mode":       "compliance",
						"cool_off_period": "1",
						"lock_duration":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lock_status": "unlocked",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lock_status": "unlocked",
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
				ImportStateVerifyIgnore: []string{"lock_mode"},
			},
		},
	})
}

var AlicloudEcsSnapshotMap12243 = map[string]string{
	"status":      " ",
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEcsSnapshotBasicDependence12243(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "zone_id" {
  default = "cn-hangzhou-i"
}

variable "region_id" {
  default = "cn-hangzhou"
}

resource "alicloud_vpc" "创建VPC" {
  is_default = false
}

resource "alicloud_vswitch" "创建交换机" {
  vpc_id       = alicloud_vpc.创建VPC.id
  vswitch_name = "快照模型测试"
  zone_id      = var.zone_id
  cidr_block   = "172.19.192.0/20"
}

resource "alicloud_security_group" "创建安全组" {
  security_group_name = "快照模型测试"
  vpc_id              = alicloud_vpc.创建VPC.id
  security_group_type = "normal"
}

resource "alicloud_ecs_instance" "创建实例" {
  system_disk {
    size      = "20"
    category  = "cloud_essd"
    disk_name = "快照模型测试"
  }
  vpc_attributes {
    vpc_id     = alicloud_vpc.创建VPC.id
    vswitch_id = alicloud_vswitch.创建交换机.id
  }
  instance_name     = "快照模型测试"
  instance_type     = "ecs.u1-c1m1.large"
  image_id          = "aliyun_3_x64_20G_alibase_20251030.vhd"
  payment_type      = "PayAsYouGo"
  security_group_id = alicloud_security_group.创建安全组.id
  status            = "Running"
}

resource "alicloud_ecs_disk" "查询磁盘" {
  disk_name = "快照模型测试"
  category  = "cloud_essd"
}


`, name)
}

// Test Ecs Snapshot. <<< Resource test cases, automatically generated.

func TestAccAliCloudECSSnapshot_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AliCloudEcsSnapshotMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%secssnapshot%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEcsSnapshotBasicDependence0)
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
					"disk_id": "${alicloud_ecs_disk_attachment.default.disk_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"disk_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.groups.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"snapshot_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"snapshot_name": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Snapshot",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Snapshot",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force"},
			},
		},
	})
}

func TestAccAliCloudECSSnapshot_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AliCloudEcsSnapshotMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%secssnapshot%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEcsSnapshotBasicDependence0)
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
					"disk_id":           "${alicloud_ecs_disk_attachment.default.disk_id}",
					"category":          "flash",
					"retention_days":    "50",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.groups.1.id}",
					"snapshot_name":     name,
					"description":       name,
					"force":             "true",
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Snapshot",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"disk_id":           CHECKSET,
						"category":          "flash",
						"retention_days":    "50",
						"resource_group_id": CHECKSET,
						"snapshot_name":     name,
						"description":       name,
						"tags.%":            "2",
						"tags.Created":      "TF",
						"tags.For":          "Snapshot",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force"},
			},
		},
	})
}

func TestAccAliCloudECSSnapshot_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AliCloudEcsSnapshotMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%secssnapshot%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEcsSnapshotBasicDependence1)
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
					"disk_id":        "${alicloud_ecs_disk_attachment.default.disk_id}",
					"retention_days": "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"disk_id":        CHECKSET,
						"retention_days": "50",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_days": "60",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_days": "60",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.groups.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"name": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Snapshot",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Snapshot",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force"},
			},
		},
	})
}

func TestAccAliCloudECSSnapshot_basic1_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_snapshot.default"
	ra := resourceAttrInit(resourceId, AliCloudEcsSnapshotMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsSnapshot")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%secssnapshot%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEcsSnapshotBasicDependence1)
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
					"disk_id":           "${alicloud_ecs_disk_attachment.default.disk_id}",
					"category":          "standard",
					"retention_days":    "50",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.groups.1.id}",
					"name":              name,
					"description":       name,
					"force":             "true",
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Snapshot",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"disk_id":           CHECKSET,
						"category":          "standard",
						"retention_days":    "50",
						"resource_group_id": CHECKSET,
						"name":              name,
						"description":       name,
						"tags.%":            "2",
						"tags.Created":      "TF",
						"tags.For":          "Snapshot",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force"},
			},
		},
	})
}

var AliCloudEcsSnapshotMap0 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
	"category":    CHECKSET,
	"status":      CHECKSET,
}

func AliCloudEcsSnapshotBasicDependence0(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_resource_manager_resource_groups" "default" {
  		status = "OK"
	}

	data "alicloud_zones" "default" {
  		available_disk_category     = "cloud_essd"
  		available_resource_creation = "VSwitch"
	}
	
	data "alicloud_images" "default" {
  		most_recent = true
  		owners      = "system"
	}
	
	data "alicloud_instance_types" "default" {
  		availability_zone    = data.alicloud_zones.default.zones.0.id
  		image_id             = data.alicloud_images.default.images.0.id
  		system_disk_category = "cloud_essd"
	}
	
	resource "alicloud_vpc" "default" {
  		vpc_name   = var.name
  		cidr_block = "192.168.0.0/16"
	}
	
	resource "alicloud_vswitch" "default" {
		vswitch_name = var.name
  		vpc_id       = alicloud_vpc.default.id
  		cidr_block   = "192.168.192.0/24"
  		zone_id      = data.alicloud_zones.default.zones.0.id
	}
	
	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}
	
	resource "alicloud_instance" "default" {
  		image_id                   = data.alicloud_images.default.images.0.id
  		instance_type              = data.alicloud_instance_types.default.instance_types.0.id
  		security_groups            = alicloud_security_group.default.*.id
  		internet_charge_type       = "PayByTraffic"
  		internet_max_bandwidth_out = "10"
  		availability_zone          = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  		instance_charge_type       = "PostPaid"
  		system_disk_category       = "cloud_essd"
  		vswitch_id                 = alicloud_vswitch.default.id
  		instance_name              = var.name
		data_disks {
			category = "cloud_essd"
			size     = 20
  		}
	}
	
	resource "alicloud_ecs_disk" "default" {
  		disk_name = var.name
  		zone_id   = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  		category  = "cloud_essd"
  		size      = 500
	}
	
	resource "alicloud_ecs_disk_attachment" "default" {
  		disk_id     = alicloud_ecs_disk.default.id
  		instance_id = alicloud_instance.default.id
	}
`, name)
}

func AliCloudEcsSnapshotBasicDependence1(name string) string {
	return fmt.Sprintf(`
	variable "name" {
 		default = "%s"
	}

	data "alicloud_resource_manager_resource_groups" "default" {
 		status = "OK"
	}

	data "alicloud_zones" "default" {
 		available_disk_category     = "cloud_efficiency"
 		available_resource_creation = "VSwitch"
	}

	data "alicloud_images" "default" {
 		most_recent = true
 		owners      = "system"
	}

	data "alicloud_instance_types" "default" {
 		availability_zone = data.alicloud_zones.default.zones.0.id
 		image_id          = data.alicloud_images.default.images.0.id
	}

	resource "alicloud_vpc" "default" {
 		vpc_name   = var.name
 		cidr_block = "192.168.0.0/16"
	}

	resource "alicloud_vswitch" "default" {
		vswitch_name = var.name
 		vpc_id       = alicloud_vpc.default.id
 		cidr_block   = "192.168.192.0/24"
 		zone_id      = data.alicloud_zones.default.zones.0.id
	}

	resource "alicloud_security_group" "default" {
 		name   = var.name
 		vpc_id = alicloud_vpc.default.id
	}

	resource "alicloud_instance" "default" {
 		image_id                   = data.alicloud_images.default.images.0.id
 		instance_type              = data.alicloud_instance_types.default.instance_types.0.id
 		security_groups            = alicloud_security_group.default.*.id
 		internet_charge_type       = "PayByTraffic"
 		internet_max_bandwidth_out = "10"
 		availability_zone          = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
 		instance_charge_type       = "PostPaid"
 		system_disk_category       = "cloud_efficiency"
 		vswitch_id                 = alicloud_vswitch.default.id
 		instance_name              = var.name
	}

	resource "alicloud_ecs_disk" "default" {
 		disk_name = var.name
 		zone_id   = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
 		category  = "cloud_efficiency"
 		size      = 500
	}

	resource "alicloud_ecs_disk_attachment" "default" {
 		instance_id = alicloud_instance.default.id
 		disk_id     = alicloud_ecs_disk.default.id
	}
`, name)
}

// lintignore: R001
func TestUnitAliCloudEcsSnapshot(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	d, _ := schema.InternalMap(p["alicloud_ecs_snapshot"].Schema).Data(nil, nil)
	dCreate, _ := schema.InternalMap(p["alicloud_ecs_snapshot"].Schema).Data(nil, nil)
	dCreate.MarkNewResource()
	for key, value := range map[string]interface{}{
		"disk_id":           "disk_id",
		"category":          `standard`,
		"name":              "name",
		"description":       "description",
		"resource_group_id": "resource_group_id",
		"tags": map[string]string{
			"Created": "TF",
			"For":     "Test",
		},
	} {
		err := dCreate.Set(key, value)
		assert.Nil(t, err)
		err = d.Set(key, value)
		assert.Nil(t, err)
	}
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)
	ReadMockResponse := map[string]interface{}{
		"Snapshots": map[string]interface{}{
			"Snapshot": []interface{}{
				map[string]interface{}{
					"Category":                   "standard",
					"Description":                "description",
					"SourceDiskId":               "disk_id",
					"InstantAccess":              true,
					"InstantAccessRetentionDays": 20,
					"ResourceGroupId":            "resource_group_id",
					"RetentionDays":              20,
					"SnapshotName":               "snapshot_name",
					"Name":                       "snapshot_name",
					"Status":                     "accomplished",
					"SnapshotId":                 "MockSnapshotId",
				},
			},
		},
		"TagResources": map[string]interface{}{
			"TagResource": []interface{}{
				map[string]interface{}{
					"TagKey":   "Created",
					"TagValue": "TF",
				},
				map[string]interface{}{
					"TagKey":   "For",
					"TagValue": "Test",
				},
			},
		},
	}

	responseMock := map[string]func(errorCode string) (map[string]interface{}, error){
		"RetryError": func(errorCode string) (map[string]interface{}, error) {
			return nil, &tea.SDKError{
				Code:       String(errorCode),
				Data:       String(errorCode),
				Message:    String(errorCode),
				StatusCode: tea.Int(400),
			}
		},
		"NotFoundError": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"NoRetryError": func(errorCode string) (map[string]interface{}, error) {
			return nil, &tea.SDKError{
				Code:       String(errorCode),
				Data:       String(errorCode),
				Message:    String(errorCode),
				StatusCode: tea.Int(400),
			}
		},
		"CreateNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			result["SnapshotId"] = "MockSnapshotId"
			return result, nil
		},
		"UpdateNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"DeleteNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"ReadNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"ReadDescribeEcsSnapshotNotFound": func(errorCode string) (map[string]interface{}, error) {
			result := map[string]interface{}{
				"Snapshots": map[string]interface{}{
					"Snapshot": []interface{}{},
				},
			}
			return result, nil
		},
	}
	// Create
	t.Run("CreateClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewEcsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})
		err := resourceAliCloudEcsSnapshotCreate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("CreateAbnormal", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotCreate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("CreateNormal", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotCreate(dCreate, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	t.Run("CreateNonRetryableError", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotCreate(dCreate, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	// Set ID for Update and Delete Method
	d.SetId("MockSnapshotId")
	// Update
	t.Run("UpdateClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewEcsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})

		err := resourceAliCloudEcsSnapshotUpdate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("UpdateHpcClusterAttributeAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"description", "snapshot_name", "name"} {
			switch p["alicloud_ecs_snapshot"].Schema[key].Type {
			case schema.TypeString:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: d.Get(key).(string), New: d.Get(key).(string) + "_update"})
			case schema.TypeBool:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.FormatBool(d.Get(key).(bool)), New: strconv.FormatBool(true)})
			case schema.TypeMap:
				diff.SetAttribute("tags.%", &terraform.ResourceAttrDiff{Old: "0", New: "2"})
				diff.SetAttribute("tags.For", &terraform.ResourceAttrDiff{Old: "", New: "Test"})
				diff.SetAttribute("tags.Created", &terraform.ResourceAttrDiff{Old: "", New: "TF"})
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_ecs_snapshot"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateHpcClusterAttributeTagsAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"description", "snapshot_name", "name", "tags"} {
			switch p["alicloud_ecs_snapshot"].Schema[key].Type {
			case schema.TypeString:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: d.Get(key).(string), New: d.Get(key).(string) + "_update"})
			case schema.TypeBool:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.FormatBool(d.Get(key).(bool)), New: strconv.FormatBool(true)})
			case schema.TypeMap:
				diff.SetAttribute("tags.%", &terraform.ResourceAttrDiff{Old: "0", New: "2"})
				diff.SetAttribute("tags.For", &terraform.ResourceAttrDiff{Old: "", New: "Test"})
				diff.SetAttribute("tags.Created", &terraform.ResourceAttrDiff{Old: "", New: "TF"})
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_ecs_snapshot"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateHpcClusterAttributeNormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"description", "snapshot_name", "name"} {
			switch p["alicloud_ecs_snapshot"].Schema[key].Type {
			case schema.TypeString:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: d.Get(key).(string), New: d.Get(key).(string) + "_update"})
			case schema.TypeBool:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.FormatBool(d.Get(key).(bool)), New: strconv.FormatBool(true)})
			case schema.TypeMap:
				diff.SetAttribute("tags.%", &terraform.ResourceAttrDiff{Old: "0", New: "2"})
				diff.SetAttribute("tags.For", &terraform.ResourceAttrDiff{Old: "", New: "Test"})
				diff.SetAttribute("tags.Created", &terraform.ResourceAttrDiff{Old: "", New: "TF"})
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_ecs_snapshot"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	// Delete
	t.Run("DeleteClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewEcsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})
		err := resourceAliCloudEcsSnapshotDelete(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("DeleteMockAbnormal", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotDelete(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("DeleteMockNormal", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		patcheDescribeEcsSnapshot := gomonkey.ApplyMethod(reflect.TypeOf(&EcsService{}), "DescribeEcsSnapshot", func(*EcsService, string) (map[string]interface{}, error) {
			return responseMock["RetryError"]("InvalidFilterKey.NotFound")
		})
		err := resourceAliCloudEcsSnapshotDelete(d, rawClient)
		patches.Reset()
		patcheDescribeEcsSnapshot.Reset()
		assert.NotNil(t, err)
	})

	//Read
	t.Run("ReadDescribeEcsSnapshotNotFound", func(t *testing.T) {
		patcheDorequest := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			NotFoundFlag := true
			noRetryFlag := false
			if NotFoundFlag {
				return responseMock["ReadDescribeEcsSnapshotNotFound"]("")
			} else if noRetryFlag {
				return responseMock["NoRetryError"]("NoRetryError")
			}
			return responseMock["ReadNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotRead(d, rawClient)
		patcheDorequest.Reset()
		assert.Nil(t, err)
	})

	t.Run("ReadDescribeEcsHpcClusterAbnormal", func(t *testing.T) {
		patcheDorequest := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			retryFlag := false
			noRetryFlag := true
			if retryFlag {
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["ReadNormal"]("")
		})
		err := resourceAliCloudEcsSnapshotRead(d, rawClient)
		patcheDorequest.Reset()
		assert.NotNil(t, err)
	})
}
