// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cr ArtifactLifecycleRule. >>> Resource test cases, automatically generated.
// Case 保留策略生命周期 5221
func TestAccAliCloudCrArtifactLifecycleRule_basic5221(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_artifact_lifecycle_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrArtifactLifecycleRuleMap5221)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrArtifactLifecycleRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrArtifactLifecycleRuleBasicDependence5221)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"auto":                "false",
					"retention_tag_count": "30",
					"scope":               "INSTANCE",
					"instance_id":         "${alicloud_cr_ee_instance.defaultnKIyBE.id}",
					"tag_regexp":          " ",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto":                "false",
						"retention_tag_count": "30",
						"scope":               "INSTANCE",
						"instance_id":         CHECKSET,
						"tag_regexp":          " ",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_tag_count": "31",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_tag_count": "31",
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

var AlicloudCrArtifactLifecycleRuleMap5221 = map[string]string{
	"create_time":                CHECKSET,
	"artifact_lifecycle_rule_id": CHECKSET,
}

func AlicloudCrArtifactLifecycleRuleBasicDependence5221(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cr_ee_instance" "defaultnKIyBE" {
  instance_name  = "test-artifact-life-rule-158"
  renewal_status = "ManualRenewal"
  image_scanner  = "ACR"
  period         = "1"
  instance_type  = "Basic"
}


`, name)
}

// Case 保留策略生命周期_换账号可用_副本1709104286506 6046
func TestAccAliCloudCrArtifactLifecycleRule_basic6046(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_artifact_lifecycle_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrArtifactLifecycleRuleMap6046)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrArtifactLifecycleRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrArtifactLifecycleRuleBasicDependence6046)
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
					"auto":                "true",
					"retention_tag_count": "30",
					"scope":               "REPO",
					"instance_id":         "${alicloud_cr_ee_instance.default2Rk8gT.id}",
					"tag_regexp":          " .*",
					"namespace_name":      "${alicloud_cr_namespace.defaultevafKF.namespace_name}",
					"repo_name":           "${alicloud_cr_repository.defaultKunw72.repository_name}",
					"schedule_time":       "WEEK",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto":                "true",
						"retention_tag_count": "30",
						"scope":               "REPO",
						"instance_id":         CHECKSET,
						"tag_regexp":          " .*",
						"namespace_name":      CHECKSET,
						"repo_name":           CHECKSET,
						"schedule_time":       "WEEK",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_tag_count": "31",
					"scope":               "INSTANCE",
					"tag_regexp":          "release-v.*",
					"schedule_time":       "MONTH",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_tag_count": "31",
						"scope":               "INSTANCE",
						"tag_regexp":          "release-v.*",
						"schedule_time":       "MONTH",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_tag_count": "28",
					"scope":               "REPO",
					"namespace_name":      "${alicloud_cr_namespace.defaultGPiaHQ.namespace_name}",
					"repo_name":           "${alicloud_cr_repository.defaultkCdOJ6.repository_name}",
					"schedule_time":       "WEEK",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_tag_count": "28",
						"scope":               "REPO",
						"namespace_name":      CHECKSET,
						"repo_name":           CHECKSET,
						"schedule_time":       "WEEK",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto":                "false",
					"retention_tag_count": "30",
					"tag_regexp":          " .*",
					"namespace_name":      "${alicloud_cr_namespace.defaultevafKF.namespace_name}",
					"repo_name":           "${alicloud_cr_repository.defaultKunw72.repository_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto":                "false",
						"retention_tag_count": "30",
						"tag_regexp":          " .*",
						"namespace_name":      CHECKSET,
						"repo_name":           CHECKSET,
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

var AlicloudCrArtifactLifecycleRuleMap6046 = map[string]string{
	"create_time":                CHECKSET,
	"artifact_lifecycle_rule_id": CHECKSET,
}

func AlicloudCrArtifactLifecycleRuleBasicDependence6046(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cr_ee_instance" "default2Rk8gT" {
  instance_name  = "auto-test-instance2"
  renewal_status = "ManualRenewal"
  image_scanner  = "ACR"
  period         = "1"
  instance_type  = "Basic"
}

resource "alicloud_cr_namespace" "defaultevafKF" {
  namespace_name   = "auto-test-ns2"
  instance_id      = alicloud_cr_ee_instance.default2Rk8gT.id
  auto_create_repo = false
}

resource "alicloud_cr_repository" "defaultKunw72" {
  repo_namespace_name = alicloud_cr_namespace.defaultevafKF.namespace_name
  tag_immutability    = false
  repo_type           = "PUBLIC"
  instance_id         = alicloud_cr_ee_instance.default2Rk8gT.id
  summary             = "dd"
  repository_name     = "auto-test-repo"
  detail              = "ss"
}

resource "alicloud_cr_namespace" "defaultGPiaHQ" {
  namespace_name   = "auto-test-ns3"
  instance_id      = alicloud_cr_ee_instance.default2Rk8gT.id
  auto_create_repo = false
}

resource "alicloud_cr_repository" "defaultkCdOJ6" {
  repo_namespace_name = alicloud_cr_namespace.defaultGPiaHQ.namespace_name
  tag_immutability    = false
  repo_type           = "PUBLIC"
  instance_id         = alicloud_cr_ee_instance.default2Rk8gT.id
  summary             = "dddd"
  repository_name     = "auto-test-repo3"
  detail              = "ddadaf"
}


`, name)
}

