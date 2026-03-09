// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test PaiAbTest Crowd. >>> Resource test cases, automatically generated.
// Case wkl_corwd_test 6540
func TestAccAliCloudPaiAbTestCrowd_basic6540(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_pai_ab_test_crowd.default"
	ra := resourceAttrInit(resourceId, AlicloudPaiAbTestCrowdMap6540)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PaiAbTestServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePaiAbTestCrowd")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccpaiabtest%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPaiAbTestCrowdBasicDependence6540)
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
					"crowd_name":   name,
					"description":  "this is a test crowd",
					"workspace_id": "45699",
					"users":        "uid1,uid2",
					"label":        "test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"crowd_name":   name,
						"description":  "this is a test crowd",
						"workspace_id": CHECKSET,
						"users":        "uid1,uid2",
						"label":        "test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"crowd_name":  name + "_update",
					"description": "this is a update function",
					"users":       "uid1,uid2,uid3",
					"label":       "test2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"crowd_name":  name + "_update",
						"description": "this is a update function",
						"users":       "uid1,uid2,uid3",
						"label":       "test2",
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

var AlicloudPaiAbTestCrowdMap6540 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudPaiAbTestCrowdBasicDependence6540(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test PaiAbTest Crowd. <<< Resource test cases, automatically generated.
