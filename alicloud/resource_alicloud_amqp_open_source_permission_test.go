// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Amqp OpenSourcePermission. >>> Resource test cases, automatically generated.
// Case 开源权限 12794
func TestAccAliCloudAmqpOpenSourcePermission_basic12794(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_open_source_permission.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpOpenSourcePermissionMap12794)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpOpenSourcePermission")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpOpenSourcePermissionBasicDependence12794)
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
					"write":       ".*",
					"read":        ".*",
					"vhost":       "${var.vhost}",
					"user_name":   "${var.user_name}",
					"instance_id": "${alicloud_amqp_instance.CreateInstance.id}",
					"configure":   ".*",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"write":       ".*",
						"read":        ".*",
						"vhost":       CHECKSET,
						"user_name":   CHECKSET,
						"instance_id": CHECKSET,
						"configure":   ".*",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudAmqpOpenSourcePermissionMap12794 = map[string]string{}

func AlicloudAmqpOpenSourcePermissionBasicDependence12794(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "instance_name" {
  default = "测试开源鉴权实例"
}

variable "vhost" {
  default = "/"
}

variable "user_name" {
  default = "Suhao123_WithPer"
}

resource "alicloud_amqp_instance" "CreateInstance" {
  period_cycle  = "Month"
  instance_name = var.instance_name
}


`, name)
}

// Test Amqp OpenSourcePermission. <<< Resource test cases, automatically generated.
