// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Mongodb Backup. >>> Resource test cases, automatically generated.
// Case bgg-test 11938
func TestAccAliCloudMongodbBackup_basic11938(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_mongodb_backup.default"
	ra := resourceAttrInit(resourceId, AlicloudMongodbBackupMap11938)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MongodbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMongodbBackup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmongodb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMongodbBackupBasicDependence11938)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shanghai"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"backup_method":           "Snapshot",
					"db_instance_id":          "${alicloud_mongodb_instance.defaultpUQD63.id}",
					"backup_retention_period": "7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"backup_method":           "Snapshot",
						"db_instance_id":          CHECKSET,
						"backup_retention_period": "7",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backup_retention_period"},
			},
		},
	})
}

var AlicloudMongodbBackupMap11938 = map[string]string{
	"status":    CHECKSET,
	"backup_id": CHECKSET,
}

func AlicloudMongodbBackupBasicDependence11938(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "zone_id" {
    default = "cn-shanghai-b"
}

variable "region_id" {
    default = "cn-shanghai"
}

variable "ipv4网段-b" {
    default = "10.0.0.0/24"
}

resource "alicloud_vpc" "defaulttAyVQI" {
        cidr_block = "10.0.0.0/8"
        vpc_name = "bgg-vpc-shanghai-b"
}

resource "alicloud_vswitch" "defaultGQF7Gl" {
        vpc_id = "${alicloud_vpc.defaulttAyVQI.id}"
        zone_id = "${var.zone_id}"
        cidr_block = "${var.ipv4网段-b}"
}

resource "alicloud_mongodb_instance" "defaultpUQD63" {
        engine_version = "5.0"
        storage_type = "cloud_essd1"
        vswitch_id = "${alicloud_vswitch.defaultGQF7Gl.id}"
        db_instance_storage = "20"
        vpc_id = "${alicloud_vpc.defaulttAyVQI.id}"
        db_instance_class = "mdb.shard.4x.large.d"
        storage_engine = "WiredTiger"
        network_type = "VPC"
        zone_id = "${var.zone_id}"
        replication_factor = "3"
        readonly_replicas = "0"
}


`, name)
}

// Case bgg-test_副本1764834653409 11969
func TestAccAliCloudMongodbBackup_basic11969(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_mongodb_backup.default"
	ra := resourceAttrInit(resourceId, AlicloudMongodbBackupMap11969)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MongodbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMongodbBackup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmongodb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMongodbBackupBasicDependence11969)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shanghai"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"backup_method":           "Snapshot",
					"db_instance_id":          "dds-uf6a01eb067a5204",
					"backup_retention_period": "7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"backup_method":           "Snapshot",
						"db_instance_id":          "dds-uf6a01eb067a5204",
						"backup_retention_period": "7",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backup_retention_period"},
			},
		},
	})
}

var AlicloudMongodbBackupMap11969 = map[string]string{
	"status":    CHECKSET,
	"backup_id": CHECKSET,
}

func AlicloudMongodbBackupBasicDependence11969(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "zone_id" {
  default = "cn-shanghai-b"
}

variable "region_id" {
  default = "cn-shanghai"
}

variable "other_zone_id" {
  default = "cn-shanghai-g"
}

resource "alicloud_vpc" "defaultIzCngN" {
  cidr_block = "10.0.0.0/8"
  vpc_name   = "bgg-vpc-shanghai-b"
}

resource "alicloud_vswitch" "defaultLWEsPM" {
  vpc_id     = alicloud_vpc.defaultIzCngN.id
  zone_id    = var.zone_id
  cidr_block = "10.0.0.0/24"
}

resource "alicloud_mongodb_instance" "default7eOftZ" {
  engine_version = "4.4"
  storage_type   = "cloud_essd1"
  vswitch_id     = alicloud_vswitch.defaultLWEsPM.id
  vpc_id         = alicloud_vpc.defaultIzCngN.id
  storage_engine = "WiredTiger"
  network_type   = "VPC"
  zone_id        = var.zone_id
}


`, name)
}

// Test Mongodb Backup. <<< Resource test cases, automatically generated.
