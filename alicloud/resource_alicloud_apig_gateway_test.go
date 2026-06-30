package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Apig Gateway. >>> Resource test cases, automatically generated.
// Case gateway_crud_alt2 12899
func TestAccAliCloudApigGateway_basic12899(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_gateway.default"
	ra := resourceAttrInit(resourceId, AlicloudApigGatewayMap12899)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigGateway")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigGatewayBasicDependence12899)
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
					"network_access_config": []map[string]interface{}{
						{
							"type": "InternetAndIntranet",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"vpc": []map[string]interface{}{
						{
							"vpc_id": "${alicloud_vpc.gateway_vpc_pre3.id}",
						},
					},
					"gateway_edition": "Professional",
					"gateway_type":    "API",
					"zones": []map[string]interface{}{
						{
							"zone_id":    "cn-hangzhou-i",
							"vswitch_id": "${alicloud_vswitch.gateway_vswitch_pre3.id}",
						},
					},
					"vswitch": []map[string]interface{}{
						{
							"vswitch_id": "${alicloud_vswitch.gateway_vswitch_pre3.id}",
						},
					},
					"zone_config": []map[string]interface{}{
						{
							"select_option": "Auto",
						},
					},
					"payment_type": "PayAsYouGo",
					"gateway_name": name,
					"spec":         "apigw.small.x1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"gateway_edition":   "Professional",
						"gateway_type":      "API",
						"zones.#":           "1",
						"payment_type":      "PayAsYouGo",
						"gateway_name":      name,
						"spec":              "apigw.small.x1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"gateway_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_name": name + "_update",
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
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"log_config", "network_access_config", "zone_config"},
			},
		},
	})
}

var AlicloudApigGatewayMap12899 = map[string]string{
	"target_version":   CHECKSET,
	"create_from":      CHECKSET,
	"version":          CHECKSET,
	"status":           CHECKSET,
	"create_time":      CHECKSET,
	"load_balancers.#": CHECKSET,
	"security_group.#": CHECKSET,
	"environments.#":   CHECKSET,
	"update_time":      CHECKSET,
	"expire_time":      CHECKSET,
}

func AlicloudApigGatewayBasicDependence12899(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "gateway_vpc_pre3" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "gateway-test-vpc3"
}

resource "alicloud_vswitch" "gateway_vswitch_pre3" {
  is_default   = false
  vpc_id       = alicloud_vpc.gateway_vpc_pre3.id
  zone_id      = "cn-hangzhou-i"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "gateway-test-vswitch3"
}


`, name)
}

// Case gateway_crud_alt 12900
func TestAccAliCloudApigGateway_basic12900(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_gateway.default"
	ra := resourceAttrInit(resourceId, AlicloudApigGatewayMap12900)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigGateway")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigGatewayBasicDependence12900)
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
					"network_access_config": []map[string]interface{}{
						{
							"type": "Intranet",
						},
					},
					"zone_config": []map[string]interface{}{
						{
							"select_option": "Manual",
						},
					},
					"vpc": []map[string]interface{}{
						{
							"vpc_id": "${alicloud_vpc.gateway_vpc_pre2.id}",
						},
					},
					"gateway_edition": "Serverless",
					"gateway_type":    "AI",
					"payment_type":    "Subscription",
					"gateway_name":    name,
					"spec":            "apigw.small.x1",
					"zones": []map[string]interface{}{
						{
							"zone_id":    "cn-hangzhou-i",
							"vswitch_id": "${alicloud_vswitch.gateway_vswitch_pre2.id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_edition": "Serverless",
						"gateway_type":    "AI",
						"payment_type":    "Subscription",
						"gateway_name":    name,
						"spec":            "apigw.small.x1",
						"zones.#":         "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"gateway_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"log_config", "network_access_config", "zone_config"},
			},
		},
	})
}

var AlicloudApigGatewayMap12900 = map[string]string{
	"target_version":   CHECKSET,
	"create_from":      CHECKSET,
	"version":          CHECKSET,
	"status":           CHECKSET,
	"create_time":      CHECKSET,
	"load_balancers.#": CHECKSET,
	"security_group.#": CHECKSET,
	"environments.#":   CHECKSET,
	"update_time":      CHECKSET,
	"expire_time":      CHECKSET,
}

func AlicloudApigGatewayBasicDependence12900(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "gateway_vpc_pre2" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "gateway-test-vpc2"
}

resource "alicloud_vswitch" "gateway_vswitch_pre2" {
  is_default   = false
  vpc_id       = alicloud_vpc.gateway_vpc_pre2.id
  zone_id      = "cn-hangzhou-i"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "gateway-test-vswitch2"
}


`, name)
}

