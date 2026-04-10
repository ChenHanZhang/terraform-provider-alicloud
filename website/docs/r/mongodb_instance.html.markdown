---
subcategory: "MongoDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_mongodb_instance"
description: |-
  Provides a Alicloud Mongodb Instance resource.
---

# alicloud_mongodb_instance

Provides a Mongodb Instance resource.

ApsaraDB for MongoDB instances integrate computing, storage, and network resources, and provide management features for database users, databases and tables, networking, and backups. You can easily customize and modify instance configurations.

For information about Mongodb Instance and how to use it, see [What is Instance](https://next.api.alibabacloud.com/document/Dds/2015-12-01/CreateDBInstance).

-> **NOTE:** Available since v1.53.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_mongodb_zones" "default" {
}

locals {
  index   = length(data.alicloud_mongodb_zones.default.zones) - 1
  zone_id = data.alicloud_mongodb_zones.default.zones[local.index].id
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.17.3.0/24"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "172.17.3.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = local.zone_id
}

resource "alicloud_mongodb_instance" "default" {
  engine_version      = "4.2"
  db_instance_class   = "dds.mongo.mid"
  db_instance_storage = 10
  vswitch_id          = alicloud_vswitch.default.id
  security_ip_list    = ["10.168.1.12", "100.69.7.112"]
  name                = var.name
  tags = {
    Created = "TF"
    For     = "example"
  }
}
```

## Argument Reference

The following arguments are supported:
* `account_description` - (Optional, Available since v1.276.0) Set a description for the account.  
  - It cannot start with http:// or https://.  
  - It must start with a Chinese character or an English letter.  
  - It can contain Chinese characters, English letters, underscores (_), hyphens (-), and digits. The length must be 2 to 256 characters.
* `account_password` - (Optional) The password for the root account. Requirements:
  - Must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
  - Special characters include: !@#$%^&*()_+-=
  - Length must be 8 to 32 characters.

-> **NOTE:**  For issues related to database connection failures caused by special characters in passwords, see [How do I resolve connection failures caused by special characters in the username or password in the connection string?](https://help.aliyun.com/document_detail/471568.html).


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `auto_pay` - (Optional, Available since v1.276.0) Specifies whether to enable auto-payment for the order. Valid values:
  - `true`: Automatically pay for the order.
  - `false`: Manually pay for the order. For more information, see [Manually renew a subscription instance](https://help.aliyun.com/document_detail/85052.html).

-> **NOTE:**  Default value: `true`.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_renew` - (Optional, Available since v1.141.0) Specifies whether auto-renewal is enabled for the instance. Valid values:  
  - `true`: Auto-renewal is enabled.  
  - `false`: Auto-renewal is disabled (default). Manual renewal is required.  

-> **NOTE:**  This parameter is optional and available only when the `ChargeType` parameter is set to `PrePaid`.

* `auto_renew_duration` - (Optional, Computed, Int, Available since v1.271.0) The auto-renewal duration, in months.

Valid values: integers from `1` to `12`.

-> **NOTE:**  This parameter is available and required only when the `AutoRenew` parameter is set to `true`.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `backup_id` - (Optional, Available since v1.276.0) Cluster backup ID.

-> **NOTE:**  - This parameter is required and mandatory only when RestoreType is 2 or 3.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `backup_interval` - (Optional, Computed, Available since v1.212.0) High-frequency backup interval. Valid values:
  - **-1**: High-frequency backup is disabled.
  - `15`: Every 15 minutes.
  - `30`: Every 30 minutes.
  - `60`: Every 1 hour.
  - `120`: Every 2 hours.
  - `180`: Every 3 hours.
  - `240`: Every 4 hours.
  - `360`: Every 6 hours.
  - `480`: Every 8 hours.
  - `720`: Every 12 hours.

-> **NOTE:**  - When `SnapshotBackupType` is set to `Standard`, this parameter must be **-1**.

-> **NOTE:**  - When `SnapshotBackupType` is set to `Flash` and this parameter is greater than 0, high-frequency backup takes effect.

