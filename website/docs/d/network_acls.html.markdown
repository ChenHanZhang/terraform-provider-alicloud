---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_network_acls"
sidebar_current: "docs-alicloud-datasource-network-acls"
description: |-
  Provides a list of Vpc Network Acl owned by an Alibaba Cloud account.
---

# alicloud_network_acls

This data source provides Vpc Network Acl available to the user.[What is Network Acl](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_network_acls" "default" {
  ids              = ["${alicloud_network_acl.default.id}"]
  name_regex       = alicloud_network_acl.default.name
  network_acl_name = "rdk-test"
  vpc_id           = "vpc-bp11lfjeaa57jxr6ovybf"
}

output "alicloud_network_acl_example_id" {
  value = data.alicloud_vpc_network_acls.default.acls.0.id
}
```

## Argument Reference

The following arguments are supported:
* `network_acl_id` - (ForceNew,Optional) The first ID of the resource
* `network_acl_name` - (ForceNew,Optional) The name of the network ACL.
* `vpc_id` - (ForceNew,Optional) The ID of the associated VPC.
* `ids` - (Optional, ForceNew, Computed) A list of Network Acl IDs.
* `network_acl_names` - (Optional, ForceNew) The name of the Network Acl. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Network Acl IDs.
* `names` - A list of name of Network Acls.
* `acls` - A list of Network Acl Entries. Each element contains the following attributes:
  * `create_time` - The creation time of the resource
  * `description` - Description of network ACL information.
  * `egress_acl_entries` - Output direction rule information.
    * `description` - Give the description information of the direction rule.
    * `destination_cidr_ip` - The destination address segment.
    * `network_acl_entry_name` - The name of the entry for the direction rule.
    * `policy` - The  authorization policy.
    * `port` - Destination port range.
    * `protocol` - Transport  layer protocol.
  * `ingress_acl_entries` - Entry direction rule information.
    * `description` - Description of the entry direction rule.
    * `network_acl_entry_name` - The name of the entry direction rule entry.
    * `policy` - The authorization policy.
    * `port` - Source port range.
    * `protocol` - Transport layer protocol.
    * `source_cidr_ip` - The source address field.
  * `network_acl_id` - The first ID of the resource
  * `network_acl_name` - The name of the network ACL.
  * `resources` - The associated resource.
    * `resource_id` - The ID of the associated resource.
    * `resource_type` - The type of the associated resource.
    * `status` - The state of the associated resource.
  * `status` - The state of the network ACL.
  * `tags` - The tags of VSwitch.
    * `tag_key` - The tag key of VSwitch.
    * `tag_value` - The tag value of VSwitch.
  * `vpc_id` - The ID of the associated VPC.
