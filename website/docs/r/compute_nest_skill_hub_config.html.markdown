---
subcategory: "Compute Nest"
layout: "alicloud"
page_title: "Alicloud: alicloud_compute_nest_skill_hub_config"
description: |-
  Provides a Alicloud Compute Nest Skill Hub Config resource.
---

# alicloud_compute_nest_skill_hub_config

Provides a Compute Nest Skill Hub Config resource.

The Skill Hub configuration information, which is primarily used to store the storage location of user-uploaded skills.

For information about Compute Nest Skill Hub Config and how to use it, see [What is Skill Hub Config](https://next.api.alibabacloud.com/document/ComputeNest/2021-06-01/CreateSkillHubConfig).

-> **NOTE:** Available since v1.285.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_compute_nest_skill_hub_config` or removing it from your configuration

Terraform cannot destroy resource `alicloud_compute_nest_skill_hub_config`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `oss_bucket_name` - (Required, ForceNew) The name of the Object Storage Service (OSS) bucket.
* `oss_region_id` - (Required, ForceNew) The region ID of the OSS bucket. For more information about how to obtain the region ID, see [Regions and endpoints](https://help.aliyun.com/document_detail/31837.html).

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<Alibaba Cloud Account ID>`.
* `create_time` - The creation time.
* `update_time` - The update time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Skill Hub Config.

## Import

Compute Nest Skill Hub Config can be imported using the id, e.g.

```shell
$ terraform import alicloud_compute_nest_skill_hub_config.example <Alibaba Cloud Account ID>
```