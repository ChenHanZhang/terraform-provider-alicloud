---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_normalization_rule"
description: |-
  Provides a Alicloud Threat Detection Normalization Rule resource.
---

# alicloud_threat_detection_normalization_rule

Provides a Threat Detection Normalization Rule resource.



For information about Threat Detection Normalization Rule and how to use it, see [What is Normalization Rule](https://next.api.alibabacloud.com/document/cloud-siem/2024-12-12/CreateNormalizationRule).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `extend_field_store_mode` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `lang` - (Optional) The language type.

-> **NOTE:** This parameter only applies during resource creation, update or deletion. If modified in isolation without other property changes, Terraform will not trigger any action.

* `normalization_category_id` - (Optional) The ID of the standardization rule classification.
* `normalization_rule_description` - (Optional) Standardized rule description.
* `normalization_rule_expression` - (Optional) Standardize rule expressions.
* `normalization_rule_format` - (Optional) Standardized rule format.
* `normalization_rule_ids` - (Optional, ForceNew, List) NormalizationRuleIds
* `normalization_rule_mode` - (Optional) Standardized rule patterns. Value:
  - both.
  - scan.
  - realtime.
* `normalization_rule_name` - (Optional) Standardized rule name.
* `normalization_rule_type` - (Optional) Normalization Rule Types
* `normalization_rule_version` - (Optional, Int) Standardize rule versions.
* `normalization_schema_id` - (Optional) The standardized structure ID.
* `order_field` - (Optional) Order Field
* `product_id` - (Optional) The product ID.
* `role_for` - (Optional, Int) The administrator switches the user ID of other drama clubs.

-> **NOTE:** This parameter only applies during resource creation, update or deletion. If modified in isolation without other property changes, Terraform will not trigger any action.

* `vendor_id` - (Optional) The ID of the vendor corresponding to the standardization rule.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Normalization Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Normalization Rule.
* `update` - (Defaults to 5 mins) Used when update the Normalization Rule.

## Import

Threat Detection Normalization Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_normalization_rule.example <normalization_rule_id>
```