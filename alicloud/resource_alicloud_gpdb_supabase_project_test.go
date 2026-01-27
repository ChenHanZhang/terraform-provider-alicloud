// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Gpdb SupabaseProject. >>> Resource test cases, automatically generated.
// Case supabase资源测试1215 11921
func TestAccAliCloudGpdbSupabaseProject_basic11921(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_gpdb_supabase_project.default"
	ra := resourceAttrInit(resourceId, AlicloudGpdbSupabaseProjectMap11921)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &GpdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeGpdbSupabaseProject")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccgpdb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudGpdbSupabaseProjectBasicDependence11921)
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
					"project_spec": "1C2G",
					"zone_id":      "cn-beijing-i",
					"vpc_id":       "${alicloud_vpc.defaultaUf3B6.id}",
					"project_name": "supabase_test",
					"security_ip_list": []string{
						"127.0.0.1"},
					"vswitch_id":             "${alicloud_vswitch.default3BTNtz.id}",
					"disk_performance_level": "PL0",
					"storage_size":           "1",
					"account_password":       "Aa123456",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"project_spec":           "1C2G",
						"zone_id":                "cn-beijing-i",
						"vpc_id":                 CHECKSET,
						"project_name":           "supabase_test",
						"security_ip_list.#":     "1",
						"vswitch_id":             CHECKSET,
						"disk_performance_level": "PL0",
						"storage_size":           "1",
						"account_password":       "Aa123456",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_ip_list": []string{
						"0.0.0.0/0"},
					"account_password": "123456Aa",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_ip_list.#": "1",
						"account_password":   "123456Aa",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_ip_list": []string{
						"127.0.0.1", "0.0.0.0/0", "140.205.11.0/24", "140.205.11.11"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_ip_list.#": "4",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_password"},
			},
		},
	})
}

var AlicloudGpdbSupabaseProjectMap11921 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudGpdbSupabaseProjectBasicDependence11921(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultaUf3B6" {
  vpc_name = "vpc_test"
}

resource "alicloud_vswitch" "default3BTNtz" {
  vpc_id       = alicloud_vpc.defaultaUf3B6.id
  cidr_block   = "172.16.18.0/24"
  vswitch_name = "vswitch_test"
  zone_id      = "cn-beijing-i"
}


`, name)
}

// Test Gpdb SupabaseProject. <<< Resource test cases, automatically generated.
