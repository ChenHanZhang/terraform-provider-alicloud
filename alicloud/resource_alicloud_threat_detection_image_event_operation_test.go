package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ThreatDetection ImageEventOperation. >>> Resource test cases, automatically generated.
// Case ImageEventOperation_buildRisk_20250214 10231
func TestAccAliCloudThreatDetectionImageEventOperation_basic10231(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_image_event_operation.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionImageEventOperationMap10231)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionImageEventOperation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionImageEventOperationBasicDependence10231)
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
					"event_type":     "buildRisk",
					"operation_code": "whitelist",
					"event_key":      "huaweicloud_ak",
					"scenarios":      "{\\\"type\\\":\\\"default\\\",\\\"value\\\":\\\"\\\"}",
					"event_name":     "华为AK",
					"conditions":     "[{\\\"condition\\\":\\\"MD5\\\",\\\"type\\\":\\\"equals\\\",\\\"value\\\":\\\"cb1ddb97bad0c443b438e6d013a6de6f\\\"}]",
					"note":           "test",
					"source":         "default",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"event_type":     "buildRisk",
						"operation_code": "whitelist",
						"event_key":      "huaweicloud_ak",
						"scenarios":      CHECKSET,
						"event_name":     "华为AK",
						"conditions":     CHECKSET,
						"note":           "test",
						"source":         "default",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"scenarios":  "{\\\"type\\\":\\\"repo\\\",\\\"value\\\":\\\"test/repo-01\\\"}",
					"conditions": "[{\\\"condition\\\":\\\"MD5\\\",\\\"type\\\":\\\"equals\\\",\\\"value\\\":\\\"0083a31cc0083a31ccf7c10367a6e783e6\\\"}]",
					"note":       "test2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"scenarios":  CHECKSET,
						"conditions": CHECKSET,
						"note":       "test2",
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

var AlicloudThreatDetectionImageEventOperationMap10231 = map[string]string{}

func AlicloudThreatDetectionImageEventOperationBasicDependence10231(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case ImageEventOperation_sensitiveFile_20250214 10229
func TestAccAliCloudThreatDetectionImageEventOperation_basic10229(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_image_event_operation.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionImageEventOperationMap10229)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionImageEventOperation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionImageEventOperationBasicDependence10229)
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
					"event_type":     "sensitiveFile",
					"operation_code": "whitelist",
					"event_key":      "huaweicloud_ak",
					"scenarios":      "{\\\"type\\\":\\\"default\\\",\\\"value\\\":\\\"\\\"}",
					"event_name":     "华为AK",
					"conditions":     "[{\\\"condition\\\":\\\"MD5\\\",\\\"type\\\":\\\"equals\\\",\\\"value\\\":\\\"cb1ddb97bad0c443b438e6d013a6de6f\\\"}]",
					"note":           "test",
					"source":         "agentless",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"event_type":     "sensitiveFile",
						"operation_code": "whitelist",
						"event_key":      "huaweicloud_ak",
						"scenarios":      CHECKSET,
						"event_name":     "华为AK",
						"conditions":     CHECKSET,
						"note":           "test",
						"source":         "agentless",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"scenarios":  "{\\\"type\\\":\\\"repo\\\",\\\"value\\\":\\\"test/repo-01\\\"}",
					"conditions": "[{\\\"condition\\\":\\\"MD5\\\",\\\"type\\\":\\\"equals\\\",\\\"value\\\":\\\"0083a31cc0083a31ccf7c10367a6e783e6\\\"}]",
					"note":       "test2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"scenarios":  CHECKSET,
						"conditions": CHECKSET,
						"note":       "test2",
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

var AlicloudThreatDetectionImageEventOperationMap10229 = map[string]string{}

