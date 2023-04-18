---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv4_gateways"
sidebar_current: "docs-alicloud-datasource-vpc-ipv4-gateways"
description: |-
  Provides a list of Vpc Ipv4 Gateway owned by an Alibaba Cloud account.
---

# alicloud_vpc_ipv4_gateways

This data source provides Vpc Ipv4 Gateway available to the user.[What is Ipv4 Gateway](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_ipv4_gateways" "default" {
  ids               = ["${alicloud_vpc_ipv4_gateway.default.id}"]
  name_regex        = alicloud_vpc_ipv4_gateway.default.name
  ipv4_gateway_name = "Ipv4GatewayNameTest"
  vpc_id            = "vpc-bp15m8cngyup50ukyy3fp"
}

output "alicloud_vpc_ipv4_gateway_example_id" {
  value = data.alicloud_vpc_ipv4_gateways.default.gateways.0.id
}
```

## Argument Reference

The following arguments are supported:
* `ipv4_gateway_id` - (ForceNew,Optional) The resource attribute field that represents the resource level 1 ID.
* `ipv4_gateway_name` - (ForceNew,Optional) Resource name
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the instance belongs.
* `tags` - (ForceNew,Optional) The tags of VSwitch.See the following `Block Tags`.
* `vpc_id` - (ForceNew,Optional) The ID of the VPC associated with the IPv4 Gateway.
* `ids` - (Optional, ForceNew, Computed) A list of Ipv4 Gateway IDs.
* `ipv4_gateway_names` - (Optional, ForceNew) The name of the Ipv4 Gateway. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

#### Block Tags

The Tags supports the following:
* `tag_key` - (ForceNew,Optional) The tag key of VSwitch.
* `tag_value` - (ForceNew,Optional) The tag value of VSwitch.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Ipv4 Gateway IDs.
* `names` - A list of name of Ipv4 Gateways.
* `gateways` - A list of Ipv4 Gateway Entries. Each element contains the following attributes:
  * `create_time` - The creation time of the resource
  * `enabled` - Enabled
  * `ipv4_gateway_description` - Description information
  * `ipv4_gateway_id` - The resource attribute field that represents the resource level 1 ID.
  * `ipv4_gateway_name` - Resource name
  * `ipv4_gateway_route_table_id` - ID of the route table associated with IPv4 Gateway
  * `resource_group_id` - The ID of the resource group to which the instance belongs.
  * `status` - The status of the resource
  * `tags` - The tags of VSwitch.
    * `tag_key` - The tag key of VSwitch.
    * `tag_value` - The tag value of VSwitch.
  * `vpc_id` - The ID of the VPC associated with the IPv4 Gateway.
