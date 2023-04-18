---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_gateway_route_table_attachment"
sidebar_current: "docs-alicloud-resource-vpc-gateway-route-table-attachment"
description: |-
  Provides a Alicloud Vpc Gateway Route Table Attachment resource.
---

# alicloud_vpc_gateway_route_table_attachment

Provides a Vpc Gateway Route Table Attachment resource.

For information about Vpc Gateway Route Table Attachment and how to use it, see [What is Gateway Route Table Attachment](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_gateway_route_table_attachment" "default" {
  route_table_id  = var.RouteTableId
  region_id       = var.RegionId
  ipv4_gateway_id = var.Ipv4GatewayId
}
```

## Argument Reference

The following arguments are supported:
* `ipv4_gateway_id` - (Required,ForceNew) The ID of the IPv4 Gateway instance.
* `route_table_id` - (Required,ForceNew) The ID of the Gateway route table to be bound.

The following arguments will be discarded. Please use new fields as soon as possible:



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.The value is formulated as `<route_table_id>:<ipv4_gateway_id>`.
* `create_time` - The creation time of the resource
* `status` - The status of the IPv4 Gateway instance. Value:-**Creating**: The function is being created.-**Created**: Created and available.-**Modifying**: is being modified.-**Deleting**: Deleted.-**Deleted**: Deleted.-**Activating**: enabled.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Gateway Route Table Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Gateway Route Table Attachment.

## Import

Vpc Gateway Route Table Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_gateway_route_table_attachment.example <route_table_id>:<ipv4_gateway_id>
```