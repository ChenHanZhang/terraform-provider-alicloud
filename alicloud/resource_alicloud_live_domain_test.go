// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Live Domain. >>> Resource test cases, automatically generated.
// Case 域名测试_Terraform 11928
func TestAccAliCloudLiveDomain_basic11928(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_live_domain.default"
	ra := resourceAttrInit(resourceId, AlicloudLiveDomainMap11928)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &LiveServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeLiveDomain")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacclive%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudLiveDomainBasicDependence11928)
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
					"domain_type":       "liveVideo",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"scope":             "overseas",
					"domain_name":       name,
					"region":            "ap-southeast-1",
					"check_url":         "http://${{ref(variable, domainName)}}/test.html",
					"top_level_domain":  " ",
					"status":            "online",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"domain_type":       "liveVideo",
						"resource_group_id": CHECKSET,
						"scope":             "overseas",
						"domain_name":       name,
						"region":            "ap-southeast-1",
						"check_url":         CHECKSET,
						"top_level_domain":  " ",
						"status":            "online",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"scope":             "global",
					"owner_id":          "1511928242963727",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"scope":             "global",
						"owner_id":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "offline",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "offline",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "online",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "online",
					}),
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
				ImportStateVerifyIgnore: []string{"check_url", "owner_id", "top_level_domain"},
			},
		},
	})
}

var AlicloudLiveDomainMap11928 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudLiveDomainBasicDependence11928(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "domain_name" {
  default = "antang-terraform-domain20251216090912120303"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Case test-domain 7582
func TestAccAliCloudLiveDomain_basic7582(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_live_domain.default"
	ra := resourceAttrInit(resourceId, AlicloudLiveDomainMap7582)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &LiveServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeLiveDomain")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacclive%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudLiveDomainBasicDependence7582)
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
					"domain_type":       "liveVideo",
					"domain_name":       name,
					"region":            "cn-shanghai",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"scope":             "domestic",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"domain_type":       "liveVideo",
						"domain_name":       name,
						"region":            "cn-shanghai",
						"resource_group_id": CHECKSET,
						"scope":             "domestic",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"scope":             "global",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"scope":             "global",
					}),
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
				ImportStateVerifyIgnore: []string{"check_url", "owner_id", "top_level_domain"},
			},
		},
	})
}

var AlicloudLiveDomainMap7582 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudLiveDomainBasicDependence7582(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "domain_name" {
  default = "antangtest01.alivecdn.com"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Test Live Domain. <<< Resource test cases, automatically generated.
