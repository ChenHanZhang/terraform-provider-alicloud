---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_cache_reserve_instance"
description: |-
  Provides a Alicloud ESA Cache Reserve Instance resource.
---

# alicloud_esa_cache_reserve_instance

Provides a ESA Cache Reserve Instance resource.



For information about ESA Cache Reserve Instance and how to use it, see [What is Cache Reserve Instance](https://next.api.alibabacloud.com/document/ESA/2024-09-10/PurchaseCacheReserve).

-> **NOTE:** Available since v1.272.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_esa_cache_reserve_instance" "default" {
  quota_gb     = "10240"
  cr_region    = "CN-beijing"
  auto_renew   = true
  period       = "1"
  payment_type = "Subscription"
  auto_pay     = true
}
```

## Argument Reference

The following arguments are supported:
* `auto_pay` - (Optional) Automatic payment.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `auto_renew` - (Optional) Whether to auto-renew:
  - `true`: Auto-renew.
  - `false`: Do not auto-renew.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `cr_region` - (Optional, ForceNew) Cache holding area
  - `HK`: Hong Kong, China
  - `CN`: Mainland China
* `payment_type` - (Required, ForceNew) Specifies whether to enable auto payment.
* `period` - (Optional, ForceNew, Int) Purchase period (unit: month).
* `quota_gb` - (Optional, Int) Cache retention specification (unit: GB).

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Instance purchase time.
* `status` - The status of the cache reserve instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cache Reserve Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Cache Reserve Instance.
* `update` - (Defaults to 5 mins) Used when update the Cache Reserve Instance.

## Import

ESA Cache Reserve Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_esa_cache_reserve_instance.example <cache_reserve_instance_id>
```