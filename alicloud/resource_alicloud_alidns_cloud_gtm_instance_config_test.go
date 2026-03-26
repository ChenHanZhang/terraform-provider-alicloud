// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Alidns CloudGtmInstanceConfig. >>> Resource test cases, automatically generated.
// Case resourceCase_20260323_w4vWin 12679
func TestAccAliCloudAlidnsCloudGtmInstanceConfig_basic12679(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alidns_cloud_gtm_instance_config.default"
	ra := resourceAttrInit(resourceId, AlicloudAlidnsCloudGtmInstanceConfigMap12679)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AlidnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAlidnsCloudGtmInstanceConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalidns%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAlidnsCloudGtmInstanceConfigBasicDependence12679)
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
					"address_pool_lb_strategy": "round_robin",
					"schedule_rr_type":         "CNAME",
					"schedule_zone_name":       "tianxuan.top",
					"enable_status":            "disable",
					"charge_type":              "postpay",
					"schedule_host_name":       "www",
					"schedule_zone_mode":       "custom",
					"ttl":                      "600",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_pool_lb_strategy": "round_robin",
						"schedule_rr_type":         "CNAME",
						"schedule_zone_name":       "tianxuan.top",
						"enable_status":            "disable",
						"charge_type":              "postpay",
						"schedule_host_name":       "www",
						"schedule_zone_mode":       "custom",
						"ttl":                      "600",
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
				ImportStateVerifyIgnore: []string{"charge_type"},
			},
		},
	})
}

var AlicloudAlidnsCloudGtmInstanceConfigMap12679 = map[string]string{
	"config_id": CHECKSET,
}

func AlicloudAlidnsCloudGtmInstanceConfigBasicDependence12679(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260323_zV1Ijx 12682
func TestAccAliCloudAlidnsCloudGtmInstanceConfig_basic12682(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alidns_cloud_gtm_instance_config.default"
	ra := resourceAttrInit(resourceId, AlicloudAlidnsCloudGtmInstanceConfigMap12682)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AlidnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAlidnsCloudGtmInstanceConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalidns%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAlidnsCloudGtmInstanceConfigBasicDependence12682)
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
					"address_pool_lb_strategy": "round_robin",
					"schedule_rr_type":         "AAAA",
					"schedule_zone_name":       "tianxuan.top",
					"enable_status":            "disable",
					"charge_type":              "postpay",
					"schedule_host_name":       "www",
					"schedule_zone_mode":       "custom",
					"ttl":                      "600",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_pool_lb_strategy": "round_robin",
						"schedule_rr_type":         "AAAA",
						"schedule_zone_name":       "tianxuan.top",
						"enable_status":            "disable",
						"charge_type":              "postpay",
						"schedule_host_name":       "www",
						"schedule_zone_mode":       "custom",
						"ttl":                      "600",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_pool_lb_strategy": "weight",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_pool_lb_strategy": "weight",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"charge_type"},
			},
		},
	})
}

var AlicloudAlidnsCloudGtmInstanceConfigMap12682 = map[string]string{
	"config_id": CHECKSET,
}

func AlicloudAlidnsCloudGtmInstanceConfigBasicDependence12682(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260320_bH90dh 12689
func TestAccAliCloudAlidnsCloudGtmInstanceConfig_basic12689(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alidns_cloud_gtm_instance_config.default"
	ra := resourceAttrInit(resourceId, AlicloudAlidnsCloudGtmInstanceConfigMap12689)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AlidnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAlidnsCloudGtmInstanceConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalidns%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAlidnsCloudGtmInstanceConfigBasicDependence12689)
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
					"address_pool_lb_strategy":  "sequence",
					"schedule_rr_type":          "A",
					"schedule_zone_name":        "tianxuan.top",
					"instance_id":               "gtm-cn-4l64pdj9901",
					"enable_status":             "disable",
					"charge_type":               "prepay",
					"schedule_host_name":        "example",
					"schedule_zone_mode":        "custom",
					"ttl":                       "600",
					"sequence_lb_strategy_mode": "preemptive",
					"remark":                    "remark",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_pool_lb_strategy":  "sequence",
						"schedule_rr_type":          "A",
						"schedule_zone_name":        "tianxuan.top",
						"instance_id":               "gtm-cn-4l64pdj9901",
						"enable_status":             "disable",
						"charge_type":               "prepay",
						"schedule_host_name":        "example",
						"schedule_zone_mode":        "custom",
						"ttl":                       "600",
						"sequence_lb_strategy_mode": "preemptive",
						"remark":                    "remark",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_status":      "enable",
					"schedule_host_name": "example-2",
					"ttl":                "60",
					"remark":             "add-test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_status":      "enable",
						"schedule_host_name": "example-2",
						"ttl":                "60",
						"remark":             "add-test",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"charge_type"},
			},
		},
	})
}

var AlicloudAlidnsCloudGtmInstanceConfigMap12689 = map[string]string{
	"config_id": CHECKSET,
}

func AlicloudAlidnsCloudGtmInstanceConfigBasicDependence12689(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Alidns CloudGtmInstanceConfig. <<< Resource test cases, automatically generated.
