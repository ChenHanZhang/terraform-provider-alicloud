// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test PaiAbTest Project. >>> Resource test cases, automatically generated.
// Case wkl_project_test 6504
func TestAccAliCloudPaiAbTestProject_basic6504(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_ab_test_project.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiAbTestProjectMap6504)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiAbTestServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiAbTestProject")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiAbTestProjectBasicDependence6504)
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
					"description":  "this is a test project",
					"project_name": name,
					"workspace_id": "45699",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  "this is a test project",
						"project_name": name,
						"workspace_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":  "this is update function",
					"project_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  "this is update function",
						"project_name": name + "_update",
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

var AlicloudPaiAbTestProjectMap6504 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudPaiAbTestProjectBasicDependence6504(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case wkl_project_test2 6625
func TestAccAliCloudPaiAbTestProject_basic6625(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_ab_test_project.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiAbTestProjectMap6625)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiAbTestServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiAbTestProject")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiAbTestProjectBasicDependence6625)
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
					"description":  "this is a create function",
					"project_name": name,
					"workspace_id": "45699",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  "this is a create function",
						"project_name": name,
						"workspace_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":  "this is a update function",
					"project_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  "this is a update function",
						"project_name": name + "_update",
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

var AlicloudPaiAbTestProjectMap6625 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudPaiAbTestProjectBasicDependence6625(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test PaiAbTest Project. <<< Resource test cases, automatically generated.
