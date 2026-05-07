---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_dataset_download_job"
description: |-
  Provides a Alicloud Cms Dataset Download Job resource.
---

# alicloud_cms_dataset_download_job

Provides a Cms Dataset Download Job resource.

Dataset Download Job.

For information about Cms Dataset Download Job and how to use it, see [What is Dataset Download Job](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateDatasetDownloadJob).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `compression_type` - (Optional, ForceNew) Compression Format
* `content_type` - (Optional, ForceNew) Result type
* `dataset_name` - (Optional, ForceNew, Computed) Dataset Name
* `job_name` - (Optional, ForceNew, Computed) Job Name
* `query` - (Optional, ForceNew) Query statement
* `workspace` - (Optional, ForceNew, Computed) Workspace

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<workspace>:<dataset_name>:<job_name>`.
* `create_time` - The creation time of the resource.
* `region_id` - The region ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Dataset Download Job.
* `delete` - (Defaults to 5 mins) Used when delete the Dataset Download Job.

## Import

Cms Dataset Download Job can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_dataset_download_job.example <workspace>:<dataset_name>:<job_name>
```