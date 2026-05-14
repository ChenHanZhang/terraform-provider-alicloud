---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_oncall_schedule"
description: |-
  Provides a Alicloud Cms Oncall Schedule resource.
---

# alicloud_cms_oncall_schedule

Provides a Cms Oncall Schedule resource.



For information about Cms Oncall Schedule and how to use it, see [What is Oncall Schedule](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateOncallSchedule).

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_cms_oncall_schedule" "default" {
  oncall_schedule_name = "OncallScheduleName"
  shift_robot_id       = "OncallScheduleName"
  rotations {
    active_days                = ["1", "2"]
    contacts                   = ["a", "b"]
    rotation_end_time          = "aaa"
    rotation_name              = "bbb"
    rotation_start_time        = "ccc"
    shift_recurrence_frequency = "DAILY"
    start_date                 = "ddd"
    time_zone                  = "zzzz"
  }
  oncall_schedule_id = "OncallScheduleId123"
}
```

## Argument Reference

The following arguments are supported:
* `oncall_schedule_id` - (Optional, ForceNew, Computed) The first ID of the resource
* `oncall_schedule_name` - (Optional) The name of the resource
* `rotations` - (Optional, List) Rotation See [`rotations`](#rotations) below.
* `shift_robot_id` - (Optional) Shift Reminder Robot

### `rotations`

The rotations supports the following:
* `active_days` - (Required, List) Week effective
* `contacts` - (Required, List) Person on duty
* `rotation_end_time` - (Required) Effective end time
* `rotation_name` - (Required) Shift Name
* `rotation_start_time` - (Required) Effective start time
* `shift_length` - (Optional, Int) Alternating Duration
* `shift_recurrence_frequency` - (Required) Shift Cycle
* `start_date` - (Required) 开始日期
* `time_zone` - (Required) Time Zone

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Oncall Schedule.
* `delete` - (Defaults to 5 mins) Used when delete the Oncall Schedule.
* `update` - (Defaults to 5 mins) Used when update the Oncall Schedule.

## Import

Cms Oncall Schedule can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_oncall_schedule.example <oncall_schedule_id>
```