---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv6_internet_bandwidths"
sidebar_current: "docs-alicloud-datasource-vpc-ipv6-internet-bandwidths"
description: |-
  Provides a list of Vpc Ipv6 Internet Bandwidth owned by an Alibaba Cloud account.
---

# alicloud_vpc_ipv6_internet_bandwidths

This data source provides Vpc Ipv6 Internet Bandwidth available to the user.[What is Ipv6 Internet Bandwidth](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_vpc_ipv6_internet_bandwidths" "default" {
  ids = ["${alicloud_vpc_ipv6_internet_bandwidth.default.id}"]
}

output "alicloud_vpc_ipv6_internet_bandwidth_example_id" {
  value = data.alicloud_vpc_ipv6_internet_bandwidths.default.bandwidths.0.id
}
```

## Argument Reference

The following arguments are supported:
* `ipv6_internet_bandwidth_id` - (ForceNew,Optional) Ipv6InternetBandwidthId
* `ids` - (Optional, ForceNew, Computed) A list of Ipv6 Internet Bandwidth IDs.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Ipv6 Internet Bandwidth IDs.
* `bandwidths` - A list of Ipv6 Internet Bandwidth Entries. Each element contains the following attributes:
  * `bandwidth` - Bandwidth
  * `internet_charge_type` - The billing method of the public network bandwidth of the IPv6 address.
  * `ipv6_address_id` - IPv6地址实例的ID The ID of the IPv6 address
  * `ipv6_gateway_id` - Ipv6GatewayId
  * `ipv6_internet_bandwidth_id` - Ipv6InternetBandwidthId
  * `payment_type` - The payment type of the resource
  * `status` - The status of the resource
