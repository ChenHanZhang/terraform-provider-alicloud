---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_havips"
sidebar_current: "docs-alicloud-datasource-havips"
description: |-
  Provides a list of Vpc Havip owned by an Alibaba Cloud account.
---

# alicloud_havips

This data source provides Vpc Havip available to the user.[What is Havip](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.203.0+.

## Example Usage

```
data "alicloud_vpc_havips" "default" {
  ids = ["${alicloud_havip.default.id}"]
}

output "alicloud_havip_example_id" {
  value = data.alicloud_vpc_havips.default.havips.0.id
}
```

## Argument Reference

The following arguments are supported:
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group to which the VPC belongs.
* `tags` - (ForceNew,Optional) The tags of PrefixList.See the following `Block Tags`.
* `ids` - (Optional, ForceNew, Computed) A list of Havip IDs.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

#### Block Tags

The Tags supports the following:
* `tag_key` - (ForceNew,Optional) The key of tag.
* `tag_value` - (ForceNew,Optional) The value of tag.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Havip IDs.
* `havips` - A list of Havip Entries. Each element contains the following attributes:
  * `associated_eip_addresses` - EIP bound to HaVip
  * `associated_instance_type` - The type of the instance that is bound to the VIIP. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
  * `associated_instances` - An ECS instance that is bound to HaVip
  * `create_time` - The creation time of the  resource
  * `description` - Dependence of a HaVip instance.
  * `ha_vip_id` - The  ID of the resource
  * `ha_vip_name` - The name of the HaVip instance
  * `ip_address` - IP address of private network
  * `master_instance_id` - The primary instance ID bound to HaVip
  * `resource_group_id` - The ID of the resource group to which the VPC belongs.
  * `status` - The status
  * `tags` - The tags of PrefixList.
    * `tag_key` - The key of tag.
    * `tag_value` - The value of tag.
  * `vswitch_id` - The switch ID to which the HaVip instance belongs
  * `vpc_id` - The VPC ID to which the HaVip instance belongs
