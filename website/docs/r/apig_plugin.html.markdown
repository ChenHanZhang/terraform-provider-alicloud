---
subcategory: "APIG"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_plugin"
description: |-
  Provides a Alicloud APIG Plugin resource.
---

# alicloud_apig_plugin

Provides a APIG Plugin resource.



For information about APIG Plugin and how to use it, see [What is Plugin](https://next.api.alibabacloud.com/document/APIG/2024-03-27/InstallPlugin).

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

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "zhenyuan-example"
}

resource "alicloud_vswitch" "defaultvswitch" {
  vpc_id       = alicloud_vpc.defaultvpc.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "192.168.15.0/24"
  vswitch_name = "zhenyuan-example"
}

resource "alicloud_apig_gateway" "defaultgateway" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    name       = "zhenyuan-example"
    vswitch_id = alicloud_vswitch.defaultvswitch.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    name   = "zhenyuan-example"
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "example"
  spec         = "apigw.small.x1"
}

resource "alicloud_apig_plugin_class" "defaultpluginclass" {
  wasm_url                      = "https://apigw-console-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/1419633767709936/plugin/jwt_logout_1732865430898.wasm?Expires=1732869031&OSSAccessKeyId=STS.NTZpvmFAdKGKHB52KB6kWAUha&Signature=wVL%2BwR2Jo5c9pawlXMgUo5QoHUw%3D&security-token=CAIS4QJ1q6Ft5B2yfSjIr5fvO8zZq55F%2FIWgSmSE1ksXOuR7rpDDgzz2IHhMeXloAuAes%2FU%2FnGlY7%2Fwclr50TZJOQkrfas992ZNM6jSvfozKp82y6%2BTxaWgGxjLBZSTg1er%2BPs%2FbLrqECNrPBnnAkihsu8iYERypQ12iN7CQlJdjda55dwKkbD1Adrw0T0kY3618D3bKMuu3ORPHm3fZCFES2jBxkmRi86%2BysKb%2Bg1j89ASrkrJK%2BtqreMD%2BNpQ8bMtFPo3rjLAsRM3oyzVN7hVGzqBygZFf9C3P1tPnWAEJvkXeaLqMr4w%2FdFYpP%2FJkAdhNqPntiPtjt%2FbNlo%2F60RFJMO9SSSPZSYunxtDDHO656mO8rCs10B8nu%2FL41fmd22tMCRpzv%2FOZ5XD%2Fr1Favb09mEA7Oy6aicEHWH3Xb4Kv%2Fi%2BYH3SbMSsGE6Qk2VbBvcuXs0I6fqtYRSdOk3bRqS5sLMduGxqAAVKMRSwn42Y6vILyPqw%2Fyus3tu%2FXHiUxEMyic9J176HmhwX0gNN4ZaE9ehrdl38ru%2F5b9e9srh4W%2Bs5XwlClc6JMlyj55PcUpg%2Fzj%2FofFK2eHrFaN%2F9XtLwpfXi47FSxFk4OymlN%2FzjRShS4y3TFg%2FZBFJYYCjbgN1P0tnxhZY3yIAA%3D"
  description                   = "example插件类"
  version_description           = "example插件类版本"
  plugin_class_name             = "zhenyuan-example"
  version                       = "1.0.0"
  alias                         = "插件类别名"
  execute_priority              = "1"
  wasm_language                 = "TinyGo"
  supported_min_gateway_version = "2.0.0"
  execute_stage                 = "UNSPECIFIED_PHASE"
}


resource "alicloud_apig_plugin" "default" {
  gateway_id      = alicloud_apig_gateway.defaultgateway.id
  plugin_class_id = alicloud_apig_plugin_class.defaultpluginclass.id
}
```

## Argument Reference

The following arguments are supported:
* `gateway_id` - (Optional, ForceNew) gateway id
* `plugin_class_id` - (Optional, ForceNew) plugin class id

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Plugin.
* `delete` - (Defaults to 5 mins) Used when delete the Plugin.

## Import

APIG Plugin can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_plugin.example <plugin_id>
```