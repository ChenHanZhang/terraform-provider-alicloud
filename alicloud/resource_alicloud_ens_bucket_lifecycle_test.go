// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ens BucketLifecycle. >>> Resource test cases, automatically generated.
// Case bucketlife测试-prefix 7122
func TestAccAliCloudEnsBucketLifecycle_basic7122(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_bucket_lifecycle.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsBucketLifecycleMap7122)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsBucketLifecycle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsBucketLifecycleBasicDependence7122)
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
					"bucket_name":               "${alicloud_ens_bucket.BUCKETNAME.id}",
					"prefix":                    "/test",
					"expiration_days":           "${var.expiration_days}",
					"status":                    "Enabled",
					"allow_same_action_overlap": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket_name":               CHECKSET,
						"prefix":                    "/test",
						"expiration_days":           CHECKSET,
						"status":                    "Enabled",
						"allow_same_action_overlap": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"prefix":          "${var.prefix}",
					"expiration_days": "${var.expiration_days_modify}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"prefix":          CHECKSET,
						"expiration_days": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"expiration_days":           "${var.expiration_days}",
					"allow_same_action_overlap": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"expiration_days":           CHECKSET,
						"allow_same_action_overlap": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"expiration_days": "${var.expiration_days_modify}",
					"status":          "Disabled",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"expiration_days": CHECKSET,
						"status":          "Disabled",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_same_action_overlap"},
			},
		},
	})
}

var AlicloudEnsBucketLifecycleMap7122 = map[string]string{}

func AlicloudEnsBucketLifecycleBasicDependence7122(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "bucket_name" {
  default = "pop-auto-lifecycle"
}

variable "rule_id" {
  default = "pop-auto-rule-id"
}

variable "expiration_days_modify" {
  default = <<EOF
1
EOF
}

variable "prefix" {
  default = "/popauto"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}

variable "expiration_days" {
  default = <<EOF
3
EOF
}

variable "expir_date" {
  default = "2024-08-15T00:00:00.000Z"
}

resource "alicloud_ens_bucket" "BUCKETNAME" {
  logical_bucket_type = "sink"
  bucket_acl          = "private"
  bucket_name         = var.bucket_name
  comment             = "bucketlifecycle资源测试"
  ens_region_id       = var.ens_region_id
}


`, name)
}

// Case bucketlife测试_全局 7507
func TestAccAliCloudEnsBucketLifecycle_basic7507(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_bucket_lifecycle.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsBucketLifecycleMap7507)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsBucketLifecycle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsBucketLifecycleBasicDependence7507)
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
					"bucket_name":               "${alicloud_ens_bucket.BUCKETNAME.id}",
					"status":                    "Enabled",
					"allow_same_action_overlap": "false",
					"created_before_date":       "${var.expir_date}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket_name":               CHECKSET,
						"status":                    "Enabled",
						"allow_same_action_overlap": "false",
						"created_before_date":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"created_before_date": "${var.expir_date_modify}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"created_before_date": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"expiration_days": "${var.expiration_days}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"expiration_days": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":          "Disabled",
					"expiration_days": "${var.expiration_days_modify}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":          "Disabled",
						"expiration_days": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_same_action_overlap"},
			},
		},
	})
}

var AlicloudEnsBucketLifecycleMap7507 = map[string]string{}

func AlicloudEnsBucketLifecycleBasicDependence7507(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "bucket_name" {
  default = "pop-auto-lifecycle-global"
}

variable "expiration_days_modify" {
  default = <<EOF
1
EOF
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}

variable "expiration_days" {
  default = <<EOF
3
EOF
}

variable "expir_date" {
  default = "2024-08-15T00:00:00.000Z"
}

variable "expir_date_modify" {
  default = "2024-08-16T00:00:00.000Z"
}

resource "alicloud_ens_bucket" "BUCKETNAME" {
  logical_bucket_type = "sink"
  bucket_acl          = "private"
  bucket_name         = var.bucket_name
  comment             = "bucketlifecycle资源测试"
  ens_region_id       = var.ens_region_id
}


`, name)
}

// Test Ens BucketLifecycle. <<< Resource test cases, automatically generated.
