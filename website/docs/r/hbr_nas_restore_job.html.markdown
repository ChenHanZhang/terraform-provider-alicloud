---
subcategory: "Hybrid Backup Recovery (HBR)"
layout: "alicloud"
page_title: "Alicloud: alicloud_hbr_nas_restore_job"
sidebar_current: "docs-alicloud-resource-hbr-nas-restore-job"
description: |-
  Provides a Alicloud Hybrid Backup Recovery (HBR) Nas Restore Job resource.
---

# alicloud\_hbr\_nas\_restore\_job

Provides a Hybrid Backup Recovery (HBR) Nas Restore Job resource.

For information about Hybrid Backup Recovery (HBR) Nas Restore Job and how to use it, see [What is Nas Restore Job](https://help.aliyun.com/).

-> **NOTE:** Available in v1.132.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_hbr_nas_restore_job" "example" {
  restore_type = "ECS_FILE"
  source_type  = "example_value"
}

```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional, ForceNew) The Cluster ID.
* `exclude` - (Optional) The exclude.
* `include` - (Optional) The include.
* `options` - (Optional, ForceNew) Recovery Options.
* `restore_type` - (Required, ForceNew) The Type of Recovery. Valid values: `ECS_FILE`, `NAS`, `OSS`.
* `snapshot_hash` - (Optional, ForceNew) Snapshot Hash.
* `snapshot_id` - (Optional, ForceNew) No Data Disk Snapshot ID.
* `source_type` - (Required, ForceNew) Type.
* `target_bucket` - (Optional) The target bucket.
* `target_client_id` - (Optional, ForceNew) The Target Client ID.
* `target_container` - (Optional) The target container.
* `target_container_cluster_id` - (Optional) The target container cluster id.
* `target_create_time` - (Optional, ForceNew) The Destination File System Creation Time.
* `target_data_source_id` - (Optional, ForceNew) The Destination ID.
* `target_file_system_id` - (Optional, ForceNew) The Destination File System ID.
* `target_instance_id` - (Optional) The target instance id.
* `target_path` - (Optional) The target path.
* `target_prefix` - (Optional) The target prefix.
* `udm_region_id` - (Optional) The udm region id.
* `vault_id` - (Optional, ForceNew) Warehouse ID.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Nas Restore Job.
* `status` - State.

## Import

Hybrid Backup Recovery (HBR) Nas Restore Job can be imported using the id, e.g.

```
$ terraform import alicloud_hbr_nas_restore_job.example <id>
```