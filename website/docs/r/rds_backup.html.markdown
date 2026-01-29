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

### Deleting `alicloud_rds_backup` or removing it from your configuration

The `alicloud_rds_backup` resource allows you to manage  `store_status = "Enabled"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `backup_job_id` - (Optional, ForceNew, Available since v1.270.0) The backup task ID.
* `backup_method` - (Optional, ForceNew, Computed) The backup type. Valid values:  
* `Logical`: logical backup (supported only for MySQL)  
* `Physical`: physical backup (supported for MySQL, SQL Server, and PostgreSQL)  
* `Snapshot`: snapshot backup (supported for all database engines)  

Default value: `Physical`.  

-> **NOTE:**  * When using logical backup, the database must contain data (the data cannot be empty).  

-> **NOTE:**  * MariaDB instances support only snapshot backup, but you must set this parameter to `Physical`.  

* `backup_retention_period` - (Optional, Int, Available since v1.270.0) When the SQL Server instance has BackupStrategy set to db, BackupMethod set to Physical, and BackupType set to FullBackup, you can specify the retention period for the backup set. Valid values: 7 to 730 days, or - 1 (permanent retention).

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `backup_strategy` - (Optional) Backup strategy. Valid values:
* `db`: Single-database backup
* `instance`: Instance-level backup

-> **NOTE:** This parameter takes effect only when the following conditions are met:

-> **NOTE:**  - MySQL: The `BackupMethod` parameter is specified with the value `Logical`.

-> **NOTE:**  - SQL Server: The `BackupType` parameter is specified with the value `FullBackup`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `backup_type` - (Optional, ForceNew, Computed) Backup type. Valid values:
  - `FullBackup`: Full backup
  - `IncrementalBackup`: Incremental backup
* `db_instance_id` - (Required, ForceNew) Instance ID. You can call DescribeDBInstances to obtain it.
* `db_name` - (Optional) List of databases. Separate multiple databases with commas (,).

-> **NOTE:**  This parameter takes effect only when the `BackupStrategy` parameter is specified and set to `db`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_instance_id>「:」<backup_id>`.
* `backup_id` - The backup set ID.
* `store_status` - Indicates whether the backup can be deleted.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Backup.
* `delete` - (Defaults to 5 mins) Used when delete the Backup.

## Import

RDS Backup can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_backup.example <db_instance_id>「:」<backup_id>
```