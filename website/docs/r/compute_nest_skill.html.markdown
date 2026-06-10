---
subcategory: "Compute Nest"
layout: "alicloud"
page_title: "Alicloud: alicloud_compute_nest_skill"
description: |-
  Provides a Alicloud Compute Nest Skill resource.
---

# alicloud_compute_nest_skill

Provides a Compute Nest Skill resource.

The term in the field of AI Agent is used to teach AI to do things according to a fixed process ".

For information about Compute Nest Skill and how to use it, see [What is Skill](https://next.api.alibabacloud.com/document/ComputeNest/2021-06-01/CreateSkill).

-> **NOTE:** Available since v1.282.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `oss_url` - (Optional) The SKILL File link of the OSS Bucket, which is passed when SourceType = UPLOAD

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `skill_description` - (Optional) The creation time of the resource
* `skill_labels` - (Optional, List) Skill Labels
* `skill_name` - (Optional, ForceNew) The name of the resource
* `skill_space_id` - (Required, ForceNew) Belongs to SkillSpaceId
* `source_skill_id` - (Optional) Official SKILL ID, passed when SourceType = COPY

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation Time.
* `update_time` - Modification Time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Skill.
* `delete` - (Defaults to 5 mins) Used when delete the Skill.
* `update` - (Defaults to 5 mins) Used when update the Skill.

## Import

Compute Nest Skill can be imported using the id, e.g.

```shell
$ terraform import alicloud_compute_nest_skill.example <skill_id>
```