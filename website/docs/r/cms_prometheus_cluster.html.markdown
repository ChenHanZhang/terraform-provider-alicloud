---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_prometheus_cluster"
description: |-
  Provides a Alicloud Cms Prometheus Cluster resource.
---

# alicloud_cms_prometheus_cluster

Provides a Cms Prometheus Cluster resource.



For information about Cms Prometheus Cluster and how to use it, see [What is Prometheus Cluster](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreatePrometheusCluster).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `hash_label_key` - (Required, ForceNew) Hash tag key
* `payment_type` - (Optional, ForceNew, Computed) The payment type of the resource
* `prometheus_cluster_name` - (Required) Prometheus cluster name
* `replicas` - (Required, ForceNew, Int) Number of copies
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `resource_type` - (Optional) Resource Type

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `storage_duration` - (Optional, Int) Storage Duration (Days)
* `tags` - (Optional, ForceNew, Map) Resource Tags

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Instance creation time, using UTC +0 time, in the format of yyyy-MM-ddTHH:mmZ.
* `region_id` - Area Id.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Prometheus Cluster.
* `delete` - (Defaults to 5 mins) Used when delete the Prometheus Cluster.
* `update` - (Defaults to 5 mins) Used when update the Prometheus Cluster.

## Import

Cms Prometheus Cluster can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_prometheus_cluster.example <prometheus_cluster_name>
```