// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ens BucketAcl. >>> Resource test cases, automatically generated.
// Case acl_test 8778
func TestAccAliCloudEnsBucketAcl_basic8778(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_bucket_acl.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsBucketAclMap8778)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsBucketAcl")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsBucketAclBasicDependence8778)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"bucket_acl":  "public-read-write",
					"bucket_name": "${alicloud_ens_bucket.defaultmRFqwe.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket_acl":  "public-read-write",
						"bucket_name": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bucket_acl": "public-read",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket_acl": "public-read",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bucket_acl": "private",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket_acl": "private",
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

var AlicloudEnsBucketAclMap8778 = map[string]string{}

func AlicloudEnsBucketAclBasicDependence8778(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "bucket_name" {
  default = "pop-autocase-acl-test"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}

resource "alicloud_ens_bucket" "defaultmRFqwe" {
  comment             = "bucket资源自动化测试"
  bucket_acl          = "private"
  bucket_name         = var.bucket_name
  logical_bucket_type = "sink"
  ens_region_id       = var.ens_region_id
}


`, name)
}

// Test Ens BucketAcl. <<< Resource test cases, automatically generated.
