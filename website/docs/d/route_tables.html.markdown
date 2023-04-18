---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_route_tables"
sidebar_current: "docs-alicloud-datasource-route-tables"
description: |-
  Provides a list of Vpc Route Table owned by an Alibaba Cloud account.
---

# alicloud_route_tables

This data source provides Vpc Route Table available to the user.[What is Route Table](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_route_tables" "default" {
  ids              = ["${alicloud_route_table.default.id}"]
  name_regex       = alicloud_route_table.default.name
  route_table_name = var.RouteTableName
  vpc_id           = var.VpcId
}

output "alicloud_route_table_example_id" {
  value = data.alicloud_vpc_route_tables.default.tables.0.id
}
```

## Argument Reference

The following arguments are supported:
* `resource_group_id` - (ForceNew,Optional) Resource group ID.
* `route_table_id` - (ForceNew,Optional) The ID of the routing table.
* `route_table_name` - (ForceNew,Optional) The name of the routing table.
* `router_id` - (ForceNew,Optional) The router ID to which the routing table belongs.
* `router_type` - (ForceNew,Optional) The router type to which the routing table belongs. Value: -VRouter: VPC Router -VBR: Boundary Router
* `vpc_id` - (ForceNew,Optional) The ID of VPC.
* `ids` - (Optional, ForceNew, Computed) A list of Route Table IDs.
* `route_table_names` - (Optional, ForceNew) The name of the Route Table. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Route Table IDs.
* `names` - A list of name of Route Tables.
* `tables` - A list of Route Table Entries. Each element contains the following attributes:
  * `associate_type` - The type of cloud resource that is bound to the routing table. Value:-**VSwitch**: switch.-**Gateway**:IPv4 Gateway.
  * `create_time` - The creation time of the routing table
  * `description` - Description of the routing table.
  * `resource_group_id` - Resource group ID.
  * `route_table_id` - The ID of the routing table.
  * `route_table_name` - The name of the routing table.
  * `route_table_type` - The type of routing table. Values: -custom: Custom routing table - System: System routing table
  * `router_id` - The router ID to which the routing table belongs.
  * `router_type` - The router type to which the routing table belongs. Value: -VRouter: VPC Router -VBR: Boundary Router
  * `status` - Routing table state
  * `tags` - The tag
    * `tag_key` - The key value of the tag.
    * `tag_value` - The value of the tag.
  * `vswitch_ids` - The ID of the switch.
  * `vpc_id` - The ID of VPC.
