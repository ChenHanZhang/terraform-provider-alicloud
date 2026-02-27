---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_custom_disk"
description: |-
  Provides a Alicloud RDS Custom Disk resource.
---

# alicloud_rds_custom_disk

Provides a RDS Custom Disk resource.

RDS User dedicated host disk.

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
  - **true (default)**: automatically completes the payment. Make sure that your account balance is sufficient. 
  - `false`: does not automatically complete the payment. An unpaid order is generated. 
&gt; If your account balance is insufficient, you can set the AutoPay parameter to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console. 

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `auto_renew` - (Optional) Specifies whether to enable auto-renewal. You must specify this parameter only when the data disk uses the subscription billing method. Valid values:
 * `true` 
  - `false` 

&gt; The auto-renewal cycle is one month for a monthly subscription. The auto-renewal cycle is one year for a yearly subscription.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `delete_with_instance` - (Optional, Available since v1.272.0) Whether to release with the instance. Value range of this parameter:
  - true: When the instance is released, the disk is released along with the instance.
  - false: When the instance is released, the disk is reserved and not released.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `description` - (Optional, ForceNew) The disk description. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
* `disk_category` - (Required) The data disk type. Valid values: 
  - `cloud_efficiency`: ultra disk. 
  - `cloud_ssd`: standard SSD 
  - `cloud_essd`: ESSD 
  - `cloud_auto` (default): Premium ESSD

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `disk_name` - (Optional, ForceNew) The disk name. It can be 2 to 128 characters in length. It supports letters in Unicode (including English, Chinese, and numbers). Can contain a colon (:), an underscore (_), a period (.), or a dash (-).
Default value: empty.
* `dry_run` - (Optional) Whether to pre-check the instance creation operation. Valid values:
  - `true`: The PreCheck operation is performed without creating an instance. Check items include request parameters, request formats, business restrictions, and inventory.
  - `false` (default): Sends a normal request and directly creates an instance after the check is passed.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `instance_charge_type` - (Optional) The billing method. Valid values: 
  - `Postpaid`: pay-as-you-go Pay-as-you-go disks do not require to be attached. You can also attach the pay-as-you-go disk to an instance of any billing method based on your business requirements. 
  - `Prepaid`: subscription Subscription disks must be attached to a subscription instance. Set `InstanceId` to the ID of a subscription instance. 

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_id` - (Optional, Available since v1.272.0) The instance ID.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `performance_level` - (Optional) The PL of the disk. Valid values: 
  - `PL1` (default): A single ESSD can deliver up to 50,000 random read/write IOPS. 
  - `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS. 
  - `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
* `period` - (Optional, Int) A reserved parameter. You do not need to specify this parameter. 

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `period_unit` - (Optional) Reserved parameters, no need to fill in.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed) The resource group ID. You can call the ListResourceGroups operation to obtain the resource group ID.
* `size` - (Required, Int) The new disk size. Unit: GiB.
* `snapshot_id` - (Optional) The snapshot that you want to use to create the disk. 
  - The snapshots of RDS Custom instances and the non-shared snapshots of ECS instances are supported. 
  - If the size of the snapshot specified by `SnapshotId` is greater than the value of `Size`, the size of the created disk is equal to the specified snapshot size. If the snapshot size is less than the `Size` value, the size of the created disk is equal to the `Size` value. 
  - You cannot create elastic ephemeral disks from snapshots. * Snapshots that were created on or before July 15, 2013 cannot be used to create disks.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `status` - (Optional, Computed) Disk status. Value Description:
  - In_use: In use.
  - Available: to be mounted.
  - Attaching: Attaching.
  - Detaching: uninstalling.
  - Creating: Creating.
  - ReIniting: Initializing.
* `tags` - (Optional, Map, Available since v1.272.0) The tag of the resource
* `type` - (Optional) The method that you want to use to resize the disk. Valid values: 
  - `offline` (default): resizes disks offline. After you resize a disk offline, you must restart the instance for the resizing operation to take effect. 
  - `online`: resizes disks online. After you resize a disk online, the resizing operation takes effect immediately and you do not need to restart the instance. 

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `zone_id` - (Required, ForceNew) The zone ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `region_id` - The region ID.

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