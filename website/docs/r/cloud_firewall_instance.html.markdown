---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_instance"
description: |-
  Provides a Alicloud Cloud Firewall Instance resource.
---

# alicloud_cloud_firewall_instance

Provides a Cloud Firewall Instance resource.

Cloud Firewall instance
.

For information about Cloud Firewall Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/product/90174.htm).

-> **NOTE:** Available since v1.139.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_cloud_firewall_instance" "PayAsYouGo" {
  payment_type = "PayAsYouGo"
}
```

### Deleting `alicloud_cloud_firewall_instance` or removing it from your configuration

The `alicloud_cloud_firewall_instance` resource allows you to manage  `payment_type = "PayAsYouGo"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `cfw_log` - (Optional, Computed) Log delivery status. Valid values:
  - `true`: Enabled
  - `false`: Disabled.
* `modify_type` - (Optional) Configuration change type. Valid values:
  - Upgrade: Upgrade.
  - Downgrade: Downgrade.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `payment_type` - (Required, ForceNew) Payment type. Valid values:
  - Subscription: Prepaid.
  - PayAsYouGo: Postpaid.
* `period` - (Optional, Int) The subscription duration, in months. For annual billing products, enter a multiple of 12.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `product_code` - (Required, ForceNew, Available since v1.273.0) The product code. You can query product codes by calling the `QueryProductList` operation.
* `product_type` - (Required, ForceNew, Available since v1.273.0) The product type.
* `renewal_duration` - (Optional, Computed, Int) The auto-renewal duration, in months.  

-> **NOTE:**  This parameter is required when `RenewalStatus` is set to `AutoRenewal`.

* `renewal_duration_unit` - (Optional, Computed) Auto-renewal cycle unit. Valid values:
  - M: Month.
  - Y: Year.

-> **NOTE:**  This parameter is required when RenewalStatus is set to AutoRenewal.

* `renewal_status` - (Optional, Computed) The auto-renewal status. Valid values:  
  - AutoRenewal: Auto-renewal is enabled.  
  - ManualRenewal: Renewal must be performed manually.  

Default value: ManualRenewal.
* `sdl` - (Optional, Available since v1.273.0) Data leakage protection activation status.
* `spec` - (Required, ForceNew) The edition of the Cloud Firewall instance. Valid values:
  - `2`: Advanced Edition  
  - `3`: Enterprise Edition  
  - `4`: Ultimate Edition  
  - `10`: Pay-as-you-go Edition

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time.
* `end_time` - The end time.
* `release_time` - Release time.
* `status` - The operational status of the Cloud Firewall instance.
* `user_status` - The status of the Cloud Firewall instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Instance.
* `update` - (Defaults to 5 mins) Used when update the Instance.

## Import

Cloud Firewall Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_instance.example <instance_id>
```