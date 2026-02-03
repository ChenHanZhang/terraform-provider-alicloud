// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test PaiStudio Algorithm. >>> Resource test cases, automatically generated.
// Case Algorithm_正式 8766
func TestAccAliCloudPaiStudioAlgorithm_basic8766(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_studio_algorithm.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiStudioAlgorithmMap8766)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiStudioServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiStudioAlgorithm")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiStudioAlgorithmBasicDependence8766)
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
					"workspace_id":          "${alicloud_pai_workspace_workspace.workspace.id}",
					"display_name":          "test",
					"algorithm_description": "算法测试用例使用",
					"algorithm_name":        name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"workspace_id":          CHECKSET,
						"display_name":          "test",
						"algorithm_description": "算法测试用例使用",
						"algorithm_name":        name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"display_name":          "test-update",
					"algorithm_description": "资源测试更新",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"display_name":          "test-update",
						"algorithm_description": "资源测试更新",
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

var AlicloudPaiStudioAlgorithmMap8766 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudPaiStudioAlgorithmBasicDependence8766(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  workspace_name = "test_work_space_1770122479"
  env_types      = ["prod"]
  description    = "Algorithm资源测试用例使用"
  display_name   = "algorithm_test_case"
}


`, name)
}

// Case Algorithm 8644
func TestAccAliCloudPaiStudioAlgorithm_basic8644(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_studio_algorithm.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiStudioAlgorithmMap8644)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiStudioServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiStudioAlgorithm")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiStudioAlgorithmBasicDependence8644)
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
					"workspace_id":          "${alicloud_pai_workspace_workspace.workspace.id}",
					"display_name":          "test",
					"algorithm_description": "算法测试用例使用",
					"algorithm_name":        name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"workspace_id":          CHECKSET,
						"display_name":          "test",
						"algorithm_description": "算法测试用例使用",
						"algorithm_name":        name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"display_name":          "test-update",
					"algorithm_description": "资源测试更新",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"display_name":          "test-update",
						"algorithm_description": "资源测试更新",
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

var AlicloudPaiStudioAlgorithmMap8644 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudPaiStudioAlgorithmBasicDependence8644(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  workspace_name = "test_work_space_1770122481"
  env_types      = ["prod"]
  description    = "Algorithm资源测试用例使用"
  display_name   = "algorithm_test_case"
}


`, name)
}

// Test PaiStudio Algorithm. <<< Resource test cases, automatically generated.
