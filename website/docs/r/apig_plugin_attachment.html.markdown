---
subcategory: "APIG"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_plugin_attachment"
description: |-
  Provides a Alicloud APIG Plugin Attachment resource.
---

# alicloud_apig_plugin_attachment

Provides a APIG Plugin Attachment resource.

plugin attachment information.

For information about APIG Plugin Attachment and how to use it, see [What is Plugin Attachment](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreatePluginAttachment).

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
  cidr_block   = "192.168.26.0/24"
  vswitch_name = "zhenyuan-example"
}

resource "alicloud_apig_gateway" "defaultgateway" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    name       = alicloud_vswitch.defaultvswitch.vswitch_name
    vswitch_id = alicloud_vswitch.defaultvswitch.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    name   = alicloud_vpc.defaultvpc.vpc_name
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "example"
  spec         = "apigw.small.x1"
}

resource "alicloud_apig_plugin" "defaultplugin" {
  plugin_class_id = "pls-cqebrgh4ckt6ppatmprc"
  gateway_id      = alicloud_apig_gateway.defaultgateway.id
}


resource "alicloud_apig_plugin_attachment" "default" {
  attach_resource_ids = ["${alicloud_apig_gateway.defaultgateway.id}"]
  enable              = false
  plugin_info {
    plugin_id     = alicloud_apig_plugin.defaultplugin.id
    gateway_id    = alicloud_apig_gateway.defaultgateway.id
    plugin_config = "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQKYm9keTogIntcImhlbGxvXCI6XCJ3b3JsZFwifSI="
  }
  attach_resource_type = "Gateway"
}
```

## Argument Reference

The following arguments are supported:
* `attach_resource_id` - (Optional, ForceNew) plugin attache resource id
* `attach_resource_ids` - (Optional, List) plugin attach resource id list
* `attach_resource_type` - (Optional, ForceNew) plugin attach resource type
* `enable` - (Optional) enable
* `environment_id` - (Optional, ForceNew) environment id
* `plugin_info` - (Optional, ForceNew, Set) information about the association of plugin with gateway instance See [`plugin_info`](#plugin_info) below.

### `plugin_info`

The plugin_info supports the following:
* `gateway_id` - (Optional, ForceNew) gateway id
* `plugin_config` - (Optional) plugin configuration
* `plugin_id` - (Optional, ForceNew) ID of plugin and gateway association

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Plugin Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Plugin Attachment.
* `update` - (Defaults to 5 mins) Used when update the Plugin Attachment.

## Import

APIG Plugin Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_plugin_attachment.example <plugin_attachment_id>
```