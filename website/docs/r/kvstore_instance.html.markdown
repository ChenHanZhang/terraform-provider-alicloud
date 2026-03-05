---
subcategory: "Tair (Redis OSS-Compatible) And Memcache (KVStore)"
layout: "alicloud"
page_title: "Alicloud: alicloud_kvstore_instance"
description: |-
  Provides a Alicloud Tair (Redis OSS-Compatible) And Memcache (KVStore) Db Instance resource.
---

# alicloud_kvstore_instance

Provides a Tair (Redis OSS-Compatible) And Memcache (KVStore) Db Instance resource.

Redis instance integrates computing, storage and network resources, and provides database users, database tables, network, backup and other management functions. It is easy to customize and change the instance configuration.

For information about Tair (Redis OSS-Compatible) And Memcache (KVStore) Db Instance and how to use it, see [What is Db Instance](https://www.alibabacloud.com/help/en/redis/developer-reference/api-r-kvstore-2015-01-01-createinstances-redis).

-> **NOTE:** Available since v1.14.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}

data "alicloud_resource_manager_resource_groups" "default" {
  status = "OK"
}

data "alicloud_kvstore_zones" "default" {
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.4.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "10.4.0.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_kvstore_zones.default.zones.0.id
}

resource "alicloud_kvstore_instance" "default" {
  db_instance_name  = var.name
  vswitch_id        = alicloud_vswitch.default.id
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
  zone_id           = data.alicloud_kvstore_zones.default.zones.0.id
  instance_class    = "redis.master.large.default"
  instance_type     = "Redis"
  engine_version    = "5.0"
  security_ips      = ["10.23.12.24"]
  config = {
    appendonly             = "yes"
    lazyfree-lazy-eviction = "yes"
  }
  tags = {
    Created = "TF",
    For     = "example",
  }
}
```

### Deleting `alicloud_kvstore_instance` or removing it from your configuration

The `alicloud_kvstore_instance` resource allows you to manage  `payment_type = "PayAsYouGo"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `auto_pay` - (Optional, Available since v1.273.0) Automatic transfer

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `auto_renew` - (Optional, Available since v1.36.0) Turn on or off automatic renewal, value:
* `true`: Enable.
* **false * *. Close.

-> **NOTE:**  default value: **false * *.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_renew_period` - (Optional, Available since v1.36.0) Auto renew period

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `backup_id` - (Optional) Backup id

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `business_info` - (Optional, Available since v1.101.0) Activity ID and business information.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `capacity` - (Optional, ForceNew, Computed, Int, Available since v1.101.0, Deprecated since v1.273.0) Capacity
* `cluster_backup_id` - (Optional, Available since v1.273.0) The ID of the backup set of the cluster.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `connection_string_prefix` - (Optional) The prefix of the public network connection address. It must be composed of lowercase letters and numbers, start with lowercase letters, and be 8 to 40 characters in length.

-> **NOTE:**  the format of the public network connection address is: ' .redis.rds.aliyuncs.com'.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `coupon_no` - (Optional, Available since v1.101.0) The default value is: 'youhuiquan_promotion_option_id_for_blank'.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `db_instance_name` - (Optional, Computed, Available since v1.101.0) Instance name
* `db_audit` - (Optional, Available since v1.273.0) Whether to enable the audit log. Valid values:
* `true`: Default value, on.
* `false`: closed.

-> **NOTE:**  If the instance is [Cluster Architecture](~~ 52228 ~~) or [Read/Write Separation Architecture](~~ 62870 ~~), the audit logs of both the data node and the proxy node are enabled or closed.

* `dry_run` - (Optional, Available since v1.128.0) Dry Run

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `effective_time` - (Optional, Available since v1.204.0) Effective Time

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `encryption_key` - (Optional, Computed, Available since v1.200.0) The custom key ID. You can call [DescribeEncryptionKeyList](~~ 302339 ~~) to obtain the key ID.

-> **NOTE:**  * If this parameter is not specified, the [Key Management Service](~~ 28935 ~~) automatically generates the key.

-> **NOTE:**  * To create a custom key, you can call the [CreateKey](~~ 28947 ~~) operation of the key management service.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `encryption_name` - (Optional, Computed, Available since v1.200.0) Encryption algorithm, default AES-CTR-256.

-> **NOTE:**  This parameter is only available when the `TDEStatus` parameter is set to **Enabled.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `engine_version` - (Optional, Computed) Engine version
* `force_trans` - (Optional, Available since v1.273.0) Whether to enable forced transmission

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `force_upgrade` - (Optional, Available since v1.101.0) Whether to force the configuration change, the value is:
  - `false`: does not force the configuration change.
  - `true`: forced configuration change, default value.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `global_security_group_ids` - (Optional, Available since v1.273.0) Instance global IP whitelist Template

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_class` - (Optional) Instance class
* `instance_release_protection` - (Optional, Available since v1.101.0) Instance release protection
* `instance_type` - (Optional, ForceNew) Instance type
* `is_auto_upgrade_open` - (Optional, Computed, Available since v1.228.0) Automatic minor version upgrade switch, value:
1: Open.
0: Closed.
Example values:
0
Enumeration value:
0
1
* `maintain_end_time` - (Optional, Computed, Available since v1.56.0) Maintain end time
* `maintain_start_time` - (Optional, Computed, Available since v1.56.0) Maintain start time
* `minorversion` - (Optional, Available since v1.273.0) The minor version to be upgraded to. Default value: `latest_version`, which is the latest version.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `network_type` - (Optional, ForceNew, Available since v1.273.0) Network type
* `new_password` - (Optional, Available since v1.273.0) The new default account password. The default account is the account named after the instance ID (for example, r-bp10noxlhcoim2 ****).

