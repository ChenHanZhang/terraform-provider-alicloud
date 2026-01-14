package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ThreatDetection ClientFileProtect. >>> Resource test cases, automatically generated.
// Case 核心文件规则Platformlinux 9094
func TestAccAliCloudThreatDetectionClientFileProtect_basic9094(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_client_file_protect.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionClientFileProtectMap9094)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionClientFileProtect")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionClientFileProtectBasicDependence9094)
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
					"status": "0",
					"file_paths": []string{
						"/usr/local"},
					"file_ops": []string{
						"CREATE"},
					"rule_action": "pass",
					"proc_paths": []string{
						"/usr/local"},
					"alert_level": "0",
					"switch_id":   "FILE_PROTECT_RULE_SWITCH_TYPE_1693474122929",
					"rule_name":   "ruleTest2_685",
					"platform":    "linux",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":       "0",
						"file_paths.#": "1",
						"file_ops.#":   "1",
						"rule_action":  "pass",
						"proc_paths.#": "1",
						"alert_level":  "0",
						"switch_id":    "FILE_PROTECT_RULE_SWITCH_TYPE_1693474122929",
						"rule_name":    CHECKSET,
						"platform":     "linux",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "1",
					"file_paths": []string{
						"/tmp", "/tmp1", "/tmp2"},
					"file_ops": []string{
						"CHMOD", "CREATE", "UPDATE"},
					"rule_action": "alert",
					"proc_paths": []string{
						"/tmp", "/tmp2", "/tmp3"},
					"alert_level": "1",
					"rule_name":   "ruleTest1_353",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":       "1",
						"file_paths.#": "3",
						"file_ops.#":   "3",
						"rule_action":  "alert",
						"proc_paths.#": "3",
						"alert_level":  "1",
						"rule_name":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "0",
					"file_paths": []string{
						"/tmp21"},
					"file_ops": []string{
						"DELETE"},
					"rule_action": "pass",
					"proc_paths": []string{
						"/tmp12"},
					"rule_name": "ruleTest_269",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":       "0",
						"file_paths.#": "1",
						"file_ops.#":   "1",
						"rule_action":  "pass",
						"proc_paths.#": "1",
						"rule_name":    CHECKSET,
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

var AlicloudThreatDetectionClientFileProtectMap9094 = map[string]string{}

func AlicloudThreatDetectionClientFileProtectBasicDependence9094(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 核心文件规则Platformwindows 9095
func TestAccAliCloudThreatDetectionClientFileProtect_basic9095(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_threat_detection_client_file_protect.default"
	ra := resourceAttrInit(resourceId, AlicloudThreatDetectionClientFileProtectMap9095)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ThreatDetectionServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeThreatDetectionClientFileProtect")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccthreatdetection%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudThreatDetectionClientFileProtectBasicDependence9095)
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
					"status": "0",
					"file_paths": []string{
						"/usr/local"},
					"file_ops": []string{
						"CREATE"},
					"rule_action": "pass",
					"proc_paths": []string{
						"/usr/local"},
					"alert_level": "0",
					"switch_id":   "FILE_PROTECT_RULE_SWITCH_TYPE_1693474122929",
					"rule_name":   "ruleTest2_314",
					"platform":    "windows",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":       "0",
						"file_paths.#": "1",
						"file_ops.#":   "1",
						"rule_action":  "pass",
						"proc_paths.#": "1",
						"alert_level":  "0",
						"switch_id":    "FILE_PROTECT_RULE_SWITCH_TYPE_1693474122929",
						"rule_name":    CHECKSET,
						"platform":     "windows",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "1",
					"file_paths": []string{
						"/tmp", "/tmp1", "/tmp2"},
					"file_ops": []string{
						"CHMOD", "CREATE", "UPDATE"},
					"rule_action": "alert",
					"proc_paths": []string{
						"/tmp", "/tmp2", "/tmp3"},
					"alert_level": "1",
					"rule_name":   "ruleTest1_749",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":       "1",
						"file_paths.#": "3",
						"file_ops.#":   "3",
						"rule_action":  "alert",
						"proc_paths.#": "3",
						"alert_level":  "1",
						"rule_name":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "0",
					"file_paths": []string{
						"/tmp21"},
					"file_ops": []string{
						"DELETE"},
					"rule_action": "pass",
					"proc_paths": []string{
						"/tmp12"},
					"rule_name": "ruleTest_446",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":       "0",
						"file_paths.#": "1",
						"file_ops.#":   "1",
						"rule_action":  "pass",
						"proc_paths.#": "1",
						"rule_name":    CHECKSET,
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

var AlicloudThreatDetectionClientFileProtectMap9095 = map[string]string{}

func AlicloudThreatDetectionClientFileProtectBasicDependence9095(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test ThreatDetection ClientFileProtect. <<< Resource test cases, automatically generated.
