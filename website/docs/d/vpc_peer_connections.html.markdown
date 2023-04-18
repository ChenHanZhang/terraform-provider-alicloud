---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_peer_connections"
sidebar_current: "docs-alicloud-datasource-vpc-peer-connections"
description: |-
  Provides a list of Vpc Peer Connection owned by an Alibaba Cloud account.
---

# alicloud_vpc_peer_connections

This data source provides Vpc Peer Connection available to the user.[What is Peer Connection](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_peer_connections" "default" {
  ids                  = ["${alicloud_vpc_peer_connection.default.id}"]
  name_regex           = alicloud_vpc_peer_connection.default.name
  peer_connection_name = "rmc-test"
  vpc_id               = "vpc-bp1m0q8m09ws87cfbdwvs"
}

output "alicloud_vpc_peer_connection_example_id" {
  value = data.alicloud_vpc_peer_connections.default.connections.0.id
}
```

## Argument Reference

The following arguments are supported:
* `peer_connection_name` - (ForceNew,Optional) The name of the resource
* `peering_id` - (ForceNew,Optional) The first ID of the resource
* `tags` - (ForceNew,Optional) The tags of PrefixList.See the following `Block Tags`.
* `vpc_id` - (ForceNew,Optional) You must create a VPC ID on the initiator of a VPC peer connection.
* `ids` - (Optional, ForceNew, Computed) A list of Peer Connection IDs.
* `peer_connection_names` - (Optional, ForceNew) The name of the Peer Connection. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

#### Block Tags

The Tags supports the following:
* `tag_key` - (ForceNew,Optional) The key of tag.
* `tag_value` - (ForceNew,Optional) The value of tag.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Peer Connection IDs.
* `names` - A list of name of Peer Connections.
* `connections` - A list of Peer Connection Entries. Each element contains the following attributes:
  * `accepting_ali_uid` - The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.-Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.-Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.> If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
  * `accepting_region_id` - The region ID of the recipient of the VPC peering connection to be created.-When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.-When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
  * `accepting_vpc_id` - The VPC ID of the receiving end of the VPC peer connection.
  * `bandwidth` - The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
  * `biz_status` - The business status of the VPC peer connection. Value:-**Normal**: Normal.-**Financialized**: Arrears locked.
  * `create_time` - The creation time of the VPC peer connection. Use UTC time in the format' YYYY-MM-DDThh:mm:ssZ '.
  * `description` - The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with http:// or https.
  * `expire_time` - The expiration time of the VPC peer connection.The specific expiration time is returned only when the **Status** (Status) of the VPC peer connection is **Accepting** (receiving) or **Expired** (Expired). When the VPC peer connection is in the remaining states, the return value is **null * *.
  * `modify_time` - The modification time of the VPC peer connection. Use UTC time in the format' YYYY-MM-DDThh:mm:ssZ '.
  * `peer_connection_name` - The name of the resource
  * `peering_id` - The first ID of the resource
  * `status` - The status of the resource
  * `tags` - The tags of PrefixList.
    * `tag_key` - The key of tag.
    * `tag_value` - The value of tag.
  * `vpc_id` - You must create a VPC ID on the initiator of a VPC peer connection.
