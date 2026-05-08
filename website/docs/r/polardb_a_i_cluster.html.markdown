---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_a_i_cluster"
description: |-
  Provides a Alicloud Polardb A I Cluster resource.
---

# alicloud_polardb_a_i_cluster

Provides a Polardb A I Cluster resource.

AI node  .

For information about Polardb A I Cluster and how to use it, see [What is A I Cluster](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateAIDBCluster).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `ack_admin` - (Optional) yes

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ai_node_type` - (Optional, ForceNew) Node type. Valid values:  
  - vnode: Managed by ACK  
  - container: Container accessible via login  
  - maas: Model-as-a-Service  
* `auto_renew` - (Optional) false  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_use_coupon` - (Optional) false

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `db_cluster_description` - (Optional) Cluster description, which supports fuzzy search.
* `db_cluster_id` - (Optional, ForceNew, Computed) Resource attribute field representing the primary resource ID.  
* `db_node_class` - (Required, ForceNew) polar.pg.g4.6xlarge.gu4
* `extension` - (Optional) maas

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `inference_engine` - (Optional) sglang

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `kube_cluster_id` - (Optional, ForceNew) The ID of the ACK cluster.  
* `kube_config` - (Optional) xxx

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `kube_management` - (Optional) self_k8s

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `kubernetes_config` - (Optional) xxx

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `management_mode` - (Optional) ack

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `mode_name` - (Optional) Qwen3-30B-A3B  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `password` - (Optional) xxx

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `pay_type` - (Required, ForceNew) Payment type. Valid values:
  - `Postpaid`: Pay-as-you-go.
  - `Prepaid`: Subscription (annual or monthly billing).
* `payment_type` - (Optional, Computed) A resource property field that represents the payment type.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `period` - (Optional) Month

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `promotion_code` - (Optional) xxx  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed) A resource property field that represents the resource group.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `security_group_id` - (Optional) sg-bp`**********`

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `storage_space` - (Optional, Int) StorageSpace.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `storage_type` - (Required, ForceNew) Storage type.  
* `tags` - (Required, Map) List of tags  
* `time_slices` - (Optional, List) Pay-as-you-go time interval   See [`time_slices`](#time_slices) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `used_time` - (Optional) 1

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `vpc_id` - (Required, ForceNew) The ID of the Virtual Private Cloud (VPC) to use when switching zones.  
* `vswitch_id` - (Required, ForceNew) The ID of the vSwitch.  

-> **NOTE:**  If VPCId is specified, VSwitchId is required.  

* `zone_id` - (Optional, ForceNew) The zone ID of the PolarDB cluster node.  

### `time_slices`

The time_slices supports the following:
* `begin_time` - (Optional, Int) 1758729600
* `end_time` - (Optional, Int) 1758733200  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the A I Cluster.
* `delete` - (Defaults to 5 mins) Used when delete the A I Cluster.
* `update` - (Defaults to 5 mins) Used when update the A I Cluster.

## Import

Polardb A I Cluster can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_a_i_cluster.example <db_cluster_id>
```