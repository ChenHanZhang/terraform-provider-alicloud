---
subcategory: "EIAM"
layout: "alicloud"
page_title: "Alicloud: alicloud_eiam_cloud_account"
description: |-
  Provides a Alicloud EIAM Cloud Account resource.
---

# alicloud_eiam_cloud_account

Provides a EIAM Cloud Account resource.

A cloud account is your unique account identifier on a cloud service provider and is used to identify your account entity.

For information about EIAM Cloud Account and how to use it, see [What is Cloud Account](https://next.api.alibabacloud.com/document/Eiam/2021-12-01/CreateCloudAccount).

-> **NOTE:** Available since v1.284.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `cloud_account_external_id` - (Required, ForceNew) The external unique identifier of the cloud account.
* `cloud_account_name` - (Optional) The cloud account name.
* `cloud_account_provider_name` - (Required, ForceNew) The identity provider name.
* `cloud_account_vendor_type` - (Required, ForceNew) The type of the cloud account.
* `description` - (Optional) The description of the cloud account.
* `instance_id` - (Required, ForceNew) The instance ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<cloud_account_id>`.
* `cloud_account_health` - The health status of the cloud account.
* `cloud_account_health_check_result` - The result of the cloud account health check.
  * `error_reason` - The cause of the error.
    * `error_code` - The error code.
    * `error_message` - The error message.
  * `last_check_time` - The time of the last check, in Unix timestamp format and milliseconds.
  * `result` - The health check result of the cloud account.
* `cloud_account_id` - The ID of the cloud account.
* `cloud_account_provider_config` - The identity provider configuration information.
  * `audience` - The audience identifier.
  * `authorization_server_id` - The authorization server ID.
  * `issuer` - Issuer.
  * `oidc_jwks_endpoint` - The signature verification public key endpoint.
* `create_time` - The creation time in Unix timestamp format, in milliseconds.
* `update_time` - The last update time in Unix timestamp format, in milliseconds.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cloud Account.
* `delete` - (Defaults to 5 mins) Used when delete the Cloud Account.
* `update` - (Defaults to 5 mins) Used when update the Cloud Account.

## Import

EIAM Cloud Account can be imported using the id, e.g.

```shell
$ terraform import alicloud_eiam_cloud_account.example <instance_id>:<cloud_account_id>
```