// Case gateway_tags 12901
func TestAccAliCloudApigGateway_basic12901(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_gateway.default"
	ra := resourceAttrInit(resourceId, AlicloudApigGatewayMap12901)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigGateway")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigGatewayBasicDependence12901)
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
					"network_access_config": []map[string]interface{}{
						{
							"type": "Internet",
						},
					},
					"vswitch": []map[string]interface{}{
						{
							"vswitch_id": "${alicloud_vswitch.gateway_tags_vswitch_pre.id}",
						},
					},
					"zone_config": []map[string]interface{}{
						{
							"select_option": "Auto",
						},
					},
					"vpc": []map[string]interface{}{
						{
							"vpc_id": "${alicloud_vpc.gateway_tags_vpc_pre.id}",
						},
					},
					"gateway_edition": "Professional",
					"gateway_type":    "API",
					"payment_type":    "PayAsYouGo",
					"gateway_name":    name,
					"spec":            "apigw.small.x1",
					"zones": []map[string]interface{}{
						{
							"zone_id":    "cn-hangzhou-i",
							"vswitch_id": "${alicloud_vswitch.gateway_tags_vswitch_pre.id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_edition": "Professional",
						"gateway_type":    "API",
						"payment_type":    "PayAsYouGo",
						"gateway_name":    name,
						"spec":            "apigw.small.x1",
						"zones.#":         "1",
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
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
				ImportStateVerifyIgnore: []string{"log_config", "network_access_config", "zone_config"},
			},
		},
	})
}

var AlicloudApigGatewayMap12901 = map[string]string{
	"target_version":   CHECKSET,
	"create_from":      CHECKSET,
	"version":          CHECKSET,
	"status":           CHECKSET,
	"create_time":      CHECKSET,
	"load_balancers.#": CHECKSET,
	"security_group.#": CHECKSET,
	"environments.#":   CHECKSET,
	"update_time":      CHECKSET,
	"expire_time":      CHECKSET,
}

func AlicloudApigGatewayBasicDependence12901(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "gateway_tags_vpc_pre" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "gateway-tags-vpc"
}

resource "alicloud_vswitch" "gateway_tags_vswitch_pre" {
  is_default   = false
  vpc_id       = alicloud_vpc.gateway_tags_vpc_pre.id
  zone_id      = "cn-hangzhou-i"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "gateway-tags-vswitch"
}


