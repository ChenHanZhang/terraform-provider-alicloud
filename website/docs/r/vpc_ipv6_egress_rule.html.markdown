---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv6_egress_rule"
sidebar_current: "docs-alicloud-resource-vpc-ipv6-egress-rule"
description: |-
  Provides a Alicloud Vpc Ipv6 Egress Rule resource.
---

# alicloud_vpc_ipv6_egress_rule

Provides a Vpc Ipv6 Egress Rule resource.

For information about Vpc Ipv6 Egress Rule and how to use it, see [What is Ipv6 Egress Rule](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_ipv6_egress_rule" "default" {
  description           = "test"
  instance_id           = "ipv6-bp11fycipeipu0ae95nz1"
  ipv6_gateway_id       = "ipv6gw-bp1kc0t2ndkccmnbnwjsu"
  ipv6_egress_rule_name = "rdk-test"
  region_id             = "cn-hangzhou"
  instance_type         = "Ipv6Address"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (ForceNew,Optional) Description
* `instance_id` - (Required,ForceNew) InstanceId
* `instance_type` - (ForceNew,Computed,Optional) InstanceType
* `ipv6_egress_rule_name` - (ForceNew,Optional) The name of the resource
* `ipv6_gateway_id` - (Required,ForceNew) Ipv6GatewayId

The following arguments will be discarded. Please use new fields as soon as possible:



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.The value is formulated as `<ipv6_gateway_id>:<ipv6_egress_rule_id>`.
* `instance_type` - InstanceType
* `ipv6_egress_rule_id` - The first ID of the resource
* `status` - The status of the resource

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ipv6 Egress Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Ipv6 Egress Rule.

## Import

Vpc Ipv6 Egress Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_ipv6_egress_rule.example <ipv6_gateway_id>:<ipv6_egress_rule_id>
```