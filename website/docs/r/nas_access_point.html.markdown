---
subcategory: "File Storage (NAS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nas_access_point"
description: |-
  Provides a Alicloud File Storage (NAS) Access Point resource.
---

# alicloud_nas_access_point

Provides a File Storage (NAS) Access Point resource.

Access Point  .

For information about File Storage (NAS) Access Point and how to use it, see [What is Access Point](https://www.alibabacloud.com/help/zh/nas/developer-reference/api-nas-2017-06-26-createaccesspoint).

-> **NOTE:** Available since v1.224.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "azone" {
  default = "cn-hangzhou-g"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

resource "alicloud_vpc" "defaultkyVC70" {
  cidr_block  = "172.16.0.0/12"
  description = "接入点测试noRootDirectory"
}

resource "alicloud_vswitch" "defaultoZAPmO" {
  vpc_id     = alicloud_vpc.defaultkyVC70.id
  zone_id    = data.alicloud_zones.default.zones.0.id
  cidr_block = "172.16.0.0/24"
}

resource "alicloud_nas_access_group" "defaultBbc7ev" {
  access_group_type = "Vpc"
  access_group_name = "${var.name}-${random_integer.default.result}"
  file_system_type  = "standard"
}

resource "alicloud_nas_file_system" "defaultVtUpDh" {
  storage_type     = "Performance"
  zone_id          = var.azone
  encrypt_type     = "0"
  protocol_type    = "NFS"
  file_system_type = "standard"
  description      = "AccessPointnoRootDirectory"
}

resource "alicloud_nas_access_point" "default" {
  vpc_id            = alicloud_vpc.defaultkyVC70.id
  access_group      = alicloud_nas_access_group.defaultBbc7ev.access_group_name
  vswitch_id        = alicloud_vswitch.defaultoZAPmO.id
  file_system_id    = alicloud_nas_file_system.defaultVtUpDh.id
  access_point_name = var.name
  posix_user {
    posix_group_id = "123"
    posix_user_id  = "123"
  }
  root_path_permission {
    owner_group_id = "1"
    owner_user_id  = "1"
    permission     = "0777"
  }
}
```

## Argument Reference

The following arguments are supported:
* `access_group` - (Required) Permission group name.
This parameter is required when the target file system is a General-purpose NAS.
Default permission group: DEFAULT_VPC_GROUP_NAME (default permission group for VPC).
* `access_point_name` - (Optional) Access point name.
* `enabled_ram` - (Optional) Specifies whether RAM policies are enabled.  
Valid values:
  - true: Enabled
  - false (default): Disabled

-> **NOTE:**  After enabling RAM policies for the access point, all RAM users are denied access to mount or access data through this access point by default. You must explicitly grant the required permissions to specific RAM users to allow them to mount and access data via the access point. If disabled, anonymous mounting is allowed. For more information about how to configure access point policies, see [Configure Access Point Policies](https://help.aliyun.com/document_detail/2545998.html).

* `file_system_id` - (Required, ForceNew) File system ID.  
* `posix_user` - (Optional, ForceNew, Computed, Set) The POSIX user. See [`posix_user`](#posix_user) below.
* `root_path` - (Optional, ForceNew, Computed) The root directory of the access point.  
The default value is "/". If the specified directory does not exist, you must also specify the OwnerUserId and OwnerGroupId parameters.
* `root_path_permission` - (Optional, ForceNew, Computed, Set) Root directory permissions. See [`root_path_permission`](#root_path_permission) below.
* `tags` - (Optional, Map, Available since v1.274.0) List of access point tags.  
* `vswitch_id` - (Required, ForceNew) The ID of the vSwitch.
* `vpc_id` - (Required, ForceNew) VPC ID.
It must be the same VPC as the ECS instance to which the file system will be mounted.

### `posix_user`

The posix_user supports the following:
* `posix_group_id` - (Optional, ForceNew, Computed, Int) POSIX group ID.
* `posix_user_id` - (Optional, ForceNew, Computed, Int) POSIX user ID.

### `root_path_permission`

The root_path_permission supports the following:
* `owner_group_id` - (Optional, ForceNew, Computed, Int) Owner group ID.  
This parameter is required when the RootDirectory does not exist.  
* `owner_user_id` - (Optional, ForceNew, Computed, Int) Owner user ID.
This parameter is required if the RootDirectory does not exist.
* `permission` - (Optional, ForceNew) POSIX permission. Default value is "0755". Constraint: It must be a four-digit octal number starting with 0.  
This field takes effect only after the OwnerUserId and OwnerGroupId parameters are specified.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<file_system_id>:<access_point_id>`.
* `access_point_id` - Access point ID.
* `create_time` - The time when the access point was created.
* `posix_user` - The POSIX user.
  * `posix_secondary_group_ids` - Secondary groups.
* `region_id` - Region ID.
* `status` - The current status of the access point.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Access Point.
* `delete` - (Defaults to 5 mins) Used when delete the Access Point.
* `update` - (Defaults to 5 mins) Used when update the Access Point.

## Import

File Storage (NAS) Access Point can be imported using the id, e.g.

```shell
$ terraform import alicloud_nas_access_point.example <file_system_id>:<access_point_id>
```