---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_vswitch_cidr_reservations"
sidebar_current: "docs-alicloud-datasource-vpc-vswitch-cidr-reservations"
description: |-
  Provides a list of Vpc Vswitch Cidr Reservation owned by an Alibaba Cloud account.
---

# alicloud_vpc_vswitch_cidr_reservations

This data source provides Vpc Vswitch Cidr Reservation available to the user.[What is Vswitch Cidr Reservation](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_vswitch_cidr_reservations" "default" {
  resource_group_id = "rg-aek2xl5qajpkquq"
  vswitch_id        = "vsw-bp19icd33jy4ji9rrnhoc"
}

output "alicloud_vpc_vswitch_cidr_reservation_example_id" {
  value = data.alicloud_vpc_vswitch_cidr_reservations.default.reservations.0.id
}
```

## Argument Reference

The following arguments are supported:
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group
* `vswitch_id` - (Required,ForceNew) The Id of the switch instance.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `reservations` - A list of Vswitch Cidr Reservation Entries. Each element contains the following attributes:
  * `cdir_reservation_type` - Reserved CIDR Block TypeValid values: prefix and Explicit. Default value: prefix
  * `cidr_reservation_cidr` - Reserved network segment CIdrBlock
  * `cidr_reservation_description` - Reserved CIDR Block Description
  * `create_time` - The creation time of the resource
  * `ip_version` - Reserved ip version of network segment, value IPv4,IPv6, default IPv4.
  * `resource_group_id` - The ID of the resource group
  * `status` - The status of the resource
  * `vswitch_cidr_reservation_id` - The first ID of the resource
  * `vswitch_cidr_reservation_name` - The name of the resource
  * `vswitch_id` - The Id of the switch instance.
  * `vpc_instance_id` - The id of the vpc instance to which the reserved CIDR block belongs.
