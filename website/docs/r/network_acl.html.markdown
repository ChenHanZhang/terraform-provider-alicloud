---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_network_acl"
sidebar_current: "docs-alicloud-resource-network-acl"
description: |-
  Provides a Alicloud Vpc Network Acl resource.
---

# alicloud_network_acl

Provides a Vpc Network Acl resource.

For information about Vpc Network Acl and how to use it, see [What is Network Acl](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_network_acl" "default" {
  description      = "test"
  vpc_id           = "vpc-bp11lfjeaa57jxr6ovybf"
  network_acl_name = "rdk-test"
  region_id        = "cn-hangzhou"
}
```

## Argument Reference

The following arguments are supported:
* `all` - (Optional) Whether to unbind all tags of the resource. Value:-**true**: untags all resources.-**false** (default): does not remove all tags of the resource.
* `description` - (Optional) Description of network ACL information.
* `egress_acl_entries` - (Computed,Optional) Output direction rule information.See the following `Block EgressAclEntries`.
* `ingress_acl_entries` - (Computed,Optional) Entry direction rule information.See the following `Block IngressAclEntries`.
* `network_acl_name` - (Optional) The name of the network ACL.
* `resources` - (Computed,Optional) The associated resource.See the following `Block Resources`.
* `source_network_acl_id` - (Optional) The ID of the copied network ACL.
* `tags` - (Optional) The tags of VSwitch.See the following `Block Tags`.
* `vpc_id` - (Required,ForceNew) The ID of the associated VPC.

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.

#### Block EgressAclEntries

The EgressAclEntries supports the following:
* `description` - (Optional) Give the description information of the direction rule.
* `destination_cidr_ip` - (Optional) The destination address segment.
* `network_acl_entry_name` - (Optional) The name of the entry for the direction rule.
* `policy` - (Optional) The  authorization policy.
* `port` - (Optional) Destination port range.
* `protocol` - (Optional) Transport  layer protocol.

#### Block IngressAclEntries

The IngressAclEntries supports the following:
* `description` - (Optional) Description of the entry direction rule.
* `network_acl_entry_name` - (Optional) The name of the entry direction rule entry.
* `policy` - (Optional) The authorization policy.
* `port` - (Optional) Source port range.
* `protocol` - (Optional) Transport layer protocol.
* `source_cidr_ip` - (Optional) The source address field.

#### Block Resources

The Resources supports the following:
* `resource_id` - (Required) The ID of the associated resource.
* `resource_type` - (Required) The type of the associated resource.

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The tag key of VSwitch.
* `tag_value` - (Optional) The tag value of VSwitch.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `create_time` - The creation time of the resource
* `egress_acl_entries` - Output direction rule information.
* `ingress_acl_entries` - Entry direction rule information.
* `network_acl_id` - The first ID of the resource
* `resources` - The associated resource.
  * `status` - The state of the associated resource.
* `status` - The state of the network ACL.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Network Acl.
* `delete` - (Defaults to 5 mins) Used when delete the Network Acl.
* `update` - (Defaults to 5 mins) Used when update the Network Acl.

## Import

Vpc Network Acl can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_network_acl.example 
```