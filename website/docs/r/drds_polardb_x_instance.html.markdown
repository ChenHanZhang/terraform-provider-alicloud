---
subcategory: "DRDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_drds_polardb_x_instance"
description: |-
  Provides a Alicloud DRDS Polardb X Instance resource.
---

# alicloud_drds_polardb_x_instance

Provides a DRDS Polardb X Instance resource. PolarDB-X Database Instance.

For information about DRDS Polardb X Instance and how to use it, see [What is Polardb X Instance](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.210.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_resource_manager_resource_group" "default4m883T" {
  display_name        = "terraform-test-all"
  resource_group_name = var.name
}

resource "alicloud_resource_manager_resource_group" "defaultL20ZLi" {
  display_name        = "terraform-test-all-2"
  resource_group_name = "${var.name}1"
}

resource "alicloud_vpc" "defaultI3SPrf" {
  vpc_name = "${var.name}2"
}

resource "alicloud_vswitch" "defaultV9mMOX" {
  vpc_id       = alicloud_vpc.defaultI3SPrf.id
  zone_id      = data.alicloud_zones.default.zones.0.id
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "${var.name}3"
}


resource "alicloud_drds_polardb_x_instance" "default" {
  topology_type     = "3azones"
  vswitch_id        = alicloud_vswitch.defaultV9mMOX.id
  primary_zone      = "cn-beijing-i"
  cn_node_count     = "2"
  dn_class          = "mysql.n4.medium.25"
  cn_class          = "polarx.x4.medium.2e"
  dn_node_count     = "2"
  secondary_zone    = "cn-beijing-k"
  tertiary_zone     = "cn-beijing-h"
  vpc_id            = alicloud_vpc.defaultI3SPrf.id
  resource_group_id = alicloud_resource_manager_resource_group.default4m883T.id
}
```

## Argument Reference

The following arguments are supported:
* `cn_class` - (Optional, ForceNew) Compute node specifications.
* `cn_node_count` - (Optional) Number of computing nodes.
* `dn_class` - (Optional, ForceNew) Storage node specifications.
* `dn_node_count` - (Optional) The number of storage nodes.
* `is_read_db_instance` - (Optional) Whether the instance is read-only.
  - **true**: Yes
  - **false**: No.
* `primary_zone` - (Optional) Primary Availability Zone.
* `resource_group_id` - (Optional, Computed) The resource group ID can be empty. This parameter is not supported for the time being.
* `secondary_zone` - (Optional) Secondary availability zone.
* `tertiary_zone` - (Optional) Third Availability Zone.
* `topology_type` - (Required) Topology type:
  - **3azones**: three available areas;
  - **1azone**: Single zone.
* `vswitch_id` - (Optional, ForceNew) The ID of the virtual switch.
* `vpc_id` - (Optional, ForceNew) The VPC ID.
* `zone` - (Optional, ForceNew) Instance availability zone.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The creation time of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Polardb X Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Polardb X Instance.
* `update` - (Defaults to 5 mins) Used when update the Polardb X Instance.

## Import

DRDS Polardb X Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_drds_polardb_x_instance.example <id>
```