---
subcategory: "Tair (Redis OSS-Compatible) And Memcache (KVStore)"
layout: "alicloud"
page_title: "Alicloud: alicloud_kvstore_account"
description: |-
  Provides a Alicloud Tair (Redis OSS-Compatible) And Memcache (KVStore) Account resource.
---

# alicloud_kvstore_account

Provides a Tair (Redis OSS-Compatible) And Memcache (KVStore) Account resource.



For information about Tair (Redis OSS-Compatible) And Memcache (KVStore) Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/redis/developer-reference/api-r-kvstore-2015-01-01-createaccount-redis).

-> **NOTE:** Available since v1.66.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}
data "alicloud_kvstore_zones" "default" {

}
data "alicloud_resource_manager_resource_groups" "default" {
  status = "OK"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.4.0.0/16"
}
resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "10.4.0.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_kvstore_zones.default.zones.0.id
}

resource "alicloud_kvstore_instance" "default" {
  db_instance_name  = var.name
  vswitch_id        = alicloud_vswitch.default.id
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
  zone_id           = data.alicloud_kvstore_zones.default.zones.0.id
  instance_class    = "redis.master.large.default"
  instance_type     = "Redis"
  engine_version    = "5.0"
  security_ips      = ["10.23.12.24"]
  config = {
    appendonly             = "yes"
    lazyfree-lazy-eviction = "yes"
  }
  tags = {
    Created = "TF",
    For     = "example",
  }
}

resource "alicloud_kvstore_account" "default" {
  account_name     = "tfexamplename"
  account_password = "YourPassword_123"
  instance_id      = alicloud_kvstore_instance.default.id
}
```

## Argument Reference

The following arguments are supported:
* `account_name` - (Required, ForceNew) Account name.
* `account_password` - (Required) The password of the account. The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and specific special characters. These special characters include ! @ # $ % ^ & * ( ) _ + - =
* `account_privilege` - (Optional, Computed) The permissions of the account. Default value: RoleReadWrite. Valid values:
RoleReadOnly: The account has the read-only permissions.
RoleReadWrite: The account has the read and write permissions.
* `account_type` - (Optional, ForceNew, Computed) Account type.
* `description` - (Optional) Account description.
* `instance_id` - (Required, ForceNew) Database instance id

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<account_name>`.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 6 mins) Used when create the Account.
* `delete` - (Defaults to 5 mins) Used when delete the Account.
* `update` - (Defaults to 7 mins) Used when update the Account.

## Import

Tair (Redis OSS-Compatible) And Memcache (KVStore) Account can be imported using the id, e.g.

```shell
$ terraform import alicloud_kvstore_account.example <instance_id>:<account_name>
```