---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_vpc_firewall_cen"
description: |-
  Provides a Alicloud Cloud Firewall Vpc Firewall Cen resource.
---

# alicloud_cloud_firewall_vpc_firewall_cen

Provides a Cloud Firewall Vpc Firewall Cen resource.

VPC cloud firewall cloud enterprise network.

For information about Cloud Firewall Vpc Firewall Cen and how to use it, see [What is Vpc Firewall Cen](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcenconfigure).

-> **NOTE:** Available since v1.194.0.

## Example Usage

Basic Usage

```terraform
# These resource primary keys should be replaced with your actual values.
resource "alicloud_cloud_firewall_vpc_firewall_cen" "default" {
  cen_id = "cen-xxx"
  local_vpc {
    network_instance_id = "vpc-xxx"
  }
  status            = "open"
  member_uid        = "14151*****827022"
  vpc_region        = "cn-hangzhou"
  vpc_firewall_name = "tf-vpc-firewall-name"
}
```

## Argument Reference

The following arguments are supported:
* `cen_id` - (Required, ForceNew) The ID of the CEN instance.
* `lang` - (Optional) The language type of the requested and received messages. Value:
  - `zh` (default): Chinese.
  - `en`: English.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `local_vpc` - (Required, ForceNew, Set) The details of the VPC. See [`local_vpc`](#local_vpc) below.
* `network_instance_id` - (Optional, ForceNew, Available since v1.286.0) The ID of the VPC instance that created the VPC firewall.
* `status` - (Required) Firewall switch status
* `vpc_firewall_name` - (Required) The name of the VPC firewall instance.
* `vpc_region` - (Required) The ID of the region to which the VPC is created.

-> **NOTE:**  For more information about supported regions, see [Supported Regions](~~ 195657 ~~).


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


### `local_vpc`

The local_vpc supports the following:
* `network_instance_id` - (Required, ForceNew) The ID of the VPC instance that created the VPC firewall.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `connect_type` - Intercommunication type, value: expressconnect: Express Channel cen: Cloud Enterprise Network.
* `local_vpc` - The details of the VPC.
  * `attachment_id` - The connection ID of the network instance.
  * `attachment_name` - The connection name of the network instance.
  * `defend_cidr_list` - The list of network segments protected by the VPC firewall.
  * `eni_list` - List of elastic network cards.
    * `eni_id` - The ID of the instance of the ENI in the VPC.
    * `eni_private_ip_address` - The private IP address of the ENI in the VPC.
  * `manual_vswitch_id` - The ID of the vSwitch specified when the routing mode is manual mode.
  * `network_instance_name` - The name of the network instance.
  * `network_instance_type` - The type of the network instance.
  * `owner_id` - The UID of the Alibaba Cloud account to which the VPC belongs.
  * `region_no` - The region ID of the VPC.
  * `route_mode` - Routing mode,.
  * `support_manual_mode` - Whether routing mode supports manual mode.
  * `transit_router_id` - The ID of the CEN-TR instance.
  * `transit_router_type` - The version of the cloud enterprise network forwarding router (CEN-TR).
  * `vpc_cidr_table_list` - The VPC network segment list.
    * `route_entry_list` - The list of route entries in the VPC.
        * `destination_cidr` - The target network segment of the VPC.
        * `next_hop_instance_id` - The ID of the next hop instance in the VPC.
    * `route_table_id` - The ID of the route table of the VPC.
  * `vpc_id` - The ID of the VPC instance.
  * `vpc_name` - The instance name of the VPC.
* `vpc_firewall_id` - VPC firewall ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 31 mins) Used when create the Vpc Firewall Cen.
* `delete` - (Defaults to 31 mins) Used when delete the Vpc Firewall Cen.
* `update` - (Defaults to 31 mins) Used when update the Vpc Firewall Cen.

## Import

Cloud Firewall Vpc Firewall Cen can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_vpc_firewall_cen.example <vpc_firewall_id>
```