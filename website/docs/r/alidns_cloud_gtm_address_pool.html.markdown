---
subcategory: "Alidns"
layout: "alicloud"
page_title: "Alicloud: alicloud_alidns_cloud_gtm_address_pool"
description: |-
  Provides a Alicloud Alidns Cloud Gtm Address Pool resource.
---

# alicloud_alidns_cloud_gtm_address_pool

Provides a Alidns Cloud Gtm Address Pool resource.

CloudGTM address pool allows you to group multiple IP addresses or domain names into a single logical entity for traffic management and failover purposes.

For information about Alidns Cloud Gtm Address Pool and how to use it, see [What is Cloud Gtm Address Pool](https://next.api.alibabacloud.com/document/Alidns/2015-01-09/CreateCloudGtmAddressPool).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_alidns_cloud_gtm_address_pool" "default" {
  address_pool_name   = "resoure-example-pool-3"
  health_judgement    = "p30_ok"
  address_pool_type   = "domain"
  enable_status       = "enable"
  address_lb_strategy = "round_robin"
}
```

## Argument Reference

The following arguments are supported:
* `address_lb_strategy` - (Optional) The load balancing strategy for addresses in the address pool
* `address_pool_name` - (Optional) Resource property field representing the resource name
* `address_pool_type` - (Optional, ForceNew) The type of the address pool
* `enable_status` - (Optional) The enable status of the address pool
* `health_judgement` - (Optional) Health check criteria for the address pool  
* `remark` - (Optional) Remarks for the address pool
* `sequence_lb_strategy_mode` - (Optional) Sequential load balancing strategy when the address load balancing policy is set to sequential  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the address pool.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cloud Gtm Address Pool.
* `delete` - (Defaults to 5 mins) Used when delete the Cloud Gtm Address Pool.
* `update` - (Defaults to 5 mins) Used when update the Cloud Gtm Address Pool.

## Import

Alidns Cloud Gtm Address Pool can be imported using the id, e.g.

```shell
$ terraform import alicloud_alidns_cloud_gtm_address_pool.example <address_pool_id>
```