---
subcategory: "Ons Mqtt For Console"
layout: "alicloud"
page_title: "Alicloud: alicloud_ons_mqtt_for_console_rmc_instance"
description: |-
  Provides a Alicloud Ons Mqtt For Console Rmc Instance resource.
---

# alicloud_ons_mqtt_for_console_rmc_instance

Provides a Ons Mqtt For Console Rmc Instance resource.

Resource Management Resource.

For information about Ons Mqtt For Console Rmc Instance and how to use it, see [What is Rmc Instance](https://next.api.alibabacloud.com/document/OnsMqtt4Console/2020-04-16/ConsoleMqttInstanceList4FroRmc).

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `ons_region_id` - (Optional, ForceNew) Private resource property.
* `ons_session_id` - (Optional, ForceNew) Private resource property.
* `prevent_cache` - (Optional, ForceNew, Int) Private resource property.
* `rmc_instance_id` - (Optional, ForceNew, Computed) A resource attribute field representing the primary resource ID.
* `resource_group_id` - (Optional, Computed) The resource attribute field that represents the resource group.
* `tags` - (Optional, ForceNew, Map) Private resource property.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `delete` - (Defaults to 5 mins) Used when delete the Rmc Instance.

## Import

Ons Mqtt For Console Rmc Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_ons_mqtt_for_console_rmc_instance.example <rmc_instance_id>
```