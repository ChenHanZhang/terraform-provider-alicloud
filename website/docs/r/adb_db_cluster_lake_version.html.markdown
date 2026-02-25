---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_db_cluster_lake_version"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Db Cluster Lake Version resource.
---

# alicloud_adb_db_cluster_lake_version

Provides a AnalyticDB for MySQL (ADB) Db Cluster Lake Version resource.

ADB MySQL DBCluster.

For information about AnalyticDB for MySQL (ADB) Db Cluster Lake Version and how to use it, see [What is Db Cluster Lake Version](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/developer-reference/api-adb-2021-12-01-createdbcluster).

-> **NOTE:** Available since v1.190.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = "ap-southeast-1"
}

data "alicloud_adb_zones" "default" {
}

data "alicloud_vpcs" "default" {
  name_regex = "^default-NODELETING$"
}

data "alicloud_vswitches" "default" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_adb_zones.default.ids.0
}

resource "alicloud_adb_db_cluster_lake_version" "default" {
  db_cluster_version            = "5.0"
  vpc_id                        = data.alicloud_vpcs.default.ids.0
  vswitch_id                    = data.alicloud_vswitches.default.ids.0
  zone_id                       = data.alicloud_adb_zones.default.ids.0
  compute_resource              = "16ACU"
  storage_resource              = "0ACU"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}
```

## Argument Reference

The following arguments are supported:
* `audit_log_status` - (Optional, Available since v1.272.0) The status of SQL audit logging. Valid values:
  - `on`: Enables SQL audit.
  - `off`: Disables SQL audit.

-> **NOTE:**  After SQL audit is disabled, all SQL audit logs are cleared. Before disabling SQL audit, query and export your SQL audit logs. For more information, see [DescribeAuditLogRecords](https://help.aliyun.com/document_detail/612426.html). When SQL audit is re-enabled, audit logs will be collected starting from the time when the feature was most recently enabled.

* `auto_renewal_period` - (Optional, Int, Available since v1.272.0) The auto-renewal period. The default value is 1. Valid values are as follows:
  - When `AutoRenewalPeriod` is set to `Month`, the value must be an integer from 1 to 11.
  - When `AutoRenewalPeriod` is set to `Year`, the value must be 1, 2, 3, or 5 (integer).

-> **NOTE:**  Longer renewal periods offer better pricing. Renewing for one year is more cost-effective than renewing for 10 or 11 months.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `auto_renewal_period_unit` - (Optional, Available since v1.272.0) The auto-renewal period unit. Valid values:
  - Year: year.
  - Month: month.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `auto_renewal_status` - (Optional, Available since v1.272.0) The auto-renewal status. Valid values:
  - `AutoRenewal`: Auto-renewal is enabled.
  - `Normal`: Manual renewal. A text message reminder is sent before the cluster expires.
  - `NotRenewal`: No renewal upon expiration. No reminders are sent except a single notification three days before expiration indicating that renewal will not occur.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `backup_set_id` - (Optional, Available since v1.211.1) The ID of the backup set used for restoration from a backup.  

-> **NOTE:**  You can call the [DescribeBackups](https://help.aliyun.com/document_detail/612318.html) operation to view the backup list of the cluster.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `capacity` - (Optional, Int, Available since v1.272.0) The size of the lake storage acceleration cache space, in GB.
* `compute_resource` - (Optional, Computed) Reserved compute resources. Valid range: 0 ACU to 4096 ACU, in increments of 16. 1 ACU is approximately equivalent to 1 vCPU and 4 GB of memory.  

-> **NOTE:**  You must include the unit when specifying this parameter.

* `db_cluster_description` - (Optional, Computed, Available since v1.198.0) The description of the cluster.
  - It cannot start with `http://` or `https://`.
  - It must be 2 to 256 characters in length.
* `db_cluster_network_type` - (Required, ForceNew, Available since v1.272.0) Cluster network type. Only `VPC` (Virtual Private Cloud) is supported.
* `db_cluster_version` - (Required, ForceNew) The version of the lakehouse cluster. Valid value: **5.0**.
* `db_cluster_ip_array_attribute` - (Optional, Available since v1.272.0) Whitelist group attribute. The default value is empty.  

