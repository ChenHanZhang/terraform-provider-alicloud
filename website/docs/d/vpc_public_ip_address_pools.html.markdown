---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_public_ip_address_pools"
sidebar_current: "docs-alicloud-datasource-vpc-public-ip-address-pools"
description: |-
  Provides a list of Vpc Public Ip Address Pool owned by an Alibaba Cloud account.
---

# alicloud_vpc_public_ip_address_pools

This data source provides Vpc Public Ip Address Pool available to the user.[What is Public Ip Address Pool](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_public_ip_address_pools" "default" {
  ids                         = ["${alicloud_vpc_public_ip_address_pool.default.id}"]
  name_regex                  = alicloud_vpc_public_ip_address_pool.default.name
  isp                         = "BGP"
  public_ip_address_pool_name = "rdk-test"
}

output "alicloud_vpc_public_ip_address_pool_example_id" {
  value = data.alicloud_vpc_public_ip_address_pools.default.pools.0.id
}
```

## Argument Reference

The following arguments are supported:
* `isp` - (ForceNew,Optional) Service providers.
* `public_ip_address_pool_name` - (ForceNew,Optional) The name of the resource
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the VPC belongs.
* `status` - (ForceNew,Optional) The status of the resource
* `tags` - (ForceNew,Optional) The tags of PrefixList.See the following `Block Tags`.
* `ids` - (Optional, ForceNew, Computed) A list of Public Ip Address Pool IDs.
* `public_ip_address_pool_names` - (Optional, ForceNew) The name of the Public Ip Address Pool. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

#### Block Tags

The Tags supports the following:
* `tag_key` - (ForceNew,Optional) The key of tag.
* `tag_value` - (ForceNew,Optional) The value of tag.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Public Ip Address Pool IDs.
* `names` - A list of name of Public Ip Address Pools.
* `pools` - A list of Public Ip Address Pool Entries. Each element contains the following attributes:
  * `create_time` - The creation time of the resource
  * `description` - Description.
  * `ip_address_remaining` - 是否还有空闲的IP地址。
  * `isp` - Service providers.
  * `public_ip_address_pool_id` - The first ID of the resource
  * `public_ip_address_pool_name` - The name of the resource
  * `resource_group_id` - The ID of the resource group to which the VPC belongs.
  * `status` - The status of the resource
  * `tags` - The tags of PrefixList.
    * `tag_key` - The key of tag.
    * `tag_value` - The value of tag.
  * `total_ip_num` - Total ip number of PublicIpAddressPool.
  * `used_ip_num` - Used ip number of PublicIpAddressPool.
