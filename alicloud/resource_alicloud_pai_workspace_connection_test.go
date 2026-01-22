// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test PaiWorkspace Connection. >>> Resource test cases, automatically generated.
// Case 连接测试用例_1.0_副本1747385800080 10797
func TestAccAliCloudPaiWorkspaceConnection_basic10797(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_workspace_connection.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiWorkspaceConnectionMap10797)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiWorkspaceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiWorkspaceConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiWorkspaceConnectionBasicDependence10797)
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
					"connection_name": name,
					"description":     "镇元资源测试用例",
					"accessibility":   "PRIVATE",
					"connection_type": "MilvusConnection",
					"secrets": map[string]interface{}{
						"\"token\"": "****",
					},
					"configs": map[string]interface{}{
						"\"database\"": "c-b1c5222fba7cxxxx",
						"\"uri\"":      "https://c-b1c5222fba7cxxxx-public.milvus.aliyuncs.com",
					},
					"workspace_id": "${alicloud_pai_workspace_workspace.defaultMr2RQE.id}",
					"resource_meta": []map[string]interface{}{
						{
							"instance_name": "test_inst",
							"instance_id":   "c-b1c5222fba7cxxxx",
						},
					},
					"models": []map[string]interface{}{
						{
							"model":        "qwen-vl-max",
							"display_name": "通义千问VL-Max",
							"model_type":   "LLM",
						},
						{
							"model":        "qwen-qwq",
							"display_name": "通义千问qwq",
							"model_type":   "LLM",
						},
						{
							"model":        "qwen-plus",
							"display_name": "通义千问plus",
							"model_type":   "LLM",
							"tool_call":    "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name": name,
						"description":     "镇元资源测试用例",
						"accessibility":   "PRIVATE",
						"connection_type": "MilvusConnection",
						"workspace_id":    CHECKSET,
						"models.#":        "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "镇元资源测试用例_测试更新",
					"secrets": map[string]interface{}{
						"\"token\"": "*****",
					},
					"configs": map[string]interface{}{
						"\"database\"": "c-b1c5222fba7cxxxx",
						"\"uri\"":      "https://c-b1c5222fba7cxxxx-internal.milvus.aliyuncs.com",
					},
					"models": []map[string]interface{}{
						{
							"model":        "qwen-vl",
							"display_name": "通义千问VL",
							"model_type":   "Embedding",
							"tool_call":    "false",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "镇元资源测试用例_测试更新",
						"models.#":    "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "镇元资源测试用例_测试更新1",
					"models":      REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "镇元资源测试用例_测试更新1",
						"models.#":    "0",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"validate_type"},
			},
		},
	})
}

var AlicloudPaiWorkspaceConnectionMap10797 = map[string]string{}

func AlicloudPaiWorkspaceConnectionBasicDependence10797(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "defaultMr2RQE" {
  description    = "test_connection_1769080359"
  display_name   = "连接管理镇元资源测试用例"
  workspace_name = "connection_1769080359"
  env_types      = ["prod"]
}


`, name)
}

// Case 连接测试用例 10691
func TestAccAliCloudPaiWorkspaceConnection_basic10691(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_workspace_connection.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiWorkspaceConnectionMap10691)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiWorkspaceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiWorkspaceConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiWorkspaceConnectionBasicDependence10691)
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
					"connection_name": name,
					"description":     "镇元资源测试用例",
					"accessibility":   "PRIVATE",
					"connection_type": "MilvusConnection",
					"secrets": map[string]interface{}{
						"\"token\"": "****",
					},
					"configs": map[string]interface{}{
						"\"database\"": "c-b1c5222fba7cxxxx",
						"\"uri\"":      "https://c-b1c5222fba7cxxxx-public.milvus.aliyuncs.com",
					},
					"workspace_id": "${alicloud_pai_workspace_workspace.defaultMr2RQE.id}",
					"resource_meta": []map[string]interface{}{
						{
							"instance_name": "test_inst",
							"instance_id":   "c-b1c5222fba7cxxxx",
						},
					},
					"models": []map[string]interface{}{
						{
							"model":        "qwen-vl-max",
							"display_name": "通义千问VL-Max",
							"model_type":   "LLM",
						},
						{
							"model":        "qwen-qwq",
							"display_name": "通义千问qwq",
							"model_type":   "LLM",
						},
						{
							"model":        "qwen-plus",
							"display_name": "通义千问plus",
							"model_type":   "LLM",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name": name,
						"description":     "镇元资源测试用例",
						"accessibility":   "PRIVATE",
						"connection_type": "MilvusConnection",
						"workspace_id":    CHECKSET,
						"models.#":        "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "镇元资源测试用例_测试更新",
					"secrets": map[string]interface{}{
						"\"token\"": "*****",
					},
					"configs": map[string]interface{}{
						"\"database\"": "c-b1c5222fba7cxxxx",
						"\"uri\"":      "https://c-b1c5222fba7cxxxx-internal.milvus.aliyuncs.com",
					},
					"models": []map[string]interface{}{
						{
							"model":        "qwen-vl",
							"display_name": "通义千问VL",
							"model_type":   "Embedding",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "镇元资源测试用例_测试更新",
						"models.#":    "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "镇元资源测试用例_测试更新1",
					"models":      REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "镇元资源测试用例_测试更新1",
						"models.#":    "0",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"validate_type"},
			},
		},
	})
}

var AlicloudPaiWorkspaceConnectionMap10691 = map[string]string{}

func AlicloudPaiWorkspaceConnectionBasicDependence10691(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "defaultMr2RQE" {
  description    = "test_connection_1769080360"
  display_name   = "连接管理镇元资源测试用例"
  workspace_name = "connection_1769080360"
  env_types      = ["prod"]
}


`, name)
}