-> **NOTE:**  Groups with the "hidden" attribute are not displayed in the console. These groups are typically used for accessing DTS or PolarDB services.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_cluster_ip_array_name` - (Optional, Available since v1.272.0) The name of the IP address whitelist group. By default, operations apply to the Default group.
  - The name must start with a lowercase letter and end with a digit or lowercase letter. It can contain lowercase letters, digits, and underscores (_), and must be 2 to 32 characters in length.
  - A single cluster supports up to 50 whitelist groups.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_version` - (Optional, Available since v1.272.0) The target kernel version to which you want to upgrade.

-> **NOTE:**  You can call the `DescribeKernelVersion` operation to query the list of kernel versions supported by the cluster.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `disk_encryption` - (Optional, ForceNew, Available since v1.245.0) Specifies whether to enable disk encryption.
* `enable_compaction_service` - (Optional, Available since v1.272.0) Specifies whether to enable the remote build service.
* `enable_default_resource_group` - (Optional) Specifies whether all reserved compute resources are allocated to the default resource group (user_default). Valid values:
  - `true` (default): Yes.
  - `false`: No.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `enable_essd_cache` - (Optional, Computed, Available since v1.272.0) Specifies whether to enable ESSD cache.
* `enable_lake_cache` - (Optional, Available since v1.272.0) Specifies whether to enable lake storage acceleration.
* `enable_ssl` - (Optional, Computed, Available since v1.245.0) Specifies whether to enable SSL encryption for connections. Valid values:
  - `true`: Enabled.
  - `false`: Disabled.
* `engine_type` - (Optional, Available since v1.272.0) The compute engine type. Valid values:
  - XIHE (`default`): Xihe compute engine.
  - SPARK: Spark compute engine.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `essd_cache_size` - (Optional, Int, Available since v1.272.0) ESSD cache size, in GB.
* `kms_id` - (Optional, ForceNew, Available since v1.245.0) The ID of the key used to encrypt cloud disk data.  

-> **NOTE:**  This parameter is used only when cloud disk encryption is enabled for the AnalyticDB for MySQL cluster.  

* `maintain_time` - (Optional, Available since v1.272.0) The maintenance window for the cluster, in the format hh:mmZ-hh:mmZ.

-> **NOTE:**  The time window must be exactly one hour long and aligned to the hour.

* `modify_mode` - (Optional, Available since v1.272.0) The method for modifying the IP whitelist. Valid values:
  - `Cover` (default): Overwrite the existing IP whitelist.
  - `Append`: Add IP addresses to the whitelist.
  - `Delete`: Remove IP addresses from the whitelist.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `payment_type` - (Optional, Computed) The billing method. Valid values:
  - `Postpaid`: Pay-as-you-go.
  - `Prepaid`: Subscription.
* `period` - (Optional, Available since v1.245.0) Specify whether a subscription cluster is billed on a yearly or monthly basis. Valid values:  
  - `Year`: Yearly subscription.  
  - `Month`: Monthly subscription.  

-> **NOTE:**  This parameter is required when PayType is set to Prepaid (subscription).


-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `product_form` - (Optional, Computed, Available since v1.245.0) Product form. Valid values:
  - `IntegrationForm`: Integrated form
  - `LegacyForm`: Lakehouse Edition
* `product_version` - (Optional, ForceNew, Computed, Available since v1.245.0) Product version. Valid values:
  - `BasicVersion`: Basic Edition.
  - `EnterpriseVersion`: Enterprise Edition.

-> **NOTE:**  This parameter is required only when `ProductForm` is set to `IntegrationForm`.

* `reserved_node_count` - (Optional, Computed, Int, Available since v1.245.0) The number of reserved nodes. 
  - The default value for Enterprise Edition is 3 nodes, with a step size of 3.
  - The default value for Basic Edition is 1 node.

-> **NOTE:**  This parameter is required only when ProductForm is set to IntegrationForm.

* `reserved_node_size` - (Optional, Computed, Available since v1.245.0) Reserved node specification, measured in ACU.
* `resource_group_id` - (Optional, Computed, Available since v1.211.1) The ID of the resource group into which the cloud resource instance is to be moved.

