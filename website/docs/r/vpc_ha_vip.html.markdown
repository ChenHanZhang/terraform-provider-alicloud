---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ha_vip"
description: |-
  Provides a Alicloud VPC Ha Vip resource.
---

# alicloud_vpc_ha_vip

Provides a VPC Ha Vip resource. Highly available virtual IP.

For information about VPC Ha Vip and how to use it, see [What is Ha Vip](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createhavip).

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
  description = "tf-test-acc-chenyi"
  vpc_name    = var.name
  cidr_block  = "192.168.0.0/16"
}

resource "alicloud_vswitch" "defaultVswitch" {
  vpc_id       = alicloud_vpc.defaultVpc.id
  cidr_block   = "192.168.0.0/21"
  vswitch_name = "${var.name}1"
  zone_id      = data.alicloud_zones.default.zones.0.id
  description  = "tf-testacc-chenyi"
}


resource "alicloud_vpc_ha_vip" "default" {
  description = "test"
  vswitch_id  = alicloud_vswitch.defaultVswitch.id
  ha_vip_name = var.name
  ip_address  = "192.168.1.101"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) The description of the HaVip instance. The length is 2 to 256 characters.
* `ha_vip_name` - (Optional) The name of the HaVip instance.
* `ip_address` - (Optional, ForceNew, Computed) The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
* `resource_group_id` - (Optional, Computed) The ID of the resource group.
* `tags` - (Optional, Map) The tags of HaVip.
* `vswitch_id` - (Required, ForceNew) The switch ID to which the HaVip instance belongs.

The following arguments will be discarded. Please use new fields as soon as possible:
* `havip_name` - (Deprecated since v1.208.0). Field 'havip_name' has been deprecated from provider version 1.208.0. New field 'ha_vip_name' instead.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The creation time of the resource.
* `status` - The status of this resource instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ha Vip.
* `delete` - (Defaults to 5 mins) Used when delete the Ha Vip.
* `update` - (Defaults to 5 mins) Used when update the Ha Vip.

## Import

VPC Ha Vip can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_ha_vip.example <id>
```