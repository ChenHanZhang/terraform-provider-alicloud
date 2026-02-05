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

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

variable "example_region_id" {
  default = "cn-beijing"
}

variable "example_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.example_zone_id
  vswitch_name = "example_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id = alicloud_vpc.vpcId.id
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}

resource "alicloud_rds_custom" "customItem" {
  amount        = "1"
  vswitch_id    = alicloud_vswitch.vSwitchId.id
  auto_renew    = false
  period        = "1"
  auto_pay      = true
  instance_type = "mysql.x2.xlarge.6cm"
  data_disk {
    category          = "cloud_essd"
    size              = "50"
    performance_level = "PL1"
  }
  status                        = "Running"
  security_group_ids            = ["${alicloud_security_group.securityGroupId.id}"]
  io_optimized                  = "optimized"
  description                   = "diskexample依赖custom"
  key_pair_name                 = alicloud_ecs_key_pair.KeyPairName.id
  zone_id                       = var.example_zone_id
  instance_charge_type          = "Prepaid"
  internet_max_bandwidth_out    = "0"
  image_id                      = "aliyun_2_1903_x64_20G_alibase_20240628.vhd"
  security_enhancement_strategy = "Active"
  period_unit                   = "Month"
  password                      = "jingyiTEST@123"
  system_disk {
    size     = "40"
    category = "cloud_essd"
  }
  host_name     = "1770258687"
  create_mode   = "0"
  spot_strategy = "NoSpot"
}

resource "alicloud_rds_custom_disk" "customdiskItem" {
  description          = "包年disk快照用"
  zone_id              = var.example_zone_id
  size                 = "40"
  instance_charge_type = "Prepaid"
  disk_category        = "cloud_ssd"
  disk_name            = "custom_disk_tosnapshot"
  auto_renew           = false
  period               = "1"
  auto_pay             = true
  period_unit          = "Month"
}


resource "alicloud_rds_custom_snapshot" "default" {
  description                   = "快照tag-带customdisk版本"
  disk_id                       = alicloud_rds_custom_disk.customdiskItem.id
  zone_id                       = var.example_zone_id
  retention_days                = "3"
  instant_access_retention_days = "1"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional, ForceNew) The description of the snapshot. It must be 2 to 256 characters in length and cannot start with 'http:// 'or 'https.
Default value: empty.
* `disk_id` - (Required, ForceNew) The cloud disk ID.
* `force` - (Optional) Whether to forcibly delete the snapshot that has been used to create a cloud disk. Value Description:
  - `true`: force deletion. The disk cannot be reinitialized after a forced delete.
  - `false` (default): Do not force deletion.

-> **NOTE:** This parameter only takes effect when deletion is triggered.

* `instant_access` - (Optional, ForceNew) This parameter is deprecated and does not need to be filled in.
* `instant_access_retention_days` - (Optional, Int) This parameter is deprecated and does not need to be filled in.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `retention_days` - (Optional, Int) Set the snapshot retention time, in days. The snapshot is automatically released after the retention period expires. The value range is 1 to 65536.
Default value: null, indicating that the snapshot will not be automatically released.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `status` - (Optional, Computed) The status of the resource

-> **NOTE:** This parameter only applies during resource creation, update or deletion. If modified in isolation without other property changes, Terraform will not trigger any action.

* `tags` - (Optional, Map) The tag of the resource
* `zone_id` - (Optional) The zone ID  of the resource

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID of the resource.

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