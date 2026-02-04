---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_db_backup_policy"
description: |-
  Provides a Alicloud RDS Backup Policy resource.
---

# alicloud_db_backup_policy

Provides a RDS Backup Policy resource.

Instance Backup settings.

For information about RDS Backup Policy and how to use it, see [What is Backup Policy](https://next.api.alibabacloud.com/document/Rds/2014-08-15/ModifyBackupPolicy).

-> **NOTE:** Available since v1.5.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}
data "alicloud_db_zones" "default" {
  engine         = "MySQL"
  engine_version = "5.6"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_db_zones.default.zones.0.id
  vswitch_name = var.name
}

resource "alicloud_db_instance" "instance" {
  engine           = "MySQL"
  engine_version   = "5.6"
  instance_type    = "rds.mysql.s1.small"
  instance_storage = "10"
  vswitch_id       = alicloud_vswitch.default.id
  instance_name    = var.name
}

resource "alicloud_db_backup_policy" "policy" {
  instance_id = alicloud_db_instance.instance.id
}
```

### Deleting `alicloud_db_backup_policy` or removing it from your configuration

Terraform cannot destroy resource `alicloud_db_backup_policy`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `archive_backup_keep_count` - (Optional, Computed, Available since v1.69.0) Number of archived backups retained. The default is 1. Value:
When ArchiveBackupKeepPolicy is set to ByMonth, the value is from 1 to 31.
When ArchiveBackupKeepPolicy is set to ByWeek, the value is from 1 to 7.
Description
When ArchiveBackupKeepPolicy is set to KeepAll, this parameter is not required.
Takes effect only when the BackupPolicyMode parameter is databackuppolicyy.
* `archive_backup_keep_policy` - (Optional, Computed, Available since v1.69.0) The retention period for archived backups. The number of backups that can be saved in this period is determined by ArchiveBackupKeepCount. The default is 0.
Value:
ByMonth: Month
ByWeek: Week
KeepAll: Keep all
Note takes effect only when the BackupPolicyMode parameter is databackuppolicyy.
* `archive_backup_retention_period` - (Optional, Computed, Available since v1.69.0) Number of days to keep Archived backups. The default value is 0, indicating that the archive backup is not enabled. Value: 30~1095.
Description:
Takes effect only when the BackupPolicyMode parameter is databackuppolicyy.
* `backup_interval` - (Optional, Computed, Available since v1.194.0) Backup interval:
  - For MySQL instances, [snapshot backup frequency](~~ 98818 ~~)(non-snapshot backup cycle) is used.
  - For SQL Server instances, log backup frequency.
* `backup_log` - (Optional, Available since v1.271.0) Log backup switch. Value: **Enable | Disabled * *
* `backup_method` - (Optional, Computed) The backup method of the instance. Return value:
* `Physical`: Physical backup
`Snapshot`: Snapshot backup

-> **NOTE:**  This parameter is returned only for SQL Server cloud disk instances.

* `backup_policy_mode` - (Optional, Available since v1.271.0) Backup type:

  - `DataBackupPolicy`: Data backup
  - `LogBackupPolicy`: log backup

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `backup_priority` - (Optional, Int, Available since v1.229.1) Set backup options. Value:
1: Priority Reserve
2: Mandatory Master Library

Description
This parameter applies only to SQL Server cluster instances.
This parameter takes effect only when BackupMethod is set to Physical. If BackupMethod is set to Snapshot,SQL Server Cluster Edition instances will be forced to back up on the primary library.
* `backup_retention_period` - (Optional, Int, Available since v1.69.0) Data backup retention days, value: 7~730.
Description
When BackupPolicyMode is databackuppolicymode, this parameter must be set.
Takes effect only when the BackupPolicyMode parameter is databackuppolicyy.
* `category` - (Optional, Computed, Available since v1.190.0) Instance series. Valid values:

  - `Basic`: Basic Edition
  - `HighAvailability`: High availability
  - `AlwaysOn`: Cluster Edition
  - `Finance`: Three-node Enterprise Edition
* `compress_type` - (Optional, Computed, Available since v1.69.0) Backup compression method, value:
  - `0`: No compression
  - `1`:zlib compression
  - `2`: parallel zlib compression
  - `4`:quicklz compression, enabling database and table recovery
  - `8`:MySQL8.0 quicklz compression but library table recovery is not supported.
* `db_instance_id` - (Required, ForceNew, Available since v1.271.0) The ID of the instance.
* `enable_backup_log` - (Optional, Computed, Available since v1.68.0) Whether to enable log Backup. Valid values:
  - `1`: indicates enabled
  - `0`: indicates closed
* `enable_increment_data_backup` - (Optional, Computed, Available since v1.229.1) Whether to enable incremental backup. Value:
* `False` (default): Off
* `True`: On

-> **NOTE:**  * this parameter takes effect only for SQL Server cloud disk instances.

-> **NOTE:**  * only takes effect when the `BackupPolicyMode` parameter is **databackuppolicyy.

* `high_space_usage_protection` - (Optional, Available since v1.69.0) If the instance usage space is greater than 80% or the remaining space is less than 5GB, whether to force Binlog cleanup:

  - `Disable`: do not clean up
  - `Enable`: clean up
* `local_log_retention_hours` - (Optional, Computed, Int, Available since v1.69.0) Log backup local retention hours.
* `local_log_retention_space` - (Optional, Computed, Available since v1.69.0) The maximum circular space usage of local logs. If the maximum circular space usage is exceeded, the earliest Binlog is cleared until the space usage is lower than this ratio. Value: 0~50. Default is not modified.
Description
When BackupPolicyMode is set to logbackuppolicymode, this parameter must be passed.
Takes effect only when the BackupPolicyMode parameter is set to LogBackupPolicy.
* `log_backup_frequency` - (Optional, Computed, Available since v1.69.0) Log backup frequency, value:

* `LogInterval`: Back up every 30 minutes;
* The default data backup cycle is the same as the data backup cycle **PreferredBackupPeriod.

-> **NOTE:**  parameter `LogBackupFrequency` is only applicable to SQL Server.

* `log_backup_local_retention_number` - (Optional, Computed, Int, Available since v1.229.1) The number of local binlogs retained. The default is 60. Value: 6~100.
Description
Takes effect only when the BackupPolicyMode parameter is set to LogBackupPolicy.
If the instance type is MySQL, you can pass in **-1**, that is, the number of reserved local binlogs is not limited.
* `log_backup_retention_period` - (Optional, Computed, Int, Available since v1.69.0) The number of days for which the log backup is retained. Valid values: 7 to 730. The log backup retention period cannot be longer than the data backup retention period.
Note
If you enable the log backup feature, you can specify the log backup retention period. This parameter is supported for instances that run MySQL and PostgreSQL.
This parameter takes effect only when BackupPolicyMode is set to DataBackupPolicy or LogBackupPolicy.
* `preferred_backup_period` - (Optional, Computed, Available since v1.69.0) Data backup cycle. Separate multiple values with commas (,). Valid values:
  - `Monday`: Monday
  - `Tuesday`: Tuesday
  - `Wednesday`: Wednesday
  - `Thursday`: Thursday
  - `Friday`: Friday
  - `Saturday`: Saturday
  - `Sunday`: Sunday
* `preferred_backup_time` - (Optional, Available since v1.69.0) Data backup time, format:  HH:mm Z- HH:mm Z(UTC time).
* `released_keep_policy` - (Optional, Computed, Available since v1.147.0) Archive backup retention policy for deleted instances. Value:
  - `None`: not reserved
  - `Lastest`: Keep the last one
  - `All`: All reserved

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Backup Policy.
* `update` - (Defaults to 5 mins) Used when update the Backup Policy.

## Import

RDS Backup Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_db_backup_policy.example <db_instance_id>
```