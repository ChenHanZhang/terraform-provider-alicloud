---
subcategory: "File Storage (NAS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nas_lifecycle_policy"
description: |-
  Provides a Alicloud File Storage (NAS) Lifecycle Policy resource.
---

# alicloud_nas_lifecycle_policy

Provides a File Storage (NAS) Lifecycle Policy resource.

Lifecycle Management Strategy.

For information about File Storage (NAS) Lifecycle Policy and how to use it, see [What is Lifecycle Policy](https://www.alibabacloud.com/help/en/doc-detail/169362.html).

-> **NOTE:** Available since v1.277.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_nas_file_system" "example" {
  protocol_type = "NFS"
  storage_type  = "Capacity"
}

resource "alicloud_nas_lifecycle_policy" "example" {
  file_system_id        = alicloud_nas_file_system.example.id
  lifecycle_policy_name = "terraform-example"
  lifecycle_rule_name   = "DEFAULT_ATIME_14"
  storage_type          = "InfrequentAccess"
  paths                 = ["/"]
}
```

## Argument Reference

The following arguments are supported:
* `file_system_id` - (Required, ForceNew) The ID of the file system.
* `lifecycle_policy_name` - (Required, ForceNew) The first ID of the resource
* `lifecycle_rule_name` - (Required) Management rules associated with lifecycle management policies.

Value:
  - DEFAULT_ATIME_14: files not accessed 14 days ago
  - DEFAULT_ATIME_30: files not accessed 30 days ago
  - DEFAULT_ATIME_60: files not accessed 60 days ago
  - DEFAULT_ATIME_90: files not accessed 90 days ago
* `path` - (Optional) Absolute path to the directory associated with the lifecycle management policy.

Only associating a single directory is supported. It must start with a forward slash (/) and be the real path in the Mount point.

-> **NOTE:**  It is recommended that you configure Paths.N to associate multiple directories at the same time.

* `paths` - (Optional, ForceNew, List) Absolute path to the directory associated with the lifecycle management policy.
Supports associating multiple directories. It must start with a forward slash (/) and is the path that exists in the Mount point. The value range of N is 1 to 10.
* `storage_type` - (Required) The type of storage after the data dump.
Default value: InfrequentAccess (low frequency media storage)

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<file_system_id>:<lifecycle_policy_name>`.
* `create_time` - The creation time of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Lifecycle Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Lifecycle Policy.
* `update` - (Defaults to 5 mins) Used when update the Lifecycle Policy.

## Import

File Storage (NAS) Lifecycle Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_nas_lifecycle_policy.example <file_system_id>:<lifecycle_policy_name>
```