---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_cycle_task"
description: |-
  Provides a Alicloud Threat Detection Cycle Task resource.
---

# alicloud_threat_detection_cycle_task

Provides a Threat Detection Cycle Task resource.

Security Center periodic task configuration.

For information about Threat Detection Cycle Task and how to use it, see [What is Cycle Task](https://next.api.alibabacloud.com/document/Sas/2018-12-03/CreateCycleTask).

-> **NOTE:** Available since v1.253.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_cycle_task" "default" {
  target_end_time   = "6"
  task_type         = "VIRUS_VUL_SCHEDULE_SCAN"
  target_start_time = "0"
  source            = "console_batch"
  task_name         = "VIRUS_VUL_SCHEDULE_SCAN"
  first_date_str    = "1650556800000"
  period_unit       = "day"
  interval_period   = "7"
  param             = jsonencode({ "targetInfo" : [{ "type" : "groupId", "name" : "TI HOST", "target" : 10597 }, { "type" : "groupId", "name" : "expense HOST", "target" : 10597 }] })
  enable            = "1"
}
```

## Argument Reference

The following arguments are supported:
* `enable` - (Required, Int) Specifies whether the feature is enabled. Valid values:
  - `1`: Enabled
  - `0`: Disabled
* `first_date_str` - (Required, Int) The first execution time.
* `interval_period` - (Required, Int) The interval period.
* `param` - (Optional) The extended information field.
* `period_unit` - (Required) The unit of the scan cycle. Valid values:
  - `day`: days
  - `hour`: hours
* `source` - (Optional) The source of the added task.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `target_end_time` - (Required, Int) The end time of the task (hour).
* `target_start_time` - (Required, Int) The start time of the task (hour).
* `task_name` - (Required, ForceNew) The name of the task.
  - **VIRUS\_VUL\_SCHEDULE\_SCAN**: Virus scan.
  - `IMAGE_SCAN`: Image scan.
  - **EMG\_VUL\_SCHEDULE\_SCAN**: Emergency vulnerability scan.
* `task_type` - (Required, ForceNew) The type of the task.
  - **VIRUS\_VUL\_SCHEDULE\_SCAN**: Virus scan.
  - `IMAGE_SCAN`: Image scan.
  - **EMG\_VUL\_SCHEDULE\_SCAN**: Emergency vulnerability scan.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cycle Task.
* `delete` - (Defaults to 5 mins) Used when delete the Cycle Task.
* `update` - (Defaults to 5 mins) Used when update the Cycle Task.

## Import

Threat Detection Cycle Task can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_cycle_task.example <config_id>
```