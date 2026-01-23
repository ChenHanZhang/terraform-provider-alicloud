---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_auto_snapshot_policy_attachment"
description: |-
  Provides a Alicloud ECS Auto Snapshot Policy Attachment resource.
---

# alicloud_ecs_auto_snapshot_policy_attachment

Provides a ECS Auto Snapshot Policy Attachment resource.

Automatic snapshot policy Mount relationship.

For information about ECS Auto Snapshot Policy Attachment and how to use it, see [What is Auto Snapshot Policy Attachment](https://www.alibabacloud.com/help/en/doc-detail/25531.htm).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_zones" "example" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_kms_key" "example" {
  description            = "terraform-example"
  pending_window_in_days = "7"
  status                 = "Enabled"
}

resource "alicloud_ecs_auto_snapshot_policy" "example" {
  name            = "terraform-example"
  repeat_weekdays = ["1", "2", "3"]
  retention_days  = -1
  time_points     = ["1", "22", "23"]
}

resource "alicloud_ecs_disk" "example" {
  zone_id     = data.alicloud_zones.example.zones.0.id
  disk_name   = "terraform-example"
  description = "Hello ecs disk."
  category    = "cloud_efficiency"
  size        = "30"
  encrypted   = true
  kms_key_id  = alicloud_kms_key.example.id
  tags = {
    Name = "terraform-example"
  }
}

resource "alicloud_ecs_auto_snapshot_policy_attachment" "example" {
  auto_snapshot_policy_id = alicloud_ecs_auto_snapshot_policy.example.id
  disk_id                 = alicloud_ecs_disk.example.id
}
```

## Argument Reference

The following arguments are supported:
* `auto_snapshot_policy_id` - (Required, ForceNew) The ID of the target automatic snapshot policy.
* `disk_id` - (Required, ForceNew) The ID of one or more cloud disks. The value is in JSON Array format. Cloud disk IDs are separated by half-angle commas (,).

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<auto_snapshot_policy_id>:<disk_id>`.
* `region_id` - The ID of the region where the automatic snapshot policy and the cloud disk are located. You can call [DescribeRegions](~~ 25609 ~~) to view the latest Alibaba Cloud region list.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Auto Snapshot Policy Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Auto Snapshot Policy Attachment.

## Import

ECS Auto Snapshot Policy Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_auto_snapshot_policy_attachment.example <auto_snapshot_policy_id>:<disk_id>
```