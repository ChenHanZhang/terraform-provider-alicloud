---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv6_cidr_block"
description: |-
  Provides a Alicloud VPC Ipv6 Cidr Block resource.
---

# alicloud_vpc_ipv6_cidr_block

Provides a VPC Ipv6 Cidr Block resource.

VPC IPv6 additional CIDR block.

For information about VPC Ipv6 Cidr Block and how to use it, see [What is Ipv6 Cidr Block](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/AssociateVpcCidrBlock).

-> **NOTE:** Available since v1.281.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_vpc_ipam_ipam" "defaultIpam" {
  operating_region_list = ["cn-hangzhou"]
}

resource "alicloud_vpc_ipam_ipam_pool" "defaultIpv6Pool" {
  ipam_scope_id  = alicloud_vpc_ipam_ipam.defaultIpam.private_default_scope_id
  pool_region_id = alicloud_vpc_ipam_ipam.defaultIpam.region_id
  ip_version     = "IPv6"
}

resource "alicloud_vpc_ipam_ipam_pool_cidr" "defaultIpv6PoolCidr" {
  ipam_pool_id = alicloud_vpc_ipam_ipam_pool.defaultIpv6Pool.id
  cidr         = "fd03:d00:a000::/48"
}

resource "alicloud_vpc" "defaultVpc" {
  cidr_block = "10.0.0.0/8"
  vpc_name   = "example-ipv6-cidr-block"
}


resource "alicloud_vpc_ipv6_cidr_block" "default" {
  ipv6_ipam_pool_id = alicloud_vpc_ipam_ipam_pool.defaultIpv6Pool.id
  vpc_id            = alicloud_vpc.defaultVpc.id
  ipv6_cidr_block   = "fd03:d00:a000::/60"
}
```

## Argument Reference

The following arguments are supported:
* `ipv6_cidr_block` - (Optional, ForceNew, Computed) An additional IPv6 CIDR block.

Both `Ipv6CidrBlock` and `Ipv6CidrMask` are optional parameters, and you can omit both. If neither is specified, the system automatically assigns an IPv6 CIDR block with a /56 prefix length to the VPC from the Alibaba Cloud GUA address pool.

-> **NOTE:**  If you specify only `Ipv6CidrBlock`, you must first call the `AllocateVpcIpv6Cidr` API to reserve the CIDR block.

-> **NOTE:**  If you specify `Ipv6IpamPoolId`, you can directly assign an IPv6 CIDR block by providing either `Ipv6CidrBlock` or `Ipv6CidrMask`, without requiring prior reservation.

* `ipv6_cidr_mask` - (Optional, Int) Add an IPv6 CIDR block to the VPC from an IPAM address pool by specifying a mask length.

-> **NOTE:**  When assigning an additional IPv6 CIDR block to a VPC from an IPAM address pool, you must specify at least one of the `Ipv6CidrBlock` or `Ipv6CidrMask` properties.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv6_ipam_pool_id` - (Optional) The ID of the IPAM IPv6 address pool instance.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv6_isp` - (Optional, ForceNew) The IPv6 address segment type of the VPC. Value:
  - `BGP` (default): Alibaba Cloud BGP IPv6.
  - `ChinaMobile`: China Mobile (single line).
  - `ChinaUnicom`: China Unicom (single line).
  - `ChinaTelecom`: China Telecom (single line).

-> **NOTE:**  If a single-line bandwidth whitelist is enabled, the field can be set to `ChinaTelecom` (China Telecom), `ChinaUnicom` (China Unicom), and `ChinaMobile` (China Mobile).

* `vpc_id` - (Required, ForceNew) The ID of the VPC.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<vpc_id>#<ipv6_cidr_block>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ipv6 Cidr Block.
* `delete` - (Defaults to 5 mins) Used when delete the Ipv6 Cidr Block.

## Import

VPC Ipv6 Cidr Block can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_ipv6_cidr_block.example <vpc_id>#<ipv6_cidr_block>
```