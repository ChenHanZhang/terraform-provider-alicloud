// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ecs DiskEncryptionByDefault. >>> Resource test cases, automatically generated.
// Case 测试账号下云盘默认加密_正式_create_true 12733
func TestAccAliCloudEcsDiskEncryptionByDefault_basic12733(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_disk_encryption_by_default.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsDiskEncryptionByDefaultMap12733)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsDiskEncryptionByDefault")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsDiskEncryptionByDefaultBasicDependence12733)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"encrypted": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"encrypted": "true",
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
					"encrypted": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"encrypted": "false",
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

var AlicloudEcsDiskEncryptionByDefaultMap12733 = map[string]string{}

func AlicloudEcsDiskEncryptionByDefaultBasicDependence12733(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 测试账号下云盘默认加密_正式_create_false 12761
func TestAccAliCloudEcsDiskEncryptionByDefault_basic12761(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_disk_encryption_by_default.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsDiskEncryptionByDefaultMap12761)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsDiskEncryptionByDefault")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsDiskEncryptionByDefaultBasicDependence12761)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"encrypted": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"encrypted": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"encrypted": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"encrypted": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"encrypted": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"encrypted": "false",
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

var AlicloudEcsDiskEncryptionByDefaultMap12761 = map[string]string{}

func AlicloudEcsDiskEncryptionByDefaultBasicDependence12761(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Ecs DiskEncryptionByDefault. <<< Resource test cases, automatically generated.
