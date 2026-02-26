---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_lake_cache"
description: |-
  Provides a Alicloud Polardb Lake Cache resource.
---

# alicloud_polardb_lake_cache

Provides a Polardb Lake Cache resource.



For information about Polardb Lake Cache and how to use it, see [What is Lake Cache](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreatePolarFs).

-> **NOTE:** Available since v1.272.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `accelerated_storage_space` - (Optional, ForceNew) Accelerated storage space, in GB.
* `accelerating_enable` - (Optional, ForceNew, Computed) Acceleration enablement. Valid values:
  - `ON`: Enabled.
  - `OFF`: Disabled (default).
* `authorized_user_ids` - (Optional) List of authorized account IDs for the cold storage instance, separated by commas (`,`).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_renew` - (Optional) Specifies whether auto-renewal is enabled. Valid values:
  - `true`: Auto-renewal is enabled.  
  - `false`: Auto-renewal is disabled.  

The default value is `false`.

-> **NOTE:**  This parameter takes effect only when the `PayType` parameter is set to `Prepaid`.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_use_coupon` - (Optional) Whether to automatically apply coupons. Valid values:
  - true (default): Apply coupons.
  - false: Do not apply coupons.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `category` - (Optional, ForceNew) AI cache series. The following series are currently supported:
  - `high_performance`: High-performance Edition  
  - `basic`: Basic Edition  
  - `cold`: Cold Storage Edition.
* `custom_bucket_count` - (Optional, Int) Number of custom buckets.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `custom_bucket_path` - (Optional, ForceNew) Custom bucket path.
* `custom_oss_ak` - (Optional, ForceNew) Custom OSS AccessKey.  
* `custom_oss_sk` - (Optional, ForceNew) Custom OSS SK.
* `db_type` - (Optional, ForceNew, Computed) Database ecosystem type. Valid values:  
  - `polardb_mysql`  
  - `polardb_pg`.
* `db_cluster_id` - (Required, ForceNew) The ID of the PolarDB cluster on which the application depends.
* `description` - (Optional, Computed) Description of the AI cache instance.
* `payment_type` - (Required, ForceNew) Payment type. Valid values:  
  - `PayAsYouGo`: Pay-as-you-go  
  - `Subscription`: Subscription (annual or monthly)  
* `period` - (Optional) This parameter is required when the billing type is `Prepaid`. It specifies whether the prepaid cluster is billed yearly or monthly.
  - `Year`: Yearly billing. When yearly or monthly billing is selected, the unit is year.
  - `Month`: Monthly billing. When yearly or monthly billing is selected, the unit is month.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `promotion_code` - (Optional) Coupon code. If not specified, the default coupon is used.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `storage_space` - (Optional, ForceNew, Int) Storage capacity, in GB.
* `storage_type` - (Optional, ForceNew, Computed) Valid values for high-performance storage type are as follows:  
  - `ESSDPL0`  
  - `ESSDPL1`  

Valid values for Basic Edition storage type are as follows:  
  - **city_redundancy (zone-redundant)**  

Valid values for Cold Storage Edition storage type are as follows:  
  - **city_redundancy (zone-redundant)**  
  - **local_redundancy (locally redundant)**.
* `tags` - (Optional, Map) List of tags.  
* `used_time` - (Optional) This parameter is required when the payment type is `Prepaid`.
  - When `Period` is set to `Month`, `UsedTime` must be an integer in the range `[1–9]`.
  - When `Period` is set to `Year`, `UsedTime` must be an integer in the range `[1–3]`.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `vpc_id` - (Optional, ForceNew, Computed) VPC ID.
* `vswitch_id` - (Optional, ForceNew, Computed) vSwitch ID.
* `zone_id` - (Optional, ForceNew, Computed) Zone ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `region_id` - The region ID.
* `status` - AI Cache instance status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Lake Cache.
* `delete` - (Defaults to 5 mins) Used when delete the Lake Cache.
* `update` - (Defaults to 5 mins) Used when update the Lake Cache.

## Import

Polardb Lake Cache can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_lake_cache.example <lake_cache_id>
```