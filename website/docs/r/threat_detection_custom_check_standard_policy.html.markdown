---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_custom_check_standard_policy"
description: |-
  Provides a Alicloud Threat Detection Custom Check Standard Policy resource.
---

# alicloud_threat_detection_custom_check_standard_policy

Provides a Threat Detection Custom Check Standard Policy resource.

The standard policy resource to which the check item belongs.

For information about Threat Detection Custom Check Standard Policy and how to use it, see [What is Custom Check Standard Policy](https://next.api.alibabacloud.com/document/Sas/2018-12-03/CreateCheckPolicy).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_custom_check_standard_policy" "default" {
  type             = "KSPM"
  policy_show_name = "lingfengTest15"
}
```

## Argument Reference

The following arguments are supported:
* `policy_show_name` - (Required) The name of the newly added category.
* `type` - (Optional) The name of the associated policy category. This parameter is required when PolicyType is set to STANDARD:
  - `AISPM`: AI Security Posture Management (AISPM).
  - `IDENTITY_PERMISSION`: Cloud Infrastructure Entitlement Management (CIEM).
  - `RISK`: Security Risk.
  - `COMPLIANCE`: Compliance Risk.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `check_type` - The type of policy to query.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Custom Check Standard Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Custom Check Standard Policy.
* `update` - (Defaults to 5 mins) Used when update the Custom Check Standard Policy.

## Import

Threat Detection Custom Check Standard Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_custom_check_standard_policy.example <policy_id>
```