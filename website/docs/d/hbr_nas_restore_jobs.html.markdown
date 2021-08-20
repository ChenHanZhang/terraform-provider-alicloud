---
subcategory: "Hybrid Backup Recovery (HBR)"
layout: "alicloud"
page_title: "Alicloud: alicloud_hbr_nas_restore_jobs"
sidebar_current: "docs-alicloud-datasource-hbr-nas-restore-jobs"
description: |-
  Provides a list of Hbr Nas Restore Jobs to the user.
---

# alicloud\_hbr\_nas\_restore\_jobs

This data source provides the Hbr Nas Restore Jobs of the current Alibaba Cloud user.

-> **NOTE:** Available in v1.132.0+.

## Example Usage

Basic Usage

```terraform
data "alicloud_hbr_nas_restore_jobs" "ids" {
  restore_type = "example_value"
  ids          = ["example_value-1", "example_value-2"]
}
output "hbr_nas_restore_job_id_1" {
  value = data.alicloud_hbr_nas_restore_jobs.ids.jobs.0.id
}
            
```

## Argument Reference

The following arguments are supported:

* `filters` - (Optional, ForceNew) The filters.
* `ids` - (Optional, ForceNew, Computed)  A list of Nas Restore Job IDs.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).
* `restore_type` - (Required, ForceNew) The Type of Recovery. Valid values: `ECS_FILE`, `NAS`, `OSS`.
* `status` - (Optional, ForceNew) State.

## Argument Reference

The following attributes are exported in addition to the arguments listed above:

* `jobs` - A list of Hbr Nas Restore Jobs. Each element contains the following attributes:
	* `actual_bytes` - The Actual Size of Snapshot.
	* `actual_items` - The Actual Number of Files.
	* `bytes_done` - Recovery Is Successful, Size.
	* `bytes_total` - The Restored Total Size.
	* `cluster_id` - The Cluster ID.
	* `complete_time` - Restore Completion Time. Unix Time in Seconds.
	* `create_time` - Creation Time.
	* `error_file` - The Error File.
	* `error_message` - Error Message.
	* `expire_time` - Restore the Expiration Time. Unix Time in Seconds.
	* `id` - The ID of the Nas Restore Job.
	* `items_done` - The Success of the Number of Files.
	* `items_total` - File the Total Number.
	* `nas_restore_job_id` - Restore Job ID.
	* `options` - Recovery Options.
	* `parent_id` - The Parent Node.
	* `progress` - The Recovery Progress 100% * 100.
	* `restore_type` - The Type of Recovery.
	* `snapshot_hash` - Snapshot Hash.
	* `snapshot_id` - No Data Disk Snapshot ID.
	* `source_type` - Type.
	* `start_time` - Restoring the Start Time. Unix Time in Seconds.
	* `status` - State.
	* `target_client_id` - The Target Client ID.
	* `target_create_time` - The Destination File System Creation Time.
	* `target_data_source_id` - The Destination ID.
	* `target_file_system_id` - The Destination File System ID.
	* `udm_detail` - The UDM Details.
	* `updated_time` - Update Time.
	* `vault_id` - Warehouse ID.