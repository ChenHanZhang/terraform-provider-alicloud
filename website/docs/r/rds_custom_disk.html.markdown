---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_custom_disk"
description: |-
  Provides a Alicloud RDS Custom Disk resource.
---

# alicloud_rds_custom_disk

Provides a RDS Custom Disk resource.

Dedicated host disk for ApsaraDB RDS users.

For information about RDS Custom Disk and how to use it, see [What is Custom Disk](https://next.api.alibabacloud.com/document/Rds/2014-08-15/CreateRCDisk).

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

variable "region_id" {
  default = "cn-beijing"
}


resource "alicloud_rds_custom_disk" "default" {
  description          = "zcc测试用例"
  zone_id              = "cn-beijing-i"
  size                 = "40"
  performance_level    = "PL1"
  instance_charge_type = "Postpaid"
  disk_category        = "cloud_essd"
  disk_name            = "custom_disk_001"
  auto_renew           = false
  period               = "1"
  auto_pay             = true
  period_unit          = "1"
}
```

## Argument Reference

The following arguments are supported:
* `auto_pay` - (Optional) Specifies whether to enable automatic payment. Valid values:
  - `true` (default): Automatic payment is enabled. You must ensure that your account balance is sufficient.
  - `false`: Only an order is generated without charging.



-> **NOTE:**  If your payment method has insufficient funds, you can set this parameter to false. In this case, an unpaid order is generated, and you can log on to the RDS console to complete the payment manually.

-> **NOTE:** .


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `auto_renew` - (Optional) Specifies whether to enable auto-renewal. This parameter applies only when you create a subscription data disk. Valid values:
  - `true`: Enable auto-renewal.
  - `false`: Disable auto-renewal.

-> **NOTE:**  If you purchase the disk on a monthly basis, the auto-renewal period is one month.  

-> **NOTE:**  If you purchase the disk on an annual basis, the auto-renewal period is one year.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `description` - (Optional, ForceNew) The disk description. It must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
* `disk_category` - (Required) The category of the data disk. Valid values:
  - `cloud_efficiency`: Ultra disk.
  - `cloud_ssd`: SSD cloud disk.
  - `cloud_essd`: ESSD cloud disk.
  - `cloud_auto` (default): High-performance cloud disk.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `disk_name` - (Optional, ForceNew) The name of the disk. It must be 2 to 128 characters in length and can contain letters defined in the Unicode standard (including English letters, Chinese characters, and digits). It can also include colons (:), underscores (_), periods (.), or hyphens (-).
* `dry_run` - (Optional) Specifies whether to perform a dry run of the instance creation operation. Valid values:
  - `true`: Performs a dry run without creating the instance. The check includes request parameters, request format, service limits, and resource availability.
  - `false` (default): Sends a normal request and creates the instance directly if all checks pass.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `instance_charge_type` - (Optional) The billing method. Valid values:
  - `Postpaid`: Pay-as-you-go. Disks with this billing method do not need to be attached to an instance. You can optionally attach them during creation to any instance regardless of its billing method.
  - `Prepaid`: Subscription. Disks with this billing method must be attached to a subscription instance. Therefore, you must specify a subscription `InstanceId` (instance ID).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_id` - (Optional, Available since v1.275.0) The ID of the instance to which the disk is attached. If `InstanceChargeType` is `Prepaid`, you must specify the ID of a prepaid instance.  

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `performance_level` - (Optional) The performance level of the disk when it is an ESSD cloud disk. Valid values:
  - `PL0`: A single disk supports up to 10,000 random read/write IOPS.
  - `PL1` (default): A single disk supports up to 50,000 random read/write IOPS.
  - `PL2`: A single disk supports up to 100,000 random read/write IOPS.
  - `PL3`: A single disk supports up to 1,000,000 random read/write IOPS.

For information about how to select an ESSD performance level, see [ESSD Cloud Disks](https://help.aliyun.com/document_detail/2859916.html).
* `period` - (Optional, Int) Reserved parameter. You do not need to specify this parameter.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `period_unit` - (Optional) Reserved parameter. You do not need to specify this parameter.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed) The ID of the resource group to which the disk belongs.
* `size` - (Required, Int) The disk size. Unit: GiB. You must specify a value for this parameter. Valid values:
  - `cloud_efficiency`: 20 to 32,768.
  - `cloud_ssd`: 20 to 32,768.
  - `cloud_auto`: 1 to 65,536.
  - `cloud_essd`: The valid range depends on the `PerformanceLevel` value.
  - PL0: 1 to 65,536.
  - PL1: 20 to 65,536.
  - PL2: 461 to 65,536.
  - PL3: 1,261 to 65,536.

If you specify the `SnapshotId` parameter and the snapshot capacity is greater than the specified `Size` value, the created disk will have the same size as the snapshot. If the snapshot capacity is smaller than the `Size` value, the created disk size will be equal to the `Size` value.
* `snapshot_id` - (Optional) The snapshot used to create the disk.
  - Both RDS Custom snapshots and ECS snapshots (non-shared) are supported.
  - If the capacity of the snapshot specified by `SnapshotId` is greater than the value of `Size`, the created disk uses the snapshot's capacity. If the snapshot capacity is less than the `Size` value, the created disk uses the `Size` value.
  - Creating ephemeral local disks from snapshots is not supported.
  - Snapshots created on or before July 15, 2013 cannot be used to create disks.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `tags` - (Optional, Map, Available since v1.275.0) The list of tags.
* `type` - (Optional) The method used to resize the disk. Valid values:
  - `offline` (default): Offline resizing. After resizing, you must restart the instance for the change to take effect.
  - `online`: Online resizing. The resizing takes effect without restarting the instance.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `zone_id` - (Required, ForceNew) The zone ID.  
This parameter is required when the `InstanceId` (the ID of the instance to which the disk is attached) is not specified.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `region_id` - The region ID.
* `status` - Disk status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 6 mins) Used when create the Custom Disk.
* `delete` - (Defaults to 5 mins) Used when delete the Custom Disk.
* `update` - (Defaults to 10 mins) Used when update the Custom Disk.

## Import

RDS Custom Disk can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_custom_disk.example <disk_id>
```