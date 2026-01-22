---
subcategory: "NAT Gateway"
layout: "alicloud"
page_title: "Alicloud: alicloud_snat_entry"
description: |-
  Provides a Alicloud N A T Gateway Snat Entry resource.
---

# alicloud_snat_entry

Provides a N A T Gateway Snat Entry resource.



For information about N A T Gateway Snat Entry and how to use it, see [What is Snat Entry](https://www.alibabacloud.com/help/en/nat-gateway/developer-reference/api-vpc-2016-04-28-createsnatentry-natgws).

-> **NOTE:** Available since v1.119.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
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
  zone_id      = data.alicloud_zones.default.zones.0.id
  vswitch_name = var.name
}

resource "alicloud_nat_gateway" "default" {
  vpc_id           = alicloud_vpc.default.id
  nat_gateway_name = var.name
  payment_type     = "PayAsYouGo"
  vswitch_id       = alicloud_vswitch.default.id
  nat_type         = "Enhanced"
}

resource "alicloud_eip_address" "default" {
  address_name = var.name
}

resource "alicloud_eip_association" "default" {
  allocation_id = alicloud_eip_address.default.id
  instance_id   = alicloud_nat_gateway.default.id
}

resource "alicloud_snat_entry" "default" {
  snat_table_id     = alicloud_nat_gateway.default.snat_table_ids
  source_vswitch_id = alicloud_vswitch.default.id
  snat_ip           = alicloud_eip_address.default.ip_address
}
```

## Argument Reference

The following arguments are supported:
* `eip_affinity` - (Optional, Int, Available since v1.241.0) Specifies whether to enable IP affinity. Valid values:
  - `0` (default): IP affinity is disabled.
  - `1`: IP affinity is enabled.

-> **NOTE:**  After IP affinity is enabled, if the SNAT entry is associated with multiple EIPs or NAT IPs, the same client accessing the same destination IP uses the same EIP or NAT IP. Otherwise, the client randomly selects an EIP or NAT IP from those associated with the SNAT entry for access.

* `network_interface_id` - (Optional, Available since v1.269.0) The ID of the elastic network interface (ENI).

-> **NOTE:**  The IPv4 addresses of the ENI will be used as the SNAT addresses.

* `snat_entry_name` - (Optional, Available since v1.71.2) The name of the SNAT entry.
The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character, and cannot start with `http://` or `https://`.
* `snat_ip` - (Optional) When adding an SNAT entry for a public NAT gateway:

* The SnatIp parameter is required.

* This parameter specifies the EIPs in the SNAT entry. Separate multiple EIPs with commas (,).

* If SnatIp specifies only one public IP address, ECS instances use this specified public IP address to access the Internet.

* If SnatIp specifies multiple public IP addresses, ECS instances randomly use one of the public IP addresses in SnatIp to access the Internet.

-> **NOTE:**  When you specify multiple EIPs to configure an SNAT IP address pool, business connections are distributed across the EIPs using a hash algorithm. Because traffic volume varies per connection, traffic distribution among the EIPs might be uneven. We recommend that you add all EIPs to the same shared bandwidth plan to prevent service degradation caused by any single EIP reaching its bandwidth limit.

When adding an SNAT entry for a VPC NAT gateway:

* This parameter specifies the NAT IP addresses in the SNAT entry. Separate multiple NAT IP addresses with commas (,).

* You must specify either the SnatIp parameter or the NetworkInterfaceId parameter, but not both.
* `snat_table_id` - (Required, ForceNew) The ID of the SNAT table to which the SNAT entry belongs.
* `source_cidr` - (Optional, ForceNew, Computed, Available since v1.71.1) Enter the CIDR block of a VPC, vSwitch, or ECS instance, or enter any custom CIDR block.

SNAT entries support the following granularities:
  - VPC granularity: The CIDR block of the VPC where the NAT gateway is deployed. All ECS instances in this VPC can access the Internet or external networks through the SNAT rule.
  - vSwitch granularity: The CIDR block of a specified vSwitch (for example, 192.168.1.0/24). All ECS instances in this vSwitch can access the Internet or external networks through the SNAT rule.
  - ECS granularity: The IP address of a specified ECS instance (for example, 192.168.1.1/32). This ECS instance can access the Internet or external networks through the SNAT rule.
  - Any custom CIDR block: All ECS instances within this CIDR block can access the Internet or external networks through the SNAT service.

-> **NOTE:**  You must specify either the `SourceCIDR` parameter or the `SourceVSwitchId` parameter, but not both.

* `source_vswitch_id` - (Optional, ForceNew, Computed) The ID of the vSwitch.

* When you add a SNAT entry for an Internet NAT gateway, this parameter indicates that all ECS instances in this vSwitch can access the Internet through the SNAT rule. If you configure multiple EIPs to form a SNAT IP address pool, service connections are distributed across these EIPs using a hash algorithm. Because traffic volume varies per connection, uneven traffic distribution among the EIPs may occur. We recommend that you add all EIPs to the same shared bandwidth plan to prevent service degradation caused by any single EIP reaching its bandwidth limit.

* When you add a SNAT entry for a VPC NAT gateway, this parameter indicates that all ECS instances in this vSwitch can access external networks through the SNAT rule.

-> **NOTE:**  You must specify either the `SourceCIDR` parameter or the `SourceVSwitchId` parameter, but not both.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<snat_table_id>:<snat_entry_id>`.
* `snat_entry_id` - The ID of the SNAT entry to be modified.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Snat Entry.
* `delete` - (Defaults to 5 mins) Used when delete the Snat Entry.
* `update` - (Defaults to 5 mins) Used when update the Snat Entry.

## Import

N A T Gateway Snat Entry can be imported using the id, e.g.

```shell
$ terraform import alicloud_snat_entry.example <snat_table_id>:<snat_entry_id>
```