// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Esa CustomResponseCodeRule. >>> Resource test cases, automatically generated.
// Case CustomResponseCodeRule_test 12096
func TestAccAliCloudEsaCustomResponseCodeRule_basic12096(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_custom_response_code_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaCustomResponseCodeRuleMap12096)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaCustomResponseCodeRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaCustomResponseCodeRuleBasicDependence12096)
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
					"page_id":      "0",
					"site_id":      "${alicloud_esa_site.resource_Site_CustomResponseCodeRule_test.id}",
					"return_code":  "400",
					"rule_enable":  "on",
					"rule":         "(http.host eq \\\"video.example.com\\\")",
					"sequence":     "1",
					"site_version": "0",
					"rule_name":    "test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"page_id":      CHECKSET,
						"site_id":      CHECKSET,
						"return_code":  CHECKSET,
						"rule_enable":  "on",
						"rule":         "(http.host eq \"video.example.com\")",
						"sequence":     "1",
						"site_version": "0",
						"rule_name":    "test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"return_code": "200",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"return_code": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rule_name": "test_modify",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rule_name": "test_modify",
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
					"rule": "(http.request.uri eq \\\"/content?page=1234\\\")",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rule": "(http.request.uri eq \"/content?page=1234\")",
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudEsaCustomResponseCodeRuleMap12096 = map[string]string{
	"config_id": CHECKSET,
}

func AlicloudEsaCustomResponseCodeRuleBasicDependence12096(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_CustomResponseCodeRule_test" {
  type         = "NS"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "overseas"
  auto_pay     = true
  plan_name    = "basic"
}

resource "alicloud_esa_site" "resource_Site_CustomResponseCodeRule_test" {
  site_name   = "hyhtest.cn"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_CustomResponseCodeRule_test.id
  coverage    = "overseas"
  access_type = "NS"
}


`, name)
}

// Test Esa CustomResponseCodeRule. <<< Resource test cases, automatically generated.
