---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv6_egress_rules"
sidebar_current: "docs-alicloud-datasource-vpc-ipv6-egress-rules"
description: |-
  Provides a list of Vpc Ipv6 Egress Rule owned by an Alibaba Cloud account.
---

# alicloud_vpc_ipv6_egress_rules

This data source provides Vpc Ipv6 Egress Rule available to the user.[What is Ipv6 Egress Rule](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_ipv6_egress_rules" "default" {
  name_regex            = alicloud_vpc_ipv6_egress_rule.default.name
  instance_id           = "ipv6-bp11fycipeipu0ae95nz1"
  instance_type         = "Ipv6Address"
  ipv6_egress_rule_name = "rdk-test"
  ipv6_gateway_id       = "ipv6gw-bp1kc0t2ndkccmnbnwjsu"
}

output "alicloud_vpc_ipv6_egress_rule_example_id" {
  value = data.alicloud_vpc_ipv6_egress_rules.default.rules.0.id
}
```

## Argument Reference

The following arguments are supported:
* `instance_id` - (ForceNew,Optional) InstanceId
* `instance_type` - (ForceNew,Optional) InstanceType
* `ipv6_egress_rule_id` - (ForceNew,Optional) The first ID of the resource
* `ipv6_egress_rule_name` - (ForceNew,Optional) The name of the resource
* `ipv6_gateway_id` - (Required,ForceNew) Ipv6GatewayId
* `ipv6_egress_rule_names` - (Optional, ForceNew) The name of the Ipv6 Egress Rule. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `names` - A list of name of Ipv6 Egress Rules.
* `rules` - A list of Ipv6 Egress Rule Entries. Each element contains the following attributes:
  * `description` - Description
  * `instance_id` - InstanceId
  * `instance_type` - InstanceType
  * `ipv6_egress_rule_id` - The first ID of the resource
  * `ipv6_egress_rule_name` - The name of the resource
  * `status` - The status of the resource