// Case 保留策略生命周期_使用固定Instance-SUCC 5605
func TestAccAliCloudCrArtifactLifecycleRule_basic5605(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cr_artifact_lifecycle_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCrArtifactLifecycleRuleMap5605)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CrServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCrArtifactLifecycleRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccr%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCrArtifactLifecycleRuleBasicDependence5605)
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
					"auto":                "true",
					"retention_tag_count": "30",
					"scope":               "REPO",
					"instance_id":         "${alicloud_cr_ee_instance.default2Rk8gT.id}",
					"tag_regexp":          " .*",
					"namespace_name":      "${alicloud_cr_namespace.defaultevafKF.namespace_name}",
					"repo_name":           "${alicloud_cr_repository.defaultKunw72.repository_name}",
					"schedule_time":       "WEEK",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto":                "true",
						"retention_tag_count": "30",
						"scope":               "REPO",
						"instance_id":         CHECKSET,
						"tag_regexp":          " .*",
						"namespace_name":      CHECKSET,
						"repo_name":           CHECKSET,
						"schedule_time":       "WEEK",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_tag_count": "31",
					"scope":               "INSTANCE",
					"tag_regexp":          "release-v.*",
					"schedule_time":       "MONTH",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_tag_count": "31",
						"scope":               "INSTANCE",
						"tag_regexp":          "release-v.*",
						"schedule_time":       "MONTH",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_tag_count": "28",
					"scope":               "REPO",
					"namespace_name":      "${alicloud_cr_namespace.defaultGPiaHQ.namespace_name}",
					"repo_name":           "${alicloud_cr_repository.defaultkCdOJ6.repository_name}",
					"schedule_time":       "WEEK",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_tag_count": "28",
						"scope":               "REPO",
						"namespace_name":      CHECKSET,
						"repo_name":           CHECKSET,
						"schedule_time":       "WEEK",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto":                "false",
					"retention_tag_count": "30",
					"tag_regexp":          " .*",
					"namespace_name":      "${alicloud_cr_namespace.defaultevafKF.namespace_name}",
					"repo_name":           "${alicloud_cr_repository.defaultKunw72.repository_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto":                "false",
						"retention_tag_count": "30",
						"tag_regexp":          " .*",
						"namespace_name":      CHECKSET,
						"repo_name":           CHECKSET,
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

var AlicloudCrArtifactLifecycleRuleMap5605 = map[string]string{
	"create_time":                CHECKSET,
	"artifact_lifecycle_rule_id": CHECKSET,
}

func AlicloudCrArtifactLifecycleRuleBasicDependence5605(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cr_ee_instance" "default2Rk8gT" {
  instance_name  = "auto-test-instance2"
  renewal_status = "ManualRenewal"
  image_scanner  = "ACR"
  period         = "1"
  instance_type  = "Basic"
}

resource "alicloud_cr_namespace" "defaultevafKF" {
  namespace_name   = "auto-test-ns2"
  instance_id      = alicloud_cr_ee_instance.default2Rk8gT.id
  auto_create_repo = false
}

resource "alicloud_cr_repository" "defaultKunw72" {
  repo_namespace_name = alicloud_cr_namespace.defaultevafKF.namespace_name
  tag_immutability    = false
  repo_type           = "PUBLIC"
  instance_id         = alicloud_cr_ee_instance.default2Rk8gT.id
  summary             = "dd"
  repository_name     = "auto-test-repo"
  detail              = "ss"
}

resource "alicloud_cr_namespace" "defaultGPiaHQ" {
  namespace_name   = "auto-test-ns3"
  instance_id      = alicloud_cr_ee_instance.default2Rk8gT.id
  auto_create_repo = false
}

resource "alicloud_cr_repository" "defaultkCdOJ6" {
  repo_namespace_name = alicloud_cr_namespace.defaultGPiaHQ.namespace_name
  tag_immutability    = false
  repo_type           = "PUBLIC"
  instance_id         = alicloud_cr_ee_instance.default2Rk8gT.id
  summary             = "dddd"
  repository_name     = "auto-test-repo3"
  detail              = "ddadaf"
}


`, name)
}

// Test Cr ArtifactLifecycleRule. <<< Resource test cases, automatically generated.
