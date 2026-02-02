---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_endpoint"
description: |-
  Provides a Alicloud Polardb Endpoint resource.
---

# alicloud_polardb_endpoint

Provides a Polardb Endpoint resource.

Public address of the PolarDB cluster primary address, default cluster address, and custom cluster address.

For information about Polardb Endpoint and how to use it, see [What is Endpoint](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateDBClusterEndpoint).

-> **NOTE:** Available since v1.80.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_polardb_node_classes" "default" {
  db_type    = "MySQL"
  db_version = "8.0"
  pay_type   = "PostPaid"
  category   = "Normal"
}

resource "alicloud_vpc" "default" {
  vpc_name   = "terraform-example"
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_polardb_node_classes.default.classes[0].zone_id
  vswitch_name = "terraform-example"
}

resource "alicloud_polardb_cluster" "default" {
  db_type       = "MySQL"
  db_version    = "8.0"
  db_node_class = data.alicloud_polardb_node_classes.default.classes.0.supported_engines.0.available_resources.0.db_node_class
  pay_type      = "PostPaid"
  vswitch_id    = alicloud_vswitch.default.id
  description   = "terraform-example"
}

resource "alicloud_polardb_endpoint" "default" {
  db_cluster_id = alicloud_polardb_cluster.default.id
  endpoint_type = "Custom"
}
```

## Argument Reference

The following arguments are supported:
* `auto_add_new_nodes` - (Optional, Computed) Specifies whether to enable automatic association of newly added nodes with the cluster endpoint. Valid values: 
  - `Enable`: enables automatic association of newly added nodes with the cluster endpoint. 
  - `Disable` (default): disables automatic association of newly added nodes with the cluster endpoint.
* `db_cluster_id` - (Required, ForceNew) The ID of the cluster.
* `db_endpoint_description` - (Optional, Available since v1.201.0) Custom cluster address name.
* `endpoint_config` - (Optional, Computed) The advanced configuration of the cluster address. The format is JSON. Currently, you can set the consistency level, transaction splitting, read-free primary database, and connection pooling.
* `endpoint_type` - (Required, ForceNew) Custom cluster address type. The value is fixed to **Custom * *.
* `nodes` - (Optional, Computed) The read load nodes that are added to the target address. Separate multiple nodes with English commas (,). The default value is all nodes.

-> **NOTE:**  * the PolarDB MySQL engine needs to pass in the node ID.
* The PolarDB PostgreSQL engine and PolarDB O engine need to pass in the node role name, such as "writer, Reader1,Reader2 '.
If the value of `ReadWriteMode` is `ReadOnly`, you can mount only one node. However, when this node fails, the address may be unavailable for up to 1 hour. Do not use it in a production environment. Therefore, it is recommended to select at least 2 nodes to improve availability.
If the value of `ReadWriteMode` is `ReadWrite`, at least two nodes must be selected.
* The PolarDB MySQL engine supports selecting any two nodes. When both nodes are read-only nodes, write requests are sent to the primary node.
* The PolarDB PostgreSQL engine and PolarDB O engine must contain the primary node.
* `read_write_mode` - (Optional, Computed) Read/write mode, the range of values is as follows:
  - `ReadWrite`: Readable and writable (automatic read/write splitting).
  - `ReadOnly`: Read-only.

The default value is **ReadOnly * *.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<db_endpoint_id>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 11 mins) Used when create the Endpoint.
* `delete` - (Defaults to 5 mins) Used when delete the Endpoint.
* `update` - (Defaults to 5 mins) Used when update the Endpoint.

## Import

Polardb Endpoint can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_endpoint.example <db_cluster_id>:<db_endpoint_id>
```