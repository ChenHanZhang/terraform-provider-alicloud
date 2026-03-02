---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_resource_group"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Resource Group resource.
---

# alicloud_adb_resource_group

Provides a AnalyticDB for MySQL (ADB) Resource Group resource.

ADB MySQL cluster resource group.

For information about AnalyticDB for MySQL (ADB) Resource Group and how to use it, see [What is Resource Group](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/latest/api-doc-adb-2019-03-15-api-doc-createdbresourcegroup).

-> **NOTE:** Available since v1.195.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_adb_zones" "default" {
}
data "alicloud_resource_manager_resource_groups" "default" {
  status = "OK"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.4.0.0/16"
}
resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "10.4.0.0/24"
  zone_id      = data.alicloud_adb_zones.default.zones[0].id
  vswitch_name = var.name
}

resource "alicloud_adb_db_cluster" "default" {
  compute_resource    = "48Core192GB"
  db_cluster_category = "MixedStorage"
  db_cluster_version  = "3.0"
  db_node_class       = "E32"
  db_node_storage     = 100
  description         = var.name
  elastic_io_resource = 1
  maintain_time       = "04:00Z-05:00Z"
  mode                = "flexible"
  payment_type        = "PayAsYouGo"
  resource_group_id   = data.alicloud_resource_manager_resource_groups.default.ids.0
  security_ips        = ["10.168.1.12", "10.168.1.11"]
  vpc_id              = alicloud_vpc.default.id
  vswitch_id          = alicloud_vswitch.default.id
  zone_id             = data.alicloud_adb_zones.default.zones[0].id
  tags = {
    Created = "TF",
    For     = "example",
  }
}

resource "alicloud_adb_resource_group" "default" {
  group_name    = "TF_EXAMPLE"
  group_type    = "batch"
  node_num      = 0
  db_cluster_id = alicloud_adb_db_cluster.default.id
}
```

## Argument Reference

The following arguments are supported:
* `cluster_mode` - (Optional, Computed, Available since v1.261.0) The mode of the resource group. Valid values:  
  - `Disable` (default): Standard mode.  
  - `AutoScale`: Auto-scaling mode.  
* `cluster_size_resource` - (Optional, Computed, Available since v1.261.0) Resource specification per cluster, measured in ACUs.  
* `db_cluster_id` - (Required, ForceNew) Data Warehouse Edition cluster ID.  

-> **NOTE:**  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all Data Warehouse Edition clusters in the specified region.  

* `engine` - (Optional, ForceNew, Computed, Available since v1.261.0) The engine type of the resource group. Valid values:
  - `AnalyticDB` (default): AnalyticDB for MySQL engine.
  - `SparkWarehouse`: SparkWarehouse engine.
* `engine_params` - (Optional, Map, Available since v1.261.0) Spark application configuration parameters applied to all Spark jobs executed by this resource group. To configure parameters for a specific Spark job, you can set them in code when submitting the job.  
* `group_name` - (Required, ForceNew) The resource group name.  
  - Must be no longer than 255 characters.  
  - Must start with a digit or an uppercase letter.  
  - Can contain digits, uppercase letters, hyphens (-), and underscores (_).  
* `group_type` - (Optional, Computed) Query type. Valid values:  
  - `interactive` (default): Interactive query mode.  
  - `batch`: Batch query mode.  
  - `Job`: Offline query mode.  

-> **NOTE:**  For more information, see [Query Execution Modes](https://help.aliyun.com/document_detail/189502.html).  

* `max_cluster_count` - (Optional, Int, Available since v1.261.0) The maximum number of clusters that can run in the resource group, up to 10.
* `max_compute_resource` - (Optional, Computed, Available since v1.261.0) Maximum reserved compute resources, measured in ACUs.  
  - When the resource group type is `Interactive`, the maximum reserved compute resources cannot exceed the cluster's currently unallocated resources, with a step size of 16 ACUs.  
  - When the resource group type is `Job`, the maximum reserved compute resources cannot exceed the cluster's currently unallocated resources, with a step size of 8 ACUs.  
* `min_cluster_count` - (Optional, Int, Available since v1.261.0) The minimum number of clusters that must be running in the resource group. The minimum value is 1.  
* `min_compute_resource` - (Optional, Computed, Available since v1.261.0) Minimum reserved compute resources, measured in ACUs.  
  - When `GroupType` is `Interactive`, the minimum reserved compute resources are 16 ACUs.  
  - When `GroupType` is `Job`, the minimum reserved compute resources are 0 ACUs.  
* `node_num` - (Optional, Computed, Int) Number of nodes. Default value is 0.  
  - Each node provides 16 vCPUs and 64 GB of memory.  
  - The number of nodes must not be too large; it must satisfy the condition: number of nodes × (16 vCPUs and 64 GB) ≤ remaining available resources in the cluster.  
* `user` - (Optional, ForceNew, Computed) List of database accounts bound to the resource group, separated by commas (,).  
* `user_list` - (Optional, List, Available since v1.272.0) A list of database accounts to be bound. Both standard and privileged database accounts can be bound.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<group_name>`.
* `connection_string` - The connection string for the resource group.
* `create_time` - Creation time.
* `port` - Port number of the resource group.
* `status` - The status of the resource group.
* `update_time` - Update time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 31 mins) Used when create the Resource Group.
* `delete` - (Defaults to 31 mins) Used when delete the Resource Group.
* `update` - (Defaults to 31 mins) Used when update the Resource Group.

## Import

AnalyticDB for MySQL (ADB) Resource Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_resource_group.example <db_cluster_id>:<group_name>
```