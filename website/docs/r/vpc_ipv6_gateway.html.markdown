---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv6_gateway"
sidebar_current: "docs-alicloud-resource-vpc-ipv6-gateway"
description: |-
  Provides a Alicloud Vpc Ipv6 Gateway resource.
---

# alicloud_vpc_ipv6_gateway

Provides a Vpc Ipv6 Gateway resource.

For information about Vpc Ipv6 Gateway and how to use it, see [What is Ipv6 Gateway](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_ipv6_gateway" "default" {
  description       = "test"
  ipv6_gateway_name = "rdk-test"
  vpc_id            = "vpc-bp1qco6n6vt5wgn4ynz9b"
  region_id         = "cn-hangzhou"
}
```

## Argument Reference

The following arguments are supported:
* `all` - (Optional) Whether to unbind all tags of the resource. Value:-**true**: untags all resources.-**false** (default): does not remove all tags of the resource.
* `description` - (Optional) Description
* `ipv6_gateway_name` - (Optional) Ipv6GatewayName
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the instance belongs.
* `tags` - (Optional) The tags of VSwitch.See the following `Block Tags`.
* `vpc_id` - (Required,ForceNew) VpcId

The following arguments will be discarded. Please use new fields as soon as possible:

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The tag key of VSwitch.
* `tag_value` - (Optional) The tag value of VSwitch.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `business_status` - BusinessStatus
* `create_time` - The creation time of the resource
* `expired_time` - ExpiredTime
* `instance_charge_type` - InstanceChargeType
* `ipv6_gateway_id` - The first ID of the resource
* `resource_group_id` - The ID of the resource group to which the instance belongs.
* `status` - The status of the resource

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ipv6 Gateway.
* `delete` - (Defaults to 5 mins) Used when delete the Ipv6 Gateway.
* `update` - (Defaults to 5 mins) Used when update the Ipv6 Gateway.

## Import

Vpc Ipv6 Gateway can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_ipv6_gateway.example 
```