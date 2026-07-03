---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_custom_snapshot"
description: |-
  Provides a Alicloud RDS Custom Snapshot resource.
---

# alicloud_rds_custom_snapshot

Provides a RDS Custom Snapshot resource.

RDS user disk snapshot.

For information about RDS Custom Snapshot and how to use it, see [What is Custom Snapshot](https://next.api.alibabacloud.com/document/Rds/2014-08-15/CreateRCSnapshot).

-> **NOTE:** Available since v1.285.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_vpc" "resourceCase_20260604_nXaKgv" {
  cidr_block = "172.16.0.0/16"
  vpc_name   = "ran_vpc1"
}

resource "alicloud_vswitch" "custom_vsw" {
  vpc_id       = alicloud_vpc.resourceCase_20260604_nXaKgv.id
  zone_id      = "cn-beijing-i"
  cidr_block   = "172.16.0.0/23"
  vswitch_name = "rante_vsw1"
}

resource "alicloud_security_group" "custom_sg" {
  security_group_name = "sg_name"
  vpc_id              = alicloud_vpc.resourceCase_20260604_nXaKgv.id
}

resource "alicloud_rds_custom_disk" "customdisk_disk" {
  description          = "cratesnapshotuse"
  zone_id              = "cn-beijing-i"
  size                 = "20"
  instance_charge_type = "Postpaid"
  performance_level    = "PL1"
  disk_category        = "cloud_essd"
  disk_name            = "snapshotuse"
}

resource "alicloud_rds_custom" "custom_instance" {
  description        = "ran资源用名称0604快照用01"
  zone_id            = "cn-beijing-i"
  auto_renew         = false
  vswitch_id         = alicloud_vswitch.custom_vsw.id
  dry_run            = false
  auto_pay           = true
  security_group_ids = ["${alicloud_security_group.custom_sg.id}"]
  system_disk {
    category          = "cloud_essd"
    size              = "40"
    performance_level = "PL1"
  }
  instance_name = "exampleransnapshotuse"
  data_disk {
    category          = "cloud_essd"
    performance_level = "PL1"
    size              = "40"
  }
  instance_type = "mysql.xa2.xlarge.8cm"
  host_name     = "ranhostexample"
  spot_strategy = "NoSpot"
  period_unit   = "Month"
}

resource "alicloud_rds_custom_disk_attachment" "customdiskattachment" {
  instance_id          = alicloud_rds_custom.custom_instance.id
  delete_with_instance = true
  disk_id              = alicloud_rds_custom_disk.customdisk_disk.id
}


resource "alicloud_rds_custom_snapshot" "default" {
  description    = "创建快照使用ran用例06041"
  zone_id        = alicloud_rds_custom_disk.customdisk_disk.zone_id
  retention_days = "15"
  disk_id        = alicloud_rds_custom_disk.customdisk_disk.id
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional, ForceNew) The description of the snapshot. It must be 2 to 256 characters in length, and cannot start with `http://` or `https://`.
Default value: empty.
* `disk_id` - (Required, ForceNew) The ID of the specified cloud disk.
* `force` - (Optional) Specifies whether to forcibly delete a snapshot that has already been used to create a cloud disk. Valid values:
  - `true`: Force deletion. After forced deletion, the disk cannot be reinitialized.
  - `false` (default): Do not force deletion.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `instant_access` - (Optional, ForceNew) This parameter has been deprecated and does not need to be specified.
* `instant_access_retention_days` - (Optional, Int) This parameter is deprecated and does not need to be specified.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed) The resource group ID. You can call ListResourceGroups to obtain it.
* `retention_days` - (Optional, Int) Specifies the retention period for the snapshot, in days. After the retention period expires, the snapshot is automatically released. Valid values: 1 to 65536.  
Default value: Empty, which means the snapshot will not be automatically released.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `status` - (Optional, Computed) A resource property field that indicates the resource status.

-> **NOTE:** This parameter is only evaluated during resource operations. Modifying it in isolation will not trigger any action.

* `tags` - (Optional, Map) Details of the tags.
* `zone_id` - (Optional) This parameter has been deprecated and does not need to be specified.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Custom Snapshot.
* `delete` - (Defaults to 5 mins) Used when delete the Custom Snapshot.
* `update` - (Defaults to 5 mins) Used when update the Custom Snapshot.

## Import

RDS Custom Snapshot can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_custom_snapshot.example <snapshot_id>
```