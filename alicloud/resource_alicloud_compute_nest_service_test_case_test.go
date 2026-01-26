// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ComputeNest ServiceTestCase. >>> Resource test cases, automatically generated.
// Case 服务测试用例 6544
func TestAccAliCloudComputeNestServiceTestCase_basic6544(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_compute_nest_service_test_case.default"
	ra := resourceAttrInit(resourceId, AlicloudComputeNestServiceTestCaseMap6544)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ComputeNestServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeComputeNestServiceTestCase")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccomputenest%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudComputeNestServiceTestCaseBasicDependence6544)
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
					"test_config":     "{ parameters: { SystemDiskSize: 100, PayType: 'PostPaid',  DataDiskSize: 40, InstanceType: '$[iact3-auto]',  AllocatePublicIp: 'true',  DataDiskCategory: 'cloud_efficiency',InstancePassword: '$[iact3-auto]',SystemDiskCategory: 'cloud_essd'} }",
					"template_name":   "模版2",
					"service_version": "1",
					"test_case_name":  "casetest",
					"service_id":      "service-b7f967cacfa04d699a8f",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"test_config":     "{ parameters: { SystemDiskSize: 100, PayType: 'PostPaid',  DataDiskSize: 40, InstanceType: '$[iact3-auto]',  AllocatePublicIp: 'true',  DataDiskCategory: 'cloud_efficiency',InstancePassword: '$[iact3-auto]',SystemDiskCategory: 'cloud_essd'} }",
						"template_name":   "模版2",
						"service_version": CHECKSET,
						"test_case_name":  "casetest",
						"service_id":      "service-b7f967cacfa04d699a8f",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"test_config":    "{ parameters: { SystemDiskSize: 150, PayType: 'PostPaid',  DataDiskSize: 40, InstanceType: '$[iact3-auto]',  AllocatePublicIp: 'true',  DataDiskCategory: 'cloud_efficiency',InstancePassword: '$[iact3-auto]',SystemDiskCategory: 'cloud_essd'} }",
					"test_case_name": "case22",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"test_config":    "{ parameters: { SystemDiskSize: 150, PayType: 'PostPaid',  DataDiskSize: 40, InstanceType: '$[iact3-auto]',  AllocatePublicIp: 'true',  DataDiskCategory: 'cloud_efficiency',InstancePassword: '$[iact3-auto]',SystemDiskCategory: 'cloud_essd'} }",
						"test_case_name": "case22",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id", "service_version"},
			},
		},
	})
}

var AlicloudComputeNestServiceTestCaseMap6544 = map[string]string{}

func AlicloudComputeNestServiceTestCaseBasicDependence6544(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test ComputeNest ServiceTestCase. <<< Resource test cases, automatically generated.
