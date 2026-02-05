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
* `auto_add_new_nodes` - (Optional, Computed) Specifies whether new nodes are automatically added to this endpoint. Valid values:
  - `Enable`: New nodes are automatically added to this endpoint.
  - `Disable` (default): New nodes are not automatically added to this endpoint.
* `db_cluster_id` - (Required, ForceNew) The cluster ID.

-> **NOTE:**  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to view detailed information about all clusters under your account, including the cluster ID.

* `db_endpoint_description` - (Optional, Available since v1.201.0) Custom cluster endpoint name.
* `endpoint_config` - (Optional, Computed) Advanced configuration for the cluster endpoint, specified in JSON format. Currently supports configuring consistency level, transaction splitting, whether the primary database accepts read requests, and connection pooling.

* Configure the load balancing policy. Format: `{"LoadBalancePolicy":"load balancing policy"}`. Valid values:
   * `0`: Connection-based load balancing (default).
   * `1`: Active request-based load balancing.
* Configure the consistency level. Format: `{"ConsistLevel":"consistency level"}`. Valid values:
    * `0`: Eventual consistency.
    * `1`: Session consistency (default).
    * `2`: Global consistency.
* Configure transaction splitting. Format: `{"DistributedTransaction":"transaction splitting"}`. Valid values:
    * `on`: Enable transaction splitting (default).
    * `off`: Disable transaction splitting.
* Configure whether the primary database accepts read requests. Format: `{"MasterAcceptReads":"primary database accepts reads"}`. Valid values:
    * `on`: The primary database accepts read requests.
    * `off`: The primary database does not accept read requests (default).
* Configure connection pooling. Format: `{"ConnectionPersist":"connection pooling"}`. Valid values:
    * `off`: Disable connection pooling (default).
    * `Session`: Enable session-level connection pooling.
    * `Transaction`: Enable transaction-level connection pooling.
* Configure parallel query. Format: `{"MaxParallelDegree":"parallel query"}`. Valid values:
    * A specific degree of concurrency. Example: `"MaxParallelDegree":"2"`.
    * `off`: Disable parallel query (default).
* Configure automatic rowstore/columnstore traffic routing. Format: `{"EnableHtapImci":"automatic rowstore/columnstore traffic routing"}`. Valid values:
    * `on`: Enable automatic rowstore/columnstore traffic routing.
    * `off`: Disable automatic rowstore/columnstore traffic routing (default).
* Configure whether to enable overload protection. Format: `{"EnableOverloadThrottle":"enable overload protection"}`. Valid values:
   * `on`: Enable overload protection.
   * `off`: Disable overload protection (default).

-> **NOTE:**  * Transaction splitting, whether the primary database accepts read requests, connection pooling, and overload protection can be configured only when the read/write mode of the PolarDB for MySQL cluster endpoint is set to **ReadWrite (Automatic Read/Write Splitting)**.

-> **NOTE:**  * When the read/write mode of the PolarDB for MySQL cluster endpoint is `ReadOnly`, both **connection-based load balancing** and **active request-based load balancing** are supported. However, cluster endpoints in **ReadWrite (Automatic Read/Write Splitting)** mode support only **active request-based load balancing**.

-> **NOTE:**  * Automatic rowstore/columnstore traffic routing can be configured only when the cluster endpoint is in **ReadWrite (Automatic Read/Write Splitting)** mode, or when the endpoint is in `ReadOnly` mode and the load balancing policy is set to **active request-based load balancing**.

-> **NOTE:**  * Only PolarDB for MySQL supports setting the consistency level to global consistency.

-> **NOTE:**  * If the `ReadWriteMode` parameter is set to `ReadOnly`, the consistency level can only be set to `0`.

-> **NOTE:**  * You can configure multiple settings simultaneously, such as consistency level, transaction splitting, whether the primary database accepts read requests, and connection pooling. Example: `{"ConsistLevel":"1","DistributedTransaction":"on","ConnectionPersist":"Session","MasterAcceptReads":"on"}`.

-> **NOTE:**  * The transaction splitting setting is constrained by the consistency level setting. For example, transaction splitting cannot be enabled when the consistency level is set to `0`, but it can be enabled when the consistency level is set to `1` or `2`.

* `endpoint_type` - (Required, ForceNew) The type of the custom cluster endpoint. The value is fixed to `Custom`.
* `nodes` - (Optional, Computed) The read load nodes to be added to the endpoint. Separate multiple nodes with commas (,). By default, all nodes are selected.

-> **NOTE:**  * For PolarDB for MySQL, you must specify node IDs.

-> **NOTE:**  * For PolarDB for PostgreSQL and PolarDB for PostgreSQL (compatible with Oracle), you must specify node role names, such as `Writer,Reader1,Reader2`.

-> **NOTE:**  * When `ReadWriteMode` is set to `ReadOnly`, you can attach only one node. However, if this node fails, the endpoint may be unavailable for up to one hour. Do not use this configuration in production environments. We recommend that you select at least two nodes to improve availability.

-> **NOTE:**  * When `ReadWriteMode` is set to `ReadWrite`, you must select at least two nodes.
    * For PolarDB for MySQL, you can select any two nodes. If both nodes are read-only, write requests are routed to the primary node.
    * For PolarDB for PostgreSQL and PolarDB for PostgreSQL (compatible with Oracle), the primary node must be included.
* `read_write_mode` - (Optional, Computed) Read/write mode. Valid values:
  - `ReadWrite`: Readable and writable (automatic read/write splitting).
  - `ReadOnly` (default): Read-only.

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