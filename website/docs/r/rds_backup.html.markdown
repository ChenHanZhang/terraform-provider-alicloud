---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_backup"
description: |-
  Provides a Alicloud RDS Backup resource.
---

# alicloud_rds_backup

Provides a RDS Backup resource.

Instance-level or database-level backup objects.

For information about RDS Backup and how to use it, see [What is Backup](https://www.alibabacloud.com/help/en/rds/developer-reference/api-rds-2014-08-15-createbackup).

-> **NOTE:** Available since v1.149.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_db_instance" "example" {
  engine                   = "MySQL"
  engine_version           = "5.6"
  instance_type            = "rds.mysql.t1.small"
  instance_storage         = "30"
  instance_charge_type     = "Postpaid"
  db_instance_storage_type = "local_ssd"
}

resource "alicloud_rds_backup" "example" {
  db_instance_id = alicloud_db_instance.example.id
}
```

## Argument Reference

The following arguments are supported:
* `backup_method` - (Optional, ForceNew, Computed) The backup type of the instance. Valid values: 
  - `Logical`: logical backup 
  - `Physical`: physical backup 
  - `Snapshot`: snapshot backup 

Default value: `Physical`. 

&gt; * You can perform a logical backup only when databases are created on the instance. &gt; * When you perform a snapshot backup on an ApsaraDB RDS for MariaDB instance, you must set this parameter to `Physical`.

&gt; * For more information about the supported backup types, see [Use the data backup feature](https://www.alibabacloud.com/help/en/doc-detail/98818.html). &gt;

 * When you perform a snapshot backup on an ApsaraDB RDS for SQL Server instance that uses cloud disks, you must set this parameter to `Snapshot`.
* `backup_retention_period` - (Optional, Int, Available since v1.270.0) When the BackupStrategy of an SQL Serevr instance is set to db, BackupMethod is set to Physical, and BackupType is set to FullBackup, you can specify the retention time of the backup set. The value is 7 to 730 days, - 1 (Long-term retention).

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `backup_strategy` - (Optional) Backup policy, value:
* `db`: Single-database backup
* `instance`: instance Backup

-> **NOTE:**  This parameter takes effect when the following conditions are met:

-> **NOTE:**  - MySQL: The `BackupMethod` parameter (value `Logical`) has been passed in.

-> **NOTE:**  - SQL Server: The `BackupType` parameter (value: `FullBackup`) has been passed in.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `backup_type` - (Optional, ForceNew, Computed) The backup type. Valid values: 
  - `FullBackup`: full backup 
  - `IncrementalBackup`: incremental backup 
* `db_instance_id` - (Required, ForceNew) The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
* `db_name` - (Optional) A list of databases. Separate multiple databases with commas (,).

-> **NOTE:**  This parameter takes effect when the following conditions are met: `BackupStrategy` parameter (value: `db`) has been passed in.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `resource_group_id` - (Optional, ForceNew, Computed, Available since v1.270.0) The ID of the resource group

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `backup_id` - The ID of the backup set.
* `status` - The status of the resource.
* `store_status` - Whether the data backup can be deleted.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Backup.
* `delete` - (Defaults to 5 mins) Used when delete the Backup.

## Import

RDS Backup can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_backup.example <backup_id>
```