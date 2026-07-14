---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_vpc_firewall"
description: |-
  Provides a Alicloud Cloud Firewall Vpc Firewall resource.
---

# alicloud_cloud_firewall_vpc_firewall

Provides a Cloud Firewall Vpc Firewall resource.

Express Connect for VPC Firewall.

For information about Cloud Firewall Vpc Firewall and how to use it, see [What is Vpc Firewall](https://www.alibabacloud.com/help/en/cloud-firewall/developer-reference/api-cloudfw-2017-12-07-createvpcfirewallconfigure).

-> **NOTE:** Available since v1.194.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_account" "current" {
}

resource "alicloud_cloud_firewall_vpc_firewall" "default" {
  vpc_firewall_name = "tf-example"
  member_uid        = data.alicloud_account.current.id
  local_vpc {
    vpc_id    = "vpc-bp1d065m6hzn1xbw8ibfd"
    region_no = "cn-hangzhou"
    local_vpc_cidr_table_list {
      local_route_table_id = "vtb-bp1lj0ddg846856chpzrv"
      local_route_entry_list {
        local_next_hop_instance_id = "ri-bp1uobww3aputjlwwkyrh"
        local_destination_cidr     = "10.1.0.0/16"
      }
    }
  }
  peer_vpc {
    vpc_id    = "vpc-bp1gcmm64o3caox84v0nz"
    region_no = "cn-hangzhou"
    peer_vpc_cidr_table_list {
      peer_route_table_id = "vtb-bp1f516f2hh4sok1ig9b5"
      peer_route_entry_list {
        peer_destination_cidr     = "10.0.0.0/16"
        peer_next_hop_instance_id = "ri-bp1thhtgf6ydr2or52l3n"
      }
    }
  }
  status = "open"
}
```

## Argument Reference

The following arguments are supported:
* `lang` - (Optional) The language of the request and response. Valid values:
  - `zh` (default): Chinese.
  - `en`: English.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `local_vpc` - (Required, ForceNew, Set) Details of the local VPC. See [`local_vpc`](#local_vpc) below.
* `local_vpc_cidr_table_list` - (Required, Available since v1.286.0) The CIDR block list of the local VPC, represented in JSON format. It includes the following parameters:
  - `RouteTableId`: The route table ID of the local VPC.
  - `RouteEntryList`: This parameter is in JSON format and includes DestinationCidr (the destination CIDR block of the local VPC) and NextHopInstanceId (the next-hop instance ID of the local VPC).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `member_uid` - (Optional, ForceNew) The UID of the Alibaba Cloud member account.
* `peer_vpc` - (Required, ForceNew, Set) The details of the peer VPC. See [`peer_vpc`](#peer_vpc) below.
* `peer_vpc_cidr_table_list` - (Required, Available since v1.286.0) The list of CIDR blocks for the peer VPC, represented in JSON format. It includes the following parameters:
  - `RouteTableId`: The ID of the route table for the peer VPC.
  - `RouteEntryList`: This parameter is in JSON format and includes DestinationCidr (the destination CIDR block of the peer VPC) and NextHopInstanceId (the ID of the next-hop instance in the peer VPC).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `status` - (Required) The status of the VPC firewall. Valid values:
  - `opened`: Enabled.
  - `closed`: Disabled.
  - `notconfigured`: The VPC firewall is not configured.
  - `configured`: The VPC firewall is configured.
* `vpc_firewall_name` - (Required) The instance name of the VPC firewall.

### `local_vpc`

The local_vpc supports the following:
* `region_no` - (Required, ForceNew) The ID of the region to which the local VPC belongs.

-> **NOTE:**  For more information about regions supported by Cloud Firewall, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).

* `vpc_id` - (Required, ForceNew) The instance ID of the local VPC.

### `peer_vpc`

The peer_vpc supports the following:
* `region_no` - (Required, ForceNew) The region ID of the peer VPC.

-> **NOTE:**  For more information about the regions supported by Cloud Firewall, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).

* `vpc_id` - (Required, ForceNew) The instance ID of the peer VPC.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `bandwidth` - The bandwidth specification of the Express Connect circuit.
* `connect_type` - The connection type of the VPC firewall.
* `local_vpc` - Details of the local VPC.
  * `eni_id` - The instance ID of the elastic network interface (ENI) in the local VPC.
  * `eni_private_ip_address` - The private IP address of the ENI in the local VPC.
  * `local_vpc_cidr_table_list` - The list of CIDR blocks for the local VPC, represented in JSON format.
    * `local_route_entry_list` - The route entry list of the local VPC.
        * `local_destination_cidr` - The destination CIDR block of the local VPC.
        * `local_next_hop_instance_id` - The ID of the next-hop instance in the local VPC.
    * `local_route_table_id` - The route table ID of the local VPC.
  * `router_interface_id` - The ID of the router interface in the local VPC.
  * `vpc_name` - The instance name of the local VPC.
* `peer_vpc` - The details of the peer VPC.
  * `eni_id` - The instance ID of the elastic network interface (ENI) in the peer VPC.
  * `eni_private_ip_address` - The private IP address of the ENI in the peer VPC.
  * `peer_vpc_cidr_table_list` - The list of CIDR blocks for the peer VPC, represented in JSON format.
    * `peer_route_entry_list` - The route entry list information of the peer VPC.
        * `peer_destination_cidr` - The destination CIDR block of the peer VPC.
        * `peer_next_hop_instance_id` - The ID of the next-hop instance in the peer VPC.
    * `peer_route_table_id` - The route table ID of the peer VPC.
  * `router_interface_id` - The ID of the router interface in the peer VPC.
  * `vpc_name` - The instance name of the peer VPC.
* `region_status` - The availability status of the region.
* `vpc_firewall_id` - The ID of the access control policy group for the VPC firewall.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 31 mins) Used when create the Vpc Firewall.
* `delete` - (Defaults to 31 mins) Used when delete the Vpc Firewall.
* `update` - (Defaults to 31 mins) Used when update the Vpc Firewall.

## Import

Cloud Firewall Vpc Firewall can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_vpc_firewall.example <vpc_firewall_id>
```