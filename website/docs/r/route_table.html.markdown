---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_route_table"
sidebar_current: "docs-alicloud-resource-route-table"
description: |-
  Provides a Alicloud Vpc Route Table resource.
---

# alicloud_route_table

Provides a Vpc Route Table resource.

For information about Vpc Route Table and how to use it, see [What is Route Table](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_route_table" "default" {
  region_id        = var.RegionId
  vpc_id           = var.VpcId
  route_table_name = var.RouteTableName
  description      = var.Description
}
```

## Argument Reference

The following arguments are supported:
* `associate_type` - (ForceNew,Computed,Optional) The type of cloud resource that is bound to the routing table. Value:-**VSwitch**: switch.-**Gateway**:IPv4 Gateway.
* `description` - (Optional) Description of the routing table.
* `route_table_name` - (Optional) The name of the routing table.
* `tags` - (Optional) The tagSee the following `Block Tags`.
* `vpc_id` - (Required,ForceNew) The ID of VPC.

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key value of the tag.
* `tag_value` - (Optional) The value of the tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `associate_type` - The type of cloud resource that is bound to the routing table. Value:-**VSwitch**: switch.-**Gateway**:IPv4 Gateway.
* `create_time` - The creation time of the routing table
* `resource_group_id` - Resource group ID.
* `route_table_id` - The ID of the routing table.
* `route_table_type` - The type of routing table. Values: -custom: Custom routing table - System: System routing table
* `router_id` - The router ID to which the routing table belongs.
* `router_type` - The router type to which the routing table belongs. Value: -VRouter: VPC Router -VBR: Boundary Router
* `status` - Routing table state
* `vswitch_ids` - The ID of the switch.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Route Table.
* `delete` - (Defaults to 10 mins) Used when delete the Route Table.
* `update` - (Defaults to 10 mins) Used when update the Route Table.

## Import

Vpc Route Table can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_route_table.example 
```