* `backup_retention_period` - (Optional, Computed, Available since v1.213.1) Number of days to retain full backups.

-> **NOTE:**  - For users created before September 10, 2021, the default retention period is 7 days.

-> **NOTE:**  - For new users created on or after September 10, 2021, the default retention period is 30 days.

* `backup_retention_policy_on_cluster_deletion` - (Optional, Int, Available since v1.235.0) Backup retention policy.  
1. 0: Delete all backup sets of the instance immediately when the instance is released.  
2. 1: Automatically create a backup when the instance is released and retain this backup set permanently.  
3. 2: Automatically create a backup when the instance is released and retain all backup sets of the cluster permanently.  
For more information, see [Long-term backup retention](https://help.aliyun.com/document_detail/2779111.html).
* `character_type` - (Optional, Available since v1.276.0) The role type of the instance. Valid values:

  - When the instance type is a sharded cluster, CharacterType is required and can be set to `db` or `cs`.
  - When the instance type is a replica set, CharacterType can be left empty or set to `normal`.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `cloud_disk_encryption_key` - (Optional, ForceNew, Available since v1.212.0) The custom key ID.  
* `configserver_list` - (Optional, ForceNew, List, Available since v1.276.0) Information about ConfigServer nodes. See [`configserver_list`](#configserver_list) below.
* `coupon_no` - (Optional, Available since v1.276.0) Specifies whether to use a coupon. Valid values:  
  - `default` or `null` (default): Use a coupon.  
  - `youhuiquan_promotion_option_id_for_blank`: Do not use a coupon.  

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `cross_backup_period` - (Optional, Available since v1.276.0) Retention period for cross-region backups.  

1. Monday: Monday  
2. Tuesday: Tuesday  
3. Wednesday: Wednesday  
4. Thursday: Thursday  
5. Friday: Friday  
6. Saturday: Saturday  
7. Sunday: Sunday  

-> **NOTE:**  Required for cross-region backup.  

-> **NOTE:**  - To specify multiple values, separate them with commas (,).  

-> **NOTE:**  - When the backup type is set to regular backup, this value must be a subset of the PreferredBackupPeriod.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `cross_backup_type` - (Optional, Available since v1.276.0) Remote backup operation policy:
  - update: Modify the remote backup policy.
  - delete: Delete the remote backup policy.

-> **NOTE:**  Required for remote backup.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `cross_log_retention_type` - (Optional, Available since v1.276.0) The retention type for cross-region log backups.
  - delay (retain for a specified period)
  - never (never expire)

-> **NOTE:**  This parameter is required for cross-region backup.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `cross_log_retention_value` - (Optional, Int, Available since v1.276.0) The retention period for cross-region log backups, which can be set from 3 to 1825 days and must be less than or equal to the value of CrossRetentionValue.

-> **NOTE:**  This parameter is required for cross-region backup.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `cross_retention_type` - (Optional, Available since v1.276.0) Cross-region backup retention type.  
  - delay (retain for a specific period)  
  - never (never expire)  

-> **NOTE:**  This parameter is required for cross-region backups.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `cross_retention_value` - (Optional, Int, Available since v1.276.0) The retention period (in days) for cross-region backups. You can set this value to a number between 3 and 1825.

-> **NOTE:**  

-> **NOTE:**  - This parameter is required for cross-region backups.

-> **NOTE:**  - This parameter must be specified when CrossRetentionType is set to `delay`.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `db_instance_description` - (Optional, Available since v1.276.0) The description or remarks for the instance.
* `db_instance_release_protection` - (Optional, Available since v1.253.0) Specifies whether to enable release protection for the instance. Valid values:  
  - `true`: Enabled.  
  - `false`: Disabled.
* `database_names` - (Optional, Available since v1.276.0) Database names.

-> **NOTE:**  When you call this operation to clone an instance, you can use this parameter to specify the databases to clone. If you do not configure this parameter, all databases of the instance will be cloned.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `db_instance_class` - (Optional) Instance type.For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html). You can also call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to view available instance types.  
  
  - For standalone instances and replica set instances, this parameter specifies the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html). You can also call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to view available instance types for standalone and replica set instances.  
  - For Serverless instances, this parameter specifies the compute capacity, with a value range of 100 to 8000.  
  