-> **NOTE:**  length is 8~32 bits, must contain uppercase letters, lowercase letters, numbers, special characters (support '! @ #$%^& *()_+-= ') at least three.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `node_type` - (Optional, Computed) Node type, value:
MASTER_SLAVE: high availability (dual copy)
STAND_ALONE: single copy
double: double copy
single: single copy
Note For Cloud Native instances, select MASTER_SLAVE or STAND_ALONE. For Classic instances, select double or single.
* `order_type` - (Optional, Available since v1.101.0) This parameter is required to change the configuration type of the package year and month instance. The value is as follows:

* `UPGRADE`: UPGRADE configuration, default value.
* `DOWNGRADE`: DOWNGRADE configuration.

-> **NOTE:**  * The value of this parameter must be set to `DOWNGRADE` when the package year and month instance is downgraded * *.
* If the target specification of the change is higher than the price of the existing specification, it is an upgraded configuration, otherwise it is a degraded configuration. For example, the price of read-write splitting version 8g (5 read-only nodes) is higher than that of 16G cluster version, and the latter is changed to the former as an upgrade configuration.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `password` - (Optional) Instance password. The length is 8-32 bits and must contain at least three of uppercase letters, lowercase letters, special characters and numbers. The allowed special characters include '! @#$%^ & *()_+-= '.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `payment_type` - (Optional, Computed, Available since v1.101.0) Payment type
* `period` - (Optional) The payment cycle. Unit is Monthly. Valid values: `1` ~ `9`, `12`, `24`, `36`, and **60 * *.

-> **NOTE:**  `ChargeType` is available and must be passed in only when the value is **PrePaid.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `port` - (Optional, ForceNew, Int, Available since v1.94.0) Port
* `private_ip` - (Optional, ForceNew, Computed) Private IP
* `read_only_count` - (Optional, Int, Available since v1.226.0) The number of read-only nodes. This parameter is only applicable to the creation of read-write splitting instances in the cloud disk version. You can use this parameter to customize the number of read-only nodes. Valid values: 1 to 5.
* `resource_group_id` - (Optional, Computed, Available since v1.86.0) Resource group id
* `restore_time` - (Optional, Available since v1.101.0) To create a new instance based on the backup data of an instance, after you specify the source instance by using the `SrcDBInstanceId`, you can specify any point in time during the backup retention period of the source instance. The system will use the backup data of the source instance at that point in time to create a new instance. The format is  yyyy-MM-dd T  HH:mm:ss Z(UTC time).

