---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_account"
description: |-
  Provides a Alicloud RDS Account resource.
---

# alicloud_rds_account

Provides a RDS Account resource.



For information about RDS Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createaccount).

-> **NOTE:** Available since v1.120.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf_example"
}

data "alicloud_db_zones" "default" {
  engine                   = "MySQL"
  engine_version           = "8.0"
  instance_charge_type     = "PostPaid"
  category                 = "HighAvailability"
  db_instance_storage_type = "local_ssd"
}

data "alicloud_db_instance_classes" "default" {
  zone_id                  = data.alicloud_db_zones.default.zones.0.id
  engine                   = "MySQL"
  engine_version           = "8.0"
  category                 = "HighAvailability"
  db_instance_storage_type = "local_ssd"
  instance_charge_type     = "PostPaid"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_db_zones.default.ids.0
  vswitch_name = var.name
}

resource "alicloud_db_instance" "default" {
  engine                   = "MySQL"
  engine_version           = "8.0"
  instance_type            = data.alicloud_db_instance_classes.default.instance_classes.0.instance_class
  instance_storage         = data.alicloud_db_instance_classes.default.instance_classes.0.storage_range.min
  vswitch_id               = alicloud_vswitch.default.id
  instance_name            = var.name
  instance_charge_type     = "Postpaid"
  monitoring_period        = 60
  db_instance_storage_type = "local_ssd"
  db_is_ignore_case        = false
}

resource "alicloud_rds_account" "default" {
  db_instance_id   = alicloud_db_instance.default.id
  account_name     = var.name
  account_password = "Example1234"
}
```

## Argument Reference

The following arguments are supported:
* `account_description` - (Optional, Computed) Account description. It can be 2 to 256 characters in length. It starts with a Chinese or English letter and can contain numbers, Chinese, English, underscores (_), and hyphens (-).
  -> **Note:** cannot start with http:// or https.
* `account_name` - (Required, ForceNew) Database account name
  -> **Description * *:
* Consists of lowercase letters, numbers, or underscores. For MySQL, uppercase letters are supported.
Start with a letter and end with a letter or number.
For MySQL, the common account name and the high-privilege account name cannot be similar to each other. For example, the high-privilege account name Test1 and the common account name cannot be test1.
Length:
* MySQL 8.0 and 5.7:2 to 32 characters.
* MySQL 5.6:2~16 characters.
* SQL Server:2 to 64 characters.
* PostgreSQL cloud disk version: 2~63 characters.
* PostgreSQL local disk version: 2 to 16 characters.
* MariaDB:2~16 characters.
Other illegal characters, see table of disabled keywords.
* `account_password` - (Required) The password of the database account.
  -> **NOTE: * *
The length is 8~32 characters.
Consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
The special character is! @#$%^& *()_+-=.
* `account_type` - (Optional, ForceNew, Computed) Account type, value:

  - Normal: Normal account (default)
  - Super: High-privilege account
  - Sysadmin: a database account with SA permissions (only RDS SQL Server instances are supported)
  - Before creating a database account with SA permissions, check whether the instance meets the prerequisites. For more information, see create a database account with SA permissions.
* `check_policy` - (Optional, Available since v1.266.0) Whether to apply password policy
* `db_instance_id` - (Required, ForceNew) The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
* `status` - (Optional, Computed) The status of the resource

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_instance_id>:<account_name>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Account.
* `delete` - (Defaults to 5 mins) Used when delete the Account.
* `update` - (Defaults to 5 mins) Used when update the Account.

## Import

RDS Account can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_account.example <db_instance_id>:<account_name>
```