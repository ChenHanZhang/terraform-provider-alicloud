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

-> **NOTE:** Available since v1.211.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "defaultlKripL" {
  vpc_name = var.name

}

resource "alicloud_vswitch" "default0amZnE" {
  vpc_id       = alicloud_vpc.defaultlKripL.id
  zone_id      = data.alicloud_zones.default.zones.0.id
  cidr_block   = "172.16.0.0/24"
  vswitch_name = var.name

}

resource "alicloud_resource_manager_resource_group" "defaultVUBk6q" {
  display_name        = "terraform-1azone-test-1"
  resource_group_name = var.name

}

resource "alicloud_resource_manager_resource_group" "default9GknYO" {
  display_name        = "terraform-1azone-test--1"
  resource_group_name = var.name

}


resource "alicloud_drds_polardb_x_instance" "default" {
  topology_type            = "3azones"
  vswitch_id               = alicloud_vswitch.default0amZnE.id
  primary_zone             = "cn-beijing-f"
  cn_node_count            = "2"
  dn_class                 = "mysql.n4.medium.25"
  cn_class                 = "polarx.x4.medium.2e"
  dn_node_count            = "2"
  vpc_id                   = alicloud_vpc.defaultlKripL.id
  primary_db_instance_name = "null"
  resource_group_id        = alicloud_resource_manager_resource_group.defaultVUBk6q.id
  secondary_zone           = "cn-beijing-k"
  tertiary_zone            = "cn-beijing-h"
}
```

## Argument Reference

The following arguments are supported:
* `cn_class` - (Required, ForceNew) Compute node specifications.
* `cn_node_count` - (Required) Number of computing nodes.
* `dn_class` - (Required, ForceNew) Storage node specifications.
* `dn_node_count` - (Required) The number of storage nodes.
* `is_read_db_instance` - (Optional) Whether the instance is read-only.
  - **true**: Yes
  - **false**: No.
* `primary_db_instance_name` - (Optional) If the instance is a read-only instance, you must specify the primary instance.
* `primary_zone` - (Required, ForceNew) Primary Availability Zone.
* `resource_group_id` - (Optional, Computed) The resource group ID can be empty. This parameter is not supported for the time being.
* `secondary_zone` - (Optional, ForceNew) Secondary availability zone.
* `tertiary_zone` - (Optional, ForceNew) Third Availability Zone.
* `topology_type` - (Required, ForceNew) Topology type:
  - **3azones**: three available areas;
  - **1azone**: Single zone.
* `vswitch_id` - (Required, ForceNew) The ID of the virtual switch.
* `vpc_id` - (Required, ForceNew) The VPC ID.
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