---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_db_proxy"
description: |-
  Provides a Alicloud RDS Db Proxy resource.
---

# alicloud_rds_db_proxy

Provides a RDS Db Proxy resource.

Database Exclusive Agent Details.

For information about RDS Db Proxy and how to use it, see [What is Db Proxy](https://next.api.alibabacloud.com/document/Rds/2014-08-15/ModifyDBProxy).

-> **NOTE:** Available since v1.193.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}
data "alicloud_db_zones" "default" {
  engine         = "MySQL"
  engine_version = "5.6"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/16"
}
resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_db_zones.default.zones.0.id
  vswitch_name = var.name
}

resource "alicloud_security_group" "default" {
  name   = var.name
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_db_instance" "default" {
  engine                   = "MySQL"
  engine_version           = "5.7"
  instance_type            = "rds.mysql.c1.large"
  instance_storage         = "20"
  instance_charge_type     = "Postpaid"
  instance_name            = var.name
  vswitch_id               = alicloud_vswitch.default.id
  db_instance_storage_type = "local_ssd"
}

resource "alicloud_db_readonly_instance" "default" {
  zone_id               = alicloud_db_instance.default.zone_id
  master_db_instance_id = alicloud_db_instance.default.id
  engine_version        = alicloud_db_instance.default.engine_version
  instance_storage      = alicloud_db_instance.default.instance_storage
  instance_type         = alicloud_db_instance.default.instance_type
  instance_name         = "${var.name}readonly"
  vswitch_id            = alicloud_vswitch.default.id
}

resource "alicloud_rds_db_proxy" "default" {
  instance_id                          = alicloud_db_instance.default.id
  instance_network_type                = "VPC"
  vpc_id                               = alicloud_db_instance.default.vpc_id
  vswitch_id                           = alicloud_db_instance.default.vswitch_id
  db_proxy_instance_num                = 2
  db_proxy_connection_prefix           = "example"
  db_proxy_connect_string_port         = 3306
  db_proxy_endpoint_read_write_mode    = "ReadWrite"
  read_only_instance_max_delay_time    = 90
  db_proxy_features                    = "TransactionReadSqlRouteOptimizeStatus:1;ConnectionPersist:1;ReadWriteSpliting:1"
  read_only_instance_distribution_type = "Custom"
  read_only_instance_weight {
    instance_id = alicloud_db_instance.default.id
    weight      = "100"
  }
  read_only_instance_weight {
    instance_id = alicloud_db_readonly_instance.default.id
    weight      = "500"
  }
}
```

## Argument Reference

The following arguments are supported:
* `causal_consist_read_timeout` - (Optional, Available since v1.269.0) The timeout period of the consistency read. Unit: milliseconds. The default value is `10` milliseconds, value: **0 to 60000**

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `config_db_proxy_features` - (Optional, Available since v1.269.0) Set the proxy connection address to enable the proxy function, which is separated by English semicolons (;). Format: 'Function 1: Opening Situation; Function 2: Opening Situation;... ', without English semicolon (;) at the end.

Function value:
* `ReadWriteSpliting`: Read/write separation
* `ConnectionPersist`: connection pool
* `Qin`: Transaction Split
* `AZProximityAccess`: nearby access
* `Causconsiststread`: Read consistency
* `HtapFilter`:HTAP line automatically shunts

Value of opening condition:
* `1`: The function has been activated
* `0`: The function is not activated

-> **NOTE:**  - RDS PostgreSQL only supports setting `ReadWriteSpliting`.

-> **NOTE:** - the nearest access function only supports MySQL exclusive proxy.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_instance_id` - (Required, ForceNew, Available since v1.269.0) The ID of the instance.
* `db_proxy_connect_string` - (Optional, ForceNew, Available since v1.269.0) The proxy terminal connection address.
* `db_proxy_connect_string_net_type` - (Optional, ForceNew, Computed, Available since v1.269.0) The network type of the proxy connection address.
  - OuterString: Extranet
  - InnerString: intranet
* `db_proxy_endpoint_id` - (Optional, ForceNew, Computed) The ID of the backend proxy terminal.
* `db_proxy_instance_num` - (Optional, Computed, Int) The number of proxy instances that are opened.
* `db_proxy_new_connect_string` - (Optional, Available since v1.269.0) Prefix of the target database proxy connection address. Custom.

-> **NOTE:**  At least one of `DBProxyNewConnectString` and `DBProxyNewConnectStringPort` is passed in.

