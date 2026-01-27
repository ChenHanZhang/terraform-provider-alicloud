---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_zonal_database"
description: |-
  Provides a Alicloud Polardb Zonal Database resource.
---

# alicloud_polardb_zonal_database

Provides a Polardb Zonal Database resource.

DB resource of PolarDB cluster on MyBase.

For information about Polardb Zonal Database and how to use it, see [What is Zonal Database](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateDatabaseZonal).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `account_name` - (Optional, ForceNew) Home account name
* `character_set_name` - (Required, ForceNew) Character set
* `db_name` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `db_cluster_id` - (Required, ForceNew) PolarDB instance ID
* `db_description` - (Optional, ForceNew) PolarDB database description

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<db_name>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Zonal Database.
* `delete` - (Defaults to 5 mins) Used when delete the Zonal Database.
* `update` - (Defaults to 5 mins) Used when update the Zonal Database.

## Import

Polardb Zonal Database can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_zonal_database.example <db_cluster_id>:<db_name>
```