---
subcategory: "ACC"
layout: "alicloud"
page_title: "Alicloud: alicloud_acc_instance_resource_configuration"
description: |-
  Provides a Alicloud ACC Instance Resource Configuration resource.
---

# alicloud_acc_instance_resource_configuration

Provides a ACC Instance Resource Configuration resource.

Configuration used for instance provisioning in the standby pool.

For information about ACC Instance Resource Configuration and how to use it, see [What is Instance Resource Configuration](https://next.api.alibabacloud.com/document/acc/2024-04-02/CreateInstanceResourceConfiguration).

-> **NOTE:** Available since v1.272.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `compute_class` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `compute_qo_s` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `cpu` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `cpu_vendor` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `extra_ephemeral_storage` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `memory` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `name` - (Required, ForceNew) The name of the resource
* `security_group_ids` - (Required, ForceNew, List) This property does not have a description in the spec, please add it before generating code.
* `vswitch_ids` - (Required, ForceNew, List) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Instance Resource Configuration.
* `delete` - (Defaults to 5 mins) Used when delete the Instance Resource Configuration.

## Import

ACC Instance Resource Configuration can be imported using the id, e.g.

```shell
$ terraform import alicloud_acc_instance_resource_configuration.example <instance_resource_configuration_id>
```