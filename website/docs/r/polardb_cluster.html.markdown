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
* `allow_shut_down` - (Optional, ForceNew, Available since v1.204.0) Whether to turn on inactive pause
* `architecture` - (Optional, ForceNew, Available since v1.270.0) CPU architecture. Value range:
  - X86
  - ARM
* `auto_renew` - (Optional, Available since v1.270.0) Whether to enable auto-renewal, with available values as follows:
  - `true`: Auto-renew.
  - `false`: Do not auto-renew.

The default value is `false`.

-> **NOTE:**  This parameter takes effect only when `PayType` is set to `Prepaid`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `backup_retention_policy_on_cluster_deletion` - (Optional, Computed) When deleting a cluster, the backup set retention policy has the following values:
* `ALL`: Keep ALL backups permanently.
* `LATEST`: Permanently retain the last backup (automatically backup before deletion).
* `NONE`: The backup set is not retained when the cluster is deleted.

When you create a cluster, the default value is `NONE`, that is, the backup set is not retained when the cluster is deleted.

-> **NOTE:**  This parameter takes effect only when `DBType` is **MySQL.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `category` - (Optional, ForceNew, Computed, Available since v1.270.0) Product series, with valid values as follows:
  - `Normal`: Cluster Edition (default)
  - `Basic`: Single-node
  - `ArchiveNormal`: High Compression Engine (X-Engine)
  - `NormalMultimaster`: Multi-master Cluster Edition
  - `SENormal`: Standard Edition
