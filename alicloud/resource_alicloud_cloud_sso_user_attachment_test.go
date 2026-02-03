package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test CloudSSO UserAttachment. >>> Resource test cases, automatically generated.
// Test CloudSso UserAttachment. >>> Resource test cases, automatically generated.
// Case UserAttachment1_online 10386
func TestAccAliCloudCloudSsoUserAttachment_basic10386(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_sso_user_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudSsoUserAttachmentMap10386)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudSSOServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudSsoUserAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudsso%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudSsoUserAttachmentBasicDependence10386)
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
					"directory_id": "${alicloud_cloud_sso_directory.defaultxZ0P6b.id}",
					"user_id":      "${alicloud_cloud_sso_user.defaultmzBZRH.user_id}",
					"group_id":     "${alicloud_cloud_sso_group.default8AFioW.group_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"directory_id": CHECKSET,
						"user_id":      CHECKSET,
						"group_id":     CHECKSET,
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

var AlicloudCloudSsoUserAttachmentMap10386 = map[string]string{}

func AlicloudCloudSsoUserAttachmentBasicDependence10386(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "group_name" {
  default = "rtgnprod1"
}

variable "user_name" {
  default = "rtunprod1"
}

variable "directory_name" {
  default = "rtdnonline1"
}

resource "alicloud_cloud_sso_directory" "defaultxZ0P6b" {
  directory_global_access_status = "Disabled"
  password_policy {
    min_password_length          = "8"
    min_password_different_chars = "8"
    max_password_age             = "90"
    password_reuse_prevention    = "1"
    max_login_attempts           = "5"
    max_password_length          = "32"
  }
  directory_name = var.directory_name
}

resource "alicloud_cloud_sso_user" "defaultmzBZRH" {
  user_name    = var.user_name
  directory_id = alicloud_cloud_sso_directory.defaultxZ0P6b.id
}

resource "alicloud_cloud_sso_group" "default8AFioW" {
  group_name   = var.group_name
  directory_id = alicloud_cloud_sso_directory.defaultxZ0P6b.id
}


`, name)
}

// Case UserAttachment1 10105
func TestAccAliCloudCloudSsoUserAttachment_basic10105(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_sso_user_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudSsoUserAttachmentMap10105)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudSSOServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudSsoUserAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudsso%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudSsoUserAttachmentBasicDependence10105)
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
					"directory_id": "${alicloud_cloud_sso_directory.defaultxZ0P6b.id}",
					"user_id":      "${alicloud_cloud_sso_user.defaultmzBZRH.user_id}",
					"group_id":     "${alicloud_cloud_sso_group.default8AFioW.group_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"directory_id": CHECKSET,
						"user_id":      CHECKSET,
						"group_id":     CHECKSET,
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

var AlicloudCloudSsoUserAttachmentMap10105 = map[string]string{}

func AlicloudCloudSsoUserAttachmentBasicDependence10105(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "group_name" {
  default = "rtgnprod1"
}

variable "user_name" {
  default = "rtunprod1"
}

variable "directory_name" {
  default = "rtdnprod1"
}

resource "alicloud_cloud_sso_directory" "defaultxZ0P6b" {
  directory_global_access_status = "Disabled"
  password_policy {
    min_password_length          = "8"
    min_password_different_chars = "8"
    max_password_age             = "90"
    password_reuse_prevention    = "1"
    max_login_attempts           = "5"
    max_password_length          = "32"
  }
  directory_name = var.directory_name
}

resource "alicloud_cloud_sso_user" "defaultmzBZRH" {
  user_name    = var.user_name
  directory_id = alicloud_cloud_sso_directory.defaultxZ0P6b.id
}

resource "alicloud_cloud_sso_group" "default8AFioW" {
  group_name   = var.group_name
  directory_id = alicloud_cloud_sso_directory.defaultxZ0P6b.id
}


`, name)
}

// Test CloudSso UserAttachment. <<< Resource test cases, automatically generated.

// Case UserAttachment1_online 10386
func TestAccAliCloudCloudSSOUserAttachment_basic10386(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_sso_user_attachment.default"
	ra := resourceAttrInit(resourceId, AliCloudCloudSSOUserAttachmentMap10386)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudSSOServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudSSOUserAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudsso%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudSSOUserAttachmentBasicDependence10386)
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
					"directory_id": "${local.directory_id}",
					"group_id":     "${alicloud_cloud_sso_group.default.group_id}",
					"user_id":      "${alicloud_cloud_sso_user.default.user_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"directory_id": CHECKSET,
						"group_id":     CHECKSET,
						"user_id":      CHECKSET,
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

var AliCloudCloudSSOUserAttachmentMap10386 = map[string]string{}

func AliCloudCloudSSOUserAttachmentBasicDependence10386(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_cloud_sso_directories" "default" {
	}

	resource "alicloud_cloud_sso_directory" "default" {
  		count          = length(data.alicloud_cloud_sso_directories.default.ids) > 0 ? 0 : 1
  		directory_name = var.name
	}

	resource "alicloud_cloud_sso_group" "default" {
  		directory_id = local.directory_id
  		group_name   = var.name
  		description  = var.name
	}

	resource "alicloud_cloud_sso_user" "default" {
  		directory_id = local.directory_id
  		user_name    = var.name
	}

	locals {
  		directory_id = length(data.alicloud_cloud_sso_directories.default.ids) > 0 ? data.alicloud_cloud_sso_directories.default.ids[0] : concat(alicloud_cloud_sso_directory.default.*.id, [""])[0]
	}
`, name)
}

// Test CloudSSO UserAttachment. <<< Resource test cases, automatically generated.
