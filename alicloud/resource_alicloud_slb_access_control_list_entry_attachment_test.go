// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Slb AccessControlListEntryAttachment. >>> Resource test cases, automatically generated.
// Case aclEntry 10570
func TestAccAliCloudSlbAccessControlListEntryAttachment_basic10570(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_slb_access_control_list_entry_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudSlbAccessControlListEntryAttachmentMap10570)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSlbAccessControlListEntryAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccslb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSlbAccessControlListEntryAttachmentBasicDependence10570)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"comment": "test-111",
					"entry":   "10.0.17.0/24",
					"acl_id":  "${alicloud_slb_acl.创建acl.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment": "test-111",
						"entry":   "10.0.17.0/24",
						"acl_id":  CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"comment": "test-test-1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"comment": "test-test-1",
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

var AlicloudSlbAccessControlListEntryAttachmentMap10570 = map[string]string{}

func AlicloudSlbAccessControlListEntryAttachmentBasicDependence10570(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_slb_acl" "创建acl" {
}


`, name)
}

// Test Slb AccessControlListEntryAttachment. <<< Resource test cases, automatically generated.
