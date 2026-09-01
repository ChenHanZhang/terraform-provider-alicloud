---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_security_group_attachment"
description: |-
  Provides a Alicloud ECS Security Group Attachment resource.
---

# alicloud_ecs_security_group_attachment

Provides a ECS Security Group Attachment resource.

Security Group mount.

For information about ECS Security Group Attachment and how to use it, see [What is Security Group Attachment](https://next.api.alibabacloud.com/document/Ecs/2014-05-26/JoinSecurityGroup).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `instance_id` - (Optional, ForceNew, Computed) The ID of the instance.

-> **NOTE:**  When this parameter passes in a value, 'NetworkInterfaceId' must be empty.

* `network_interface_id` - (Optional, ForceNew) The ID of the Eni.

-> **NOTE:**  When the value of this parameter is passed in, 'InstanceId' must be empty.

* `security_group_id` - (Required, ForceNew) The ID of the security group. You can call [DescribeSecurityGroups](~~ 25556 ~~) to view the available security groups.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<security_group_id>`.
* `region_id` - The region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Security Group Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Security Group Attachment.

## Import

ECS Security Group Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_security_group_attachment.example <instance_id>:<security_group_id>
```