---
subcategory: "EIAM"
layout: "alicloud"
page_title: "Alicloud: alicloud_eiam_credential_provider"
description: |-
  Provides a Alicloud EIAM Credential Provider resource.
---

# alicloud_eiam_credential_provider

Provides a EIAM Credential Provider resource.

A credential provider used to generate authentication tokens, which manages the root trust or meta-configuration for generating authentication tokens.

For information about EIAM Credential Provider and how to use it, see [What is Credential Provider](https://next.api.alibabacloud.com/document/Eiam/2021-12-01/CreateCredentialProvider).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `credential_provider_identifier` - (Optional, ForceNew, Computed) The business identifier of the credential provider. It is a human-readable unique identifier.
* `credential_provider_name` - (Optional, Computed) Credential provider name.  
* `credential_provider_type` - (Required, ForceNew) The type of the credential provider.
* `description` - (Optional) Description.
* `instance_id` - (Required, ForceNew) Instance ID.
* `status` - (Optional, Computed) The status of the credential provider.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<credential_provider_id>`.
* `create_time` - The creation time of the credential provider, in Unix timestamp format.
* `credential_provider_config` - Credential provider configuration.
  * `jwt_provider_config` - JWT provider configuration.
    * `issuer` - The JWT issuer.
    * `jwks_endpoint` - The JWKs endpoint URL.
  * `provider_credential_ids` - A list of credential IDs corresponding to the sensitive configuration of the credential provider.
* `credential_provider_id` - List of credential provider IDs.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Credential Provider.
* `delete` - (Defaults to 5 mins) Used when delete the Credential Provider.
* `update` - (Defaults to 5 mins) Used when update the Credential Provider.

## Import

EIAM Credential Provider can be imported using the id, e.g.

```shell
$ terraform import alicloud_eiam_credential_provider.example <instance_id>:<credential_provider_id>
```