// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test DataWorks WorkflowDefinition. >>> Resource test cases, automatically generated.
// Case NormalCase 7809
func TestAccAliCloudDataWorksWorkflowDefinition_basic7809(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_workflow_definition.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksWorkflowDefinitionMap7809)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksWorkflowDefinition")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksWorkflowDefinitionBasicDependence7809)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shanghai"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"workflow_definition_name": name,
					"project_id":               "${alicloud_data_works_project.CreateProject.id}",
					"spec":                     "{\\n    \\\"metadata\\\": {\\n        \\\"uuid\\\": \\\"4634978808809548226\\\"\\n    },\\n    \\\"kind\\\": \\\"CycleWorkflow\\\",\\n    \\\"version\\\": \\\"1.1.0\\\",\\n    \\\"spec\\\": {\\n        \\\"name\\\": \\\"OpenAPI测试工作流Demo\\\",\\n        \\\"id\\\": \\\"4634978808809548226\\\",\\n        \\\"type\\\": \\\"CycleWorkflow\\\",\\n        \\\"workflows\\\": [\\n            {\\n                \\\"script\\\": {\\n                    \\\"path\\\": \\\"莫泣/OpenAPI测试/工作流测试/OpenAPI测试工作流Demo\\\",\\n                    \\\"runtime\\\": {\\n                        \\\"command\\\": \\\"WORKFLOW\\\"\\n                    },\\n                    \\\"id\\\": \\\"6980027813686443486\\\"\\n                },\\n                \\\"id\\\": \\\"4634978808809548226\\\",\\n                \\\"trigger\\\": {\\n                    \\\"type\\\": \\\"Scheduler\\\",\\n                    \\\"id\\\": \\\"6525678244703541090\\\",\\n                    \\\"cron\\\": \\\"00 02 00 * * ?\\\",\\n                    \\\"startTime\\\": \\\"1970-01-01 00:00:00\\\",\\n                    \\\"endTime\\\": \\\"9999-01-01 00:00:00\\\",\\n                    \\\"timezone\\\": \\\"Asia/Shanghai\\\",\\n                    \\\"delaySeconds\\\": 0\\n                },\\n                \\\"strategy\\\": {\\n                    \\\"timeout\\\": 0,\\n                    \\\"instanceMode\\\": \\\"T+1\\\",\\n                    \\\"rerunMode\\\": \\\"Allowed\\\",\\n                    \\\"rerunTimes\\\": 3,\\n                    \\\"rerunInterval\\\": 180000,\\n                    \\\"failureStrategy\\\": \\\"Break\\\"\\n                },\\n                \\\"name\\\": \\\"OpenAPI测试工作流Demo\\\",\\n                \\\"inputs\\\": {},\\n                \\\"outputs\\\": {\\n                    \\\"nodeOutputs\\\": [\\n                        {\\n                            \\\"data\\\": \\\"4634978808809548226\\\",\\n                            \\\"artifactType\\\": \\\"NodeOutput\\\",\\n                            \\\"refTableName\\\": \\\"OpenAPI测试工作流Demo\\\",\\n                            \\\"isDefault\\\": true\\n                        }\\n                    ]\\n                },\\n                \\\"nodes\\\": [],\\n                \\\"dependencies\\\": []\\n            }\\n        ]\\n    }\\n}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"workflow_definition_name": name,
						"project_id":               CHECKSET,
						"spec":                     CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec": "{\\n    \\\"metadata\\\": {\\n        \\\"uuid\\\": \\\"4634978808809548226\\\"\\n    },\\n    \\\"kind\\\": \\\"CycleWorkflow\\\",\\n    \\\"version\\\": \\\"1.1.0\\\",\\n    \\\"spec\\\": {\\n        \\\"name\\\": \\\"OpenAPI测试工作流Demo\\\",\\n        \\\"id\\\": \\\"4634978808809548226\\\",\\n        \\\"type\\\": \\\"CycleWorkflow\\\",\\n        \\\"workflows\\\": [\\n            {\\n                \\\"script\\\": {\\n                    \\\"path\\\": \\\"莫泣/OpenAPI测试/工作流测试/OpenAPI测试工作流Demo\\\",\\n                    \\\"runtime\\\": {\\n                        \\\"command\\\": \\\"WORKFLOW\\\"\\n                    },\\n                    \\\"id\\\": \\\"6980027813686443486\\\"\\n                },\\n                \\\"id\\\": \\\"4634978808809548226\\\",\\n                \\\"trigger\\\": {\\n                    \\\"type\\\": \\\"Scheduler\\\",\\n                    \\\"id\\\": \\\"6525678244703541090\\\",\\n                    \\\"cron\\\": \\\"00 00 05 * * ?\\\",\\n                    \\\"startTime\\\": \\\"1970-01-01 00:00:00\\\",\\n                    \\\"endTime\\\": \\\"9999-01-01 00:00:00\\\",\\n                    \\\"timezone\\\": \\\"Asia/Shanghai\\\",\\n                    \\\"delaySeconds\\\": 0\\n                },\\n                \\\"strategy\\\": {\\n                    \\\"timeout\\\": 0,\\n                    \\\"instanceMode\\\": \\\"T+1\\\",\\n                    \\\"rerunMode\\\": \\\"Allowed\\\",\\n                    \\\"rerunTimes\\\": 3,\\n                    \\\"rerunInterval\\\": 180000,\\n                    \\\"failureStrategy\\\": \\\"Break\\\"\\n                },\\n                \\\"name\\\": \\\"OpenAPI测试工作流Demo\\\",\\n                \\\"inputs\\\": {},\\n                \\\"outputs\\\": {\\n                    \\\"nodeOutputs\\\": [\\n                        {\\n                            \\\"data\\\": \\\"4634978808809548226\\\",\\n                            \\\"artifactType\\\": \\\"NodeOutput\\\",\\n                            \\\"refTableName\\\": \\\"OpenAPI测试工作流Demo\\\",\\n                            \\\"isDefault\\\": true\\n                        }\\n                    ]\\n                },\\n                \\\"nodes\\\": [],\\n                \\\"dependencies\\\": []\\n            }\\n        ]\\n    }\\n}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"workflow_definition_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"workflow_definition_name": name + "_update",
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"spec"},
			},
		},
	})
}

var AlicloudDataWorksWorkflowDefinitionMap7809 = map[string]string{
	"create_time":            CHECKSET,
	"workflow_definition_id": CHECKSET,
}

func AlicloudDataWorksWorkflowDefinitionBasicDependence7809(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_data_works_project" "CreateProject" {
  description      = "新版数据开发测试terraform接入的工作空间"
  project_name     = "tf_test_datastudio5_2"
  pai_task_enabled = false
  display_name     = "tf_test_datastudio5"
}


`, name)
}

// Test DataWorks WorkflowDefinition. <<< Resource test cases, automatically generated.
