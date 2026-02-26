package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ThreatDetection WebLockConfig. >>> Resource test cases, automatically generated.
// Case WebLockConfig20251218 12128
func TestAccAliCloudThreatDetectionWebLockConfig_basic12128(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_web_lock_config.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionWebLockConfigMap12128)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionWebLockConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionWebLockConfigBasicDependence12128)
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
					"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif;png",
					"exclusive_file":      "txt.txt",
					"uuid":                "${alicloud_threat_detection_web_lock_bind.defaultaAkZPl.id}",
					"exclusive_dir":       "test2",
					"defence_mode":        "block",
					"mode":                "whitelist",
					"local_backup_dir":    "/usr/local/bak1/",
					"exclusive_file_type": "txt",
					"dir":                 "/opt/test6/",
					"lang":                "en",
					"inclusive_file":      "t4.txt",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif;png",
						"exclusive_file":      "txt.txt",
						"uuid":                CHECKSET,
						"exclusive_dir":       "test2",
						"defence_mode":        "block",
						"mode":                "whitelist",
						"local_backup_dir":    "/usr/local/bak1/",
						"exclusive_file_type": "txt",
						"dir":                 "/opt/test6/",
						"lang":                "en",
						"inclusive_file":      "t4.txt",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif",
					"defence_mode":        "audit",
					"local_backup_dir":    "/usr/local/bak3/",
					"dir":                 "/opt/test8/",
					"lang":                "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif",
						"defence_mode":        "audit",
						"local_backup_dir":    "/usr/local/bak3/",
						"dir":                 "/opt/test8/",
						"lang":                "zh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"exclusive_file":      "txt5.txt",
					"exclusive_dir":       "test1",
					"defence_mode":        "block",
					"mode":                "blacklist",
					"local_backup_dir":    "/usr/local/bak3",
					"exclusive_file_type": "png",
					"inclusive_file":      "t5.txt",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"exclusive_file":      "txt5.txt",
						"exclusive_dir":       "test1",
						"defence_mode":        "block",
						"mode":                "blacklist",
						"local_backup_dir":    "/usr/local/bak3",
						"exclusive_file_type": "png",
						"inclusive_file":      "t5.txt",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudThreatDetectionWebLockConfigMap12128 = map[string]string{
	"config_id": CHECKSET,
}

func AlicloudThreatDetectionWebLockConfigBasicDependence12128(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_threat_detection_asset" "defaultgUQ4PP" {
}

resource "alicloud_threat_detection_web_lock_bind" "defaultaAkZPl" {
  inclusive_file_type = "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif;png"
  status              = "off"
  exclusive_file      = "txt.txt"
  uuid                = alicloud_threat_detection_asset.defaultgUQ4PP.id
  exclusive_dir       = "test2"
  defence_mode        = "block"
  mode                = "whitelist"
  local_backup_dir    = "/usr/local/bak/"
  exclusive_file_type = "txt"
  dir                 = "/opt/test/"
}


`, name)
}

// Case WebLockConfig 9102
func TestAccAliCloudThreatDetectionWebLockConfig_basic9102(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_web_lock_config.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionWebLockConfigMap9102)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionWebLockConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionWebLockConfigBasicDependence9102)
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
					"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif;png",
					"exclusive_file":      "txt.txt",
					"uuid":                "${alicloud_threat_detection_web_lock_bind.defaultaAkZPl.id}",
					"exclusive_dir":       "test2",
					"defence_mode":        "block",
					"mode":                "whitelist",
					"local_backup_dir":    "/usr/local/bak1/",
					"exclusive_file_type": "txt",
					"dir":                 "/opt/test6/",
					"lang":                "en",
					"inclusive_file":      "t4.txt",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif;png",
						"exclusive_file":      "txt.txt",
						"uuid":                CHECKSET,
						"exclusive_dir":       "test2",
						"defence_mode":        "block",
						"mode":                "whitelist",
						"local_backup_dir":    "/usr/local/bak1/",
						"exclusive_file_type": "txt",
						"dir":                 "/opt/test6/",
						"lang":                "en",
						"inclusive_file":      "t4.txt",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif",
					"defence_mode":        "audit",
					"local_backup_dir":    "/usr/local/bak3/",
					"dir":                 "/opt/test8/",
					"lang":                "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif",
						"defence_mode":        "audit",
						"local_backup_dir":    "/usr/local/bak3/",
						"dir":                 "/opt/test8/",
						"lang":                "zh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"exclusive_file":      "txt5.txt",
					"exclusive_dir":       "test1",
					"defence_mode":        "block",
					"mode":                "blacklist",
					"local_backup_dir":    "/usr/local/bak3",
					"exclusive_file_type": "png",
					"inclusive_file":      "t5.txt",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"exclusive_file":      "txt5.txt",
						"exclusive_dir":       "test1",
						"defence_mode":        "block",
						"mode":                "blacklist",
						"local_backup_dir":    "/usr/local/bak3",
						"exclusive_file_type": "png",
						"inclusive_file":      "t5.txt",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudThreatDetectionWebLockConfigMap9102 = map[string]string{
	"config_id": CHECKSET,
}

func AlicloudThreatDetectionWebLockConfigBasicDependence9102(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_threat_detection_asset" "defaultgUQ4PP" {
}

resource "alicloud_threat_detection_web_lock_bind" "defaultaAkZPl" {
  inclusive_file_type = "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg;gif;png"
  status              = "off"
  exclusive_file      = "txt.txt"
  uuid                = alicloud_threat_detection_asset.defaultgUQ4PP.id
  exclusive_dir       = "test2"
  defence_mode        = "block"
  mode                = "whitelist"
  local_backup_dir    = "/usr/local/bak/"
  exclusive_file_type = "txt"
  dir                 = "/opt/test/"
}


`, name)
}

// Test ThreatDetection WebLockConfig. <<< Resource test cases, automatically generated.

// Case 1
func TestAccAlicloudThreatDetectionWebLockConfig_basic1875(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_web_lock_config.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionWebLockConfigMap1875)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SasService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionWebLockConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sThreatDetectionWebLockConfig%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionWebLockConfigBasicDependence1875)
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
					"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg",
					"uuid":                "${data.alicloud_threat_detection_assets.default.ids.0}",
					"mode":                "whitelist",
					"local_backup_dir":    "/usr/local/aegis/bak",
					"dir":                 "/tmp/",
					"defence_mode":        "audit",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"inclusive_file_type": "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg",
						"uuid":                CHECKSET,
						"mode":                "whitelist",
						"local_backup_dir":    "/usr/local/aegis/bak",
						"dir":                 "/tmp/",
						"defence_mode":        "audit",
					}),
				),
			}, {
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudThreatDetectionWebLockConfigMap1875 = map[string]string{}

func AlicloudThreatDetectionWebLockConfigBasicDependence1875(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_threat_detection_assets" "default" {
  machine_types = "ecs"
}

`, name)
}
