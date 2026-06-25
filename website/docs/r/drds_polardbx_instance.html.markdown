---
subcategory: "Distributed Relational Database Service (DRDS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_drds_polardbx_instance"
description: |-
  Provides a Alicloud Distributed Relational Database Service (DRDS) Polardbx Instance resource.
---

# alicloud_drds_polardbx_instance

Provides a Distributed Relational Database Service (DRDS) Polardbx Instance resource.

PolarDB-X Database Instance.

For information about Distributed Relational Database Service (DRDS) Polardbx Instance and how to use it, see [What is Polardbx Instance](https://www.alibabacloud.com/help/en/polardb/polardb-for-xscale/api-createdbinstance-1).

-> **NOTE:** Available since v1.211.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}
provider "alicloud" {
  region = "ap-southeast-1"
}
data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}
resource "alicloud_vpc" "example" {
  vpc_name = var.name
}
resource "alicloud_vswitch" "example" {
  vpc_id       = alicloud_vpc.example.id
  zone_id      = data.alicloud_zones.default.zones.0.id
  cidr_block   = "172.16.0.0/24"
  vswitch_name = var.name
}
resource "alicloud_drds_polardbx_instance" "default" {
  topology_type  = "3azones"
  vswitch_id     = alicloud_vswitch.example.id
  primary_zone   = "ap-southeast-1a"
  cn_node_count  = "2"
  dn_class       = "mysql.n4.medium.25"
  cn_class       = "polarx.x4.medium.2e"
  dn_node_count  = "2"
  secondary_zone = "ap-southeast-1b"
  tertiary_zone  = "ap-southeast-1c"
  vpc_id         = alicloud_vpc.example.id
}
```

## Argument Reference

The following arguments are supported:
* `cn_class` - (Required) Compute node specifications.
* `cn_node_count` - (Required, Int) Number of computing nodes.
* `db_instance_name` - (Optional, ForceNew, Available since v1.283.0) DBInstanceName
* `description` - (Optional, Available since v1.268.0) Instance remarks
* `dn_class` - (Required) Storage node specifications.
* `dn_node_count` - (Required, Int) The number of storage nodes.
* `dn_storage_space` - (Optional, Available since v1.283.0) Storage node disk space size
* `engine_version` - (Optional, ForceNew, Computed, Available since v1.268.0) Engine version, default 5.7
* `is_read_db_instance` - (Optional, Available since v1.268.0) Whether the instance is read-only.
  - `true`: Yes
  - `false`: No

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `primary_db_instance_name` - (Optional, Available since v1.268.0) If the instance is a read-only instance, you must specify the primary instance.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `primary_zone` - (Required, ForceNew) Primary Availability Zone.
* `resource_group_id` - (Optional, Computed) The resource group ID can be empty. This parameter is not supported for the time being.
* `secondary_zone` - (Optional, ForceNew) Secondary availability zone.
* `specified_dnscale` - (Optional, ForceNew, Available since v1.283.0) SpecifiedDNScale
* `specified_dnspec_map_json` - (Optional, ForceNew, Available since v1.283.0) SpecifiedDNSpecMapJson
* `storage_type` - (Optional, Computed, Available since v1.283.0) Storage type
  - Local disk: custom_local_ssd
  - Cloud disk: cloud_auto
* `tertiary_zone` - (Optional, ForceNew) Third Availability Zone.
* `topology_type` - (Required, ForceNew) Topology type:
  - `3azones`: three available areas;
  - `1azone`: Single zone.
* `vswitch_id` - (Required, ForceNew) The ID of the virtual switch.
* `vpc_id` - (Required, ForceNew) The VPC ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `db_node_class` - Node specifications:.
* `db_node_count` - The number of instance nodes.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 61 mins) Used when create the Polardbx Instance.
* `delete` - (Defaults to 61 mins) Used when delete the Polardbx Instance.
* `update` - (Defaults to 61 mins) Used when update the Polardbx Instance.

## Import

Distributed Relational Database Service (DRDS) Polardbx Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_drds_polardbx_instance.example <polardbx_instance_id>
```