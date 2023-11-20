---
subcategory: "ENS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ens_load_balancer"
description: |-
  Provides a Alicloud ENS Load Balancer resource.
---

# alicloud_ens_load_balancer

Provides a ENS Load Balancer resource. 

For information about ENS Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.212.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_ens_network" "network" {
  network_name = var.name

  description   = "LoadBalancerNetworkDescription_autotest"
  cidr_block    = "192.168.2.0/24"
  ens_region_id = "cn-hangzhou-44"
}

resource "alicloud_ens_v_switch" "switch" {
  description  = "LoadBalancerVSwitchDescription_autotest"
  cidr_block   = "192.168.2.0/24"
  vswitch_name = var.name

  ens_region_id = "cn-hangzhou-44"
  network_id    = alicloud_ens_network.network.id
}


resource "alicloud_ens_load_balancer" "default" {
  load_balancer_name = var.name

  payment_type       = "PostPaid"
  ens_region_id      = "cn-hangzhou-44"
  load_balancer_spec = "elb.s1.small"
  vswitch_id         = alicloud_ens_v_switch.switch.id
  network_id         = alicloud_ens_network.network.id
}
```

## Argument Reference

The following arguments are supported:
* `ens_region_id` - (Required, ForceNew) The ID of the ENS node.
* `load_balancer_name` - (Optional) Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with http:// and https.
* `load_balancer_spec` - (Required, ForceNew) Specifications of the Server Load Balancer instanceExample value: elb.s2.medium.
* `network_id` - (Required, ForceNew) The network ID of the created edge load balancing (ELB) instance.Example value: n-5sax03dh2eyagujgsn7z9 * * * *.
* `payment_type` - (Required, ForceNew) Server Load Balancer Instance Payment TypeValue: PostPaid.
* `vswitch_id` - (Required, ForceNew) The ID of the vSwitch to which the VPC instance belongsExample value: vsw-5s78haoys9oylle6ln71m * * * *.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The creation Time (UTC) of the load balancing instance.
* `status` - The status of the SLB instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Load Balancer.
* `delete` - (Defaults to 5 mins) Used when delete the Load Balancer.
* `update` - (Defaults to 5 mins) Used when update the Load Balancer.

## Import

ENS Load Balancer can be imported using the id, e.g.

```shell
$ terraform import alicloud_ens_load_balancer.example <id>
```