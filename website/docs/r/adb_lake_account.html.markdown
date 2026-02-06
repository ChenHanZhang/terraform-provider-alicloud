---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_lake_account"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Lake Account resource.
---

# alicloud_adb_lake_account

Provides a AnalyticDB for MySQL (ADB) Lake Account resource.

Account of the DBClusterLakeVesion.

For information about AnalyticDB for MySQL (ADB) Lake Account and how to use it, see [What is Lake Account](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-modifyaccountprivileges/).

-> **NOTE:** Available since v1.214.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "VPCID" {
  vpc_name = var.name

  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = var.name

  cidr_block = "172.16.0.0/24"
}

resource "alicloud_adb_db_cluster_lake_version" "CreateInstance" {
  storage_resource       = "0ACU"
  zone_id                = "cn-hangzhou-k"
  vpc_id                 = alicloud_vpc.VPCID.id
  vswitch_id             = alicloud_vswitch.VSWITCHID.id
  db_cluster_description = var.name
  compute_resource       = "16ACU"
  db_cluster_version     = "5.0"
  payment_type           = "PayAsYouGo"
  security_ips           = "127.0.0.1"
}


resource "alicloud_adb_lake_account" "default" {
  db_cluster_id    = alicloud_adb_db_cluster_lake_version.CreateInstance.id
  account_type     = "Super"
  account_name     = "tfnormal"
  account_password = "normal@2023"
  account_privileges {
    privilege_type = "Database"
    privilege_object {
      database = "MYSQL"
    }

    privileges = [
      "select",
      "update"
    ]
  }
  account_privileges {
    privilege_type = "Table"
    privilege_object {
      database = "INFORMATION_SCHEMA"
      table    = "ENGINES"
    }

    privileges = [
      "update"
    ]
  }
  account_privileges {
    privilege_type = "Column"
    privilege_object {
      table    = "COLUMNS"
      column   = "PRIVILEGES"
      database = "INFORMATION_SCHEMA"
    }

    privileges = [
      "update"
    ]
  }

  account_description = var.name
}
```

## Argument Reference

The following arguments are supported:
* `account_description` - (Optional) The description of the account
* `account_name` - (Required, ForceNew) The name of the account
* `account_password` - (Required) AccountPassword
* `account_privileges` - (Optional, Computed, List) List of permissions granted. See [`account_privileges`](#account_privileges) below.
* `account_type` - (Optional, ForceNew) The type of the account
* `db_cluster_id` - (Required, ForceNew) The DBCluster ID
* `engine` - (Optional, ForceNew, Computed, Available since v1.271.0) The engine, AnalyticDB.
* `ram_user_list` - (Optional, Computed, List, Available since v1.271.0) RamUsers

### `account_privileges`

The account_privileges supports the following:
* `privilege_object` - (Optional, Set) Object associated to privileges See [`privilege_object`](#account_privileges-privilege_object) below.
* `privilege_type` - (Optional) The type of privileges
* `privileges` - (Optional, List) privilege list

### `account_privileges-privilege_object`

The account_privileges-privilege_object supports the following:
* `column` - (Optional, Computed) The name of column
* `database` - (Optional, Computed) The name of database
* `table` - (Optional, Computed) The name of table

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<account_name>`.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Lake Account.
* `delete` - (Defaults to 5 mins) Used when delete the Lake Account.
* `update` - (Defaults to 5 mins) Used when update the Lake Account.

## Import

AnalyticDB for MySQL (ADB) Lake Account can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_lake_account.example <db_cluster_id>:<account_name>
```