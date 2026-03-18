---
subcategory: "Alidns"
layout: "alicloud"
page_title: "Alicloud: alicloud_alidns_cloud_gtm_address_pool"
description: |-
  Provides a Alicloud Alidns Cloud Gtm Address Pool resource.
---

# alicloud_alidns_cloud_gtm_address_pool

Provides a Alidns Cloud Gtm Address Pool resource.



For information about Alidns Cloud Gtm Address Pool and how to use it, see [What is Cloud Gtm Address Pool](https://next.api.alibabacloud.com/document/Alidns/2015-01-09/CreateCloudGtmAddressPool).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `address_lb_strategy` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `address_pool_name` - (Optional) The name of the resource
* `address_pool_type` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `enable_status` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `health_judgement` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `remark` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `sequence_lb_strategy_mode` - (Optional) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.

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