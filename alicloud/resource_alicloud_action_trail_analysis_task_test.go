// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ActionTrail AnalysisTask. >>> Resource test cases, automatically generated.
// Case AnalysisTask模型测试用例 12519
func TestAccAliCloudActionTrailAnalysisTask_basic12519(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_action_trail_analysis_task.default"
	ra := resourceAttrInit(resourceId, AlicloudActionTrailAnalysisTaskMap12519)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ActionTrailServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeActionTrailAnalysisTask")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccactiontrail%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudActionTrailAnalysisTaskBasicDependence12519)
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
					"end_time":         "2026-04-10CST1717:0404:040428800",
					"query_conditions": "{\\\"access_key_ids\\\": [\\\"LTAI4GDuVyEUfNcnWCsXgbbR\\\"]}",
					"start_time":       "2026-04-10CST1717:0404:040428800",
					"query_type":       "AccessKey",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"end_time":         CHECKSET,
						"query_conditions": CHECKSET,
						"start_time":       CHECKSET,
						"query_type":       "AccessKey",
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

var AlicloudActionTrailAnalysisTaskMap12519 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudActionTrailAnalysisTaskBasicDependence12519(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test ActionTrail AnalysisTask. <<< Resource test cases, automatically generated.
