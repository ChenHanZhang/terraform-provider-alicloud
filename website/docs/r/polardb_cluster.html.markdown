---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_cluster"
description: |-
  Provides a Alicloud Polardb Db Cluster resource.
---

# alicloud_polardb_cluster

Provides a Polardb Db Cluster resource.

POLARDB uses a cluster architecture. A cluster contains a master node and multiple read nodes.

For information about Polardb Db Cluster and how to use it, see [What is Db Cluster](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateDBCluster).

-> **NOTE:** Available since v1.66.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_polardb_node_classes" "default" {
  db_type    = "MySQL"
  db_version = "8.0"
  category   = "Normal"
  pay_type   = "PostPaid"
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

  db_cluster_ip_array {
    db_cluster_ip_array_name = "default"
    security_ips             = ["1.2.3.4", "1.2.3.5"]
  }
  db_cluster_ip_array {
    db_cluster_ip_array_name = "default2"
    security_ips             = ["1.2.3.6"]
  }
}
```

## Argument Reference

The following arguments are supported:
* `allow_shut_down` - (Optional, ForceNew, Available since v1.204.0) Specifies whether to enable auto-pause when inactive. Valid values:
  - `true`: Enables auto-pause.
  - `false`: Disables auto-pause (default).

-> **NOTE:**  This parameter is supported only for Serverless clusters.

* `architecture` - (Optional, ForceNew, Available since v1.271.0) The CPU architecture. Valid values are as follows:  
  - `X86`  
  - `ARM`.  
* `auto_renew` - (Optional, Available since v1.271.0) Specifies whether auto-renewal is enabled. Valid values:
  - `true`: Auto-renewal is enabled.
  - `false`: Auto-renewal is disabled.

Default value: `false`.

-> **NOTE:**  This parameter takes effect only when the `PayType` parameter is set to `Prepaid`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `backup_retention_policy_on_cluster_deletion` - (Optional, Computed) The backup retention policy when the cluster is deleted. Valid values:
* `ALL`: Retain all backups permanently.
* `LATEST`: Retain only the last backup (an automatic backup is created before deletion).
* `NONE`: Do not retain any backups when the cluster is deleted.

By default, this parameter is set to `NONE` when a cluster is created, which means no backups are retained upon cluster deletion.

-> **NOTE:**  This parameter takes effect only when `DBType` is `MySQL`.

-> **NOTE:**  This parameter is not supported for Serverless clusters.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `category` - (Optional, ForceNew, Computed, Available since v1.271.0) The database edition. Valid values:
* `Normal`: Cluster Edition (default)
* `Basic`: Single-node Edition
* `ArchiveNormal`: High-compression engine (X-Engine)
* `NormalMultimaster`: Multi-master Cluster Edition
* `SENormal`: Standard Edition

-> **NOTE:**  * `Basic` is supported by **MySQL 5.6**, **5.7**, **8.0**, **PostgreSQL 14**, and **Oracle Syntax Compatible Edition 2.0**.

-> **NOTE:**  * `ArchiveNormal` and `NormalMultimaster` are supported by **MySQL 8.0**.

-> **NOTE:**  * `SENormal` is supported by **MySQL 5.6**, **5.7**, **8.0**, and **PostgreSQL 14**.

For more information about database editions, see [Database Editions](https://help.aliyun.com/document_detail/183258.html).
* `clone_data_point` - (Optional) The point in time from which data is cloned. Valid values:
  - `LATEST`: Data at the latest point in time.
  - `BackupID`: A historical backup set ID. Specify an actual backup set ID.
  - `Timestamp`: A historical point in time. Specify an exact time in the format `YYYY-MM-DDThh:mm:ssZ` (UTC time).

Default value: `LATEST`.

-> **NOTE:**  If `CreationOption` is set to `CloneFromRDS`, this parameter can only be set to `LATEST`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `cluster_network_type` - (Optional, ForceNew, Available since v1.271.0) The network type of the cluster.
* `collector_status` - (Optional, Computed) Specifies whether to enable or disable SQL collection. Valid values:
  - Enable: enables SQL collection.
  - Disabled: disables SQL collection.
* `creation_option` - (Optional, Computed) The creation method. Valid values are as follows:

* `Normal`: Creates a brand-new PolarDB cluster. For console operations, see the following documentation:

    * [Create a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/58769.html)
    * [Create a PolarDB for PostgreSQL cluster](https://help.aliyun.com/document_detail/118063.html)
    * [Create a PolarDB for PostgreSQL (compatible with Oracle) cluster](https://help.aliyun.com/document_detail/118182.html)

* `CloneFromPolarDB`: Clones data from an existing PolarDB cluster to a new PolarDB cluster. For console operations, see the following documentation:

    * [Clone a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/87966.html)
    * [Clone a PolarDB for PostgreSQL cluster](https://help.aliyun.com/document_detail/118108.html)
    * [Clone a PolarDB for PostgreSQL (compatible with Oracle) cluster](https://help.aliyun.com/document_detail/118221.html)

* `RecoverFromRecyclebin`: Restores data from a released PolarDB cluster to a new PolarDB cluster. For console operations, see the following documentation:

    * [Restore a released PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/164880.html)
    * [Restore a released PolarDB for PostgreSQL cluster](https://help.aliyun.com/document_detail/432844.html)
    * [Restore a released PolarDB for PostgreSQL (compatible with Oracle) cluster](https://help.aliyun.com/document_detail/424632.html)

* `CloneFromRDS`: Clones data from an existing RDS instance to a new PolarDB cluster. For console operations, see [One-click cloning of RDS for MySQL to PolarDB for MySQL](https://help.aliyun.com/document_detail/121812.html).

* `MigrationFromRDS`: Migrates data from an existing RDS instance to a new PolarDB cluster. The created PolarDB cluster is in read-only mode and has binary logging enabled by default. For console operations, see [One-click upgrade from RDS for MySQL to PolarDB for MySQL](https://help.aliyun.com/document_detail/121582.html).

* `CreateGdnStandby`: Creates a secondary cluster. For console operations, see [Add a secondary cluster](https://help.aliyun.com/document_detail/160381.html).

* `UpgradeFromPolarDB`: Upgrades and migrates from an existing PolarDB cluster. For console operations, see [Major version upgrade](https://help.aliyun.com/document_detail/459712.html).

The default value is `Normal`.

-> **NOTE:**  When `DBType` is `MySQL` and `DBVersion` is **8.0**, this parameter can be set to `CreateGdnStandby`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_node_num` - (Optional, Int) The number of nodes for Standard Edition and Enterprise Edition. Valid values:
  - Standard Edition: 1 to 8 (supports 1 read-write node and up to 7 read-only nodes).
  - Enterprise Edition: 1 to 16 (supports 1 read-write node and up to 15 read-only nodes).

