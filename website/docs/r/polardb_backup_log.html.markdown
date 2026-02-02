---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_backup_log"
description: |-
  Provides a Alicloud Polardb Backup Log resource.
---

# alicloud_polardb_backup_log

Provides a Polardb Backup Log resource.

Backup log list.

For information about Polardb Backup Log and how to use it, see [What is Backup Log](https://next.api.alibabacloud.com/document/polardb/2017-08-01/ModifyLogBackupPolicy).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_polardb_backup_log" "default" {
}
```

### Deleting `alicloud_polardb_backup_log` or removing it from your configuration

Terraform cannot destroy resource `alicloud_polardb_backup_log`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `db_cluster_id` - (Required, ForceNew) The ID of the cluster.

-> **NOTE:**  You can call the [DescribeDBClusters](~~ 98094 ~~) operation to view all cluster information in the target region, including cluster ID.

* `log_backup_another_region_region` - (Optional) Log backup region
* `log_backup_another_region_retention_period` - (Optional) Log backup retention period
* `log_backup_retention_period` - (Optional, Int) Log retention period. The value range is as follows:
  - 3~7300: log retention period, unit Day.
  - - 1: Permanently retained.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Backup Log.
* `update` - (Defaults to 5 mins) Used when update the Backup Log.

## Import

Polardb Backup Log can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_backup_log.example <db_cluster_id>
```