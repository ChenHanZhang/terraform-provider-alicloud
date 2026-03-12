---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_evaluation_task"
description: |-
  Provides a Alicloud Cms Evaluation Task resource.
---

# alicloud_cms_evaluation_task

Provides a Cms Evaluation Task resource.



For information about Cms Evaluation Task and how to use it, see [What is Evaluation Task](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateEvaluationTask).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `channel` - (Optional) Channel source.
* `config` - (Optional, Map) Task configuration.
* `data_filter` - (Optional) DataFilter.
* `data_type` - (Optional, ForceNew) Data source type.
* `description` - (Optional) Task description.
* `evaluators` - (Required, List) Evaluator configuration.   See [`evaluators`](#evaluators) below.
* `run_strategies` - (Optional) Run strategies.
* `status` - (Optional, Computed) Task status.
* `tags` - (Optional, Map) Attribute tags.
* `task_id` - (Optional, Computed) Task ID.
* `task_mode` - (Optional, ForceNew) Task mode.
* `task_name` - (Required, ForceNew) Task name.

### `evaluators`

The evaluators supports the following:
* `config` - (Optional, Map) Evaluator configuration.  
* `filters` - (Optional, Map) Data filtering configuration.  
* `name` - (Required) Evaluator name.  
* `result_name` - (Required) Metric name.  
* `result_type` - (Optional) Result type.  
* `variable_mapping` - (Required, Map) Evaluator variable mapping.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Resource property field representing the creation time.
* `region_id` - Resource property field representing the region.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Evaluation Task.
* `delete` - (Defaults to 5 mins) Used when delete the Evaluation Task.
* `update` - (Defaults to 5 mins) Used when update the Evaluation Task.

## Import

Cms Evaluation Task can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_evaluation_task.example <task_id>
```