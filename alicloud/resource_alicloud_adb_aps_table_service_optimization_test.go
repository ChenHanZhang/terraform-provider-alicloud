package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Adb ApsTableServiceOptimization. >>> Resource test cases, automatically generated.
// Case 5532
func TestAccAliCloudAdbApsTableServiceOptimization_basic5532(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_aps_table_service_optimization.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbApsTableServiceOptimizationMap5532)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbApsTableServiceOptimization")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbapstableserviceoptimization%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbApsTableServiceOptimizationBasicDependence5532)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "off",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "on",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "on",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudAdbApsTableServiceOptimizationMap5532 = map[string]string{
	"status": "on",
}

func AlicloudAdbApsTableServiceOptimizationBasicDependence5532(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "db_cluster_id" {
  default = "amv-bp1u30028ta370f7"
}


`, name)
}

// Case 5532  twin
func TestAccAliCloudAdbApsTableServiceOptimization_basic5532_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_aps_table_service_optimization.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbApsTableServiceOptimizationMap5532)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbApsTableServiceOptimization")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbapstableserviceoptimization%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbApsTableServiceOptimizationBasicDependence5532)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"status":        "on",
					"db_cluster_id": "${var.db_cluster_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":        "on",
						"db_cluster_id": CHECKSET,
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

// Test Adb ApsTableServiceOptimization. <<< Resource test cases, automatically generated.
