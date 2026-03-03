---
subcategory: "Event Bridge"
layout: "alicloud"
page_title: "Alicloud: alicloud_event_bridge_namespace"
description: |-
  Provides a Alicloud Event Bridge Namespace resource.
---

# alicloud_event_bridge_namespace

Provides a Event Bridge Namespace resource.



For information about Event Bridge Namespace and how to use it, see [What is Namespace](https://next.api.alibabacloud.com/document/eventbridge/2020-04-01/CreateNamespace).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `catalog` - (Optional, Computed) This property does not have a description in the spec, please add it before generating code.
* `comment` - (Optional) The name of the resource
* `name` - (Required, ForceNew) The first ID of the resource

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<catalog>:<name>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Namespace.
* `delete` - (Defaults to 5 mins) Used when delete the Namespace.
* `update` - (Defaults to 5 mins) Used when update the Namespace.

## Import

Event Bridge Namespace can be imported using the id, e.g.

```shell
$ terraform import alicloud_event_bridge_namespace.example <catalog>:<name>
```