func AlicloudThreatDetectionImageEventOperationBasicDependence10229(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case ImageEventOperation_maliciousFile_20250214 10230
func TestAccAliCloudThreatDetectionImageEventOperation_basic10230(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_image_event_operation.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionImageEventOperationMap10230)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionImageEventOperation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionImageEventOperationBasicDependence10230)
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
					"event_type":     "maliciousFile",
					"operation_code": "whitelist",
					"event_key":      "huaweicloud_ak",
					"scenarios":      "{\\\"type\\\":\\\"default\\\",\\\"value\\\":\\\"\\\"}",
					"event_name":     "华为AK",
					"conditions":     "[{\\\"condition\\\":\\\"MD5\\\",\\\"type\\\":\\\"equals\\\",\\\"value\\\":\\\"cb1ddb97bad0c443b438e6d013a6de6f\\\"}]",
					"note":           "test",
					"source":         "agentless",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"event_type":     "maliciousFile",
						"operation_code": "whitelist",
						"event_key":      "huaweicloud_ak",
						"scenarios":      CHECKSET,
						"event_name":     "华为AK",
						"conditions":     CHECKSET,
						"note":           "test",
						"source":         "agentless",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"scenarios":  "{\\\"type\\\":\\\"repo\\\",\\\"value\\\":\\\"test/repo-01\\\"}",
					"conditions": "[{\\\"condition\\\":\\\"MD5\\\",\\\"type\\\":\\\"equals\\\",\\\"value\\\":\\\"0083a31cc0083a31ccf7c10367a6e783e6\\\"}]",
					"note":       "test2",
					"source":     "default",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"scenarios":  CHECKSET,
						"conditions": CHECKSET,
						"note":       "test2",
						"source":     "default",
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

var AlicloudThreatDetectionImageEventOperationMap10230 = map[string]string{}

func AlicloudThreatDetectionImageEventOperationBasicDependence10230(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case ImageEventOperation_TestCase_maliciousFile 4600
func TestAccAliCloudThreatDetectionImageEventOperation_basic4600(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_image_event_operation.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionImageEventOperationMap4600)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionImageEventOperation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionImageEventOperationBasicDependence4600)
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
					"event_type":     "maliciousFile",
					"operation_code": "whitelist",
					"event_key":      "huaweicloud_ak",
					"scenarios":      "{\\\"type\\\":\\\"default\\\",\\\"value\\\":\\\"\\\"}",
					"event_name":     "华为AK",
					"conditions":     "[{\\\"condition\\\":\\\"MD5\\\",\\\"type\\\":\\\"equals\\\",\\\"value\\\":\\\"0083a31cc0083a31ccf7c10367a6e783e\\\"}]",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"event_type":     "maliciousFile",
						"operation_code": "whitelist",
						"event_key":      "huaweicloud_ak",
						"scenarios":      CHECKSET,
						"event_name":     "华为AK",
						"conditions":     CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"scenarios": "{\\\"type\\\":\\\"repo\\\",\\\"value\\\":\\\"test/repo-01\\\"}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"scenarios": CHECKSET,
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

var AlicloudThreatDetectionImageEventOperationMap4600 = map[string]string{}

func AlicloudThreatDetectionImageEventOperationBasicDependence4600(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case ImageEventOperation_TestCase_sensitiveFile 4588
func TestAccAliCloudThreatDetectionImageEventOperation_basic4588(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_image_event_operation.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionImageEventOperationMap4588)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionImageEventOperation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionImageEventOperationBasicDependence4588)
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
					"event_type":     "sensitiveFile",
					"operation_code": "whitelist",
					"event_key":      "huaweicloud_ak",
					"scenarios":      "{\\\"type\\\":\\\"default\\\",\\\"value\\\":\\\"\\\"}",
					"event_name":     "华为AK",
					"conditions":     "[{\\\"condition\\\":\\\"MD5\\\",\\\"type\\\":\\\"equals\\\",\\\"value\\\":\\\"0083a31cc0083a31ccf7c10367a6e783e\\\"}]",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"event_type":     "sensitiveFile",
						"operation_code": "whitelist",
						"event_key":      "huaweicloud_ak",
						"scenarios":      CHECKSET,
						"event_name":     "华为AK",
						"conditions":     CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"scenarios": "{\\\"type\\\":\\\"repo\\\",\\\"value\\\":\\\"test/repo-01\\\"}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"scenarios": CHECKSET,
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

var AlicloudThreatDetectionImageEventOperationMap4588 = map[string]string{}

func AlicloudThreatDetectionImageEventOperationBasicDependence4588(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test ThreatDetection ImageEventOperation. <<< Resource test cases, automatically generated.
