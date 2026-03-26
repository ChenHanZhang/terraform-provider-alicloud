// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cms PrometheusView. >>> Resource test cases, automatically generated.
// Case prom-view 9280
func TestAccAliCloudCmsPrometheusView_basic9280(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_prometheus_view.default"
	ra := resourceAttrInit(resourceId, AlicloudCmsPrometheusViewMap9280)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCmsPrometheusView")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCmsPrometheusViewBasicDependence9280)
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
					"status":               "Running",
					"prometheus_view_name": name,
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"version":              "V2",
					"prometheus_instances": []map[string]interface{}{
						{
							"prometheus_instance_id": "c1e94e934622f4e42a7f88e1744e858b1",
							"region_id":              "cn-hangzhou",
							"user_id":                "1707387247477962",
						},
					},
					"workspace": "prometheus-1511928242963727",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":                 "Running",
						"prometheus_view_name":   name,
						"resource_group_id":      CHECKSET,
						"version":                "V2",
						"prometheus_instances.#": "1",
						"workspace":              CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"prometheus_view_name":  name + "_update",
					"auth_free_read_policy": "0.0.0.0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"prometheus_view_name":  name + "_update",
						"auth_free_read_policy": "0.0.0.0",
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

var AlicloudCmsPrometheusViewMap9280 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudCmsPrometheusViewBasicDependence9280(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Test Cms PrometheusView. <<< Resource test cases, automatically generated.
