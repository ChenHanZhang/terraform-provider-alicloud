---
subcategory: "RocketMQ"
layout: "alicloud"
page_title: "Alicloud: alicloud_rocketmq_instance_endpoint"
description: |-
  Provides a Alicloud Rocketmq Instance Endpoint resource.
---

# alicloud_rocketmq_instance_endpoint

Provides a Rocketmq Instance Endpoint resource.

Instance endpoint  .

For information about Rocketmq Instance Endpoint and how to use it, see [What is Instance Endpoint](https://next.api.alibabacloud.com/document/RocketMQ/2022-08-01/CreateInstanceEndpoint).

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `instance_id` - (Required, ForceNew) Instance ID
* `security_group_ids` - (Optional, List) List of security groups  
* `vswitch_ids` - (Required, ForceNew, List) List of vSwitches
* `vpc_id` - (Required, ForceNew) VPC ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<id>`.
* `id` - Endpoint ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Instance Endpoint.
* `delete` - (Defaults to 5 mins) Used when delete the Instance Endpoint.
* `update` - (Defaults to 5 mins) Used when update the Instance Endpoint.

## Import

Rocketmq Instance Endpoint can be imported using the id, e.g.

```shell
$ terraform import alicloud_rocketmq_instance_endpoint.example <instance_id>:<id>
```