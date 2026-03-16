---
subcategory: "File Storage (NAS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nas_snapshot"
description: |-
  Provides a Alicloud File Storage (NAS) Snapshot resource.
---

# alicloud_nas_snapshot

Provides a File Storage (NAS) Snapshot resource.

Extreme NAS snapshot.

For information about File Storage (NAS) Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/en/doc-detail/126549.html).

-> **NOTE:** Available since v1.152.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "testacc"
}

data "alicloud_nas_zones" "default" {
  file_system_type = "extreme"
}

resource "alicloud_nas_file_system" "default" {
  file_system_type = "extreme"
  protocol_type    = "NFS"
  zone_id          = data.alicloud_nas_zones.default.zones.0.zone_id
  storage_type     = "standard"
  capacity         = 100
}

resource "alicloud_nas_snapshot" "default" {
  file_system_id = alicloud_nas_file_system.default.id
  description    = var.name
  retention_days = 20
  snapshot_name  = var.name
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional, ForceNew) The snapshot description.
* `file_system_id` - (Required, ForceNew) The ID of the file system.
* `retention_days` - (Optional, ForceNew, Int) The retention time of the snapshot.
* `snapshot_name` - (Optional, ForceNew) The snapshot name.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `status` - Snapshot status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 20 mins) Used when create the Snapshot.
* `delete` - (Defaults to 5 mins) Used when delete the Snapshot.

## Import

File Storage (NAS) Snapshot can be imported using the id, e.g.

```shell
$ terraform import alicloud_nas_snapshot.example <snapshot_id>
```