---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_prefix_list"
description: |-
  Provides a Alicloud VPC Prefix List resource.
---

# alicloud_vpc_prefix_list

Provides a VPC Prefix List resource.

This resource is used to create a prefix list.

For information about VPC Prefix List and how to use it, see [What is Prefix List](https://www.alibabacloud.com/help/zh/virtual-private-cloud/latest/creatvpcprefixlist).

-> **NOTE:** Available since v1.182.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-testacc-example"
}

resource "alicloud_resource_manager_resource_group" "defaultRg" {
  display_name        = "tf-testacc-chenyi"
  resource_group_name = var.name
}

resource "alicloud_resource_manager_resource_group" "changeRg" {
  display_name        = "tf-testacc-chenyi-change"
  resource_group_name = "${var.name}1"
}


resource "alicloud_vpc_prefix_list" "default" {
  max_entries             = 50
  resource_group_id       = alicloud_resource_manager_resource_group.defaultRg.id
  prefix_list_description = "test"
  ip_version              = "IPV4"
  prefix_list_name        = var.name
  entrys {
    cidr        = "192.168.0.0/16"
    description = "test"
  }
}
```

## Argument Reference

The following arguments are supported:
* `entries` - (Optional, Computed, List, Available since v1.274.0) The CIDR address block list of the prefix list. See [`entries`](#entries) below.
* `ip_version` - (Optional, ForceNew, Computed) The IP version of the prefix list. Value:
  - `IPV4`:IPv4 version.
  - `IPV6`:IPv6 version.
* `max_entries` - (Optional, Computed, Int) The maximum number of entries for CIDR address blocks in the prefix list.
* `prefix_list_description` - (Optional) The description of the prefix list.
It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
* `prefix_list_name` - (Optional) The name of the prefix list. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, periods (.), underscores (_), and hyphens (-).
* `resource_group_id` - (Optional, Computed, Available since v1.205.0) The ID of the resource group to which the PrefixList belongs.
* `tags` - (Optional, Map, Available since v1.205.0) The tags of PrefixList.

The following arguments will be discarded. Please use new fields as soon as possible:
* `entrys` - (Deprecated since v1.274.0). Field 'entrys' has been deprecated from provider version 1.274.0. New field 'entries' instead.

### `entries`

The entries supports the following:
* `cidr` - (Optional, Available since v1.274.0) The CIDR address block of the prefix list.
* `description` - (Optional, Available since v1.274.0) The description of the cidr entry. It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The time when the prefix list was created.
* `prefix_list_association` - The association list information of the prefix list.
  * `owner_id` - The ID of the Alibaba Cloud account (primary account) to which the prefix list belongs.
  * `prefix_list_id` - The instance ID of the prefix list.
  * `reason` - Reason when the association fails.
  * `region_id` - The region ID of the prefix list to be queried.
  * `resource_id` - The ID of the associated resource.
  * `resource_type` - The associated resource type.
  * `resource_uid` - The ID of the Alibaba Cloud account (primary account) to which the resource bound to the prefix list belongs.
  * `status` - The association status of the prefix list.
* `prefix_list_id` - The ID of the query Prefix List.
* `region_id` - The region ID of the resource.
* `share_type` - The share type of the prefix list.
* `status` - Resource attribute fields that represent the status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Prefix List.
* `delete` - (Defaults to 5 mins) Used when delete the Prefix List.
* `update` - (Defaults to 5 mins) Used when update the Prefix List.

## Import

VPC Prefix List can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_prefix_list.example <prefix_list_id>
```