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
* `audit_log_status` - (Optional, Available since v1.270.0) AuditLogStatus
* `auto_renewal_period` - (Optional, Int, Available since v1.270.0) AutoRenewPeriod.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `auto_renewal_period_unit` - (Optional, Available since v1.270.0) Unit of renewal duration. The default value is Month.
Year: Year.
Month: Month.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `auto_renewal_status` - (Optional, Available since v1.270.0) RenewalStatus

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `backup_set_id` - (Optional, Available since v1.211.1) BackupSetId

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `capacity` - (Optional, Int, Available since v1.270.0) Lake storage acceleration space size, unit is GB
* `compute_resource` - (Optional, Computed) ComputeResource
* `db_cluster_description` - (Optional, Computed, Available since v1.198.0) DBClusterDescription
* `db_cluster_network_type` - (Required, ForceNew, Available since v1.270.0) DBClusterNetworkType
* `db_cluster_version` - (Required, ForceNew) DBClusterVersion
* `db_cluster_ip_array_attribute` - (Optional, Available since v1.270.0) DbClusterIpArrayAttribute

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_cluster_ip_array_name` - (Optional, Available since v1.270.0) DbClusterIpArrayName

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_version` - (Optional, Available since v1.270.0) Kernel version

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `disk_encryption` - (Optional, ForceNew, Available since v1.245.0) DiskEncryption
* `enable_compaction_service` - (Optional, Available since v1.270.0) Remote build service switch
* `enable_default_resource_group` - (Optional) EnableDefaultResourceGroup

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `enable_essd_cache` - (Optional, Computed, Available since v1.270.0) After the cloud disk cache mode is enabled, full data is stored in OSS. EIU cloud disks are used for caching, and existing hot data is not affected.
* `enable_lake_cache` - (Optional, Available since v1.270.0) Lake Storage Acceleration Switch
* `enable_ssl` - (Optional, Computed, Available since v1.245.0) EnableSsl
* `engine_type` - (Optional, Available since v1.270.0) The computing engine type. Value Description:
  - XIHE (default): xi-he computing engine.
  - SPARK: The Spark computing engine.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `essd_cache_size` - (Optional, Int, Available since v1.270.0) Cloud Disk cache size, in GB
* `kms_id` - (Optional, ForceNew, Available since v1.245.0) KmsId
* `maintain_time` - (Optional, Available since v1.270.0) MaintainTime
* `modify_mode` - (Optional, Available since v1.270.0) ModifyMode

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `payment_type` - (Optional, ForceNew, Computed) The paymen type of the resource
* `period` - (Optional, Available since v1.245.0) Period

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `product_form` - (Optional, Computed, Available since v1.245.0) ProductForm
* `product_version` - (Optional, ForceNew, Computed, Available since v1.245.0) ProductVersion
* `reserved_node_count` - (Optional, Computed, Int, Available since v1.245.0) ReservedNodeCount
* `reserved_node_size` - (Optional, Computed, Available since v1.245.0) ReservedNodeSize
* `resource_group_id` - (Optional, Computed, Available since v1.211.1) The ID of the resource group
* `restore_to_time` - (Optional) RestoreToTime

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `restore_type` - (Optional, Available since v1.211.1) RestoreType

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `secondary_vswitch_id` - (Optional, ForceNew, Available since v1.248.0) SecondaryVSwitchId
* `secondary_zone_id` - (Optional, ForceNew, Available since v1.248.0) SecondaryZoneId
* `security_ips` - (Optional, Computed, Available since v1.198.0) SecurityIps
* `source_db_cluster_id` - (Optional, Available since v1.211.1) SourceDbClusterId

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `storage_resource` - (Optional, Computed) StorageResource
* `switch_mode` - (Optional, Int, Available since v1.270.0) Execution time, value:
  - 0: Execute immediately. Default value.
  - 1: Execute during The maintainable time period.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `tags` - (Optional, Map, Available since v1.270.0) The tag of the resource
* `used_time` - (Optional, Available since v1.270.0) UsedTime

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `vpc_id` - (Required, ForceNew) VPCId
* `vswitch_id` - (Required, ForceNew) VSwitchId
* `zone_id` - (Required, ForceNew) The zone ID  of the resource

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `commodity_code` - CommodityCode.
* `connection_string` - ConnectionString.
* `create_time` - The creation time of the resource.
* `engine` - Engine.
* `engine_version` - EngineVersion.
* `expire_time` - ExpireTime.
* `expired` - Expired.
* `lock_mode` - LockMode.
* `lock_reason` - LockReason.
* `port` - Port.
* `region_id` - The region ID of the resource.
* `status` - The status of the resource.

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