---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_flow_log"
sidebar_current: "docs-alicloud-resource-vpc-flow-log"
description: |-
  Provides a Alicloud Vpc Flow Log resource.
---

# alicloud_vpc_flow_log

Provides a Vpc Flow Log resource.

For information about Vpc Flow Log and how to use it, see [What is Flow Log](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
没有资源测试用例，请先通过资源测试用例后再生成示例代码。
```

## Argument Reference

The following arguments are supported:
* `aggregation_interval` - (Computed,Optional) Data aggregation interval
* `description` - (Optional) The Description of flow log.
* `flow_log_name` - (Optional) The flow log name.
* `log_store_name` - (Required,ForceNew) The log store name.
* `project_name` - (Required,ForceNew) The project name.
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the VPC belongs.
* `resource_id` - (Required,ForceNew) The resource id.
* `resource_type` - (Required,ForceNew) The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
* `status` - (Computed,Optional) The status of  flow log.
* `tags` - (Optional) The tags of PrefixList.See the following `Block Tags`.
* `traffic_path` - (ForceNew,Computed,Optional) 采集的流量路径。取值：    all（默认值）：表示全量采集。     internetGateway：表示公网流量采集。
* `traffic_type` - (Required,ForceNew) The traffic type.

The following arguments will be discarded. Please use new fields as soon as possible:

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of tag.
* `tag_value` - (Optional) The value of tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `aggregation_interval` - Data aggregation interval
* `business_status` - Business status
* `create_time` - the time of creation.
* `flow_log_id` - The flow log ID.
* `resource_group_id` - The ID of the resource group to which the VPC belongs.
* `status` - The status of  flow log.
* `traffic_path` - 采集的流量路径。取值：    all（默认值）：表示全量采集。     internetGateway：表示公网流量采集。

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Flow Log.
* `delete` - (Defaults to 5 mins) Used when delete the Flow Log.
* `update` - (Defaults to 5 mins) Used when update the Flow Log.

## Import

Vpc Flow Log can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_flow_log.example 
```