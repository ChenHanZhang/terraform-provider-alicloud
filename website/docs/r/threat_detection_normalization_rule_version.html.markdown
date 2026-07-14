---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_normalization_rule_version"
description: |-
  Provides a Alicloud Threat Detection Normalization Rule Version resource.
---

# alicloud_threat_detection_normalization_rule_version

Provides a Threat Detection Normalization Rule Version resource.

Normalization rule version.

For information about Threat Detection Normalization Rule Version and how to use it, see [What is Normalization Rule Version](https://next.api.alibabacloud.com/document/cloud-siem/2024-12-12/GetNormalizationRuleVersion).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `lang` - (Optional) The language type of the API request.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `normalization_rule_id` - (Optional, ForceNew, Computed) The ID of the standardization rule.
* `normalization_rule_version` - (Optional, Computed, Int) The current version of the normalization rule.
* `role_for` - (Optional, Int) The administrator switches the ID of another user.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<normalization_rule_id>:<normalization_rule_version>`.
* `region_id` - The region ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `delete` - (Defaults to 5 mins) Used when delete the Normalization Rule Version.

## Import

Threat Detection Normalization Rule Version can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_normalization_rule_version.example <normalization_rule_id>:<normalization_rule_version>
```