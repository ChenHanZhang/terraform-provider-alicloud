---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_v_switch_cidr_reservation"
description: |-
  Provides a Alicloud VPC V Switch Cidr Reservation resource.
---

# alicloud_vpc_v_switch_cidr_reservation

Provides a VPC V Switch Cidr Reservation resource. The reserved network segment of the vswitch. This resource type can be used only in ap-southeast region.

For information about VPC V Switch Cidr Reservation and how to use it, see [What is V Switch Cidr Reservation](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/610154).

-> **NOTE:** Available since v1.208.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "defaultVpc" {
  vpc_name   = var.name
  cidr_block = "10.0.0.0/8"
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id       = alicloud_vpc.defaultVpc.id
  cidr_block   = "10.0.0.0/20"
  vswitch_name = "${var.name}1"
  zone_id      = data.alicloud_zones.default.zones.0.id
}


resource "alicloud_vpc_v_switch_cidr_reservation" "default" {
  ip_version                    = "IPv4"
  vswitch_id                    = alicloud_vswitch.defaultVSwitch.id
  cidr_reservation_description  = "test"
  cidr_reservation_cidr         = "10.0.10.0/24"
  vswitch_cidr_reservation_name = var.name
  cidr_reservation_type         = "Prefix"
}
```

## Argument Reference

The following arguments are supported:
* `cidr_reservation_cidr` - (Optional, ForceNew, Computed, Available since v1.205.0) Reserved network segment CIdrBlock.
* `cidr_reservation_description` - (Optional, Available since v1.205.0) The description of the reserved CIDR block.
* `cidr_reservation_mask` - (Optional, Available since v1.205.0) Reserved segment mask.
* `cidr_reservation_type` - (Optional, ForceNew, Computed, Available since v1.205.0) Reserved CIDR Block Type.Valid values: `Prefix`. Default value: Prefix.
* `ip_version` - (Optional, ForceNew, Computed, Available since v1.205.0) Reserved ip version of network segment, valid values: `IPv4`, `IPv6`, default IPv4.
* `vswitch_cidr_reservation_name` - (Optional, Available since v1.205.0) The name of the resource.
* `vswitch_id` - (Required, ForceNew, Available since v1.205.0) The Id of the switch instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<vswitch_id>:<vswitch_cidr_reservation_id>`.
* `create_time` - The creation time of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the V Switch Cidr Reservation.
* `delete` - (Defaults to 5 mins) Used when delete the V Switch Cidr Reservation.
* `update` - (Defaults to 5 mins) Used when update the V Switch Cidr Reservation.

## Import

VPC V Switch Cidr Reservation can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_v_switch_cidr_reservation.example <vswitch_id>:<vswitch_cidr_reservation_id>
```