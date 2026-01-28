---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_backup_file"
description: |-
  Provides a Alicloud RDS Backup File resource.
---

# alicloud_rds_backup_file

Provides a RDS Backup File resource.

Details of user backup.

For information about RDS Backup File and how to use it, see [What is Backup File](https://next.api.alibabacloud.com/document/Rds/2014-08-15/ImportUserBackupFile).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `backup_file` - (Optional) The JSON Array that describes the backup file information in the OSS Bucket. Example:
'{"Bucket":"test", "Object":"test/test_db_employees.xb","Location":"ap-southeast-1"}'

The description of each parameter in the array is as follows:
  - `Bucket`: the name of the OSS Bucket where the backup file resides. You can call [GetBucket](~~ 31965 ~~) to get it.
  - `Object`: the detailed path of the directory where the backup file is located. You can call [GetObject](~~ 31980 ~~) to get it.
  - `Location`: the ID of the region where the OSS Bucket is located. You can call [GetBucketLocation](~~ 31967 ~~) to get it.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `bucket_region` - (Optional) The region ID of the OSS Bucket where the user-created MySQL 5.7 backup file resides. Callable DescribeRegions to get.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `build_replication` - (Optional) Whether to automatically create copy, value description:
  - true: Yes, the 'MasterInfo' parameter is required.
  - false: No.

-> **NOTE:**  takes effect only for native replication instances. You must pass the 'DBInstanceId' parameter when calling the API.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `comment` - (Optional) Comment
* `db_instance_id` - (Optional) The instance ID.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `engine_version` - (Required, ForceNew) EngineVersion
* `master_info` - (Optional) Build a JSON Array of Master information copied by MySQL (case sensitive). Example:

'''
{"masterIp":"172.20.xx.xx","masterPort":"3306","masterUser":"replica","masterPassword":"W33uopkehBQ="}

'''

The description of each parameter in the array is as follows:
  -'masterIp': master library IP.
  -'masterPort': Main Library port.
  -'masterUser': the master database replication account.
  -'masterPassword': the password of the master database replication account, which requires Base64 encryption.

-> **NOTE:**  takes effect only for native replication instances. You must pass the 'DBInstanceId' parameter when calling the API.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `mode` - (Optional) Import mode, value description:
  - oss: Download and import backups from OSS.
  - stream: Import the backup over the network.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `resource_group_id` - (Optional, Computed) The ID of the resource group. Callable DescribeDBInstanceAttribute to get.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `restore_size` - (Optional, ForceNew) RestoreSize
* `retention` - (Optional) Retention
* `source_info` - (Optional) A case-sensitive JSON Array that provides the source information of the full backup. Example:

'''
{"sourceIp":"172.20.xx
.xx","sourcePort":"9999"}

'''

The description of each parameter in the array is as follows:
  -'Source': source IP address.
  -'sourcePort': the port on which the source Netcat listens.

-> **NOTE:**  takes effect only for native replication instances. You must pass the 'DBInstanceId' parameter when calling the API.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `zone_id` - (Optional, ForceNew) ZoneId

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Backup File.
* `delete` - (Defaults to 5 mins) Used when delete the Backup File.
* `update` - (Defaults to 5 mins) Used when update the Backup File.

## Import

RDS Backup File can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_backup_file.example <backup_id>
```