---
subcategory: "Data Works"
layout: "alicloud"
page_title: "Alicloud: alicloud_data_works_route"
description: |-
  Provides a Alicloud Data Works Route resource.
---

# alicloud_data_works_route

Provides a Data Works Route resource.

Resource group network routing rules.

For information about Data Works Route and how to use it, see [What is Route](https://next.api.alibabacloud.com/document/dataworks-public/2024-05-18/CreateRoute).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-beijing"
}

resource "alicloud_vpc" "default5Bia4h" {
  description = "network_default_resgv2061"
  vpc_name    = "network_default_resgv2061"
  cidr_block  = "10.0.0.0/8"
}

resource "alicloud_vswitch" "defaultss7s7F" {
  description  = "network_default_resg102"
  vpc_id       = alicloud_vpc.default5Bia4h.id
  zone_id      = "cn-beijing-g"
  vswitch_name = "network_default_resg102"
  cidr_block   = "10.0.0.0/24"
}

resource "alicloud_data_works_dw_resource_group" "defaultVJvKvl" {
  default_vswitch_id  = alicloud_vswitch.defaultss7s7F.id
  default_vpc_id      = alicloud_vpc.default5Bia4h.id
  remark              = "route_example"
  payment_type        = "PostPaid"
  resource_group_name = "route_example0008"
}

resource "alicloud_vpc" "defaulte4zhaL" {
  description = "network_default_resgv2010"
  vpc_name    = "network_default_resgv2010"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default675v38" {
  description  = "network_default_resg010"
  vpc_id       = alicloud_vpc.defaulte4zhaL.id
  zone_id      = "cn-beijing-g"
  vswitch_name = "network_default_resg011"
  cidr_block   = "172.16.0.0/24"
}

resource "alicloud_data_works_network" "defaultwEWYyK" {
  vpc_id               = alicloud_vpc.defaulte4zhaL.id
  vswitch_id           = alicloud_vswitch.default675v38.id
  dw_resource_group_id = alicloud_data_works_dw_resource_group.defaultVJvKvl.id
}


resource "alicloud_data_works_route" "default" {
  destination_cidr     = "198.162.0.0/24"
  network_id           = alicloud_data_works_network.defaultwEWYyK.id
  dw_resource_group_id = alicloud_data_works_dw_resource_group.defaultVJvKvl.id
}
```

## Argument Reference

The following arguments are supported:
* `destination_cidr` - (Required) CIDR of the routing destination
* `dw_resource_group_id` - (Optional, ForceNew, Computed) ID of the resource group to which the route belongs.
* `network_id` - (Required, ForceNew, Int) The ID of the network resource to which the route belongs.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Time when route information was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Route.
* `delete` - (Defaults to 5 mins) Used when delete the Route.
* `update` - (Defaults to 5 mins) Used when update the Route.

## Import

Data Works Route can be imported using the id, e.g.

```shell
$ terraform import alicloud_data_works_route.example <route_id>
```