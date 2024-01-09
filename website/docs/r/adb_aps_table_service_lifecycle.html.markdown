---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_aps_table_service_lifecycle"
description: |-
  Provides a Alicloud ADB Aps Table Service Lifecycle resource.
---

# alicloud_adb_aps_table_service_lifecycle

Provides a ADB Aps Table Service Lifecycle resource. 

For information about ADB Aps Table Service Lifecycle and how to use it, see [What is Aps Table Service Lifecycle](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.215.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "db_cluster_id" {
  default = "amv-bp1u30028ta370f7"
}


resource "alicloud_adb_aps_table_service_lifecycle" "default" {
  status         = "on"
  strategy_type  = "KEEP_BY_TIME"
  strategy_name  = "Lifecycle-202312270000001"
  strategy_value = "100"
  db_cluster_id  = var.db_cluster_id
  operation_tables {
    table_names   = ["schema001_tb"]
    database_name = "schema001_db"
  }
  strategy_desc = "100天生命周期管理"
}
```

## Argument Reference

The following arguments are supported:
* `db_cluster_id` - (Required, ForceNew) ADB instance name.
* `operation_tables` - (Required) Lifecycle Management Processing Table List. See [`operation_tables`](#operation_tables) below.
* `status` - (Optional, Computed) Lifecycle Management Workload Status.
* `strategy_desc` - (Optional) Description of lifecycle management Workload.
* `strategy_name` - (Required, ForceNew) Lifecycle Management Name.
* `strategy_type` - (Required, ForceNew) Types of lifecycle management workloads.
* `strategy_value` - (Required) Lifecycle Management Policy Values.

### `operation_tables`

The operation_tables supports the following:
* `database_name` - (Required) Name that uniquely identifies the database.
* `process_all` - (Required) Used to identify whether to process all tables of the current library.
* `table_names` - (Optional) List of processed table names.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<db_cluster_id>:<aps_job_id>`.
* `aps_job_id` - The first ID of the resource.
* `create_time` - Lifecycle management Workload creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Aps Table Service Lifecycle.
* `delete` - (Defaults to 5 mins) Used when delete the Aps Table Service Lifecycle.
* `update` - (Defaults to 5 mins) Used when update the Aps Table Service Lifecycle.

## Import

ADB Aps Table Service Lifecycle can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_aps_table_service_lifecycle.example <db_cluster_id>:<aps_job_id>
```