// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Rds GadInstance. >>> Resource test cases, automatically generated.
// Case 测试创建-杨涛 1020
func TestAccAliCloudRdsGadInstance_basic1020(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_gad_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsGadInstanceMap1020)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsGadInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsGadInstanceBasicDependence1020)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
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

var AlicloudRdsGadInstanceMap1020 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudRdsGadInstanceBasicDependence1020(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Rds GadInstance. <<< Resource test cases, automatically generated.
