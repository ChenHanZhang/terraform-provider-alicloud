---
subcategory: "Event Bridge"
layout: "alicloud"
page_title: "Alicloud: alicloud_event_bridge_table"
description: |-
  Provides a Alicloud Event Bridge Table resource.
---

# alicloud_event_bridge_table

Provides a Event Bridge Table resource.



For information about Event Bridge Table and how to use it, see [What is Table](https://next.api.alibabacloud.com/document/eventbridge/2020-04-01/CreateTable).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `catalog` - (Optional, ForceNew, Computed) Name of the catalog where the table resides  
* `columns` - (Optional, ForceNew, Set) Column information associated with the table   See [`columns`](#columns) below.
* `comment` - (Optional) Description of the table  
* `name` - (Required, ForceNew) Name of the table  
* `namespace` - (Optional, ForceNew, Computed) Name of the namespace where the table resides  
* `retention_policy` - (Optional, ForceNew, Set) Retention policy for the table's lifecycle   See [`retention_policy`](#retention_policy) below.

### `columns`

The columns supports the following:
* `comment` - (Optional) Comment for the column  
* `name` - (Optional) Name of the column  
* `type` - (Optional) Data type of the column  

### `retention_policy`

The retention_policy supports the following:
* `cold_ttl` - (Optional, Int) Retention duration for cold storage of the table  
* `hot_ttl` - (Optional, Int) Retention duration for hot storage of the table  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<catalog>:<namespace>:<name>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Table.
* `delete` - (Defaults to 5 mins) Used when delete the Table.
* `update` - (Defaults to 5 mins) Used when update the Table.

## Import

Event Bridge Table can be imported using the id, e.g.

```shell
$ terraform import alicloud_event_bridge_table.example <catalog>:<namespace>:<name>
```