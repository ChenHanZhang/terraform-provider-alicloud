---
subcategory: "PAI"
layout: "alicloud"
page_title: "Alicloud: alicloud_pai_job_template"
description: |-
  Provides a Alicloud PAI Job Template resource.
---

# alicloud_pai_job_template

Provides a PAI Job Template resource.



For information about PAI Job Template and how to use it, see [What is Job Template](https://next.api.alibabacloud.com/document/pai-dlc/2020-12-03/CreateJobTemplate).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `description` - (Optional) The description of the job template
* `metadata` - (Optional, Map) User-defined key-value metadata
* `template_name` - (Required) The name of the job template
* `workspace_id` - (Required, ForceNew) The workspace ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Job Template.
* `delete` - (Defaults to 5 mins) Used when delete the Job Template.
* `update` - (Defaults to 5 mins) Used when update the Job Template.

## Import

PAI Job Template can be imported using the id, e.g.

```shell
$ terraform import alicloud_pai_job_template.example <template_id>
```