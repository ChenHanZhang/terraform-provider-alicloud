---
subcategory: "EIAM"
layout: "alicloud"
page_title: "Alicloud: alicloud_eiam_credential"
description: |-
  Provides a Alicloud EIAM Credential resource.
---

# alicloud_eiam_credential

Provides a EIAM Credential resource.

Credentials used to authenticate the associated principal, which are typically long-lived, support credential versioning, and need to be rotated. For example, the password and key pair of the system user in the Linux host.

For information about EIAM Credential and how to use it, see [What is Credential](https://next.api.alibabacloud.com/document/Eiam/2021-12-01/CreateCredential).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `credential_identifier` - (Required, ForceNew) Business identifier of the credential.  
* `credential_name` - (Required) Credential name.  
* `credential_scenario_label` - (Optional, ForceNew) Usage scenario label of the credential.  
* `credential_subject_id` - (Optional, ForceNew) ID of the subject to which the credential belongs.  
* `credential_subject_type` - (Optional, ForceNew) The type of principal to which the credential belongs.  
* `credential_type` - (Required, ForceNew) Credential type.  
* `description` - (Required) Description.
* `instance_id` - (Required, ForceNew) EIAM instance ID.  
* `status` - (Optional, Computed) Credential status.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<credential_id>`.
* `create_time` - Creation time.
* `credential_id` - Credential ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Credential.
* `delete` - (Defaults to 5 mins) Used when delete the Credential.
* `update` - (Defaults to 5 mins) Used when update the Credential.

## Import

EIAM Credential can be imported using the id, e.g.

```shell
$ terraform import alicloud_eiam_credential.example <instance_id>:<credential_id>
```