---
subcategory: "Vpc"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc"
sidebar_current: "docs-alicloud-resource-vpc"
description: |-
  Provides a Alicloud Vpc Vpc resource.
---

# alicloud_vpc

Provides a Vpc Vpc resource.

For information about Vpc Vpc and how to use it, see [What is Vpc](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc" "default" {
  description = "test"
  region_id   = var.RegionId
}
```

## Argument Reference

The following arguments are supported:
* `cidr_block` - (Computed,Optional) The IPv4 CIDR block of the VPC.
* `classic_link_enabled` - (Optional) The status of ClassicLink function.
* `description` - (Optional) The description of the VPC.
* `dry_run` - (ForceNew,Optional) Whether to PreCheck this request only. Value:-**true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
* `enable_ipv6` - (Optional) Whether to enable the IPv6 network segment. Value:-**false** (default): not enabled.-**true**: on.
* `ipv6_cidr_block` - (Computed,Optional) The ipv6 cidr block of vpc.
* `ipv6_isp` - (Optional) The IPv6 address segment type of the VPC. Value:-**BGP** (default): Alibaba Cloud BGP IPv6.-**ChinaMobile**: China Mobile (single line).-**ChinaUnicom**: China Unicom (single line).-**ChinaTelecom**: China Telecom (single line).> If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the VPC belongs.
* `secondary_cidr_blocks` - (Optional) Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.
* `tags` - (Optional) The tags of Vpc.See the following `Block Tags`.
* `user_cidrs` - (ForceNew,Computed,Optional) A list of user CIDRs.
* `vpc_name` - (Optional) The name of the VPC.

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
* `router_table_id` - Field 'router_table_id' has been deprecated from provider version 1.204.0. New field 'route_table_id' instead.

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of tag.
* `tag_value` - (Optional) The value of tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `cidr_block` - The IPv4 CIDR block of the VPC.
* `create_time` - The time at which the VPC was created.
* `dhcp_options_set_id` - The ID of the DHCP option set.
* `ipv6_cidr_block` - The ipv6 cidr block of vpc.
* `ipv6_cidr_blocks` - The IPv6 network segment of the VPC.
* `is_default` - Indicates whether to query the default VPC in the specified region. Valid values:  true (default): All VPCs in the specified region are queried. false: The default VPC is not queried.
* `resource_group_id` - The ID of the resource group to which the VPC belongs.
* `route_table_id` - The ID of the route table to query.
* `router_id` - The ID of the VRouter.
* `status` - The status of the VPC. Valid values:  Pending: The VPC is being configured. Available: The VPC is available.
* `user_cidrs` - A list of user CIDRs.
* `vswitch_ids` - A list of VSwitches in the VPC.
* `vpc_id` - The ID of the VPC.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Vpc.
* `delete` - (Defaults to 5 mins) Used when delete the Vpc.
* `update` - (Defaults to 5 mins) Used when update the Vpc.

## Import

Vpc Vpc can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_vpc.example 
```