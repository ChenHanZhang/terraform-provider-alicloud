---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_traffic_mirror_filter"
sidebar_current: "docs-alicloud-resource-vpc-traffic-mirror-filter"
description: |-
  Provides a Alicloud Vpc Traffic Mirror Filter resource.
---

# alicloud_vpc_traffic_mirror_filter

Provides a Vpc Traffic Mirror Filter resource.

For information about Vpc Traffic Mirror Filter and how to use it, see [What is Traffic Mirror Filter](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_traffic_mirror_filter" "default" {
  traffic_mirror_filter_description = "流量镜像描述信息1"
  region_id                         = "cn-hangzhou"
  traffic_mirror_filter_name        = "TrafficMirrorFilterNameTest"
}
```

## Argument Reference

The following arguments are supported:
* `egress_rules` - (ForceNew,Computed,Optional) EgressRulesSee the following `Block EgressRules`.
* `ingress_rules` - (ForceNew,Computed,Optional) IngressRulesSee the following `Block IngressRules`.
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the VPC belongs.
* `tags` - (Optional) The tags of PrefixList.See the following `Block Tags`.
* `traffic_mirror_filter_description` - (Optional) The description of the TrafficMirrorFilter.
* `traffic_mirror_filter_name` - (Optional) The name of the TrafficMirrorFilter.

The following arguments will be discarded. Please use new fields as soon as possible:

#### Block EgressRules

The EgressRules supports the following:
* `action` - (Required,ForceNew) Action.
* `destination_cidr_block` - (ForceNew,Optional) DestinationCidrBlock.
* `destination_port_range` - (ForceNew,Optional) DestinationPortRange.
* `priority` - (ForceNew,Optional) Priority.
* `protocol` - (Required,ForceNew) Protocol.
* `source_cidr_block` - (ForceNew,Optional) SourceCidrBlock.
* `source_port_range` - (ForceNew,Optional) SourcePortRange.

#### Block IngressRules

The IngressRules supports the following:
* `action` - (Required,ForceNew) Action.
* `destination_cidr_block` - (ForceNew,Optional) DestinationCidrBlock.
* `destination_port_range` - (ForceNew,Optional) DestinationPortRange.
* `priority` - (ForceNew,Optional) Priority.
* `protocol` - (Required,ForceNew) Protocol.
* `source_cidr_block` - (ForceNew,Optional) SourceCidrBlock.
* `source_port_range` - (ForceNew,Optional) SourcePortRange.

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of tag.
* `tag_value` - (Optional) The value of tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `egress_rules` - EgressRules
  * `traffic_mirror_filter_rule_status` - TrafficMirrorFilterRuleStatus.
* `ingress_rules` - IngressRules
  * `traffic_mirror_filter_rule_status` - TrafficMirrorFilterRuleStatus.
* `resource_group_id` - The ID of the resource group to which the VPC belongs.
* `status` - The statusID of the resource
* `traffic_mirror_filter_id` - The first ID of the resource

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Traffic Mirror Filter.
* `delete` - (Defaults to 5 mins) Used when delete the Traffic Mirror Filter.
* `update` - (Defaults to 5 mins) Used when update the Traffic Mirror Filter.

## Import

Vpc Traffic Mirror Filter can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_traffic_mirror_filter.example 
```