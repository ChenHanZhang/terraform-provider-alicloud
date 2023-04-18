---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv6_internet_bandwidth"
sidebar_current: "docs-alicloud-resource-vpc-ipv6-internet-bandwidth"
description: |-
  Provides a Alicloud Vpc Ipv6 Internet Bandwidth resource.
---

# alicloud_vpc_ipv6_internet_bandwidth

Provides a Vpc Ipv6 Internet Bandwidth resource.

For information about Vpc Ipv6 Internet Bandwidth and how to use it, see [What is Ipv6 Internet Bandwidth](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_vpc_ipv6_internet_bandwidth" "default" {
  bandwidth            = 10
  ipv6_address_id      = "ipv6-bp11fycipeipu0ae95nz1"
  ipv6_gateway_id      = "ipv6gw-bp1kc0t2ndkccmnbnwjsu"
  region_id            = "cn-hangzhou"
  internet_charge_type = "PayByBandwidth"
}
```

## Argument Reference

The following arguments are supported:
* `bandwidth` - (Required) Bandwidth
* `internet_charge_type` - (ForceNew,Computed,Optional) The billing method of the public network bandwidth of the IPv6 address.
* `ipv6_address_id` - (Required) IPv6地址实例的ID The ID of the IPv6 address
* `ipv6_gateway_id` - (Required,ForceNew) Ipv6GatewayId

The following arguments will be discarded. Please use new fields as soon as possible:



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `internet_charge_type` - The billing method of the public network bandwidth of the IPv6 address.
* `ipv6_internet_bandwidth_id` - Ipv6InternetBandwidthId
* `payment_type` - The payment type of the resource
* `status` - The status of the resource

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ipv6 Internet Bandwidth.
* `delete` - (Defaults to 5 mins) Used when delete the Ipv6 Internet Bandwidth.
* `update` - (Defaults to 5 mins) Used when update the Ipv6 Internet Bandwidth.

## Import

Vpc Ipv6 Internet Bandwidth can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_ipv6_internet_bandwidth.example 
```