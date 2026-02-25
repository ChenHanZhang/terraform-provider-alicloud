package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test MessageService Endpoint. >>> Resource test cases, automatically generated.
// Case Endpoint测试用例 9855
func TestAccAliCloudMessageServiceEndpoint_basic9855(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_message_service_endpoint.default"
	ra := resourceAttrInit(resourceId, AlicloudMessageServiceEndpointMap9855)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MessageServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMessageServiceEndpoint")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmessageservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMessageServiceEndpointBasicDependence9855)
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
					"endpoint_type": "public",
					"cidr_list": []map[string]interface{}{
						{
							"cidr":         "0.0.0.0/0",
							"acl_strategy": "allow",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"endpoint_type": "public",
						"cidr_list.#":   "1",
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
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cidr_list": []map[string]interface{}{
						{
							"cidr":         "0.0.0.0/0",
							"acl_strategy": "allow",
						},
						{
							"cidr":         "192.168.1.1",
							"acl_strategy": "allow",
						},
						{
							"cidr":         "192.168.1.2",
							"acl_strategy": "allow",
						},
						{
							"cidr":         "192.168.1.3/23",
							"acl_strategy": "allow",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cidr_list.#": "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cidr_list": []map[string]interface{}{
						{
							"cidr":         "0.0.0.0/0",
							"acl_strategy": "allow",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cidr_list.#": "1",
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

var AlicloudMessageServiceEndpointMap9855 = map[string]string{}

func AlicloudMessageServiceEndpointBasicDependence9855(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case Endpoint测试用例_v1.1 10073
func TestAccAliCloudMessageServiceEndpoint_basic10073(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_message_service_endpoint.default"
	ra := resourceAttrInit(resourceId, AlicloudMessageServiceEndpointMap10073)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MessageServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMessageServiceEndpoint")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmessageservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMessageServiceEndpointBasicDependence10073)
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
					"endpoint_type": "public",
					"cidr_list": []map[string]interface{}{
						{
							"cidr":         "0.0.0.0/0",
							"acl_strategy": "allow",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"endpoint_type": "public",
						"cidr_list.#":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cidr_list": []map[string]interface{}{
						{
							"cidr":         "0.0.0.0/0",
							"acl_strategy": "allow",
						},
						{
							"cidr":         "192.168.1.1",
							"acl_strategy": "allow",
						},
						{
							"cidr":         "192.168.1.2",
							"acl_strategy": "allow",
						},
						{
							"cidr":         "192.168.1.3/23",
							"acl_strategy": "allow",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cidr_list.#": "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cidr_list": []map[string]interface{}{
						{
							"cidr":         "0.0.0.0/0",
							"acl_strategy": "allow",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cidr_list.#": "1",
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

var AlicloudMessageServiceEndpointMap10073 = map[string]string{}

func AlicloudMessageServiceEndpointBasicDependence10073(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test MessageService Endpoint. <<< Resource test cases, automatically generated.

// Case Endpoint测试用例 9856
func TestAccAliCloudMessageServiceEndpoint_basic9856(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_message_service_endpoint.default"
	ra := resourceAttrInit(resourceId, AliCloudMessageServiceEndpointMap9855)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MessageServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMessageServiceEndpoint")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%smessageserviceendpoint%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudMessageServiceEndpointBasicDependence9855)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  nil,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"endpoint_type":    "public",
					"endpoint_enabled": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"endpoint_type":    "public",
						"endpoint_enabled": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"endpoint_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"endpoint_enabled": "true",
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
