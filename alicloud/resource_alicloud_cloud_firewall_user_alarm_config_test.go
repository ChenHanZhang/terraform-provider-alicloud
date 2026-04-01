package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test CloudFirewall UserAlarmConfig. >>> Resource test cases, automatically generated.
// Case 告警测试_Notify_ContactConfig空 12698
func TestAccAliCloudCloudFirewallUserAlarmConfig_basic12698(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_user_alarm_config.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallUserAlarmConfigMap12698)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallUserAlarmConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallUserAlarmConfigBasicDependence12698)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"alarm_config": []map[string]interface{}{
						{
							"alarm_value":    "on",
							"alarm_type":     "bandwidth",
							"alarm_period":   "1",
							"alarm_hour":     "0",
							"alarm_notify":   "0",
							"alarm_week_day": "0",
						},
					},
					"use_default_contact": "1",
					"alarm_lang":          "zh",
					"lang":                "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alarm_config.#":      "1",
						"use_default_contact": CHECKSET,
						"alarm_lang":          "zh",
						"lang":                "zh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang", "use_default_contact"},
			},
		},
	})
}

var AlicloudCloudFirewallUserAlarmConfigMap12698 = map[string]string{}

func AlicloudCloudFirewallUserAlarmConfigBasicDependence12698(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case test 12700
func TestAccAliCloudCloudFirewallUserAlarmConfig_basic12700(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_user_alarm_config.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallUserAlarmConfigMap12700)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallUserAlarmConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallUserAlarmConfigBasicDependence12700)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"use_default_contact": "1",
					"alarm_config": []map[string]interface{}{
						{
							"alarm_value":  "on",
							"alarm_type":   "bandwidth",
							"alarm_period": "1",
							"alarm_hour":   "0",
						},
					},
					"contact_config": []map[string]interface{}{
						{
							"status":       "1",
							"mobile_phone": "18966985698",
							"name":         "test",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"use_default_contact": CHECKSET,
						"alarm_config.#":      "1",
						"contact_config.#":    "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang", "use_default_contact"},
			},
		},
	})
}

var AlicloudCloudFirewallUserAlarmConfigMap12700 = map[string]string{}

func AlicloudCloudFirewallUserAlarmConfigBasicDependence12700(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 告警测试 12506
func TestAccAliCloudCloudFirewallUserAlarmConfig_basic12506(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_user_alarm_config.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallUserAlarmConfigMap12506)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallUserAlarmConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallUserAlarmConfigBasicDependence12506)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"alarm_config": []map[string]interface{}{
						{
							"alarm_value":    "on",
							"alarm_type":     "bandwidth",
							"alarm_period":   "1",
							"alarm_hour":     "0",
							"alarm_notify":   "0",
							"alarm_week_day": "0",
						},
					},
					"use_default_contact": "1",
					"contact_config": []map[string]interface{}{
						{
							"status":       "1",
							"email":        "shanshan.xl@alibaba-inc.com",
							"mobile_phone": "13585310666",
							"name":         "闪闪",
						},
					},
					"alarm_lang": "zh",
					"lang":       "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alarm_config.#":      "1",
						"use_default_contact": CHECKSET,
						"contact_config.#":    "1",
						"alarm_lang":          "zh",
						"lang":                "zh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang", "use_default_contact"},
			},
		},
	})
}

var AlicloudCloudFirewallUserAlarmConfigMap12506 = map[string]string{}

func AlicloudCloudFirewallUserAlarmConfigBasicDependence12506(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 告警测试_Notify 12512
func TestAccAliCloudCloudFirewallUserAlarmConfig_basic12512(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_user_alarm_config.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallUserAlarmConfigMap12512)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallUserAlarmConfig")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallUserAlarmConfigBasicDependence12512)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"alarm_config": []map[string]interface{}{
						{
							"alarm_value":    "on",
							"alarm_type":     "bandwidth",
							"alarm_period":   "1",
							"alarm_hour":     "0",
							"alarm_notify":   "0",
							"alarm_week_day": "0",
						},
					},
					"use_default_contact": "1",
					"alarm_lang":          "zh",
					"lang":                "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alarm_config.#":      "1",
						"use_default_contact": CHECKSET,
						"alarm_lang":          "zh",
						"lang":                "zh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang", "use_default_contact"},
			},
		},
	})
}

var AlicloudCloudFirewallUserAlarmConfigMap12512 = map[string]string{}

func AlicloudCloudFirewallUserAlarmConfigBasicDependence12512(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test CloudFirewall UserAlarmConfig. <<< Resource test cases, automatically generated.
