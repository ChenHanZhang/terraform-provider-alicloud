// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test DataWorks Function. >>> Resource test cases, automatically generated.
// Case NormalCase 7820
func TestAccAliCloudDataWorksFunction_basic7820(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_function.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksFunctionMap7820)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksFunction")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksFunctionBasicDependence7820)
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
					"function_name": name,
					"project_id":    "${alicloud_data_works_project.CreateProject.id}",
					"spec":          "{\\n    \\\"version\\\": \\\"1.1.0\\\",\\n    \\\"kind\\\": \\\"Function\\\",\\n    \\\"spec\\\": {\\n        \\\"functions\\\": [\\n            {\\n                \\\"name\\\": \\\"OpenAPIFunc\\\",\\n                \\\"id\\\": \\\"5806679648885952157\\\",\\n                \\\"script\\\": {\\n                    \\\"content\\\": \\\"{  \\\\\\\"uuid\\\\\\\": \\\\\\\"5806679648885952157\\\\\\\",  \\\\\\\"name\\\\\\\": \\\\\\\"OpenAPIFunc\\\\\\\"}\\\",\\n                    \\\"path\\\": \\\"莫泣/OpenAPI测试/函数测试/OpenAPIFunc\\\",\\n                    \\\"runtime\\\": {\\n                        \\\"command\\\": \\\"ODPS_FUNCTION\\\"\\n                    }\\n                }\\n            }\\n        ]\\n    }\\n}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"function_name": name,
						"project_id":    CHECKSET,
						"spec":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec": "{\\n    \\\"version\\\": \\\"1.1.0\\\",\\n    \\\"kind\\\": \\\"Function\\\",\\n    \\\"spec\\\": {\\n        \\\"functions\\\": [\\n            {\\n                \\\"name\\\": \\\"OpenAPIFunc\\\",\\n                \\\"id\\\": \\\"5806679648885952157\\\",\\n                \\\"script\\\": {\\n                    \\\"content\\\": \\\"{  \\\\\\\"uuid\\\\\\\": \\\\\\\"5806679648885952157\\\\\\\",  \\\\\\\"name\\\\\\\": \\\\\\\"OpenAPIFunc\\\\\\\"}\\\",\\n                    \\\"path\\\": \\\"莫泣/OpenAPI测试/函数测试/OpenAPIFunc\\\",\\n                    \\\"runtime\\\": {\\n                        \\\"command\\\": \\\"ODPS_FUNCTION\\\"\\n                    }\\n                },\\n                \\\"type\\\":\\\"math\\\"\\n            }\\n        ]\\n    }\\n}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"function_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"function_name": name + "_update",
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

var AlicloudDataWorksFunctionMap7820 = map[string]string{
	"function_id": CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudDataWorksFunctionBasicDependence7820(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_data_works_project" "CreateProject" {
  description      = "新版数据开发测试terraform接入的工作空间"
  project_name     = "tf_test_datastudio5_4"
  pai_task_enabled = false
  display_name     = "tf_test_datastudio5"
}


`, name)
}

// Test DataWorks Function. <<< Resource test cases, automatically generated.
