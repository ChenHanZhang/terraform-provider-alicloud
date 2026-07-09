---
subcategory: "EIAM"
layout: "alicloud"
page_title: "Alicloud: alicloud_eiam_application_federated_credential"
description: |-
  Provides a Alicloud EIAM Application Federated Credential resource.
---

# alicloud_eiam_application_federated_credential

Provides a EIAM Application Federated Credential resource.

Application Federated Credential.

For information about EIAM Application Federated Credential and how to use it, see [What is Application Federated Credential](https://next.api.alibabacloud.com/document/Eiam/2021-12-01/CreateApplicationFederatedCredential).

-> **NOTE:** Available since v1.285.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `application_federated_credential_name` - (Required, ForceNew) Application Federated Credential Name
* `application_federated_credential_type` - (Required, ForceNew) Application Federation Credential Type
* `application_id` - (Required, ForceNew) 应用ID
* `attribute_mappings` - (Optional, List) Attribute Mapping See [`attribute_mappings`](#attribute_mappings) below.
* `description` - (Optional) 应用联邦凭证描述
* `federated_credential_provider_id` - (Required, ForceNew) Application Federation credential provider ID
* `instance_id` - (Required, ForceNew) EAIM 实例ID
* `verification_condition` - (Optional) 验证条件

### `attribute_mappings`

The attribute_mappings supports the following:
* `source_value_expression` - (Optional) Source Value Expression
* `target_field` - (Optional) Target Field

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<application_id>:<application_federated_credential_id>`.
* `application_federated_credential_id` - 
* `create_time` - Creation time.
* `last_used_time` - Recently Used.
* `status` - Application Federation Credential Status.
* `update_time` - Update time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Application Federated Credential.
* `delete` - (Defaults to 5 mins) Used when delete the Application Federated Credential.
* `update` - (Defaults to 5 mins) Used when update the Application Federated Credential.

## Import

EIAM Application Federated Credential can be imported using the id, e.g.

```shell
$ terraform import alicloud_eiam_application_federated_credential.example <instance_id>:<application_id>:<application_federated_credential_id>
```