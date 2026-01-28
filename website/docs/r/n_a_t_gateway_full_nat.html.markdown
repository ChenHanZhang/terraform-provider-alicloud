---
subcategory: "NAT Gateway"
layout: "alicloud"
page_title: "Alicloud: alicloud_n_a_t_gateway_full_nat"
description: |-
  Provides a Alicloud N A T Gateway Full Nat resource.
---

# alicloud_n_a_t_gateway_full_nat

Provides a N A T Gateway Full Nat resource.



For information about N A T Gateway Full Nat and how to use it, see [What is Full Nat](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/CreateFullNatEntry).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_vpc" "VPC" {
  cidr_block = "172.16.0.0/16"
  vpc_name   = "tf-example-natgw-vpc"
}

resource "alicloud_vswitch" "VSwitch" {
  vpc_id       = alicloud_vpc.VPC.id
  zone_id      = "eu-central-1b"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "tf-example-natgw-vsw"
}

resource "alicloud_nat_gateway" "defaultZ41qqy" {
  description      = "terraform-example"
  nat_gateway_name = "tf-example-fullnat-vpcnat"
  auto_pay         = false
  nat_type         = "Enhanced"
  nat_gateway_private_info {
    vswitch_id = alicloud_vswitch.VSwitch.id
  }
  vpc_id             = alicloud_vpc.VPC.id
  network_type       = "intranet"
  icmp_reply_enabled = false
}

resource "alicloud_privatelink_vpc_endpoint_service" "defaultQvfbOM" {
  payer                  = "Endpoint"
  auto_accept_connection = true
  service_resource_type  = "vpcNat"
}

resource "alicloud_privatelink_vpc_endpoint_service_resource" "defaultlkgL5E" {
  zone_id       = "eu-central-1b"
  resource_id   = alicloud_nat_gateway.defaultZ41qqy.id
  resource_type = "vpcNat"
  service_id    = alicloud_privatelink_vpc_endpoint_service.defaultQvfbOM.id
}

resource "alicloud_vpc" "defaultRiLKXY" {
  cidr_block = "172.16.0.0/16"
  vpc_name   = "tf-example-natgw-ep-vpc"
}

resource "alicloud_vswitch" "defaultujMnA1" {
  vpc_id       = alicloud_vpc.defaultRiLKXY.id
  zone_id      = "eu-central-1a"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "tf-example-natgw-ep-vsw"
}

resource "alicloud_security_group" "defaultb6PMC7" {
  description         = "sg"
  security_group_name = "tf-example-ep-sg"
  vpc_id              = alicloud_vpc.defaultRiLKXY.id
  security_group_type = "normal"
}

resource "alicloud_privatelink_vpc_endpoint" "defaultX3ZwsR" {
  service_id        = alicloud_privatelink_vpc_endpoint_service.defaultQvfbOM.id
  vpc_id            = alicloud_vpc.defaultRiLKXY.id
  endpoint_type     = "Reverse"
  vpc_endpoint_name = "tf-vpcnat-example-ep"
}

resource "alicloud_privatelink_vpc_endpoint_zone" "defaultXXRrPq" {
  zone_id     = alicloud_vswitch.defaultujMnA1.zone_id
  endpoint_id = alicloud_privatelink_vpc_endpoint.defaultX3ZwsR.id
  vswitch_id  = alicloud_vswitch.defaultujMnA1.id
}


resource "alicloud_n_a_t_gateway_full_nat" "default" {
}
```

## Argument Reference

The following arguments are supported:
* `access_ip` - (Required) The backend IP address used for FULLNAT address translation in the FULLNAT entry.
* `access_port` - (Required) The backend port used for port mapping in the FULLNAT entry. Valid values: `1` to `65535`.
* `full_nat_entry_description` - (Optional) The description of the FULLNAT entry.
The description must be 2 to 128 characters in length, start with a letter (uppercase or lowercase) or a Chinese character, and cannot start with `http://` or `https://`.
* `full_nat_entry_name` - (Optional) The new name for the FULLNAT entry to be modified.
The name must be 2 to 128 characters in length, start with a letter or a Chinese character, and cannot start with `http://` or `https://`.
* `full_nat_table_id` - (Required, ForceNew) The ID of the FULLNAT table to which the queried FULLNAT entry belongs.

-> **NOTE:**  You must specify at least one of the `FullNatTableId` or `NatGatewayId` parameters.

* `ip_protocol` - (Required) The protocol type for port forwarding. Valid values:
  - `TCP`: forwards TCP packets.
  - `UDP`: forwards UDP packets.
* `nat_ip` - (Required) The NAT IP address used for address translation in the FULLNAT entry.
* `nat_ip_port` - (Optional) The frontend port used for port mapping in the FULLNAT entry. Valid values: `1` to `65535`.
* `network_interface_id` - (Required) The ID of the elastic network interface (ENI) to be modified.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<full_nat_table_id>:<full_nat_entry_id>`.
* `create_time` - The creation time of the FULLNAT entry.
* `full_nat_entry_id` - The ID of the FULLNAT entry.
* `status` - The status of the FULLNAT entry.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Full Nat.
* `delete` - (Defaults to 5 mins) Used when delete the Full Nat.
* `update` - (Defaults to 5 mins) Used when update the Full Nat.

## Import

N A T Gateway Full Nat can be imported using the id, e.g.

```shell
$ terraform import alicloud_n_a_t_gateway_full_nat.example <full_nat_table_id>:<full_nat_entry_id>
```