---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_traffic_mirror_session"
sidebar_current: "docs-alicloud-resource-vpc-traffic-mirror-session"
description: |-
  Provides a Alicloud Vpc Traffic Mirror Session resource.
---

# alicloud_vpc_traffic_mirror_session

Provides a Vpc Traffic Mirror Session resource.

For information about Vpc Traffic Mirror Session and how to use it, see [What is Traffic Mirror Session](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_traffic_mirror_session" "default" {
  traffic_mirror_target_id           = "eni-bp12l2j2g12gpinep803"
  priority                           = 111
  packet_length                      = 1500
  traffic_mirror_session_description = "rmc-test"
  traffic_mirror_session_name        = "rmc-test"
  region_id                          = "cn-hangzhou"
  traffic_mirror_filter_id           = "tmf-bp1y9l7pmmx0ji994cfha"
  traffic_mirror_target_type         = "NetworkInterface"
  virtual_network_id                 = 1
  traffic_mirror_source_ids = [
  "eni-bp1db3ueezr53l1hy827"]
  resource_group_id = "rg-acfm3w3jp6ypy4i"
  enabled           = false
}
```

## Argument Reference

The following arguments are supported:
* `dry_run` - (ForceNew,Optional) Whether to PreCheck only this request, value:-**true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request and directly creates a mirror session after checking.
* `enabled` - (Optional) Enabled
* `packet_length` - (ForceNew,Computed,Optional) PacketLength
* `priority` - (Required) Priority
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the VPC belongs.
* `tags` - (Optional) The tags of PrefixList.See the following `Block Tags`.
* `traffic_mirror_filter_id` - (Required) TrafficMirrorFilterId
* `traffic_mirror_session_description` - (Optional) TrafficMirrorSessionDescription
* `traffic_mirror_session_name` - (Optional) TrafficMirrorSessionName
* `traffic_mirror_source_ids` - (Required) TrafficMirrorSourceIds
* `traffic_mirror_target_id` - (Required) TrafficMirrorTargetId
* `traffic_mirror_target_type` - (Required) TrafficMirrorTargetType
* `virtual_network_id` - (Computed,Optional) VirtualNetworkId

The following arguments will be discarded. Please use new fields as soon as possible:

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of tag.
* `tag_value` - (Optional) The value of tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `packet_length` - PacketLength
* `resource_group_id` - The ID of the resource group to which the VPC belongs.
* `status` - The status of the resource
* `traffic_mirror_session_business_status` - TrafficMirrorSessionBusinessStatus
* `traffic_mirror_session_id` - The first ID of the resource
* `virtual_network_id` - VirtualNetworkId

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Traffic Mirror Session.
* `delete` - (Defaults to 5 mins) Used when delete the Traffic Mirror Session.
* `update` - (Defaults to 5 mins) Used when update the Traffic Mirror Session.

## Import

Vpc Traffic Mirror Session can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_traffic_mirror_session.example 
```