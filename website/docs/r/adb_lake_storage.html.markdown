---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_lake_storage"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Lake Storage resource.
---

# alicloud_adb_lake_storage

Provides a AnalyticDB for MySQL (ADB) Lake Storage resource.

ADB Lake storage resource definition.

For information about AnalyticDB for MySQL (ADB) Lake Storage and how to use it, see [What is Lake Storage](https://next.api.alibabacloud.com/document/adb/2021-12-01/CreateLakeStorage).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_vpc" "VPC" {
  dry_run    = false
  cidr_block = "172.16.0.0/12"
  vpc_name   = "APS结果集导出VPC"
}

resource "alicloud_vswitch" "VSWITCH" {
  vpc_id       = alicloud_vpc.VPC.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "APS结果集导出VSwitch"
}

resource "alicloud_adb_db_cluster_lake_version" "ADB" {
  storage_resource              = "24ACU"
  security_ips                  = "127.0.0.1"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.VPC.id
  vswitch_id                    = alicloud_vswitch.VSWITCH.id
  compute_resource              = "16ACU"
  db_cluster_version            = "5.0"
  payment_type                  = "Postpaid"
  enable_default_resource_group = true
  db_cluster_description        = "APS结果集导出OpenAPI集测20260206151502021010"
}

resource "alicloud_adb_lake_account" "LakeAccount" {
  db_cluster_id       = alicloud_adb_db_cluster_lake_version.ADB.id
  account_type        = "Super"
  account_name        = "apsexample"
  account_password    = "apsTest2024!"
  account_description = "OPENAPIexample"
}


resource "alicloud_adb_lake_storage" "default" {
  db_cluster_id = alicloud_adb_db_cluster_lake_version.ADB.id
  description   = "描述1"
}
```

## Argument Reference

The following arguments are supported:
* `db_cluster_id` - (Optional, ForceNew, Computed) The lake storage that specifies and mounts to a specific ADB master instance
* `description` - (Optional) Lake Storage Description
* `permissions` - (Optional, List) When creating a lake storage, the RAM account and primary account for the operation are automatically granted permissions, and additional primary account authorization can be added here. See [`permissions`](#permissions) below.

### `permissions`

The permissions supports the following:
* `account` - (Optional) Account ID
* `read` - (Optional) Read Permissions
* `type` - (Required) Account Type
* `write` - (Optional) Write permission

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<lake_storage_id>`.
* `create_time` - The time when the lake storage was created.
* `lake_storage_id` - The unique identifier used to recognize the specific lake storage.
* `region_id` - The ID of the region where the lake storage is located.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Lake Storage.
* `delete` - (Defaults to 5 mins) Used when delete the Lake Storage.
* `update` - (Defaults to 5 mins) Used when update the Lake Storage.

## Import

AnalyticDB for MySQL (ADB) Lake Storage can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_lake_storage.example <db_cluster_id>:<lake_storage_id>
```