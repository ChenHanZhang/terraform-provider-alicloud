// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test AliKafka AllowedIpAttachment. >>> Resource test cases, automatically generated.
// Case 白名单操作生命周期 9872
func TestAccAliCloudAliKafkaAllowedIpAttachment_basic9872(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ali_kafka_allowed_ip_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaAllowedIpAttachmentMap9872)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaAllowedIpAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaAllowedIpAttachmentBasicDependence9872)
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
					"description":       "${var.desc}",
					"instance_id":       "${alicloud_alikafka_instance_v2.defaultBzMsgy.id}",
					"port_range":        "9092/9092",
					"allowed_list_ip":   "${var.inner_ip}",
					"allowed_list_type": "vpc",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       CHECKSET,
						"instance_id":       CHECKSET,
						"port_range":        "9092/9092",
						"allowed_list_ip":   CHECKSET,
						"allowed_list_type": "vpc",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"description"},
			},
		},
	})
}

var AlicloudAliKafkaAllowedIpAttachmentMap9872 = map[string]string{}

func AlicloudAliKafkaAllowedIpAttachmentBasicDependence9872(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region" {
  default = "cn-beijing"
}

variable "inner_ip" {
  default = "192.168.1.89"
}

variable "desc" {
  default = "test-516"
}

resource "alicloud_vpc" "defaultbqyuKu" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultjvEBtT" {
  vpc_id     = alicloud_vpc.defaultbqyuKu.id
  zone_id    = "cn-beijing-a"
  cidr_block = "172.18.0.0/23"
}

resource "alicloud_alikafka_instance_v2" "defaultBzMsgy" {
  deploy_type = "5"
  spec_type   = "normal"
  config      = "{\"auto.create.topics.enable\":\"true\",\"enable.acl\":\"true\",\"enable.vpc_sasl_ssl\":\"false\",\"kafka.ssl.bit\":\"4096\",\"log.retention.hours\":\"72\",\"message.max.bytes\":\"1048576\",\"num.partitions\":\"3\",\"offsets.retention.minutes\":\"10080\"}"
  zone_id     = "cn-beijing-a"
  vswitch_id  = alicloud_vswitch.defaultjvEBtT.id
  vpc_id      = alicloud_vpc.defaultbqyuKu.id
  paid_type   = "3"
  serverless_config {
    reserved_publish_capacity   = "60"
    reserved_subscribe_capacity = "60"
  }
}


`, name)
}

// Test AliKafka AllowedIpAttachment. <<< Resource test cases, automatically generated.
