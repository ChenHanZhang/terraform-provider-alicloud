---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_delivery_task"
description: |-
  Provides a Alicloud Cms Delivery Task resource.
---

# alicloud_cms_delivery_task

Provides a Cms Delivery Task resource.



For information about Cms Delivery Task and how to use it, see [What is Delivery Task](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateDeliveryTask).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `data_source_id` - (Required) Data source ID
* `external_labels` - (Optional, Map) External labels
* `label_filters` - (Optional, Map) label filters
* `label_filters_type` - (Optional) Filter Type
* `resource_group_id` - (Optional, Computed) Resource group ID

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `sink_list` - (Required, List) Delivery Sink List See [`sink_list`](#sink_list) below.
* `status` - (Optional, Computed) The status of the resource
* `tags` - (Optional, Map) The tag of the resource
* `task_description` - (Optional) Task Description
* `task_name` - (Required) The name of the resource

### `sink_list`

The sink_list supports the following:
* `sink_configs` - (Optional, Map) Delivery sink configuration
* `sink_type` - (Required) Delivery sink type

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `region_id` - The region ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Delivery Task.
* `delete` - (Defaults to 5 mins) Used when delete the Delivery Task.
* `update` - (Defaults to 5 mins) Used when update the Delivery Task.

## Import

Cms Delivery Task can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_delivery_task.example <task_id>
```