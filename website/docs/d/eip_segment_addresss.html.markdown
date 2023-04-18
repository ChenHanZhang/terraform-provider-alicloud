---
subcategory: "EIP"
layout: "alicloud"
page_title: "Alicloud: alicloud_eip_segment_addresss"
sidebar_current: "docs-alicloud-datasource-eip-segment-addresss"
description: |-
  Provides a list of Eip Segment Address owned by an Alibaba Cloud account.
---

# alicloud_eip_segment_addresss

This data source provides Eip Segment Address available to the user.[What is Segment Address](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_eip_segment_addresss" "default" {
  ids        = ["${alicloud_eip_segment_address.default.id}"]
  name_regex = alicloud_eip_segment_address.default.name
}

output "alicloud_eip_segment_address_example_id" {
  value = data.alicloud_eip_segment_addresss.default.addresss.0.id
}
```

## Argument Reference

The following arguments are supported:
* `segment_instance_id` - (ForceNew,Optional) Instance ID of consecutive EIP groups
* `ids` - (Optional, ForceNew, Computed) A list of Segment Address IDs.
* `segment_address_names` - (Optional, ForceNew) The name of the Segment Address. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Segment Address IDs.
* `names` - A list of name of Segment Addresss.
* `addresss` - A list of Segment Address Entries. Each element contains the following attributes:
  * `create_time` - The time when the contiguous Elastic IP address group was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
  * `descritpion` - The description of the contiguous Elastic IP address group.
  * `ip_count` - The number of IP addresses contained in the contiguous Elastic IP address group.
  * `segment` - The CIDR block and mask of the contiguous Elastic IP address group.
  * `segment_address_name` - The name of the contiguous Elastic IP address group.
  * `segment_instance_id` - Instance ID of consecutive EIP groups
  * `status` - The status of the resource
  * `zone` - The zone of the EIP.This parameter is returned only for whitelist users that are visible to the zone.
