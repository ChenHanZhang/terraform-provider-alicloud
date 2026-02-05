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
* `account_description` - (Optional, Computed) Account description, which must be 2 to 256 characters in length. It must start with a Chinese character or an English letter and can contain digits, Chinese characters, English letters, underscores (_), and hyphens (-).  

-> **NOTE:** It cannot start with `http://` or `https://`.

* `account_name` - (Required, ForceNew) Database account name.

-> **NOTE:**  The name must be unique and consist of uppercase letters (supported only by MySQL), lowercase letters, digits, or underscores. For specific naming restrictions, see the documentation for each database engine: [Create a MySQL account](https://help.aliyun.com/document_detail/96089.html), [Create a PostgreSQL account](https://help.aliyun.com/document_detail/96753.html), [Create a SQL Server account](https://help.aliyun.com/document_detail/95810.html), [Create a MariaDB account](https://help.aliyun.com/document_detail/97132.html).

* `account_password` - (Required) Password for the database account.

-> **NOTE:**  * Must be 8 to 32 characters in length.

-> **NOTE:**  * Must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters (`!@#$%^&*()_+-=`).

* `account_type` - (Optional, ForceNew, Computed) Account type. Valid values:
  - `Normal` (default): Standard account.
  - `Super`: Privileged account. You can create at most one such account per instance.
  - `Sysadmin` (SQL Server instances only): Database account with SA privileges. Before creating this account, verify that your instance meets the [prerequisites](https://help.aliyun.com/document_detail/170736.html).
  - `GlobalRO` (SQL Server instances only): Global read-only account. You can create at most two such accounts per instance. The instance must run SQL Server 2016 or later and belong to the dedicated or general-purpose instance family.
* `check_policy` - (Optional, Available since v1.266.0) Specifies whether to apply a password policy. Valid values:  
  - `true`: Apply a password policy to the account.  
  - `false`: Remove the password policy applied to the account.
* `db_instance_id` - (Required, ForceNew) Instance ID. You can call DescribeDBInstances to obtain it.
* `status` - (Optional, Computed) Account status. Valid values:  
  - `Unavailable`: Unavailable  
  - `Available`: Available

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