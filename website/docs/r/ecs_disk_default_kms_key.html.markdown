---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_disk_default_kms_key"
description: |-
  Provides a Alicloud Ecs Disk Default Kms Key resource.
---

# alicloud_ecs_disk_default_kms_key

Provides a Ecs Disk Default Kms Key resource.

The encryption key used by default for cloud storage encryption.

For information about Ecs Disk Default Kms Key and how to use it, see [What is Disk Default Kms Key](https://next.api.alibabacloud.com/document/Ecs/2014-05-26/ModifyDiskDefaultKMSKeyId).

-> **NOTE:** Available since v1.278.0.

-> **NOTE:** Destroying this resource will reset the default CMK to the account's AliCloud-managed default CMK for Ecs.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_kms_keys" "default" {
  filters = "[{\"Key\":\"KeyState\",\"Values\":[\"Enabled\"]}]"
}

resource "alicloud_ecs_disk_default_kms_key" "default" {
  kms_key_id = data.alicloud_kms_keys.default.ids.0
}
```

## Argument Reference

The following arguments are supported:
* `kms_key_id` - (Required) The ID of the KMS key.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Disk Default Kms Key.
* `delete` - (Defaults to 5 mins) Used when delete the Disk Default Kms Key.
* `update` - (Defaults to 5 mins) Used when update the Disk Default Kms Key.

## Import

Ecs Disk Default Kms Key can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_disk_default_kms_key.example <region_id>
```
