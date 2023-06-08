---
subcategory: "EIP"
layout: "alicloud"
page_title: "Alicloud: alicloud_eip_segment_address"
description: |-
  Provides a Alicloud EIP Segment Address resource.
---

# alicloud_eip_segment_address

Provides a EIP Segment Address resource. 

For information about EIP Segment Address and how to use it, see [What is Segment Address](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.207.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_resource_manager_resource_group" "defaultyHWGFu" {
  display_name        = "test03"
  resource_group_name = var.name
}


resource "alicloud_eip_segment_address" "default" {
  eip_mask             = "28"
  bandwidth            = "5"
  isp                  = "BGP"
  internet_charge_type = "PayByBandwidth"
  resource_group_id    = alicloud_resource_manager_resource_group.defaultyHWGFu.id
}
```


## Argument Reference

The following arguments are supported:
* `bandwidth` - (Optional, ForceNew) The peak bandwidth of the EIP. Unit: Mbps. When the value of instancargetype is PostPaid and the value of InternetChargeType is PayByBandwidth, the range of Bandwidth is 1 to 500. If the value of instancargetype is PostPaid and the value of InternetChargeType is PayByTraffic, the range of Bandwidth is 1 to 200. When instancargetype is set to PrePaid, the range of Bandwidth is 1 to 1000. The default value is 5 Mbps.
* `eip_mask` - (Required, ForceNew) Mask of consecutive EIPs. Value:28: For a single call, the system will allocate 16 consecutive EIPs.27: For a single call, the system will allocate 32 consecutive EIPs.26: For a single call, the system will allocate 64 consecutive EIPs.25: For a single call, the system will allocate 128 consecutive EIPs.24: For a single call, the system will allocate 256 consecutive EIPs.
* `internet_charge_type` - (Optional, ForceNew) Continuous EIP billing method, value:PayByBandwidth (default): Billing based on fixed bandwidth.PayByTraffic: Billing by usage flow.
* `isp` - (Optional, ForceNew) Service providers.
* `netmode` - (Optional, ForceNew) 当前属性没有在镇元上录入属性描述，请补充后再生成代码。.
* `resource_group_id` - (Optional, ForceNew, Computed) The ID of the resource group.
* `zone` - (Optional, ForceNew) The zone of the EIP.This parameter is returned only for whitelist users that are visible to the zone.



## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The time when the contiguous Elastic IP address group was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
* `status` - The status of the resource.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Segment Address.
* `delete` - (Defaults to 5 mins) Used when delete the Segment Address.

## Import

EIP Segment Address can be imported using the id, e.g.

```shell
$ terraform import alicloud_eip_segment_address.example <id>
```