---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_gateway_route_table_attachments"
sidebar_current: "docs-alicloud-datasource-vpc-gateway-route-table-attachments"
description: |-
  Provides a list of Vpc Gateway Route Table Attachment owned by an Alibaba Cloud account.
---

# alicloud_vpc_gateway_route_table_attachments

This data source provides Vpc Gateway Route Table Attachment available to the user.[What is Gateway Route Table Attachment](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_gateway_route_table_attachments" "default" {
  ipv4_gateway_id = var.Ipv4GatewayId
}

output "alicloud_vpc_gateway_route_table_attachment_example_id" {
  value = data.alicloud_vpc_gateway_route_table_attachments.default.attachments.0.id
}
```

## Argument Reference

The following arguments are supported:
* `ipv4_gateway_id` - (ForceNew,Optional) The ID of the IPv4 Gateway instance.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `attachments` - A list of Gateway Route Table Attachment Entries. Each element contains the following attributes:
  * `ipv4_gateway_id` - The ID of the IPv4 Gateway instance.
  * `route_table_id` - The ID of the Gateway route table to be bound.
  * `status` - The status of the IPv4 Gateway instance. Value:-**Creating**: The function is being created.-**Created**: Created and available.-**Modifying**: is being modified.-**Deleting**: Deleted.-**Deleted**: Deleted.-**Activating**: enabled.
