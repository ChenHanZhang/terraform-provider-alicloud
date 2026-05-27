---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_normalization_schema"
description: |-
  Provides a Alicloud Threat Detection Normalization Schema resource.
---

# alicloud_threat_detection_normalization_schema

Provides a Threat Detection Normalization Schema resource.

Alibaba Cloud Security Data Normalization Schema.

For information about Threat Detection Normalization Schema and how to use it, see [What is Normalization Schema](https://next.api.alibabacloud.com/document/cloud-siem/2024-12-12/CreateNormalizationSchema).

-> **NOTE:** Available since v1.280.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `lang` - (Optional) The language type for requests and responses. Valid values:
  - `zh` (default): Chinese.
  - `en`: English.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `normalization_category_id` - (Optional, ForceNew) The ID of the normalization rule category.
* `normalization_fields` - (Optional, ForceNew, List) The list of standard fields corresponding to the normalization schema. See [`normalization_fields`](#normalization_fields) below.
* `normalization_schema_description` - (Optional) Description of the normalization schema.
* `normalization_schema_id` - (Required, ForceNew) The ID of the normalization rule category.
* `normalization_schema_name` - (Optional) The name of the normalization schema.
* `normalization_schema_type` - (Required) The type of the normalization schema. Valid values:
  - log
  - entity
  - incident.
* `role_for` - (Optional, Int) The user ID used when an administrator switches to the perspective of another member.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `target_log_store` - (Required, ForceNew) The Log Service Logstore.

### `normalization_fields`

The normalization_fields supports the following:
* `normalization_field_description` - (Optional) A detailed description of the normalized field.
* `normalization_field_example` - (Optional) Example value of the normalized field.
* `normalization_field_from` - (Optional) The source of the normalized field.
* `normalization_field_json_index_all` - (Optional) Indicates whether all keys in a JSON-type normalized field are indexed.
* `normalization_field_json_keys` - (Optional, List) The list of keys for a JSON-type normalized field. See [`normalization_field_json_keys`](#normalization_fields-normalization_field_json_keys) below.
* `normalization_field_name` - (Required) The name of the normalized field.
* `normalization_field_required` - (Optional) Indicates whether the standard field is required.
* `normalization_field_requirement` - (Optional, ForceNew) This indicator specifies whether the field is required.
* `normalization_field_reserved` - (Optional, ForceNew) This indicator specifies whether the field is reserved by the system.
* `normalization_field_tokenize` - (Optional) Indicates whether the standard field is tokenized.
* `normalization_field_type` - (Required) Name of the normalized field.

### `normalization_fields-normalization_field_json_keys`

The normalization_fields-normalization_field_json_keys supports the following:
* `normalization_field_description` - (Optional) Description of the standard field key for JSON type.  
* `normalization_field_example` - (Optional) Example of a standard field key for JSON.
* `normalization_field_from` - (Optional) Source of the standard field for JSON type.  
* `normalization_field_name` - (Required) Name of the standard field key for JSON type.  
* `normalization_field_required` - (Optional) Indicates whether the standard field key for JSON is required.
* `normalization_field_tokenize` - (Optional) Indicates whether the standard field key for JSON is tokenized.
* `normalization_field_type` - (Required) Type of the standard field key for JSON type.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time.
* `normalization_schema_from` - Standardized structure Source: preset-predefined, custom-custom.
* `target_store_view` - The Log Service StoreView.
* `update_time` - The update time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Normalization Schema.
* `delete` - (Defaults to 5 mins) Used when delete the Normalization Schema.
* `update` - (Defaults to 5 mins) Used when update the Normalization Schema.

## Import

Threat Detection Normalization Schema can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_normalization_schema.example <normalization_schema_id>
```