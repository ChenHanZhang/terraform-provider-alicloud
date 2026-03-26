---
subcategory: "ACC"
layout: "alicloud"
page_title: "Alicloud: alicloud_acc_standby_pool"
description: |-
  Provides a Alicloud ACC Standby Pool resource.
---

# alicloud_acc_standby_pool

Provides a ACC Standby Pool resource.

Pre-warmed pool  .

For information about ACC Standby Pool and how to use it, see [What is Standby Pool](https://next.api.alibabacloud.com/document/acc/2024-04-02/CreateStandbyPool).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `capacity` - (Required, Int) Capacity of the pre-warmed pool  
* `instance_resource_configuration_id` - (Required, ForceNew) Pre-warmed instance configuration profile ID  
* `provision_policy` - (Required, ForceNew) Pre-warmed instance provisioning policy  
* `standby_pool_name` - (Required) A resource property field representing the resource name.  
* `standby_state` - (Required, ForceNew) Target state of the pre-warmed instance  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - A resource property field representing the region.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Standby Pool.
* `delete` - (Defaults to 5 mins) Used when delete the Standby Pool.
* `update` - (Defaults to 5 mins) Used when update the Standby Pool.

## Import

ACC Standby Pool can be imported using the id, e.g.

```shell
$ terraform import alicloud_acc_standby_pool.example <standby_pool_id>
```