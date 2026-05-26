// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cr ScanRule. >>> Resource test cases, automatically generated.
// Case ScanRule-1_pl_副本1763544286754 11851
func TestAccAliCloudCrScanRule_basic11851(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_scan_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrScanRuleMap11851)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrScanRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrScanRuleBasicDependence11851)
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
					"repo_tag_filter_pattern": ".*",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"scan_type":               "VUL",
					"rule_name":               "161",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"bb", "cc", "dd", "ee"},
					"instance_id": "${alicloud_cr_ee_instance.default2Aqoce.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": ".*",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"scan_type":               "VUL",
						"rule_name":               CHECKSET,
						"namespaces.#":            "1",
						"repo_names.#":            "4",
						"instance_id":             CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "a",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "AUTO",
					"rule_name":               "aa",
					"namespaces": []string{
						"aa", "bb", "cc", "dd"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "a",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "AUTO",
						"rule_name":               "aa",
						"namespaces.#":            "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "bb",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"rule_name":               "bb",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"aa"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "bb",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"rule_name":               "bb",
						"namespaces.#":            "1",
						"repo_names.#":            "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "cc",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "dd",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "cc",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "dd",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "dd",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "MANUAL",
					"rule_name":               "ff",
					"namespaces": []string{
						"aa", "bb"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "dd",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "MANUAL",
						"rule_name":               "ff",
						"namespaces.#":            "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "aa",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "gg",
					"namespaces":              []string{},
					"repo_names":              []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "aa",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "gg",
						"namespaces.#":            "0",
						"repo_names.#":            "0",
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

var AlicloudCrScanRuleMap11851 = map[string]string{
	"scan_rule_id": CHECKSET,
	"create_time":  CHECKSET,
	"update_time":  CHECKSET,
}

func AlicloudCrScanRuleBasicDependence11851(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cr_ee_instance" "default2Aqoce" {
  default_oss_bucket = "false"
  renewal_status     = "ManualRenewal"
  period             = "1"
  instance_name      = "pl-test-2"
  payment_type       = "Subscription"
  instance_type      = "Basic"
}


`, name)
}

// Case ScanRule-1_pl 11745
func TestAccAliCloudCrScanRule_basic11745(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_scan_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrScanRuleMap11745)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrScanRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrScanRuleBasicDependence11745)
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
					"repo_tag_filter_pattern": ".*",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"scan_type":               "VUL",
					"rule_name":               "254",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"bb", "cc", "dd", "ee"},
					"instance_id": "${alicloud_cr_ee_instance.default2Aqoce.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": ".*",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"scan_type":               "VUL",
						"rule_name":               CHECKSET,
						"namespaces.#":            "1",
						"repo_names.#":            "4",
						"instance_id":             CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "a",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "AUTO",
					"rule_name":               "aa",
					"namespaces": []string{
						"aa", "bb", "cc", "dd"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "a",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "AUTO",
						"rule_name":               "aa",
						"namespaces.#":            "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "bb",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"rule_name":               "bb",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"aa"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "bb",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"rule_name":               "bb",
						"namespaces.#":            "1",
						"repo_names.#":            "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "cc",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "dd",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "cc",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "dd",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "dd",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "MANUAL",
					"rule_name":               "ff",
					"namespaces": []string{
						"aa", "bb"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "dd",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "MANUAL",
						"rule_name":               "ff",
						"namespaces.#":            "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "aa",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "gg",
					"namespaces":              []string{},
					"repo_names":              []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "aa",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "gg",
						"namespaces.#":            "0",
						"repo_names.#":            "0",
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

var AlicloudCrScanRuleMap11745 = map[string]string{
	"scan_rule_id": CHECKSET,
	"create_time":  CHECKSET,
	"update_time":  CHECKSET,
}

func AlicloudCrScanRuleBasicDependence11745(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cr_ee_instance" "default2Aqoce" {
  default_oss_bucket = "false"
  renewal_status     = "ManualRenewal"
  period             = "1"
  instance_name      = "pl-test-2"
  payment_type       = "Subscription"
  instance_type      = "Basic"
}


`, name)
}

