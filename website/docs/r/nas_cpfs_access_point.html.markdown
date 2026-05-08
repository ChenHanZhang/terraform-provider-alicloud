---
subcategory: "File Storage (NAS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nas_cpfs_access_point"
description: |-
  Provides a Alicloud File Storage (NAS) Cpfs Access Point resource.
---

# alicloud_nas_cpfs_access_point

Provides a File Storage (NAS) Cpfs Access Point resource.

CPFS access point.

For information about File Storage (NAS) Cpfs Access Point and how to use it, see [What is Cpfs Access Point](https://next.api.alibabacloud.com/document/NAS/2017-06-26/CreateCpfsAccessPoint).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Description.
* `file_system_id` - (Required, ForceNew) File system ID.
* `root_directory` - (Optional, ForceNew, Set) Root directory of the access point. See [`root_directory`](#root_directory) below.

### `root_directory`

The root_directory supports the following:
* `root_path` - (Optional, ForceNew) Root directory of the access point.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<file_system_id>:<access_point_id>`.
* `access_point_id` - A resource property field representing the primary ID of the resource.
* `create_time` - The time when the access point was created.
* `region_id` - A resource property field representing the region.
* `root_directory` - Root directory of the access point.
  * `root_path_status` - Current status of the root directory.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cpfs Access Point.
* `delete` - (Defaults to 5 mins) Used when delete the Cpfs Access Point.
* `update` - (Defaults to 5 mins) Used when update the Cpfs Access Point.

## Import

File Storage (NAS) Cpfs Access Point can be imported using the id, e.g.

```shell
$ terraform import alicloud_nas_cpfs_access_point.example <file_system_id>:<access_point_id>
```