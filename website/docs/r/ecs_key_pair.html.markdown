---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_key_pair"
description: |-
  Provides a Alicloud ECS Key Pair resource.
---

# alicloud_ecs_key_pair

Provides a ECS Key Pair resource.



For information about ECS Key Pair and how to use it, see [What is Key Pair](https://www.alibabacloud.com/help/en/doc-detail/51771.htm).

-> **NOTE:** Available since v1.121.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_ecs_key_pair" "example" {
  key_pair_name = "key_pair_name"
}

// Using name prefix to build key pair
resource "alicloud_ecs_key_pair" "prefix" {
  key_name_prefix = "terraform-test-key-pair-prefix"
}

// Import an existing public key to build a alicloud key pair
resource "alicloud_ecs_key_pair" "publickey" {
  key_pair_name = "my_public_key"
  public_key    = "ssh-rsa AAAAB3Nza12345678qwertyuudsfsg"
}
```

## Argument Reference

The following arguments are supported:
* `key_pair_name` - (Required, ForceNew) The name of the key pair. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
* `public_key` - (Optional) The public key of the key pair.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `resource_group_id` - (Optional, Computed) The ID of the resource group to which to add the key pair.
* `tags` - (Optional, Map) A mapping of tags to assign to the resource.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The time when the key pair was created.
* `finger_print` - The fingerprint of the key pair.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Key Pair.
* `delete` - (Defaults to 5 mins) Used when delete the Key Pair.
* `update` - (Defaults to 5 mins) Used when update the Key Pair.

## Import

ECS Key Pair can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_key_pair.example <key_pair_name>
```