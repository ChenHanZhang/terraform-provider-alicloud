---
subcategory: "File Storage (NAS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nas_auto_snapshot_policy"
description: |-
  Provides a Alicloud File Storage (NAS) Auto Snapshot Policy resource.
---

# alicloud_nas_auto_snapshot_policy

Provides a File Storage (NAS) Auto Snapshot Policy resource.

Automatic snapshot policy.

For information about File Storage (NAS) Auto Snapshot Policy and how to use it, see [What is Auto Snapshot Policy](https://www.alibabacloud.com/help/en/doc-detail/135662.html)).

-> **NOTE:** Available since v1.153.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_nas_auto_snapshot_policy" "default" {
  time_points               = ["0", "1", "2"]
  retention_days            = "1"
  repeat_weekdays           = ["2", "3", "4"]
  auto_snapshot_policy_name = var.name
  file_system_type          = "extreme"
}
```

## Argument Reference

The following arguments are supported:
* `auto_snapshot_policy_name` - (Optional) The name of the automatic snapshot policy.

Constraints:
  - The name must be 2 to 128 characters in length and can contain letters or Chinese characters.
  - The name must start with a letter (uppercase or lowercase) or a Chinese character.
  - The name can contain digits, colons (:), underscores (_), or hyphens (-), but it cannot start with `http://` or `https://`.
  - By default, this parameter is empty.
* `file_system_type` - (Optional, ForceNew, Computed, Available since v1.223.2) File system type.
Valid value: extreme (Extreme NAS).
* `repeat_weekdays` - (Required, List) The days of the week on which automatic snapshots are repeated.
Cycle: Week.
Valid values: 1 to 7. For example, 1 indicates Monday. If multiple automatic snapshots need to be created within a week, you can specify multiple time points, separated by commas (,). You can specify up to seven time points.
* `retention_days` - (Optional, Computed, Int) Retention period for automatic snapshots.

Unit: days

Valid values:
  - - 1 (default): Snapshots are retained permanently and automatically deleted only when the snapshot quota is reached.
  - 1 to 65536: Snapshots are retained for the specified number of days and automatically released after the retention period expires.
* `time_points` - (Required, List) The time points at which automatic snapshots are created.
Unit: hours
Valid values: 0 to 23, representing the 24 hourly time points from 00:00 to 23:00. For example, 1 represents 01:00.
If multiple automatic snapshots need to be created in a single day, you can specify multiple time points separated by commas (,). Up to 24 time points are supported.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `region_id` - Region ID.
* `status` - Status of the automatic snapshot policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Auto Snapshot Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Auto Snapshot Policy.
* `update` - (Defaults to 5 mins) Used when update the Auto Snapshot Policy.

## Import

File Storage (NAS) Auto Snapshot Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_nas_auto_snapshot_policy.example <auto_snapshot_policy_id>
```