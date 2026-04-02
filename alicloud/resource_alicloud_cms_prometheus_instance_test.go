// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cms PrometheusInstance. >>> Resource test cases, automatically generated.
// Case promInstance有rg、tag 8018
func TestAccAliCloudCmsPrometheusInstance_basic8018(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_prometheus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCmsPrometheusInstanceMap8018)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCmsPrometheusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCmsPrometheusInstanceBasicDependence8018)
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
					"status":                   "Running",
					"archive_duration":         "60",
					"resource_group_id":        "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"prometheus_instance_name": name,
					"auth_free_read_policy":    "1.1.1.1",
					"auth_free_write_policy":   "2.2.2.2",
					"storage_duration":         "30",
					"enable_auth_free_read":    "true",
					"enable_auth_free_write":   "true",
					"workspace":                "prometheus-1511928242963727",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":                   "Running",
						"archive_duration":         "60",
						"resource_group_id":        CHECKSET,
						"prometheus_instance_name": name,
						"auth_free_read_policy":    "1.1.1.1",
						"auth_free_write_policy":   "2.2.2.2",
						"storage_duration":         "30",
						"enable_auth_free_read":    "true",
						"enable_auth_free_write":   "true",
						"workspace":                CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"archive_duration":         "90",
					"resource_group_id":        "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"prometheus_instance_name": name + "_update",
					"enable_auth_free_read":    "false",
					"enable_auth_free_write":   "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"archive_duration":         "90",
						"resource_group_id":        CHECKSET,
						"prometheus_instance_name": name + "_update",
						"enable_auth_free_read":    "false",
						"enable_auth_free_write":   "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"archive_duration": "180",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"archive_duration": "180",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "Test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF-update",
						"tags.For":     "Test-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "0",
						"tags.Created": REMOVEKEY,
						"tags.For":     REMOVEKEY,
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

var AlicloudCmsPrometheusInstanceMap8018 = map[string]string{
	"payment_type": CHECKSET,
	"create_time":  CHECKSET,
	"region_id":    CHECKSET,
}

func AlicloudCmsPrometheusInstanceBasicDependence8018(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Test Cms PrometheusInstance. <<< Resource test cases, automatically generated.
