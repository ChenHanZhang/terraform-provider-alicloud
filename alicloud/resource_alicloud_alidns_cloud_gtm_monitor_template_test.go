// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Alidns CloudGtmMonitorTemplate. >>> Resource test cases, automatically generated.
// Case 模版资源用例-1 12646
func TestAccAliCloudAlidnsCloudGtmMonitorTemplate_basic12646(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alidns_cloud_gtm_monitor_template.default"
	ra := resourceAttrInit(resourceId, AlicloudAlidnsCloudGtmMonitorTemplateMap12646)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AlidnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAlidnsCloudGtmMonitorTemplate")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalidns%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAlidnsCloudGtmMonitorTemplateBasicDependence12646)
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
					"ip_version":       "IPv4",
					"timeout":          "3000",
					"evaluation_count": "2",
					"protocol":         "ping",
					"failure_rate":     "50",
					"extend_info":      "{\\\"packetLossRate\\\":10,\\\"packetNum\\\":5}",
					"name":             name,
					"interval":         "60",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ip_version":       "IPv4",
						"timeout":          CHECKSET,
						"evaluation_count": CHECKSET,
						"protocol":         "ping",
						"failure_rate":     "50",
						"extend_info":      CHECKSET,
						"name":             name,
						"interval":         CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"timeout":          "5000",
					"evaluation_count": "{\\\"packetLossRate\\\":10,\\\"packetNum\\\":20}",
					"failure_rate":     "100",
					"extend_info":      "60",
					"remark":           "test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"timeout":          CHECKSET,
						"evaluation_count": CHECKSET,
						"failure_rate":     "100",
						"extend_info":      CHECKSET,
						"remark":           "test",
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

var AlicloudAlidnsCloudGtmMonitorTemplateMap12646 = map[string]string{}

func AlicloudAlidnsCloudGtmMonitorTemplateBasicDependence12646(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Alidns CloudGtmMonitorTemplate. <<< Resource test cases, automatically generated.
