---
subcategory: "Tair (Redis OSS-Compatible) And Memcache (KVStore)"
layout: "alicloud"
page_title: "Alicloud: alicloud_redis_tair_instance"
description: |-
  Provides a Alicloud Tair (Redis OSS-Compatible) And Memcache (KVStore) Tair Instance resource.
---

# alicloud_redis_tair_instance

Provides a Tair (Redis OSS-Compatible) And Memcache (KVStore) Tair Instance resource.

Describes the creation, deletion, and query operations for Tair instances.

For information about Tair (Redis OSS-Compatible) And Memcache (KVStore) Tair Instance and how to use it, see [What is Tair Instance](https://www.alibabacloud.com/help/en/tair).

-> **NOTE:** Available since v1.206.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = "cn-hangzhou"
}

variable "name" {
  default = "tf-example"
}

data "alicloud_kvstore_zones" "default" {
  product_type = "Tair_rdb"
}

data "alicloud_vpcs" "default" {
  name_regex = "default-NODELETING"
}

data "alicloud_vswitches" "default" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_kvstore_zones.default.zones.0.id
}

locals {
  vswitch_id = data.alicloud_vswitches.default.ids.0
  zone_id    = data.alicloud_kvstore_zones.default.zones.0.id
}

data "alicloud_resource_manager_resource_groups" "default" {
}

