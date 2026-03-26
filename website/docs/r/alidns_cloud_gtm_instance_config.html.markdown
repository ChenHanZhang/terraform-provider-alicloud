---
subcategory: "Alidns"
layout: "alicloud"
page_title: "Alicloud: alicloud_alidns_cloud_gtm_instance_config"
description: |-
  Provides a Alicloud Alidns Cloud Gtm Instance Config resource.
---

# alicloud_alidns_cloud_gtm_instance_config

Provides a Alidns Cloud Gtm Instance Config resource.

CloudGtm instance configuration.

For information about Alidns Cloud Gtm Instance Config and how to use it, see [What is Cloud Gtm Instance Config](https://next.api.alibabacloud.com/document/Alidns/2015-01-09/CreateCloudGtmInstanceConfig).

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


resource "alicloud_alidns_cloud_gtm_instance_config" "default" {
  address_pool_lb_strategy = "round_robin"
  schedule_rr_type         = "CNAME"
  schedule_zone_name       = "tianxuan.top"
  enable_status            = "disable"
  charge_type              = "postpay"
  schedule_host_name       = "www"
  schedule_zone_mode       = "custom"
  ttl                      = "600"
}
```

## Argument Reference

The following arguments are supported:
* `address_pool_lb_strategy` - (Optional) Address pool load balancing strategy  
* `charge_type` - (Optional) Billing type

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `enable_status` - (Optional) Enable status  
* `instance_id` - (Optional, ForceNew, Computed) Instance ID associated with the instance configuration  
* `remark` - (Optional) Remarks  
* `schedule_host_name` - (Optional) Host record name  
* `schedule_rr_type` - (Optional, ForceNew) Scheduling RR type
* `schedule_zone_mode` - (Optional, ForceNew) Domain access mode, which can be either custom access domain or system-assigned access domain
* `schedule_zone_name` - (Optional) Zone name
* `sequence_lb_strategy_mode` - (Optional) Sequential load balancing strategy used when the address pool load balancing policy is set to sequential
* `ttl` - (Optional, Int) TTL value  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<config_id>:<instance_id>`.
* `config_id` - Resource property field representing the top-level resource ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cloud Gtm Instance Config.
* `delete` - (Defaults to 5 mins) Used when delete the Cloud Gtm Instance Config.
* `update` - (Defaults to 5 mins) Used when update the Cloud Gtm Instance Config.

## Import

Alidns Cloud Gtm Instance Config can be imported using the id, e.g.

```shell
$ terraform import alicloud_alidns_cloud_gtm_instance_config.example <config_id>:<instance_id>
```