---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_zonal_endpoint"
description: |-
  Provides a Alicloud Polardb Zonal Endpoint resource.
---

# alicloud_polardb_zonal_endpoint

Provides a Polardb Zonal Endpoint resource.

PolarDB cluster link address resource on MyBase.

For information about Polardb Zonal Endpoint and how to use it, see [What is Zonal Endpoint](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateDBClusterEndpointZonal).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "db_cluster_nodes_configs" {
  description = "The advanced configuration for all nodes in the cluster except for the RW node, including db_node_class, hot_replica_mode, and imci_switch properties."
  type = map(object({
    db_node_class    = string
    db_node_role     = optional(string, null)
    hot_replica_mode = optional(string, null)
    imci_switch      = optional(string, null)
  }))
  default = {
    db_node_1 = {
      db_node_class = "polar.mysql.x4.medium.c"
      db_node_role  = "Writer"
    }
    db_node_2 = {
      db_node_class = "polar.mysql.x4.medium.c"
      db_node_role  = "Reader"
    }
  }
}

resource "alicloud_ens_network" "default" {
  network_name = "terraform-example"

  description   = "LoadBalancerNetworkDescription_test"
  cidr_block    = "192.168.2.0/24"
  ens_region_id = "tr-Istanbul-1"
}

resource "alicloud_ens_vswitch" "default" {
  description  = "LoadBalancerVSwitchDescription_test"
  cidr_block   = "192.168.2.0/24"
  vswitch_name = "terraform-example"

  ens_region_id = "tr-Istanbul-1"
  network_id    = alicloud_ens_network.default.id
}

resource "alicloud_polardb_zonal_db_cluster" "default" {
  db_node_class = "polar.mysql.x4.medium.c"
  description   = "terraform-example"
  ens_region_id = "tr-Istanbul-1"
  vpc_id        = alicloud_ens_network.default.id
  vswitch_id    = alicloud_ens_vswitch.default.id
  db_cluster_nodes_configs = {
    for node, config in var.db_cluster_nodes_configs : node => jsonencode({ for k, v in config : k => v if v != null })
  }
}

resource "alicloud_polardb_zonal_endpoint" "default" {
  db_cluster_id        = alicloud_polardb_zonal_db_cluster.default.id
  db_cluster_nodes_ids = alicloud_polardb_zonal_db_cluster.default.db_cluster_nodes_ids
  endpoint_config      = {}
  nodes_key            = ["db_node_1", "db_node_2"]
  read_write_mode      = "ReadWrite"
}
```

## Argument Reference

The following arguments are supported:
* `auto_add_new_nodes` - (Optional, ForceNew) Whether a new node is automatically added to this address, the value range is as follows:
  - `Enable`: new nodes are automatically added to this address.
  - `Disable`: new nodes are not automatically added to this address.

The default value is **Disable * *.
* `db_cluster_id` - (Required, ForceNew) DBClusterId
* `db_endpoint_description` - (Optional) Custom cluster address name.
* `endpoint_config` - (Optional) The advanced configuration of the cluster address. The format is JSON. Currently, you can set the consistency level, transaction splitting, read-free primary database, and connection pooling.
* `endpoint_type` - (Required, ForceNew) Custom cluster address type. The value is fixed to **Custom * *.
* `nodes` - (Optional) The read load nodes that are added to the target address. Separate multiple nodes with English commas (,). The default value is all nodes.

-> **NOTE:**  * the PolarDB MySQL engine needs to pass in the node ID.
* The PolarDB PostgreSQL engine and PolarDB O engine need to pass in the node role name, such as "writer, Reader1,Reader2 '.
If the value of `ReadWriteMode` is `ReadOnly`, you can mount only one node. However, when this node fails, the address may be unavailable for up to 1 hour. Do not use it in a production environment. Therefore, it is recommended to select at least 2 nodes to improve availability.
If the value of `ReadWriteMode` is `ReadWrite`, at least two nodes must be selected.
* The PolarDB MySQL engine supports selecting any two nodes. When both nodes are read-only nodes, write requests are sent to the primary node.
* The PolarDB PostgreSQL engine and PolarDB O engine must contain the primary node.
* `read_write_mode` - (Optional) Read/write mode, the range of values is as follows:
  - `ReadWrite`: Readable and writable (automatic read/write splitting).
  - `ReadOnly`: Read-only.

The default value is **ReadOnly * *.
* `resource_group_id` - (Optional, Computed) The ID of the resource group

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<db_endpoint_id>`.
* `db_endpoint_id` - The first ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Zonal Endpoint.
* `delete` - (Defaults to 5 mins) Used when delete the Zonal Endpoint.
* `update` - (Defaults to 5 mins) Used when update the Zonal Endpoint.

## Import

Polardb Zonal Endpoint can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_zonal_endpoint.example <db_cluster_id>:<db_endpoint_id>
```