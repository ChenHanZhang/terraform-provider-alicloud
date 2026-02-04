// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Eiam Application. >>> Resource test cases, automatically generated.
// Case test-application_copy 4438
func TestAccAliCloudEiamApplication_basic4438(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eiam_application.default"
	ra := resourceAttrInit(resourceId, AlicloudEiamApplicationMap4438)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EiamServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEiamApplication")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceiam%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEiamApplicationBasicDependence4438)
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
					"instance_id":             "${alicloud_eiam_instance.defaultiNPi8X.id}",
					"application_source_type": "urn:alibaba:idaas:app:source:standard",
					"sso_type":                "saml2",
					"application_name":        name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id":             CHECKSET,
						"application_source_type": "urn:alibaba:idaas:app:source:standard",
						"sso_type":                "saml2",
						"application_name":        name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-lcw-update-desc",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test-lcw-update-desc",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"application_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"application_name": name + "_update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"saml_sso_config": []map[string]interface{}{
						{
							"sp_sso_acs_url": "https://signin.aliyun.com/saml-role/sso",
							"sp_entity_id":   "urn:alibaba:cloudcomputing",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"grant_scopes": []string{
						"urn:alibaba:idaas:scope:user:read_all"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"grant_scopes.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sso_status": "disabled",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sso_status": "disabled",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sso_status": "enabled",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sso_status": "enabled",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "disabled",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "disabled",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "enabled",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "enabled",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"application_template_params"},
			},
		},
	})
}

var AlicloudEiamApplicationMap4438 = map[string]string{}

func AlicloudEiamApplicationBasicDependence4438(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_eiam_instance" "defaultiNPi8X" {
  description = "lcw-test"
}


`, name)
}

// Test Eiam Application. <<< Resource test cases, automatically generated.
