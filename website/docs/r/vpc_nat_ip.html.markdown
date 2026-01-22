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
* `dry_run` - (Optional, Computed) Only pre-This request value: true: send request does not means that traffic can be routed to the NAT IP ADDRESS. Check items including AccessKey is valid, RAM, EPROM, A/D and the user's authorization and whether the fill in the necessary parameters. If not, then returns the corresponding error. If the checks pass, the returns error code DryRunOperation. false (default): sending a normal request, through the inspection of the 2xx HTTP status code and this means that traffic can be routed to the NAT IP ADDRESS.

-> **NOTE:** This parameter only applies during resource creation, update or deletion. If modified in isolation without other property changes, Terraform will not trigger any action.

* `ipv4_prefix` - (Optional, ForceNew, Available since v1.269.0) This property does not have a description in the spec, please add it before generating code.
* `nat_gateway_id` - (Required, ForceNew) NAT IP ADDRESS belongs to the VPC NAT gateway instance ID.
* `nat_ip` - (Optional, ForceNew, Computed) The created NAT IP address.
If you do not specify this IP address, the system randomly assigns an IP address to your NAT IP address segment.
* `nat_ip_cidr` - (Required, ForceNew) NAT IP ADDRESS of the address segment.
* `nat_ip_description` - (Optional) NAT IP ADDRESS description of information. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
* `nat_ip_name` - (Optional) NAT IP ADDRESS the name of the root directory. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, half a period (.), underscore (_) and dash (-). But do not start with http:// or https:// at the beginning.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<nat_gateway_id>:<nat_ip_id>`.
* `nat_ip_id` - NAT IP ADDRESS instance ID.
* `status` - The status of the resource

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