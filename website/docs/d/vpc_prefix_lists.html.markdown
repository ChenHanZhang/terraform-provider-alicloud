---
subcategory: "Vpc"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_prefix_lists"
sidebar_current: "docs-alicloud-datasource-vpc-prefix-lists"
description: |-
  Provides a list of Vpc Prefix List owned by an Alibaba Cloud account.
---

# alicloud_vpc_prefix_lists

This data source provides Vpc Prefix List available to the user.[What is Prefix List](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_prefix_lists" "default" {
  ids              = ["${alicloud_vpc_prefix_list.default.id}"]
  name_regex       = alicloud_vpc_prefix_list.default.name
  prefix_list_name = "rdk-test2"
  tags {
    tag_key   = "k0"
    tag_value = "v0"
  }
}

output "alicloud_vpc_prefix_list_example_id" {
  value = data.alicloud_vpc_prefix_lists.default.lists.0.id
}
```

## Argument Reference

The following arguments are supported:
* `prefix_list_name` - (ForceNew,Optional) The name of the prefix list.
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the VPC belongs.
* `tags` - (ForceNew,Optional) The tags of PrefixList.See the following `Block Tags`.
* `ids` - (Optional, ForceNew, Computed) A list of Prefix List IDs.
* `prefix_list_names` - (Optional, ForceNew) The name of the Prefix List. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

#### Block Tags

The Tags supports the following:
* `tag_key` - (ForceNew,Optional) The key of tag.
* `tag_value` - (ForceNew,Optional) The value of tag.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Prefix List IDs.
* `names` - A list of name of Prefix Lists.
* `lists` - A list of Prefix List Entries. Each element contains the following attributes:
  * `create_time` - The time when the prefix list was created.
  * `ip_version` - The IP version of the prefix list. Value:-**ipv4**:IPv4 version.-**ipv6**:IPv6.
  * `max_entries` - The maximum number of entries for CIDR address blocks in the prefix list.
  * `owner_id` - The Alibaba Cloud account (primary account) to which the prefix list belongs.
  * `prefix_list_description` - The description of the prefix list.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with http:// or https.
  * `prefix_list_id` - The ID of the query Prefix List.
  * `prefix_list_name` - The name of the prefix list.
  * `resource_group_id` - The ID of the resource group to which the VPC belongs.
  * `share_type` - The share type of the prefix list. Value:-**Shared**: indicates that the prefix list is a Shared prefix list.-Null: indicates that the prefix list is not a shared prefix list.
  * `status` - The status of the resource
