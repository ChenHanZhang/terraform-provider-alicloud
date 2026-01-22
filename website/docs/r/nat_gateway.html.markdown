---
subcategory: "NAT Gateway"
layout: "alicloud"
page_title: "Alicloud: alicloud_nat_gateway"
description: |-
  Provides a Alicloud N A T Gateway Nat Gateway resource.
---

# alicloud_nat_gateway

Provides a N A T Gateway Nat Gateway resource.



For information about N A T Gateway Nat Gateway and how to use it, see [What is Nat Gateway](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/CreateNatGateway).

-> **NOTE:** Available since v1.121.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf_example"
}

data "alicloud_enhanced_nat_available_zones" "default" {
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.0.0.0/8"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  zone_id      = data.alicloud_enhanced_nat_available_zones.default.zones.0.zone_id
  cidr_block   = "10.10.0.0/20"
  vpc_id       = alicloud_vpc.default.id
}

resource "alicloud_nat_gateway" "default" {
  vpc_id           = alicloud_vpc.default.id
  nat_gateway_name = var.name
  payment_type     = "PayAsYouGo"
  vswitch_id       = alicloud_vswitch.default.id
  nat_type         = "Enhanced"
}
```

## Argument Reference

The following arguments are supported:
* `access_mode` - (Optional, Computed, Set, Available since v1.235.0) Configuration of the reverse access VPC-NAT access mode. See [`access_mode`](#access_mode) below.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `auto_pay` - (Optional, ForceNew, Available since v1.269.0) Indicates whether auto-payment is enabled.
  - `true`: Auto-payment is enabled, and orders are automatically paid.
  - `false` (default): Auto-payment is disabled, and you must complete payment in the Order Center after an order is generated.
* `description` - (Optional) The description of the NAT gateway.
The description can be empty, or contain 2 to 256 characters and must not start with `http://` or `https://`.
* `eip_bind_mode` - (Required, Available since v1.184.0) The EIP binding mode of the NAT gateway. Valid values: this parameter can be left empty; if specified, it must be `NAT`, indicating the standard EIP binding mode.


-> **NOTE:**  - You can only change the mode from `MULTI_BINDED` to `NAT`. Changing from `NAT` to `MULTI_BINDED` is not supported. For more information about the `MULTI_BINDED` mode, see [CreateNatGateway](~~120219~~).
  - During the EIP binding mode switch, network connections experience a brief interruption lasting several seconds (the duration increases with the number of EIPs; currently, configuration changes are supported only for NAT gateways bound to up to five EIPs). We recommend performing this operation during off-peak business hours.
  - After switching the EIP binding mode to `NAT`, the public NAT gateway becomes compatible with IPv4 gateways. However, when one public EIP is bound, it consumes one private IP address from the vSwitch where the NAT gateway resides. Ensure that sufficient private IP addresses are available in the vSwitch. If no private IP addresses are available in the vSwitch, you cannot bind new EIPs.
* `enable_session_log` - (Optional, Available since v1.269.0) Specifies whether to enable session logging. Valid values:
  - `true`: Session logging is enabled.
  - `false`: Session logging is disabled.
* `force_delete` - (Optional, Available since v1.269.0) Specifies whether to forcibly delete the NAT gateway. Valid values:
  - `true`: Forces deletion. When this value is set to `true`, the system performs the following actions automatically:

    - Deletes SNAT entries associated with the NAT gateway.

    - Deletes DNAT entries associated with the NAT gateway.

    - Disassociates any Elastic IP addresses (EIPs) bound to the NAT gateway.

    - Deletes any NAT bandwidth packages associated with the NAT gateway.
  - `false` (default): Does not force deletion. When this value is set to `false`, you must perform the following actions manually before deleting the NAT gateway:

    - Delete any NAT bandwidth packages associated with the NAT gateway.

    - Delete all SNAT entries.

    - Delete all DNAT entries.

    - Disassociate any bound EIPs.

-> **NOTE:** This parameter only takes effect when deletion is triggered.

* `icmp_reply_enabled` - (Optional, Computed, Available since v1.235.0) Indicates whether ICMP reply is enabled. Valid values:
  - `true` (default): Enabled.
  - `false`: Disabled.
* `internet_charge_type` - (Optional, ForceNew, Computed) The billing method of the NAT gateway instance. Valid values:
  - `PayBySpec`: Billed based on fixed specifications.
  - `PayByLcu`: Billed based on usage.
