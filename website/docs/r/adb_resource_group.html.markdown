---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_resource_group"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Resource Group resource.
---

# alicloud_adb_resource_group

Provides a AnalyticDB for MySQL (ADB) Resource Group resource.

ResourceGroup for ADBMySQL.

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
* `cluster_mode` - (Optional, Computed, Available since v1.261.0) ClusterMode
* `cluster_size_resource` - (Optional, Computed, Available since v1.261.0) ClusterSizeResource
* `db_cluster_id` - (Required, ForceNew) DBClusterId
* `engine` - (Optional, ForceNew, Computed, Available since v1.261.0) EngineType
* `engine_params` - (Optional, Map, Available since v1.261.0) EngineParams
* `group_name` - (Required, ForceNew) The name of the resource pool, which cannot exceed 64 bytes in length.
* `group_type` - (Optional, Computed) Query type, value description:
  - `batch`: Batch query mode.
  - `interactive`: interactive Query mode.

-> **NOTE:**  For more information, see [Query execution mode](~~ 189502 ~~).

* `max_cluster_count` - (Optional, Int, Available since v1.261.0) MaxClusterCount
* `max_compute_resource` - (Optional, Computed, Available since v1.261.0) MaxComputeResource
* `min_cluster_count` - (Optional, Int, Available since v1.261.0) MinClusterCount
* `min_compute_resource` - (Optional, Computed, Available since v1.261.0) MinComputeResource
* `node_num` - (Optional, Computed, Int) The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.

-> **NOTE:**  You can use the [DescribeDBResourcePool](~~ 190594 ~~) interface to view the number of nodes whose resource name is USER_DEFAULT.

* `user` - (Optional, ForceNew, Computed) Binding User
* `user_list` - (Optional, List, Available since v1.271.0) UserList

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<group_name>`.
* `connection_string` - ConnectionString.
* `create_time` - Creation time.
* `port` - Port.
* `status` - Status of ResourceGroup.
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