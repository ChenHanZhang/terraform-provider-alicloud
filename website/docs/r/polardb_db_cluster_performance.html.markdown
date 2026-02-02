---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_db_cluster_performance"
description: |-
  Provides a Alicloud Polardb Db Cluster Performance resource.
---

# alicloud_polardb_db_cluster_performance

Provides a Polardb Db Cluster Performance resource.

Cluster performance log.

For information about Polardb Db Cluster Performance and how to use it, see [What is Db Cluster Performance](https://next.api.alibabacloud.com/document/polardb/2017-08-01/DescribeDBClusterPerformance).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-beijing"
}


resource "alicloud_polardb_db_cluster_performance" "default" {
}
```

### Deleting `alicloud_polardb_db_cluster_performance` or removing it from your configuration

Terraform cannot destroy resource `alicloud_polardb_db_cluster_performance`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `db_cluster_id` - (Required, ForceNew) The ID of the cluster.
* `db_node_id` - (Optional, ForceNew) The ID of the database cluster node.
* `end_time` - (Optional, ForceNew) Query end time. The format is' yyyy-MM-ddTHH:mm:ssZ'(UTC time).
* `start_time` - (Optional, ForceNew) The start time of the query. The format is' yyyy-MM-ddTHH:mm:ssZ'(UTC time).

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 


## Import

Polardb Db Cluster Performance can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_db_cluster_performance.example <db_cluster_id>
```