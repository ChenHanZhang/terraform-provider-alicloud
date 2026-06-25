---
subcategory: "Distributed Relational Database Service (DRDS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_drds_polardbx_instance"
description: |-
  Provides a Alicloud Distributed Relational Database Service (DRDS) Polardbx Instance resource.
---

# alicloud_drds_polardbx_instance

Provides a Distributed Relational Database Service (DRDS) Polardbx Instance resource.

PolarDB-X database instance.

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
* `cn_class` - (Required) Compute node specifications:
  - **polarx.x4.medium.2e**: 2 cores, 8 GB
  - **polarx.x4.large.2e**: 4 cores, 16 GB
  - **polarx.x8.large.2e**: 4 cores, 32 GB
  - **polarx.x4.xlarge.2e**: 8 cores, 32 GB
  - **polarx.x8.xlarge.2e**: 8 cores, 64 GB
  - **polarx.x4.2xlarge.2e**: 16 cores, 64 GB
  - **polarx.x8.2xlarge.2e**: 16 cores, 128 GB
  - **polarx.x4.4xlarge.2e**: 32 cores, 128 GB
  - **polarx.x8.4xlarge.2e**: 32 cores, 256 GB
  - **polarx.st.8xlarge.2e**: 60 cores, 470 GB
  - **polarx.st.12xlarge.2e**: 90 cores, 720 GB
* `cn_node_count` - (Required, Int) Number of compute nodes.
* `description` - (Optional, Available since v1.268.0) Database description.
* `dn_class` - (Required) Storage node specifications:
  - **mysql.n4.medium.25**: 2 cores, 8 GB  
  - **mysql.n4.large.25**: 4 cores, 16 GB  
  - **mysql.x8.large.25**: 4 cores, 32 GB  
  - **mysql.n4.xlarge.25**: 8 cores, 32 GB  
  - **mysql.x8.xlarge.25**: 8 cores, 64 GB  
  - **mysql.n4.2xlarge.25**: 16 cores, 64 GB  
  - **mysql.x8.2xlarge.25**: 16 cores, 128 GB  
  - **mysql.x4.4xlarge.25**: 32 cores, 128 GB  
  - **mysql.x8.4xlarge.25**: 32 cores, 256 GB  
  - **mysql.st.8xlarge.25**: 60 cores, 470 GB  
  - **mysql.st.12xlarge.25**: 90 cores, 720 GB.  
* `dn_node_count` - (Required, Int) Number of storage nodes.  
* `dn_storage_space` - (Optional, Available since v1.283.0) Disk space size of the storage node.
* `engine_version` - (Optional, ForceNew, Computed, Available since v1.268.0) The MySQL engine version, which can be 5.7 or 8.0.
* `is_read_db_instance` - (Optional, Available since v1.268.0) Indicates whether the instance is a read-only instance.
  - `true`: Yes  
  - `false`: No.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `primary_db_instance_name` - (Optional, Available since v1.268.0) If this is a read-only instance, you must specify the primary instance.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `primary_zone` - (Required, ForceNew) Primary zone.  
* `resource_group_id` - (Optional, Computed) Resource group ID. This parameter can be empty and is currently unsupported.
* `secondary_zone` - (Optional, ForceNew) Secondary zone.
* `specified_dnscale` - (Optional, Available since v1.283.0, Deprecated since v1.283.0) SpecifiedDNScale  
* `specified_dnspec_map_json` - (Optional, Available since v1.283.0) SpecifiedDNSpecMapJson
* `storage_type` - (Optional, Computed, Available since v1.283.0) Storage type:  
  - Local disk: custom_local_ssd  
  - Cloud disk: cloud_auto
* `switch_time_mode` - (Optional, Available since v1.283.0) Switch mode:
  - 0: Switch immediately.
  - 1: Switch during maintenance window.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `tertiary_zone` - (Optional, ForceNew) Third availability zone.
* `topology_type` - (Required, ForceNew) Topology type:
  - `3azones`: Three availability zones.
  - `1azone`: Single availability zone.
* `vswitch_id` - (Required, ForceNew) Virtual switch ID.
* `vpc_id` - (Required, ForceNew) VPC ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `db_node_class` - The node specification of the instance.
* `db_node_count` - Number of instance nodes.
* `status` - The status of the instance.

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