---
subcategory: "APIG"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_plugin_class"
description: |-
  Provides a Alicloud APIG Plugin Class resource.
---

# alicloud_apig_plugin_class

Provides a APIG Plugin Class resource.

plugin class info.

For information about APIG Plugin Class and how to use it, see [What is Plugin Class](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreatePluginClass).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_apig_plugin_class" "default" {
  wasm_url                      = "https://apigw-console-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/1419633767709936/plugin/jwt_logout_1732865430898.wasm?Expires=1732869031&OSSAccessKeyId=STS.NTZpvmFAdKGKHB52KB6kWAUha&Signature=wVL%2BwR2Jo5c9pawlXMgUo5QoHUw%3D&security-token=CAIS4QJ1q6Ft5B2yfSjIr5fvO8zZq55F%2FIWgSmSE1ksXOuR7rpDDgzz2IHhMeXloAuAes%2FU%2FnGlY7%2Fwclr50TZJOQkrfas992ZNM6jSvfozKp82y6%2BTxaWgGxjLBZSTg1er%2BPs%2FbLrqECNrPBnnAkihsu8iYERypQ12iN7CQlJdjda55dwKkbD1Adrw0T0kY3618D3bKMuu3ORPHm3fZCFES2jBxkmRi86%2BysKb%2Bg1j89ASrkrJK%2BtqreMD%2BNpQ8bMtFPo3rjLAsRM3oyzVN7hVGzqBygZFf9C3P1tPnWAEJvkXeaLqMr4w%2FdFYpP%2FJkAdhNqPntiPtjt%2FbNlo%2F60RFJMO9SSSPZSYunxtDDHO656mO8rCs10B8nu%2FL41fmd22tMCRpzv%2FOZ5XD%2Fr1Favb09mEA7Oy6aicEHWH3Xb4Kv%2Fi%2BYH3SbMSsGE6Qk2VbBvcuXs0I6fqtYRSdOk3bRqS5sLMduGxqAAVKMRSwn42Y6vILyPqw%2Fyus3tu%2FXHiUxEMyic9J176HmhwX0gNN4ZaE9ehrdl38ru%2F5b9e9srh4W%2Bs5XwlClc6JMlyj55PcUpg%2Fzj%2FofFK2eHrFaN%2F9XtLwpfXi47FSxFk4OymlN%2FzjRShS4y3TFg%2FZBFJYYCjbgN1P0tnxhZY3yIAA%3D"
  description                   = "镇元example插件类"
  version_description           = "example插件类版本"
  plugin_class_name             = "example-pluginclass"
  version                       = "1.0.2"
  execute_stage                 = "UNSPECIFIED_PHASE"
  wasm_language                 = "TinyGo"
  execute_priority              = "1"
  alias                         = "插件类别名"
  supported_min_gateway_version = "2.0.0"
}
```

## Argument Reference

The following arguments are supported:
* `alias` - (Optional, ForceNew) plugin class alias
* `description` - (Required, ForceNew) plugin class description 
* `execute_priority` - (Required, Int) plugin execute priority

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `execute_stage` - (Required) plugin execute stage

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `plugin_class_name` - (Required, ForceNew) The name of the plugin class
* `supported_min_gateway_version` - (Optional) supported minimum gateway version

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `version` - (Required) plugin class version

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `version_description` - (Required) version description

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `wasm_language` - (Required, ForceNew) wasm language
* `wasm_url` - (Required) wasm url

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - publish status of plugin class 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Plugin Class.
* `delete` - (Defaults to 5 mins) Used when delete the Plugin Class.

## Import

APIG Plugin Class can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_plugin_class.example <plugin_class_id>
```