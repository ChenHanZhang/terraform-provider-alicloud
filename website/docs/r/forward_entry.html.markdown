---
subcategory: "NAT Gateway"
layout: "alicloud"
page_title: "Alicloud: alicloud_forward_entry"
description: |-
  Provides a Alicloud N A T Gateway Forward Entry resource.
---

# alicloud_forward_entry

Provides a N A T Gateway Forward Entry resource.

DNAT table entry.

For information about N A T Gateway Forward Entry and how to use it, see [What is Forward Entry](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/CreateForwardEntry).

-> **NOTE:** Available since v1.119.1.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "forward-entry-example-name"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/21"
  zone_id      = data.alicloud_zones.default.zones[0].id
  vswitch_name = var.name
}

resource "alicloud_nat_gateway" "default" {
  vpc_id               = alicloud_vpc.default.id
  internet_charge_type = "PayByLcu"
  nat_gateway_name     = var.name
  nat_type             = "Enhanced"
  vswitch_id           = alicloud_vswitch.default.id
}

resource "alicloud_eip_address" "default" {
  address_name = var.name
}

resource "alicloud_eip_association" "default" {
  allocation_id = alicloud_eip_address.default.id
  instance_id   = alicloud_nat_gateway.default.id
}

resource "alicloud_forward_entry" "default" {
  forward_table_id = alicloud_nat_gateway.default.forward_table_ids
  external_ip      = alicloud_eip_address.default.ip_address
  external_port    = "80"
  ip_protocol      = "tcp"
  internal_ip      = "172.16.0.3"
  internal_port    = "8080"
}
```

## Argument Reference

The following arguments are supported:
* `external_ip` - (Required) The public IP address in the DNAT entry. The public IP address is used by the ECS instance to receive requests from the Internet.
* `external_port` - (Required) The external port in the DNAT entry. The external port is used by the ECS instance to receive requests from the Internet.
* `forward_entry_name` - (Optional, Computed) The name of the DNAT entry.
* `forward_table_id` - (Required, ForceNew) The ID of the DNAT table to which the DNAT entry belongs.
* `internal_ip` - (Required) The private IP address that is mapped to the public IP address in the DNAT entry.
* `internal_port` - (Required) The internal port that is mapped to the external port in the DNAT entry.
* `ip_protocol` - (Required) The type of the protocol.
* `port_break` - (Optional) Whether to enable Port breakout, value:
  - `true`: Enable Port breakout.
  - `false` (default): Do not enable Port breakout.

-> **NOTE:**  When the DNAT entry and SNAT entry use the same public IP address, if you need to configure a port number greater than 1024, you must specify `PortBreak` to **true * *.


-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<forward_table_id>:<forward_entry_id>`.
* `forward_entry_id` - The ID of the DNAT entry.
* `status` - The state of the DNAT entry

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Forward Entry.
* `delete` - (Defaults to 5 mins) Used when delete the Forward Entry.
* `update` - (Defaults to 5 mins) Used when update the Forward Entry.

## Import

N A T Gateway Forward Entry can be imported using the id, e.g.

```shell
$ terraform import alicloud_forward_entry.example <forward_table_id>:<forward_entry_id>
```