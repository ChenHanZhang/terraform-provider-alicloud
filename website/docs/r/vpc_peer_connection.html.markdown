---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_peer_connection"
sidebar_current: "docs-alicloud-resource-vpc-peer-connection"
description: |-
  Provides a Alicloud Vpc Peer Connection resource.
---

# alicloud_vpc_peer_connection

Provides a Vpc Peer Connection resource.

For information about Vpc Peer Connection and how to use it, see [What is Peer Connection](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_peer_connection" "default" {
  description          = "rmc-test"
  region_id            = "cn-hangzhou"
  accepting_region_id  = "cn-hangzhou"
  accepting_ali_uid    = 1891593620094065
  peer_connection_name = "rmc-test"
  accepting_vpc_id     = "vpc-bp1qfrofhdeabhbjmhv7o"
  vpc_id               = "vpc-bp1m0q8m09ws87cfbdwvs"
}
```

## Argument Reference

The following arguments are supported:
* `accepting_ali_uid` - (ForceNew,Optional) The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.-Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.-Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.> If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
* `accepting_region_id` - (Required,ForceNew) The region ID of the recipient of the VPC peering connection to be created.-When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.-When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
* `accepting_vpc_id` - (Required,ForceNew) The VPC ID of the receiving end of the VPC peer connection.
* `bandwidth` - (Computed,Optional) The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
* `delete_all` - (Optional) Delete All tags due to mapping UntagResources input parameter All
* `description` - (Optional) The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with http:// or https.
* `peer_connection_name` - (Optional) The name of the resource
* `status` - (Computed,Optional) The status of the resource
* `tags` - (Optional) The tags of PrefixList.See the following `Block Tags`.
* `vpc_id` - (Required,ForceNew) You must create a VPC ID on the initiator of a VPC peer connection.

The following arguments will be discarded. Please use new fields as soon as possible:

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of tag.
* `tag_value` - (Optional) The value of tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `bandwidth` - The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
* `biz_status` - The business status of the VPC peer connection. Value:-**Normal**: Normal.-**Financialized**: Arrears locked.
* `create_time` - The creation time of the VPC peer connection. Use UTC time in the format' YYYY-MM-DDThh:mm:ssZ '.
* `expire_time` - The expiration time of the VPC peer connection.The specific expiration time is returned only when the **Status** (Status) of the VPC peer connection is **Accepting** (receiving) or **Expired** (Expired). When the VPC peer connection is in the remaining states, the return value is **null * *.
* `modify_time` - The modification time of the VPC peer connection. Use UTC time in the format' YYYY-MM-DDThh:mm:ssZ '.
* `peering_id` - The first ID of the resource
* `status` - The status of the resource

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Peer Connection.
* `delete` - (Defaults to 5 mins) Used when delete the Peer Connection.
* `update` - (Defaults to 5 mins) Used when update the Peer Connection.

## Import

Vpc Peer Connection can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_peer_connection.example 
```