* `nat_gateway_name` - (Optional, Computed) The name of the NAT gateway.
The name must be 2 to 128 characters in length, start with a letter (uppercase or lowercase) or a Chinese character, and can contain digits, underscores (_), and hyphens (-).
If this parameter is not specified, the system assigns a default name to the NAT gateway.
* `nat_gateway_private_info` - (Required, ForceNew, Set, Available since v1.269.0) Private IP address information of the NAT gateway. See [`nat_gateway_private_info`](#nat_gateway_private_info) below.
* `nat_type` - (Required, ForceNew, Available since v1.102.0) The type of the public NAT gateway. The current valid value is `Enhanced`, which indicates an enhanced NAT gateway.
* `network_type` - (Required, ForceNew, Available since v1.136.0) The type of the created NAT gateway. Valid values:
  - `internet`: Public NAT gateway.
  - `intranet`: VPC NAT gateway.
* `payment_type` - (Optional, ForceNew, Computed) The billing method of the NAT gateway. Valid values:

  - *PostPaid** (default): Pay-as-you-go.

For more information, see [Billing of public NAT gateways](~~48126~~) and [Billing of VPC NAT gateways](~~270913~~).
* `private_link_enabled` - (Optional, Available since v1.235.0) PrivateLink is disabled by default. Setting this parameter to true enables PrivateLink.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `security_protection_enabled` - (Optional, Available since v1.269.0) Indicates whether the firewall feature is enabled. Valid values:
  - `false` (default): Disabled.

-> **NOTE:**  This parameter has been deprecated.>.


-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `spec` - (Optional, ForceNew) The specification of the Internet NAT gateway instance. When `InternetChargeType` is `PayBySpec`, the following values are returned:
  - `Small`: Small.
  - `Middle`: Medium.
  - `Large`: Large.

When `InternetChargeType` is `PayByLcu`, this field returns an empty value.
* `vpc_id` - (Required, ForceNew) The ID of the virtual private cloud (VPC) to which the NAT gateway belongs.

### `access_mode`

The access_mode supports the following:
* `mode_value` - (Optional, Computed) Valid values for the access mode:
  - `route`: Route mode.
  - `tunnel`: Tunnel mode.

-> **NOTE:**  When this field is set, the `PrivateLinkEnabled` field must be `true`.

* `tunnel_type` - (Optional, Computed) The tunnel mode type:
  - `geneve`: Geneve type.

-> **NOTE:**  This value is valid only when the access mode is set to tunnel mode.


### `nat_gateway_private_info`

The nat_gateway_private_info supports the following:
* `vswitch_id` - (Required, ForceNew, Available since v1.269.0) The ID of the vSwitch to which the NAT gateway belongs.

When you create a NAT gateway, you must specify the vSwitch to which the NAT gateway belongs. The system assigns an available private IP address from the vSwitch to the NAT gateway.
  - If you want to create a NAT gateway in an existing vSwitch, ensure that the zone where the vSwitch resides supports NAT gateway creation and that the vSwitch has available IP addresses.
  - If you have not created a vSwitch yet, first create a vSwitch in a zone that supports NAT gateway creation, and then specify this vSwitch for the NAT gateway.

-> **NOTE:**  You can call the [ListEnhancedNatGatewayAvailableZones](~~182292~~) operation to query the zones where NAT gateway resources are available, and call the [DescribeVSwitches](~~35748~~) operation to query the number of available IP addresses in a vSwitch.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time, expressed in ISO 8601 standard format using UTC time. The format is: YYYY-MM-DDThh:mmZ.
* `deletion_protection` - Deletion protection.
* `forward_table_ids` - The ID of the port forwarding table.
* `nat_gateway_private_info` - Private IP address information of the NAT gateway.
  * `eni_instance_id` - The elastic network interface (ENI) instance ID.
  * `eni_type` - The elastic network interface (ENI) type.
  * `iz_no` - Zone ID.
  * `max_bandwidth` - Maximum bandwidth.
  * `max_session_establish_rate` - Maximum number of new connections per second.
  * `max_session_quota` - Maximum number of concurrent connections.
  * `private_ip_address` - Private IP address.
* `region_id` - The region ID of the EIP that you want to associate with a cloud service instance.
* `snat_table_ids` - The SNAT table IDs of the NAT gateway.
* `status` - The status of the NAT gateway. , and the status remains `Creating` until the operation completes., and the status remains `Modifying` during the process., and the status remains `Deleting` during the process., and the status remains `Converting` during the process.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Nat Gateway.
* `delete` - (Defaults to 5 mins) Used when delete the Nat Gateway.
* `update` - (Defaults to 5 mins) Used when update the Nat Gateway.

## Import

N A T Gateway Nat Gateway can be imported using the id, e.g.

```shell
$ terraform import alicloud_nat_gateway.example <nat_gateway_id>
```