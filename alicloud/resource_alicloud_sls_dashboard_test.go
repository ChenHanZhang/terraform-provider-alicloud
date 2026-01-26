// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Sls Dashboard. >>> Resource test cases, automatically generated.
// Case Dashboard_Terraform 12338
func TestAccAliCloudSlsDashboard_basic12338(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_sls_dashboard.default"
	ra := resourceAttrInit(resourceId, AlicloudSlsDashboardMap12338)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SlsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSlsDashboard")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccsls%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSlsDashboardBasicDependence12338)
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
					"description": "description",
					"charts": []map[string]interface{}{
						{
							"action": map[string]interface{}{
								"\"name\"":   "open_logstore",
								"\"title\"":  "查看该省份日志",
								"\"events\"": "[{\\\"type\\\":\\\"click\\\",\\\"parameters\\\":{\\\"logstore\\\":\\\"access-log\\\",\\\"query\\\":\\\"ip_to_province(remote_addr): $${province}\\\",\\\"timeRange\\\":\\\"last_1h\\\"}}]",
							},
							"display": map[string]interface{}{
								"\"height\"": "5",
								"\"width\"":  "5",
								"\"xPos\"":   "0",
								"\"yPos\"":   "0",
								"\"xAxis\"":  "[\\\"province\\\"]",
								"\"yAxis\"":  "[\\\"pv\\\"]",
							},
							"title": "eqfqdasdwqd",
							"type":  "linepro",
							"search": map[string]interface{}{
								"\"logstore\"": "access-log",
								"\"query\"":    "* | SELECT date_format(__time__ - __time__ % 60, '%H:%i:%s') as time, count(1) as count GROUP BY time ORDER BY time",
								"\"start\"":    "-3600",
								"\"end\"":      "now",
							},
						},
					},
					"dashboard_name": name,
					"project_name":   "${var.project_name}",
					"display_name":   "asdsadsada",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":    "description",
						"charts.#":       "1",
						"dashboard_name": name,
						"project_name":   CHECKSET,
						"display_name":   "asdsadsada",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"charts": []map[string]interface{}{
						{
							"title": "adasdasdasd",
							"type":  "linepro",
							"search": map[string]interface{}{
								"\"logstore\"": "access-log",
								"\"query\"":    "* | SELECT date_format(__time__ - __time__ % 60, '%H:%i:%s') as time, count(1) as count GROUP BY time ORDER BY time",
								"\"start\"":    "-3600",
								"\"end\"":      "now",
							},
							"display": map[string]interface{}{
								"\"height\"": "5",
								"\"width\"":  "5",
								"\"xPos\"":   "0",
								"\"yPos\"":   "0",
								"\"xAxis\"":  "[\\\"province\\\"]",
								"\"yAxis\"":  "[\\\"pv\\\"]",
							},
							"action": map[string]interface{}{
								"\"name\"":   "open_logstore",
								"\"title\"":  "查看该省份日志",
								"\"events\"": "[{\\\"type\\\":\\\"click\\\",\\\"parameters\\\":{\\\"logstore\\\":\\\"access-log\\\",\\\"query\\\":\\\"ip_to_province(remote_addr): $${province}\\\",\\\"timeRange\\\":\\\"last_1h\\\"}}]",
							},
						},
					},
					"display_name": "asdqdawd",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"charts.#":     "1",
						"display_name": "asdqdawd",
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

var AlicloudSlsDashboardMap12338 = map[string]string{}

func AlicloudSlsDashboardBasicDependence12338(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "project_name" {
  default = "test-projects213"
}

variable "dashboard_name" {
  default = "testdashboard"
}

resource "alicloud_log_project" "defaultP71Nui" {
  description = "Description"
  name        = "test-projects213"
}


`, name)
}

// Test Sls Dashboard. <<< Resource test cases, automatically generated.
