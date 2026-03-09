// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ram MFADevice. >>> Resource test cases, automatically generated.
// Case  MFADevice测试 9169
func TestAccAliCloudRamMFADevice_basic9169(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ram_mfadevice.default"
	ra := resourceAttrInit(resourceId, AlicloudRamMFADeviceMap9169)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RamServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRamMFADevice")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccram%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRamMFADeviceBasicDependence9169)
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
					"virtual_mfa_device_name": "zaijiuTestMFA1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"virtual_mfa_device_name": "zaijiuTestMFA1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"virtual_mfa_device_name"},
			},
		},
	})
}

var AlicloudRamMFADeviceMap9169 = map[string]string{
	"activate_date": CHECKSET,
}

func AlicloudRamMFADeviceBasicDependence9169(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case  MFADevice测试_副本1745496171386 10739
func TestAccAliCloudRamMFADevice_basic10739(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ram_mfadevice.default"
	ra := resourceAttrInit(resourceId, AlicloudRamMFADeviceMap10739)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RamServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRamMFADevice")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccram%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRamMFADeviceBasicDependence10739)
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
					"virtual_mfa_device_name": "zaijiuTestMFA1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"virtual_mfa_device_name": "zaijiuTestMFA1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"virtual_mfa_device_name"},
			},
		},
	})
}

var AlicloudRamMFADeviceMap10739 = map[string]string{
	"activate_date": CHECKSET,
}

func AlicloudRamMFADeviceBasicDependence10739(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Ram MFADevice. <<< Resource test cases, automatically generated.
