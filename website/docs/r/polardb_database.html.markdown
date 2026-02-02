---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_database"
description: |-
  Provides a Alicloud Polardb Database resource.
---

# alicloud_polardb_database

Provides a Polardb Database resource.

Manage linked databases.

For information about Polardb Database and how to use it, see [What is Database](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateDatabase).

-> **NOTE:** Available since v1.66.0.

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

resource "alicloud_polardb_database" "default" {
  db_cluster_id = alicloud_polardb_cluster.default.id
  db_name       = "terraform-example"
}
```

## Argument Reference

The following arguments are supported:
* `account_name` - (Optional) Account name
* `character_set_name` - (Required, ForceNew) The character set that is used by the cluster. For more information, see [Character set tables](https://www.alibabacloud.com/help/en/doc-detail/99716.html).
* `collate` - (Optional, Available since v1.265.0) The locale setting that specifies the collation of the new database.

-> **NOTE:** - The locale must be compatible with the character set set set by the `CharacterSetName` parameter.
  - This parameter is required when the cluster is the PolarDB O engine or the PolarDB PostgreSQL engine. This parameter is not required when the cluster is the PolarDB MySQL engine.
For the value range of this parameter, log on to the PolarDB console. On The **Configuration and Management**> **Database Management** tab, click **Create Database.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `ctype` - (Optional, Available since v1.265.0) The locale setting that specifies the character classification of the database.

-> **NOTE:** - The locale must be compatible with the character set set set by the `CharacterSetName` parameter.
  - Consistent with `Collate` incoming information.
  - This parameter is required when the cluster is the PolarDB O engine or the PolarDB PostgreSQL engine. This parameter is not required when the cluster is the PolarDB MySQL engine.
For the value range of this parameter, log on to the PolarDB console. On The **Configuration and Management**> **Database Management** tab, click **Create Database.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_cluster_id` - (Required, ForceNew) The ID of cluster.
* `db_name` - (Required, ForceNew) The name of the database.
* `db_description` - (Optional, Computed) The description of the database. The description must meet the following requirements:
  - It cannot start with `http://` or `https://`.
  - It must be 2 to 256 characters in length.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<db_name>`.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Database.
* `delete` - (Defaults to 30 mins) Used when delete the Database.
* `update` - (Defaults to 5 mins) Used when update the Database.

## Import

Polardb Database can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_database.example <db_cluster_id>:<db_name>
```