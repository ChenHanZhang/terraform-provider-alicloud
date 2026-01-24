// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Apig PluginClass. >>> Resource test cases, automatically generated.
// Case pluginclass-test 9272
func TestAccAliCloudApigPluginClass_basic9272(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_plugin_class.default"
	ra := resourceAttrInit(resourceId, AlicloudApigPluginClassMap9272)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigPluginClass")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigPluginClassBasicDependence9272)
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
					"wasm_url":                      "https://apigw-console-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/1419633767709936/plugin/jwt_logout_1732865430898.wasm?Expires=1732869031&OSSAccessKeyId=STS.NTZpvmFAdKGKHB52KB6kWAUha&Signature=wVL%2BwR2Jo5c9pawlXMgUo5QoHUw%3D&security-token=CAIS4QJ1q6Ft5B2yfSjIr5fvO8zZq55F%2FIWgSmSE1ksXOuR7rpDDgzz2IHhMeXloAuAes%2FU%2FnGlY7%2Fwclr50TZJOQkrfas992ZNM6jSvfozKp82y6%2BTxaWgGxjLBZSTg1er%2BPs%2FbLrqECNrPBnnAkihsu8iYERypQ12iN7CQlJdjda55dwKkbD1Adrw0T0kY3618D3bKMuu3ORPHm3fZCFES2jBxkmRi86%2BysKb%2Bg1j89ASrkrJK%2BtqreMD%2BNpQ8bMtFPo3rjLAsRM3oyzVN7hVGzqBygZFf9C3P1tPnWAEJvkXeaLqMr4w%2FdFYpP%2FJkAdhNqPntiPtjt%2FbNlo%2F60RFJMO9SSSPZSYunxtDDHO656mO8rCs10B8nu%2FL41fmd22tMCRpzv%2FOZ5XD%2Fr1Favb09mEA7Oy6aicEHWH3Xb4Kv%2Fi%2BYH3SbMSsGE6Qk2VbBvcuXs0I6fqtYRSdOk3bRqS5sLMduGxqAAVKMRSwn42Y6vILyPqw%2Fyus3tu%2FXHiUxEMyic9J176HmhwX0gNN4ZaE9ehrdl38ru%2F5b9e9srh4W%2Bs5XwlClc6JMlyj55PcUpg%2Fzj%2FofFK2eHrFaN%2F9XtLwpfXi47FSxFk4OymlN%2FzjRShS4y3TFg%2FZBFJYYCjbgN1P0tnxhZY3yIAA%3D",
					"description":                   "镇元测试插件类",
					"version_description":           "测试插件类版本",
					"plugin_class_name":             name,
					"version":                       "1.0.2",
					"execute_stage":                 "UNSPECIFIED_PHASE",
					"wasm_language":                 "TinyGo",
					"execute_priority":              "1",
					"alias":                         "插件类别名",
					"supported_min_gateway_version": "2.0.0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"wasm_url":                      "https://apigw-console-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/1419633767709936/plugin/jwt_logout_1732865430898.wasm?Expires=1732869031&OSSAccessKeyId=STS.NTZpvmFAdKGKHB52KB6kWAUha&Signature=wVL%2BwR2Jo5c9pawlXMgUo5QoHUw%3D&security-token=CAIS4QJ1q6Ft5B2yfSjIr5fvO8zZq55F%2FIWgSmSE1ksXOuR7rpDDgzz2IHhMeXloAuAes%2FU%2FnGlY7%2Fwclr50TZJOQkrfas992ZNM6jSvfozKp82y6%2BTxaWgGxjLBZSTg1er%2BPs%2FbLrqECNrPBnnAkihsu8iYERypQ12iN7CQlJdjda55dwKkbD1Adrw0T0kY3618D3bKMuu3ORPHm3fZCFES2jBxkmRi86%2BysKb%2Bg1j89ASrkrJK%2BtqreMD%2BNpQ8bMtFPo3rjLAsRM3oyzVN7hVGzqBygZFf9C3P1tPnWAEJvkXeaLqMr4w%2FdFYpP%2FJkAdhNqPntiPtjt%2FbNlo%2F60RFJMO9SSSPZSYunxtDDHO656mO8rCs10B8nu%2FL41fmd22tMCRpzv%2FOZ5XD%2Fr1Favb09mEA7Oy6aicEHWH3Xb4Kv%2Fi%2BYH3SbMSsGE6Qk2VbBvcuXs0I6fqtYRSdOk3bRqS5sLMduGxqAAVKMRSwn42Y6vILyPqw%2Fyus3tu%2FXHiUxEMyic9J176HmhwX0gNN4ZaE9ehrdl38ru%2F5b9e9srh4W%2Bs5XwlClc6JMlyj55PcUpg%2Fzj%2FofFK2eHrFaN%2F9XtLwpfXi47FSxFk4OymlN%2FzjRShS4y3TFg%2FZBFJYYCjbgN1P0tnxhZY3yIAA%3D",
						"description":                   "镇元测试插件类",
						"version_description":           "测试插件类版本",
						"plugin_class_name":             name,
						"version":                       "1.0.2",
						"execute_stage":                 "UNSPECIFIED_PHASE",
						"wasm_language":                 "TinyGo",
						"execute_priority":              "1",
						"alias":                         "插件类别名",
						"supported_min_gateway_version": "2.0.0",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"execute_priority", "execute_stage", "supported_min_gateway_version", "version", "version_description", "wasm_url"},
			},
		},
	})
}

var AlicloudApigPluginClassMap9272 = map[string]string{
	"status": CHECKSET,
}

func AlicloudApigPluginClassBasicDependence9272(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Apig PluginClass. <<< Resource test cases, automatically generated.
