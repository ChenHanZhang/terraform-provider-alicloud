---
subcategory: "Express Connect"
layout: "alicloud"
page_title: "Alicloud: alicloud_express_connect_virtual_border_router"
description: |-
  Provides a Alicloud Express Connect Virtual Border Router resource.
---

# alicloud_express_connect_virtual_border_router

Provides a Express Connect Virtual Border Router resource.

VBR VBR instance.

For information about Express Connect Virtual Border Router and how to use it, see [What is Virtual Border Router](https://www.alibabacloud.com/help/en/doc-detail/44854.htm).

-> **NOTE:** Available since v1.134.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = "cn-hangzhou"
}

variable "name" {
  default = "terraform-example"
}

data "alicloud_express_connect_physical_connections" "default" {
  name_regex = "^preserved-NODELETING"
}

resource "random_integer" "default" {
  min = 1
  max = 2999
}

resource "alicloud_express_connect_virtual_border_router" "default" {
  local_gateway_ip           = "10.0.0.1"
  peer_gateway_ip            = "10.0.0.2"
  peering_subnet_mask        = "255.255.255.252"
  physical_connection_id     = data.alicloud_express_connect_physical_connections.default.connections.0.id
  virtual_border_router_name = var.name
  vlan_id                    = random_integer.default.id
  min_rx_interval            = 1000
  min_tx_interval            = 1000
  detect_multiplier          = 10
}
```

## Argument Reference

The following arguments are supported:
* `bandwidth` - (Optional, Computed, Int) The bandwidth of the VBR instance. Unit: Mbps.
  - When creating a VBR instance for an exclusive leased line, the values are `50`, `100`, `200`, `300`, `400`, `500`, `1000`, `2048`, `5120`, `8192`, `10240`, `20480`, `40960`, `50120`, `61440`, and **102400 * *.
  - When creating a VBR instance for a shared line, you do not need to configure it. The bandwidth of the VBR is the bandwidth set when creating a shared physical line.

* `circuit_code` - (Optional) The circuit code provided by the operator for the physical connection.
* `description` - (Optional) The description information of the VBR.
* `detect_multiplier` - (Optional, Computed, Int) Multiple of detection time.
That is, the maximum number of connection packet losses allowed by the receiver to send messages, which is used to detect whether the link is normal.
Valid values: **3 to 10 * *.
* `enable_ipv6` - (Optional) Whether IPv6 is enabled.
  - `true`: on.
  - `false`: closed.
* `local_gateway_ip` - (Required) The IPv4 address on the Alibaba Cloud side of the VBR instance.
* `local_ipv6_gateway_ip` - (Optional) The IPv6 address on the Alibaba Cloud side of the VBR instance.
* `min_rx_interval` - (Optional, Computed, Int) Configure the receiving interval of BFD packets. Values: **200 to 1000**, in ms.
* `min_tx_interval` - (Optional, Computed, Int) Configure the sending interval of BFD packets. Value: **200~1000**, unit: ms.
* `mtu` - (Optional, Computed, Int, Available since v1.263.0) Maximum transmission unit
* `peer_gateway_ip` - (Required) The IPv4 address of the client side of the VBR instance.
* `peer_ipv6_gateway_ip` - (Optional) The IPv6 address of the client side of the VBR instance.
* `peering_ipv6_subnet_mask` - (Optional) The subnet masks of the Alibaba Cloud-side IPv6 and the customer-side IPv6 of The VBR instance.
* `peering_subnet_mask` - (Required) The subnet masks of the Alibaba Cloud-side IPv4 and the customer-side IPv4 of The VBR instance.
* `physical_connection_id` - (Required, ForceNew) The ID of the physical connection to which the VBR belongs.
* `resource_group_id` - (Optional, Computed, Available since v1.263.0) The ID of the resource group
* `sitelink_enable` - (Optional, Available since v1.263.0) Whether to allow inter-IDC communication
* `status` - (Optional, Computed) Status of the VBR
* `tags` - (Optional, Map, Available since v1.263.0) The tag of the resource
* `vbr_owner_id` - (Optional) The account ID of the VBR instance owner.
The default value is the logon Alibaba Cloud account ID.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `virtual_border_router_name` - (Optional) The name of the VBR instance.
* `vlan_id` - (Required, Int) The VLAN ID of the VBR instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `access_point_id` - The ID of the physical connection point.
* `activation_time` - The time when the VBR is first activated.
* `associated_cens` - The associated CEN instance.
  * `cen_id` - The ID of the CEN instance.
  * `cen_owner_id` - The ID of the account to which the CEN instance belongs.
  * `cen_status` - Cloud enterprise network status.
* `associated_physical_connections` - The associated physical connection information.
  * `circuit_code` - The circuit code provided by the operator for the physical connection.
  * `enable_ipv6` - Whether IPv6 is enabled.
  * `local_gateway_ip` - The IPv4 address on the Alibaba Cloud side of the VBR instance.
  * `local_ipv6_gateway_ip` - The IPv6 address on the Alibaba Cloud side of the VBR instance.
  * `peer_gateway_ip` - The IPv4 address of the client side of the VBR instance.
  * `peer_ipv6_gateway_ip` - The IPv6 address of the client side of the VBR instance.
  * `peering_ipv6_subnet_mask` - The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.
  * `peering_subnet_mask` - The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.
  * `physical_connection_business_status` - The business status of the physical connection.
  * `physical_connection_id` - The ID of the physical connection.
  * `physical_connection_owner_uid` - The ID of the account of the physical connection owner.
  * `physical_connection_status` - The status of the physical connection.
  * `status` - VBR instance status.
  * `vlan_id` - The VLAN ID of the VBR instance.
  * `vlan_interface_id` - The router interface ID of the VBR, which can be used as the next hop of the VBR route.
* `cloud_box_instance_id` - The ID of the cloud box instance.
* `create_time` - The creation time of the VBR.
* `pconn_vbr_expire_time` - The overdue time of the billing VBR.
* `physical_connection_business_status` - The business status of the physical connection.
* `physical_connection_owner_uid` - The ID of the account to which the physical connection belongs.
* `physical_connection_status` - The status of the physical connection.
* `recovery_time` - The last time VBR returned from the Terminated state to the Active state.
* `route_table_id` - The ID of the route table of the VBR.
* `termination_time` - The time when VBR was last terminated.
* `type` - VBR type.
* `vlan_interface_id` - The ID of the VBR router interface.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Virtual Border Router.
* `delete` - (Defaults to 5 mins) Used when delete the Virtual Border Router.
* `update` - (Defaults to 5 mins) Used when update the Virtual Border Router.

## Import

Express Connect Virtual Border Router can be imported using the id, e.g.

```shell
$ terraform import alicloud_express_connect_virtual_border_router.example <vbr_id>
```