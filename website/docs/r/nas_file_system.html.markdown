---
subcategory: "File Storage (NAS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nas_file_system"
description: |-
  Provides a Alicloud File Storage (NAS) File System resource.
---

# alicloud_nas_file_system

Provides a File Storage (NAS) File System resource.

File system instance.

For information about File Storage (NAS) File System and how to use it, see [What is File System](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-createfilesystem).

-> **NOTE:** Available since v1.33.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

data "alicloud_nas_zones" "default" {
  file_system_type = "standard"
}

resource "alicloud_nas_file_system" "default" {
  protocol_type    = "NFS"
  storage_type     = "Capacity"
  description      = var.name
  encrypt_type     = 1
  file_system_type = "standard"
  recycle_bin {
    status        = "Enable"
    reserved_days = "10"
  }
  nfs_acl {
    enabled = true
  }
  zone_id = data.alicloud_nas_zones.default.zones.0.zone_id
}
```

## Argument Reference

The following arguments are supported:
* `capacity` - (Optional, Computed, Int, Available since v1.140.0) File system capacity. Unit: GiB.  
This parameter is required and valid only when FileSystemType is set to extreme, cpfs, or cpfsse.

For available values, refer to the actual specifications on the purchase pages:

  - [Extreme NAS Pay-as-you-go Purchase Page](https://common-buy.aliyun.com/?commodityCode=nas_extreme_post#/buy)
  - [CPFS Pay-as-you-go Purchase Page](https://common-buy.aliyun.com/?commodityCode=nas_cpfs_post#/buy)



  - [Extreme NAS Pay-as-you-go Purchase Page](https://common-buy-intl.alibabacloud.com/?commodityCode=nas_extpost_public_intl#/buy)
  - [CPFS Pay-as-you-go Purchase Page](https://common-buy-intl.alibabacloud.com/?spm=5176.nas_overview.0.0.7ea01dbft0dTui&commodityCode=nas_cpfspost_public_intl#/buy)
.
* `description` - (Optional) Description of the file system.  

Constraints:  
  - The description must be 2 to 128 characters in length and can contain letters, digits, Chinese characters, colons (:), underscores (_), or hyphens (-).  
  - It must start with a letter or a Chinese character and cannot start with `http://` or `https://`.
* `encrypt_type` - (Optional, ForceNew, Computed, Int, Available since v1.121.2) Whether the file system is encrypted.

Data at rest in the file system is encrypted using keys managed by Key Management Service (KMS). Data is automatically decrypted during read and write operations without requiring manual decryption.

Valid values:
  - 0 (default): Not encrypted.
  - 1: NAS-managed key. Supported when FileSystemType is set to standard or extreme.
  - 2: Customer-managed key. Supported when FileSystemType is set to standard or extreme.

-> **NOTE:**  - Extreme NAS: The customer-managed key (KMS) feature is supported in all regions except China East 1 Finance Cloud.

-> **NOTE:**  - General-purpose NAS: The customer-managed key (KMS) feature is supported in all regions.

* `file_system_type` - (Optional, ForceNew, Computed, Available since v1.140.0) The file system type.

Valid values:
  - standard (default): General-purpose NAS.
  - extreme: Extreme NAS.
  - cpfs: Cloud Parallel File System (CPFS) with local redundancy.
  - cpfsse: Cloud Parallel File System (CPFS) SE with zone-redundant storage.
