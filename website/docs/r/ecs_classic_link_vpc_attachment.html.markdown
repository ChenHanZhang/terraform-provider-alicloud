---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_classic_link_vpc_attachment"
description: |-
  Provides a Alicloud ECS Classic Link Vpc Attachment resource.
---

# alicloud_ecs_classic_link_vpc_attachment

Provides a ECS Classic Link Vpc Attachment resource.

Classic network mount.

For information about ECS Classic Link Vpc Attachment and how to use it, see [What is Classic Link Vpc Attachment](https://next.api.alibabacloud.com/document/Ecs/2014-05-26/AttachClassicLinkVpc).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `instance_id` - (Required, ForceNew) The ID of the instance.
* `vpc_id` - (Required, ForceNew) VPC ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<vpc_id>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Classic Link Vpc Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Classic Link Vpc Attachment.

## Import

ECS Classic Link Vpc Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_classic_link_vpc_attachment.example <instance_id>:<vpc_id>
```