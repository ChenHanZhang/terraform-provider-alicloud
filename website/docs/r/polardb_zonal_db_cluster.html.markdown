---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_zonal_db_cluster"
description: |-
  Provides a Alicloud Polardb Zonal Db Cluster resource.
---

# alicloud_polardb_zonal_db_cluster

Provides a Polardb Zonal Db Cluster resource.

MyBase cluster of PolarDB.

For information about Polardb Zonal Db Cluster and how to use it, see [What is Zonal Db Cluster](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateDBCluster).

-> **NOTE:** Available since v1.261.0.

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
```

## Argument Reference

The following arguments are supported:
* `architecture` - (Optional, Available since v1.270.0) CPU architecture. Valid values:
  - `X86`
  - `ARM`.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `auto_renew_status` - (Optional, Available since v1.270.0) The auto-renewal status. Valid values:
  - `AutoRenewal`: Auto-renewal is enabled.
  - `Normal`: Manual renewal is required.
  - `NotRenewal`: The cluster will not be renewed.

The default value is `AutoRenewal`.

-> **NOTE:**  If you set this parameter to `NotRenewal`, the system stops sending expiration reminders and only sends a non-renewal reminder three days before expiration.

* `cluster_network_type` - (Optional, Available since v1.270.0) The cluster network type. Only Virtual Private Cloud (VPC) is supported, and the value is fixed as `VPC`.  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `creation_category` - (Optional, ForceNew, Computed) The edition. Valid values:
* `Normal`: Cluster Edition (default)
* `Basic`: Single-node Edition
* `ArchiveNormal`: X-Engine (high compression)
* `NormalMultimaster`: Multi-master Cluster Edition
* `SENormal`: Standard Edition

-> **NOTE:**  * `Basic` is supported for **MySQL 5.6**, **5.7**, **8.0**, **PostgreSQL 14**, and **Oracle Syntax Compatible 2.0**.

-> **NOTE:**  * `ArchiveNormal` and `NormalMultimaster` are supported for **MySQL 8.0**.

-> **NOTE:**  * `SENormal` is supported for **MySQL 5.6**, **5.7**, **8.0**, and **PostgreSQL 14**.

For more information about editions, see [Editions](https://help.aliyun.com/document_detail/183258.html).
* `db_cluster_description` - (Optional, ForceNew, Available since v1.270.0) The cluster description, which supports fuzzy queries.  
* `db_minor_version` - (Optional, ForceNew, Computed) The minor version number of the database engine. Valid values include:
  - **8.0.2**
  - **8.0.1**

-> **NOTE:**  This parameter takes effect only when the `DBType` parameter is `MySQL` and the `DBVersion` parameter is **8.0**.

* `db_node` - (Optional, List, Available since v1.270.0) Details of the node information.   See [`db_node`](#db_node) below.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_node_class` - (Required) Node specification. For more information, see the following documentation:
  - PolarDB for MySQL: [Compute Node Specifications](https://help.aliyun.com/document_detail/102542.html).
  - PolarDB for PostgreSQL (compatible with Oracle): [Compute Node Specifications](https://help.aliyun.com/document_detail/207921.html).
  - PolarDB for PostgreSQL: [Compute Node Specifications](https://help.aliyun.com/document_detail/209380.html).

-> **NOTE:**  - To create a Serverless cluster in PolarDB for MySQL Cluster Edition, specify **polar.mysql.sl.small**.

-> **NOTE:**  - To create a Serverless cluster in PolarDB for MySQL Standard Edition, specify **polar.mysql.sl.small.c**.

-> **NOTE:**  - To create a Serverless cluster in PolarDB for PostgreSQL (compatible with Oracle) or PolarDB for PostgreSQL, specify **polar.pg.sl.small.c**.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_node_num` - (Optional, Int, Available since v1.270.0) The number of nodes for Standard Edition and Enterprise Edition. Valid values:  
  - Standard Edition: 1 to 8 (supports 1 read/write node and up to 7 read-only nodes).  
  - Enterprise Edition: 1 to 16 (supports 1 read/write node and up to 15 read-only nodes).  

-> **NOTE:**  - Enterprise Edition has 2 nodes by default, while Standard Edition has 1 node by default.  

-> **NOTE:**  - This parameter is supported only for PolarDB for MySQL.  

-> **NOTE:**  - Changing the number of nodes for multi-primary clusters is currently not supported.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_node_target_class` - (Optional, Available since v1.270.0) The target specification for all nodes. For more information, see [Compute Node Specifications](https://help.aliyun.com/document_detail/102542.html).
* `db_type` - (Required, ForceNew) The database engine type. Valid values are as follows:
  - `MySQL`
  - `PostgreSQL`
  - `Oracle`.
* `db_version` - (Required, ForceNew) The database engine version number.  
  - Valid values for MySQL versions are as follows:  
    * **5.6**  
    * **5.7**  
    * **8.0**  
  - Valid values for PostgreSQL versions are as follows:  
    * `11`  
    * `14`  
    * `15`  
      
      
      > For PolarDB for PostgreSQL, if you need to create a Serverless cluster, only version 14 is supported.  
    
      
  - Valid values for Oracle versions are as follows:  
    * `11`  
    * `14`.  
* `duration` - (Optional, Available since v1.270.0) Specifies the auto-renewal duration for the instance. Valid values are as follows:
  - When `PeriodUnit` is `Month`, valid values are `[1,2,3,6,12]`.
  - When `PeriodUnit` is `Year`, valid values are `[1-3]`.

The default value is `1`.
* `ens_region_id` - (Optional) EnsRegionId.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `hot_standby_cluster` - (Optional, Available since v1.270.0) Whether to enable the hot standby cluster. Valid values:
  - `ON` (default): Enable storage hot standby cluster.
  - `OFF`: Disable hot standby cluster.
  - `STANDBY`: Enable hot standby cluster.
  - `EQUAL`: Enable both storage and compute hot standby.
  - `3AZ`: Strong data consistency across multiple zones.

-> **NOTE:**  `STANDBY` applies only to PolarDB for PostgreSQL.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `modify_type` - (Optional, Computed, Available since v1.270.0) The modification type. Valid values:
  - `Upgrade`: Upgrade the specification.
  - `Downgrade`: Downgrade the specification.
* `pay_type` - (Required, ForceNew) The payment type. Valid values are as follows:
  - `Postpaid`: Pay-as-you-go.
  - `Prepaid`: Subscription (yearly or monthly).
* `period` - (Optional, Available since v1.270.0) This parameter is required when the payment type is `Prepaid`. It specifies whether the prepaid cluster is billed yearly or monthly.
  - `Year`: Yearly subscription. When you select a yearly or monthly subscription, the unit is years.
  - `Month`: Monthly subscription. When you select a yearly or monthly subscription, the unit is months.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `period_unit` - (Optional, Available since v1.270.0) The unit of the renewal period. Valid values:
  - `Year`: Year.
  - `Month`: Month.

Default value: `Month`.
* `storage_auto_scale` - (Optional, Available since v1.270.0) Specifies whether auto scaling is enabled for storage in Standard Edition clusters. Valid values:  
  - Enable: Enables auto scaling for storage.  
  - Disable: Disables auto scaling for storage.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `storage_pay_type` - (Optional, ForceNew, Computed) The billing method for storage. Valid values:  
  - `Postpaid`: Pay-as-you-go (billed by capacity).  
  - `Prepaid`: Subscription (billed by allocated space).
* `storage_space` - (Optional, Int) Storage space for capacity-based billing (subscription). Unit: GB.

-> **NOTE:**  - For PolarDB MySQL Enterprise Edition, the storage space ranges from 10 to 50000.

-> **NOTE:**  - For PolarDB MySQL Standard Edition, the storage space ranges from 20 to 64000.

-> **NOTE:**  - When the storage type of the Standard Edition is ESSDAUTOPL, the storage space ranges from 40 to 64000, with a minimum step size of 10 (i.e., only values such as 40, 50, 60, and so on are allowed).

* `storage_type` - (Optional, ForceNew) Valid values for Enterprise Edition storage types:
  - `PSL5`
  - `PSL4`

Valid values for Standard Edition storage types:
  - `ESSDPL0`
  - `ESSDPL1`
  - `ESSDPL2`
  - `ESSDPL3`
  - `ESSDAUTOPL`.
* `target_minor_version` - (Optional) TargetMinorVersion.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `used_time` - (Optional) This parameter is required when the payment type is `Prepaid`.
  - When `Period` is `Month`, `UsedTime` must be an integer in the range `[1–9]`.
  - When `Period` is `Year`, `UsedTime` must be an integer in the range `[1–3]`.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `vswitch_id` - (Optional, ForceNew) Virtual switch ID.
* `vpc_id` - (Optional, ForceNew, Computed) The ID of the Virtual Private Cloud (VPC).
* `zone_id` - (Optional, ForceNew, Available since v1.270.0) The zone ID.

-> **NOTE:**  You can view available zones by using the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation.


### `db_node`

The db_node supports the following:
* `db_node_id` - (Optional, Available since v1.270.0) DBNodeId.  
* `target_class` - (Optional, Available since v1.270.0) TargetClass.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The time when the cluster was created.
* `region_id` - The region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Zonal Db Cluster.
* `delete` - (Defaults to 5 mins) Used when delete the Zonal Db Cluster.
* `update` - (Defaults to 38 mins) Used when update the Zonal Db Cluster.

## Import

Polardb Zonal Db Cluster can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_zonal_db_cluster.example <db_cluster_id>
```