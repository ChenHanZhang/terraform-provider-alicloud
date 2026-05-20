// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Amqp OpenSourceAccount. >>> Resource test cases, automatically generated.
// Case 开源用户 12792
func TestAccAliCloudAmqpOpenSourceAccount_basic12792(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_open_source_account.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpOpenSourceAccountMap12792)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpOpenSourceAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpOpenSourceAccountBasicDependence12792)
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
					"user_name":   "${var.user_name}",
					"description": "${var.user_name}",
					"password":    "${var.user_name}",
					"instance_id": "${alicloud_amqp_instance.CreateInstance.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_name":   CHECKSET,
						"description": CHECKSET,
						"password":    CHECKSET,
						"instance_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "${var.user_name_update}",
					"password":    "${var.user_name_update}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": CHECKSET,
						"password":    CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

var AlicloudAmqpOpenSourceAccountMap12792 = map[string]string{}

func AlicloudAmqpOpenSourceAccountBasicDependence12792(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable " _instance_name" {
  default = "测试开源鉴权实例"
}

variable "user_name" {
  default = "Suhao123_"
}

variable "user_name_update" {
  default = "Suhao456_"
}

resource "alicloud_amqp_instance" "CreateInstance" {
  period_cycle  = "Month"
  instance_name = var.instance_name
}


`, name)
}

// Test Amqp OpenSourceAccount. <<< Resource test cases, automatically generated.