* `db_endpoint_aliases` - (Optional, Available since v1.269.0) The comment information of the agent terminal.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_endpoint_min_slave_count` - (Optional, Available since v1.269.0) Minimum number of reserved instances.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_endpoint_operator` - (Optional, Available since v1.269.0) Operation type, value:
  - `Modify`: The default value. Modify the proxy terminal.
  - `Create`: Creates a proxy terminal.
  - `Delete`: deletes the proxy terminal.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_endpoint_read_write_mode` - (Optional, Available since v1.269.0) Read/write type, value:
* `ReadWrite`: connects to the Master instance and accepts write requests.
* `ReadOnly`: The default value. The primary instance is not connected and cannot accept write requests.

-> **NOTE:**  * When the value of `DbEndpointOperator` is `Create`, this parameter must be specified.

-> **NOTE:**  * In the RDS MySQL instance, when the value of this parameter is changed from `ReadWrite` to `ReadOnly`, the transaction splitting function is disabled.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_proxy_instance_type` - (Optional, Computed, Available since v1.230.0) Database proxy instance type, value:
  - `common`: General Purpose Agent
  - `exclusive`: exclusive proxy (default)
* `db_proxy_new_connect_string_port` - (Optional, Available since v1.269.0) Port of the destination database proxy connection address. Custom.

-> **NOTE:**  At least one of `DBProxyNewConnectString` and `DBProxyNewConnectStringPort` is passed in.

* `db_proxy_ssl_enabled` - (Optional, Computed) The operation to perform SSL encryption. The value is:
* 0: disable SSL encryption
* 1: Enable SSL encryption or modify the SSL encryption address
* 2: Update Certificate Validity Period

-> **NOTE:**  Performing the above operations will restart the instance. Please be cautious.

* `effective_specific_time` - (Optional, Computed) The specified time takes effect. Format:  yyyy-MM-dd T  HH:mm:ss Z(UTC time).

-> **NOTE:**  When `EffectiveTime` is set to `SpecificTime`, this parameter must be set.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `effective_time` - (Optional, Computed) Effective time, value:

  - `Immediate`: Effective immediately.
  - `MaintainTime`: takes effect during the O & M period. For details, see ModifyDBInstanceMaintainTime.
  - `SpecificTime`: The specified time takes effect.

Default value: **MaintainTime * *.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `persistent_connection_status` - (Optional, Available since v1.269.0) Whether to turn on the connection hold. Value:
  - `Enabled`: open connection hold
  - `Disabled`: Turn off connection hold

-> **NOTE:** - only RDS MySQL supports this parameter.

-> **NOTE:** - The value of `ConfigDBProxyService` is `Modify` when the connection retention status is modified * *.

* `read_only_instance_distribution_type` - (Optional, Computed) Read weight assignment mode. Value:

* `Standard`: The default value, which is automatically assigned by specification weight.
* `Custom`: The Custom weight.

-> **NOTE:**  This parameter needs to be passed in only when read/write splitting is enabled. For details about the permission assignment mode, see [Read Weight Assignment](~~ 96076 ~~) for MySQL, and [Activate and Configure Database Proxy Service](~~ 418272 ~~) for PostgreSQL.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `read_only_instance_max_delay_time` - (Optional, ForceNew, Computed) The maximum latency threshold for read-only instances in read/write splitting. When the latency of a read-only instance exceeds this value, read traffic is not sent to the instance. Unit: seconds. If this parameter is not passed, the original value will be maintained. Value: `0` ~ `3600`.

-> **NOTE:** - This parameter needs to be passed in only when read/write splitting is enabled.

-> **NOTE:** - Default value: The default value is `30` seconds when the read/write attribute is read/write (read/write separation), and the default value is **-1** (disabled) when the read/write attribute is read-only.

* `read_only_instance_weight` - (Optional, Computed) Custom read weight allocation, that is, the read weight of the input Master instance and read-only instance. Increment by 100, maximum 10000, format:
  - Regular instance: '{"Master instance ID":"weight","read-only instance ID":"weight"...}'

Example: '{"rm-uf6wjk5 ****":"500","rr-tfhfgk5xxx":"200"...}'
  - RDS MySQL Cluster Edition instance: '{"read-only instance ID":"weight","DBClusterNode":{"primary node ID":"weight","secondary node ID":"weight","secondary node ID":"weight"...}}'

Example: '{"rr-tfhfgk5 ****":"200","DBClusterNode":{"rn-2z ****":"0","rn-2z ****":"400","rn-2z ****":"400"...}}'

-> **NOTE:**  `DBClusterNode` is the request information unique to the cluster edition, including the `NodeID` and `weight` of the active and standby nodes.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `vswitch_id` - (Optional) The ID of the vSwitch to which the instance belongs. You can call the DescribeDBInstanceAttribute interface to query.

-> **NOTE:**  This parameter is required when enabling the database proxy for RDS MySQL cloud disk or RDS PostgreSQL.


-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `vpc_id` - (Optional) The VPC ID of the instance. You can call the DescribeDBInstanceAttribute interface to query.

-> **NOTE:**  This parameter is required when enabling the database proxy for RDS MySQL cloud disk or RDS PostgreSQL.


-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `db_proxy_connect_string_port` - The port of the proxy connection address.
* `db_proxy_endpoint_aliases` - The comment information of the agent terminal.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Db Proxy.
* `delete` - (Defaults to 5 mins) Used when delete the Db Proxy.
* `update` - (Defaults to 5 mins) Used when update the Db Proxy.

## Import

RDS Db Proxy can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_db_proxy.example <db_instance_id>
```