// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test OpenSearch AppGroupCredential. >>> Resource test cases, automatically generated.
// Case app_group_credential_test 7515
func TestAccAliCloudOpenSearchAppGroupCredential_basic7515(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_open_search_app_group_credential.default"
	ra := resourceAttrInit(resourceId, AlicloudOpenSearchAppGroupCredentialMap7515)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OpenSearchServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOpenSearchAppGroupCredential")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccopensearch%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOpenSearchAppGroupCredentialBasicDependence7515)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shanghai"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"enabled":      "true",
					"type":         "api-token",
					"app_group_id": "${alicloud_open_search_app_group.defaultWQxaIV.group_id}",
					"dry_run":      "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enabled":      "true",
						"type":         "api-token",
						"app_group_id": CHECKSET,
						"dry_run":      "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enabled": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enabled": "false",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dry_run"},
			},
		},
	})
}

var AlicloudOpenSearchAppGroupCredentialMap7515 = map[string]string{
	"app_group_credential_id": CHECKSET,
}

func AlicloudOpenSearchAppGroupCredentialBasicDependence7515(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_open_search_app_group" "defaultWQxaIV" {
  quota {
    spec             = "opensearch.share.common"
    doc_size         = "1"
    compute_resource = "20"
  }
  app_group_name = "credential_test"
  type           = "standard"
}


`, name)
}

// Test OpenSearch AppGroupCredential. <<< Resource test cases, automatically generated.
