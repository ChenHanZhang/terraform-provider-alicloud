---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_traffic_mirror_sessions"
sidebar_current: "docs-alicloud-datasource-vpc-traffic-mirror-sessions"
description: |-
  Provides a list of Vpc Traffic Mirror Session owned by an Alibaba Cloud account.
---

# alicloud_vpc_traffic_mirror_sessions

This data source provides Vpc Traffic Mirror Session available to the user.[What is Traffic Mirror Session](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_traffic_mirror_sessions" "default" {
  ids                         = ["${alicloud_vpc_traffic_mirror_session.default.id}"]
  name_regex                  = alicloud_vpc_traffic_mirror_session.default.name
  enabled                     = false
  priority                    = 111
  resource_group_id           = "rg-acfm3w3jp6ypy4i"
  traffic_mirror_filter_id    = "tmf-bp1y9l7pmmx0ji994cfha"
  traffic_mirror_session_name = "rmc-test"
  traffic_mirror_source_ids = [
  "eni-bp1db3ueezr53l1hy827"]
  traffic_mirror_target_id = "eni-bp12l2j2g12gpinep803"
  virtual_network_id       = 1
}

output "alicloud_vpc_traffic_mirror_session_example_id" {
  value = data.alicloud_vpc_traffic_mirror_sessions.default.sessions.0.id
}
```

## Argument Reference

The following arguments are supported:
* `enabled` - (ForceNew,Optional) Enabled
* `priority` - (ForceNew,Optional) Priority
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the VPC belongs.
* `tags` - (ForceNew,Optional) The tags of PrefixList.See the following `Block Tags`.
* `traffic_mirror_filter_id` - (ForceNew,Optional) TrafficMirrorFilterId
* `traffic_mirror_session_id` - (ForceNew,Optional) The first ID of the resource
* `traffic_mirror_session_name` - (ForceNew,Optional) TrafficMirrorSessionName
* `traffic_mirror_source_ids` - (ForceNew,Optional) TrafficMirrorSourceIds
* `traffic_mirror_target_id` - (ForceNew,Optional) TrafficMirrorTargetId
* `virtual_network_id` - (ForceNew,Optional) VirtualNetworkId
* `ids` - (Optional, ForceNew, Computed) A list of Traffic Mirror Session IDs.
* `traffic_mirror_session_names` - (Optional, ForceNew) The name of the Traffic Mirror Session. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

#### Block Tags

The Tags supports the following:
* `tag_key` - (ForceNew,Optional) The key of tag.
* `tag_value` - (ForceNew,Optional) The value of tag.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Traffic Mirror Session IDs.
* `names` - A list of name of Traffic Mirror Sessions.
* `sessions` - A list of Traffic Mirror Session Entries. Each element contains the following attributes:
  * `enabled` - Enabled
  * `packet_length` - PacketLength
  * `priority` - Priority
  * `resource_group_id` - The ID of the resource group to which the VPC belongs.
  * `status` - The status of the resource
  * `tags` - The tags of PrefixList.
    * `tag_key` - The key of tag.
    * `tag_value` - The value of tag.
  * `traffic_mirror_filter_id` - TrafficMirrorFilterId
  * `traffic_mirror_session_business_status` - TrafficMirrorSessionBusinessStatus
  * `traffic_mirror_session_description` - TrafficMirrorSessionDescription
  * `traffic_mirror_session_id` - The first ID of the resource
  * `traffic_mirror_session_name` - TrafficMirrorSessionName
  * `traffic_mirror_source_ids` - TrafficMirrorSourceIds
  * `traffic_mirror_target_id` - TrafficMirrorTargetId
  * `traffic_mirror_target_type` - TrafficMirrorTargetType
  * `virtual_network_id` - VirtualNetworkId