-> **NOTE:**  You must configure at least one of this parameter or the `DBInstanceStorage` parameter.

* `db_instance_storage` - (Optional, Int) The storage capacity of the instance.Valid values range from 10 GB to 3000 GB, in increments of 10 GB. The actual valid values depend on the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).

  - For single-node and replica set instances, valid storage capacity values range from 10 GB to 3000 GB, in increments of 10 GB. The actual valid values depend on the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
  - For Serverless instances, valid storage capacity values range from 1 GB to 100 GB, in increments of 1 GB.


-> **NOTE:**  - You must specify at least one of this parameter or the `DBInstanceClass` parameter.

-> **NOTE:**  - Downgrading the storage capacity is currently not supported.

* `dest_region` - (Optional, Available since v1.276.0) The region where the backup is stored.

-> **NOTE:**  Required for cross-region backups.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `effective_time` - (Optional, Available since v1.215.0) The effective time for migrating to a different zone. Valid values:
  - `Immediately` (default): Takes effect immediately.
  - `MaintainTime`: Takes effect during the instance maintenance window.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `enable_backup_log` - (Optional, Computed, Int, Available since v1.230.1) Specifies whether to enable log backup. Valid values:  
  - `0`: Disabled (default).  
  - `1`: Enabled.
* `enable_cross_log_backup` - (Optional, Int, Available since v1.276.0) Specifies whether to enable cross-region log backup.  

-> **NOTE:**  This parameter is required for cross-region backups.  

-> **NOTE:**  - 1: Enable (must be set to 1 for sharded clusters; must be set to 1 for replica sets if you want to support point-in-time recovery across regions)  

-> **NOTE:**  - 0: Disable.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `encrypted` - (Optional, ForceNew, Available since v1.212.0) Specifies whether to enable disk encryption.
* `encryption_key` - (Optional, Computed, Available since v1.212.0) The custom key ID.  
Custom keys are supported only in the following regions. Default keys are used for encryption in all other regions:  
  - Singapore (ap-southeast-1)  
  - Hangzhou (cn-hangzhou)  
  - Shanghai (cn-shanghai)  
  - Beijing (cn-beijing)  
  - Shenzhen (cn-shenzhen)  
  - Hong Kong (cn-hongkong)  
  - Malaysia (ap-southeast-3).  

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `encryptor_name` - (Optional, Computed, Available since v1.212.0) Encryption method. Valid value: **aes-256-cbc**.

-> **NOTE:**  This parameter is available only when the `TdeStatus` parameter is set to `enabled`.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `engine` - (Required, ForceNew, Available since v1.276.0) Database engine. Valid value: `MongoDB`.
* `engine_version` - (Required) Database version. Valid values:  
  - **8.0**  
  - **7.0**  
  - **6.0**  
  - **5.0**  
  - **4.4**  
  - **4.2**  
  - **4.0**  

-> **NOTE:**  When you call this operation to clone an instance or restore an instance from the recycle bin, the value of this parameter must be the same as that of the source instance.

-> **NOTE:** Versions 3.4 and earlier are no longer available for sale.

* `global_security_group_list` - (Optional, List, Available since v1.257.0) The ID of the IP whitelist template.
* `hidden_zone_id` - (Optional, ForceNew, Available since v1.199.0) The zone where the hidden node is deployed to enable multi-zone deployment. Valid values:  
  - **cn-hangzhou-g**: Hangzhou Zone G  
  - **cn-hangzhou-h**: Hangzhou Zone H  
  - **cn-hangzhou-i**: Hangzhou Zone I  
  - **cn-hongkong-b**: Hong Kong Zone B  
  - **cn-hongkong-c**: Hong Kong Zone C  
  - **cn-hongkong-d**: Hong Kong Zone D  
  - **cn-wulanchabu-a**: Ulanqab Zone A  
  - **cn-wulanchabu-b**: Ulanqab Zone B  
  - **cn-wulanchabu-c**: Ulanqab Zone C  
  - **ap-southeast-1a**: Singapore Zone A  
  - **ap-southeast-1b**: Singapore Zone B  
  - **ap-southeast-1c**: Singapore Zone C  
  - **ap-southeast-5a**: Jakarta Zone A  
  - **ap-southeast-5b**: Jakarta Zone B  
  - **ap-southeast-5c**: Jakarta Zone C  
  - **eu-central-1a**: Frankfurt Zone A  
  - **eu-central-1b**: Frankfurt Zone B  
  - **eu-central-1c**: Frankfurt Zone C  

