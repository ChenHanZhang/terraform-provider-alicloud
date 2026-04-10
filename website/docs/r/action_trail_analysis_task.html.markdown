---
subcategory: "Actiontrail"
layout: "alicloud"
page_title: "Alicloud: alicloud_action_trail_analysis_task"
description: |-
  Provides a Alicloud Action Trail Analysis Task resource.
---

# alicloud_action_trail_analysis_task

Provides a Action Trail Analysis Task resource.



For information about Action Trail Analysis Task and how to use it, see [What is Analysis Task](https://next.api.alibabacloud.com/document/Actiontrail/2020-07-06/CreateAnalysisTask).

-> **NOTE:** Available since v1.276.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_action_trail_analysis_task" "default" {
  end_time         = "2026-04-10CST1717:0404:070728800"
  query_conditions = "{\"access_key_ids\": [\"LTAI4GDuVyEUfNcnWCsXgbbR\"]}"
  start_time       = "2026-04-10CST1717:0404:070728800"
  query_type       = "AccessKey"
}
```

## Argument Reference

The following arguments are supported:
* `end_time` - (Required, ForceNew) The end time of the analysis time range.
* `query_conditions` - (Required, ForceNew) Query conditions.
* `query_type` - (Required, ForceNew) The type of the analysis task.
* `start_time` - (Required, ForceNew) The start time of the analysis time range.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - A resource property field that represents the creation time.
* `status` - A resource property field that represents the resource status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Analysis Task.
* `delete` - (Defaults to 5 mins) Used when delete the Analysis Task.

## Import

Action Trail Analysis Task can be imported using the id, e.g.

```shell
$ terraform import alicloud_action_trail_analysis_task.example <analysis_task_id>
```