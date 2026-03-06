---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_network_interface"
description: |-
  Provides a Alicloud ECS Network Interface resource.
---

# alicloud_ecs_network_interface

Provides a ECS Network Interface resource.

Elastic network card.

For information about ECS Network Interface and how to use it, see [What is Network Interface](https://www.alibabacloud.com/help/en/doc-detail/58504.htm).

-> **NOTE:** Available since v1.123.1.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "192.168.0.0/24"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "192.168.0.0/24"
  zone_id      = data.alicloud_zones.default.zones.0.id
  vpc_id       = alicloud_vpc.default.id
}

resource "alicloud_security_group" "default" {
  name   = var.name
  vpc_id = alicloud_vpc.default.id
}

data "alicloud_resource_manager_resource_groups" "default" {
  status = "OK"
}

resource "alicloud_ecs_network_interface" "default" {
  network_interface_name = var.name
  vswitch_id             = alicloud_vswitch.default.id
  security_group_ids     = [alicloud_security_group.default.id]
  description            = "Basic test"
  primary_ip_address     = "192.168.0.2"
  tags = {
    Created = "TF",
    For     = "Test",
  }
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Description
* `ipv4_prefix` - (Optional, ForceNew, List, Available since v1.273.0) Specify one or more IPv4 prefixes for an ENI
* `ipv6_prefix` - (Optional, ForceNew, List, Available since v1.273.0) Specify one or more IPv6 prefixes for an ENI
* `ipv6_sets` - (Optional, List, Available since v1.273.0) The IPv6 address set assigned to the ENI. See [`ipv6_sets`](#ipv6_sets) below.
* `network_interface_name` - (Optional, Computed) Name of ENI
* `primary_ip_address` - (Optional, ForceNew, Computed) The private IP address of the instance
* `private_ip_sets` - (Optional, List, Available since v1.273.0) A collection of privateipsets. See [`private_ip_sets`](#private_ip_sets) below.
* `qo_s_config` - (Optional, Set, Available since v1.273.0) QoS limit config See [`qo_s_config`](#qo_s_config) below.
* `queue_number` - (Optional, Computed, Int) Number of queues for network cards
* `resource_group_id` - (Optional, ForceNew, Computed) The resource group id
* `resource_type` - (Optional, Available since v1.273.0) Resource type

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `security_group_ids` - (Optional, Computed, List) The security group Collection to which
* `tags` - (Optional, Map) Resource label
* `vswitch_id` - (Required, ForceNew) VPC virtual switch ID

### `ipv6_sets`

The ipv6_sets supports the following:
* `ipv6_address` - (Optional, Available since v1.273.0) The IPv6 address specified for the ENI.

### `private_ip_sets`

The private_ip_sets supports the following:
* `private_ip_address` - (Optional, Available since v1.273.0) The private IP address of the instance.

### `qo_s_config`

The qo_s_config supports the following:
* `enable_qo_s` - (Optional, Available since v1.273.0) config enable
* `qo_s` - (Optional, Set, Available since v1.273.0) Qos cofiguration See [`qo_s`](#qo_s_config-qo_s) below.

### `qo_s_config-qo_s`

The qo_s_config-qo_s supports the following:
* `bandwidth_rx` - (Optional, Int, Available since v1.273.0) Maximum intranet inbound bandwidth limit
* `bandwidth_tx` - (Optional, Int, Available since v1.273.0) Maximum intranet outgoing bandwidth limit
* `concurrent_connections` - (Optional, Int, Available since v1.273.0) Maximum value of connection quantity
* `pps_rx` - (Optional, Int, Available since v1.273.0) Inbound network packet transmission and reception capability of the intranet
* `pps_tx` - (Optional, Int, Available since v1.273.0) Intranet outbound network packet sending and receiving capability

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `private_ip_sets` - A collection of privateipsets.
  * `associated_public_ip` - The elastic IP address associated with the elastic network card.
    * `allocation_id` - > This parameter is under test and is not yet open for use.
    * `public_ip_address` - The IP address of the elastic network.
  * `primary` - Is the IP address of the primary private network.
* `status` - Status of ENI.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Network Interface.
* `delete` - (Defaults to 5 mins) Used when delete the Network Interface.
* `update` - (Defaults to 5 mins) Used when update the Network Interface.

## Import

ECS Network Interface can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_network_interface.example <network_interface_id>
```