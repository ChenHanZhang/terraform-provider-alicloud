// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cr DiagnosisTask. >>> Resource test cases, automatically generated.
// Case 推荐转换-资源-725: DiagnosisTask 8515
func TestAccAliCloudCrDiagnosisTask_basic8515(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_diagnosis_task.default"
	ra := resourceAttrInit(resourceId, AlicloudCrDiagnosisTaskMap8515)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrDiagnosisTask")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrDiagnosisTaskBasicDependence8515)
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
					"instance_id": "cri-nbuedifs8yxf03dq",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id": "cri-nbuedifs8yxf03dq",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"targets"},
			},
		},
	})
}

var AlicloudCrDiagnosisTaskMap8515 = map[string]string{
	"status":            CHECKSET,
	"diagnosis_time":    CHECKSET,
	"diagnosis_task_id": CHECKSET,
}

func AlicloudCrDiagnosisTaskBasicDependence8515(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "diagnosis_task_id" {
  default = "bd-d26c373b47b74e398d7abaed27"
}

variable "instance_id" {
  default = "cri-nbuedifs8yxf03dq"
}

variable "task_type" {
  default = "BUILD_DIAGNOSIS"
}

variable "related_id" {
  default = "4623D623-DEF7-1344-B99E-2CAD63F0160F"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_cr_ee_instance" "Instance" {
}


`, name)
}

// Test Cr DiagnosisTask. <<< Resource test cases, automatically generated.
