---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv6_gateways"
sidebar_current: "docs-alicloud-datasource-vpc-ipv6-gateways"
description: |-
  Provides a list of Vpc Ipv6 Gateway owned by an Alibaba Cloud account.
---

# alicloud_vpc_ipv6_gateways

This data source provides Vpc Ipv6 Gateway available to the user.[What is Ipv6 Gateway](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_ipv6_gateways" "default" {
  ids               = ["${alicloud_vpc_ipv6_gateway.default.id}"]
  name_regex        = alicloud_vpc_ipv6_gateway.default.name
  ipv6_gateway_name = "rdk-test"
  vpc_id            = "vpc-bp1qco6n6vt5wgn4ynz9b"
}

output "alicloud_vpc_ipv6_gateway_example_id" {
  value = data.alicloud_vpc_ipv6_gateways.default.gateways.0.id
}
```

## Argument Reference

The following arguments are supported:
* `ipv6_gateway_id` - (ForceNew,Optional) The first ID of the resource
* `ipv6_gateway_name` - (ForceNew,Optional) Ipv6GatewayName
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the instance belongs.
* `tags` - (ForceNew,Optional) The tags of VSwitch.See the following `Block Tags`.
* `vpc_id` - (ForceNew,Optional) VpcId
* `ids` - (Optional, ForceNew, Computed) A list of Ipv6 Gateway IDs.
* `ipv6_gateway_names` - (Optional, ForceNew) The name of the Ipv6 Gateway. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

#### Block Tags

The Tags supports the following:
* `tag_key` - (ForceNew,Optional) The tag key of VSwitch.
* `tag_value` - (ForceNew,Optional) The tag value of VSwitch.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Ipv6 Gateway IDs.
* `names` - A list of name of Ipv6 Gateways.
* `gateways` - A list of Ipv6 Gateway Entries. Each element contains the following attributes:
  * `business_status` - BusinessStatus
  * `create_time` - The creation time of the resource
  * `description` - Description
  * `expired_time` - ExpiredTime
  * `instance_charge_type` - InstanceChargeType
  * `ipv6_gateway_id` - The first ID of the resource
  * `ipv6_gateway_name` - Ipv6GatewayName
  * `resource_group_id` - The ID of the resource group to which the instance belongs.
  * `status` - The status of the resource
  * `tags` - The tags of VSwitch.
    * `tag_key` - The tag key of VSwitch.
    * `tag_value` - The tag value of VSwitch.
  * `vpc_id` - VpcId
