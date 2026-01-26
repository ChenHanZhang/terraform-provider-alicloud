---
subcategory: "Log Service (SLS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_sls_ingest_processor"
description: |-
  Provides a Alicloud Log Service (SLS) Ingest Processor resource.
---

# alicloud_sls_ingest_processor

Provides a Log Service (SLS) Ingest Processor resource.



For information about Log Service (SLS) Ingest Processor and how to use it, see [What is Ingest Processor](https://next.api.alibabacloud.com/document/Sls/2020-12-30/PutIngestProcessor).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `configuration` - (Required, Set) Configuration details. See [`configuration`](#configuration) below.
* `description` - (Optional) Ingest processor description.
* `display_name` - (Required) Ingest processor display name.
* `processor_name` - (Required, ForceNew) Ingest processor identity.
* `project_name` - (Required, ForceNew) The Project name.

### `configuration`

The configuration supports the following:
* `parse_fail` - (Optional) SPL handles failed behavior. keep is used to retain the original data, and drop is used to discard the original data.
* `spl` - (Optional) The SPL statement.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<project_name>:<processor_name>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ingest Processor.
* `delete` - (Defaults to 5 mins) Used when delete the Ingest Processor.
* `update` - (Defaults to 5 mins) Used when update the Ingest Processor.

## Import

Log Service (SLS) Ingest Processor can be imported using the id, e.g.

```shell
$ terraform import alicloud_sls_ingest_processor.example <project_name>:<processor_name>
```