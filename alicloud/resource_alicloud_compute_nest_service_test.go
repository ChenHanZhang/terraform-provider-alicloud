// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ComputeNest Service. >>> Resource test cases, automatically generated.
// Case 自动化测试service创建服务 3293
func TestAccAliCloudComputeNestService_basic3293(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_compute_nest_service.default"
	ra := resourceAttrInit(resourceId, AlicloudComputeNestServiceMap3293)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ComputeNestServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeComputeNestService")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccomputenest%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudComputeNestServiceBasicDependence3293)
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
					"deploy_type":  "operation",
					"duration":     "2592000",
					"policy_names": "AliyunComputeNestPolicyForReadOnly",
					"service_info": []map[string]interface{}{
						{
							"locale":            "zh-CN",
							"short_description": "自动化测试service创建服务",
							"image":             "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/05498c54-80ea-445f-9a44-1e2ba25b7462.png",
							"name":              "自动化测试service创建服务",
						},
					},
					"deploy_metadata": []map[string]interface{}{
						{
							"supplier_deploy_metadata": []map[string]interface{}{
								{
									"deploy_timeout": "7200",
								},
							},
						},
					},
					"service_type":       "operation",
					"approval_type":      "Manual",
					"version_name":       "自动化测试service创建服务",
					"operation_metadata": "{\\\"PrometheusConfigMap\\\":{}}",
					"share_type":         "Public",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":        "operation",
						"duration":           "2592000",
						"policy_names":       "AliyunComputeNestPolicyForReadOnly",
						"service_info.#":     "1",
						"service_type":       "operation",
						"approval_type":      "Manual",
						"version_name":       "自动化测试service创建服务",
						"operation_metadata": CHECKSET,
						"share_type":         "Public",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudComputeNestServiceMap3293 = map[string]string{
	"status": CHECKSET,
}

func AlicloudComputeNestServiceBasicDependence3293(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 私有化部署加更新 7596
func TestAccAliCloudComputeNestService_basic7596(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_compute_nest_service.default"
	ra := resourceAttrInit(resourceId, AlicloudComputeNestServiceMap7596)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ComputeNestServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeComputeNestService")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccomputenest%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudComputeNestServiceBasicDependence7596)
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
					"deploy_type":  "ros",
					"duration":     "2592000",
					"policy_names": "AliyunComputeNestPolicyForFullAccess",
					"service_info": []map[string]interface{}{
						{
							"locale":            "zh-CN",
							"short_description": "自动化测试service创建服务",
							"image":             "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/05498c54-80ea-445f-9a44-1e2ba25b7462.png",
							"name":              "自动化测试service创建服务-632",
						},
						{
							"locale":            "en-US",
							"short_description": "test",
							"image":             "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/05498c54-80ea-445f-9a44-1e2ba25b7462.png",
							"name":              "testName",
						},
					},
					"service_type":        "private",
					"approval_type":       "Manual",
					"version_name":        "自动化测试service创建服务",
					"operation_metadata":  "{\\\"PrometheusConfigMap\\\":{}}",
					"share_type":          "Public",
					"is_support_operated": "true",
					"deploy_metadata": []map[string]interface{}{
						{},
					},
					"trial_duration":    "7",
					"upgrade_metadata":  "{\\\"Description\\\":\\\"测试升级\\\",\\\"SupportRollback\\\":true,\\\"SupportUpgradeFromVersions\\\":[],\\\"UpgradeComponents\\\":[\\\"Configuration\\\",\\\"Resource\\\"]}",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":         "ros",
						"duration":            "2592000",
						"policy_names":        "AliyunComputeNestPolicyForFullAccess",
						"service_info.#":      "2",
						"service_type":        "private",
						"approval_type":       "Manual",
						"version_name":        "自动化测试service创建服务",
						"operation_metadata":  CHECKSET,
						"share_type":          "Public",
						"is_support_operated": "true",
						"trial_duration":      "7",
						"upgrade_metadata":    CHECKSET,
						"resource_group_id":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"duration":     "7200",
					"policy_names": "AliyunComputeNestPolicyForReadOnly",
					"service_info": []map[string]interface{}{
						{
							"locale":            "zh-CN",
							"short_description": "自动测试运行-605",
							"image":             "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/99005496-6390-45d8-a74e-323a2c315545.jpeg",
							"name":              "自动测试运行-65",
						},
						{
							"locale":            "en-US",
							"short_description": "test2",
							"image":             "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/99005496-6390-45d8-a74e-323a2c315545.jpeg",
							"name":              "test2-xx",
						},
					},
					"version_name":       "自动测试运行-4",
					"operation_metadata": "{\\\"ModifyParametersConfig\\\":[{\\\"TemplateName\\\":\\\"模板1\\\",\\\"Operation\\\":[{\\\"Name\\\":\\\"升级\\\",\\\"Description\\\":\\\"升级\\\",\\\"Type\\\":\\\"Upgrade\\\",\\\"SupportPredefinedParameters\\\":true,\\\"EnableLogging\\\":false}]}],\\\"PrometheusConfigMap\\\":{\\\"模板1\\\":{\\\"EnablePrometheus\\\":false}}}",
					"deploy_metadata": []map[string]interface{}{
						{},
					},
					"trial_duration":   "8",
					"upgrade_metadata": "{\\\"Description\\\":\\\"测试升级2\\\",\\\"SupportRollback\\\":true,\\\"SupportUpgradeFromVersions\\\":[],\\\"UpgradeComponents\\\":[\\\"Configuration\\\"]}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"duration":           "7200",
						"policy_names":       "AliyunComputeNestPolicyForReadOnly",
						"service_info.#":     "2",
						"version_name":       CHECKSET,
						"operation_metadata": CHECKSET,
						"trial_duration":     "8",
						"upgrade_metadata":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"is_support_operated": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"is_support_operated": "false",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudComputeNestServiceMap7596 = map[string]string{
	"status": CHECKSET,
}

func AlicloudComputeNestServiceBasicDependence7596(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Test ComputeNest Service. <<< Resource test cases, automatically generated.
