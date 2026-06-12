---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc"
description: |-
  Provides a Alicloud VPC VPC resource.
---

# alicloud_vpc

Provides a VPC VPC resource.

A VPC instance represents a virtual private cloud that you have created. You have full control over your VPC, such as selecting the IP address range, configuring route tables and gateways, and using Alibaba Cloud resources like Elastic Compute Service (ECS), ApsaraDB RDS, and Server Load Balancer within your custom VPC.

For information about VPC VPC and how to use it, see [What is VPC](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/what-is-a-vpc).

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}


resource "alicloud_vpc" "default" {
  description = "example description"
  cidr_block  = "10.0.0.0/8"
  vpc_name    = var.name
  enable_ipv6 = true
}
```

## Argument Reference

The following arguments are supported:
* `cidr_block` - (Optional, Computed) The CIDR block of the VPC.
  - We recommend that you use private IPv4 addresses defined in RFC 1918 as the primary IPv4 CIDR block for the VPC, with a subnet mask between /16 and /28. Examples include 10.0.0.0/16, 172.16.0.0/16, and 192.168.0.0/16.
  - You can also use custom CIDR blocks as the primary IPv4 CIDR block for the VPC, except for 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, and their subnets.
* `classic_link_enabled` - (Optional) Indicates whether the ClassicLink feature is enabled. Valid values:
  - `true`: Enabled.
  - `false` (default): Disabled.
* `description` - (Optional) The description of the VPC.
The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
* `dns_hostname_status` - (Optional, Computed, Available since v1.240.0) Specifies whether to enable the DNS hostname feature. Valid values:
  - `false` (default): Disabled.
  - `true`: Enabled.
* `dry_run` - (Optional, Available since v1.119.0) Specifies whether to only precheck the request. Valid values:
  - `true`: Sends a dry-run request without creating a VPC. The check includes validation of required parameters, request format, and service limits. If the check fails, an error is returned. If the check passes, the error code `DryRunOperation` is returned.
  - `false` (default): Sends a normal request. If the check passes, an HTTP 2xx status code is returned and the VPC is created.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `enable_ipv6` - (Optional, Computed, Available since v1.119.0) Specifies whether to enable an IPv6 CIDR block. Valid values:
  - `false` (default): Disabled.
  - `true`: Enabled.
* `force_delete` - (Optional, Available since v1.248.0) Specifies whether to forcibly delete the VPC. Valid values:
  - `true`: Forces deletion.
  - `false` (default): Does not force deletion.

You can forcibly delete a VPC only if it contains one of the following resources:
  - Only an IPv4 gateway and routes pointing to the IPv4 gateway.
  - Only an IPv6 gateway and routes pointing to the IPv6 gateway.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `ipv4_cidr_mask` - (Optional, Int, Available since v1.240.0) Allocate a VPC from an IPAM address pool by specifying a subnet mask.

-> **NOTE:**  When creating a VPC from an IPAM address pool, you must specify at least one of the parameters CidrBlock or Ipv4CidrMask.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv4_ipam_pool_id` - (Optional) The ID of the IPv4 IPAM address pool instance.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv6_cidr_block` - (Optional, ForceNew, Computed, Deprecated since v1.282.0) The IPv6 CIDR block of the VPC. When IPv6 is enabled for a VPC, the system assigns an IPv6 CIDR block. To specify an IPv6 CIDR block, you must first call the [AllocateVpcIpv6Cidr](https://help.aliyun.com/document_detail/448916.html) operation to reserve the specified IPv6 address block and then pass it in.
* `ipv6_isp` - (Optional, Computed, Deprecated since v1.282.0) The type of IPv6 address block for the VPC. Valid values:
  - `BGP` (default): Alibaba Cloud BGP IPv6.
  - `ChinaMobile`: China Mobile (single ISP).
  - `ChinaUnicom`: China Unicom (single ISP).
  - `ChinaTelecom`: China Telecom (single ISP).

-> **NOTE:**  If you are a user who has been granted access to single-ISP bandwidth through a whitelist, you can set this parameter to `ChinaTelecom` (China Telecom), `ChinaUnicom` (China Unicom), or `ChinaMobile` (China Mobile).

* `resource_group_id` - (Optional, Computed, Available since v1.115) The ID of the resource group into which the cloud resource instance is to be moved.

-> **NOTE:**  A resource group is a mechanism for grouping and managing resources under an Alibaba Cloud account. Resource groups help you address complexities related to resource grouping and authorization management within a single Alibaba Cloud account. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html).

* `route_table_id` - (Optional, ForceNew, Computed) The ID of the route table automatically created by the system after a default VPC is created.
* `secondary_cidr_blocks` - (Optional, Computed, List, Deprecated since v1.185.0) The additional IPv4 CIDR blocks to be added. The CIDR blocks must meet the following requirements:
  - We recommend that you use private IPv4 addresses specified in RFC 1918 as additional IPv4 CIDR blocks for the VPC. The recommended subnet mask length is between 16 and 28 bits. Examples: 10.0.0.0/16, 172.16.0.0/16, and 192.168.0.0/16.
  - You can also use custom CIDR blocks outside of 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, and their subnets as additional IPv4 CIDR blocks for the VPC.

Configuration limits:
  - The CIDR block cannot start with 0, and the recommended subnet mask length is between 16 and 28 bits.
  - The additional CIDR block must not overlap with the primary CIDR block of the VPC or any previously added additional CIDR blocks.

-> **NOTE:**  When you add an additional CIDR block to a VPC without using an IPAM address pool, you must specify either the `SecondaryCidrBlock` parameter or the `Ipv6CidrBlock` parameter, but not both.

* `system_route_table_description` - (Optional) The description of the route table.
The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
* `system_route_table_name` - (Optional) The name of the route table.
The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
* `system_route_table_route_propagation_enable` - (Optional, Computed, Available since v1.248.0) You can control whether dynamic routes are received by enabling or disabling route propagation. Valid values:
  - `true` (default): Enable.
  - `false`: Disable.
* `tags` - (Optional, Map, Available since v1.55.3) Tags of the resource.
* `user_cidrs` - (Optional, ForceNew, Computed, List, Available since v1.119.0) User CIDR blocks. To specify multiple CIDR blocks, separate them with commas (,). A maximum of three CIDR blocks are supported.
For more information about user CIDR blocks, see "What is a user CIDR block?" in the [VPC FAQ](https://help.aliyun.com/document_detail/185311.html).
* `vpc_name` - (Optional, Computed, Available since v1.119.0) The name of the VPC.
The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - (Deprecated since v1.119.0). Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
* `router_table_id` - (Deprecated since v1.282.0). Field 'router_table_id' has been deprecated from provider version 1.282.0. New field 'route_table_id' instead.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the VPC.
* `dhcp_options_set_id` - The ID of the DHCP options set.
* `ipv6_cidr_blocks` - The IPv6 CIDR block information of the VPC.
  * `ipv6_cidr_block` - The IPv6 CIDR block of the VPC.
  * `ipv6_isp` - **BGP** (default): Alibaba Cloud BGP IPv6.
* `is_default` - Indicates whether the VPC is the default VPC.
* `router_id` - The router ID of the VPC.
* `status` - The status of the VPC.
* `vswitch_ids` - A list of vSwitches in the VPC.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the VPC.
* `delete` - (Defaults to 5 mins) Used when delete the VPC.
* `update` - (Defaults to 5 mins) Used when update the VPC.

## Import

VPC VPC can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc.example <vpc_id>
```