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

resource "alicloud_resource_manager_resource_group" "group" {
  display_name        = "terraform-test-1zone"
  resource_group_name = var.name
}

resource "alicloud_vpc" "vpc" {
  vpc_name   = "${var.name}1"
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vsw" {
  vpc_id     = alicloud_vpc.vpc.id
  zone_id    = data.alicloud_zones.default.zones.0.id
  cidr_block = "172.16.0.0/24"
}


resource "alicloud_drds_polardb_x_instance" "default" {
  topology_type            = "1azone"
  engine_version           = "5.7"
  zone                     = "cn-beijing-h"
  vswitch_id               = alicloud_vswitch.vsw.id
  primary_db_instance_name = "null"
  db_node_count            = 2
  vpc_id                   = alicloud_vpc.vpc.id
  network_type             = "VPC"
  primary_zone             = "cn-beijing-h"
  payment_type             = "Postpaid"
  db_node_class            = "polarx.x4.medium.2e"
  resource_group_id        = alicloud_resource_manager_resource_group.group.id
}
```

## Argument Reference

The following arguments are supported:
* `auto_renew` - (Optional) Whether to automatically renew. The default value is true.
  - **true**: Yes
  - **false**: No.
* `cn_node_count` - (Optional) Number of computing nodes.
* `db_node_class` - (Required, ForceNew) Node specifications:
  - **polarx.x4.medium.2e**:2 cores 8g
  - **polarx.x4.large.2e**:4 core 16g
  - **polarx.x8.large.2e**:4 core 32g
  - **polarx.x4.xlarge.2e**:8 cores 32g
  - **polarx.x8.xlarge.2e**:8 cores 64g
  - **polarx.x4.2xlarge.2e**:16 cores 64g
  - **polarx.x8.2xlarge.2e**:16 cores 128g
  - **polarx.x4.4xlarge.2e**:32 core 128g
  - **polarx.x8.4xlarge.2e**:32 cores 256G
  - **polarx.st.8xlarge.2e**:60 cores 470g
  - **polarx.st.12xlarge.2e**:90 core 720g.
* `db_node_count` - (Required) The number of instance nodes. The minimum number is 2.
* `dn_node_count` - (Optional) The number of storage nodes.
* `engine_version` - (Optional) Fixed as 2.0 and cannot be changed.
* `is_read_db_instance` - (Optional) Whether the instance is read-only.
  - **true**: Yes
  - **false**: No.
* `network_type` - (Optional, ForceNew) The network type. Only the VPC network is supported.
* `payment_type` - (Required, ForceNew) The payment type of the resource.
* `period` - (Optional) The charging cycle, Year and Month can only be selected, and Hour is selected by default.
* `primary_db_instance_name` - (Optional) If the instance is a read-only instance, you must specify the primary instance.
* `primary_zone` - (Optional) Primary Availability Zone.
* `resource_group_id` - (Optional, Computed) The resource group ID can be empty. This parameter is not supported for the time being.
* `secondary_zone` - (Optional) Secondary availability zone.
* `tertiary_zone` - (Optional) Third Availability Zone.
* `topology_type` - (Optional) Topology type:
  - **3azones**: three available areas;
  - **1azone**: Single zone.
* `used_time` - (Optional) Prepaid, you can choose to buy a few months or years.When> Period is set to Year, the supported values of this parameter are 1, 2, and 3.
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