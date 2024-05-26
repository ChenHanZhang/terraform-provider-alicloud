package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Oss BucketStyle. >>> Resource test cases, automatically generated.
// Case Style指定Category 6688
func TestAccAliCloudOssBucketStyle_basic6688(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_style.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketStyleMap6688)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketStyle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sossbucketstyle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketStyleBasicDependence6688)
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
					"bucket":     "${alicloud_oss_bucket.CreateBucket.bucket}",
					"style_name": "style-372",
					"content":    "image/resize,p_75,w_75",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":     CHECKSET,
						"style_name": CHECKSET,
						"content":    "image/resize,p_75,w_75",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"category": "document",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"category": "document",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"content": "image/resize,p_75,w_70",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"content": "image/resize,p_75,w_70",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"category": "video",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"category": "video",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bucket":     "${alicloud_oss_bucket.CreateBucket.bucket}",
					"style_name": "style-679",
					"content":    "image/resize,p_75,w_75",
					"category":   "document",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":     CHECKSET,
						"style_name": CHECKSET,
						"content":    "image/resize,p_75,w_75",
						"category":   "document",
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

var AlicloudOssBucketStyleMap6688 = map[string]string{
	"category":    CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudOssBucketStyleBasicDependence6688(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_oss_bucket" "CreateBucket" {
  storage_class = "Standard"
  bucket        = var.name
}


`, name)
}

// Case BucketStyle测试 6687
func TestAccAliCloudOssBucketStyle_basic6687(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_style.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketStyleMap6687)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketStyle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sossbucketstyle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketStyleBasicDependence6687)
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
					"bucket":     "${alicloud_oss_bucket.CreateBucket.bucket}",
					"style_name": "style-373",
					"content":    "image/resize,p_75,w_75",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":     CHECKSET,
						"style_name": CHECKSET,
						"content":    "image/resize,p_75,w_75",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"content": "image/resize,p_75,w_70",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"content": "image/resize,p_75,w_70",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bucket":     "${alicloud_oss_bucket.CreateBucket.bucket}",
					"style_name": "style-293",
					"content":    "image/resize,p_75,w_75",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":     CHECKSET,
						"style_name": CHECKSET,
						"content":    "image/resize,p_75,w_75",
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

var AlicloudOssBucketStyleMap6687 = map[string]string{
	"category":    CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudOssBucketStyleBasicDependence6687(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_oss_bucket" "CreateBucket" {
  storage_class = "Standard"
  bucket        = var.name
}


`, name)
}

// Case Style指定Category 6688  twin
func TestAccAliCloudOssBucketStyle_basic6688_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_style.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketStyleMap6688)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketStyle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sossbucketstyle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketStyleBasicDependence6688)
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
					"bucket":     "${alicloud_oss_bucket.CreateBucket.bucket}",
					"style_name": "style-29",
					"content":    "image/resize,p_75,w_75",
					"category":   "document",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":     CHECKSET,
						"style_name": CHECKSET,
						"content":    "image/resize,p_75,w_75",
						"category":   "document",
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

// Case BucketStyle测试 6687  twin
func TestAccAliCloudOssBucketStyle_basic6687_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_style.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketStyleMap6687)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketStyle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sossbucketstyle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketStyleBasicDependence6687)
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
					"bucket":     "${alicloud_oss_bucket.CreateBucket.bucket}",
					"style_name": "style-878",
					"content":    "image/resize,p_75,w_75",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":     CHECKSET,
						"style_name": CHECKSET,
						"content":    "image/resize,p_75,w_75",
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

// Case Style指定Category 6688  raw
func TestAccAliCloudOssBucketStyle_basic6688_raw(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_style.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketStyleMap6688)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketStyle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sossbucketstyle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketStyleBasicDependence6688)
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
					"bucket":     "${alicloud_oss_bucket.CreateBucket.bucket}",
					"style_name": "style-9",
					"content":    "image/resize,p_75,w_75",
					"category":   "document",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":     CHECKSET,
						"style_name": CHECKSET,
						"content":    "image/resize,p_75,w_75",
						"category":   "document",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"content":  "image/resize,p_75,w_70",
					"category": "video",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"content":  "image/resize,p_75,w_70",
						"category": "video",
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

// Case BucketStyle测试 6687  raw
func TestAccAliCloudOssBucketStyle_basic6687_raw(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_oss_bucket_style.default"
	ra := resourceAttrInit(resourceId, AlicloudOssBucketStyleMap6687)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &OssServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeOssBucketStyle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sossbucketstyle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudOssBucketStyleBasicDependence6687)
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
					"bucket":     "${alicloud_oss_bucket.CreateBucket.bucket}",
					"style_name": "style-365",
					"content":    "image/resize,p_75,w_75",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bucket":     CHECKSET,
						"style_name": CHECKSET,
						"content":    "image/resize,p_75,w_75",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"content": "image/resize,p_75,w_70",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"content": "image/resize,p_75,w_70",
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

// Test Oss BucketStyle. <<< Resource test cases, automatically generated.
