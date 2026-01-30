---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_snapshot"
description: |-
  Provides a Alicloud ECS Snapshot resource.
---

# alicloud_ecs_snapshot

Provides a ECS Snapshot resource.



For information about ECS Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/en/doc-detail/25524.htm).

-> **NOTE:** Available since v1.120.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_zones" "default" {
  available_disk_category     = "cloud_essd"
  available_resource_creation = "VSwitch"
}

data "alicloud_images" "default" {
  most_recent = true
  owners      = "system"
}

data "alicloud_instance_types" "default" {
  availability_zone    = data.alicloud_zones.default.zones.0.id
  image_id             = data.alicloud_images.default.images.0.id
  system_disk_category = "cloud_essd"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "192.168.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "192.168.192.0/24"
  zone_id      = data.alicloud_zones.default.zones.0.id
}

resource "alicloud_security_group" "default" {
  security_group_name = var.name
  vpc_id              = alicloud_vpc.default.id
}

resource "alicloud_instance" "default" {
  image_id                   = data.alicloud_images.default.images.0.id
  instance_type              = data.alicloud_instance_types.default.instance_types.0.id
  security_groups            = alicloud_security_group.default.*.id
  internet_charge_type       = "PayByTraffic"
  internet_max_bandwidth_out = "10"
  availability_zone          = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  instance_charge_type       = "PostPaid"
  system_disk_category       = "cloud_essd"
  vswitch_id                 = alicloud_vswitch.default.id
  instance_name              = var.name
  data_disks {
    category = "cloud_essd"
    size     = 20
  }
}

resource "alicloud_ecs_disk" "default" {
  disk_name = var.name
  zone_id   = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  category  = "cloud_essd"
  size      = 500
}

resource "alicloud_ecs_disk_attachment" "default" {
  disk_id     = alicloud_ecs_disk.default.id
  instance_id = alicloud_instance.default.id
}

resource "alicloud_ecs_snapshot" "default" {
  disk_id        = alicloud_ecs_disk_attachment.default.disk_id
  category       = "standard"
  retention_days = 20
}
```

## Argument Reference

The following arguments are supported:
* `category` - (Optional, ForceNew, Computed) The category of the snapshot. Valid values:
  - `standard`: Normal snapshot.
  - `flash`: Local snapshot.
* `cool_off_period` - (Optional, Int, Available since v1.270.0) Compliance mode cooling off period. Unit: hours.
* `description` - (Required) The description of the snapshot.
* `disk_id` - (Optional, ForceNew) The ID of the disk.
* `encrypted` - (Optional, ForceNew, Available since v1.270.0) Specifies whether the snapshot is encrypted. Default value: false.
* `force` - (Optional) Specifies whether to force delete the snapshot that has been used to create disks. Valid values:
  - `true`: Force deletes the snapshot. After the snapshot is force deleted, the disks created from the snapshot cannot be re-initialized.
  - `false`: Does not force delete the snapshot.

-> **NOTE:** This parameter only takes effect when deletion is triggered.

* `from_region_id` - (Optional, Available since v1.270.0) The ID of the destination region of the new snapshot.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `instant_access` - (Optional, Deprecated since v1.270.0) Field `instant_access` has been deprecated from provider version 1.231.0.
* `instant_access_retention_days` - (Optional, ForceNew, Int, Deprecated since v1.270.0) Field `instant_access_retention_days` has been deprecated from provider version 1.231.0.
* `kms_key_id` - (Optional, ForceNew, Available since v1.270.0) The ID of the Key Management Service (KMS) key that is used for the data disk.
* `lock_duration` - (Optional, Int, Available since v1.270.0) Lock is long. The snapshot lock automatically expires after the lock duration expires. Unit: days.
* `lock_mode` - (Optional, Available since v1.270.0) Lock mode. Value range:
  - compliance: Lock the snapshot in compliance mode. Snapshots locked in compliance mode cannot be unlocked by any user and can only be deleted after the lock duration expires. Users cannot shorten the lock duration, but users with the corresponding RAM permissions can extend the lock duration at any time. When locking a snapshot in compliance mode, you can optionally specify a cooling off period.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `resource_group_id` - (Optional, Computed) The ID of the resource group. **NOTE:** From version 1.239.0, `resource_group_id` can be modified.
* `retention_days` - (Optional, Int) The retention period of the snapshot. Valid values: `1` to `65536`. **NOTE:** From version 1.231.0, `retention_days` can be modified.
* `snapshot_name` - (Required) The name of the snapshot.
* `source_snapshot_id` - (Optional, Available since v1.270.0) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `tags` - (Optional, Map) A mapping of tags to assign to the resource.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The time when the snapshot was created.
* `region_id` - The region ID of the snapshot.
* `status` - The status of the Snapshot.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Snapshot.
* `delete` - (Defaults to 5 mins) Used when delete the Snapshot.
* `update` - (Defaults to 5 mins) Used when update the Snapshot.

## Import

ECS Snapshot can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_snapshot.example <snapshot_id>
```