---
subcategory: "NAT Gateway"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_nat_ip"
description: |-
  Provides a Alicloud N A T Gateway Nat Ip resource.
---

# alicloud_vpc_nat_ip

Provides a N A T Gateway Nat Ip resource.

NAT IP address instance.

For information about N A T Gateway Nat Ip and how to use it, see [What is Nat Ip](https://www.alibabacloud.com/help/doc-detail/281976.htm).

-> **NOTE:** Available since v1.136.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_zones" "example" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "example" {
  vpc_name   = "terraform-example"
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "example" {
  vpc_id       = alicloud_vpc.example.id
  cidr_block   = "172.16.0.0/21"
  zone_id      = data.alicloud_zones.example.zones.0.id
  vswitch_name = "terraform-example"
}

resource "alicloud_nat_gateway" "example" {
  vpc_id               = alicloud_vpc.example.id
  internet_charge_type = "PayByLcu"
  nat_gateway_name     = "terraform-example"
  description          = "terraform-example"
  nat_type             = "Enhanced"
  vswitch_id           = alicloud_vswitch.example.id
  network_type         = "intranet"
}

resource "alicloud_vpc_nat_ip_cidr" "example" {
  nat_ip_cidr             = "192.168.0.0/16"
  nat_gateway_id          = alicloud_nat_gateway.example.id
  nat_ip_cidr_description = "terraform-example"
  nat_ip_cidr_name        = "terraform-example"
}

resource "alicloud_vpc_nat_ip" "example" {
  nat_ip             = "192.168.0.37"
  nat_gateway_id     = alicloud_nat_gateway.example.id
  nat_ip_description = "example_value"
  nat_ip_name        = "example_value"
  nat_ip_cidr        = alicloud_vpc_nat_ip_cidr.example.nat_ip_cidr
}

```

## Argument Reference

The following arguments are supported:
* `dry_run` - (Optional, Computed) Specifies whether to only precheck the request. Valid values:
  - `true`: Sends a dry-run request without creating the NAT IP address. The check includes verifying the validity of the AccessKey, RAM user permissions, and whether all required parameters are provided. If the check fails, an error is returned. If the check passes, the error code `DryRunOperation` is returned.
  - `false` (default): Sends a normal request. If the check passes, a 2xx HTTP status code is returned and the NAT IP address is created.

-> **NOTE:** This parameter only applies during resource creation, update or deletion. If modified in isolation without other property changes, Terraform will not trigger any action.

* `ipv4_prefix` - (Optional, ForceNew, Available since v1.269.0) The IPv4 prefix address block to be created.
This IPv4 prefix address block must reside within the reserved CIDR block of the vSwitch where the NAT gateway is deployed, and the reserved CIDR block must not be occupied. The prefix mask must be /28.
* `nat_gateway_id` - (Required, ForceNew) The ID of the VPC NAT gateway instance to which the queried NAT IP address belongs.
* `nat_ip` - (Optional, ForceNew, Computed) The NAT IP address to be created.
If you do not specify this IP address, the system randomly assigns an IP address from your NAT IP CIDR block.
* `nat_ip_cidr` - (Required, ForceNew) The CIDR block of the NAT IP address to be queried.
* `nat_ip_description` - (Optional) The description of the NAT IP address to be modified.
The description must be 2 to 256 characters in length, start with a letter or Chinese character, and cannot start with `http://` or `https://`.
* `nat_ip_name` - (Optional) The name of the NAT IP address to be modified.
The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter or a Chinese character and cannot start with `http://` or `https://`.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<nat_gateway_id>:<nat_ip_id>`.
* `nat_ip_id` - The ID of the NAT IP address instance to be queried. `N` ranges from `1` to `20`.
* `status` - The status of the queried NAT IP address. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Nat Ip.
* `delete` - (Defaults to 5 mins) Used when delete the Nat Ip.
* `update` - (Defaults to 5 mins) Used when update the Nat Ip.

## Import

N A T Gateway Nat Ip can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_nat_ip.example <nat_gateway_id>:<nat_ip_id>
```