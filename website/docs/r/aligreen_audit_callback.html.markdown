---
subcategory: "Aligreen"
layout: "alicloud"
page_title: "Alicloud: alicloud_aligreen_audit_callback"
description: |-
  Provides a Alicloud Aligreen Audit Callback resource.
---

# alicloud_aligreen_audit_callback

Provides a Aligreen Audit Callback resource.

Callback notifications sent after moderation is completed.

For information about Aligreen Audit Callback and how to use it, see [What is Audit Callback](https://next.api.alibabacloud.com/document/Green/2017-08-23/CreateAuditCallback).

-> **NOTE:** Available since v1.228.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform_example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_aligreen_audit_callback" "default" {
  crypt_type           = "SM3"
  audit_callback_name  = var.name
  url                  = "https://www.aliyun.com/"
  callback_types       = ["aliyunAudit", "selfAduit", "example"]
  callback_suggestions = ["block", "review", "pass"]
}
```

## Argument Reference

The following arguments are supported:
* `audit_callback_name` - (Required, ForceNew) The name of the callback configuration.
* `callback_suggestions` - (Required, List) The moderation result.
* `callback_types` - (Required, List) The notification type.
* `crypt_type` - (Required) The encryption algorithm for callbacks.
* `url` - (Required) The callback URL.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Audit Callback.
* `delete` - (Defaults to 5 mins) Used when delete the Audit Callback.
* `update` - (Defaults to 5 mins) Used when update the Audit Callback.

## Import

Aligreen Audit Callback can be imported using the id, e.g.

```shell
$ terraform import alicloud_aligreen_audit_callback.example <audit_callback_name>
```