// Case ScanRule-1_pl_副本1761702766466 11751
func TestAccAliCloudCrScanRule_basic11751(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_scan_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrScanRuleMap11751)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrScanRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrScanRuleBasicDependence11751)
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
					"repo_tag_filter_pattern": ".*",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"scan_type":               "VUL",
					"rule_name":               "670",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"bb", "cc", "dd", "ee"},
					"instance_id": "${alicloud_cr_ee_instance.defaultT1WvNH.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": ".*",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"scan_type":               "VUL",
						"rule_name":               CHECKSET,
						"namespaces.#":            "1",
						"repo_names.#":            "4",
						"instance_id":             CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "a",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "AUTO",
					"rule_name":               "aa",
					"namespaces": []string{
						"aa", "bb", "cc", "dd"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "a",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "AUTO",
						"rule_name":               "aa",
						"namespaces.#":            "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "bb",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"rule_name":               "bb",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"aa"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "bb",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"rule_name":               "bb",
						"namespaces.#":            "1",
						"repo_names.#":            "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "cc",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "dd",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "cc",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "dd",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "dd",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "MANUAL",
					"rule_name":               "ff",
					"namespaces": []string{
						"aa", "bb"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "dd",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "MANUAL",
						"rule_name":               "ff",
						"namespaces.#":            "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "aa",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "gg",
					"namespaces":              []string{},
					"repo_names":              []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "aa",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "gg",
						"namespaces.#":            "0",
						"repo_names.#":            "0",
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

var AlicloudCrScanRuleMap11751 = map[string]string{
	"scan_rule_id": CHECKSET,
	"create_time":  CHECKSET,
	"update_time":  CHECKSET,
}

func AlicloudCrScanRuleBasicDependence11751(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "common_region" {
  default = "cn-hangzhou"
}

resource "alicloud_cr_ee_instance" "defaultT1WvNH" {
  default_oss_bucket = "true"
  renewal_status     = "ManualRenewal"
  period             = "1"
  instance_name      = "pl-test-scan-rule"
  payment_type       = "Subscription"
  instance_type      = "Basic"
}


`, name)
}

// Case ScanRule-1 11563
func TestAccAliCloudCrScanRule_basic11563(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_scan_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrScanRuleMap11563)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrScanRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrScanRuleBasicDependence11563)
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
					"repo_tag_filter_pattern": ".*",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"scan_type":               "VUL",
					"instance_id":             "${alicloud_cr_ee_instance.defaultST7wU7.id}",
					"rule_name":               "210",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"bb", "cc", "dd", "ee"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": ".*",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"scan_type":               "VUL",
						"instance_id":             CHECKSET,
						"rule_name":               CHECKSET,
						"namespaces.#":            "1",
						"repo_names.#":            "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "a",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "AUTO",
					"rule_name":               "aa",
					"namespaces": []string{
						"aa", "bb", "cc", "dd"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "a",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "AUTO",
						"rule_name":               "aa",
						"namespaces.#":            "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "bb",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"rule_name":               "bb",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"aa"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "bb",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"rule_name":               "bb",
						"namespaces.#":            "1",
						"repo_names.#":            "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "cc",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "dd",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "cc",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "dd",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "dd",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "MANUAL",
					"rule_name":               "ff",
					"namespaces": []string{
						"aa", "bb"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "dd",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "MANUAL",
						"rule_name":               "ff",
						"namespaces.#":            "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "aa",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "gg",
					"namespaces":              []string{},
					"repo_names":              []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "aa",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "gg",
						"namespaces.#":            "0",
						"repo_names.#":            "0",
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

var AlicloudCrScanRuleMap11563 = map[string]string{
	"scan_rule_id": CHECKSET,
	"create_time":  CHECKSET,
	"update_time":  CHECKSET,
}

func AlicloudCrScanRuleBasicDependence11563(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cr_ee_instance" "defaultST7wU7" {
  instance_name      = "test-654"
  period             = "1"
  renewal_status     = "ManualRenewal"
  image_scanner      = "ACR"
  instance_type      = "Basic"
  default_oss_bucket = "false"
  payment_type       = "Subscription"
}


`, name)
}

// Case c s 11592
func TestAccAliCloudCrScanRule_basic11592(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_scan_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrScanRuleMap11592)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrScanRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrScanRuleBasicDependence11592)
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
					"repo_tag_filter_pattern": ".*",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"scan_type":               "VUL",
					"instance_id":             "${alicloud_cr_ee_instance.defaultST7wU7.id}",
					"rule_name":               "12",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"bb", "cc", "dd", "ee"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": ".*",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"scan_type":               "VUL",
						"instance_id":             CHECKSET,
						"rule_name":               CHECKSET,
						"namespaces.#":            "1",
						"repo_names.#":            "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "a",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "AUTO",
					"rule_name":               "aa",
					"namespaces": []string{
						"aa", "bb", "cc", "dd"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "a",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "AUTO",
						"rule_name":               "aa",
						"namespaces.#":            "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "bb",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"rule_name":               "bb",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"aa"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "bb",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"rule_name":               "bb",
						"namespaces.#":            "1",
						"repo_names.#":            "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "cc",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "dd",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "cc",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "dd",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "dd",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "MANUAL",
					"rule_name":               "ff",
					"namespaces": []string{
						"aa", "bb"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "dd",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "MANUAL",
						"rule_name":               "ff",
						"namespaces.#":            "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "aa",
					"scan_scope":              "INSTANCE",
					"trigger_type":            "AUTO",
					"rule_name":               "gg",
					"namespaces":              []string{},
					"repo_names":              []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "aa",
						"scan_scope":              "INSTANCE",
						"trigger_type":            "AUTO",
						"rule_name":               "gg",
						"namespaces.#":            "0",
						"repo_names.#":            "0",
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

var AlicloudCrScanRuleMap11592 = map[string]string{
	"scan_rule_id": CHECKSET,
	"create_time":  CHECKSET,
	"update_time":  CHECKSET,
}

func AlicloudCrScanRuleBasicDependence11592(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cr_ee_instance" "defaultST7wU7" {
  instance_name      = "test-54"
  period             = "1"
  renewal_status     = "ManualRenewal"
  image_scanner      = "ACR"
  instance_type      = "Basic"
  default_oss_bucket = "false"
  payment_type       = "Subscription"
}


`, name)
}

// Case pl_test 11565
func TestAccAliCloudCrScanRule_basic11565(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_scan_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrScanRuleMap11565)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrScanRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrScanRuleBasicDependence11565)
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
					"repo_tag_filter_pattern": ".*",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"scan_type":               "VUL",
					"instance_id":             "cri-k6vwi42u6gfehg1o",
					"rule_name":               "293",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": ".*",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"scan_type":               "VUL",
						"instance_id":             "cri-k6vwi42u6gfehg1o",
						"rule_name":               CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "a",
					"scan_scope":              "NAMESPACE",
					"trigger_type":            "AUTO",
					"rule_name":               "aa",
					"namespaces": []string{
						"aa", "bb"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "a",
						"scan_scope":              "NAMESPACE",
						"trigger_type":            "AUTO",
						"rule_name":               "aa",
						"namespaces.#":            "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repo_tag_filter_pattern": "bb",
					"scan_scope":              "REPO",
					"trigger_type":            "MANUAL",
					"rule_name":               "bb",
					"namespaces": []string{
						"aa"},
					"repo_names": []string{
						"aa", "bb"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repo_tag_filter_pattern": "bb",
						"scan_scope":              "REPO",
						"trigger_type":            "MANUAL",
						"rule_name":               "bb",
						"namespaces.#":            "1",
						"repo_names.#":            "2",
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

var AlicloudCrScanRuleMap11565 = map[string]string{
	"scan_rule_id": CHECKSET,
	"create_time":  CHECKSET,
	"update_time":  CHECKSET,
}

func AlicloudCrScanRuleBasicDependence11565(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Cr ScanRule. <<< Resource test cases, automatically generated.
