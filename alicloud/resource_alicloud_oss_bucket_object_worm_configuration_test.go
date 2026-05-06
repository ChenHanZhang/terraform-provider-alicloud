// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Oss BucketObjectWormConfiguration. >>> Resource test cases, automatically generated.
// Case 测试ObjectWorm_依赖资源 12777
func TestAccAliCloudOssBucketObjectWormConfiguration_basic12777(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_object_worm_configuration.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketObjectWormConfigurationMap12777)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketObjectWormConfiguration")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccoss%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketObjectWormConfigurationBasicDependence12777)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"bucket_name":         "${alicloud_oss_bucket.defaultQf8G0L.id}",
					"object_worm_enabled": "Enabled",
					"rule": []map[string]interface{}{
						{
							"default_retention": []map[string]interface{}{
								{
									"mode": "COMPLIANCE",
									"days": "1",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket_name":         CHECKSET,
						"object_worm_enabled": "Enabled",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule": []map[string]interface{}{
						{
							"default_retention": []map[string]interface{}{
								{
									"mode": "COMPLIANCE",
									"days": "2",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule": []map[string]interface{}{
						{
							"default_retention": []map[string]interface{}{
								{
									"mode":  "COMPLIANCE",
									"years": "1",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule": []map[string]interface{}{
						{
							"default_retention": []map[string]interface{}{
								{
									"mode":  "COMPLIANCE",
									"years": "2",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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

var AlicloudOssBucketObjectWormConfigurationMap12777 = map[string]string{}

func AlicloudOssBucketObjectWormConfigurationBasicDependence12777(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_oss_bucket" "defaultQf8G0L" {
  storage_class = "Standard"
}

resource "alicloud_oss_bucket_versioning" "defaultosxikW" {
  status = "Enabled"
  bucket = alicloud_oss_bucket.defaultQf8G0L.id
}


`, name)
}

// Case 测试ObjectWorm_Year 12778
func TestAccAliCloudOssBucketObjectWormConfiguration_basic12778(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_object_worm_configuration.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketObjectWormConfigurationMap12778)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketObjectWormConfiguration")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccoss%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketObjectWormConfigurationBasicDependence12778)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"bucket_name":         "${alicloud_oss_bucket.defaultQf8G0L.id}",
					"object_worm_enabled": "Enabled",
					"rule": []map[string]interface{}{
						{
							"default_retention": []map[string]interface{}{
								{
									"mode":  "COMPLIANCE",
									"years": "1",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket_name":         CHECKSET,
						"object_worm_enabled": "Enabled",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule": []map[string]interface{}{
						{},
					},
				}),
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

var AlicloudOssBucketObjectWormConfigurationMap12778 = map[string]string{}

func AlicloudOssBucketObjectWormConfigurationBasicDependence12778(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_oss_bucket" "defaultQf8G0L" {
  storage_class = "Standard"
}

resource "alicloud_oss_bucket_versioning" "defaultosxikW" {
  status = "Enabled"
  bucket = alicloud_oss_bucket.defaultQf8G0L.id
}


`, name)
}

// Test Oss BucketObjectWormConfiguration. <<< Resource test cases, automatically generated.
