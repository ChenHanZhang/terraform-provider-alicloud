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
* `access_mode` - (Optional, ForceNew, Computed, Set, Available since v1.235.0) Configuration of the reverse access VPC-NAT access mode. See [`access_mode`](#access_mode) below.
* `auto_pay` - (Optional, Available since v1.270.0) Indicates whether auto-payment is enabled.
  - `true`: Auto-payment is enabled, and the system automatically pays for the order.
  - `false` (default): Auto-payment is disabled. After an order is generated, you must complete the payment in the Order Center.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `description` - (Optional) The description of the NAT gateway.
The description can be empty, or contain 2 to 256 characters and must not start with `http://` or `https://`.
* `eip_bind_mode` - (Required, ForceNew, Available since v1.184.0) The EIP binding mode of the NAT gateway. Valid values: This parameter can be left empty. If specified, it must be `NAT`, which indicates the standard EIP binding mode.


-> **NOTE:**  - You can only change the mode from `MULTI_BINDED` to `NAT`. Changing from `NAT` to `MULTI_BINDED` is not supported. For more information about the `MULTI_BINDED` mode, see [CreateNatGateway](https://help.aliyun.com/document_detail/120219.html).
  - During the EIP binding mode switch, network connections experience brief interruptions lasting several seconds (the duration increases with the number of EIPs; currently, configuration changes are supported only for NAT gateways bound to up to 5 EIPs). We recommend performing this operation during off-peak hours.
  - After switching the EIP binding mode to `NAT`, the public NAT gateway becomes compatible with IPv4 gateways. However, when one public EIP is bound, it consumes one private IP address from the vSwitch where the NAT gateway resides. Ensure that sufficient private IP addresses are available in the vSwitch. If no free private IP addresses remain in the vSwitch, you cannot bind additional EIPs.
* `enable_session_log` - (Optional, Available since v1.270.0) Specifies whether to enable session logging. Valid values:
  - `true`: Session logging is enabled.
  - `false`: Session logging is disabled.
* `force_delete` - (Optional, Available since v1.270.0) Specifies whether to forcibly delete the NAT gateway. Valid values:
  - `true`: Force deletion. If this value is specified, the system performs the following actions:

    - If the NAT gateway has SNAT rules, the system automatically deletes them.

    - If the NAT gateway has DNAT rules, the system automatically deletes them.

    - If the NAT gateway is associated with an EIP, the system automatically unbinds it.

    - If the NAT gateway has an undeleted NAT bandwidth package, the system automatically deletes it.
  - `false` (default): Do not force deletion. If this value is specified, you must perform the following actions before deleting the NAT gateway:

    - Delete any NAT bandwidth packages associated with the NAT gateway.

    - Delete all SNAT rules configured on the NAT gateway.

    - Delete all DNAT rules configured on the NAT gateway.

    - Unbind any EIPs associated with the NAT gateway.

-> **NOTE:** This parameter only takes effect when deletion is triggered.

* `icmp_reply_enabled` - (Optional, Computed, Available since v1.235.0) Specifies whether to enable ICMP reply.

Valid values:
  - `true` (default): Enable ICMP reply.
  - `false`: Disable ICMP reply.
* `internet_charge_type` - (Optional, ForceNew, Computed) The billing type of the NAT gateway instance. Valid values:
  - `PayBySpec`: billed based on fixed specifications.
  - `PayByLcu`: billed based on usage.
* `log_delivery` - (Optional, ForceNew, Set, Available since v1.270.0) Session log configuration information. See [`log_delivery`](#log_delivery) below.
* `nat_gateway_name` - (Optional, Computed) The name of the NAT gateway.
The name must be 2 to 128 characters in length, start with a letter (uppercase or lowercase) or a Chinese character, and can contain digits, underscores (_), and hyphens (-).
If this parameter is not specified, the system creates a default name for the NAT gateway.
* `nat_gateway_private_info` - (Required, ForceNew, Set, Available since v1.270.0) Private IP information of the NAT gateway. See [`nat_gateway_private_info`](#nat_gateway_private_info) below.
* `nat_type` - (Required, ForceNew, Available since v1.102.0) The type of the public NAT gateway. The current value is `Enhanced`, which refers to the enhanced NAT gateway.
* `network_type` - (Required, ForceNew, Available since v1.136.0) The type of NAT gateway to create. Valid values:
  - `internet`: Public NAT gateway.
  - `intranet`: VPC NAT gateway.
* `payment_type` - (Optional, ForceNew, Computed) The billing method of the NAT gateway. Valid values:

  - *PostPaid** (default): Pay-as-you-go.

For more information, see [Billing of public NAT gateways](https://help.aliyun.com/document_detail/48126.html) and [Billing of VPC NAT gateways](https://help.aliyun.com/document_detail/270913.html).
* `private_link_enabled` - (Optional, ForceNew, Available since v1.235.0) Specifies whether to enable PrivateLink. Valid values:
  - true: PrivateLink is enabled.
  - false (default): PrivateLink is disabled.
* `security_protection_enabled` - (Optional, Available since v1.270.0, Deprecated since v1.270.0) Indicates whether the firewall feature is enabled. Valid values:
  - `false` (default): Disabled.

-> **NOTE:**  This parameter is deprecated.>


-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `spec` - (Optional, ForceNew) The specification of the public NAT gateway instance. When `InternetChargeType` is set to `PayBySpec`, the following values are returned:
  - `Small`: Small.
  - `Middle`: Medium.
  - `Large`: Large.

When `InternetChargeType` is set to `PayByLcu`, this field returns an empty value.
* `vpc_id` - (Required, ForceNew) The ID of the Virtual Private Cloud (VPC) to which the NAT gateway belongs.

### `access_mode`

The access_mode supports the following:
* `mode_value` - (Optional, ForceNew, Computed) Access mode values:
  - `route`: Route mode.
  - `tunnel`: Tunnel mode.

-> **NOTE:**  When this field is set, the `PrivateLinkEnabled` field must be set to `true`.

* `tunnel_type` - (Optional, ForceNew, Computed) Tunnel mode type:
  - `geneve`: Geneve type.

-> **NOTE:**  This value is valid only when the access mode is set to tunnel mode.


### `log_delivery`

The log_delivery supports the following:
* `log_delivery_type` - (Optional, Available since v1.270.0) The type of session log delivery. Valid value: `sls`, which refers to Alibaba Cloud Log Service (SLS).
* `log_destination` - (Optional, Available since v1.270.0) The destination address for session log delivery. Value format: acs:log:${regionName}:${projectOwnerAliUid}:project/${projectName}/logstore/${logstoreName}.

### `nat_gateway_private_info`

The nat_gateway_private_info supports the following:
* `vswitch_id` - (Required, ForceNew, Available since v1.270.0) The ID of the vSwitch to which the NAT gateway belongs.

When creating a NAT gateway, you must specify the vSwitch to which the NAT gateway belongs. The system assigns an available private IP address from within the vSwitch to the NAT gateway.
  - If you want to create a NAT gateway in an existing vSwitch, ensure that the zone where the vSwitch resides supports NAT gateway creation and that the vSwitch has available IP addresses.
  - If you have not created a vSwitch yet, first create one in a zone that supports NAT gateway creation, and then specify this vSwitch for the NAT gateway.

-> **NOTE:**  You can call the [ListEnhancedNatGatewayAvailableZones](https://help.aliyun.com/document_detail/182292.html) operation to query the zones where NAT gateway resources are available, and call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query the number of available IP addresses in a vSwitch.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time.
* `deletion_protection` - Deletion protection.
* `forward_table_ids` - The ID of the DNAT table.
* `log_delivery` - Session log configuration information.
  * `deliver_logs_error_message` - Error message for session log delivery failure.
  * `delivery_status` - The status of session log delivery.
* `nat_gateway_private_info` - Private IP information of the NAT gateway.
  * `eni_instance_id` - The ID of the elastic network interface (ENI) instance.
  * `eni_type` - The type of the elastic network interface (ENI).
  * `iz_no` - The zone ID of the NAT gateway instance.
  * `max_bandwidth` - The maximum bandwidth.
  * `max_session_establish_rate` - The maximum number of new connections per second.
  * `max_session_quota` - The maximum number of concurrent connections.
  * `private_ip_address` - The private IP address.
* `region_id` - The region ID of the Elastic IP address (EIP) that you want to associate with a cloud service instance.
* `snat_table_ids` - The ID of the SNAT table.
* `status` - The status of the NAT gateway.

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