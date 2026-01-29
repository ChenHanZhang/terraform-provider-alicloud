---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_dedicated_host"
description: |-
  Provides a Alicloud RDS Dedicated Host resource.
---

# alicloud_rds_dedicated_host

Provides a RDS Dedicated Host resource.

Exclusive host group host.

For information about RDS Dedicated Host and how to use it, see [What is Dedicated Host](https://next.api.alibabacloud.com/document/Rds/2014-08-15/DescribeDedicatedHosts).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_rds_dedicated_host` or removing it from your configuration

Terraform cannot destroy resource `alicloud_rds_dedicated_host`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `dedicated_host_group_id` - (Optional, ForceNew) The ID of the host group to which the host belongs.
* `dedicated_host_id` - (Optional, ForceNew, Computed) The host ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 


## Import

RDS Dedicated Host can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_dedicated_host.example <dedicated_host_id>
```