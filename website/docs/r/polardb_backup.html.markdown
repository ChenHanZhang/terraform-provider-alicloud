---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_backup"
description: |-
  Provides a Alicloud Polardb Backup resource.
---

# alicloud_polardb_backup

Provides a Polardb Backup resource.

Backing up objects at the cluster or database level.

For information about Polardb Backup and how to use it, see [What is Backup](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateBackup).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_polardb_backup" "default" {
  db_cluster_id = "pc-bp1593bkhcc4frx36"
}
```

## Argument Reference

The following arguments are supported:
* `backup_frequency` - (Optional) Backup frequency. Valid values are as follows:  
  - `Normal` (default): Standard backup performed once per day at a scheduled time.  
  - **2/24H**: Enhanced backup performed every 2 hours.  
  - **3/24H**: Enhanced backup performed every 3 hours.  
  - **4/24H**: Enhanced backup performed every 4 hours.  

-> **NOTE:**  * After enhanced backup is enabled, all backups completed within any 24-hour period are retained. For backups older than 24 hours, the system retains only the first backup completed after 00:00 UTC each day and deletes all others.  

-> **NOTE:**  * After enhanced backup is enabled, the `PreferredBackupPeriod` parameter defaults to all days of the week (Monday through Sunday).  

* `backup_retention_policy_on_cluster_deletion` - (Optional) Whether to retain backups when a cluster is deleted:
  - `ALL`: Retains all backups permanently.
  - `LATEST`: Retains only the latest backup permanently.
  - `NONE`: Does not retain any backup sets.

-> **NOTE:**  Default value: NONE.

* `db_cluster_id` - (Required, ForceNew) The cluster ID.

-> **NOTE:**  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to view information about all clusters in the target region, including their cluster IDs.

* `data_level1_backup_frequency` - (Optional) Backup frequency. Valid values are as follows:
  - `Normal` (default): Standard backup performed once per day at a scheduled time.
  - **2/24H**: High-frequency backup performed every 2 hours.
  - **3/24H**: High-frequency backup performed every 3 hours.
  - **4/24H**: High-frequency backup performed every 4 hours.

-> **NOTE:**  * PolarDB for PostgreSQL (compatible with Oracle) and PolarDB for PostgreSQL do not support this parameter.

-> **NOTE:**  * If cross-region backup is not supported in the region where your PolarDB for MySQL cluster is deployed, this parameter is not supported. For regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).

-> **NOTE:**  * After advanced backup is enabled, we recommend that you do not use this parameter. Instead, use the AdvancedDataPolicies parameter.

* `data_level1_backup_period` - (Optional) The primary backup cycle. Valid values are as follows:  
* `Monday`: Monday  
* `Tuesday`: Tuesday  
* `Wednesday`: Wednesday  
* `Thursday`: Thursday  
* `Friday`: Friday  
* `Saturday`: Saturday  
* `Sunday`: Sunday  

-> **NOTE:**  * * You must select at least two days. Separate multiple values with commas (,).  

-> **NOTE:** * * This parameter is not supported by PolarDB for PostgreSQL (compatible with Oracle) or PolarDB for PostgreSQL.  

-> **NOTE:** * * If the region where your PolarDB for MySQL cluster resides does not support cross-region backup, this parameter is unavailable. For regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).  

-> **NOTE:** * * After advanced backup is enabled, we recommend that you do not use this parameter. Use the `AdvancedDataPolicies` parameter instead.  

* `data_level1_backup_retention_period` - (Optional, Computed) Retention period for level-1 backups, in days. Valid values range from 3 to 14.  

-> **NOTE:**  * After advanced backup is enabled, this parameter no longer takes effect. Use the `AdvancedDataPolicies` parameter instead.  

* `data_level1_backup_time` - (Optional) The time window for automatic backups, in the format `hh:mmZ-hh:mmZ` (UTC). The specified time must be on the hour and span exactly one hour, such as `14:00Z-15:00Z`.

-> **NOTE:**  * This parameter is not supported by PolarDB for PostgreSQL (compatible with Oracle) or PolarDB for PostgreSQL.

-> **NOTE:**  * If cross-region backup is not supported in the region where your PolarDB for MySQL cluster resides, this parameter is not supported. For regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).

* `data_level2_backup_another_region_region` - (Optional) The region for cross-region Level-2 backup. For regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).

-> **NOTE:**  After advanced backup is enabled, we recommend that you do not use this parameter. Use the AdvancedDataPolicies parameter instead.

* `data_level2_backup_another_region_retention_period` - (Optional) Retention period for cross-region Level-2 backup. Valid values are as follows:
  - `0`: Disables the Level-2 backup feature.
  - **30–7300**: Retention period for Level-2 backup, in days.
  - **-1**: Retains Level-2 backups permanently.

-> **NOTE:** 

-> **NOTE:** - When you create a cluster, the default value is `0`, which disables cross-region Level-2 backup.

-> **NOTE:** - After advanced backup is enabled, we recommend that you do not use this parameter. Use the AdvancedDataPolicies parameter instead.

* `data_level2_backup_period` - (Optional) The secondary backup cycle. Valid values are as follows:  
* `Monday`: Monday  
* `Tuesday`: Tuesday  
* `Wednesday`: Wednesday  
* `Thursday`: Thursday  
* `Friday`: Friday  
* `Saturday`: Saturday  
* `Sunday`: Sunday  

-> **NOTE:**  * * You must select at least two days. Separate multiple values with commas (,).  

-> **NOTE:** * * This parameter is not supported by PolarDB for PostgreSQL (compatible with Oracle) or PolarDB for PostgreSQL.  

-> **NOTE:** * * If the region where your PolarDB for MySQL cluster resides does not support cross-region backup, this parameter is unavailable. For regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).  

-> **NOTE:** * * After advanced backup is enabled, we recommend that you do not use this parameter. Use the `AdvancedDataPolicies` parameter instead.  

* `data_level2_backup_retention_period` - (Optional, Computed) Retention period for secondary backups. Valid values are as follows:
 * 0: Disables the secondary backup feature.
 * 30–7300: Retention period for secondary backups, in days.
 * - 1: Retains secondary backups indefinitely.

 >

-> **NOTE:** - * When you create a cluster, the default value is `0`, which disables the secondary backup feature.

-> **NOTE:** - * After advanced backup is enabled, we recommend that you do not use this parameter. Instead, use the AdvancedDataPolicies parameter.

* `preferred_backup_period` - (Optional, Computed) Data backup cycle. Valid values are as follows:
  - Monday: Monday
  - Tuesday: Tuesday
  - Wednesday: Wednesday
  - Thursday: Thursday
  - Friday: Friday
  - Saturday: Saturday
  - Sunday: Sunday

-> **NOTE:**  After advanced backup is enabled, we recommend that you do not use this parameter. Instead, use the AdvancedDataPolicies parameter.

* `preferred_backup_time` - (Optional) Time window for performing automatic backups, in UTC. The format is `hh:mmZ-hh:mmZ`. The specified time must be on the hour and span exactly one hour, such as `14:00Z-15:00Z`.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Start time of the backup job in UTC.
* `status` - The backup status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Backup.
* `delete` - (Defaults to 5 mins) Used when delete the Backup.
* `update` - (Defaults to 5 mins) Used when update the Backup.

## Import

Polardb Backup can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_backup.example <db_cluster_id>
```