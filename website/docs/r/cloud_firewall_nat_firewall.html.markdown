---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_nat_firewall"
description: |-
  Provides a Alicloud Cloud Firewall Nat Firewall resource.
---

# alicloud_cloud_firewall_nat_firewall

Provides a Cloud Firewall Nat Firewall resource.



For information about Cloud Firewall Nat Firewall and how to use it, see [What is Nat Firewall](https://www.alibabacloud.com/help/zh/cloud-firewall/developer-reference/api-cloudfw-2017-12-07-createsecurityproxy).

-> **NOTE:** Available since v1.224.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shenzhen"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "defaultikZ0gD" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = var.name
}

resource "alicloud_vswitch" "defaultp4O7qi" {
  vpc_id       = alicloud_vpc.defaultikZ0gD.id
  cidr_block   = "172.16.6.0/24"
  vswitch_name = var.name
  zone_id      = data.alicloud_zones.default.zones.0.id
}

resource "alicloud_nat_gateway" "default2iRZpC" {
  eip_bind_mode    = "MULTI_BINDED"
  vpc_id           = alicloud_vpc.defaultikZ0gD.id
  nat_gateway_name = var.name
  payment_type     = "PayAsYouGo"
  vswitch_id       = alicloud_vswitch.defaultp4O7qi.id
  nat_type         = "Enhanced"
  network_type     = "internet"
}

resource "alicloud_eip_address" "defaultyiRwgs" {
  address_name = var.name
}

resource "alicloud_eip_association" "defaults2MTuO" {
  instance_id   = alicloud_nat_gateway.default2iRZpC.id
  allocation_id = alicloud_eip_address.defaultyiRwgs.id
  mode          = "NAT"
  instance_type = "Nat"
}

resource "alicloud_snat_entry" "defaultAKE43g" {
  snat_ip           = alicloud_eip_address.defaultyiRwgs.ip_address
  snat_table_id     = alicloud_nat_gateway.default2iRZpC.snat_table_ids
  source_vswitch_id = alicloud_vswitch.defaultp4O7qi.id
}

resource "alicloud_cloud_firewall_nat_firewall" "default" {
  nat_gateway_id = alicloud_nat_gateway.default2iRZpC.id
  nat_route_entry_list {
    nexthop_type     = "NatGateway"
    route_table_id   = alicloud_vpc.defaultikZ0gD.route_table_id
    nexthop_id       = alicloud_nat_gateway.default2iRZpC.id
    destination_cidr = "0.0.0.0/0"
  }

  firewall_switch = "close"
  vswitch_auto    = "true"
  status          = "closed"
  region_no       = "cn-shenzhen"
  lang            = "zh"
  proxy_name      = "CFW-example"
  vswitch_id      = alicloud_snat_entry.defaultAKE43g.source_vswitch_id
  strict_mode     = "0"
  vpc_id          = alicloud_vpc.defaultikZ0gD.id
  vswitch_cidr    = "172.16.5.0/24"
}
```

## Argument Reference

The following arguments are supported:
* `firewall_switch` - (Optional) The security protection switch. Valid values:
  - `open`: enabled
  - `close`: disabled.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `lang` - (Optional) The language of the request and response. Valid values:
  - `zh` (default): Chinese
  - `en`: English.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `nat_gateway_id` - (Required, ForceNew) The ID of the NAT gateway to query.
* `nat_route_entry_list` - (Required, ForceNew, List) The list of routes to be switched for the NAT gateway. See [`nat_route_entry_list`](#nat_route_entry_list) below.
* `proxy_name` - (Required, ForceNew) The name of the NAT firewall. The name must be 4 to 50 characters in length and can contain uppercase and lowercase letters, Chinese characters, digits, and underscores (_). It cannot start with an underscore.
* `region_no` - (Required, ForceNew) The region ID of the VPC.

-> **NOTE:**  For more information about regions supported by Cloud Firewall, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).

* `status` - (Optional, Computed) The status of Cloud Firewall. Valid values:
  - configuring: Creating
  - deleting: Deleting
  - normal: Normal
  - abnormal: Abnormal
  - opening: Enabling
  - closing: Disabling
  - closed: Disabled
* `strict_mode` - (Optional, ForceNew, Int) Specifies whether to enable strict mode.
  - 1: Enable strict mode.
  - 0: Disable strict mode.
* `vswitch_id` - (Optional) The vSwitch ID. This parameter is required when manual mode is enabled for vSwitches.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `vpc_id` - (Required, ForceNew) The ID of the VPC instance.
* `vswitch_auto` - (Optional) Specifies whether to enable automatic mode for vSwitches. Valid values:
  - `true`: Automatic mode
  - `false`: Manual mode

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `vswitch_cidr` - (Optional) The CIDR block of the vSwitch. This parameter is required when automatic mode is enabled for vSwitches.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


### `nat_route_entry_list`

The nat_route_entry_list supports the following:
* `destination_cidr` - (Required, ForceNew) The destination CIDR block of the default route.
* `nexthop_id` - (Required, ForceNew) The next hop address of the original NAT gateway.
* `nexthop_type` - (Required, ForceNew) The network type of the next hop. Valid value: NatGateway (NAT gateway).
* `route_table_id` - (Required, ForceNew) The route table to which the default route of the NAT gateway belongs.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `ali_uid` - The UID of the Alibaba Cloud account.
* `member_uid` - The UID of the member account that belongs to the current Alibaba Cloud account.
* `nat_gateway_name` - The name of the NAT gateway.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Nat Firewall.
* `delete` - (Defaults to 13 mins) Used when delete the Nat Firewall.
* `update` - (Defaults to 5 mins) Used when update the Nat Firewall.

## Import

Cloud Firewall Nat Firewall can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_nat_firewall.example <proxy_id>
```