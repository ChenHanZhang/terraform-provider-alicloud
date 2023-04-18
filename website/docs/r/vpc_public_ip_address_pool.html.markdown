---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_public_ip_address_pool"
sidebar_current: "docs-alicloud-resource-vpc-public-ip-address-pool"
description: |-
  Provides a Alicloud Vpc Public Ip Address Pool resource.
---

# alicloud_vpc_public_ip_address_pool

Provides a Vpc Public Ip Address Pool resource.

For information about Vpc Public Ip Address Pool and how to use it, see [What is Public Ip Address Pool](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_public_ip_address_pool" "default" {
  description                 = "rdk-test"
  public_ip_address_pool_name = "rdk-test"
  isp                         = "BGP"
  region_id                   = "cn-beijing"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Description.
* `isp` - (ForceNew,Computed,Optional) Service providers.
* `public_ip_address_pool_name` - (Optional) The name of the resource
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the VPC belongs.
* `tags` - (Optional) The tags of PrefixList.See the following `Block Tags`.

The following arguments will be discarded. Please use new fields as soon as possible:

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of tag.
* `tag_value` - (Optional) The value of tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `create_time` - The creation time of the resource
* `ip_address_remaining` - 是否还有空闲的IP地址。
* `isp` - Service providers.
* `public_ip_address_pool_id` - The first ID of the resource
* `resource_group_id` - The ID of the resource group to which the VPC belongs.
* `status` - The status of the resource
* `total_ip_num` - Total ip number of PublicIpAddressPool.
* `used_ip_num` - Used ip number of PublicIpAddressPool.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Public Ip Address Pool.
* `delete` - (Defaults to 5 mins) Used when delete the Public Ip Address Pool.
* `update` - (Defaults to 5 mins) Used when update the Public Ip Address Pool.

## Import

Vpc Public Ip Address Pool can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_public_ip_address_pool.example 
```