-> **NOTE:**  A resource group is a mechanism for grouping and managing resources within an Alibaba Cloud account. Resource groups help you manage resource grouping and authorization complexities within a single Alibaba Cloud account. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html).

* `restore_to_time` - (Optional) The point in time to which the instance is restored based on a backup set.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `restore_type` - (Optional, Available since v1.211.1) The restoration method. Valid values:
  - `backup`: Restore from a backup set. You must also specify the `BackupSetId` and `SourceDBClusterId` parameters.
  - `timepoint`: Restore to a specific point in time. You must also specify the `RestoreToTime` and `SourceDBClusterId` parameters.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `secondary_vswitch_id` - (Optional, ForceNew, Available since v1.248.0) The secondary vSwitch ID.  

-> **NOTE:**  The value of this parameter must be different from that of the VSwitchId parameter.

* `secondary_zone_id` - (Optional, ForceNew, Available since v1.248.0) The secondary zone ID.  

-> **NOTE:**  The value of this parameter must be different from that of the ZoneId parameter.

* `security_ips` - (Optional, Computed, Available since v1.198.0) The IP address whitelist for the cluster. Separate multiple IP addresses with commas (,). Duplicate entries are not allowed, and up to 500 IP addresses are supported. The following formats are supported:
  - IP address format, for example: 10.23.xx.xx.
  - CIDR notation, for example: 10.23.xx.xx/24 (Classless Inter-Domain Routing; the number 24 indicates the prefix length, which can range from 1 to 32).
* `source_db_cluster_id` - (Optional, Available since v1.211.1) The source instance ID of the AnalyticDB for MySQL Data Warehouse Edition cluster. Specify this parameter to restore a Lakehouse Edition cluster from a Data Warehouse Edition cluster.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `storage_resource` - (Optional, Computed) Reserved storage resources. Valid values: 0 ACU to 2064 ACU, in increments of 24. 1 ACU is approximately equivalent to 1 vCPU and 4 GB memory.  

-> **NOTE:**  You must include the unit when specifying this parameter.

* `switch_mode` - (Optional, Int, Available since v1.272.0) The execution time. Valid values:
* `0`: Execute immediately (default).
* `1`: Execute during the maintenance window.

-> **NOTE:**  You can call [ModifyDBClusterMaintainTime](https://help.aliyun.com/document_detail/612236.html) to modify the maintenance window of the cluster.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `tags` - (Optional, Map, Available since v1.272.0) The tag information of the cluster.
* `used_time` - (Optional, Available since v1.272.0) Specifies the subscription duration for a subscription cluster. Valid values:  
  - When `Period` is set to Year, the value of UsedTime must be an integer from 1 to 3.  
  - When `Period` is set to Month, the value of UsedTime must be an integer from 1 to 9.  

-> **NOTE:**  This parameter is required when `PayType` is set to `Prepaid` (subscription).  


-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `vpc_id` - (Required, ForceNew) The ID of the Virtual Private Cloud (VPC).
* `vswitch_id` - (Required, ForceNew) The ID of the vSwitch.
* `zone_id` - (Required, ForceNew) The zone ID.

-> **NOTE:**  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the zone IDs available for a specified lakehouse cluster.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `commodity_code` - The commodity code for sales.
* `connection_string` - The internal or public endpoint for which you want to create or update a server certificate.
* `create_time` - Resource property field representing the creation time.
* `engine` - The engine, AnalyticDB.
* `engine_version` - The engine version.
* `expire_time` - Expiration time.
* `expired` - Indicates whether the cluster instance has expired.
* `lock_mode` - Lock mode of the cluster instance:.
* `lock_reason` - The reason for locking.
* `port` - The port of the instance.
* `region_id` - The region ID.
* `status` - The cluster status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 9 mins) Used when create the Db Cluster Lake Version.
* `delete` - (Defaults to 5 mins) Used when delete the Db Cluster Lake Version.
* `update` - (Defaults to 47 mins) Used when update the Db Cluster Lake Version.

## Import

AnalyticDB for MySQL (ADB) Db Cluster Lake Version can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_db_cluster_lake_version.example <db_cluster_id>
```