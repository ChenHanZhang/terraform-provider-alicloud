// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Adb LakeStorage. >>> Resource test cases, automatically generated.
// Case ADB湖存储 6570
func TestAccAliCloudAdbLakeStorage_basic6570(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_lake_storage.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbLakeStorageMap6570)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbLakeStorage")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbLakeStorageBasicDependence6570)
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
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.ADB.id}",
					"description":   "描述1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id": CHECKSET,
						"description":   "描述1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "描述2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "描述2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "描述3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "描述3",
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

var AlicloudAdbLakeStorageMap6570 = map[string]string{
	"lake_storage_id": CHECKSET,
	"create_time":     CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbLakeStorageBasicDependence6570(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "VPC" {
  dry_run    = false
  cidr_block = "172.16.0.0/12"
  vpc_name   = "APS结果集导出VPC"
}

resource "alicloud_vswitch" "VSWITCH" {
  vpc_id       = alicloud_vpc.VPC.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "APS结果集导出VSwitch"
}

resource "alicloud_adb_db_cluster_lake_version" "ADB" {
  storage_resource              = "24ACU"
  security_ips                  = "127.0.0.1"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.VPC.id
  vswitch_id                    = alicloud_vswitch.VSWITCH.id
  compute_resource              = "16ACU"
  db_cluster_version            = "5.0"
  payment_type                  = "Postpaid"
  enable_default_resource_group = true
  db_cluster_description        = "APS结果集导出OpenAPI集测20260206151502020202"
}

resource "alicloud_adb_lake_account" "LakeAccount" {
  db_cluster_id       = alicloud_adb_db_cluster_lake_version.ADB.id
  account_type        = "Super"
  account_name        = "apstest"
  account_password    = "apsTest2024!"
  account_description = "OPENAPI测试"
}


`, name)
}

// Test Adb LakeStorage. <<< Resource test cases, automatically generated.
