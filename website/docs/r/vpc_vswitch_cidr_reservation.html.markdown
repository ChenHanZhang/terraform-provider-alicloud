---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_vswitch_cidr_reservation"
sidebar_current: "docs-alicloud-resource-vpc-vswitch-cidr-reservation"
description: |-
  Provides a Alicloud Vpc Vswitch Cidr Reservation resource.
---

# alicloud_vpc_vswitch_cidr_reservation

Provides a Vpc Vswitch Cidr Reservation resource.

For information about Vpc Vswitch Cidr Reservation and how to use it, see [What is Vswitch Cidr Reservation](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_vswitch_cidr_reservation" "default" {
  resource_group_id             = "rg-aek2xl5qajpkquq"
  ip_version                    = "IPv4"
  vswitch_id                    = "vsw-bp19icd33jy4ji9rrnhoc"
  region_id                     = "cn-hangzhou"
  vswitch_cidr_reservation_name = "rdk-test"
}
```

## Argument Reference

The following arguments are supported:
* `cdir_reservation_type` - (ForceNew,Optional) Reserved CIDR Block TypeValid values: prefix and Explicit. Default value: prefix
* `cidr_reservation_cidr` - (ForceNew,Optional) Reserved network segment CIdrBlock
* `cidr_reservation_description` - (Optional) Reserved CIDR Block Description
* `cidr_reservation_mask` - (ForceNew,Optional) Reserved segment mask
* `ip_version` - (ForceNew,Optional) Reserved ip version of network segment, value IPv4,IPv6, default IPv4.
* `resource_group_id` - (Computed,Optional) The ID of the resource group
* `resource_type` - (Optional) The cloud resource type of the resource group needs to be modified. Valid values:-**vpc**: vpc.-**eip**: eip.-**bandwidthpackage**: Shared bandwidth.
* `vswitch_cidr_reservation_name` - (Optional) The name of the resource
* `vswitch_id` - (Required,ForceNew) The Id of the switch instance.

The following arguments will be discarded. Please use new fields as soon as possible:



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.The value is formulated as `<vswitch_id>:<vswitch_cidr_reservation_id>`.
* `create_time` - The creation time of the resource
* `resource_group_id` - The ID of the resource group
* `status` - The status of the resource
* `vswitch_cidr_reservation_id` - The first ID of the resource
* `vpc_instance_id` - The id of the vpc instance to which the reserved CIDR block belongs.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Vswitch Cidr Reservation.
* `delete` - (Defaults to 5 mins) Used when delete the Vswitch Cidr Reservation.
* `update` - (Defaults to 5 mins) Used when update the Vswitch Cidr Reservation.

## Import

Vpc Vswitch Cidr Reservation can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_vswitch_cidr_reservation.example <vswitch_id>:<vswitch_cidr_reservation_id>
```