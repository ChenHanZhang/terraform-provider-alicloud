// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cms AddonRelease. >>> Resource test cases, automatically generated.
// Case AddonRelease常规测试 8556
func TestAccAliCloudCmsAddonRelease_basic8556(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_addon_release.default"
	ra := resourceAttrInit(resourceId, AlicloudCmsAddonReleaseMap8556)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCmsAddonRelease")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCmsAddonReleaseBasicDependence8556)
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
					"addon_version":         "0.0.2",
					"aliyun_lang":           "zh",
					"addon_name":            "cs-gpu",
					"integration_policy_id": "env-79520003404088b43c8aabaf4bb8",
					"workspace":             "sls-mall",
					"addon_release_name":    name,
					"dry_run":               "true",
					"values":                "{\\\"install\\\":{\\\"mode\\\":\\\"auto-install\\\",\\\"listenPort\\\":\\\"9400\\\"},\\\"discoverMode\\\":\\\"instances\\\",\\\"discover\\\":{\\\"instances\\\":\\\"worker-k8s-for-cs-c126d87c76218487e83ab322017f11b44\\\"},\\\"scrapeInterval\\\":\\\"15\\\",\\\"enableSecuritecs-nodeyGroupInjection\\\":\\\"true\\\",\\\"metricTags\\\":\\\"\\\"}",
					"env_type":              "CS",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"addon_version":         "0.0.2",
						"aliyun_lang":           "zh",
						"addon_name":            "cs-gpu",
						"integration_policy_id": "env-79520003404088b43c8aabaf4bb8",
						"workspace":             "sls-mall",
						"addon_release_name":    name,
						"dry_run":               "true",
						"values":                CHECKSET,
						"env_type":              "CS",
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
				ImportStateVerifyIgnore: []string{"dry_run", "values"},
			},
		},
	})
}

var AlicloudCmsAddonReleaseMap8556 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudCmsAddonReleaseBasicDependence8556(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Cms AddonRelease. <<< Resource test cases, automatically generated.
