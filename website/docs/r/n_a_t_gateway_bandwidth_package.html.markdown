---
subcategory: "NAT Gateway"
layout: "alicloud"
page_title: "Alicloud: alicloud_n_a_t_gateway_bandwidth_package"
description: |-
  Provides a Alicloud N A T Gateway Bandwidth Package resource.
---

# alicloud_n_a_t_gateway_bandwidth_package

Provides a N A T Gateway Bandwidth Package resource.



For information about N A T Gateway Bandwidth Package and how to use it, see [What is Bandwidth Package](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/DescribeBandwidthPackages).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_n_a_t_gateway_bandwidth_package` or removing it from your configuration

Terraform cannot destroy resource `alicloud_n_a_t_gateway_bandwidth_package`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `bandwidth_package_id` - (Required, ForceNew) The ID of the Internet shared bandwidth.
* `payment_type` - (Optional, ForceNew, Computed) PaymentType
* `resource_group_id` - (Optional, ForceNew, Computed) The ID of the resource group.
* `status` - (Optional, ForceNew, Computed) The status of the Internet Shared Bandwidth instance. Default value: **Available * *.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The ID of the region where the internet shared bandwidth instance resides.


## Import

N A T Gateway Bandwidth Package can be imported using the id, e.g.

```shell
$ terraform import alicloud_n_a_t_gateway_bandwidth_package.example <bandwidth_package_id>
```