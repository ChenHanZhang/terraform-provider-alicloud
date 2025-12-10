// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test AliKafka Acl. >>> Resource test cases, automatically generated.
// Case v3 创建ACL生命周期 9888
func TestAccAliCloudAliKafkaAcl_basic9888(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ali_kafka_acl.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaAclMap9888)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaAcl")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaAclBasicDependence9888)
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
					"username":                  "${alicloud_alikafka_sasl_user.defaulty2f9ds.username}",
					"acl_operation_types":       "IDEMPOTENT_WRITE",
					"acl_operation_type":        "IDEMPOTENT_WRITE",
					"acl_permission_type":       "ALLOW",
					"host":                      "*",
					"acl_resource_pattern_type": "LITERAL",
					"acl_resource_name":         "*",
					"instance_id":               "${alicloud_alikafka_instance_v2.defaultZgBJfR.id}",
					"acl_resource_type":         "CLUSTER",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"username":                  CHECKSET,
						"acl_operation_types":       "IDEMPOTENT_WRITE",
						"acl_operation_type":        "IDEMPOTENT_WRITE",
						"acl_permission_type":       "ALLOW",
						"host":                      "*",
						"acl_resource_pattern_type": "LITERAL",
						"acl_resource_name":         "*",
						"instance_id":               CHECKSET,
						"acl_resource_type":         "CLUSTER",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"acl_operation_types"},
			},
		},
	})
}

var AlicloudAliKafkaAclMap9888 = map[string]string{}

func AlicloudAliKafkaAclBasicDependence9888(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region" {
  default = "cn-beijing"
}

resource "alicloud_vpc" "default0ppNRd" {
  ipv4_cidr_mask = "24"
  is_default     = false
  cidr_block     = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultn6iyAD" {
  vpc_id     = alicloud_vpc.default0ppNRd.id
  cidr_block = "172.18.0.0/23"
  zone_id    = "cn-beijing-a"
}

resource "alicloud_alikafka_instance_v2" "defaultZgBJfR" {
  deploy_type = "5"
  spec_type   = "normal"
  config      = "{\"auto.create.topics.enable\":\"true\",\"enable.acl\":\"true\",\"enable.vpc_sasl_ssl\":\"false\",\"kafka.ssl.bit\":\"4096\",\"log.retention.hours\":\"72\",\"message.max.bytes\":\"1048576\",\"num.partitions\":\"3\",\"offsets.retention.minutes\":\"10080\"}"
  zone_id     = "cn-beijing-a"
  vswitch_id  = alicloud_vswitch.defaultn6iyAD.id
  paid_type   = "3"
  serverless_config {
    reserved_publish_capacity   = "60"
    reserved_subscribe_capacity = "60"
  }
  vpc_id = alicloud_vpc.default0ppNRd.id
}

resource "alicloud_alikafka_sasl_user" "defaulty2f9ds" {
  type        = "scram"
  username    = "qwoeiuqwoieurandom866"
  password    = "123123"
  instance_id = alicloud_alikafka_instance_v2.defaultZgBJfR.id
}


`, name)
}

// Test AliKafka Acl. <<< Resource test cases, automatically generated.
