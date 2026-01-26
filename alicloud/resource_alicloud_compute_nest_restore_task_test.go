// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ComputeNest RestoreTask. >>> Resource test cases, automatically generated.
// Case 测试备份恢复 10475
func TestAccAliCloudComputeNestRestoreTask_basic10475(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_compute_nest_restore_task.default"
	ra := resourceAttrInit(resourceId, AlicloudComputeNestRestoreTaskMap10475)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ComputeNestServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeComputeNestRestoreTask")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccomputenest%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudComputeNestRestoreTaskBasicDependence10475)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"service_instance_id": "${alicloud_compute_nest_service_instance.defaultslP8EP.id}",
					"backup_id":           "${alicloud_compute_nest_backup.defaulthJ0HNs.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"service_instance_id": CHECKSET,
						"backup_id":           CHECKSET,
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

var AlicloudComputeNestRestoreTaskMap10475 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudComputeNestRestoreTaskBasicDependence10475(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_compute_nest_service" "defaultpV4cW6" {
  deploy_type = "ros"
  service_info {
    locale            = "zh-CN"
    short_description = "自动化测试备份service创建服务"
    image             = "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/05498c54-80ea-445f-9a44-1e2ba25b7462.png"
    name              = "自动化测试备份service创建服务-7"
  }
  policy_names = "AliyunComputeNestPolicyForFullAccess"
  deploy_metadata {
  }
  service_type       = "private"
  approval_type      = "Manual"
  version_name       = "自动化测试service创建服务"
  operation_metadata = "{\"PrometheusConfigMap\":{\"模板1\":{\"EnablePrometheus\":false}},\"ModifyParametersConfig\":[], \"SupportBackup\":true,\"StatusOperationConfigs\":[{\"TemplateName\":\"模板1\",\"SupportOperations\":[\"Start\",\"Stop\"]}]}"
}

resource "alicloud_compute_nest_service_instance" "defaultslP8EP" {
  parameters = "{\"RegionId\":\"cn-hangzhou\",\"PayType\":\"PostPaid\",\"ZoneId\":\"cn-hangzhou-i\",\"VpcCidrBlock\":\"192.168.0.0/16\",\"VSwitchCidrBlock\":\"192.168.11.0/24\",\"EcsInstanceType\":\"ecs.g6.large\",\"SystemDiskCategory\":\"cloud_essd\",\"SystemDiskSize\":40,\"InstancePassword\":\"liuzheng121@\",\"InstanceCount\":1,\"UserEnablePrometheus\":false}"
  service {
    service_id = alicloud_compute_nest_service.defaultpV4cW6.id
  }
  template_name = "模板1"
}

resource "alicloud_compute_nest_backup" "defaulthJ0HNs" {
  service_instance_id = alicloud_compute_nest_service_instance.defaultslP8EP.id
  description         = "fdsfdsfsd_909"
}


`, name)
}

// Test ComputeNest RestoreTask. <<< Resource test cases, automatically generated.
