// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test MessageService AccountLogging. >>> Resource test cases, automatically generated.
// Case resourceCase_20260714_KVFLqK 13024
func TestAccAliCloudMessageServiceAccountLogging_basic13024(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_message_service_account_logging.default"
	ra := resourceAttrInit(resourceId, AlicloudMessageServiceAccountLoggingMap13024)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MessageServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMessageServiceAccountLogging")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmessageservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMessageServiceAccountLoggingBasicDependence13024)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"log_store_name":        "cloudspec-test-logstore",
					"project_name":          "cloudspec-test-project",
					"log_enabled":           "true",
					"message_trace_enabled": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"log_store_name":        "cloudspec-test-logstore",
						"project_name":          "cloudspec-test-project",
						"log_enabled":           "true",
						"message_trace_enabled": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"log_store_name":        "cloudspec-test-logstore-updated",
					"project_name":          "cloudspec-test-project-updated",
					"message_trace_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"log_store_name":        "cloudspec-test-logstore-updated",
						"project_name":          "cloudspec-test-project-updated",
						"message_trace_enabled": "true",
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

var AlicloudMessageServiceAccountLoggingMap13024 = map[string]string{}

func AlicloudMessageServiceAccountLoggingBasicDependence13024(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_log_project" "project" {
  description = "cloudspec test project for AccountLogging"
  name        = "cloudspec-test-project"
}

resource "alicloud_log_store" "logstore" {
  append_meta      = false
  retention_period = "30"
  shard_count      = "2"
  project          = "cloudspec-test-project"
  name             = "cloudspec-test-logstore"
}


`, name)
}

// Test MessageService AccountLogging. <<< Resource test cases, automatically generated.
