// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Schedulerx WorkFlow. >>> Resource test cases, automatically generated.
// Case 预发环境_20250318_乌兰察布region 10567
func TestAccAliCloudSchedulerxWorkFlow_basic10567(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_schedulerx_work_flow.default"
	ra := resourceAttrInit(resourceId, AlicloudSchedulerxWorkFlowMap10567)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SchedulerxServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSchedulerxWorkFlow")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccschedulerx%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSchedulerxWorkFlowBasicDependence10567)
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
					"timezone":         "GTM+7",
					"description":      "workflow资源用例自动生成的任务",
					"workflow_name":    name,
					"max_concurrency":  "10",
					"time_expression":  "0 0 18 1 */1 ?",
					"namespace":        "${alicloud_schedulerx_namespace.CreateNameSpace.id}",
					"group_id":         "${alicloud_schedulerx_app_group.CreateAppGroup.group_id}",
					"time_type":        "1",
					"status":           "Disable",
					"namespace_source": "schedulerx",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"timezone":         "GTM+7",
						"description":      "workflow资源用例自动生成的任务",
						"workflow_name":    name,
						"max_concurrency":  "10",
						"time_expression":  "0 0 18 1 */1 ?",
						"namespace":        CHECKSET,
						"group_id":         CHECKSET,
						"time_type":        "1",
						"status":           "Disable",
						"namespace_source": "schedulerx",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":     "workflow资源用例自动生成的工作流",
					"workflow_name":   name + "_update",
					"time_expression": "/",
					"time_type":       "100",
					"status":          "Enable",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":     "workflow资源用例自动生成的工作流",
						"workflow_name":   name + "_update",
						"time_expression": "/",
						"time_type":       "100",
						"status":          "Enable",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Disable",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "Disable",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"namespace_source", "timezone"},
			},
		},
	})
}

var AlicloudSchedulerxWorkFlowMap10567 = map[string]string{
	"work_flow_id": CHECKSET,
}

func AlicloudSchedulerxWorkFlowBasicDependence10567(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_schedulerx_namespace" "CreateNameSpace" {
  namespace_name = "test-namespace-pop-autotest"
  description    = "由workflow 资源测试用例前置步骤创建"
}

resource "alicloud_schedulerx_app_group" "CreateAppGroup" {
  description    = "由workflow 资源测试用例前置步骤创建"
  enable_log     = false
  namespace      = alicloud_schedulerx_namespace.CreateNameSpace.id
  group_id       = "test-appgroup-pop-autotest"
  app_name       = "test-appgroup-pop-autotest"
  app_version    = "2"
  namespace_name = alicloud_schedulerx_namespace.CreateNameSpace.namespace_name
  app_type       = "2"
  max_jobs       = "100"
}


`, name)
}

// Test Schedulerx WorkFlow. <<< Resource test cases, automatically generated.
