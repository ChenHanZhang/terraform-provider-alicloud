---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_instance"
description: |-
  Provides a Alicloud Cloud Firewall Instance resource.
---

# alicloud_cloud_firewall_instance

Provides a Cloud Firewall Instance resource.

Cloud Firewall Instance.

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
* `auto_asset_protection` - (Optional, Computed, Available since v1.283.0) Specifies whether to enable automatic traffic redirection. Valid values: - `true`: Enabled. - `false`: Disabled.
* `cfw_bandwidth` - (Optional, Computed, Int, Available since v1.287.0) Firewall extended bandwidth
* `cfw_elastic_traffic` - (Optional, Available since v1.287.0) Elastic traffic processing capacity

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `cfw_endpoint` - (Optional, Int, Available since v1.287.0) The number of additional Cloud Firewall instances.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `cfw_log` - (Optional) The status of log delivery. Valid values:
  - `true`: Enabled.
  - `false`: Disabled.
* `cfw_log_storage` - (Optional, Int) The log storage capacity.
* `cfw_ndlp_enable` - (Optional, Available since v1.287.0) Sensitive data leakage detection

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `immediately_release` - (Optional, Int, Available since v1.287.0) Required for unsubscription scenarios. 1: Indicates immediate release. 0: Indicates stopping the instance first based on the stop policy. Only supported by specific products. The default value is immediate release.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `modify_type` - (Optional) The configuration change type. Valid values:
  - Upgrade: Upgrade.
  - Downgrade: Downgrade.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `payment_type` - (Required) The billing method. Valid values:
  - Subscription: subscription.
  - PayAsYouGo: pay-as-you-go.
* `period` - (Optional, Int) The subscription period. Unit: months. For yearly subscriptions, enter a multiple of 12.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `product_code` - (Required, Available since v1.287.0) The product code. You can call the `QueryProductList` operation to query the product code.

-> **NOTE:**  If a region is specified, this parameter cannot be empty.

* `product_type` - (Required, ForceNew, Available since v1.287.0) The product type.
* `renewal_duration` - (Optional, Computed, Int) The auto-renewal period, in months.

-> **NOTE:**  Required when `RenewalStatus` is set to `AutoRenewal`.

* `renewal_duration_unit` - (Optional, Computed) The unit of the auto-renewal period. Valid values:
  - M: month.
  - Y: year.

-> **NOTE:**  Required when RenewalStatus is set to AutoRenewal.

* `renewal_status` - (Optional, Computed) The auto-renewal status. Valid values:
  - AutoRenewal: Auto-renewal is enabled.
  - ManualRenewal: Manual renewal is used.

Default value: ManualRenewal.
* `sdl` - (Optional, Available since v1.287.0) The activation status of the data leakage prevention feature.
* `spec` - (Required, ForceNew) The edition of the Cloud Firewall instance. Valid values:
  - `2`: Premium Edition
  - `3`: Enterprise Edition
  - `4`: Ultimate Edition
  - `10`: Pay-As-You-Go Edition.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time.
* `end_time` - The end time.
* `release_time` - The release time.
* `status` - The production status of the Cloud Firewall instance.
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