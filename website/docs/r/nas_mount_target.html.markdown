---
subcategory: "File Storage (NAS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nas_mount_target"
description: |-
  Provides a Alicloud File Storage (NAS) Mount Target resource.
---

# alicloud_nas_mount_target

Provides a File Storage (NAS) Mount Target resource.

Mount target for a file system.

For information about File Storage (NAS) Mount Target and how to use it, see [What is Mount Target](https://www.alibabacloud.com/help/en/doc-detail/27531.htm).

-> **NOTE:** Available since v1.34.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_nas_zones" "default" {
  file_system_type = "extreme"
}

locals {
  count_size = length(data.alicloud_nas_zones.default.zones)
  zone_id    = data.alicloud_nas_zones.default.zones[local.count_size - 1].zone_id
}

resource "alicloud_vpc" "example" {
  vpc_name   = "terraform-example"
  cidr_block = "172.17.3.0/24"
}

resource "alicloud_vswitch" "example" {
  vswitch_name = alicloud_vpc.example.vpc_name
  cidr_block   = alicloud_vpc.example.cidr_block
  vpc_id       = alicloud_vpc.example.id
  zone_id      = local.zone_id
}

resource "alicloud_nas_file_system" "example" {
  protocol_type    = "NFS"
  storage_type     = "advance"
  file_system_type = "extreme"
  capacity         = "100"
  zone_id          = local.zone_id
}

resource "alicloud_nas_access_group" "example" {
  access_group_name = "access_group_xxx"
  access_group_type = "Vpc"
  description       = "test_access_group"
  file_system_type  = "extreme"
}

resource "alicloud_nas_mount_target" "example" {
  file_system_id    = alicloud_nas_file_system.example.id
  access_group_name = alicloud_nas_access_group.example.access_group_name
  vswitch_id        = alicloud_vswitch.example.id
  vpc_id            = alicloud_vpc.example.id
  network_type      = alicloud_nas_access_group.example.access_group_type
}
```

## Argument Reference

The following arguments are supported:
* `access_group_name` - (Required) The name of the permission group.
This parameter is required when the target file system is General-purpose NAS or Extreme NAS.
Default permission group: DEFAULT_VPC_GROUP_NAME (the default permission group for VPC).
* `access_point_access_only` - (Optional, Available since v1.278.0) Specifies whether the VPC mount target allows access only through access points (APs). This parameter applies only to CPFS file systems for intelligent computing.
* `dual_stack` - (Optional, Available since v1.247.0) Specifies whether to create an IPv6-enabled mount target.

Valid values:
  - true: creates an IPv6-enabled mount target
  - false (default): does not create an IPv6-enabled mount target

-> **NOTE:**  Currently, IPv6 is supported only in mainland China regions for Extreme NAS, and the file system must have IPv6 enabled.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `file_system_id` - (Required, ForceNew) File system ID.
  - General-purpose NAS: 31a8e4****.
  - Extreme NAS: Must start with `extreme-`, for example, extreme-0015****.
  - CPFS: Must start with `cpfs-`, for example, cpfs-125487****.
* `network_type` - (Required, ForceNew, Available since v1.208.1) Network type of the mount target. The value `Vpc` indicates a Virtual Private Cloud (VPC).
* `security_group_id` - (Optional) The ID of the security group.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `status` - (Optional, Computed) The status of the mount target.

Valid values:
  - Active: available
  - Inactive: unavailable

-> **NOTE:**  Currently, only General-purpose NAS supports changing the mount target status.

* `tags` - (Optional, Map, Available since v1.278.0) An array of tags. The array length must be from 1 to 20. If multiple tag objects are included in the array, their keys must be unique.
* `vswitch_id` - (Optional, ForceNew) The ID of the vSwitch.
This field is required and valid only when the network type is Virtual Private Cloud (VPC).  
For example:  
When NetworkType=VPC, VSwitchId is required.
* `vpc_id` - (Optional, ForceNew, Computed) Virtual Private Cloud (VPC) ID.
This field is required and meaningful only when the network type is VPC.  
For example:  
When NetworkType=VPC, VpcId is required.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<file_system_id>:<mount_target_domain>`.
* `mount_target_domain` - IPv4 mount target.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Mount Target.
* `delete` - (Defaults to 5 mins) Used when delete the Mount Target.
* `update` - (Defaults to 5 mins) Used when update the Mount Target.

## Import

File Storage (NAS) Mount Target can be imported using the id, e.g.

```shell
$ terraform import alicloud_nas_mount_target.example <file_system_id>:<mount_target_domain>
```