// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ecs DiskDefaultKMSKey. >>> Resource test cases, automatically generated.
// Case 修改云盘加密默认密钥_有前置密钥生成 12754
func TestAccAliCloudEcsDiskDefaultKMSKey_basic12754(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_disk_default_k_m_s_key.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsDiskDefaultKMSKeyMap12754)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsDiskDefaultKMSKey")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsDiskDefaultKMSKeyBasicDependence12754)
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
					"kms_key_id": "${alicloud_kms_key.defaultC3EYIX.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"kms_key_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"kms_key_id": "${alicloud_kms_key.defaultDDxPUN.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"kms_key_id": CHECKSET,
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

var AlicloudEcsDiskDefaultKMSKeyMap12754 = map[string]string{}

func AlicloudEcsDiskDefaultKMSKeyBasicDependence12754(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cidr_block_vpc" {
  default = "172.16.0.0/12"
}

variable "cidr_block_vsw_k" {
  default = "172.17.0.0/16"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "zone_id_k" {
  default = "cn-hangzhou-k"
}

variable "zone_id_j" {
  default = "cn-hangzhou-j"
}

resource "alicloud_vpc" "default5RztkC" {
  is_default = false
  cidr_block = var.cidr_block_vpc
}

resource "alicloud_vswitch" "defaultTHeK3U" {
  vpc_id     = alicloud_vpc.default5RztkC.id
  zone_id    = var.zone_id_k
  cidr_block = var.cidr_block_vsw_k
}

resource "alicloud_kms_instance" "defaultLbatIU" {
  vpc_num         = "7"
  key_num         = "2000"
  renew_period    = "1"
  secret_num      = "0"
  product_version = "3"
  renew_status    = "AutoRenewal"
  vpc_id          = alicloud_vpc.default5RztkC.id
  vswitch_ids     = ["${alicloud_vswitch.defaultTHeK3U.id}"]
  zone_ids        = ["cn-hangzhou-k", "${var.zone_id_j}"]
  spec            = "2000"
}

resource "alicloud_kms_key" "defaultC3EYIX" {
  origin           = "Aliyun_KMS"
  status           = "Enabled"
  protection_level = "SOFTWARE"
  key_spec         = "Aliyun_AES_256"
  key_usage        = "ENCRYPT/DECRYPT"
  dkms_instance_id = alicloud_kms_instance.defaultLbatIU.id
}

resource "alicloud_kms_key" "defaultDDxPUN" {
  origin           = "Aliyun_KMS"
  status           = "Enabled"
  protection_level = "SOFTWARE"
  key_spec         = "Aliyun_AES_256"
  key_usage        = "ENCRYPT/DECRYPT"
  dkms_instance_id = alicloud_kms_instance.defaultLbatIU.id
  description      = "KmsKey02"
}


`, name)
}

// Case 无预置条件-Delete测试 12751
func TestAccAliCloudEcsDiskDefaultKMSKey_basic12751(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_disk_default_k_m_s_key.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsDiskDefaultKMSKeyMap12751)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsDiskDefaultKMSKey")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsDiskDefaultKMSKeyBasicDependence12751)
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
					"kms_key_id": "key-hzz69de01adkn3bltbvzf",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"kms_key_id": "key-hzz69de01adkn3bltbvzf",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"kms_key_id": "key-hzz69e078a6kbdssykgf6",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"kms_key_id": "key-hzz69e078a6kbdssykgf6",
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

var AlicloudEcsDiskDefaultKMSKeyMap12751 = map[string]string{}

func AlicloudEcsDiskDefaultKMSKeyBasicDependence12751(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 修改云盘加密默认密钥_有前置密钥生成_副本1776417819030 12760
func TestAccAliCloudEcsDiskDefaultKMSKey_basic12760(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_disk_default_k_m_s_key.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsDiskDefaultKMSKeyMap12760)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsDiskDefaultKMSKey")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsDiskDefaultKMSKeyBasicDependence12760)
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
					"kms_key_id": "${alicloud_kms_key.defaultC3EYIX.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"kms_key_id": CHECKSET,
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudEcsDiskDefaultKMSKeyMap12760 = map[string]string{}

func AlicloudEcsDiskDefaultKMSKeyBasicDependence12760(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cidr_block_vpc" {
  default = "172.16.0.0/12"
}

variable "cidr_block_vsw_k" {
  default = "172.17.0.0/16"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "zone_id_k" {
  default = "cn-hangzhou-k"
}

variable "zone_id_j" {
  default = "cn-hangzhou-j"
}

resource "alicloud_vpc" "default5RztkC" {
  cidr_block = var.cidr_block_vpc
}

resource "alicloud_vswitch" "defaultTHeK3U" {
  vpc_id     = alicloud_vpc.default5RztkC.id
  zone_id    = var.zone_id_k
  cidr_block = var.cidr_block_vsw_k
}

resource "alicloud_kms_instance" "defaultLbatIU" {
  vpc_num         = "7"
  key_num         = "2000"
  renew_period    = "1"
  secret_num      = "0"
  product_version = "3"
  renew_status    = "AutoRenewal"
  vpc_id          = alicloud_vpc.default5RztkC.id
  vswitch_ids     = ["${alicloud_vswitch.defaultTHeK3U.id}"]
  zone_ids        = ["cn-hangzhou-k", "${var.zone_id_j}"]
  spec            = "2000"
}

resource "alicloud_kms_key" "defaultC3EYIX" {
  origin           = "Aliyun_KMS"
  status           = "Enabled"
  protection_level = "SOFTWARE"
  key_spec         = "Aliyun_AES_256"
  key_usage        = "ENCRYPT/DECRYPT"
  dkms_instance_id = alicloud_kms_instance.defaultLbatIU.id
}


`, name)
}

// Test Ecs DiskDefaultKMSKey. <<< Resource test cases, automatically generated.