`, name)
}

// Case gateway_crud 12902
func TestAccAliCloudApigGateway_basic12902(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_gateway.default"
	ra := resourceAttrInit(resourceId, AlicloudApigGatewayMap12902)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigGateway")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigGatewayBasicDependence12902)
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
					"network_access_config": []map[string]interface{}{
						{
							"type": "Internet",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"vpc": []map[string]interface{}{
						{
							"vpc_id": "${alicloud_vpc.gateway_vpc_pre.id}",
						},
					},
					"gateway_edition": "Professional",
					"gateway_type":    "API",
					"zones": []map[string]interface{}{
						{
							"zone_id":    "cn-hangzhou-i",
							"vswitch_id": "${alicloud_vswitch.gateway_vswitch_pre.id}",
						},
					},
					"vswitch": []map[string]interface{}{
						{
							"vswitch_id": "${alicloud_vswitch.gateway_vswitch_pre.id}",
						},
					},
					"zone_config": []map[string]interface{}{
						{
							"select_option": "Auto",
						},
					},
					"payment_type": "PayAsYouGo",
					"gateway_name": name,
					"spec":         "apigw.small.x1",
					"log_config": []map[string]interface{}{
						{
							"sls": []map[string]interface{}{
								{
									"enable": "true",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"gateway_edition":   "Professional",
						"gateway_type":      "API",
						"zones.#":           "1",
						"payment_type":      "PayAsYouGo",
						"gateway_name":      name,
						"spec":              "apigw.small.x1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"gateway_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_name": name + "_update",
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
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"log_config", "network_access_config", "zone_config"},
			},
		},
	})
}

var AlicloudApigGatewayMap12902 = map[string]string{
	"target_version":   CHECKSET,
	"create_from":      CHECKSET,
	"version":          CHECKSET,
	"status":           CHECKSET,
	"create_time":      CHECKSET,
	"load_balancers.#": CHECKSET,
	"security_group.#": CHECKSET,
	"environments.#":   CHECKSET,
	"update_time":      CHECKSET,
	"expire_time":      CHECKSET,
}

func AlicloudApigGatewayBasicDependence12902(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "gateway_vpc_pre" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "gateway-test-vpc"
}

resource "alicloud_vswitch" "gateway_vswitch_pre" {
  is_default   = false
  vpc_id       = alicloud_vpc.gateway_vpc_pre.id
  zone_id      = "cn-hangzhou-i"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "gateway-test-vswitch"
}


`, name)
}

// Case apigateway_postpay_zones 11572
func TestAccAliCloudApigGateway_basic11572(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_gateway.default"
	ra := resourceAttrInit(resourceId, AlicloudApigGatewayMap11572)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigGateway")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigGatewayBasicDependence11572)
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
					"gateway_name": name,
					"spec":         "apigw.small.x1",
					"vpc": []map[string]interface{}{
						{
							"vpc_id": "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
					},
					"network_access_config": []map[string]interface{}{
						{
							"type": "Intranet",
						},
					},
					"zone_config": []map[string]interface{}{
						{
							"select_option": "Manual",
						},
					},
					"vswitch": []map[string]interface{}{
						{
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
						},
					},
					"log_config": []map[string]interface{}{
						{
							"sls": []map[string]interface{}{
								{
									"enable": "false",
								},
							},
						},
					},
					"payment_type": "PayAsYouGo",
					"gateway_type": "API",
					"zones": []map[string]interface{}{
						{
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"zone_id":    "${alicloud_vswitch.defaultVSwitch.zone_id}",
							"name":       "杭州 可用区J",
						},
						{
							"vswitch_id": "${alicloud_vswitch.defaultAskkJp.id}",
							"zone_id":    "${alicloud_vswitch.defaultAskkJp.zone_id}",
							"name":       "杭州 可用区K",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_name": name,
						"spec":         "apigw.small.x1",
						"payment_type": "PayAsYouGo",
						"gateway_type": "API",
						"zones.#":      "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"gateway_name":      name + "_update",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_name":      name + "_update",
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"log_config", "network_access_config", "zone_config"},
			},
		},
	})
}

var AlicloudApigGatewayMap11572 = map[string]string{
	"target_version":   CHECKSET,
	"create_from":      CHECKSET,
	"version":          CHECKSET,
	"status":           CHECKSET,
	"create_time":      CHECKSET,
	"load_balancers.#": CHECKSET,
	"security_group.#": CHECKSET,
	"environments.#":   CHECKSET,
	"update_time":      CHECKSET,
	"expire_time":      CHECKSET,
}

func AlicloudApigGatewayBasicDependence11572(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id     = alicloud_vpc.defaultvpc.id
  cidr_block = "172.16.4.0/24"
  zone_id    = "cn-hangzhou-j"
}

resource "alicloud_vswitch" "defaultAskkJp" {
  vpc_id     = alicloud_vpc.defaultvpc.id
  cidr_block = "172.16.13.0/24"
  zone_id    = "cn-hangzhou-k"
}


`, name)
}

// Test Apig Gateway. <<< Resource test cases, automatically generated.
