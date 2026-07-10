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
* `application_federated_credential_name` - (Required, ForceNew) The name of the application federated credential.
* `application_federated_credential_type` - (Required, ForceNew) The type of the application federated credential.
* `application_id` - (Required, ForceNew) The application ID.
* `attribute_mappings` - (Optional, List) The attribute mappings. See [`attribute_mappings`](#attribute_mappings) below.
* `description` - (Optional) The description of the application federated credential.
* `federated_credential_provider_id` - (Required, ForceNew) The federated credential provider ID.
* `instance_id` - (Required, ForceNew) The instance ID.
* `oidc_verification_config` - (Optional, ForceNew, Set) The OIDC structured configuration (structured mode + oidc type). See [`oidc_verification_config`](#oidc_verification_config) below.
* `pkcs7_verification_config` - (Optional, Set) PKCS#7 structured configuration (structured mode + pkcs7 type). See [`pkcs7_verification_config`](#pkcs7_verification_config) below.
* `verification_condition` - (Optional) The verification condition.
* `verification_mode` - (Optional, ForceNew) The verification mode. Valid values:
  - freedom: Free mode. You can set the verification condition parameters (VerificationCondition) by using expressions.
  - structured: Structured mode. You can configure the relevant settings based on the specific type.

### `attribute_mappings`

The attribute_mappings supports the following:
* `source_value_expression` - (Optional) The source value expression.
* `target_field` - (Optional) The target field.

### `oidc_verification_config`

The oidc_verification_config supports the following:
* `azure_vm_config` - (Optional, ForceNew, Set) The Azure VM scenario configuration. See [`azure_vm_config`](#oidc_verification_config-azure_vm_config) below.
* `gcp_vm_config` - (Optional, Set) The configuration for the GCP VM scenario. See [`gcp_vm_config`](#oidc_verification_config-gcp_vm_config) below.
* `generic_config` - (Optional, Set) The configuration for generic scenarios. See [`generic_config`](#oidc_verification_config-generic_config) below.
* `kubernetes_config` - (Optional, Set) The configuration for Kubernetes scenarios. See [`kubernetes_config`](#oidc_verification_config-kubernetes_config) below.
* `profile` - (Optional) The standard OIDC profile type. Valid values:
  - generic: General scenarios
  - kubernetes: Kubernetes scenarios
  - gcp_vm: GCP virtual machine scenarios
  - azure_vm: Azure virtual machine scenarios

### `oidc_verification_config-azure_vm_config`

The oidc_verification_config-azure_vm_config supports the following:
* `principal_id` - (Optional) The principal identity ID.
* `resource_group_name` - (Optional) The resource group name.
* `subscription_id` - (Optional) The subscription ID.
* `vm_names` - (Optional, List) The list of virtual machine names.

### `oidc_verification_config-gcp_vm_config`

The oidc_verification_config-gcp_vm_config supports the following:
* `instance_ids` - (Optional, List) The list of VM instance IDs. A maximum of 10 IDs are supported.
* `project_id` - (Optional) The project ID.
* `service_account_id` - (Optional) The sub claim corresponding to the service account.

### `oidc_verification_config-generic_config`

The oidc_verification_config-generic_config supports the following:
* `subject` - (Optional) The subject identifier.

### `oidc_verification_config-kubernetes_config`

The oidc_verification_config-kubernetes_config supports the following:
* `namespace` - (Optional) The Kubernetes namespace.
* `pod_name_prefix` - (Optional) The pod name prefix.
* `service_account_name` - (Optional) The name of the Kubernetes service account.

### `pkcs7_verification_config`

The pkcs7_verification_config supports the following:
* `instance_ids` - (Optional, List) The list of allowed instance IDs. A maximum of 10 instance IDs are supported.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<application_id>:<application_federated_credential_id>`.
* `application_federated_credential_id` - The application federated credential ID.
* `create_time` - The creation time.
* `last_used_time` - The last used time.
* `status` - The status of the application federated credential.
* `update_time` - The update time.

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