-> **NOTE:**  - This parameter is available only for instances that use cloud disks.  

-> **NOTE:**  - The value of this parameter must be different from the values of the `ZoneId` and `SecondaryZoneId` parameters.

* `high_frequency_backup_retention` - (Optional, Int, Available since v1.276.0) Retention period (in days) for high-frequency backups. Before using this parameter, ensure that the BackupInterval parameter has been specified. The default retention period is one day.
* `instance_type` - (Optional, Available since v1.276.0) Instance type. Valid values:
  - replicate 
  - sharding

-> **NOTE:** - Required when the backup restoration type is "rebuild from a deleted instance."

-> **NOTE:** - Required when the backup restoration type is "clone from a remote backup."


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `log_backup_retention_period` - (Optional, Computed, Int, Available since v1.230.1) The number of days to retain log backups. The default value is 7 days.
Valid values: 7 to 730.
* `maintain_end_time` - (Optional, Computed, Available since v1.56.0) The end time of the maintenance window for the instance, in the format HH:mmZ (UTC time).  

-> **NOTE:**  The end time of the maintenance window must be later than the start time.  

* `maintain_start_time` - (Optional, Computed, Available since v1.56.0) The start time of the maintenance window for the instance, in the format HH:mmZ (UTC time).
* `mongos_list` - (Optional, ForceNew, List, Available since v1.276.0) Information about Mongos nodes. See [`mongos_list`](#mongos_list) below.
* `network_type` - (Optional, ForceNew, Computed) The network type of the instance. Valid values:  
  - `Classic`: Classic network  
  - `VPC`: VPC network  
* `order_type` - (Optional, Available since v1.134.0) The configuration change type. Valid values:  
  - `UPGRADE`: Default value, indicating an upgrade.  
  - `DOWNGRADE`: Downgrade.  

-> **NOTE:**  You can configure this parameter when the instance uses the subscription billing method.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `parameters` - (Optional, Computed, List, Available since v1.203.0) The parameters to be modified and their corresponding values, formatted as a JSON string. Example: {"ParameterName1":"ParameterValue1","ParameterName2":"ParameterValue2"}.

-> **NOTE:** You can call the [DescribeParameterTemplates](https://help.aliyun.com/document_detail/67618.html) operation to query the list of default parameter templates.
 See [`parameters`](#parameters) below.
* `payment_type` - (Optional, Computed, Available since v1.276.0) The billing method of the instance. Valid values:  
  - `PrePaid`: Subscription (monthly or yearly)  
  - `PostPaid`: Pay-as-you-go  
* `period` - (Optional, Computed, Int) Subscription duration, in months. Valid values: `1` to `9`, `12`, `24`, and `36`.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `preferred_backup_period` - (Optional, Available since v1.276.0) The backup cycle. Valid values:  
  - `Monday`: Monday  
  - `Tuesday`: Tuesday  
  - `Wednesday`: Wednesday  
  - `Thursday`: Thursday  
  - `Friday`: Friday  
  - `Saturday`: Saturday  
  - `Sunday`: Sunday  

-> **NOTE:** For data security, back up your MongoDB instance at least twice a week.  

-> **NOTE:**  To specify multiple values, separate them with commas (,).  

* `preferred_backup_time` - (Optional, Computed, Available since v1.276.0) The time window for performing backups, in the format HH:mmZ-HH:mmZ (UTC time).

-> **NOTE:**  The time window must be exactly 1 hour.

* `preserve_one_each_hour` - (Optional, Available since v1.276.0) Specifies whether to enable hourly sparse backup.  
  - true: When the backup frequency is set to minutes, all snapshots within one hour from the current time are retained. For snapshots older than one hour but within 24 hours, only the first snapshot after each full hour is retained.  
  - false: All snapshots within the retention period for high-frequency backups are retained.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `pricing_cycle` - (Optional, Available since v1.276.0) Billing cycle unit for the instance.  
Valid values:  
  - **Month:** Month  
  - **Year:** Year  

Default value: Month.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `protocol_type` - (Optional, ForceNew, Computed, Available since v1.276.0) The protocol type for access. Valid values:
  - `mongodb`: MongoDB protocol
  - `dynamodb`: DynamoDB protocol
* `provisioned_iops` - (Optional, Int, Available since v1.229.0) Provisioned performance (IOPS). Valid values: 0 to 50,000.
* `readonly_replicas` - (Optional, Computed, Available since v1.199.0) The number of read-only nodes. Valid values: `0` to `5`.  

If the instance has only **Classic Network** and `VPC` enabled, you must enable public network access or release the Classic Network endpoint before you can modify the number of read-only nodes.  

-> **NOTE:**  You can log on to the ApsaraDB for MongoDB console and view the enabled network types on the **Database Connection** page.  

* `replication_factor` - (Optional, Computed, Int) The number of nodes in the instance.  

-> **NOTE:**  This parameter is returned only for replica set instances.  

* `resource_group_id` - (Optional, Computed, Available since v1.161.0) Resource group ID. For more information about resource groups, see [View basic resource group information](https://help.aliyun.com/document_detail/151181.html).  

-> **NOTE:**  This parameter is supported only on the China site.

* `restore_time` - (Optional, Available since v1.271.0) Select the point in time to which you want to restore. You can specify any point in time within the last 7 days. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).  

-> **NOTE:**  You must configure this parameter only when you call this operation to clone an instance by point in time. You must also configure the `SrcDBInstanceId` parameter.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `restore_type` - (Optional, Available since v1.276.0) Instance restoration from backup.
  - 1: Restore the instance to a specified point in time.
  - 2: Restore a released instance to a specified backup set.
  - 3: Restore the instance to a specified remote backup set.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `role_arn` - (Optional, Computed, Available since v1.212.0) The ARN of the specified role. Format: `acs:ram::$accountID:role/$roleName`.

-> **NOTE:**  - `$accountID`: The Alibaba Cloud account ID. You can view it by logging into the Alibaba Cloud console, hovering over your avatar in the upper-right corner, and clicking Security Settings.

-> **NOTE:**  - `$roleName`: The RAM role name. You can view it by logging into the RAM console, clicking RAM Roles in the left-side navigation pane, and checking the list of RAM role names.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `secondary_zone_id` - (Optional, ForceNew, Available since v1.199.0) The zone where the secondary node is deployed to enable multi-zone deployment. Valid values:
  - **cn-hangzhou-g**: Hangzhou Zone G
  - **cn-hangzhou-h**: Hangzhou Zone H
  - **cn-hangzhou-i**: Hangzhou Zone I
  - **cn-hongkong-b**: Hong Kong Zone B
  - **cn-hongkong-c**: Hong Kong Zone C
  - **cn-hongkong-d**: Hong Kong Zone D
  - **cn-wulanchabu-a**: Ulanqab Zone A
  - **cn-wulanchabu-b**: Ulanqab Zone B
  - **cn-wulanchabu-c**: Ulanqab Zone C
  - **ap-southeast-1a**: Singapore Zone A
  - **ap-southeast-1b**: Singapore Zone B
  - **ap-southeast-1c**: Singapore Zone C
  - **ap-southeast-5a**: Jakarta Zone A
  - **ap-southeast-5b**: Jakarta Zone B
  - **ap-southeast-5c**: Jakarta Zone C
  - **eu-central-1a**: Frankfurt Zone A
  - **eu-central-1b**: Frankfurt Zone B
  - **eu-central-1c**: Frankfurt Zone C

-> **NOTE:**  - This parameter is available only for instances with cloud disks.

-> **NOTE:**  - The value of this parameter must be different from the values of `ZoneId` and `HiddenZoneId`.

* `security_group_id` - (Optional, Available since v1.73.0) ECS security group ID.

-> **NOTE:**  - A single instance can be associated with up to 10 ECS security groups.

-> **NOTE:**  - You can call the ECS [DescribeSecurityGroups](https://help.aliyun.com/document_detail/25556.html) operation to query security group IDs in the target region.

* `security_ips` - (Optional, Available since v1.276.0) The list of IP addresses in the IP whitelist group. Separate multiple IP addresses with commas (,). Duplicate entries are not allowed, and up to 1,000 IP addresses can be added. The following formats are supported:  
  - IP address format, for example: 127.0.0.1.  
  - CIDR notation, for example: 127.0.0.1/24 (Classless Inter-Domain Routing; 24 indicates the prefix length, which ranges from 1 to 32).  
* `shard_list` - (Optional, ForceNew, List, Available since v1.276.0) Information about shard nodes. See [`shard_list`](#shard_list) below.
* `snapshot_backup_type` - (Optional, Computed, Available since v1.212.0) The snapshot backup type. Valid values:
  - `Flash`: Second-level backup.
  - `Standard`: Standard backup (default).
* `src_db_instance_id` - (Optional, Available since v1.271.0) The ID of the source instance.

-> **NOTE:**  When you call this operation to clone an instance, you must specify this parameter and also configure either the `BackupId` or `RestoreTime` parameter. When you call this operation to restore an instance from the recycle bin, you only need to specify this parameter; you do not need to configure `BackupId` or `RestoreTime`.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `src_region` - (Optional, Available since v1.276.0) The region where the instance is located.

-> **NOTE:**  

-> **NOTE:**  - Required when restoring a deleted instance.

-> **NOTE:**  - Required for remote backup.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `ssl_status` - (Optional, Computed) The operation to perform on the SSL feature. Valid values:  
  - `Open`: Enable SSL encryption.  
  - `Close`: Disable SSL encryption.  
  - `Update`: Update the SSL certificate.
* `storage_engine` - (Optional, ForceNew, Computed) Storage engine of the instance. The value is fixed to `WiredTiger`.  

-> **NOTE:**  - When you call this operation to clone an instance or restore an instance from the recycle bin, this parameter must be the same as that of the source instance.  

-> **NOTE:**  - For information about the constraints on storage engines and versions, see [Versions and storage engines](https://help.aliyun.com/document_detail/61906.html).

* `storage_type` - (Optional, Computed, Available since v1.199.0) Storage type. Valid values:
  - `cloud_essd1`: ESSD PL1 cloud disk.
  - `cloud_essd2`: ESSD PL2 cloud disk.
  - `cloud_essd3`: ESSD PL3 cloud disk.
  - `local_ssd`: SSD local disk.

-> **NOTE:**  - Instances of version 4.4 or later support only cloud disks. If this parameter is not specified, `cloud_essd1` is used by default.

-> **NOTE:**  - Instances of version 4.2 or earlier support only local disks. If this parameter is not specified, `local_ssd` is used by default.

* `switch_mode` - (Optional, Int, Available since v1.276.0) The version upgrade mode for the instance. Valid values:
  - `0`: Upgrade immediately.
  - `1`: Upgrade during the maintenance window.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `tag` - (Optional, List, Available since v1.276.0) This property does not have a description in the spec, please add it before generating code. See [`tag`](#tag) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `tags` - (Optional, ForceNew, Map) User-defined tags.
* `tde_status` - (Optional, Computed, Available since v1.73.0) TDE status. Valid values:
  - `enabled`: TDE is enabled.
  - `disabled`: TDE is disabled.

-> **NOTE:**  When TDE is disabled, the `RoleARN`, `EncryptionKey`, and `EncryptorName` parameters are not returned.

* `vswitch_id` - (Required) The ID of the vSwitch in the destination zone for migration.

-> **NOTE:**  This parameter is available and required only when the network type of the instance is Virtual Private Cloud (VPC).

* `vpc_auth_mode` - (Optional, ForceNew, Available since v1.276.0) Indicates whether password-free access over Virtual Private Cloud (VPC) is enabled. Valid values:
  - `Open`: Password-free access over VPC is enabled.
  - `Close`: Password-free access over VPC is disabled.
* `vpc_id` - (Required, ForceNew) The ID of the Virtual Private Cloud (VPC).
* `zone_id` - (Optional, Computed) The ID of the destination zone for migration.

-> **NOTE:**  - The destination zone must be in the same region as the current instance.

-> **NOTE:**  - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query zone IDs.


### `configserver_list`

The configserver_list supports the following:
* `node_class` - (Required, ForceNew, Available since v1.276.0) Specification of the ConfigServer node.
* `node_storage` - (Required, ForceNew, Int, Available since v1.276.0) Storage capacity of the Configserver node.  

### `mongos_list`

The mongos_list supports the following:
* `node_class` - (Required, ForceNew, Available since v1.276.0) Specification of the Mongos node.

### `parameters`

The parameters supports the following:
* `parameter_name` - (Optional, Available since v1.276.0) Parameter name.
* `parameter_value` - (Optional, Available since v1.276.0) Parameter value.

### `shard_list`

The shard_list supports the following:
* `node_class` - (Required, ForceNew, Available since v1.276.0) Specification of the shard node.
* `node_storage` - (Required, ForceNew, Int, Available since v1.276.0) Storage capacity of the shard node.
* `readonly_replicas` - (Optional, ForceNew, Int, Available since v1.276.0) Number of read-only replicas in the shard node. Valid values: `0` to `5`.

-> **NOTE:**  This parameter is currently supported only on the China site.


### `tag`

The tag supports the following:
* `key` - (Optional, ForceNew, Available since v1.276.0) This property does not have a description in the spec, please add it before generating code.
* `value` - (Optional, ForceNew, Available since v1.276.0) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `configserver_list` - Information about ConfigServer nodes.
  * `connect_string` - Connection address of the Configserver node.
  * `max_connections` - Maximum number of connections for the Configserver node.
  * `max_iops` - Maximum IOPS of the Configserver node.
  * `node_description` - Name of the ConfigServer node.
  * `node_id` - ID of the Configserver node.
  * `port` - Port of the Configserver node.
  * `status` - Status of the ConfigServer node.
* `create_time` - Creation time.
* `mongos_list` - Information about Mongos nodes.
  * `connect_sting` - Connection address of the Mongos node.
  * `max_connections` - Maximum number of connections for the Mongos node.
  * `max_iops` - Maximum IOPS of the Mongos node.
  * `node_description` - Mongos node name.
  * `node_id` - Mongos node ID.
  * `port` - Connection port of the Mongos node.
  * `vpc_id` - VPC ID of the Mongos node.
  * `vswitch_id` - VSwitch ID of the Mongos node.
  * `vpc_cloud_instance_id` - VPC instance ID of the Mongos node.
* `region_id` - The region ID of the instance.
* `security_ip_list` - Whitelist groups used for database access control.
  * `security_ip_group_attribute` - Attribute of the IP whitelist group.
  * `security_ip_group_name` - Group name.
  * `security_ips` - IP addresses included in the default whitelist group.
* `shard_list` - Information about shard nodes.
  * `connect_string` - Connection address of the shard node.
  * `max_connections` - Maximum number of connections for the shard node.
  * `max_iops` - Maximum IOPS of the shard node.
  * `node_description` - Name of the shard node.
  * `node_id` - ID of the shard node.
  * `port` - Port of the shard node.
* `status` - The status of the instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 120 mins) Used when create the Instance.
* `delete` - (Defaults to 17 mins) Used when delete the Instance.
* `update` - (Defaults to 73 mins) Used when update the Instance.

## Import

Mongodb Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_mongodb_instance.example <instance_id>
```