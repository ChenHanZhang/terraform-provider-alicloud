---
subcategory: "Vpc"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_prefix_list"
sidebar_current: "docs-alicloud-resource-vpc-prefix-list"
description: |-
  Provides a Alicloud Vpc Prefix List resource.
---

# alicloud_vpc_prefix_list

Provides a Vpc Prefix List resource.

For information about Vpc Prefix List and how to use it, see [What is Prefix List](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_prefix_list" "default" {
  max_entries             = 50
  prefix_list_description = "test"
  ip_version              = "ipv4"
  prefix_list_name        = "rdk-test2"
  region_id               = "cn-hangzhou"
  entries {
    cidr        = "192.168.0.0/16"
    description = "test16"
  }
  entries {
    cidr        = "192.168.0.0/15"
    description = "test15"
  }
  tags {
    tag_key   = "k0"
    tag_value = "v0"
  }
}
```

## Argument Reference

The following arguments are supported:
* `entries` - (Computed,Optional) The CIDR address block list of the prefix list.See the following `Block Entries`.
* `ip_version` - (ForceNew,Computed,Optional) The IP version of the prefix list. Value:-**ipv4**:IPv4 version.-**ipv6**:IPv6.
* `max_entries` - (Computed,Optional) The maximum number of entries for CIDR address blocks in the prefix list.
* `prefix_list_description` - (Optional) The description of the prefix list.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with http:// or https.
* `prefix_list_name` - (Optional) The name of the prefix list.
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the VPC belongs.
* `tags` - (Optional) The tags of PrefixList.See the following `Block Tags`.

The following arguments will be discarded. Please use new fields as soon as possible:

#### Block Entries

The Entries supports the following:
* `cidr` - (Optional) cidr block.
* `description` - (Optional) the description of the cidr entry.

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of tag.
* `tag_value` - (Optional) The value of tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `create_time` - The time when the prefix list was created.
* `entries` - The CIDR address block list of the prefix list.
* `ip_version` - The IP version of the prefix list. Value:-**ipv4**:IPv4 version.-**ipv6**:IPv6.
* `max_entries` - The maximum number of entries for CIDR address blocks in the prefix list.
* `owner_id` - The Alibaba Cloud account (primary account) to which the prefix list belongs.
* `prefix_list_association` - The association list information of the prefix list.
  * `owner_id` - The ID of the Alibaba Cloud account (primary account) to which the prefix list belongs.
  * `prefix_list_id` - The instance ID of the prefix list.
  * `reason` - Reason when the association fails.
  * `region_id` - The region ID of the prefix list to be queried.
  * `resource_id` - The ID of the associated resource.
  * `resource_type` - The associated resource type. Value:-**Vprouter table**: The VPC route table.-**trRouteTable**: the routing table of the forwarding router.
  * `resource_uid` - The ID of the Alibaba Cloud account (primary account) to which the resource bound to the prefix list belongs.
  * `status` - The association status of the prefix list. Value:-**Created**: Success.-**ModifyFailed**: is not associated with the latest version.-**Creating**: Creating.-**Modifying**: Modifying.-**Deleting**: Deleting.-**Deleted**: Deleted.
* `prefix_list_id` - The ID of the query Prefix List.
* `resource_group_id` - The ID of the resource group to which the VPC belongs.
* `share_type` - The share type of the prefix list. Value:-**Shared**: indicates that the prefix list is a Shared prefix list.-Null: indicates that the prefix list is not a shared prefix list.
* `status` - The status of the resource

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Prefix List.
* `delete` - (Defaults to 5 mins) Used when delete the Prefix List.
* `update` - (Defaults to 5 mins) Used when update the Prefix List.

## Import

Vpc Prefix List can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_prefix_list.example 
```