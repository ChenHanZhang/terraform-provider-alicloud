---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_custom"
description: |-
  Provides a Alicloud RDS Custom resource.
---

# alicloud_rds_custom

Provides a RDS Custom resource.

Dedicated RDS User host.

For information about RDS Custom and how to use it, see [What is Custom](https://next.api.alibabacloud.com/document/Rds/2014-08-15/RunRCInstances).

-> **NOTE:** Available since v1.247.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-beijing"
}

variable "cluster_id" {
  default = "c18c40b2b336840e2b2bbf8ab291758e2"
}

variable "deploymentsetid" {
  default = "ds-2ze78ef5kyj9eveue92m"
}

variable "vswtich-id" {
  default = "example_vswitch"
}

variable "vpc_name" {
  default = "beijing111"
}

variable "example_region_id" {
  default = "cn-beijing"
}

variable "description" {
  default = "ran_1-08_rccreatenodepool_api"
}

variable "example_zone_id" {
  default = "cn-beijing-h"
}

variable "securitygroup_name" {
  default = "rds_custom_init_sg_cn_beijing"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  vpc_name = var.vpc_name
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  zone_id      = var.example_zone_id
  vswitch_name = var.vswtich-id
  cidr_block   = "172.16.5.0/24"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_name = var.securitygroup_name
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


resource "alicloud_rds_custom" "default" {
  amount        = "1"
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
  description                   = var.description
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
  host_name         = "1743386110"
  create_mode       = "0"
  spot_strategy     = "NoSpot"
  vswitch_id        = alicloud_vswitch.vSwitchId.id
  support_case      = "eni"
  deployment_set_id = var.deploymentsetid
  dry_run           = false
}
```

## Argument Reference

The following arguments are supported:
* `amount` - (Optional, Int) Represents the number of instances created

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_pay` - (Optional) Specifies whether to enable the automatic payment feature. Valid values: * `true` (default): enables the feature. You must make sure that your account balance is sufficient. * `false`: disables the feature. An unpaid order is generated. &gt; If your account balance is insufficient, you can set AutoPay to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console. 

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `auto_renew` - (Optional) Specifies whether to enable auto-renewal for the instance. Valid values: * `true` (default) * `false`

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `auto_use_coupon` - (Optional, Available since v1.272.0) Whether to automatically use the coupon, value:
* `true` (default): Yes.
* `false`: No.

-> **NOTE:**  After using the coupon, if the allocation reduction operation is required, the amount deducted by the coupon will not be refunded.


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `create_mode` - (Optional) Whether to allow joining the ACK cluster. When this parameter is set to `1`, the created instance can be added to the ACK cluster through The `AttachRCInstances` API to efficiently manage container applications.
  - `1`: Yes.
  - `0` (default): No.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `data_disk` - (Optional, ForceNew, List) The list of data disks. See [`data_disk`](#data_disk) below.
* `deletion_protection` - (Optional, ForceNew, Available since v1.272.0) Whether the release protection function is enabled. Value:
  - `true`: On
  - `false` (default): Off
* `deployment_set_id` - (Optional, ForceNew) The ID of the deployment set.
* `description` - (Optional, ForceNew) Instance description. It must be 2 to 256 characters in length and cannot start with http:// or https.
* `direction` - (Optional) The type of the change that you want to perform on the instance. Valid values: &gt; This parameter is optional. The system can automatically determine whether the instance change is an upgrade or a downgrade. If you want to specify this parameter, take note of the following items: * `Upgrade` (default): upgrades the instance type. Make sure that your account balance is sufficient. * `Down`: downgrades the instance type. If the new instance type specified by InstanceType has lower specifications than the current instance type, set Direction to Down.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `dry_run` - (Optional) Whether to pre-check the operation of creating an instance. Valid values:
  - `true`: The PreCheck operation is performed without creating an instance. Check items include request parameters, request formats, business restrictions, and inventory.
  - `false` (default): Sends a normal request and directly creates an instance after the check is passed.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `force` - (Optional) Whether to forcibly release the running instance. Value: true/false

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `force_stop` - (Optional) Specifies whether to forcefully stop the instance. Valid values: * `true` * `false` (default)

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `host_name` - (Optional) The instance host name.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `image_id` - (Optional) The ID of the image used by the instance.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_charge_type` - (Optional) The Payment type. Currently, only `Prepaid` (package year and month) types are supported.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_type` - (Required) The new instance type. For more information about the instance types that are supported by RDS Custom instances, see [Instance types of RDS Custom instances](https://www.alibabacloud.com/help/en/doc-detail/2844823.html). 
* `internet_charge_type` - (Optional) Reserved parameters are not supported.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `internet_max_bandwidth_out` - (Optional, Int) The reserved parameter. This parameter is not supported.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `io_optimized` - (Optional) Reserved parameters are not supported.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `key_pair_name` - (Optional) The name of the AccessKey pair. You can specify only one name.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `password` - (Optional) The account and password of the instance.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `period` - (Optional, Int) The subscription duration of the instance. Default value: `1`. 

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `period_unit` - (Optional) The unit of duration of the year-to-month billing method. Value range:
  - `Year`: Year
  - `Month` (default): Month

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `promotion_code` - (Optional, Available since v1.272.0) Coupon Code.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `reboot_time` - (Optional, Available since v1.272.0) The restart time of the instance.
  - When `RebootWhenFinished` is set to `false`, `must` set the restart time within 48 hours.
  - According to the ISO 8601 standard, use UTC +0 time. The format is: 'yyyy-MM-ddTHH:mmZ '.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `reboot_when_finished` - (Optional, Available since v1.272.0) Whether to restart the instance immediately after the change of configuration. Value range:
  - `true` (default): Yes.
  - `false`: No.

-> **NOTE:**  If the instance is in the `paused` status, even if you set 'RebootWhenFinished = true', the instance will remain in the original status and will not be restarted.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `security_enhancement_strategy` - (Optional) The reserved parameter. This parameter is not supported. 

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `security_group_ids` - (Optional, ForceNew, List) Security group list
* `spot_strategy` - (Optional, Available since v1.252.0) The bidding strategy for pay-as-you-go instances. This parameter takes effect when the value of `InstanceChargeType` is set to **PostPaid. Value range:
  - `NoSpot`: normal pay-as-you-go instances.
  - `SpotAsPriceGo`: The system automatically bids and follows the actual price in the current market.

Default value: **NoSpot * *.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `status` - (Optional, Computed) The status of the resource
* `support_case` - (Optional, Available since v1.252.0) Supported scenarios: createMode:supportCase, for example: NATIVE("0", "eni"),RCK("1", "rck"),ACK_EDGE("1", "edge");

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `system_disk` - (Optional, ForceNew, Set) System disk specifications. See [`system_disk`](#system_disk) below.
* `tags` - (Optional, Map) The tag of the resource
* `user_data` - (Optional, ForceNew, Available since v1.272.0) Instance custom data. The maximum size of the original data is 32kB.
Do not pass confidential information, such as passwords and private keys, in clear text. If you really need to pass it in, encrypt it and use Base64 encoding before sending it, and then decrypt it in the instance. The following is an example of converting a script to a Base64 string:
'''
echo - n '#! /bin/sh
echo "Hello World"' | base64 - w 0
'''
* `user_data_in_base64` - (Optional, Available since v1.272.0) Whether the custom data is encoded in Base64 mode.
  - `true`: Yes.
  - `false` (default): No.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `vswitch_id` - (Required, ForceNew) The ID of the virtual switch. The zone in which the vSwitch is located must correspond to the zone ID entered in ZoneId.
The network type InstanceNetworkType must be VPC.
* `zone_id` - (Optional, ForceNew) The zone ID  of the resource

### `data_disk`

The data_disk supports the following:
* `category` - (Optional, ForceNew) Data disk type, value:
  - `cloud_efficiency`: The ultra cloud disk.
  - `cloud_ssd`:SSD cloud disk.
  - `cloud_essd` (default): The ESSD cloud disk.
  - `cloud_auto`: High-Performance Cloud disks.
* `device` - (Optional, ForceNew, Available since v1.272.0) The mount point of the data disk.

-> **NOTE:**  This parameter is only used in full mirror (complete mirror) scenarios. You can set this parameter to the Mount point of the data disk in the full image, and modify the corresponding **DataDisk.Size** and **DataDisk.Category** parameters to modify the type and Size of the data disk in the full image.

* `encrypted` - (Optional, ForceNew, Available since v1.272.0) Whether to encrypt the cloud disk. Value range:
  - `true`: Yes
  - `false` (default): No
* `performance_level` - (Optional, ForceNew) The performance level when the data disk is an ESSD cloud disk. For more information about the performance differences of ESSD cloud disks, see [ESSD cloud disks](~~ 2859916 ~~). Value:
  - `PL0`
  - `PL1` (default)
  - `PL2`
  - `PL3`
* `size` - (Optional, ForceNew, Int) Data disk size, in GiB. Value range:
  - cloud_efficiency:20~32,768.
  - cloud_ssd:20~32,768.
  - cloud_auto:1~65,536.
  - cloud_essd: The specific value range is related to the value of **DataDisk.PerformanceLevel**.
  - PL0:1~65,536.
  - PL1:20~65,536.
  - PL2:461~65,536.
  - PL3:1,261~65,536.

If The **DataDisk.SnapshotId** parameter is specified and the capacity of the corresponding snapshot is greater than the value of **DataDisk.Size**, the Size of the created disk is the same as that of the snapshot. If the snapshot capacity is less than the value of **DataDisk.Size**, the Size of the created cloud disk is set to **DataDisk.Size**.
* `snapshot_id` - (Optional, ForceNew, Available since v1.272.0) The snapshot used to create the data disk.
  - If the snapshot capacity of **DataDisk.SnapshotId** is greater than **DataDisk.Size**, the Size of the created disk is the same as that of the snapshot. If the snapshot capacity is less than the value of **DataDisk.Size**, the Size of the created cloud disk is set to **DataDisk.Size**.
  - Snapshots are not supported for creating elastic temporary disks.
  - Snapshots on or before July 15, 2013 cannot be used to create cloud disks.

### `system_disk`

The system_disk supports the following:
* `category` - (Optional, ForceNew) System disk type, value:
  - `cloud_efficiency`: The ultra cloud disk.
  - `cloud_ssd`:SSD cloud disk.
  - **cloud_essd (default)**: the ESSD cloud disk.
  - `cloud_auto`: High-Performance Cloud disks.
* `performance_level` - (Optional, ForceNew, Available since v1.272.0) The performance level when the system disk is an ESSD cloud disk. For more information about the performance differences of ESSD cloud disks, see [ESSD cloud disks](~~ 2859916 ~~). Value:
  - `PL0`
  - `PL1` (default)
  - `PL2`
  - `PL3`
* `size` - (Optional, ForceNew) System disk size, unit: GiB. The value must be greater than or equal to the image size corresponding to the `ImageId` parameter. Value range:
  - `cloud_efficiency`:20~2048.
  - `cloud_ssd`:20~2048.
  - `cloud_auto`:1 to 2048.
  - `cloud_essd`: The specific value range is related to the value of **SystemDisk.PerformanceLevel**.
  - PL0:1~2048.
  - PL1:20~2048.
  - PL2:461~2048.
  - PL3:1,261~2048.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Custom.
* `delete` - (Defaults to 5 mins) Used when delete the Custom.
* `update` - (Defaults to 7 mins) Used when update the Custom.

## Import

RDS Custom can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_custom.example <instance_id>
```