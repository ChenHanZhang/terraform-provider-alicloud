---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_network_interface"
description: |-
  Provides a Alicloud ECS Network Interface resource.
---

# alicloud_ecs_network_interface

Provides a ECS Network Interface resource.

Elastic Network Interface  .

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
  security_group_name = var.name
  vpc_id              = alicloud_vpc.default.id
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
* `delete_on_release` - (Optional, Computed, Available since v1.273.0) Specifies whether to retain the network interface when the associated instance is released. Valid values:
  - true: Do not retain the network interface.
  - false: Retain the network interface.
* `description` - (Optional) The description of the elastic network interface. It must be 2 to 256 characters in length, and cannot start with `http://` or `https://`.
Default value: empty.
* `ipv4_prefix` - (Optional, ForceNew, List, Available since v1.273.0) Specify one or more IPv4 prefixes for the elastic network interface. Valid values for N: 1 to 10.  

-> **NOTE:**  If you need to assign IPv4 prefixes to the elastic network interface, you must specify either Ipv4Prefix.N or Ipv4PrefixCount, but not both.

* `ipv6_prefix` - (Optional, ForceNew, List, Available since v1.273.0) One or more IPv6 prefixes to be reclaimed. The value range for N is 1 to 10.
* `ipv6_sets` - (Optional, List, Available since v1.273.0) Specify one or more IPv6 addresses for the elastic network interface. You can configure up to 10 IPv6 addresses, so the value range for N is 1 to 10.

Example value: Ipv6Address.1=2001:db8:1234:1a00::****

-> **NOTE:**  When calling this API operation, you must specify either the `Ipv6Addresses.N` parameter or the `Ipv6AddressCount` parameter, but not both.
 See [`ipv6_sets`](#ipv6_sets) below.
* `network_interface_name` - (Optional, Computed) The name of the elastic network interface. The name must be 2 to 128 characters in length, start with a letter (uppercase or lowercase) or a Chinese character, and cannot start with `http://` or `https://`. It can contain letters (including English, Chinese, and digits) as defined in the Unicode "Letter" category, as well as colons (:), underscores (_), periods (.), or hyphens (-).
Default value: empty.
* `primary_ip_address` - (Optional, ForceNew, Computed) The primary private IP address of the elastic network interface.
The specified IP address must be an available IP address within the CIDR block of the associated vSwitch. If no IP address is specified, an available IP address from the vSwitch is randomly assigned by default.
* `private_ip_sets` - (Optional, List, Available since v1.273.0) A collection of PrivateIpSet entries. See [`private_ip_sets`](#private_ip_sets) below.
* `qo_s_config` - (Optional, Set, Available since v1.273.0) QoS rate limiting settings. See [`qo_s_config`](#qo_s_config) below.
* `queue_number` - (Optional, Computed, Int) The number of queues for the ENI.

  - If the ENI is a primary ENI, the default number of queues for the instance type is returned.

  - If the ENI is a secondary ENI:
    * If the secondary ENI is in the InUse state:
        * If the queue number has not been modified, the default number of queues for secondary ENIs of the instance type is returned.
        * If the queue number has been modified, the modified value is returned.
    * If the secondary ENI is in the Available state:
        * If the queue number has not been modified, an empty value is returned.
        * If the queue number has been modified, the modified value is returned.
* `resource_group_id` - (Optional, ForceNew, Computed) The ID of the enterprise resource group to which the instance belongs. When you use this parameter to filter resources, the number of returned resources cannot exceed 1,000.

-> **NOTE:**  Filtering by the default resource group is not supported.

* `resource_type` - (Optional, Available since v1.273.0) The resource type. Valid values:
  - instance: ECS instance.
  - disk: Disk.
  - snapshot: Snapshot.
  - image: Image.
  - securitygroup: Security group.
  - volume: Volume.
  - eni: Elastic Network Interface (ENI).
  - ddh: Dedicated Host (DDH).
  - ddhcluster: Dedicated Host cluster.
  - keypair: SSH key pair.
  - launchtemplate: Launch template.
  - reservedinstance: Reserved Instance.
  - snapshotpolicy: Automatic snapshot policy.
  - elasticityassurance: Elasticity Assurance.
  - capacityreservation: Capacity Reservation.
  - command: Cloud Assistant command.
  - invocation: Cloud Assistant command execution or file delivery result.
  - activation: Cloud Assistant managed instance activation code.
  - managedinstance: Cloud Assistant managed instance.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `security_group_ids` - (Optional, Computed, List) The IDs of one or more security groups to which the ENI belongs. The security groups and the ENI must reside in the same Virtual Private Cloud (VPC). The valid range for N depends on the quota for the number of security groups that an ENI can join. For more information, see [Limits](https://help.aliyun.com/document_detail/25412.html).

-> **NOTE:**  When calling the API, you must specify either `SecurityGroupId` or `SecurityGroupIds.N`, but not both.

* `tags` - (Optional, Map) The list of tags for the elastic network interface.
* `vswitch_id` - (Required, ForceNew) The vSwitch ID of the elastic network interface. The private IP address of the elastic network interface is assigned from the available IP addresses within the CIDR block of the specified vSwitch.

-> **NOTE:** The elastic network interface and the instance to be attached must reside in the same zone but can belong to different vSwitches.


### `ipv6_sets`

The ipv6_sets supports the following:
* `ipv6_address` - (Optional, Available since v1.273.0) The IPv6 address specified for the elastic network interface.

### `private_ip_sets`

The private_ip_sets supports the following:
* `private_ip_address` - (Optional, Available since v1.273.0) The private IP address of the instance.

### `qo_s_config`

The qo_s_config supports the following:
* `enable_qo_s` - (Optional, Available since v1.273.0) Specifies whether to enable QoS rate limiting.
* `qo_s` - (Optional, Set, Available since v1.273.0) QoS rate limiting settings. See [`qo_s`](#qo_s_config-qo_s) below.

### `qo_s_config-qo_s`

The qo_s_config-qo_s supports the following:
* `bandwidth_rx` - (Optional, Int, Available since v1.273.0) Maximum inbound bandwidth limit over the internal network.
Unit: kbit/s, step size: 1 000 (1 Mbps), valid range: [50 000, +∞).
* `bandwidth_tx` - (Optional, Int, Available since v1.273.0) Maximum outbound bandwidth limit over the internal network.
Unit: kbit/s, step size: 1 000 (1 Mbps), valid range: [50 000, +∞).
* `concurrent_connections` - (Optional, Int, Available since v1.273.0) The maximum number of concurrent sessions.
Step size: 10,000. Valid range: [10,000, +∞).
* `pps_rx` - (Optional, Int, Available since v1.273.0) Inbound packet processing capability over the internal network.
Unit: pps, step size: 10 000, valid range: [10 000, +∞).
* `pps_tx` - (Optional, Int, Available since v1.273.0) Outbound packet processing capability over the internal network.
Unit: pps, step size: 10 000, valid range: [10 000, +∞).

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `private_ip_sets` - A collection of PrivateIpSet entries.
  * `associated_public_ip` - The Elastic IP Address (EIP) associated with the elastic network interface.
    * `allocation_id` - > This parameter is currently in邀测 and is not publicly available for use.
    * `public_ip_address` - The Elastic IP Address (EIP).
  * `primary` - Indicates whether this is the primary private IP address.
* `region_id` - The ID of the region where the elastic network interface is to be created.
* `status` - The status of the elastic network interface.

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