* `keytab` - (Optional, Available since v1.248.0) A Base64-encoded string of the keytab file content.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `keytab_md5` - (Optional, Available since v1.248.0) The MD5 hash of the keytab file content.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `kms_key_id` - (Optional, ForceNew, Computed, Available since v1.140.0) KMS key ID.
This parameter is required only when EncryptType=2.
* `nfs_acl` - (Optional, Computed, Set, Available since v1.236.0) NFS ACL. See [`nfs_acl`](#nfs_acl) below.
* `options` - (Optional, Computed, Set, Available since v1.248.0) Options. See [`options`](#options) below.
* `protocol_type` - (Required, ForceNew) The file transfer protocol type.
  - When FileSystemType is set to standard, valid values are NFS and SMB.
  - When FileSystemType is set to extreme, the valid value is NFS.
  - When FileSystemType is set to cpfs, the valid value is cpfs.
  - When FileSystemType is set to cpfsse, the valid value is cpfs.
* `recycle_bin` - (Optional, Computed, Set) Recycle Bin. See [`recycle_bin`](#recycle_bin) below.
* `redundancy_type` - (Optional, ForceNew, Computed, Available since v1.267.0) Storage redundancy type. This parameter applies only to CPFS SE.
Valid value: ZRS.
* `redundancy_vswitch_ids` - (Optional, ForceNew, List, Available since v1.267.0) List of zone-redundant virtual switch IDs.  
This parameter is required when RedundancyType is set to ZRS. You must specify three switch IDs from three different zones.
* `resource_group_id` - (Optional, Computed, Available since v1.236.0) The ID of the new resource group.  
You can view the resource group ID in the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups?).
* `smb_acl` - (Optional, Computed, Set, Available since v1.248.0) SMB ACL. See [`smb_acl`](#smb_acl) below.
* `snapshot_id` - (Optional, Available since v1.236.0) Snapshot ID.

Supported only for Extreme NAS with the advanced storage specification.  

-> **NOTE:**  When you create a file system from a snapshot, the version of the created file system matches the version of the source file system of the snapshot. For example, if the source file system of the snapshot is version 1 and you need to create a version 2 file system, first create a file system A from the snapshot. Then create another file system B that meets the version 2 configuration, copy the data from file system A to file system B, and migrate your workloads to file system B after the copy completes.  


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `storage_type` - (Required, ForceNew) Storage type.
  - When FileSystemType=standard, valid values are: Performance, Capacity, and Premium.
  - When FileSystemType=extreme, valid values are: standard and advance.
  - When FileSystemType=cpfs, valid values are: advance_100 (100 MB/s/TiB baseline), advance_200 (200 MB/s/TiB baseline), and economic.
  - When FileSystemType=cpfsse, the only valid value is: advance_100 (100 MB/s/TiB baseline).
* `tags` - (Optional, Map, Available since v1.153.0) An array of tags.  
The array length must be between 1 and 20. If multiple tag objects are included in the array, the tag keys (Key) must be unique.
* `vswitch_id` - (Optional, ForceNew, Available since v1.153.0) VSwitch ID.
  - When FileSystemType=cpfs, you must specify this parameter.
  - When FileSystemType is not cpfs, this parameter is reserved for future use and currently has no effect. You do not need to configure it.
* `vpc_id` - (Optional, ForceNew, Available since v1.153.0) The ID of the Virtual Private Cloud (VPC).
  - This parameter is required when FileSystemType is set to cpfs or cpfsse.
  - When FileSystemType is set to standard or extreme, this parameter is reserved for future use and currently has no effect. You do not need to configure it.
* `zone_id` - (Optional, ForceNew, Computed) Zone ID.

A zone is a physical area within the same region that has independent power and network infrastructure.

When FileSystemType=standard, this parameter is optional. If not specified, the system randomly selects a zone that meets the requirements based on ProtocolType and StorageType.

When FileSystemType=extreme or FileSystemType=cpfs, this parameter is required.

-> **NOTE:**   - File systems can communicate with ECS instances across different zones within the same region.

-> **NOTE:**   - We recommend deploying your file system and ECS instances in the same zone to avoid latency caused by cross-zone communication.


### `nfs_acl`

The nfs_acl supports the following:
* `enabled` - (Optional, Computed) Indicates whether the NFS ACL feature is enabled.  
  - true: The NFS ACL feature is enabled.  
  - false: The NFS ACL feature is disabled.

### `options`

The options supports the following:
* `enable_oplock` - (Optional, Computed) Specifies whether to enable the OpLock feature.  
Valid values:  
  - true: Enables the feature.  
  - false: Disables the feature.  

-> **NOTE:**  This parameter applies only to file systems that use the SMB protocol.

* `vsc_access_point_access_only` - (Optional, Available since v1.278.0) Specifies whether to restrict access to the intelligent computing VSC mount point so that it can be accessed only through an access point (AP). This parameter applies only to intelligent computing CPFS file systems.

### `recycle_bin`

The recycle_bin supports the following:
* `reserved_days` - (Optional, Computed, Int) Retention period for files in the Recycle Bin. Unit: days.
Valid values: 1 to 180.
Default value: 3.  
* `status` - (Optional, Computed) Recycle Bin status.

Valid values:
  - Enable: The Recycle Bin is enabled.
  - Disable: The Recycle Bin is disabled.

### `smb_acl`

The smb_acl supports the following:
* `enable_anonymous_access` - (Optional, Computed, Available since v1.248.0) Specifies whether anonymous access is allowed.
  - true: Anonymous access is allowed.
  - false (default): Anonymous access is not allowed.
* `enabled` - (Optional, Computed, Available since v1.248.0) Specifies whether to enable the SMB AD ACL feature.
  - true: Enables the SMB AD ACL feature.
  - false: Disables the SMB AD ACL feature.
* `encrypt_data` - (Optional, Available since v1.248.0) Specifies whether to enable transport encryption.
  - true: Enables transport encryption.
  - false (default): Disables transport encryption.
* `home_dir_path` - (Optional, Available since v1.248.0) The home directory path for each user. The file path must adhere to the following rules:
  - Use forward slashes (/) or backslashes (\) as separators.
  - Each path segment must not contain any of the following characters: `":|?*`.
  - The length of each path segment must be between 0 and 255 characters.
  - The total path length must be between 0 and 32767 characters.

For example, if the home directory is `/home`, the file system automatically creates the directory `/home/A` when user A logs in. If `/home/A` already exists, the creation step is skipped.

-> **NOTE:**  User A must have permission to create directories; otherwise, the `/home/A` directory cannot be created.

* `reject_unencrypted_access` - (Optional, Available since v1.248.0) Specifies whether to reject unencrypted clients.
  - true: Rejects unencrypted clients.
  - false (default): Does not reject unencrypted clients.
* `super_admin_sid` - (Optional) The ID of the super administrator. The ID must follow these rules:
  - It must start with `S`, and no other letters are allowed after the initial `S`.
  - It must contain at least three hyphens (-) separating components.

Examples: `S-1-5-22` or `S-1-5-22-23`.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time of the file system.
* `recycle_bin` - Recycle Bin.
  * `enable_time` - Time when the Recycle Bin was enabled.
  * `secondary_size` - Storage usage of infrequent access data in the Recycle Bin.
  * `size` - Storage usage of files in the Recycle Bin.
* `region_id` - The region ID where the zone is located.
* `status` - The status of the file system.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the File System.
* `delete` - (Defaults to 20 mins) Used when delete the File System.
* `update` - (Defaults to 10 mins) Used when update the File System.

## Import

File Storage (NAS) File System can be imported using the id, e.g.

```shell
$ terraform import alicloud_nas_file_system.example <file_system_id>
```