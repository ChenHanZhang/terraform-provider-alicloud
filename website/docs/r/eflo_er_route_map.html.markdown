---
subcategory: "Eflo"
layout: "alicloud"
page_title: "Alicloud: alicloud_eflo_er_route_map"
description: |-
  Provides a Alicloud Eflo Er Route Map resource.
---

# alicloud_eflo_er_route_map

Provides a Eflo Er Route Map resource.

Lingjun HUB routing strategy.

For information about Eflo Er Route Map and how to use it, see [What is Er Route Map](https://next.api.alibabacloud.com/document/eflo/2022-05-30/CreateErRouteMap).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-wulanchabu"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}

resource "alicloud_vpc" "VPC" {
  cidr_block = "192.168.0.0/16"
}

resource "alicloud_vswitch" "VSW" {
  vpc_id     = alicloud_vpc.VPC.id
  zone_id    = var.zone_id
  cidr_block = "192.168.0.0/24"
}

resource "alicloud_eflo_vcc" "VCC" {
  connection_type = "VPC"
  zone_id         = var.zone_id
  vswitch_id      = alicloud_vswitch.VSW.id
  vpc_id          = alicloud_vpc.VPC.id
  bandwidth       = "1000"
}

resource "alicloud_eflo_vpd" "VPD" {
  cidr     = "10.0.0.0/8"
  vpd_name = "example-route-map"
}

resource "alicloud_eflo_er" "ER" {
  er_name        = "er-example-routemap"
  master_zone_id = var.zone_id
}

resource "alicloud_eflo_er_attachment" "ER_AT_VPD" {
  resource_tenant_id     = "1511928242963727"
  instance_id            = alicloud_eflo_vpd.VPD.id
  auto_receive_all_route = false
  er_id                  = alicloud_eflo_er.ER.id
  instance_type          = "VPD"
  er_attachment_name     = "example-route-map-tf"
}

resource "alicloud_eflo_er_attachment" "ER_AT_VCC" {
  resource_tenant_id     = "1511928242963727"
  instance_id            = alicloud_eflo_vcc.VCC.id
  er_id                  = alicloud_eflo_er.ER.id
  instance_type          = "VCC"
  er_attachment_name     = "example-route-map-tf"
  auto_receive_all_route = false
}


resource "alicloud_eflo_er_route_map" "default" {
  transmission_instance_type  = "VPD"
  action                      = "permit"
  reception_instance_type     = "VCC"
  description                 = "录入策略VPD-VCC"
  reception_instance_id       = alicloud_eflo_vcc.VCC.id
  er_id                       = alicloud_eflo_er.ER.id
  reception_instance_owner    = "1511928242963727"
  transmission_instance_owner = "1511928242963727"
  transmission_instance_id    = alicloud_eflo_vpd.VPD.id
  er_route_map_num            = "1001"
  destination_cidr_block      = "0.0.0.0/0"
}
```

## Argument Reference

The following arguments are supported:
* `action` - (Required, ForceNew) Strategic behavior
* `description` - (Optional) Lingjun HUB routing policy description information
* `destination_cidr_block` - (Required, ForceNew) Destination Network segment
* `er_id` - (Required, ForceNew) Lingjun HUB
* `er_route_map_num` - (Required, ForceNew, Int) Policy number
* `reception_instance_id` - (Required, ForceNew) Receive instance ID
* `reception_instance_owner` - (Optional, ForceNew) The tenant ID of the receiving instance.
* `reception_instance_type` - (Required, ForceNew) Receive instance type
* `transmission_instance_id` - (Required, ForceNew) Publish instance ID
* `transmission_instance_owner` - (Optional, ForceNew) The ID of the tenant to which the publish instance belongs.
* `transmission_instance_type` - (Required, ForceNew) Publish instance type

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<er_id>:<er_route_map_id>`.
* `create_time` - Creation time.
* `er_route_map_id` - Routing Policy ID.
* `region_id` - The region ID of the resource.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Er Route Map.
* `delete` - (Defaults to 5 mins) Used when delete the Er Route Map.
* `update` - (Defaults to 5 mins) Used when update the Er Route Map.

## Import

Eflo Er Route Map can be imported using the id, e.g.

```shell
$ terraform import alicloud_eflo_er_route_map.example <er_id>:<er_route_map_id>
```