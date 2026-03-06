// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ens Bucket. >>> Resource test cases, automatically generated.
// Case Bucket资源测试-sink-jiaozuo-4-private 7121
func TestAccAliCloudEnsBucket_basic7121(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_bucket.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsBucketMap7121)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsBucket")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsBucketBasicDependence7121)
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
					"comment":             "bucket资源自动化测试",
					"bucket_acl":          "private",
					"bucket_name":         name,
					"logical_bucket_type": "sink",
					"ens_region_id":       "${var.ens_region_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment":             "bucket资源自动化测试",
						"bucket_acl":          "private",
						"bucket_name":         name,
						"logical_bucket_type": "sink",
						"ens_region_id":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"comment": "bucket资源自动化测试-修改",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment": "bucket资源自动化测试-修改",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dispatch_scope", "ens_region_id"},
			},
		},
	})
}

var AlicloudEnsBucketMap7121 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEnsBucketBasicDependence7121(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "bucket_name" {
  default = "pop-autotest-sink-jiaozuo4-private"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}


`, name)
}

// Case Bucket资源测试-sink-jiaozuo4-prw 7127
func TestAccAliCloudEnsBucket_basic7127(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_bucket.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsBucketMap7127)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsBucket")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsBucketBasicDependence7127)
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
					"comment":             "bucket资源自动化测试",
					"bucket_acl":          "public-read-write",
					"bucket_name":         name,
					"logical_bucket_type": "sink",
					"ens_region_id":       "${var.ens_region_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment":             "bucket资源自动化测试",
						"bucket_acl":          "public-read-write",
						"bucket_name":         name,
						"logical_bucket_type": "sink",
						"ens_region_id":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"comment": "bucket资源自动化测试-修改",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment": "bucket资源自动化测试-修改",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dispatch_scope", "ens_region_id"},
			},
		},
	})
}

var AlicloudEnsBucketMap7127 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEnsBucketBasicDependence7127(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "bucket_name" {
  default = "pop-autotest-sink-jiaozuo4-prw"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}


`, name)
}

// Case Bucket资源测试-sink-jiaozuo4-public-read 7128
func TestAccAliCloudEnsBucket_basic7128(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_bucket.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsBucketMap7128)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsBucket")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsBucketBasicDependence7128)
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
					"comment":             "bucket资源自动化测试",
					"bucket_acl":          "public-read",
					"bucket_name":         name,
					"logical_bucket_type": "sink",
					"ens_region_id":       "${var.ens_region_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment":             "bucket资源自动化测试",
						"bucket_acl":          "public-read",
						"bucket_name":         name,
						"logical_bucket_type": "sink",
						"ens_region_id":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"comment": "bucket资源自动化测试-修改",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment": "bucket资源自动化测试-修改",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dispatch_scope", "ens_region_id"},
			},
		},
	})
}

var AlicloudEnsBucketMap7128 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEnsBucketBasicDependence7128(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "bucket_name" {
  default = "pop-auto-sink-jiaozuo4-pub-rd"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}


`, name)
}

// Case Bucket资源测试-standard-private_domestic 7134
func TestAccAliCloudEnsBucket_basic7134(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_bucket.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsBucketMap7134)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsBucket")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsBucketBasicDependence7134)
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
					"comment":             "bucket资源自动化测试",
					"bucket_acl":          "private",
					"bucket_name":         name,
					"logical_bucket_type": "standard",
					"dispatch_scope":      "domestic",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment":             "bucket资源自动化测试",
						"bucket_acl":          "private",
						"bucket_name":         name,
						"logical_bucket_type": "standard",
						"dispatch_scope":      "domestic",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"comment": "bucket资源自动化测试-修改",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment": "bucket资源自动化测试-修改",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dispatch_scope", "ens_region_id"},
			},
		},
	})
}

var AlicloudEnsBucketMap7134 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEnsBucketBasicDependence7134(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "bucket_name" {
  default = "pop-autocase-standard-domestic"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}


`, name)
}

// Case Bucket资源测试-standard-private_oversea 8806
func TestAccAliCloudEnsBucket_basic8806(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_bucket.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsBucketMap8806)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsBucket")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsBucketBasicDependence8806)
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
					"comment":             "bucket资源自动化测试",
					"bucket_acl":          "private",
					"bucket_name":         name,
					"logical_bucket_type": "standard",
					"dispatch_scope":      "oversea",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment":             "bucket资源自动化测试",
						"bucket_acl":          "private",
						"bucket_name":         name,
						"logical_bucket_type": "standard",
						"dispatch_scope":      "oversea",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"comment": "bucket资源自动化测试-修改",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment": "bucket资源自动化测试-修改",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dispatch_scope", "ens_region_id"},
			},
		},
	})
}

var AlicloudEnsBucketMap8806 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEnsBucketBasicDependence8806(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "bucket_name" {
  default = "pop-autotest-standard-oversea-new"
}


`, name)
}

// Test Ens Bucket. <<< Resource test cases, automatically generated.
