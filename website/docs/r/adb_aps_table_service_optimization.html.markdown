---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_aps_table_service_optimization"
description: |-
  Provides a Alicloud ADB Aps Table Service Optimization resource.
---

# alicloud_adb_aps_table_service_optimization

Provides a ADB Aps Table Service Optimization resource. 

For information about ADB Aps Table Service Optimization and how to use it, see [What is Aps Table Service Optimization](https://www.alibabacloud.com/help/en/).

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


resource "alicloud_adb_aps_table_service_optimization" "default" {
  db_cluster_id = var.db_cluster_id
}
```

### Deleting `alicloud_adb_aps_table_service_optimization` or removing it from your configuration

Terraform cannot destroy resource `alicloud_adb_aps_table_service_optimization`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `db_cluster_id` - (Required, ForceNew) The ID of the ADB instance to which the resource belongs.
* `status` - (Optional, Computed) Optimization policy on state.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Aps Table Service Optimization.
* `update` - (Defaults to 5 mins) Used when update the Aps Table Service Optimization.

## Import

ADB Aps Table Service Optimization can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_aps_table_service_optimization.example <id>
```