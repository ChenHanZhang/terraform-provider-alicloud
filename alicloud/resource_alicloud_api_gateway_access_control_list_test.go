package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ApiGateway AccessControlList. >>> Resource test cases, automatically generated.
// Case ipv4测试用例 6378
func TestAccAliCloudApiGatewayAccessControlList_basic6378(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_api_gateway_access_control_list.default"
	ra := resourceAttrInit(resourceId, AlicloudApiGatewayAccessControlListMap6378)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApiGatewayServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApiGatewayAccessControlList")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApiGatewayAccessControlListBasicDependence6378)
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
					"access_control_list_name": name,
					"acl_entrys": []map[string]interface{}{
						{
							"acl_entry_ip":      "128.0.0.1/32",
							"acl_entry_comment": "test comment",
						},
					},
					"address_ip_version": "ipv4",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_control_list_name": name,
						"acl_entrys.#":             "1",
						"address_ip_version":       "ipv4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"acl_entrys": []map[string]interface{}{
						{
							"acl_entry_ip":      "163.235.23.45/32",
							"acl_entry_comment": "test",
						},
						{
							"acl_entry_ip":      "128.0.0.1/32",
							"acl_entry_comment": "modify comment",
						},
						{
							"acl_entry_ip":      "123.45.34.23/32",
							"acl_entry_comment": "tetete",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"acl_entrys.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"acl_entrys": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"acl_entrys.#": "0",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"acl_entrys"},
			},
		},
	})
}

var AlicloudApiGatewayAccessControlListMap6378 = map[string]string{}

func AlicloudApiGatewayAccessControlListBasicDependence6378(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_api_gateway_instance" "defaultxywS8c" {
  instance_name = "1768552974"
  instance_spec = "api.s1.small"
  https_policy  = "HTTPS2_TLS1_0"
  zone_id       = "cn-hangzhou-MAZ6"
  payment_type  = "PayAsYouGo"
}


`, name)
}

// Case ipv6测试用例 6402
func TestAccAliCloudApiGatewayAccessControlList_basic6402(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_api_gateway_access_control_list.default"
	ra := resourceAttrInit(resourceId, AlicloudApiGatewayAccessControlListMap6402)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApiGatewayServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApiGatewayAccessControlList")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApiGatewayAccessControlListBasicDependence6402)
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
					"address_ip_version":       "ipv6",
					"access_control_list_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_ip_version":       "ipv6",
						"access_control_list_name": name,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"acl_entrys"},
			},
		},
	})
}

var AlicloudApiGatewayAccessControlListMap6402 = map[string]string{}

func AlicloudApiGatewayAccessControlListBasicDependence6402(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_api_gateway_instance" "defaultVKfqBN" {
  instance_name = "1768552976"
  instance_spec = "api.s1.small"
  https_policy  = "HTTPS2_TLS1_0"
  zone_id       = "cn-hangzhou-MAZ6"
  payment_type  = "PayAsYouGo"
}


`, name)
}

// Test ApiGateway AccessControlList. <<< Resource test cases, automatically generated.
