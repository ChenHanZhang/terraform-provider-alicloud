// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cr InternetEndpoint. >>> Resource test cases, automatically generated.
// Case resource_InternetEndpoint_test 12950
func TestAccAliCloudCrInternetEndpoint_basic12950(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_internet_endpoint.default"
	ra := resourceAttrInit(resourceId, AlicloudCrInternetEndpointMap12950)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrInternetEndpoint")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrInternetEndpointBasicDependence12950)
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
					"instance_id": "${alicloud_cr_ee_instance.internet_endpoint_pre_Instance.id}",
					"entries": []map[string]interface{}{
						{
							"comment": "jj",
							"entry":   "127.0.0.9/32",
						},
						{
							"entry": "127.0.0.10/32",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id": CHECKSET,
						"entries.#":   "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"entries": []map[string]interface{}{
						{
							"comment": "aa",
							"entry":   "127.0.0.2/32",
						},
						{
							"entry": "127.0.0.3/32",
						},
						{
							"comment": "bb",
							"entry":   "127.0.0.4/32",
						},
						{
							"comment": "cc",
							"entry":   "127.0.0.5/32",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"entries.#": "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"entries": []map[string]interface{}{
						{
							"comment": "aa",
							"entry":   "127.0.0.2/32",
						},
						{
							"comment": "ff",
							"entry":   "127.0.0.5/32",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"entries.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"entries": []map[string]interface{}{
						{
							"comment": "gg",
							"entry":   "127.0.0.6/32",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"entries.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"entries": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"entries.#": "0",
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

var AlicloudCrInternetEndpointMap12950 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCrInternetEndpointBasicDependence12950(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cr_ee_instance" "internet_endpoint_pre_Instance" {
  instance_name      = "cspec-test-instance"
  default_oss_bucket = "true"
  payment_type       = "Subscription"
  period             = "1"
  instance_type      = "Basic"
}


`, name)
}

// Test Cr InternetEndpoint. <<< Resource test cases, automatically generated.
