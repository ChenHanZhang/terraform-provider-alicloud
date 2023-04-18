---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_flow_logs"
sidebar_current: "docs-alicloud-datasource-vpc-flow-logs"
description: |-
  Provides a list of Vpc Flow Log owned by an Alibaba Cloud account.
---

# alicloud_vpc_flow_logs

This data source provides Vpc Flow Log available to the user.[What is Flow Log](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
没有资源测试用例，请先通过资源测试用例后再生成示例代码。
```

## Argument Reference

The following arguments are supported:
* `description` - (ForceNew,Optional) The Description of flow log.
* `flow_log_id` - (ForceNew,Optional) The flow log ID.
* `flow_log_name` - (ForceNew,Optional) The flow log name.
* `log_store_name` - (ForceNew,Optional) The log store name.
* `project_name` - (ForceNew,Optional) The project name.
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the VPC belongs.
* `resource_id` - (ForceNew,Optional) The resource id.
* `resource_type` - (ForceNew,Optional) The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
* `status` - (ForceNew,Optional) The status of  flow log.
* `tags` - (ForceNew,Optional) The tags of PrefixList.See the following `Block Tags`.
* `traffic_type` - (ForceNew,Optional) The traffic type.
* `ids` - (Optional, ForceNew, Computed) A list of Flow Log IDs.
* `flow_log_names` - (Optional, ForceNew) The name of the Flow Log. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

#### Block Tags

The Tags supports the following:
* `tag_key` - (ForceNew,Optional) The key of tag.
* `tag_value` - (ForceNew,Optional) The value of tag.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Flow Log IDs.
* `names` - A list of name of Flow Logs.
* `logs` - A list of Flow Log Entries. Each element contains the following attributes:
  * `aggregation_interval` - Data aggregation interval
  * `business_status` - Business status
  * `create_time` - the time of creation.
  * `description` - The Description of flow log.
  * `flow_log_id` - The flow log ID.
  * `flow_log_name` - The flow log name.
  * `log_store_name` - The log store name.
  * `project_name` - The project name.
  * `resource_group_id` - The ID of the resource group to which the VPC belongs.
  * `resource_id` - The resource id.
  * `resource_type` - The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
  * `status` - The status of  flow log.
  * `tags` - The tags of PrefixList.
    * `tag_key` - The key of tag.
    * `tag_value` - The value of tag.
  * `traffic_path` - 采集的流量路径。取值：    all（默认值）：表示全量采集。     internetGateway：表示公网流量采集。
  * `traffic_type` - The traffic type.
