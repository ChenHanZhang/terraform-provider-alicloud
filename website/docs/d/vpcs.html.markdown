---
subcategory: "Vpc"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpcs"
sidebar_current: "docs-alicloud-datasource-vpcs"
description: |-
  Provides a list of Vpc Vpc owned by an Alibaba Cloud account.
---

# alicloud_vpcs

This data source provides Vpc Vpc available to the user.[What is Vpc](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_vpcs" "default" {
  ids        = ["${alicloud_vpc.default.id}"]
  name_regex = alicloud_vpc.default.name
}

output "alicloud_vpc_example_id" {
  value = data.alicloud_vpc_vpcs.default.vpcs.0.id
}
```

## Argument Reference

The following arguments are supported:
* `dhcp_options_set_id` - (ForceNew,Optional) The ID of the DHCP option set.
* `dry_run` - (ForceNew,Optional) Whether to PreCheck this request only. Value:-**true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
* `is_default` - (ForceNew,Optional) Indicates whether to query the default VPC in the specified region. Valid values:  true (default): All VPCs in the specified region are queried. false: The default VPC is not queried.
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the VPC belongs.
* `vpc_id` - (ForceNew,Optional) The ID of the VPC.
* `vpc_name` - (ForceNew,Optional) The name of the VPC.
* `ids` - (Optional, ForceNew, Computed) A list of Vpc IDs.
* `vpc_names` - (Optional, ForceNew) The name of the Vpc. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Vpc IDs.
* `names` - A list of name of Vpcs.
* `vpcs` - A list of Vpc Entries. Each element contains the following attributes:
  * `cidr_block` - The IPv4 CIDR block of the VPC.
  * `create_time` - The time at which the VPC was created.
  * `description` - The description of the VPC.
  * `dhcp_options_set_id` - The ID of the DHCP option set.
  * `ipv6_cidr_block` - The ipv6 cidr block of vpc.
  * `ipv6_cidr_blocks` - The IPv6 network segment of the VPC.
    * `ipv6_cidr_block` - The IPv6 network segment of the VPC.
  * `is_default` - Indicates whether to query the default VPC in the specified region. Valid values:  true (default): All VPCs in the specified region are queried. false: The default VPC is not queried.
  * `resource_group_id` - The ID of the resource group to which the VPC belongs.
  * `router_id` - The ID of the VRouter.
  * `secondary_cidr_blocks` - Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
  * `status` - The status of the VPC. Valid values:  Pending: The VPC is being configured. Available: The VPC is available.
  * `tags` - The tags of Vpc.
    * `tag_key` - The key of tag.
    * `tag_value` - The value of tag.
  * `user_cidrs` - A list of user CIDRs.
  * `vswitch_ids` - A list of VSwitches in the VPC.
  * `vpc_id` - The ID of the VPC.
  * `vpc_name` - The name of the VPC.
