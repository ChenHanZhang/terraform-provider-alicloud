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
* `external_ip` - (Required) - When querying DNAT entries of an Internet NAT gateway, this parameter indicates the Elastic IP address used in the DNAT entry to provide public network access.
  - When querying DNAT entries of a VPC NAT gateway, this parameter indicates the NAT IP address used for access from external networks.
* `external_port` - (Required) - When configuring a DNAT entry for an Internet NAT gateway, specify the external port or port range that requires port forwarding.

    - The entered port range must be between `1` and `65535`.
    - To forward a range of ports, separate the start and end ports with a forward slash (/), for example, `10/20`.
    - If `ExternalPort` is set to a port range, `InternalPort` must also be set to a port range with the same number of ports. For example, if `ExternalPort` is set to `10/20`, `InternalPort` must be set to `80/90`.
  - When configuring a DNAT entry for a VPC NAT gateway, this parameter specifies the port on the NAT IP address that is accessible from external networks. Valid values: `1` to `65535`.
* `forward_entry_name` - (Optional, Computed) The name of the DNAT rule.
The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character, and cannot start with `http://` or `https://`.
* `forward_table_id` - (Required, ForceNew) The ID of the DNAT table.

-> **NOTE:**  You must specify at least one of the `ForwardTableId` or `NatGatewayId` parameters.

* `internal_ip` - (Required) - When you configure a DNAT entry for an Internet NAT gateway, this parameter specifies the private IP address of the ECS instance that requires public network communication. This private IP address must meet the following requirements:

    - It must belong to the CIDR block of the VPC where the NAT gateway is deployed.

    - It must be assigned to an ECS instance that does not have an EIP bound. The DNAT entry takes effect only under this condition.
  - When you configure a DNAT entry for a VPC NAT gateway, this parameter specifies the private IP address that communicates through the DNAT rule.
* `internal_port` - (Required) - When you configure a DNAT entry for an Internet NAT gateway, this parameter specifies the internal port or port range that requires port forwarding. Valid values: `1` to `65535`.
  - When you configure a DNAT entry for a VPC NAT gateway, this parameter specifies the destination ECS instance port to be mapped. Valid values: `1` to `65535`.
* `ip_protocol` - (Required) The protocol type. Valid values:
  - `TCP`: forwards TCP packets.
  - `UDP`: forwards UDP packets.
  - `Any`: forwards packets of all protocols. If `IpProtocol` is set to `Any`, both `ExternalPort` and `InternalPort` must also be set to `Any` to implement DNAT IP mapping.
* `port_break` - (Optional) Specifies whether to enable port break. Valid values:
  - `true`: Enables port break.
  - `false` (default): Disables port break.

-> **NOTE:**  If a DNAT entry and an SNAT entry share the same public IP address and you need to configure a port number greater than 1024, you must set `PortBreak` to `true`.


-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<forward_table_id>:<forward_entry_id>`.
* `forward_entry_id` - The ID of the DNAT entry to be modified.
* `status` - The status of the DNAT entry.

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