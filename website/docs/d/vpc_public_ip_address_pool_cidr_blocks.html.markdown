---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_public_ip_address_pool_cidr_blocks"
sidebar_current: "docs-alicloud-datasource-vpc-public-ip-address-pool-cidr-blocks"
description: |-
  Provides a list of Vpc Public Ip Address Pool Cidr Block owned by an Alibaba Cloud account.
---

# alicloud_vpc_public_ip_address_pool_cidr_blocks

This data source provides Vpc Public Ip Address Pool Cidr Block available to the user.[What is Public Ip Address Pool Cidr Block](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_public_ip_address_pool_cidr_blocks" "default" {
  cidr_block                = "47.118.126.0/25"
  public_ip_address_pool_id = "pippool-bp1aqw7z6cyd2u7ayvw7v"
}

output "alicloud_vpc_public_ip_address_pool_cidr_block_example_id" {
  value = data.alicloud_vpc_public_ip_address_pool_cidr_blocks.default.blocks.0.id
}
```

## Argument Reference

The following arguments are supported:
* `cidr_block` - (ForceNew,Optional) IP地址池CIDR Block网段
* `public_ip_address_pool_id` - (Required,ForceNew) IP地址池的实例ID
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `blocks` - A list of Public Ip Address Pool Cidr Block Entries. Each element contains the following attributes:
  * `cidr_block` - IP地址池CIDR Block网段
  * `create_time` - The creation time of the resource
  * `public_ip_address_pool_id` - IP地址池的实例ID
  * `status` - The status of the resource
  * `total_ip_num` - IP地址池中网段的可用IP地址总数
  * `used_ip_num` - IP地址池中网段的已用IP地址数
