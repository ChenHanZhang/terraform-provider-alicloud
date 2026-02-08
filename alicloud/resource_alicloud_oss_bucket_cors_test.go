package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Oss BucketCors. >>> Resource test cases, automatically generated.
// Case BucketCors测试 6634
func TestAccAliCloudOssBucketCors_basic6634(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_cors.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketCorsMap6634)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketCors")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccoss%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketCorsBasicDependence6634)
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
					"response_vary": "true",
					"cors_rule": []map[string]interface{}{
						{
							"allowed_methods": []string{
								"GET"},
							"allowed_origins": []string{
								"*"},
							"allowed_headers": []string{
								"x-oss-test", "x-oss-abc", "x-oss-123"},
							"max_age_seconds": "1000",
							"expose_headers": []string{
								"osstest"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":        CHECKSET,
						"response_vary": "true",
						"cors_rule.#":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"response_vary": "false",
					"cors_rule": []map[string]interface{}{
						{
							"allowed_methods": []string{
								"POST", "HEAD", "GET", "DELETE"},
							"allowed_origins": []string{
								"oss.aliyuncs.com", "1.test.com", "2.test.com"},
							"allowed_headers": []string{
								"*"},
							"max_age_seconds": "100",
							"expose_headers": []string{
								"osstest", "abc", "test"},
						},
						{
							"allowed_methods": []string{
								"GET"},
							"allowed_origins": []string{
								"*"},
							"allowed_headers": []string{
								"*"},
							"max_age_seconds": "200",
							"expose_headers":  []string{},
						},
						{
							"allowed_methods": []string{
								"PUT"},
							"allowed_origins": []string{
								"allow.aliyuncs.com.", "*.aliyuncs.com"},
							"allowed_headers": []string{
								"test-oss"},
							"expose_headers": []string{},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"response_vary": "false",
						"cors_rule.#":   "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cors_rule": []map[string]interface{}{
						{
							"allowed_methods": []string{
								"GET", "PUT"},
							"allowed_origins": []string{
								"*"},
							"allowed_headers": []string{
								"test", "abc"},
						},
						{
							"allowed_methods": []string{
								"GET"},
							"allowed_origins": []string{
								"*"},
							"allowed_headers": []string{
								"abc", "def"},
							"max_age_seconds": "150",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cors_rule.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cors_rule": []map[string]interface{}{
						{
							"allowed_methods": []string{
								"GET", "PUT", "POST", "HEAD"},
							"allowed_origins": []string{
								"test.a.com", "test.b.com", "test.c.com"},
							"allowed_headers": []string{
								"test-a", "test-b", "test-c", "test-d"},
							"expose_headers": []string{
								"test-a", "test-b", "test-c"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cors_rule.#": "1",
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

var AlicloudOssBucketCorsMap6634 = map[string]string{}

func AlicloudOssBucketCorsBasicDependence6634(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_oss_bucket" "CreateBucket" {
        storage_class = "Standard"
}


`, name)
}

// Test Oss BucketCors. <<< Resource test cases, automatically generated.
