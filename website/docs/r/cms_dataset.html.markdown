---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_dataset"
description: |-
  Provides a Alicloud Cms Dataset resource.
---

# alicloud_cms_dataset

Provides a Cms Dataset resource.



For information about Cms Dataset and how to use it, see [What is Dataset](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateDataset).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `dataset_name` - (Required, ForceNew) The name of the resource
* `description` - (Optional) Description
* `schema` - (Required, ForceNew, Map) Schema
* `workspace` - (Required, ForceNew) Workspace

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<workspace>:<dataset_name>`.
* `create_time` - The creation time of the resource.
* `region_id` - The region ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Dataset.
* `delete` - (Defaults to 5 mins) Used when delete the Dataset.
* `update` - (Defaults to 5 mins) Used when update the Dataset.

## Import

Cms Dataset can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_dataset.example <workspace>:<dataset_name>
```