// Case 连接测试用例_1.0 10796
func TestAccAliCloudPaiWorkspaceConnection_basic10796(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_workspace_connection.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiWorkspaceConnectionMap10796)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiWorkspaceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiWorkspaceConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiWorkspaceConnectionBasicDependence10796)
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
					"connection_name": name,
					"description":     "镇元资源测试用例",
					"accessibility":   "PRIVATE",
					"connection_type": "MilvusConnection",
					"secrets": map[string]interface{}{
						"\"token\"": "****",
					},
					"configs": map[string]interface{}{
						"\"database\"": "c-b1c5222fba7cxxxx",
						"\"uri\"":      "https://c-b1c5222fba7cxxxx-public.milvus.aliyuncs.com",
					},
					"workspace_id": "${alicloud_pai_workspace_workspace.defaultMr2RQE.id}",
					"resource_meta": []map[string]interface{}{
						{
							"instance_name": "test_inst",
							"instance_id":   "c-b1c5222fba7cxxxx",
						},
					},
					"models": []map[string]interface{}{
						{
							"model":        "qwen-vl-max",
							"display_name": "通义千问VL-Max",
							"model_type":   "LLM",
						},
						{
							"model":        "qwen-qwq",
							"display_name": "通义千问qwq",
							"model_type":   "LLM",
						},
						{
							"model":        "qwen-plus",
							"display_name": "通义千问plus",
							"model_type":   "LLM",
							"tool_call":    "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name": name,
						"description":     "镇元资源测试用例",
						"accessibility":   "PRIVATE",
						"connection_type": "MilvusConnection",
						"workspace_id":    CHECKSET,
						"models.#":        "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "镇元资源测试用例_测试更新",
					"secrets": map[string]interface{}{
						"\"token\"": "*****",
					},
					"configs": map[string]interface{}{
						"\"database\"": "c-b1c5222fba7cxxxx",
						"\"uri\"":      "https://c-b1c5222fba7cxxxx-internal.milvus.aliyuncs.com",
					},
					"models": []map[string]interface{}{
						{
							"model":        "qwen-vl",
							"display_name": "通义千问VL",
							"model_type":   "Embedding",
							"tool_call":    "false",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "镇元资源测试用例_测试更新",
						"models.#":    "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "镇元资源测试用例_测试更新1",
					"models":      REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "镇元资源测试用例_测试更新1",
						"models.#":    "0",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"validate_type"},
			},
		},
	})
}

var AlicloudPaiWorkspaceConnectionMap10796 = map[string]string{}

func AlicloudPaiWorkspaceConnectionBasicDependence10796(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_pai_workspace_workspace" "defaultMr2RQE" {
  description    = "test_connection_1769080360"
  display_name   = "连接管理镇元资源测试用例"
  workspace_name = "connection_1769080360"
  env_types      = ["prod"]
}


`, name)
}

// Test PaiWorkspace Connection. <<< Resource test cases, automatically generated.
