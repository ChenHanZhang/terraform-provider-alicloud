package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Oss BucketLogging. >>> Resource test cases, automatically generated.
// Case BucketLogging测试 6484
func TestAccAliCloudOssBucketLogging_basic6484(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_logging.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketLoggingMap6484)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketLogging")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccoss%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketLoggingBasicDependence6484)
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
					"bucket":        "${alicloud_oss_bucket.CreateBucket.id}",
					"target_bucket": "${alicloud_oss_bucket.CreateBucket.id}",
					"target_prefix": "log/",
					"logging_role":  "test-role",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":        CHECKSET,
						"target_bucket": CHECKSET,
						"target_prefix": "log/",
						"logging_role":  "test-role",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"target_bucket": "${alicloud_oss_bucket.CreateLoggingBucket.id}",
					"target_prefix": "logging/",
					"logging_role":  "new-test-role",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"target_bucket": CHECKSET,
						"target_prefix": "logging/",
						"logging_role":  "new-test-role",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"target_bucket": "${alicloud_oss_bucket.CreateBucket.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"target_bucket": CHECKSET,
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

var AlicloudOssBucketLoggingMap6484 = map[string]string{}

func AlicloudOssBucketLoggingBasicDependence6484(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_oss_bucket" "CreateBucket" {
  storage_class = "Standard"
}


`, name)
}

// Test Oss BucketLogging. <<< Resource test cases, automatically generated.
