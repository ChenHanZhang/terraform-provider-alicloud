// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ThreatDetection VulAutoConfig. >>> Resource test cases, automatically generated.
// Case VulAutoConfig 12047
func TestAccAliCloudThreatDetectionVulAutoConfig_basic12047(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_vul_auto_config.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionVulAutoConfigMap12047)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionVulAutoConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionVulAutoConfigBasicDependence12047)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"all_uuid":          "1",
					"start_time":        "1796903984000",
					"necessity":         "asap",
					"target_end_time":   "6",
					"type":              "cron",
					"enable":            "0",
					"target_start_time": "0",
					"period_unit":       "day",
					"need_snapshot":     "0",
					"rules":             "[{\\\"Type\\\":\\\"cve\\\",\\\"Name\\\":\\\"alilinux3:3:ALINUX3-SA-2025:0163\\\"}]",
					"snapshot_name":     "测试1",
					"snapshot_time":     "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"all_uuid":          "1",
						"start_time":        "1796903984000",
						"necessity":         "asap",
						"target_end_time":   "6",
						"type":              "cron",
						"enable":            "0",
						"target_start_time": "0",
						"period_unit":       "day",
						"need_snapshot":     "0",
						"rules":             CHECKSET,
						"snapshot_name":     "测试1",
						"snapshot_time":     "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"all_uuid":      "0",
					"start_time":    "1767024000000",
					"necessity":     "later",
					"type":          "once",
					"need_snapshot": "1",
					"snapshot_name": "测试",
					"snapshot_time": "3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"all_uuid":      "0",
						"start_time":    "1767024000000",
						"necessity":     "later",
						"type":          "once",
						"need_snapshot": "1",
						"snapshot_name": "测试",
						"snapshot_time": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"necessity":         "nntf",
					"target_end_time":   "12",
					"type":              "cron",
					"target_start_time": "6",
					"period_unit":       "week",
					"snapshot_time":     "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"necessity":         "nntf",
						"target_end_time":   "12",
						"type":              "cron",
						"target_start_time": "6",
						"period_unit":       "week",
						"snapshot_time":     "2",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"rules"},
			},
		},
	})
}

var AlicloudThreatDetectionVulAutoConfigMap12047 = map[string]string{}

func AlicloudThreatDetectionVulAutoConfigBasicDependence12047(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test ThreatDetection VulAutoConfig. <<< Resource test cases, automatically generated.
