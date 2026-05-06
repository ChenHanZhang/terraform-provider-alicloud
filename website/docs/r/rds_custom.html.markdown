---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_custom"
description: |-
  Provides a Alicloud RDS Custom resource.
---

# alicloud_rds_custom

Provides a RDS Custom resource.

RDS Dedicated Host for users.

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
* `amount` - (Optional, Int) The number of RDS Custom instances to create. This parameter applies only when you create multiple RDS Custom instances at a time.  
Valid values: `1` to `5`. Default value: `1`.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_pay` - (Optional) Specifies whether to enable auto-payment. Valid values:
  - `true` (default): Enables auto-payment. You must ensure that your account balance is sufficient.
  - `false`: Generates an order without charging you.

-> **NOTE:**  If your payment method has insufficient funds, you can set the AutoPay parameter to false. In this case, an unpaid order is generated, and you can log on to the ApsaraDB RDS console to complete the payment manually.

-> **NOTE:** 


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `auto_renew` - (Optional) Specifies whether the instance is automatically renewed. This parameter applies only when you create a subscription instance. Valid values:  
* `true`: Yes  
* `false`: No  

-> **NOTE:**  * If you purchase the instance on a monthly basis, the auto-renewal period is one month.  

-> **NOTE:**  * If you purchase the instance on an annual basis, the auto-renewal period is one year.  


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `auto_use_coupon` - (Optional, Available since v1.278.0) Specifies whether to automatically apply coupons. Valid values:
* `true` (default): Yes.
* `false`: No.

-> **NOTE:**  If a coupon is applied and you later downgrade the instance configuration, the amount covered by the coupon will not be refunded.


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `create_mode` - (Optional) Specifies whether the instance can be added to an ACK cluster. When this parameter is set to `1`, the created instance can be added to an ACK cluster by using the `AttachRCInstances` API operation, enabling efficient management of containerized applications.
  - `1`: Yes.
  - `0` (default): No.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `data_disk` - (Optional, ForceNew, Computed, List) The list of data disks. See [`data_disk`](#data_disk) below.
* `deletion_protection` - (Optional, Available since v1.278.0) Specifies whether to enable deletion protection. Valid values:  
  - `true`: Enables deletion protection.  
  - `false` (default): Disables deletion protection.  
* `deployment_set_id` - (Optional, ForceNew) Deployment set ID.
* `description` - (Optional) The description of the instance. It must be 2 to 256 characters in length and cannot start with http:// or https://.  
* `direction` - (Optional) The instance configuration change type. Valid values:  

-> **NOTE:**  You do not need to specify this parameter because the system can automatically determine whether to upgrade or downgrade the instance. If you choose to specify it, follow the logic below:  
  - `Up` (default): Upgrade the instance specification. Ensure that your account has sufficient balance.  
  - `Down`: Downgrade the instance specification. Set Direction=Down when the InstanceType value specifies an instance type lower than the current instance type.  

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `dry_run` - (Optional) Specifies whether to perform a dry run of the instance creation operation. Valid values:
  - `true`: Performs a dry run without creating the instance. The check includes request parameters, request format, service limits, and inventory availability.
  - `false` (default): Sends a normal request and creates the instance directly after passing the checks.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `enable_jumbo_frame` - (Optional, Available since v1.278.0) Whether the instance has the Jumbo Frame feature enabled. Parameter value range:
false: Jumbo frames are disabled. The MTU value for all network interfaces under this instance (including the primary and secondary NICs) is 1500.
true: Enables Jumbo Frames. The MTU value for all network interfaces under this instance (including the primary and secondary NICs) is set to 8500.
Default value: false.
Explanation
Only certain instance types from the 8th generation and later support enabling the Jumbo Frame feature. For more information, see ECS Instance MTU.
* `force` - (Optional) Specifies whether to forcibly release a running instance. Valid values:  
  - `Yes`: Force release  
  - `No` (default): Do not force release.  

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `force_stop` - (Optional) Specifies whether to force stop the instance. Valid values:
  - `true`: Force stops the instance.
  - `false` (default): Stops the instance normally.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `host_name` - (Optional) The hostname of the instance.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `image_id` - (Optional) The ID of the image used by the instance.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_charge_type` - (Optional) Billing method. Valid values:
  - `Prepaid`: Subscription.
  - `Postpaid`: Pay-as-you-go.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_name` - (Optional, Available since v1.278.0) The length is 2~128 characters. It can contain uppercase and lowercase letters, Chinese characters, digits, periods (), underscores (), colons (:), or hyphens (-). The default value is the InstanceId of the instance. When you create multiple RdsCustom instances, you can set the ordered instance names in batches, and can contain square brackets ([]) and commas (,). For specific operations, see [Create an RDS Custom instance](https://help.aliyun.com/zh/rds/apsaradb-rds-for-mysql/create-an-rds-custom-instance? spm = a2c4g.11186623.0.0.36 ef7288jg7aZD#00481 f9ba381u).
* `instance_type` - (Required) The target instance type to which you want to change the configuration. For the instance types supported by RDS Custom instances, see [RDS Custom Instance Types](https://help.aliyun.com/document_detail/2844823.html).  
* `internet_charge_type` - (Optional) Reserved parameter. Not supported currently.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `internet_max_bandwidth_out` - (Optional, Int) Maximum public outbound bandwidth for Custom SQL Server instances, in Mbit/s.
Valid values: 0 to 1024. Default value: 0.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `io_optimized` - (Optional) This parameter is reserved and not currently supported.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `key_pair_name` - (Optional) The name of the key pair. Only a single key pair name is supported.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `password` - (Optional) The account password for the instance. It must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Supported special characters include: `()~!@#$%^&*-_+=|{}[]:;',.?/`.  

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `period` - (Optional, Int) The subscription duration of the resource. Default value: `1`.  

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `period_unit` - (Optional) The time unit for subscription billing. Valid values:  
  - `Year`: Year  
  - `Month` (default): Month.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `private_ip_address` - (Optional, ForceNew, Available since v1.278.0) The private IP address of the instance. When setting a private IP address for an ECS instance of the VPC type, you must select the free CIDR block of the vSwitch (VSwitchId).
