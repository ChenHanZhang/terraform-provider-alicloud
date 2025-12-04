package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ESA HttpResponseHeaderModificationRule. >>> Resource test cases, automatically generated.
// Case httpResponseHeaderModificationRule_test
func TestAccAliCloudESAHttpResponseHeaderModificationRulehttpResponseHeaderModificationRule_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_http_response_header_modification_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAHttpResponseHeaderModificationRulehttpResponseHeaderModificationRule_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaHttpResponseHeaderModificationRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAHttpResponseHeaderModificationRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAHttpResponseHeaderModificationRulehttpResponseHeaderModificationRule_testBasicDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":     "${data.alicloud_esa_sites.default.sites.0.id}",
					"rule_enable": "on",
					"response_header_modification": []map[string]interface{}{

						{
							"value":     "add",
							"operation": "add",
							"name":      "testadd",
						},

						{
							"operation": "del",
							"name":      "testdel",
						},

						{
							"value":     "modify",
							"operation": "modify",
							"name":      "testmodify",
						},
					},
					"rule":         "(http.host eq \\\"video.example.com\\\")",
					"site_version": "0",
					"rule_name":    "testResponseHeader",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule_name": "testResponseHeader_modify",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule_enable": "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule": "(http.request.uri eq \\\"/content?page=1234\\\")",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"response_header_modification": []map[string]interface{}{

						{
							"value":     "add1",
							"operation": "add",
							"name":      "testadd1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule_enable": "on",
					"rule":        "(http.host eq \\\"api.example.com\\\")",
					"rule_name":   "test_httpResponseHeader_last",
				}),
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AliCloudESAHttpResponseHeaderModificationRulehttpResponseHeaderModificationRule_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAHttpResponseHeaderModificationRulehttpResponseHeaderModificationRule_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
  site_name = "gositecdn.cn"
}

`, name)
}

// Test ESA HttpResponseHeaderModificationRule. <<< Resource test cases, automatically generated.
// Test Esa HttpResponseHeaderModificationRule. >>> Resource test cases, automatically generated.
// Case httpResponseHeaderModificationRule_test 11931
func TestAccAliCloudEsaHttpResponseHeaderModificationRule_basic11931(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_http_response_header_modification_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaHttpResponseHeaderModificationRuleMap11931)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaHttpResponseHeaderModificationRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaHttpResponseHeaderModificationRuleBasicDependence11931)
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
					"site_id":     "${alicloud_esa_site.resource_Site_HttpResponseHeaderModificationRule_test.id}",
					"rule_enable": "on",
					"response_header_modification": []map[string]interface{}{
						{
							"type":      "static",
							"value":     "add",
							"operation": "add",
							"name":      "testadd",
						},
						{
							"type":      "static",
							"operation": "del",
							"name":      "testdel",
						},
						{
							"type":      "static",
							"value":     "modify",
							"operation": "modify",
							"name":      "testmodify",
						},
					},
					"rule":         "(http.host eq \\\\\\\"video.example.com\\\\\\\")",
					"sequence":     "1",
					"site_version": "0",
					"rule_name":    "testResponseHeader",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"site_id":                        CHECKSET,
						"rule_enable":                    "on",
						"response_header_modification.#": "3",
						"rule":                           "(http.host eq \\\"video.example.com\\\")",
						"sequence":                       "1",
						"site_version":                   "0",
						"rule_name":                      "testResponseHeader",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule_name": "testResponseHeader_modify",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rule_name": "testResponseHeader_modify",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule_enable": "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rule_enable": "off",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule": "(http.request.uri eq \\\\\\\"/content?page=1234\\\\\\\")",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rule": "(http.request.uri eq \\\"/content?page=1234\\\")",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"response_header_modification": []map[string]interface{}{
						{
							"type":      "dynamic",
							"value":     "ip.geoip.country",
							"operation": "add",
							"name":      "x-ip-country",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"response_header_modification.#": "1",
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

var AlicloudEsaHttpResponseHeaderModificationRuleMap11931 = map[string]string{
	"config_id": CHECKSET,
}

func AlicloudEsaHttpResponseHeaderModificationRuleBasicDependence11931(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_HttpResponseHeaderModificationRule_test" {
  type         = "NS"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "overseas"
  auto_pay     = true
  plan_name    = "basic"
}

resource "alicloud_esa_site" "resource_Site_HttpResponseHeaderModificationRule_test" {
  site_name   = "pltestpl.cn"
  instance_id = alicloud_esa_rate_plan_instance.resource_HttpResponseHeaderModificationRule_test.id
  coverage    = "overseas"
  access_type = "NS"
}


`, name)
}

// Test Esa HttpResponseHeaderModificationRule. <<< Resource test cases, automatically generated.
