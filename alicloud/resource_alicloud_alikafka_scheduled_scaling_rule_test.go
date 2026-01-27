// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Alikafka ScheduledScalingRule. >>> Resource test cases, automatically generated.
// Case 定时策略全生命周期-网段-195.0.0.0/25 12412
func TestAccAliCloudAlikafkaScheduledScalingRule_basic12412(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_scheduled_scaling_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudAlikafkaScheduledScalingRuleMap12412)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AlikafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAlikafkaScheduledScalingRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAlikafkaScheduledScalingRuleBasicDependence12412)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"schedule_type":        "repeat",
					"reserved_sub_flow":    "200",
					"reserved_pub_flow":    "200",
					"time_zone":            "GMT+8",
					"duration_minutes":     "100",
					"first_scheduled_time": "1769480056",
					"enable":               "true",
					"repeat_type":          "Weekly",
					"weekly_types": []string{
						"${var.weekly_types}"},
					"rule_name":   "${var.rulename}",
					"instance_id": "${ alicloud_alikafka_instance.default1wRjcq.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"schedule_type":        "repeat",
						"reserved_sub_flow":    "200",
						"reserved_pub_flow":    "200",
						"time_zone":            "GMT+8",
						"duration_minutes":     "100",
						"first_scheduled_time": CHECKSET,
						"enable":               "true",
						"repeat_type":          "Weekly",
						"weekly_types.#":       "1",
						"rule_name":            CHECKSET,
						"instance_id":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable": "false",
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

var AlicloudAlikafkaScheduledScalingRuleMap12412 = map[string]string{}

func AlicloudAlikafkaScheduledScalingRuleBasicDependence12412(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "weekly_types" {
  default = "Monday"
}

variable "rulename" {
  default = "test"
}

variable "region" {
  default = "cn-beijing"
}

resource "alicloud_vpc" "defaultqSIhPN" {
  cidr_block = "195.0.0.0/24"
}

resource "alicloud_vswitch" "defaulttqV6cH" {
  vpc_id     = alicloud_vpc.defaultqSIhPN.id
  cidr_block = "195.0.0.0/25"
  zone_id    = "cn-beijing-a"
}

resource " alicloud_alikafka_instance" "default1wRjcq" {
  deploy_type   = "5"
  spec_type     = "normal"
  deploy_module = "vpc"
  vswitch_id    = alicloud_vswitch.defaulttqV6cH.id
  vpc_id        = alicloud_vpc.defaultqSIhPN.id
  paid_type     = "3"
  serverless_config {
    reserved_publish_capacity   = "60"
    reserved_subscribe_capacity = "60"
  }
}


`, name)
}

// Test Alikafka ScheduledScalingRule. <<< Resource test cases, automatically generated.