&gt; * `MySQL` **5.6**, **5.7**, **8.0**, `PostgreSQL` `14`, and **Oracle Syntax Compatible 2.0** support `Basic`.
&gt; * `MySQL` **8.0** supports `ArchiveNormal` and `NormalMultimaster`.
&gt; * `MySQL` **5.6**, **5.7**, **8.0**, and `PostgreSQL` `14` support `SENormal`.
For more information about product series, see [Product Series](https://www.alibabacloud.com/help/en/doc-detail/183258.html).
* `clone_data_point` - (Optional) The time node of the cloned data. Valid values:
  - `LATEST`: the LATEST time point.
  - `BackupID`: the ID of the historical backup set. Please enter the ID of the specific backup set.
  - `Timestamp`: historical time point, please pass in the specific time format: 'YYYY-MM-DDThh:mm:ssZ'(UTC time).

The default value is **LATEST * *.

-> **NOTE:**  If `CreationOption` is `CloneFromRDS`, the value of this parameter can only be **LATEST * *.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `cluster_network_type` - (Optional, ForceNew, Available since v1.270.0) DBClusterNetworkType
* `collector_status` - (Optional, Computed) CollectorStatus
* `creation_option` - (Optional, Computed) Creation method, with the following values supported: 
* `Normal`: Creates a brand new PolarDB cluster. For console operations, refer to the following documents: 
    * [Create a PolarDB MySQL Edition Database Cluster](https://www.alibabacloud.com/help/en/doc-detail/58769.html) 
    * [Create a PolarDB PostgreSQL Edition Database Cluster](https://www.alibabacloud.com/help/en/doc-detail/118063.html) 
    * [Create a PolarDB PostgreSQL Edition (Oracle Compatible) Database Cluster](https://www.alibabacloud.com/help/en/doc-detail/118182.html) 
* `CloneFromPolarDB`: Clones data from an existing PolarDB cluster to a new PolarDB cluster. For console operations, refer to the following documents: 
    * [Clone a PolarDB MySQL Edition Cluster](https://www.alibabacloud.com/help/en/doc-detail/87966.html) 
    * [Clone a PolarDB PostgreSQL Edition Cluster](https://www.alibabacloud.com/help/en/doc-detail/118108.html) 
    * [Clone a PolarDB PostgreSQL Edition (Oracle Compatible) Cluster](https://www.alibabacloud.com/help/en/doc-detail/118221.html) 
* `RecoverFromRecyclebin`: Recovers data from a released PolarDB cluster to a new PolarDB cluster. For console operations, refer to the following documents: 
   * [Restore a Released PolarDB MySQL Edition Cluster](https://www.alibabacloud.com/help/en/doc-detail/164880.html) 
   * [Restore a Released PolarDB PostgreSQL Edition Cluster](https://www.alibabacloud.com/help/en/doc-detail/432844.html) 
   * [Restore a Released PolarDB PostgreSQL Edition (Oracle Compatible) Cluster](https://www.alibabacloud.com/help/en/doc-detail/424632.html) 
* `CloneFromRDS`: Clones data from an existing RDS instance to a new PolarDB cluster. Console operation guide is available at [One-click Clone from RDS MySQL to PolarDB MySQL Edition](https://www.alibabacloud.com/help/en/doc-detail/121812.html). 
* `MigrationFromRDS`: Migrates data from an existing RDS instance to a new PolarDB cluster. The created PolarDB cluster operates in read-only mode with Binlog enabled by default. Console operation guide is at [One-click Upgrade from RDS MySQL to PolarDB MySQL Edition](https://www.alibabacloud.com/help/en/doc-detail/121582.html). 
* `CreateGdnStandby`: Creates a standby cluster. Console operation guide can be found at [Add Standby Cluster](https://www.alibabacloud.com/help/en/doc-detail/160381.html). 
* `UpgradeFromPolarDB`: Upgrades and migrates from PolarDB. Console operation guide is detailed in [Major Version Upgrade](https://www.alibabacloud.com/help/en/doc-detail/459712.html). 

The default value is `Normal`. 

-> **NOTE:**  When `DBType` is `MySQL` and `DBVersion` is **8.0**, this parameter can also take the value `CreateGdnStandby`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_node_num` - (Optional, Int) The number of nodes. This parameter is supported for Standard Edition clusters. Valid values:
  - `1` (default): only one primary node.
  - `2`: one read-only node and one primary node.
&gt; * By default, an Enterprise Edition cluster has two nodes and a Standard Edition cluster has one node.
  - This parameter is supported only for PolarDB for MySQL clusters.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_cluster_ip_array_attribute` - (Optional, Available since v1.270.0) IP whitelist grouping attributes. The console is not visible when set to **hidden.

-> **NOTE:** - IP whitelist groups displayed in the console do not support hiding.

-> **NOTE:** - This parameter is supported only when the value of `WhiteListType` is **IP.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_cluster_ip_array_name` - (Optional, Available since v1.270.0) The name of the IP whitelist group. The group name must be 2 to 120 characters in length and consist of lowercase letters and numbers. It starts with a letter and ends with a letter or number.
  - If the name of the whitelist group does not exist, the whitelist group is created.
  - If the name of the incoming whitelist group already exists, the whitelist group will be modified.
  - Modify the default grouping if it is not passed in.

-> **NOTE:**  - 1 cluster supports up to 50 IP whitelist groups.

-> **NOTE:** - This parameter is supported only when the value of `WhiteListType` is **IP.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_minor_version` - (Optional, ForceNew, Computed) Database engine minor version number.
* `db_node_class` - (Required) DBNodeClass

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_node_type` - (Optional, Available since v1.270.0) The node type. The fixed value is DLNode.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_type` - (Required, ForceNew) DBType
* `db_version` - (Required, ForceNew) Database engine version number.

The range of the MySQL version number is as follows:
  - ** 5.6 * * *
  - ** 5.7 * * *
  - ** 8.0 * * *

The value range of PostgreSQL version number is as follows:
  - ** 11 * * *
  - ** 14 * * *
The Oracle version number is 11.
* `default_time_zone` - (Optional, Computed) Cluster timezone (UTC), with selectable values ranging from **-12:00** to **+13:00** at whole-hour intervals, e.g., **00:00**. The default value is `SYSTEM`, which matches the Region&#39;s timezone. 

-> **NOTE:**  This parameter applies only when `DBType` is `MySQL`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `description` - (Optional, Computed) DBClusterDescription
* `duration` - (Optional, Available since v1.270.0) Duration
* `encrypt_new_tables` - (Optional) Whether to enable automatic encryption of all newly created tables. Value:
  - `ON`: ON
  - `OFF`: OFF
* `encryption_key` - (Optional) The ID of the custom key.
* `existed_endpoint_switch_type` - (Optional, Available since v1.270.0) Set existed endpoint switch rule

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `from_time_service` - (Optional, Available since v1.208.1) Perform the replacement of the primary zone immediately or periodically. Value:
  - false (default): timed execution
  - true: Execute immediately

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `gdn_id` - (Optional) The ID of the global database network.

-> **NOTE:**  This parameter is required when `CreationOption` is **CreateGdnStandby.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `hot_standby_cluster` - (Optional, ForceNew, Computed, Available since v1.203.0) Whether to enable the hot standby cluster. Values are as follows: 
  - `ON` (default): Enables the hot standby cluster. 
  - `OFF`: Disables the hot standby cluster. 
  - `STANDBY`: Enables the hot standby cluster for the standard edition. 

-> **NOTE:** The default value for standard edition clusters is `STANDBY`.

* `is_switch_over_for_disaster` - (Optional, Available since v1.270.0) Whether to revert to the original Main Zone. The value range is as follows:
true: restitution of the original primary zone
false: Change the primary zone.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `loose_polar_log_bin` - (Optional, Computed) Enable Binlog feature, valid values are as follows: 
  - `ON`: Cluster enables the Binlog feature. 
  - `OFF`: Cluster disables the Binlog feature. 

-> **NOTE:**  This parameter takes effect only when the `DBType` parameter is set to `MySQL`.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `loose_xengine` - (Optional, Computed, Available since v1.232.0) Enable the X-Engine storage engine function. The value range is as follows:
  - `ON`: Enable the X-Engine engine for the cluster
  - `OFF`: The cluster X-Engine engine is disabled.

-> **NOTE:**  This parameter takes effect only when `CreationOption` is not equal to `CreateGdnStandby`,`DBType` is `MySQL`, and `DBVersion` is **8.0. The memory specification of the node on which the X-Engine engine is enabled must be greater than or equal to 16GB.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `loose_xengine_use_memory_pct` - (Optional, Computed, Available since v1.232.0) Set the ratio of enabled X-Engine storage engines. The value is an integer from 10 to 90.

-> **NOTE:**  This parameter takes effect only when the parameter `LooseXEngine` is set to **ON.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `lower_case_table_names` - (Optional, Computed) Whether the table name is case sensitive, the value range is as follows:
* `1`: Case-insensitive
* `0`: Case Sensitive

The default value is **1 * *.

-> **NOTE:**  This parameter takes effect only when `DBType` is **MySQL.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `maintain_time` - (Optional, Computed) MaintainTime
* `modify_mode` - (Optional, Available since v1.270.0) How to modify the IP address whitelist. The value range is as follows:
  - `Cover`: overwrites the original IP address whitelist (default value).
  - `Append`: appends an IP address.
  - `Delete`: deletes an IP address.

-> **NOTE:**  This parameter is supported only when the value of `WhiteListType` is **IP.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `modify_type` - (Optional) Change type, the range of values is as follows:
  - `Upgrade`: Upgrade specification
  - `Downgrade`: Downgrade specification

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `parameter_group_id` - (Optional) Parameter template ID. 

-> **NOTE:**  You can view the list of parameter templates in the target region, including the parameter template ID, by calling the [DescribeParameterGroups](https://www.alibabacloud.com/help/en/doc-detail/207178.html) interface.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `payment_type` - (Required, Available since v1.270.0) The paymen type of the resource
* `period` - (Optional) If the payment type is `Prepaid`, this parameter is required. It specifies whether the prepaid cluster is on a monthly or yearly basis. 
  - `Year`: Yearly subscription. 
  - `Month`: Monthly subscription.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `period_unit` - (Optional, Available since v1.270.0) PeriodUnit
* `planned_end_time` - (Optional, Available since v1.208.1) The latest time to start the target scheduled task. The format is YYYY-MM-DDThh:mm:ssZ(UTC).

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `planned_flashing_off_time` - (Optional, Available since v1.270.0) Specifies the flash start time. The format is YYYY-MM-DDThh:mm:ssZ(UTC).

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `planned_start_time` - (Optional, Available since v1.208.1) Start execution timing (that is, the earliest time when the upgrade configuration task is executed within the target time period). The format is YYYY-MM-DDThh:mm:ssZ(UTC).

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `provisioned_iops` - (Optional, Computed, Int, Available since v1.229.1) ESSD AutoPL IOPS preconfigured for reading and writing cloud disks. Possible values: 0 ~ min{50,000, 1000 * capacity-baseline performance}. 
 Benchmark performance = min{1,800+50 * capacity, 50000}. 
 This parameter is supported only when StorageType is set to ESSDAUTOPL. 
* `proxy_class` - (Optional) Database agent specification

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `proxy_type` - (Optional, ForceNew) Database agent
* `renewal_status` - (Optional) RenewalStatus
* `resource_group_id` - (Optional, Computed) The ID of the resource group.
* `role_arn` - (Optional, Computed) A global resource descriptor for a role that specifies a specific role. For more information, see [RAM role Overview](~~ 93689 ~~).

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `scale_max` - (Optional, ForceNew, Available since v1.204.0) Maximum scaling limit for a single node. The value range is: 1 PCU~32 PCU. 
&gt; Only supported by Serverless clusters.
* `scale_min` - (Optional, ForceNew, Available since v1.204.0) Lower limit of single node scaling
* `scale_ro_num_max` - (Optional, ForceNew, Available since v1.204.0) The upper limit of the number of read-only nodes.
* `scale_ro_num_min` - (Optional, ForceNew, Available since v1.204.0) Lower limit for number of read-only nodes
* `security_group_ids` - (Optional, Computed) The ID of the security group. Separate multiple security groups with commas (,).

-> **NOTE:** - One cluster supports a maximum of three security groups.

-> **NOTE:** - This parameter is supported only when the value of `WhiteListType` is **SecurityGroup.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `security_ips` - (Optional, Computed) The secutity ips.
* `serverless_type` - (Optional, ForceNew, Available since v1.204.0) Serverless type. The current value is fixed to `AgileServerless` (sensitive state).
* `source_resource_id` - (Optional) Source RDS instance ID or source PolarDB cluster ID. This parameter is mandatory only when `CreationOption` is set to `MigrationFromRDS`, `CloneFromRDS`, `CloneFromPolarDB`, or `RecoverFromRecyclebin`. 
  - If `CreationOption` is `MigrationFromRDS` or `CloneFromRDS`, you need to input the source RDS instance ID. The source RDS instance version must be RDS MySQL 5.6, 5.7, or 8.0 High Availability edition. 
  - If `CreationOption` is `CloneFromPolarDB`, you need to input the source PolarDB cluster ID. The DBType of the cloned cluster will default to match the source cluster. For example, if the source cluster is MySQL 8.0, the cloned cluster must also have `DBType` set to `MySQL` and `DBVersion` to **8.0**. 
  - If `CreationOption` is `RecoverFromRecyclebin`, you need to input the released source PolarDB cluster ID. The DBType of the cluster being recovered from the recycle bin must match the source cluster. For example, if the source cluster was MySQL 8.0, the recovered cluster must also have `DBType` set to `MySQL` and `DBVersion` to **8.0**.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `standby_az` - (Optional, Computed) The zone in which the hot standby cluster is stored. Suitable for standard 3AZ scenarios.

-> **NOTE:**  This parameter takes effect only when the multi-zone strong data consistency is enabled.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `storage_auto_scale` - (Optional, Available since v1.270.0) Storage automatic expansion switch

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `storage_pay_type` - (Optional, ForceNew, Computed, Available since v1.210.0) The storage billing type, with valid values as follows: 
  - Postpaid: Pay-as-you-go (hourly). 
  - Prepaid: Pay-per-use based on space (subscription).
* `storage_space` - (Optional, Computed, Int, Available since v1.203.0) Storage space billed by space (package year and month). Unit: GB.
* `storage_type` - (Optional, ForceNew, Computed, Available since v1.203.0) Enterprise edition storage types include: 
  - `PSL5` 
  - `PSL4` 

Standard edition storage types include: 
  - `ESSDPL0` 
  - `ESSDPL1` 
  - `ESSDPL2` 
  - `ESSDPL3` 
  - `ESSDAUTOPL`
* `storage_upper_bound` - (Optional, Int, Available since v1.270.0) Automatic storage capacity limit

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `strict_consistency` - (Optional, ForceNew, Computed, Available since v1.239.0) Whether the cluster has multi-zone data consistency enabled. Value range:
  - `ON`: indicates that multi-zone data consistency is enabled, applicable to standard 3AZ scenarios.
  - `OFF`: indicates that multi-zone data consistency is not enabled.
* `sub_category` - (Optional, Computed) Cluster subseries. The range of values is as follows:
  - `normal_exclusive`: exclusive specification
  - `normal_general`: General specification
* `tags` - (Optional, Map) Tags
* `tde_status` - (Optional) Enables TDE encryption. Valid values are as follows: 
  - `true`: Enabled. 
  - `false`: Disabled (default). 

-> **NOTE:**  * This parameter takes effect only when `DBType` is `PostgreSQL` or `Oracle`. 

-> **NOTE:**  * You can call the [ModifyDBClusterTDE](https://www.alibabacloud.com/help/en/doc-detail/167982.html) interface to enable TDE encryption for a PolarDB MySQL cluster. 

-> **NOTE:**  * Once the TDE feature is enabled, it cannot be disabled.

* `used_time` - (Optional, Available since v1.270.0) If the payment type is `Prepaid`, this parameter is required.
  - When `Period` is `Month`, `UsedTime` should be an integer within `[1-9]`.
  - When `Period` is `Year`, `UsedTime` should be an integer within `[1-3]`.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `vswitch_id` - (Optional, Computed) The VSwitchId of db cluster
* `vpc_id` - (Optional, Computed, Available since v1.185.0) VpcId
* `white_list_type` - (Optional, Available since v1.270.0) The whitelist type. The value range is as follows:
  - `IP`:IP whitelist group.
  - `SecurityGroup`: Security Group.

The default value is **IP * *.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `zone_id` - (Optional, Computed) ZoneId
* `zone_type` - (Optional, Available since v1.270.0) Type of the replacement zone. The value range is as follows:
  - `Primary`: The Primary zone.
  - `Standby`: Standby area.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - CreateTime.
* `parameters` - DBClusterParameters.
  * `checking_code` - The checking code.
  * `data_type` - The data type.
  * `force_restart` - Is force restart.
  * `is_modifiable` - Is modifable.
  * `name` - ParameterName.
  * `parameter_description` - The description of parameter.
  * `parameter_status` - The status of parameter.
  * `value` - ParameterValue.
* `region_id` - RegionId.
* `seconds_until_auto_pause` - The duration of the paused probe is enabled (it needs to be adjusted by AllowShutDown to true).
* `status` - The status of the resource.

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