-> **NOTE:**  After setting `SrcDBInstanceId`, you must select one of the `BackupId` and `RestoreTime` parameters to specify the backup data.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `retention` - (Optional, Int, Available since v1.273.0) Retention time
* `role_arn` - (Optional, Computed, Available since v1.200.0) Specify the global Resource descriptor ARN(Alibaba Cloud Resource Name) information of the role to be authorized. After the authorization is completed, the relevant key management service can be used in the format: 'acs:ram ::$ accountID:role/$roleName '.

-> **NOTE:**  * '$accountID': the ID of the cloud account. You can log on to the Alibaba Cloud console, hover your mouse over the profile picture in the upper-right corner, and click **Security Settings** to view the picture.

-> **NOTE:**  * '$roleName': the name of the RAM role. The value is fixed to **AliyunRdsInstanceEncryptionDefaultRole * *.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `secondary_zone_id` - (Optional, ForceNew, Available since v1.128.0) The ID of the standby zone. You can call the [DescribeZones](~~ 94527 ~~) operation.

-> **NOTE:**  This parameter can be passed in to realize that the primary and standby data nodes are distributed in different available areas, realize disaster tolerance across available areas, and can withstand computer room-level failures.

* `security_group_id` - (Optional, Available since v1.76.0) Security group id
* `security_ip_group_attribute` - (Optional, Available since v1.101.0) Security IP group attribute
* `security_ip_group_name` - (Optional, Computed, Available since v1.101.0) Security IP group name
* `security_ips` - (Optional, Computed, List) Security IPs
* `shard_count` - (Optional, Computed, Int, Available since v1.208.0) The number of slices. This parameter is only applicable to creating a cloud disk cluster architecture instance. You can use this parameter to customize the number of slices.
* `slave_read_only_count` - (Optional, Int, Available since v1.226.0) Specifies the number of read-only nodes in the secondary zone when creating a multi-zone read/write splitting instance. The sum of this parameter and ReadOnlyCount cannot be greater than 9.
Note: When creating a multi-zone read/write splitting instance, you must specify both the slaveadonlycount and SecondaryZoneId parameters.
* `source_biz` - (Optional, Available since v1.273.0) The source of the call. This parameter is only used for internal maintenance and does not need to be passed in.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `src_db_instance_id` - (Optional, Available since v1.273.0) To create a new instance based on the backup data of an instance, first specify the ID of the source instance in this parameter, and then specify the backup data to be used by using the `BackupId` or `RestoreTime` parameter.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ssl_enabled` - (Optional, ForceNew, Available since v1.273.0) SSL encryption status:
  - `Enable`: enabled.
  - `Disable`: not enabled.
* `tags` - (Optional, Map, Available since v1.55.3) Tags
* `tde_status` - (Optional, Computed, Available since v1.200.0) Enable the TDE function. Set the value to **Enabled * *.

-> **NOTE:**  TDE function cannot be turned off for the time being after it is turned on. Please evaluate whether it affects the business before opening. For more information, see [enable TDE for transparent data encryption](~~ 265913 ~~).


-> **NOTE:** This parameter is only evaluated during resource operations. Modifying it in isolation will not trigger any action.

* `token` - (Optional, Available since v1.273.0) Used to guarantee the idempotence of the request. The value of this parameter is generated by the client. It must be unique between different requests and must be case sensitive and have no more than 64 ASCII characters.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `vswitch_id` - (Optional, ForceNew) Vswitch id
* `vpc_auth_mode` - (Optional) Vpc auth mode
* `vpc_id` - (Optional, ForceNew, Available since v1.273.0) Vpc id
* `zone_id` - (Optional, ForceNew, Computed, Available since v1.101.0) Zone id

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `bandwidth` - Bandwidth.
* `config` - Config.
* `connection_domain` - Connection domain.
* `create_time` - Create time.
* `end_time` - End time.
* `qps` - QPS.
* `region_id` - Region id.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 35 mins) Used when create the Db Instance.
* `delete` - (Defaults to 21 mins) Used when delete the Db Instance.
* `update` - (Defaults to 40 mins) Used when update the Db Instance.

## Import

Tair (Redis OSS-Compatible) And Memcache (KVStore) Db Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_kvstore_instance.example <db_instance_id>
```