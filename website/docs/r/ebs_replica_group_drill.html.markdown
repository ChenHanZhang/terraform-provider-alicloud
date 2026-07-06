---
subcategory: "Elastic Block Storage(EBS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ebs_replica_group_drill"
description: |-
  Provides a Alicloud Elastic Block Storage(EBS) Replica Group Drill resource.
---

# alicloud_ebs_replica_group_drill

Provides a Elastic Block Storage(EBS) Replica Group Drill resource.



For information about Elastic Block Storage(EBS) Replica Group Drill and how to use it, see [What is Replica Group Drill](https://next.api.alibabacloud.com/document/ebs/2021-07-30/StartReplicaGroupDrill).

-> **NOTE:** Available since v1.215.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_ebs_replica_group_drill" "default" {
  group_id = "pg-m1H9aaOUIGsDUwgZ"
}
```

## Argument Reference

The following arguments are supported:
* `group_id` - (Required, ForceNew) The ID of the replication group. You can use the [describediskreplicaggroups](~~ 426614 ~~) interface to query the asynchronous replication group list to obtain the value of the replication group ID input parameter.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<group_id>:<replica_group_drill_id>`.
* `pairs_info` - Copy pair information.
  * `drill_disk_id` - Drill disk ID.
  * `drill_disk_status` - Drill disk status.
  * `pair_id` - Copy the ID of the pair.
* `recover_point` - The recovery point for the walkthrough.
* `replica_group_drill_id` - The first ID of the resource.
* `start_at` - Walkthrough start time.
* `status` - Walkthrough status.
* `status_message` - The error message indicating the task execution failure.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Replica Group Drill.
* `delete` - (Defaults to 5 mins) Used when delete the Replica Group Drill.

## Import

Elastic Block Storage(EBS) Replica Group Drill can be imported using the id, e.g.

```shell
$ terraform import alicloud_ebs_replica_group_drill.example <group_id>:<replica_group_drill_id>
```