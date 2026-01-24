// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Apig Plugin. >>> Resource test cases, automatically generated.
// Case pluginTest 8906
func TestAccAliCloudApigPlugin_basic8906(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_plugin.default"
	ra := resourceAttrInit(resourceId, AlicloudApigPluginMap8906)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigPlugin")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigPluginBasicDependence8906)
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
					"gateway_id":      "${alicloud_apig_gateway.defaultgateway.id}",
					"plugin_class_id": "${alicloud_apig_plugin_class.defaultpluginclass.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_id":      CHECKSET,
						"plugin_class_id": CHECKSET,
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

var AlicloudApigPluginMap8906 = map[string]string{}

func AlicloudApigPluginBasicDependence8906(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "zhenyuan-test"
}

resource "alicloud_vswitch" "defaultvswitch" {
  vpc_id       = alicloud_vpc.defaultvpc.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "192.168.15.0/24"
  vswitch_name = "zhenyuan-test"
}

resource "alicloud_apig_gateway" "defaultgateway" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    name       = "zhenyuan-test"
    vswitch_id = alicloud_vswitch.defaultvswitch.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    name   = "zhenyuan-test"
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "test"
  spec         = "apigw.small.x1"
}

resource "alicloud_apig_plugin_class" "defaultpluginclass" {
  wasm_url                      = "https://apigw-console-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/1419633767709936/plugin/jwt_logout_1732865430898.wasm?Expires=1732869031&OSSAccessKeyId=STS.NTZpvmFAdKGKHB52KB6kWAUha&Signature=wVL%2BwR2Jo5c9pawlXMgUo5QoHUw%3D&security-token=CAIS4QJ1q6Ft5B2yfSjIr5fvO8zZq55F%2FIWgSmSE1ksXOuR7rpDDgzz2IHhMeXloAuAes%2FU%2FnGlY7%2Fwclr50TZJOQkrfas992ZNM6jSvfozKp82y6%2BTxaWgGxjLBZSTg1er%2BPs%2FbLrqECNrPBnnAkihsu8iYERypQ12iN7CQlJdjda55dwKkbD1Adrw0T0kY3618D3bKMuu3ORPHm3fZCFES2jBxkmRi86%2BysKb%2Bg1j89ASrkrJK%2BtqreMD%2BNpQ8bMtFPo3rjLAsRM3oyzVN7hVGzqBygZFf9C3P1tPnWAEJvkXeaLqMr4w%2FdFYpP%2FJkAdhNqPntiPtjt%2FbNlo%2F60RFJMO9SSSPZSYunxtDDHO656mO8rCs10B8nu%2FL41fmd22tMCRpzv%2FOZ5XD%2Fr1Favb09mEA7Oy6aicEHWH3Xb4Kv%2Fi%2BYH3SbMSsGE6Qk2VbBvcuXs0I6fqtYRSdOk3bRqS5sLMduGxqAAVKMRSwn42Y6vILyPqw%2Fyus3tu%2FXHiUxEMyic9J176HmhwX0gNN4ZaE9ehrdl38ru%2F5b9e9srh4W%2Bs5XwlClc6JMlyj55PcUpg%2Fzj%2FofFK2eHrFaN%2F9XtLwpfXi47FSxFk4OymlN%2FzjRShS4y3TFg%2FZBFJYYCjbgN1P0tnxhZY3yIAA%3D"
  description                   = "测试插件类"
  version_description           = "测试插件类版本"
  plugin_class_name             = "zhenyuan-test"
  version                       = "1.0.0"
  alias                         = "插件类别名"
  execute_priority              = "1"
  wasm_language                 = "TinyGo"
  supported_min_gateway_version = "2.0.0"
  execute_stage                 = "UNSPECIFIED_PHASE"
}


`, name)
}

// Test Apig Plugin. <<< Resource test cases, automatically generated.