-> **NOTE:**  - Enterprise Edition has 2 nodes by default, while Standard Edition has 1 node by default.

-> **NOTE:**  - This parameter is supported only for PolarDB for MySQL.

-> **NOTE:**  - Changing the number of nodes in a multi-primary cluster is currently not supported.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_cluster_ip_array_attribute` - (Optional, Available since v1.271.0) IP whitelist group attribute. When set to `hidden`, the group is not visible in the console.

-> **NOTE:**  - IP whitelist groups already displayed in the console cannot be hidden.

-> **NOTE:**  - This parameter can be configured only when `WhiteListType` is set to `IP`.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_cluster_ip_array_name` - (Optional, Available since v1.271.0) The name of the IP address whitelist group. The name must be 2 to 120 characters in length, contain only lowercase letters and digits, start with a letter, and end with a letter or digit.
  - If the specified whitelist group name does not exist, the system creates a new whitelist group.
  - If the specified whitelist group name already exists, the system modifies the existing whitelist group.
  - If this parameter is not specified, the system modifies the default group.

-> **NOTE:**  - A cluster supports up to 50 IP address whitelist groups.

-> **NOTE:**  - This parameter is configurable only when `WhiteListType` is set to `IP`.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_minor_version` - (Optional, ForceNew, Computed) Database engine minor version number. Valid values:
  - **8.0.2**
  - **8.0.1**

-> **NOTE:**  This parameter takes effect only when the `DBType` parameter is `MySQL` and the `DBVersion` parameter is **8.0**.

* `db_node_class` - (Required) The node specification. For more information, see the following documentation:
  - PolarDB for MySQL: [Compute Node Specifications](https://help.aliyun.com/document_detail/102542.html)
  - PolarDB for PostgreSQL (compatible with Oracle): [Compute Node Specifications](https://help.aliyun.com/document_detail/207921.html)
  - PolarDB for PostgreSQL: [Compute Node Specifications](https://help.aliyun.com/document_detail/209380.html)

-> **NOTE:**  - To create a Serverless cluster for PolarDB for MySQL Cluster Edition, specify **polar.mysql.sl.small**.

-> **NOTE:**  - To create a Serverless cluster for PolarDB for MySQL Standard Edition, specify **polar.mysql.sl.small.c**.

-> **NOTE:**  - To create a Serverless cluster for PolarDB for PostgreSQL Cluster Edition, specify **polar.pg.sl.small**.

-> **NOTE:**  - To create a Serverless cluster for PolarDB for PostgreSQL Standard Edition, specify **polar.pg.sl.small.c**.

-> **NOTE:**  - To create a Serverless cluster for PolarDB for PostgreSQL (compatible with Oracle), specify **polar.o.sl.small**.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_node_type` - (Optional, Available since v1.271.0) Node type. When modifying the AI node specifications, this value must be DLNode.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_type` - (Required, ForceNew) The database engine type. Valid values are as follows:
  - `MySQL`
  - `PostgreSQL`
  - `Oracle`
* `db_version` - (Required, ForceNew) The major version number of the database engine. Valid values are as follows:  
  - **8.0**  
  - **5.7**  
  - **5.6**.  
* `default_time_zone` - (Optional, Computed) The time zone of the cluster (UTC). Valid values are all full-hour offsets from **-12:00** to **+13:00**, such as **00:00**. The default value is `SYSTEM`, which means the time zone matches the region's time zone.

-> **NOTE:**  This parameter takes effect only when `DBType` is set to `MySQL`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `description` - (Optional, Computed) The description of the cluster.
* `duration` - (Optional, Available since v1.271.0) The auto-renewal duration for the instance. Valid values are as follows:  
  - When `PeriodUnit` is `Month`, valid values are `[1, 2, 3, 6, 12]`.  
  - When `PeriodUnit` is `Year`, valid values are `[1–3]`.  

The default value is `1`.  
* `encrypt_new_tables` - (Optional) Specifies whether to automatically encrypt all newly created tables. Valid values:
  - `ON`: Enabled
  - `OFF`: Disabled

-> **NOTE:**  This parameter takes effect only when the database engine is MySQL-compatible.

* `encryption_key` - (Optional) The ID of the custom encryption key.
* `existed_endpoint_switch_type` - (Optional, Available since v1.271.0) Specifies whether to switch existing endpoints. Valid values:
  - `NONE`: Do not switch existing endpoints.
  - `ALL`: Switch existing endpoints.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `from_time_service` - (Optional, Available since v1.208.1) Specifies whether to immediately or schedule the zone change operation. Valid values:  
  - false (default): Schedule the operation.  
  - true: Execute immediately.  

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `gdn_id` - (Optional) The global database network ID.

-> **NOTE:**  This parameter is required when `CreationOption` is set to `CreateGdnStandby`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `hot_standby_cluster` - (Optional, ForceNew, Computed, Available since v1.203.0) Specifies whether to enable the hot standby cluster. Valid values:
  - `ON` (default): Enables storage hot standby.
  - `OFF`: Disables hot standby.
  - `STANDBY`: Enables hot standby cluster.
  - `EQUAL`: Enables both storage and compute hot standby.
  - `3AZ`: Ensures strong data consistency across multiple zones.

-> **NOTE:**  `STANDBY` applies only to PolarDB for PostgreSQL.

* `is_switch_over_for_disaster` - (Optional, Available since v1.271.0) Specifies whether to switch back to the original zone. Valid values:
  - true: Switches back to the original zone.
  - false: Does not switch back to the original zone.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `loose_polar_log_bin` - (Optional, Computed) Enable the Binlog feature. Valid values are as follows:
  - `ON`: Enables the Binlog feature for the cluster.
  - `OFF`: Disables the Binlog feature for the cluster.

-> **NOTE:**  This parameter takes effect only when `DBType` is set to `MySQL`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `loose_xengine` - (Optional, Computed, Available since v1.232.0) Enables the X-Engine storage engine. Valid values:
  - `ON`: The cluster enables the X-Engine.
  - `OFF`: The cluster disables the X-Engine.

-> **NOTE:**  This parameter takes effect only when `CreationOption` is not `CreateGdnStandby`, `DBType` is `MySQL`, and `DBVersion` is **8.0**. The memory specification of nodes with X-Engine enabled must be at least 8 GB.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `loose_xengine_use_memory_pct` - (Optional, Computed, Available since v1.232.0) Specifies the percentage of memory allocated to the X-Engine storage engine. Valid values are integers from 10 to 90.

-> **NOTE:**  This parameter takes effect only when the `LooseXEngine` parameter is set to `ON`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `lower_case_table_names` - (Optional, Computed) Specifies whether table names are case-sensitive. Valid values:
* `1`: Case-insensitive
* `0`: Case-sensitive

Default value: `1`.

-> **NOTE:**  This parameter takes effect only when `DBType` is `MySQL`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `maintain_time` - (Optional, Computed) Maintenance window for the cluster, in the format `HH:mmZ-HH:mmZ`. For example, `16:00Z-17:00Z` indicates that routine maintenance can be performed from 00:00 to 01:00 (UTC+08:00).

-> **NOTE:**  The maintenance window must start on the hour and last for one hour.

* `modify_mode` - (Optional, Available since v1.271.0) The method for modifying the IP address whitelist. Valid values:
  - `Cover`: Overwrites the existing IP address whitelist (default).
  - `Append`: Adds new IP addresses.
  - `Delete`: Removes specified IP addresses.

-> **NOTE:**  This parameter can be configured only when `WhiteListType` is set to `IP`.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `modify_type` - (Optional) Modification type. Valid values are as follows:  
  - `Upgrade`: Upgrade specifications  
  - `Downgrade`: Downgrade specifications

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `parameter_group_id` - (Optional) The parameter group ID.

-> **NOTE:**  You can call the [DescribeParameterGroups](https://help.aliyun.com/document_detail/207178.html) operation to view the list of parameter groups in the target region, including their IDs.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `payment_type` - (Required, Available since v1.271.0) The billing method of the instance. Valid values:  
  - `Postpaid`: Pay-as-you-go (postpaid)  
  - `Prepaid`: Subscription (prepaid)
* `period` - (Optional) Specifies whether a subscription instance is billed on a yearly or monthly basis. Valid values:  
  - `Year`: Yearly subscription  
  - `Month`: Monthly subscription  

-> **NOTE:**  This parameter is required when `PayType` is set to `Prepaid`.  


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `period_unit` - (Optional, Available since v1.271.0) The unit of the renewal duration. Valid values:
  - `Year`: Year
  - `Month`: Month

Default value: `Month`.
* `planned_end_time` - (Optional, Available since v1.208.1) The latest time to start executing the scheduled task. The time must be in the format `YYYY-MM-DDThh:mm:ssZ` (UTC).

-> **NOTE:**  * The latest time must be at least 30 minutes later than the start time.

-> **NOTE:**  * If `PlannedStartTime` is specified but this parameter is not, the default latest start time for the scheduled task is `start time + 30 minutes`. For example, if `PlannedStartTime` is set to `2021-01-14T09:00:00Z` and this parameter is left empty, the scheduled task will start no later than `2021-01-14T09:30:00Z`.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `planned_flashing_off_time` - (Optional, Available since v1.271.0) The scheduled start time of the flash disconnection. The format is YYYY-MM-DDThh:mm:ssZ (UTC).

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `planned_start_time` - (Optional, Available since v1.208.1) The earliest time to start executing a scheduled scale-up or scale-down task within the specified time window. The time format is `YYYY-MM-DDThh:mm:ssZ` (UTC).

-> **NOTE:**  * This parameter takes effect only when `ModifyType` is set to `Upgrade` or `Downgrade`.

-> **NOTE:**  * The start time must be within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, the valid start time range is from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.

-> **NOTE:**  * If this parameter is left empty, the upgrade task is executed immediately.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `provisioned_iops` - (Optional, Computed, Int, Available since v1.229.1) The provisioned read/write IOPS for ESSD AutoPL cloud disks. Valid values: 0 to min{50,000, 1,000 × capacity − baseline performance}.
Baseline performance = min{1,800 + 50 × capacity, 50,000}.
This parameter is supported only when StorageType is set to ESSDAUTOPL..
* `proxy_class` - (Optional) The specification of the Standard Edition database proxy. Valid values:
  - **polar.maxscale.g2.medium.c**: 2 cores.
  - **polar.maxscale.g2.large.c**: 4 cores.
  - **polar.maxscale.g2.xlarge.c**: 8 cores.
  - **polar.maxscale.g2.2xlarge.c**: 16 cores.
  - **polar.maxscale.g2.3xlarge.c**: 24 cores.
  - **polar.maxscale.g2.4xlarge.c**: 32 cores.
  - **polar.maxscale.g2.8xlarge.c**: 64 cores.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `proxy_type` - (Optional, ForceNew) The database proxy type. Valid values:
  - `EXCLUSIVE`: Enterprise Exclusive Edition
  - `GENERAL`: Enterprise General Edition

-> **NOTE:**  The proxy type must match the type corresponding to the cluster's node specifications. Specifically:

-> **NOTE:**  - If the node specification is a general-purpose specification, the proxy type must be Enterprise General Edition.

-> **NOTE:**  - If the node specification is an exclusive specification, the proxy type must be Enterprise Exclusive Edition.

* `renewal_status` - (Optional) The auto-renewal status. Valid values are as follows:  
  - `AutoRenewal`: Auto-renewal is enabled.  
  - `Normal`: Manual renewal is required.  
  - `NotRenewal`: The instance will not be renewed.  

The default value is `AutoRenewal`.  

-> **NOTE:**  If this parameter is set to `NotRenewal`, the system stops sending expiration reminders and sends only a non-renewal reminder three days before expiration.  

* `resource_group_id` - (Optional, Computed) The ID of the resource group after the change.
* `role_arn` - (Optional, Computed) The Alibaba Cloud Resource Name (ARN) of the role used to specify a particular role. For more information, see [Overview of RAM roles](https://help.aliyun.com/document_detail/93689.html).

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `scale_max` - (Optional, ForceNew, Available since v1.204.0) Maximum scaling capacity per node. Valid values: 1 PCU to 32 PCU.

-> **NOTE:**  This parameter is supported only for Serverless clusters.

* `scale_min` - (Optional, ForceNew, Available since v1.204.0) Minimum scaling limit per node. Valid values: 1 PCU to 31 PCU.

-> **NOTE:**  This parameter is supported only for Serverless clusters.

* `scale_ro_num_max` - (Optional, ForceNew, Available since v1.204.0) Maximum number of read-only nodes allowed for scaling. Valid values: 0 to 15.

-> **NOTE:**  This parameter is supported only for Serverless clusters.

* `scale_ro_num_min` - (Optional, ForceNew, Available since v1.204.0) The minimum number of read-only nodes for scaling. Valid values: 0 to 15.  

-> **NOTE:**  This parameter is supported only for Serverless clusters.

* `security_group_ids` - (Optional, Computed) The IDs of security groups, separated by commas (,).  

-> **NOTE:**  - A cluster can be associated with up to three security groups.  

-> **NOTE:**  - This parameter is configurable only when `WhiteListType` is set to `SecurityGroup`.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `security_ips` - (Optional, Computed) IP addresses or CIDR blocks in the IP whitelist group. A maximum of 1,000 IP addresses or CIDR blocks are supported across all IP whitelist groups. Separate multiple IP addresses with commas (,). The following formats are supported:
  - IP address format, for example: 10.23.12.24.
  - CIDR format, for example: 10.23.12.24/24 (Classless Inter-Domain Routing; 24 indicates the prefix length, which can range from 1 to 32).

-> **NOTE:**  This parameter is configurable only when `WhiteListType` is set to `IP`.

* `serverless_type` - (Optional, ForceNew, Available since v1.204.0) The Serverless type. The current value is fixed as `AgileServerless` (Agile).

-> **NOTE:**  This parameter is supported only for Serverless clusters.

* `source_resource_id` - (Optional) The source ApsaraDB RDS instance ID or source PolarDB cluster ID. This parameter is required only when `CreationOption` is set to `MigrationFromRDS`, `CloneFromRDS`, `CloneFromPolarDB`, or `RecoverFromRecyclebin`.  
  - If `CreationOption` is set to `MigrationFromRDS` or `CloneFromRDS`, you must specify the source ApsaraDB RDS instance ID for this parameter. The source RDS instance must be a High Availability Edition of ApsaraDB RDS for MySQL 5.6, 5.7, or 8.0.  

  - If `CreationOption` is set to `CloneFromPolarDB`, you must specify the source PolarDB cluster ID for this parameter. By default, the cloned cluster uses the same `DBType` as the source cluster. For example, if the source cluster runs MySQL 8.0, you must set `DBType` to `MySQL` and `DBVersion` to **8.0** for the cloned cluster.  
  - If `CreationOption` is set to `RecoverFromRecyclebin`, you must specify the ID of a released source PolarDB cluster for this parameter. The recovered cluster must use the same `DBType` as the source cluster. For example, if the source cluster runs MySQL 8.0, you must set `DBType` to `MySQL` and `DBVersion` to **8.0** for the recovered cluster.  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `standby_az` - (Optional, Computed) The zone where the hot standby cluster is deployed.

-> **NOTE:**  This parameter takes effect only when the hot standby cluster or multi-zone strong data consistency is enabled.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `storage_auto_scale` - (Optional, Available since v1.271.0) Specifies whether to enable automatic storage scaling for Standard Edition clusters. Valid values:
  - Enable: Enables automatic storage scaling.
  - Disable: Disables automatic storage scaling.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `storage_pay_type` - (Optional, ForceNew, Computed, Available since v1.210.0) The billing method for storage. Valid values:
  - Postpaid: Pay-as-you-go based on capacity usage.
  - Prepaid: Subscription-based billing based on allocated storage space.
* `storage_space` - (Optional, Computed, Int, Available since v1.203.0) The storage space for subscription (monthly or yearly) billing. Unit: GB.

-> **NOTE:**  - For PolarDB MySQL Enterprise Edition, the valid range is 10 to 50,000.

-> **NOTE:**  - For PolarDB MySQL Standard Edition, the valid range is 20 to 64,000.

-> **NOTE:**  - When the storage type of the Standard Edition is ESSDAUTOPL, the valid range is 40 to 64,000 with a minimum step size of 10 (for example, 40, 50, 60, and so on).

* `storage_type` - (Optional, ForceNew, Computed, Available since v1.203.0) For Enterprise Edition, valid storage types are:
  - `PSL5`
  - `PSL4`

For Standard Edition, valid storage types are:
  - `ESSDPL0`
  - `ESSDPL1`
  - `ESSDPL2`
  - `ESSDPL3`
  - `ESSDAUTOPL`.
* `storage_upper_bound` - (Optional, Int, Available since v1.271.0) The upper limit for automatic storage expansion of Standard Edition clusters. Unit: GB.

-> **NOTE:**  This parameter takes effect only when the `StorageAutoScale` parameter is set to `Enable`.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `strict_consistency` - (Optional, ForceNew, Computed, Available since v1.239.0) Specifies whether multi-zone strong data consistency is enabled for the cluster. Valid values:
  - `ON`: Multi-zone strong data consistency is enabled. This applies to Standard Edition clusters deployed across three availability zones (3AZ).
  - `OFF`: Multi-zone strong data consistency is not enabled.
* `sub_category` - (Optional, Computed) The cluster sub-series. Valid values:
  - `normal_exclusive`: Dedicated specifications
  - `normal_general`: General-purpose specifications

If you switch between dedicated and general-purpose specifications, this field is required.
* `tags` - (Optional, Map) A list of tags.
* `tde_status` - (Optional) Modifies the Transparent Data Encryption (TDE) status. The only valid value is `Enable`.
* `used_time` - (Optional, Available since v1.271.0) Subscription duration. Valid values:
  - When `Period` is set to `Year`, `UsedTime` can be 1 to 3.
  - When `Period` is set to `Month`, `UsedTime` can be 1 to 9.

-> **NOTE:**  This parameter is required when `PayType` is set to `Prepaid`.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `vswitch_id` - (Optional, Computed) The ID of the vSwitch.  

-> **NOTE:**  If VPCId is specified, VSwitchId is required.  

* `vpc_id` - (Optional, Computed, Available since v1.185.0) The VPC ID used when switching zones.
* `white_list_type` - (Optional, Available since v1.271.0) The whitelist type. Valid values:
  - `IP`: IP address whitelist group.
  - `SecurityGroup`: Security group.

Default value: `IP`.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `zone_id` - (Optional, Computed) Zone ID.

-> **NOTE:**  You can view available zones by using the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation.

* `zone_type` - (Optional, Available since v1.271.0) The type of zone to switch to. Valid values:
  - `Primary`: Primary zone.
  - `Standby`: Standby zone.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Cluster creation time.
* `parameters` - The parameters that describe the PolarDB cluster.
  * `checking_code` - The valid value range for the parameter.
  * `data_type` - The data type of the parameter value.
  * `force_restart` - Indicates whether a restart is required for the change to take effect.
  * `is_modifiable` - Indicates whether the parameter is modifiable.
  * `name` - The name of the parameter.
  * `parameter_description` - The description of the parameter.
  * `parameter_status` - The status of the parameter.
  * `value` - The parameter value.
* `region_id` - The region ID.
* `seconds_until_auto_pause` - The idle timeout duration before auto-pause.
* `status` - The task status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 17 mins) Used when create the Db Cluster.
* `delete` - (Defaults to 10 mins) Used when delete the Db Cluster.
* `update` - (Defaults to 29 mins) Used when update the Db Cluster.

## Import

Polardb Db Cluster can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_cluster.example <db_cluster_id>
```