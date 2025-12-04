---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_routine_route"
description: |-
  Provides a Alicloud ESA Routine Route resource.
---

# alicloud_esa_routine_route

Provides a ESA Routine Route resource.



For information about ESA Routine Route and how to use it, see [What is Routine Route](https://next.api.alibabacloud.com/document/ESA/2024-09-10/CreateRoutineRoute).

-> **NOTE:** Available since v1.265.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `bypass` - (Optional) Bypass mode. Value range:
  - on: Open
  - off: off
* `fallback` - (Optional) Spare
* `route_enable` - (Optional) Routing switch. Value range:
  - on: Open
  - off: off
* `route_name` - (Optional) The route name.
* `routine_name` - (Required, ForceNew) The edge function Routine name.
* `rule` - (Optional) The rule content.
* `sequence` - (Optional, Int) Rule execution order.
* `site_id` - (Required, ForceNew) Site Id

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<site_id>:<routine_name>:<config_id>`.
* `config_id` - Config Id

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Routine Route.
* `delete` - (Defaults to 5 mins) Used when delete the Routine Route.
* `update` - (Defaults to 5 mins) Used when update the Routine Route.

## Import

ESA Routine Route can be imported using the id, e.g.

```shell
$ terraform import alicloud_esa_routine_route.example <site_id>:<routine_name>:<config_id>
```