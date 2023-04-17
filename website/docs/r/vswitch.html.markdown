---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vswitch"
sidebar_current: "docs-alicloud-resource-vswitch"
description: |-
  Provides a Alicloud Vpc Vswitch resource.
---

# alicloud_vswitch

Provides a Vpc Vswitch resource.

For information about Vpc Vswitch and how to use it, see [What is Vswitch](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vswitch" "default" {
  description     = "test"
  zone_id         = "cn-hangzhou-j"
  vpc_id          = "vpc-bp1vzgj2t30917m8tlfwd"
  cidr_block      = "10.50.0.0/24"
  vswitch_name    = "Rdk-test"
  region_id       = "cn-hangzhou"
  ipv6_cidr_block = "12"
}
```

## Argument Reference

The following arguments are supported:
* `all` - (Optional) Whether to unbind all tags of the resource. Value:-**true**: untags all resources.-**false** (default): does not remove all tags of the resource.
* `cidr_block` - (Required,ForceNew) The IPv4 CIDR block of the VSwitch.
* `description` - (Optional) The description of VSwitch.
* `enable_ipv6` - (Optional) Whether the IPv6 function is enabled in the switch. Value:-**true**: enables IPv6.-**false** (default): IPv6 is not enabled.
* `ipv6_cidr_block_mask` - (Optional) The IPv6 CIDR block of the VSwitch.
* `tags` - (Optional) The tags of VSwitch.See the following `Block Tags`.
* `vswitch_name` - (Optional) The name of the VSwitch.
* `vpc_id` - (Required,ForceNew) The VPC ID.
* `zone_id` - (Required,ForceNew) The zone ID  of the resource

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.
* `availability_zone` - Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The tag key of VSwitch.
* `tag_value` - (Optional) The tag value of VSwitch.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `available_ip_address_count` - The number of available IP addresses.
* `create_time` - The creation time of the VSwitch.
* `ipv6_cidr_block` - The IPv6 CIDR block of the VSwitch.
* `is_default` - Indicates whether the VSwitch is a default VSwitch.
* `network_acl_id` - The ID of the network ACL.
* `resource_group_id` - The resource group id of VSwitch.
* `route_table_id` - The route table id
* `status` - The status of the resource
* `vswitch_id` - The ID of the VSwitch.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Vswitch.
* `delete` - (Defaults to 5 mins) Used when delete the Vswitch.
* `update` - (Defaults to 5 mins) Used when update the Vswitch.

## Import

Vpc Vswitch can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_vswitch.example 
```