resource "alicloud_redis_tair_instance" "default" {
  payment_type       = "Subscription"
  period             = "1"
  instance_type      = "tair_rdb"
  zone_id            = local.zone_id
  instance_class     = "tair.rdb.2g"
  shard_count        = "2"
  vswitch_id         = local.vswitch_id
  vpc_id             = data.alicloud_vpcs.default.ids.0
  tair_instance_name = var.name
}
```

### Deleting `alicloud_redis_tair_instance` or removing it from your configuration

The `alicloud_redis_tair_instance` resource allows you to manage  `payment_type = "PayAsYouGo"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `auto_renew` - (Optional) Specifies whether auto-renewal is enabled. Valid values:
  - `true`: Auto-renewal is enabled.
  - `false` (default): Auto-renewal is disabled.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_renew_period` - (Optional) Auto-renewal period, in months. Valid values: `1`, `2`, `3`, `6`, `12`.

-> **NOTE:**  This parameter is required when `AutoRenew` is set to `true`.  


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `backup_id` - (Optional, Available since v1.233.1) You can specify the backup set ID of the source instance in this parameter. The system creates a new instance based on the data stored in the specified backup set. You can call [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) to query the BackupId. If the source instance is a cluster instance, you must specify the backup set IDs of all shards in the source instance, separated by commas, for example, "10**,11**,15**".

-> **NOTE:**  If your instance uses a cloud-native architecture and is a cluster instance, we recommend that you call [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) to obtain the cluster backup set ID (for example, "cb-xx") and then use the ClusterBackupId request parameter to clone the cluster instance. This avoids the need to manually enter the backup set IDs of individual shards.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `cluster_backup_id` - (Optional, Available since v1.224.0) Some new cluster architectures support specifying a cluster backup set ID. You can obtain this ID by calling the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation.  
  - If supported, you can directly specify the cluster backup set ID without providing the `BackupId` parameter.  
  - If not supported, specify the backup set IDs of each shard in the source instance in the `BackupId` parameter, separated by commas, such as "2158****20,2158****22".  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `connection_string_prefix` - (Optional, Available since v1.235.0) The prefix of the connection string. It must consist of lowercase letters and digits, start with a lowercase letter, and be 8 to 40 characters in length.  

-> **NOTE:**   

-> **NOTE:**  The connection string format is: .redis.rds.aliyuncs.com.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `effective_time` - (Optional) The time when the configuration change takes effect. Valid values:
  - `Immediately`: The change takes effect immediately. This is the default value.
  - `MaintainTime`: The change takes effect during the maintenance window of the instance. You can call [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) to modify the maintenance window.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `encryption_key` - (Optional, Available since v1.273.0) The custom key ID. You can call [DescribeEncryptionKeyList](https://help.aliyun.com/document_detail/473860.html) to obtain it.

-> **NOTE:**  * If this parameter is not specified, a key is automatically generated by [Key Management Service](https://help.aliyun.com/document_detail/28935.html).

-> **NOTE:**  * To create a custom key, you can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation of Key Management Service.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `encryption_name` - (Optional, Available since v1.273.0) The encryption algorithm. The default value is AES-CTR-256.  

-> **NOTE:**  This parameter is available only when the `TDEStatus` parameter is set to `Enabled`.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `engine_version` - (Optional, Computed) The database version. Default value: **1.0**. The valid values vary by Tair instance type:
  - `Tair_rdb`: Tair in-memory instances are compatible with Redis 5.0, Redis 6.0, and Redis 7.0 protocols. Specify **5.0**, **6.0**, or **7.0**.
  - `Tair_scm`: Tair persistent memory instances are compatible with the Redis 6.0 protocol. Specify **1.0**.
  - `Tair_essd`: Tair disk-based instances (ESSD or SSD) are compatible with the Redis 6.0 protocol. Specify **1.0** for ESSD-based instances or **2.0** for SSD-based instances.
* `force_upgrade` - (Optional) Specifies whether to force specification changes. Valid values:
  - `false`: Do not force specification changes.
  - `true`: Force specification changes. This is the default value.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `global_instance_id` - (Optional, Available since v1.233.1) Specifies whether to create this new instance as a child instance in a distributed instance. This method allows you to create a distributed instance.
* To create this new instance as the first child instance, enter `true`.
* To create this new instance as the second or third child instance, enter the ID of the distributed instance, for example, gr-bp14rkqrhac****.
* If you do not want to create a distributed instance, leave this field empty.

-> **NOTE:**  If you want to create a distributed instance, the new instance must be of the Tair in-memory type.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_class` - (Required) The target instance type code after modification. You can call [DescribeAvailableResource](https://help.aliyun.com/document_detail/473765.html) to query the instance types available for configuration changes in the zone where the instance resides.

-> **NOTE:**  For details about each instance type, see [Instance Type Lookup Guide](https://help.aliyun.com/document_detail/26350.html).

* `instance_type` - (Required, ForceNew) The storage medium. Valid values:
  - `tair_rdb`: In-memory.
  - `tair_scm`: Persistent memory.
  - `tair_essd`: Disk-based.
* `intranet_bandwidth` - (Optional, Computed, Int, Available since v1.233.1) The total target bandwidth of the instance.  

-> **NOTE:**  If the instance is a cluster instance, the total target bandwidth must be divisible by the number of shards. This operation is not supported for read/write splitting instances with the new architecture.>.  

* `modify_mode` - (Optional, Available since v1.233.1) The modification mode. Valid values:  
  - `Cover` (default): Overwrites the original whitelist.  
  - `Append`: Adds to the whitelist.  
  - `Delete`: Deletes the specified whitelist entry.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `node_type` - (Optional, Computed) Node type. Valid values:  
* `MASTER_SLAVE`: High availability (dual replicas)  
* `STAND_ALONE`: Single replica  
* `double`: Dual replicas  
* `single`: Single replica  

-> **NOTE:**  For cloud-native instances, select `MASTER_SLAVE` or `STAND_ALONE`. For classic instances, select `double` or `single`.  

* `param_no_loose_sentinel_enabled` - (Optional, Computed, Available since v1.233.1) Sentinel compatibility mode, applicable to non-cluster instances. Valid values:  
* `no` (default): Disabled  
* `yes`: Enabled  

-> **NOTE:**  For more information, see [Sentinel Compatibility Mode](https://help.aliyun.com/document_detail/178911.html).  

* `param_no_loose_sentinel_password_free_access` - (Optional, Computed, Available since v1.237.0) When Sentinel mode is enabled, specifies whether to allow password-free execution of Sentinel-related commands. Valid values:
  - `no` (default): Disabled.
  - `yes`: Enabled. After this option is enabled, you can execute Sentinel commands without a password on any connection and use the SENTINEL command to listen to the +switch-master channel.
* `param_no_loose_sentinel_password_free_commands` - (Optional, Computed, Available since v1.239.0) After enabling Sentinel mode and the `ParamNoLooseSentinelPasswordFreeAccess` parameter, you can use this parameter to add additional password-free commands (empty by default).

-> **NOTE:**  * After configuration, the specified commands can be executed without a password on any connection. Proceed with caution.

-> **NOTE:**  * Commands must be in lowercase letters and separated by commas (,).

* `param_repl_mode` - (Optional, Computed, Available since v1.233.1) Replication mode:
  - `async` (default): Asynchronous mode
  - `semisync`: Semi-synchronous mode.
* `param_semisync_repl_timeout` - (Optional, Computed, Available since v1.233.1) The degradation threshold for semi-synchronous mode. This parameter applies only to semi-synchronous replication and is measured in milliseconds (ms). Valid values range from 10 to 60,000, with a default value of 500.  

-> **NOTE:**  If the replication delay exceeds this threshold, the replication mode automatically switches to asynchronous. Once the delay is resolved, the mode automatically reverts to semi-synchronous.  

* `param_sentinel_compat_enable` - (Optional, Computed, Available since v1.233.1) Sentinel compatibility mode, applicable to cluster instances in proxy connection mode or read/write splitting instances. Valid values:  
* `0` (default): Disabled  
* `1`: Enabled  

-> **NOTE:**  For more information, see [Sentinel Compatibility Mode](https://help.aliyun.com/document_detail/178911.html).  

* `password` - (Optional) The password of the instance, which must meet the following requirements:
  - It must be 8 to 32 characters in length.
  - It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Supported special characters are `!@#$%^&*()_+-=`.
* `payment_type` - (Optional, Computed) The billing method to switch to. Valid values:
  - `PrePaid`: Subscription. If you specify this value, you must also specify the `Period` parameter.
  - `PostPaid`: Pay-as-you-go.
* `period` - (Optional, Int) The billing cycle, measured in months. Valid values: `1` to `9`, `12`, `24`, `36`, and `60`.  

-> **NOTE:**  This parameter is required when `ChargeType` is set to `PrePaid`.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `port` - (Optional, ForceNew, Computed, Int) The service port of the instance. Valid values: 1 to 65535. Default value: 6379.
* `read_only_count` - (Optional, Int) The number of read-only nodes in the primary zone. This parameter applies only when creating a cloud-native read/write splitting instance.
* For a standard architecture instance, the value ranges from 1 to 9.
* For a cluster architecture instance, the value ranges from 1 to 4, indicating the number of read-only nodes per data shard.

-> **NOTE:**  If you are creating a multi-zone instance, you can use this parameter together with the SlaveReadOnlyCount parameter to customize the number of read-only nodes in the primary and secondary zones.

-> **NOTE:**  * For a standard architecture instance, the sum of this parameter and SlaveReadOnlyCount must not exceed 9.

-> **NOTE:**  * For a cluster architecture instance, the sum of this parameter and SlaveReadOnlyCount must not exceed 4.


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `recover_config_mode` - (Optional, Available since v1.233.1) Specifies whether to restore account, kernel parameters, and whitelist information from the specified backup set when creating an instance. For example, to restore account information, set this parameter to `account`.  

By default, this parameter is empty, which indicates that account, kernel parameters, and whitelist information are not restored from the backup set.  

-> **NOTE:**  This parameter applies only to cloud-native instances and requires that the original backup set already contains account, kernel parameters, and whitelist information. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to check whether the RecoverConfigMode parameter of a specified backup set includes the aforementioned information.  


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed) The ID of the destination resource group.

-> **NOTE:**  * You can obtain a list of resource group IDs by calling the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation or through the console. For more information, see [View basic information about a resource group](https://help.aliyun.com/document_detail/151181.html).

-> **NOTE:**  * Before modifying the resource group to which an instance belongs, you can call the [ListResources](https://help.aliyun.com/document_detail/158866.html) operation to view the current resource group of the instance.

* `role_arn` - (Optional, Available since v1.273.0) The Alibaba Cloud Resource Name (ARN) of the role to be authorized. After authorization, you can use the related Key Management Service (KMS) features. Format: `acs:ram::$accountID:role/$roleName`.  

-> **NOTE:**  * `$accountID`: Your Alibaba Cloud account ID. To view it, log on to the Alibaba Cloud console, move the pointer over your profile picture in the upper-right corner, and click **Security Settings**.  

-> **NOTE:**  * `$roleName`: The RAM role name. The value is fixed as `AliyunRdsInstanceEncryptionDefaultRole`.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `secondary_zone_id` - (Optional, ForceNew) The ID of the secondary zone. You can call [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) to query zone IDs.

-> **NOTE:**  The value of this parameter must differ from that of the ZoneId parameter, and you cannot specify an ID that represents multiple zones.

* `security_group_id` - (Optional) The security group ID to be configured. You can specify up to 10 security group IDs, separated by commas (,).

-> **NOTE:**  This will overwrite the original settings. For valid values, refer to the API operation for querying basic information about your created security groups: [DescribeSecurityGroups](https://next.api.aliyun.com/api/Ecs/2014-05-26/DescribeSecurityGroups).

* `security_ip_group_name` - (Optional, Computed, Available since v1.233.1) The name of the IP address whitelist group.
You cannot modify system-generated whitelist groups. If this parameter is not specified, the default group is modified by default.
* `security_ips` - (Optional, Computed, Available since v1.233.1) IP addresses in the IP whitelist group, up to 1,000 entries. Separate multiple IP addresses with commas. Supported formats include: 0.0.0.0/0, 10.23.12.24, or 10.23.12.24/24 (CIDR notation, where /24 indicates the prefix length, ranging from 1 to 32).  
* `shard_count` - (Optional, Computed, Int) The number of shards to delete. The system deletes the specified number of shard nodes starting from the end.

-> **NOTE:**  For example, if an instance originally has five shard nodes named db-0, db-1, db-2, db-3, and db-4, and you set this parameter to 2, the system deletes db-3 and db-4.

* `slave_read_only_count` - (Optional, Int) The number of read-only nodes in the secondary zone.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `src_db_instance_id` - (Optional, Available since v1.233.1) If you need to create a new instance based on a backup set from an existing instance, specify the source instance ID in this parameter.

-> **NOTE:**  Then, use one of the following parameters to specify the backup set or point in time: `BackupId`, `ClusterBackupId` (recommended for cloud-native cluster architecture), or `RestoreTime`. This parameter must be used together with one of these three parameters.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ssl_enabled` - (Optional, Computed) Modify TLS (SSL) settings. Valid values:  
  - `Disable`: Disable TLS (SSL).  
  - `Enable`: Enable TLS (SSL).  
  - `Update`: Update the certificate.  
* `storage_performance_level` - (Optional, ForceNew) The storage type. Valid values: `essd_pl1`, `essd_pl2`, and `essd_pl3`.  

-> **NOTE:**  This parameter is required only when `InstanceType` is set to `tair_essd` and you are creating an ESSD-based instance.

* `storage_size_gb` - (Optional, ForceNew, Computed, Int) The storage capacity of disk-based instances. The valid range varies depending on the instance type. For more information, see [Disk-based Instance Types](https://help.aliyun.com/document_detail/2527111.html).

-> **NOTE:**  This parameter is required only when `InstanceType` is set to `tair_essd` and you are creating an ESSD-based instance. For Tair disk-based `SSD` instances, the storage capacity is predefined as a fixed value based on the instance type, so you do not need to specify this parameter.

* `tde_status` - (Optional, Available since v1.273.0) Specifies whether to enable Transparent Data Encryption (TDE). Valid value: `Enabled`.

-> **NOTE:**  After TDE is enabled, it cannot be disabled. Evaluate the potential impact on your business before enabling this feature. For more information, see [Enable Transparent Data Encryption (TDE)](https://help.aliyun.com/document_detail/265913.html).

* `tags` - (Optional, Map) The tag information of the instance.
* `tair_instance_name` - (Optional) The instance name must meet the following requirements:
  - It must be 2 to 80 characters in length.
  - It must start with a letter (uppercase or lowercase) or a Chinese character, and must not contain spaces or special characters: `@/:=”{[]}`.
* `vswitch_id` - (Required, ForceNew) The ID of the vSwitch in the specified VPC. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation of Virtual Private Cloud (VPC) to obtain this ID.
* `vpc_auth_mode` - (Optional, Computed, Available since v1.233.1) The authentication mode for Virtual Private Cloud (VPC), with the following valid values:
* `Open`: Password-based authentication is enabled.
* `Close`: Password-based authentication is disabled, enabling password-free access.

-> **NOTE:**  The default value is `Open`.

* `vpc_id` - (Required, ForceNew) The ID of the Virtual Private Cloud (VPC). You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation of VPC to obtain this ID.
* `zone_id` - (Required, ForceNew) Primary zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query available zones. Use this parameter to specify the zone where the instance will be created.  

-> **NOTE:**  You can also specify the `SecondaryZoneId` parameter to designate a secondary zone. The primary and secondary nodes will then be deployed in the specified primary and secondary zones, respectively, implementing a同城 dual-center active-standby architecture. For example, set `ZoneId` to "cn-hangzhou-h" and `SecondaryZoneId` to "cn-hangzhou-g".  


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `architecture_type` - The architecture type.
* `connection_domain` - The internal endpoint of the instance.
* `create_time` - The instance creation time, in the format yyyy-MM-ddTHH:mm:ssZ (UTC time).
* `max_connections` - The maximum number of connections allowed for the instance.
* `network_type` - The network type.
* `region_id` - The region ID.
* `status` - The instance status.
* `tair_instance_id` - The ID of the instance to query.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 60 mins) Used when create the Tair Instance.
* `delete` - (Defaults to 30 mins) Used when delete the Tair Instance.
* `update` - (Defaults to 60 mins) Used when update the Tair Instance.

## Import

Tair (Redis OSS-Compatible) And Memcache (KVStore) Tair Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_redis_tair_instance.example <tair_instance_id>
```