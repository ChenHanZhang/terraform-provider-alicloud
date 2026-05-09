// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cms AggTaskGroup. >>> Resource test cases, automatically generated.
// Case aggTaskGroup 8019
func TestAccAliCloudCmsAggTaskGroup_basic8019(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_agg_task_group.default"
	ra := resourceAttrInit(resourceId, AlicloudCmsAggTaskGroupMap8019)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCmsAggTaskGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCmsAggTaskGroupBasicDependence8019)
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
					"target_prometheus_id":       "rw-ac05cec04ad891d38d3fe1c5c7c9",
					"source_prometheus_id":       "rw-ac05cec04ad891d38d3fe1c5c7c9",
					"agg_task_group_name":        name,
					"agg_task_group_config":      `---\ngroups:\n- name: test-group\n  interval: 1m\n  rules:\n  - record: cpu_user_util:1m\n    expr: cpu_user_util\n`,
					"override_if_exists":         "true",
					"max_retries":                "20",
					"status":                     "Running",
					"schedule_time_expr":         "@m",
					"description":                "描述测试",
					"agg_task_group_config_type": "RecordingRuleYaml",
					"schedule_mode":              "FixedRate",
					"precheck_string":            "{\\\"policy\\\":\\\"skip\\\",\\\"prometheusId\\\":\\\"rw-ac05cec04ad891d38d3fe1c5c7c9\\\",\\\"query\\\":\\\"scalar(sum(count_over_time(up{job=\\\\\\\"_arms/kubelet/cadvisor\\\\\\\"}[15s])) / 21)\\\",\\\"threshold\\\":1.0,\\\"timeout\\\":13,\\\"type\\\":\\\"none\\\"}",
					"from_time":                  "1727409939",
					"to_time":                    "0",
					"max_run_time_in_seconds":    "200",
					"delay":                      "31",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"target_prometheus_id":       "rw-ac05cec04ad891d38d3fe1c5c7c9",
						"source_prometheus_id":       "rw-ac05cec04ad891d38d3fe1c5c7c9",
						"agg_task_group_name":        name,
						"agg_task_group_config":      CHECKSET,
						"override_if_exists":         "true",
						"max_retries":                "20",
						"status":                     "Running",
						"schedule_time_expr":         "@m",
						"description":                "描述测试",
						"agg_task_group_config_type": "RecordingRuleYaml",
						"schedule_mode":              "FixedRate",
						"precheck_string":            CHECKSET,
						"from_time":                  "1727409939",
						"to_time":                    "0",
						"max_run_time_in_seconds":    "200",
						"delay":                      "31",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"agg_task_group_config":   `---\ngroups:\n- name: test-group-update\n  interval: 1m\n  rules:\n  - record: cpu_user_util:1m\n    expr: cpu_user_util\n`,
					"max_retries":             "10",
					"schedule_time_expr":      "@s",
					"description":             "描述更新",
					"precheck_string":         "{\\\"policy\\\":\\\"skip\\\",\\\"prometheusId\\\":\\\"rw-ac05cec04ad891d38d3fe1c5c7c9\\\",\\\"query\\\":\\\"scalar(sum(count_over_time(up{job=\\\\\\\"_arms/kubelet/cadvisor\\\\\\\"}[15s])) / 21)\\\",\\\"threshold\\\":0.9,\\\"timeout\\\":13,\\\"type\\\":\\\"none\\\"}",
					"from_time":               "1727400939",
					"max_run_time_in_seconds": "300",
					"delay":                   "33",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"agg_task_group_config":   CHECKSET,
						"max_retries":             "10",
						"schedule_time_expr":      "@s",
						"description":             "描述更新",
						"precheck_string":         CHECKSET,
						"from_time":               "1727400939",
						"max_run_time_in_seconds": "300",
						"delay":                   "33",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"agg_task_group_config": `---\ngroups:\n- name: test-group-update-2\n  interval: 1m\n  rules:\n  - record: cpu_user_util:1m\n    expr: cpu_user_util\n`,
					"status":                "Stopped",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"agg_task_group_config": CHECKSET,
						"status":                "Stopped",
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
				ImportStateVerifyIgnore: []string{"agg_task_group_config_type", "override_if_exists"},
			},
		},
	})
}

var AlicloudCmsAggTaskGroupMap8019 = map[string]string{
	"agg_task_group_id": CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudCmsAggTaskGroupBasicDependence8019(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Cms AggTaskGroup. <<< Resource test cases, automatically generated.
