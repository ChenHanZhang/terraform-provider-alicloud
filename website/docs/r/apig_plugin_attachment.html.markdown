---
subcategory: "Cloud Native API Gateway (APIG)"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_plugin_attachment"
description: |-
  Provides a Alicloud APIG Plugin Attachment resource.
---

# alicloud_apig_plugin_attachment

Provides a APIG Plugin Attachment resource.

Plug-in attachment information.

For information about APIG Plugin Attachment and how to use it, see [What is Plugin Attachment](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreatePluginAttachment).

-> **NOTE:** Available since v1.285.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_vpc" "pa_vpc_pre" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pa-example-vpc"
}

resource "alicloud_vswitch" "pa_vswitch_pre" {
  is_default   = false
  vpc_id       = alicloud_vpc.pa_vpc_pre.id
  zone_id      = "cn-hangzhou-i"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pa-example-vswitch"
}

resource "alicloud_apig_gateway" "pa_gateway_pre" {
  network_access_config {
    type = "Internet"
  }
  vswitch {
    vswitch_id = alicloud_vswitch.pa_vswitch_pre.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    vpc_id = alicloud_vpc.pa_vpc_pre.id
  }
  gateway_type = "API"
  payment_type = "PayAsYouGo"
  gateway_name = "pa-example-gateway"
  spec         = "apigw.small.x1"
  log_config {
    sls {
      enable = true
    }
  }
}

resource "alicloud_apig_plugin" "pa_plugin_pre" {
  plugin_class_id = "pls-crpqb35lhtgo800k2m86"
  gateway_id      = alicloud_apig_gateway.pa_gateway_pre.id
}

resource "alicloud_apig_environment" "pa_env_pre" {
  environment_name = "pa-example-env"
  gateway_id       = alicloud_apig_gateway.pa_gateway_pre.id
}

resource "alicloud_apig_http_api" "pa_http_api_pre" {
  http_api_name = "pa-example-httpapi"
  protocols     = ["HTTP"]
  type          = "Rest"
  base_path     = "/pa-example"
}


resource "alicloud_apig_plugin_attachment" "default" {
  attach_resource_ids = ["${alicloud_apig_http_api.pa_http_api_pre.id}"]
  environment_id      = alicloud_apig_environment.pa_env_pre.id
  enable              = true
  plugin_info {
    plugin_config = "eyJ0ZXN0IjoiaGVsbG8ifQ=="
    plugin_id     = alicloud_apig_plugin.pa_plugin_pre.id
    gateway_id    = alicloud_apig_gateway.pa_gateway_pre.id
  }
  attach_resource_type = "HttpApi"
}
```

## Argument Reference

The following arguments are supported:
* `attach_resource_id` - (Optional, ForceNew) The ID of the attached resource.
* `attach_resource_ids` - (Optional, List) The list of mount point IDs.
* `attach_resource_type` - (Optional, ForceNew) The type of the resource to which the plug-in is attached, such as GatewayRoute, Gateway, GatewayDomain, HttpApi, or Operation.
* `enable` - (Optional) Specifies whether to enable the feature. Default value: false.
* `environment_id` - (Optional, ForceNew) The environment ID.
* `plugin_info` - (Optional, ForceNew, Set) The association information between the plug-in and the gateway instance. See [`plugin_info`](#plugin_info) below.

### `plugin_info`

The plugin_info supports the following:
* `gateway_id` - (Optional, ForceNew) The gateway instance ID.
* `plugin_config` - (Optional) The Base64-encoded content of the original plug-in configuration.
* `plugin_id` - (Optional, ForceNew) The plug-in ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `attach_resource_names` - The list of names of resources to which the plug-in is attached.
* `attach_resource_parent_ids` - The list of parent node IDs for the resources to which the plug-in is attached.
* `plugin_class_info` - The type information of the attached plug-in.
  * `direction` - The direction in which the plug-in acts on traffic: InBound, OutBound, or Both.
  * `execute_priority` - The execution priority of the plug-in.
  * `execute_stage` - The execution stage of the plug-in.
  * `name` - The name of the plug-in.
  * `type` - The plug-in type.

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