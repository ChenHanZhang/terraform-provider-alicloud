---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_traffic_mirror_filters"
sidebar_current: "docs-alicloud-datasource-vpc-traffic-mirror-filters"
description: |-
  Provides a list of Vpc Traffic Mirror Filter owned by an Alibaba Cloud account.
---

# alicloud_vpc_traffic_mirror_filters

This data source provides Vpc Traffic Mirror Filter available to the user.[What is Traffic Mirror Filter](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_traffic_mirror_filters" "default" {
  ids                        = ["${alicloud_vpc_traffic_mirror_filter.default.id}"]
  name_regex                 = alicloud_vpc_traffic_mirror_filter.default.name
  traffic_mirror_filter_name = "TrafficMirrorFilterNameTest"
}

output "alicloud_vpc_traffic_mirror_filter_example_id" {
  value = data.alicloud_vpc_traffic_mirror_filters.default.filters.0.id
}
```

## Argument Reference

The following arguments are supported:
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the VPC belongs.
* `traffic_mirror_filter_name` - (ForceNew,Optional) The name of the TrafficMirrorFilter.
* `ids` - (Optional, ForceNew, Computed) A list of Traffic Mirror Filter IDs.
* `traffic_mirror_filter_names` - (Optional, ForceNew) The name of the Traffic Mirror Filter. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Traffic Mirror Filter IDs.
* `names` - A list of name of Traffic Mirror Filters.
* `filters` - A list of Traffic Mirror Filter Entries. Each element contains the following attributes:
  * `egress_rules` - EgressRules
    * `action` - Action.
    * `destination_cidr_block` - DestinationCidrBlock.
    * `destination_port_range` - DestinationPortRange.
    * `priority` - Priority.
    * `protocol` - Protocol.
    * `source_cidr_block` - SourceCidrBlock.
    * `source_port_range` - SourcePortRange.
    * `traffic_mirror_filter_rule_status` - TrafficMirrorFilterRuleStatus.
  * `ingress_rules` - IngressRules
    * `action` - Action.
    * `destination_cidr_block` - DestinationCidrBlock.
    * `destination_port_range` - DestinationPortRange.
    * `priority` - Priority.
    * `protocol` - Protocol.
    * `source_cidr_block` - SourceCidrBlock.
    * `source_port_range` - SourcePortRange.
    * `traffic_mirror_filter_rule_status` - TrafficMirrorFilterRuleStatus.
  * `resource_group_id` - The ID of the resource group to which the VPC belongs.
  * `status` - The statusID of the resource
  * `tags` - The tags of PrefixList.
    * `tag_key` - The key of tag.
    * `tag_value` - The value of tag.
  * `traffic_mirror_filter_description` - The description of the TrafficMirrorFilter.
  * `traffic_mirror_filter_id` - The first ID of the resource
  * `traffic_mirror_filter_name` - The name of the TrafficMirrorFilter.
