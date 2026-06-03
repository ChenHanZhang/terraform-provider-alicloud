---
subcategory: "Alikafka Streaming"
layout: "alicloud"
page_title: "Alicloud: alicloud_alikafka_streaming_compute_instance"
description: |-
  Provides a Alicloud Alikafka Streaming Compute Instance resource.
---

# alicloud_alikafka_streaming_compute_instance

Provides a Alikafka Streaming Compute Instance resource.

Stream computing instance resource  .

For information about Alikafka Streaming Compute Instance and how to use it, see [What is Compute Instance](https://next.api.alibabacloud.com/document/AlikafkaStreaming/2026-02-02/StartComputeInstance).

-> **NOTE:** Available since v1.281.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `instance_id` - (Required, ForceNew) A resource attribute field representing the primary resource ID.
* `instance_name` - (Optional) A resource attribute field representing the resource name.
* `selected_zones` - (Optional) The zone where the resource is deployed.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `service_version` - (Optional, ForceNew) Version number of the currently running software or service engine  
* `vswitch_ids` - (Required, ForceNew, List) An array of IDs of associated virtual switches (vSwitches).
* `vpc_id` - (Required, ForceNew) The unique identifier of the Virtual Private Cloud (VPC).

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `charge_type` - Billing type.
* `create_time` - A resource attribute field representing the creation time.
* `cu` - Compute Unit.
* `expire_time` - Expiration time of the resource.
* `order_id` - Order ID.
* `service_status` - Service status.
* `total_jobs` - Total number of jobs accumulated in the current cluster or instance group.
* `total_running_jobs` - The number of tasks currently running.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Compute Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Compute Instance.
* `update` - (Defaults to 5 mins) Used when update the Compute Instance.

## Import

Alikafka Streaming Compute Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_alikafka_streaming_compute_instance.example <instance_id>
```