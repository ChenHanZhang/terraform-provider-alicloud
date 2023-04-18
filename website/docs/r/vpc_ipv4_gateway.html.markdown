---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv4_gateway"
sidebar_current: "docs-alicloud-resource-vpc-ipv4-gateway"
description: |-
  Provides a Alicloud Vpc Ipv4 Gateway resource.
---

# alicloud_vpc_ipv4_gateway

Provides a Vpc Ipv4 Gateway resource.

For information about Vpc Ipv4 Gateway and how to use it, see [What is Ipv4 Gateway](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_ipv4_gateway" "default" {
  ipv4_gateway_name        = "Ipv4GatewayNameTest"
  vpc_id                   = "vpc-bp15m8cngyup50ukyy3fp"
  region_id                = "cn-hangzhou"
  ipv4_gateway_description = "Ipv4GatewayDescription测试用例"
}
```

## Argument Reference

The following arguments are supported:
* `all` - (Optional) Whether to unbind all tags of the resource. Value:-**true**: untags all resources.-**false** (default): does not remove all tags of the resource.
* `enabled` - (Optional) Enabled
* `ipv4_gateway_description` - (Optional) Description information
* `ipv4_gateway_name` - (Optional) Resource name
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the instance belongs.
* `tags` - (Optional) The tags of VSwitch.See the following `Block Tags`.
* `vpc_id` - (Required,ForceNew) The ID of the VPC associated with the IPv4 Gateway.

The following arguments will be discarded. Please use new fields as soon as possible:

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The tag key of VSwitch.
* `tag_value` - (Optional) The tag value of VSwitch.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `create_time` - The creation time of the resource
* `ipv4_gateway_id` - The resource attribute field that represents the resource level 1 ID.
* `ipv4_gateway_route_table_id` - ID of the route table associated with IPv4 Gateway
* `resource_group_id` - The ID of the resource group to which the instance belongs.
* `status` - The status of the resource

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ipv4 Gateway.
* `delete` - (Defaults to 5 mins) Used when delete the Ipv4 Gateway.
* `update` - (Defaults to 5 mins) Used when update the Ipv4 Gateway.

## Import

Vpc Ipv4 Gateway can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_ipv4_gateway.example 
```