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
* `category` - (Optional, ForceNew, Computed) The snapshot type. Valid values:  
  - Standard: Standard snapshot.  
  - Flash: Local snapshot.  

-> **NOTE:**  This parameter will soon be deprecated. Standard snapshots for ESSD cloud disks have been upgraded to [instant access by default](https://help.aliyun.com/document_detail/193667.html). You do not need to configure this feature explicitly, and no additional cost is incurred.  

* `cool_off_period` - (Optional, Int, Available since v1.272.0) Compliance mode cool-off period. Unit: hours.
* `description` - (Required) The description of the new snapshot. The description must be 2 to 256 characters in length and cannot start with http:// or https://.  
Default value: empty.  
* `disk_id` - (Optional, ForceNew) The ID of the specified disk device.
* `encrypted` - (Optional, ForceNew, Available since v1.272.0) Specifies whether to filter encrypted snapshots. Default value: false.  
* `force` - (Optional) Specifies whether to forcibly delete a snapshot that has already been used to create a cloud disk. Valid values:  
  - true: Forces deletion. After forced deletion, the disk cannot be reinitialized.  
  - false: Does not force deletion.  

Default value: false.  

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `instant_access` - (Optional, Deprecated since v1.272.0) Specifies whether to enable the snapshot instant access feature. Valid values:  
  - true: Enables the feature. This feature is supported only for ESSD cloud disks.  
  - false: Disables the feature. A standard snapshot is created.  

Default value: false.  

-> **NOTE:**  This parameter is deprecated. Standard snapshots for ESSD cloud disks have been upgraded to [instant access by default](https://help.aliyun.com/document_detail/193667.html). You do not need to configure this feature explicitly, and no additional cost is incurred.  

* `instant_access_retention_days` - (Optional, ForceNew, Int, Deprecated since v1.272.0) Specifies the retention period for the snapshot instant access feature. After the retention period expires, the snapshot is automatically released. This parameter takes effect only when `InstantAccess=true`. Unit: days. Valid values: 1 to 65535.

By default, this parameter uses the same value as `RetentionDays`.

-> **NOTE:**  This parameter is deprecated. Standard snapshots of ESSD disks now support [instant access by default](https://help.aliyun.com/document_detail/193667.html). You do not need to configure this feature explicitly, and no additional charges apply.

* `kms_key_id` - (Optional, ForceNew, Available since v1.272.0) The customer master key (CMK) of Key Management Service (KMS) in the destination region.
* `lock_duration` - (Optional, Int, Available since v1.272.0) Lock duration. After this duration expires, the snapshot lock automatically becomes invalid. Unit: days.
* `lock_mode` - (Optional, Available since v1.272.0) Lock mode. Valid values: 
  - compliance: Locks the snapshot in compliance mode. A snapshot locked in compliance mode cannot be unlocked by any user and can only be deleted after the lock duration expires. Users cannot shorten the lock duration, but users with appropriate RAM permissions can extend the lock duration at any time. When locking a snapshot in compliance mode, you can optionally specify a cool-down period.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `lock_status` - (Optional, Available since v1.272.0) Lock status. Valid values: 
  - compliance-cooloff: The snapshot is locked in compliance mode but is still within the cool-off period. The snapshot cannot be deleted, but users with appropriate RAM permissions can unlock it, extend or shorten the cool-off period, or extend or shorten the lock duration. 
  - compliance: The snapshot is locked in compliance mode and the cool-off period has ended. The snapshot cannot be unlocked or deleted, but users with appropriate RAM permissions can extend the lock duration. 
  - expired: The snapshot was previously locked, but the lock duration has ended and the lock has expired. The snapshot is currently unlocked and can be deleted.
* `resource_group_id` - (Optional, Computed) The resource group ID. When you use this parameter to filter resources, the number of returned resources cannot exceed 1,000.  

-> **NOTE:**  Filtering by the default resource group is not supported.  

* `retention_days` - (Optional, Int) Number of days to retain the snapshot. The retention period starts from the snapshot's creation time (CreationTime). After a standard snapshot is created, it must be retained for at least 14 days before it can be archived.

Archived snapshots have a minimum retention requirement of 60 days in the archive tier. When calculating the retention duration for an archived snapshot, the time already spent in the standard tier is deducted. If you delete an archived snapshot before it has been retained in the archive tier for at least 60 days, you will be charged for 60 days of archive storage fees. For more information, see [Snapshot billing](https://help.aliyun.com/document_detail/56159.html).

Valid values: [74, 65536].

-> **NOTE:** If this parameter is not specified, the snapshot is retained permanently.

* `snapshot_name` - (Required) The name of the snapshot. The name must be 2 to 128 characters in length, and must start with a letter (uppercase or lowercase) or a Chinese character. It can contain letters (including English and Chinese characters as defined in the Unicode Letter category), ASCII digits (0–9), colons (:), underscores (_), periods (.), or hyphens (-).  

-> **NOTE:**  The name cannot start with http:// or https://. To avoid conflicts with automatic snapshot names, the name cannot start with `auto`.  

* `source_region_id` - (Optional, ForceNew, Available since v1.272.0) The region ID of the source snapshot. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
* `source_snapshot_id` - (Optional, ForceNew, Available since v1.272.0) Source snapshot ID.
* `tags` - (Optional, Map) Tag information for the new snapshot.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `region_id` - The ID of the destination region for the new snapshot.
* `status` - The status of the snapshot.

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