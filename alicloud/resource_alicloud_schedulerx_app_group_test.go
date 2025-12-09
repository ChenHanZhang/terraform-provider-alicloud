package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Schedulerx AppGroup. >>> Resource test cases, automatically generated.
// Case 预发环境_20251105_乌兰察布(代码只部署到乌兰察布，用这个用例测试) 11495
func TestAccAliCloudSchedulerxAppGroup_basic11495(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_schedulerx_app_group.default"
	ra := resourceAttrInit(resourceId, AlicloudSchedulerxAppGroupMap11495)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SchedulerxServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSchedulerxAppGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccschedulerx%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSchedulerxAppGroupBasicDependence11495)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"namespace":                "${alicloud_schedulerx_namespace.CreateNameSpace.id}",
					"group_id":                 "test-appgroup-pop-autotest",
					"description":              "appgroup 资源用例自动生成",
					"monitor_contacts_json":    "[{\\\"name\\\":\\\"David\\\"},{\\\"name\\\":\\\"Grace\\\"}]",
					"enable_log":               "false",
					"app_name":                 "test-appgroup-pop-autotest",
					"app_version":              "1",
					"namespace_name":           "default",
					"monitor_config_json":      "{\\\"sendChannel\\\":\\\"sms,ding\\\",\\\"alarmType\\\":\\\"Contacts\\\",\\\"webhookIsAtAll\\\":false,\\\"useNotificationPolicy\\\":true}",
					"app_type":                 "1",
					"max_jobs":                 "100",
					"namespace_source":         "schedulerx",
					"schedule_busy_workers":    "false",
					"notification_policy_name": "test-notification",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"namespace":                CHECKSET,
						"group_id":                 "test-appgroup-pop-autotest",
						"description":              "appgroup 资源用例自动生成",
						"monitor_contacts_json":    CHECKSET,
						"enable_log":               "false",
						"app_name":                 "test-appgroup-pop-autotest",
						"app_version":              CHECKSET,
						"namespace_name":           "default",
						"monitor_config_json":      CHECKSET,
						"app_type":                 "1",
						"max_jobs":                 "100",
						"namespace_source":         "schedulerx",
						"schedule_busy_workers":    "false",
						"notification_policy_name": "test-notification",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":              "appgroup 资源用例自动生成_update",
					"monitor_contacts_json":    "[{\\\"name\\\":\\\"Brian\\\"},{\\\"name\\\":\\\"Laura\\\"}]",
					"enable_log":               "true",
					"app_version":              "2",
					"monitor_config_json":      "{\\\"sendChannel\\\":\\\"ding\\\",\\\"alarmType\\\":\\\"Contacts\\\",\\\"webhookIsAtAll\\\":false,\\\"useNotificationPolicy\\\":true}",
					"notification_policy_name": "test-all-everyday",
					"max_concurrency":          "500",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":              "appgroup 资源用例自动生成_update",
						"monitor_contacts_json":    CHECKSET,
						"enable_log":               "true",
						"app_version":              CHECKSET,
						"monitor_config_json":      CHECKSET,
						"notification_policy_name": "test-all-everyday",
						"max_concurrency":          "500",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_type", "enable_log", "max_concurrency", "namespace_name", "namespace_source", "schedule_busy_workers"},
			},
		},
	})
}

var AlicloudSchedulerxAppGroupMap11495 = map[string]string{}

func AlicloudSchedulerxAppGroupBasicDependence11495(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_schedulerx_namespace" "CreateNameSpace" {
  namespace_name = "test-namespace-pop-autotest"
  description    = "由appgroup 资源测试用例前置步骤创建"
}


`, name)
}

// Test Schedulerx AppGroup. <<< Resource test cases, automatically generated.
