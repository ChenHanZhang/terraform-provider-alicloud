---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_account"
description: |-
  Provides a Alicloud Polardb Account resource.
---

# alicloud_polardb_account

Provides a Polardb Account resource.

Database account information.

For information about Polardb Account and how to use it, see [What is Account](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateAccount).

-> **NOTE:** Available since v1.67.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_polardb_node_classes" "default" {
  db_type    = "MySQL"
  db_version = "8.0"
  pay_type   = "PostPaid"
  category   = "Normal"
}

resource "alicloud_vpc" "default" {
  vpc_name   = "terraform-example"
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_polardb_node_classes.default.classes[0].zone_id
  vswitch_name = "terraform-example"
}

resource "alicloud_polardb_cluster" "default" {
  db_type       = "MySQL"
  db_version    = "8.0"
  db_node_class = data.alicloud_polardb_node_classes.default.classes.0.supported_engines.0.available_resources.0.db_node_class
  pay_type      = "PostPaid"
  vswitch_id    = alicloud_vswitch.default.id
  description   = "terraform-example"
}

resource "alicloud_polardb_account" "default" {
  db_cluster_id       = alicloud_polardb_cluster.default.id
  account_name        = "terraform_example"
  account_password    = "Example1234"
  account_description = "terraform-example"
}
```

## Argument Reference

The following arguments are supported:
* `account_description` - (Optional) The account number Notes shall meet the following requirements:
  - Cannot start with' http:// 'or' https.
  - 2 to 256 characters in length.
* `account_lock_state` - (Optional, Computed, Available since v1.265.0) Account lock status. Value range:
  - `UnLock`: not locked.
  - `Lock`: locked.
* `account_name` - (Required, ForceNew) The account name must meet the following requirements:

  - Start with a lowercase letter and end with a letter or number.
  - Consists of lowercase letters, numbers, or underscores.
  - The length is 2 to 16 characters.
  - You cannot use some reserved usernames, such as root and admin.
* `account_password` - (Required) The account password. Must meet the following requirements:

  - Contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
  - Be 8 to 32 characters in length.
  - Special characters include !@#$%^&*()_+-=.
* `account_password_valid_time` - (Optional, Computed, Available since v1.265.0) Password expiration time
* `account_type` - (Optional, ForceNew, Computed) The type of the account. Valid values: 
* `Normal`: standard account 
* `Super`: privileged account. 

-> **NOTE:** * If you leave this parameter empty, the default value `Super` is used. 
* You can create multiple privileged accounts for a PolarDB for PostgreSQL (Compatible with Oracle) cluster or a PolarDB for PostgreSQL cluster. A privileged account has more permissions than a standard account. For more information, see [Create a database account](https://www.alibabacloud.com/help/en/doc-detail/68508.html). 
* You can create only one privileged account for a PolarDB for MySQL cluster. A privileged account has more permissions than a standard account. For more information, see [Create a database account](https://www.alibabacloud.com/help/en/doc-detail/68508.html).
* `db_cluster_id` - (Required, ForceNew) The cluster ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<account_name>`.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 8 mins) Used when create the Account.
* `delete` - (Defaults to 8 mins) Used when delete the Account.
* `update` - (Defaults to 5 mins) Used when update the Account.

## Import

Polardb Account can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_account.example <db_cluster_id>:<account_name>
```