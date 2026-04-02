---
subcategory: "Global Accelerator (GA)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ga_service_extension"
description: |-
  Provides a Alicloud Global Accelerator (GA) Service Extension resource.
---

# alicloud_ga_service_extension

Provides a Global Accelerator (GA) Service Extension resource.

A plugin bundle that extends Global Accelerator (GA) service capabilities  .

For information about Global Accelerator (GA) Service Extension and how to use it, see [What is Service Extension](https://next.api.alibabacloud.com/document/Ga/2019-11-20/CreateServiceExtension).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `components` - (Optional, List) The resource property field that describes the list of service components associated with the service extension. See [`components`](#components) below.
* `description` - (Optional) The description of the service extension.
* `name` - (Required) The name of the service extension  
* `resource_group_id` - (Optional, Computed) A resource property field representing the resource group.
* `resources` - (Optional, List) A resource property field that describes the list of associated resources linked to the service extension   See [`resources`](#resources) below.
* `tag` - (Optional, List) Tags applied to the service extension. See [`tag`](#tag) below.

### `components`

The components supports the following:
* `config` - (Required) The resource property field that describes the configuration of the service component associated with the service extension.
* `fail_policy` - (Required) The resource property field that describes the failure handling policy of the service component associated with the service extension.
* `priority` - (Required, Int) The resource property field that describes the priority of the service component associated with the service extension.
* `service_component_id` - (Required) The resource property field that describes the ID of the service component associated with the service extension.
* `timeout` - (Required, Int) The resource property field that describes the timeout duration of the service component associated with the service extension.

### `resources`

The resources supports the following:
* `accelerator_id` - (Required) The resource attribute field that describes the Global Accelerator ID linked to the service extension  
* `associate_id` - (Optional, ForceNew) The resource attribute field that describes the association ID linked to the service extension  
* `resource_id` - (Required) A resource property field that describes the ID of the associated resource linked to the service extension  
* `resource_type` - (Required) A resource property field that describes the type of the associated resource linked to the service extension
* `update_time` - (Optional, ForceNew) A resource property field that describes the update time of the associated resource linked to the service extension  

### `tag`

The tag supports the following:
* `key` - (Optional) The key of the tag  
* `value` - (Optional) The value of the tag  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `components` - The resource property field that describes the list of service components associated with the service extension.
  * `component_name` - A resource property field that describes the name of the service component associated with the service extension.
  * `create_time` - The resource property field that describes the creation time of the service component associated with the service extension.
  * `update_time` - A resource property field that describes the update time of the service component associated with the service extension.
* `resources` - A resource property field that describes the list of associated resources linked to the service extension.
  * `create_time` - A resource property field that describes the creation time of the associated resource linked to the service extension.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Service Extension.
* `delete` - (Defaults to 5 mins) Used when delete the Service Extension.
* `update` - (Defaults to 5 mins) Used when update the Service Extension.

## Import

Global Accelerator (GA) Service Extension can be imported using the id, e.g.

```shell
$ terraform import alicloud_ga_service_extension.example <service_extension_id>
```