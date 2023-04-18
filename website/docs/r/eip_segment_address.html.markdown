---
subcategory: "EIP"
layout: "alicloud"
page_title: "Alicloud: alicloud_eip_segment_address"
sidebar_current: "docs-alicloud-resource-eip-segment-address"
description: |-
  Provides a Alicloud Eip Segment Address resource.
---

# alicloud_eip_segment_address

Provides a Eip Segment Address resource.

For information about Eip Segment Address and how to use it, see [What is Segment Address](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_eip_segment_address" "default" {
  resource_group_id = "rg-acfmv7pftjmydwa"
  isp               = "BGP"
  region_id         = "cn-shanghai"
  eip_mask          = "28"
}
```

## Argument Reference

The following arguments are supported:
* `bandwidth` - (ForceNew,Optional) 当前属性没有在镇元上录入属性描述，请补充后再生成代码。
* `eip_mask` - (Required,ForceNew) Mask of consecutive EIPs. Value:28: For a single call, the system will allocate 16 consecutive EIPs.27: For a single call, the system will allocate 32 consecutive EIPs.26: For a single call, the system will allocate 64 consecutive EIPs.25: For a single call, the system will allocate 128 consecutive EIPs.24: For a single call, the system will allocate 256 consecutive EIPs.
* `internet_charge_type` - (ForceNew,Optional) Continuous EIP billing method, value:PayByBandwidth (default): Billing based on fixed bandwidth.PayByTraffic: Billing by usage flow.
* `isp` - (ForceNew,Optional) Service providers.
* `netmode` - (ForceNew,Optional) 当前属性没有在镇元上录入属性描述，请补充后再生成代码。
* `resource_group_id` - (ForceNew,Computed,Optional) The ID of the resource group.

The following arguments will be discarded. Please use new fields as soon as possible:



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `create_time` - The time when the contiguous Elastic IP address group was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
* `descritpion` - The description of the contiguous Elastic IP address group.
* `ip_count` - The number of IP addresses contained in the contiguous Elastic IP address group.
* `resource_group_id` - The ID of the resource group.
* `segment` - The CIDR block and mask of the contiguous Elastic IP address group.
* `segment_address_name` - The name of the contiguous Elastic IP address group.
* `segment_instance_id` - Instance ID of consecutive EIP groups
* `status` - The status of the resource
* `zone` - The zone of the EIP.This parameter is returned only for whitelist users that are visible to the zone.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Segment Address.
* `delete` - (Defaults to 5 mins) Used when delete the Segment Address.

## Import

Eip Segment Address can be imported using the id, e.g.

```shell
$ terraform import alicloud_eip_segment_address.example 
```