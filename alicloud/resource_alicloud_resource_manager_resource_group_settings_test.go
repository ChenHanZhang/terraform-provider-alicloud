// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ResourceManager ResourceGroupSettings. >>> Resource test cases, automatically generated.
// Case resourceCase_20260625_XRVjjt 13025
func TestAccAliCloudResourceManagerResourceGroupSettings_basic13025(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_resource_manager_resource_group_settings.default"
	ra := resourceAttrInit(resourceId, AlicloudResourceManagerResourceGroupSettingsMap13025)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ResourceManagerServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeResourceManagerResourceGroupSettings")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccresourcemanager%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudResourceManagerResourceGroupSettingsBasicDependence13025)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_admin_setting_status":        "true",
					"resource_group_notification_setting_status": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_admin_setting_status":        "true",
						"resource_group_notification_setting_status": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_admin_setting_status":        "false",
					"resource_group_notification_setting_status": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_admin_setting_status":        "false",
						"resource_group_notification_setting_status": "false",
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

var AlicloudResourceManagerResourceGroupSettingsMap13025 = map[string]string{}

func AlicloudResourceManagerResourceGroupSettingsBasicDependence13025(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test ResourceManager ResourceGroupSettings. <<< Resource test cases, automatically generated.
