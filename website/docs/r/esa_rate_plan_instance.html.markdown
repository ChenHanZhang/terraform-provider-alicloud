---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_rate_plan_instance"
description: |-
  Provides a Alicloud ESA Rate Plan Instance resource.
---

# alicloud_esa_rate_plan_instance

Provides a ESA Rate Plan Instance resource.



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

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `auto_renew` - (Optional) Whether auto-renewal is enabled:  
true: Auto-renewal is enabled.  
false: Auto-renewal is disabled.  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `coverage` - (Optional) Acceleration region:  
domestic: Mainland China only.  
global: Global.  
overseas: Global (excluding Mainland China).  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `max_burst_gbps` - (Optional, Available since v1.271.0) The maximum domestic protection bandwidth for DDoS instances, measured in Gbps.
* `payment_type` - (Optional, Computed) Resource property field indicating the payment type.  
* `period` - (Optional, Int) Subscription period (in months).  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `plan_name` - (Optional) Plan name.

For China site accounts:  
  - `entranceplan`: Free Edition  
  - `basic`: Basic Edition  
  - `medium`: Standard Edition  
  - `high`: Advanced Edition  

For International site accounts:  
  - `entranceplan_intl`: Entrance Edition  
  - `basicplan_intl`: Pro Edition  
  - `vipplan_intl`: Premium Edition  
* `type` - (Optional) Site access type:  
NS: NS access.  
CNAME: CNAME access.  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The purchase time of the rate plan instance.
* `instance_status` - renewing: Renewing.
* `status` - A resource attribute field that indicates the resource status.

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