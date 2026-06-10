---
subcategory: "Compute Nest"
layout: "alicloud"
page_title: "Alicloud: alicloud_compute_nest_skill_space"
description: |-
  Provides a Alicloud Compute Nest Skill Space resource.
---

# alicloud_compute_nest_skill_space

Provides a Compute Nest Skill Space resource.

A logical set of skills.

For information about Compute Nest Skill Space and how to use it, see [What is Skill Space](https://next.api.alibabacloud.com/document/ComputeNest/2021-06-01/CreateSkillSpace).

-> **NOTE:** Available since v1.282.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `skill_space_description` - (Required, ForceNew) Description
* `skill_space_name` - (Required, ForceNew) Name

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation Time.
* `update_time` - Modification Time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Skill Space.
* `delete` - (Defaults to 5 mins) Used when delete the Skill Space.

## Import

Compute Nest Skill Space can be imported using the id, e.g.

```shell
$ terraform import alicloud_compute_nest_skill_space.example <skill_space_id>
```