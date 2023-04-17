---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vswitchs"
sidebar_current: "docs-alicloud-datasource-vswitchs"
description: |-
  Provides a list of Vpc Vswitch owned by an Alibaba Cloud account.
---

# alicloud_vswitchs

This data source provides Vpc Vswitch available to the user.[What is Vswitch](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_vswitchs" "default" {
  ids          = ["${alicloud_vswitch.default.id}"]
  vswitch_name = "Rdk-test"
  vpc_id       = "vpc-bp1vzgj2t30917m8tlfwd"
  zone_id      = "cn-hangzhou-j"
}

output "alicloud_vswitch_example_id" {
  value = data.alicloud_vpc_vswitchs.default.vswitchs.0.id
}
```

## Argument Reference

The following arguments are supported:
* `is_default` - (ForceNew,Optional) Indicates whether the VSwitch is a default VSwitch.
* `resource_group_id` - (ForceNew,Optional) The resource group id of VSwitch.
* `route_table_id` - (ForceNew,Optional) The route table id
* `vswitch_id` - (ForceNew,Optional) The ID of the VSwitch.
* `vswitch_name` - (ForceNew,Optional) The name of the VSwitch.
* `vpc_id` - (ForceNew,Optional) The VPC ID.
* `zone_id` - (ForceNew,Optional) The zone ID  of the resource
* `ids` - (Optional, ForceNew, Computed) A list of Vswitch IDs.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Vswitch IDs.
* `vswitchs` - A list of Vswitch Entries. Each element contains the following attributes:
  * `available_ip_address_count` - The number of available IP addresses.
  * `cidr_block` - The IPv4 CIDR block of the VSwitch.
  * `create_time` - The creation time of the VSwitch.
  * `description` - The description of VSwitch.
  * `ipv6_cidr_block` - The IPv6 CIDR block of the VSwitch.
  * `ipv6_cidr_block_mask` - The IPv6 CIDR block of the VSwitch.
  * `is_default` - Indicates whether the VSwitch is a default VSwitch.
  * `network_acl_id` - The ID of the network ACL.
  * `resource_group_id` - The resource group id of VSwitch.
  * `route_table_id` - The route table id
  * `status` - The status of the resource
  * `tags` - The tags of VSwitch.
    * `tag_key` - The tag key of VSwitch.
    * `tag_value` - The tag value of VSwitch.
  * `vswitch_id` - The ID of the VSwitch.
  * `vswitch_name` - The name of the VSwitch.
  * `vpc_id` - The VPC ID.
  * `zone_id` - The zone ID  of the resource
