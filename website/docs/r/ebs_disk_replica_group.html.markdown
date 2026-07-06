---
subcategory: "Elastic Block Storage(EBS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ebs_disk_replica_group"
description: |-
  Provides a Alicloud Elastic Block Storage(EBS) Disk Replica Group resource.
---

# alicloud_ebs_disk_replica_group

Provides a Elastic Block Storage(EBS) Disk Replica Group resource.

Consistent replication groups used for block storage-based replication services.

For information about Elastic Block Storage(EBS) Disk Replica Group and how to use it, see [What is Disk Replica Group](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/creatediskreplicagroup).

-> **NOTE:** Available since v1.187.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}
data "alicloud_regions" "default" {
  current = true
}
data "alicloud_ebs_regions" "default" {
  region_id = data.alicloud_regions.default.regions.0.id
}

resource "alicloud_ebs_disk_replica_group" "default" {
  source_region_id      = data.alicloud_regions.default.regions.0.id
  source_zone_id        = data.alicloud_ebs_regions.default.regions[0].zones[0].zone_id
  destination_region_id = data.alicloud_regions.default.regions.0.id
  destination_zone_id   = data.alicloud_ebs_regions.default.regions[0].zones[1].zone_id
  group_name            = var.name
  description           = var.name
  rpo                   = 900
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) The description of the consistent replication group.
* `destination_region_id` - (Required, ForceNew) The region ID of the disaster recovery site.
* `destination_zone_id` - (Required, ForceNew) The zone ID of the disaster recovery site.
* `disk_replica_group_name` - (Optional, Computed, Available since v1.245.0) The name of the consistent replication group. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
* `one_shot` - (Optional, Available since v1.245.0) Specifies whether to perform an immediate synchronization. Valid values:
  - true: Data synchronization starts immediately.
  - false: Data synchronization starts after the RPO period elapses.

Default value: false.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `pair_ids` - (Optional, List, Available since v1.245.0) The list of replication pair IDs contained in the consistent replication group.
* `rpo` - (Optional, ForceNew, Int) The Recovery Point Objective (RPO) value configured for the consistent replication group. Unit: seconds. Only 900 seconds is supported.
* `resource_group_id` - (Optional, Computed, Available since v1.245.0) The ID of the enterprise resource group to which the consistent replication group belongs.
* `reverse_replicate` - (Optional, Available since v1.245.0) The reverse replication switch. A value of false indicates that the original replication direction is restored. A value of true indicates that reverse replication is performed. Default value: true.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `source_region_id` - (Required, ForceNew) The region ID of the consistent replication group, which is the same as the region ID of the production site.
* `source_zone_id` - (Required, ForceNew) The zone ID of the production site.
* `status` - (Optional, Computed) The status of the replication pair-consistent group. Valid values:
  - invalid: Invalid. This status indicates that an exception exists in the replication pairs within the replication pair-consistent group.
  - creating: Creating.
  - created: Created.
  - create_failed: Creation Failed.
  - manual_syncing: One-time Synchronizing. If this is the first one-time synchronization, the status is also displayed as Synchronizing.
  - syncing: Synchronizing. The group enters this status when asynchronous data replication between the primary and secondary disks occurs after the initial synchronization.
  - normal: Normal. The group enters this status when data replication within the current cycle of asynchronous replication is complete.
  - stopping: Stopping.
  - stopped: Stopped.
  - stop_failed: Stop Failed.
  - failovering: Failing Over.
  - failovered: Failover Completed.
  - failover_failed: Failover Failed.
  - reprotecting: Performing Reverse Replication.
  - reprotect_failed: Reverse Replication Failed.
  - deleting: Deleting.
  - delete_failed: Deletion Failed.
  - deleted: Deleted.
* `tags` - (Optional, Map, Available since v1.245.0) A collection consisting of resources and their tags, containing information such as resource IDs, resource types, and tag key-value pairs.

The following arguments will be discarded. Please use new fields as soon as possible:
* `group_name` - (Deprecated since v1.285.0). Field 'group_name' has been deprecated from provider version 1.285.0. New field 'disk_replica_group_name' instead.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `last_recover_point` - The time when the last asynchronous replication operation on the consistent replication group was completed.
* `pair_number` - The number of replication pairs in the replication pair-consistent group.
* `primary_region` - The initial source region of the replication group.
* `primary_zone` - The initial source zone of the replication group.
* `site` - The source of site information for replication pairs and consistent replication groups.
* `standby_region` - The initial destination region of the replication group.
* `standby_zone` - The initial destination zone of the replication group.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Disk Replica Group.
* `delete` - (Defaults to 5 mins) Used when delete the Disk Replica Group.
* `update` - (Defaults to 20 mins) Used when update the Disk Replica Group.

## Import

Elastic Block Storage(EBS) Disk Replica Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_ebs_disk_replica_group.example <replica_group_id>
```