---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_public_ip_address_pool_cidr_block"
sidebar_current: "docs-alicloud-resource-vpc-public-ip-address-pool-cidr-block"
description: |-
  Provides a Alicloud Vpc Public Ip Address Pool Cidr Block resource.
---

# alicloud_vpc_public_ip_address_pool_cidr_block

Provides a Vpc Public Ip Address Pool Cidr Block resource.

For information about Vpc Public Ip Address Pool Cidr Block and how to use it, see [What is Public Ip Address Pool Cidr Block](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_public_ip_address_pool_cidr_block" "default" {
  region_id                 = "cn-hangzhou"
  public_ip_address_pool_id = "pippool-bp1aqw7z6cyd2u7ayvw7v"
  cidr_block                = "47.118.126.0/25"
}
```

## Argument Reference

The following arguments are supported:
* `cidr_block` - (ForceNew,Computed,Optional) IP地址池CIDR Block网段
* `cidr_mask` - (ForceNew,Optional) IP address and network segment mask.After you enter the mask, the system automatically allocates the IP address network segment.Value range: **24** to **28 * *.> **CidrBlock** and **CidrMask** cannot be configured at the same time. Select one of them to configure.
* `public_ip_address_pool_id` - (Required,ForceNew) IP地址池的实例ID

The following arguments will be discarded. Please use new fields as soon as possible:



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.The value is formulated as `<public_ip_address_pool_id>:<cidr_block>`.
* `cidr_block` - IP地址池CIDR Block网段
* `create_time` - The creation time of the resource
* `status` - The status of the resource
* `total_ip_num` - IP地址池中网段的可用IP地址总数
* `used_ip_num` - IP地址池中网段的已用IP地址数

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Public Ip Address Pool Cidr Block.
* `delete` - (Defaults to 5 mins) Used when delete the Public Ip Address Pool Cidr Block.

## Import

Vpc Public Ip Address Pool Cidr Block can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_public_ip_address_pool_cidr_block.example <public_ip_address_pool_id>:<cidr_block>
```