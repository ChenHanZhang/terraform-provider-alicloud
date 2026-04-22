---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_rate_plan_instance"
description: |-
  Provides a Alicloud ESA Rate Plan Instance resource.
---

# alicloud_esa_rate_plan_instance

Provides a ESA Rate Plan Instance resource.

Rate plan  .

For information about ESA Rate Plan Instance and how to use it, see [What is Rate Plan Instance](https://www.alibabacloud.com/help/en/edge-security-acceleration/esa/product-overview/query-package-information).

-> **NOTE:** Available since v1.234.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_esa_rate_plan_instance" "default" {
  type         = "NS"
  auto_renew   = true
  period       = "1"
  payment_type = "Subscription"
  coverage     = "overseas"
  plan_name    = "basic"
  auto_pay     = true
}
```

## Argument Reference

The following arguments are supported:
* `auto_pay` - (Optional) Whether auto-payment is enabled.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `auto_renew` - (Optional) Specifies whether auto-renewal is enabled:  
true: Auto-renewal is enabled.  
false: Auto-renewal is disabled.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `coverage` - (Optional) Acceleration region:
domestic: Mainland China only.  
global: Global.  
overseas: Global (excluding Mainland China).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `payment_type` - (Optional, ForceNew, Computed) A resource attribute field that represents the payment type.  
* `period` - (Optional, Int) Subscription period (unit: months).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `plan_name` - (Optional) Plan name.

For accounts on the China site:  
  - `entranceplan`: Free  
  - `basic`: Basic  
  - `medium`: Standard  
  - `high`: Advanced  

For accounts on the International site:  
  - `entranceplan_intl`: Entrance  
  - `basicplan_intl`: Pro  
  - `vipplan_intl`: Premium  
* `type` - (Optional) Site access type:
NS: NS access.  
CNAME: CNAME access.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The purchase time of the rate plan instance.
* `instance_status` - renewing: Renewing.
* `status` - A resource property field that represents the resource status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Rate Plan Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Rate Plan Instance.
* `update` - (Defaults to 5 mins) Used when update the Rate Plan Instance.

## Import

ESA Rate Plan Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_esa_rate_plan_instance.example <instance_id>
```