* `promotion_code` - (Optional, Available since v1.278.0) The coupon code.  

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `reboot` - (Optional, Available since v1.278.0) Whether the instance is restarted. Value Description:
  - `true`: Restart
  - `false` (default): do not restart

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `reboot_time` - (Optional, Available since v1.278.0) The scheduled restart time for the instance.
  - When `RebootWhenFinished` is set to `false`, you `must` specify a restart time within the next 48 hours.
  - The time must follow the ISO 8601 standard in UTC+0 format: `yyyy-MM-ddTHH:mmZ`.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `reboot_when_finished` - (Optional, Available since v1.278.0) Specifies whether to restart the instance immediately after configuration changes are completed. Valid values:
  - `true` (default): Yes.
  - `false`: No.

-> **NOTE:**  If the instance is in the `Stopped` state, it remains in that state even if `RebootWhenFinished=true` is set, and no restart operation is performed.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `resource_group_id` - (Optional, Computed) The resource group ID. You can call ListResourceGroups to obtain it.
* `security_enhancement_strategy` - (Optional) Reserved parameter, currently unsupported.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `security_group_ids` - (Optional, ForceNew, List) The IDs of the security groups to which the instance belongs. Instances within the same security group can access each other. The maximum number of instances that a security group can contain depends on the security group type. For more information, see the "Security Groups" section in [Limits](https://help.aliyun.com/document_detail/25412.html).  

-> **NOTE:**  The SecurityGroupId parameter determines the network type of the instance. For example, if the specified security group uses the Virtual Private Cloud (VPC) network type, the instance will be of the VPC type, and you must also specify the VSwitchId parameter.  

* `spot_strategy` - (Optional, Available since v1.252.0) The spot strategy for pay-as-you-go instances. This parameter takes effect only when the `InstanceChargeType` parameter is set to `PostPaid`. Valid values:  
  - `NoSpot`: A standard pay-as-you-go instance.  
  - `SpotAsPriceGo`: The system automatically bids based on the current market price.  

Default value: `NoSpot`.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `status` - (Optional, Computed) The status of the instance. Valid values:
  - `Pending`: Being created.
  - `Running`: Running.
  - `Starting`: Starting.
  - `Stopping`: Stopping.
  - `Stopped`: Stopped.
* `support_case` - (Optional, Available since v1.252.0) The deployment type of RDS Custom. Valid values:
  - `eni`: Dual ENIs.
  - `edge`: Edge node pool.
  - `share`: VPC.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `system_disk` - (Optional, ForceNew, Set) System disk specifications. See [`system_disk`](#system_disk) below.
* `tags` - (Optional, Map) Details of the queried instances and their tags.
* `user_data` - (Optional, ForceNew, Available since v1.278.0) Custom data for the instance. The raw data can be up to 32 KB in size.  
Do not pass sensitive information, such as passwords and private keys, in plaintext. If you must pass such information, encrypt it first, encode it in Base64, and then transmit it. Decrypt the data inside the instance after transmission. The following example shows how to convert a script into a Base64-encoded string:  
```
echo - n '#!/bin/sh
echo "Hello World"' | base64 - w 0
```  
* `user_data_in_base64` - (Optional, Available since v1.278.0) Specifies whether the custom data is Base64-encoded.
  - `true`: Yes.
  - `false` (default): No.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `vswitch_id` - (Required, ForceNew) The virtual switch ID of the target instance. If you are creating a VPC-type ApsaraDB RDS Custom instance, you must specify the virtual switch ID. In this case, the security group and the virtual switch must belong to the same Virtual Private Cloud (VPC).  

-> **NOTE:**  If you specify the VSwitchId parameter, the ZoneId parameter you set must match the zone where the specified virtual switch is located. Alternatively, you can omit the ZoneId parameter, and the system will automatically select the zone of the specified virtual switch.  

* `zone_id` - (Optional, ForceNew) The zone ID of the instance. You can call the DescribeZones operation to query the list of available zones.  

-> **NOTE:**  If you specify the VSwitchId parameter, the specified ZoneId must match the zone where the vSwitch is located. Alternatively, you can leave ZoneId unspecified, and the system will automatically select the zone of the specified vSwitch.  


### `data_disk`

The data_disk supports the following:
* `category` - (Optional, ForceNew) The type of data disk. Valid values:
  - `cloud_efficiency`: Ultra disk.
  - `cloud_ssd`: SSD cloud disk.
  - `cloud_essd` (default): ESSD cloud disk.
  - `cloud_auto`: High-performance cloud disk.
* `device` - (Optional, ForceNew, Available since v1.278.0) The mount point of the data disk.

-> **NOTE:**  This parameter applies only to full-image (whole-machine image) scenarios. You can set this parameter to the mount point corresponding to the data disk in the full image and modify the **DataDisk.Size** and **DataDisk.Category** parameters to change the disk type and size of the data disk in the full image.

* `encrypted` - (Optional, ForceNew, Available since v1.278.0) Specifies whether to encrypt the cloud disk. Valid values:
  - `true`: Yes.
  - `false` (default): No.
* `performance_level` - (Optional, ForceNew) The performance level of an ESSD cloud disk. For information about performance differences among ESSD cloud disks, see [ESSD Cloud Disks](https://help.aliyun.com/document_detail/2859916.html). Valid values:
  - `PL0`
  - `PL1` (default)
  - `PL2`
  - `PL3`.
* `size` - (Optional, ForceNew, Int) The size of the data disk, in GiB. Valid ranges:
  - `cloud_efficiency`: 20 to 32,768.
  - `cloud_ssd`: 20 to 32,768.
  - `cloud_auto`: 1 to 65,536.
  - `cloud_essd`: The valid range depends on the value of **DataDisk.PerformanceLevel**.
  - `PL0`: 1 to 65,536.
  - `PL1`: 20 to 65,536.
  - `PL2`: 461 to 65,536.
  - `PL3`: 1,261 to 65,536.

If you specify the **DataDisk.SnapshotId** parameter and the snapshot capacity is greater than the value of **DataDisk.Size**, the created cloud disk will have the same size as the snapshot. If the snapshot capacity is less than the value of **DataDisk.Size**, the created cloud disk will have the size specified by **DataDisk.Size**.
* `snapshot_id` - (Optional, ForceNew, Available since v1.278.0) The snapshot used to create the data disk.
  - If the capacity of the snapshot specified by **DataDisk.SnapshotId** is greater than the value of **DataDisk.Size**, the created cloud disk has the same size as the snapshot. If the snapshot capacity is smaller than the value of **DataDisk.Size**, the created cloud disk size equals the value of **DataDisk.Size**.
  - Snapshots cannot be used to create elastic temporary disks.
  - Snapshots created on or before July 15, 2013, cannot be used to create cloud disks.

### `system_disk`

The system_disk supports the following:
* `category` - (Optional, ForceNew) System disk category. Valid values:
  - `cloud_efficiency`: Ultra disk.
  - `cloud_ssd`: SSD cloud disk.
  - `cloud_essd` (default): ESSD cloud disk.
  - `cloud_auto`: High-performance cloud disk.
* `performance_level` - (Optional, ForceNew, Available since v1.278.0) Performance level for ESSD cloud disks. For details about performance differences among ESSD cloud disks, see [ESSD Cloud Disks](https://help.aliyun.com/document_detail/2859916.html). Valid values:
  - `PL0`
  - `PL1` (default)
  - `PL2`
  - `PL3`
* `size` - (Optional, ForceNew) System disk size, in GiB. The value must be greater than or equal to the size of the image specified by the `ImageId` parameter. Valid ranges:
  - `cloud_efficiency`: 20–2048.
  - `cloud_ssd`: 20–2048.
  - `cloud_auto`: 1–2048.
  - `cloud_essd`: The valid range depends on the value of **SystemDisk.PerformanceLevel**.
  - PL0: 1–2048.
  - PL1: 20–2048.
  - PL2: 461–2048.
  - PL3: 1,261–2048.

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