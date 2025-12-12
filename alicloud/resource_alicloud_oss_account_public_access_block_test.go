package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Oss AccountPublicAccessBlock. >>> Resource test cases, automatically generated.
// Case Account阻止公共访问测试 6554
func TestAccAliCloudOssAccountPublicAccessBlock_basic6554(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_account_public_access_block.default"
	ra := resourceAttrInit(resourceId, AlicloudOssAccountPublicAccessBlockMap6554)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssAccountPublicAccessBlock")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccoss%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssAccountPublicAccessBlockBasicDependence6554)
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
					"block_public_access": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"block_public_access": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"block_public_access": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"block_public_access": "false",
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

var AlicloudOssAccountPublicAccessBlockMap6554 = map[string]string{}

func AlicloudOssAccountPublicAccessBlockBasicDependence6554(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Oss AccountPublicAccessBlock. <<< Resource test cases, automatically generated.
