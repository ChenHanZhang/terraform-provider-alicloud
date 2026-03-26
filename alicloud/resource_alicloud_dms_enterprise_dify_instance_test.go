// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test DmsEnterprise DifyInstance. >>> Resource test cases, automatically generated.
// Case difyinstance测试 12559
func TestAccAliCloudDmsEnterpriseDifyInstance_basic12559(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_dms_enterprise_dify_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudDmsEnterpriseDifyInstanceMap12559)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DMSEnterpriseServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDmsEnterpriseDifyInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdmsenterprise%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDmsEnterpriseDifyInstanceBasicDependence12559)
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
					"dify_instance_name": name,
					"workspace_id":       "12345678",
					"payment_type":       "PREPAY",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dify_instance_name": name,
						"workspace_id":       CHECKSET,
						"payment_type":       "PREPAY",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"dify_instance_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dify_instance_name": name + "_update",
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

var AlicloudDmsEnterpriseDifyInstanceMap12559 = map[string]string{
	"status": CHECKSET,
}

func AlicloudDmsEnterpriseDifyInstanceBasicDependence12559(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test DmsEnterprise DifyInstance. <<< Resource test cases, automatically generated.
