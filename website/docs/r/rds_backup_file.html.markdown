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

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `backup_file` - (Optional) A JSON array that describes the backup file information in an OSS bucket. Example:  
`{"Bucket":"test", "Object":"test/test_db_employees.xb","Location":"ap-southeast-1"}`  

The parameters in the array are described as follows:  
  - `Bucket`: The name of the OSS bucket where the backup file is stored. You can obtain it by calling [GetBucket](https://help.aliyun.com/document_detail/31965.html).  
  - `Object`: The full path of the directory containing the backup file. You can obtain it by calling [GetObject](https://help.aliyun.com/document_detail/31980.html).  
  - `Location`: The region ID of the OSS bucket. You can obtain it by calling [GetBucketLocation](https://help.aliyun.com/document_detail/31967.html).  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `bucket_region` - (Optional) The region ID of the OSS bucket where the self-managed MySQL 5.7 backup file is stored. You can obtain it by calling DescribeRegions.  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `build_replication` - (Optional) Specifies whether to automatically set up replication. Valid values:
  - true: Yes. The `MasterInfo` parameter is required.
  - false: No.

-> **NOTE:**  This parameter applies only to native replication instances and requires the `DBInstanceId` parameter to be specified when calling the API.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `comment` - (Optional) Remarks for the user's backup to be imported.
* `db_instance_id` - (Optional) Instance ID.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `engine_version` - (Required, ForceNew) The database version.  
* `master_info` - (Optional) A case-sensitive JSON array containing master information for setting up MySQL replication. Example:

```
{"masterIp":"172.20.xx.xx","masterPort":"3306","masterUser":"replica","masterPassword":"W33uopkehBQ="}
```

Descriptions of the parameters in the array are as follows:
  - `masterIp`: IP address of the primary instance.
  - `masterPort`: Port of the primary instance.
  - `masterUser`: Replication account on the primary instance.
  - `masterPassword`: Password for the replication account on the primary instance, Base64-encoded.

-> **NOTE:**  Applies only to native replication instances. The `DBInstanceId` parameter must be provided when calling the API.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `mode` - (Optional) The import mode. Valid values:  
  - oss: Import the backup by downloading it from OSS.  
  - stream: Import the backup over the network.  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `resource_group_id` - (Optional, Computed) The resource group ID. You can obtain it by calling the DescribeDBInstanceAttribute operation.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `restore_size` - (Optional, ForceNew) The storage space required to restore the user's backup. Unit: GB.

-> **NOTE:**  * Defaults to 5 times the size of the backup file.

-> **NOTE:**  * Minimum value is 20.

* `retention` - (Optional) The retention period for user backup files. Unit: days. The value must be an integer of type Integer greater than `0`.
* `source_info` - (Optional) A case-sensitive JSON array that provides source information for full backup. Example:

```
{"sourceIp":"172.20.xx.xx","sourcePort":"9999"}
```

The parameters in the array are described as follows:
  - `sourceIp`: The IP address of the source.
  - `sourcePort`: The port on which Netcat listens at the source.

-> **NOTE:**  This parameter applies only to native replication instances and requires the `DBInstanceId` parameter to be specified when calling the API.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `zone_id` - (Optional, ForceNew) The zone ID where the user backup is located.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The start time of the user backup import, in Unix timestamp format.
* `status` - Query the status of the user's backup file.

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