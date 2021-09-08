---
subcategory: "DTS"
layout: "alicloud"
page_title: "Alicloud: alicloud_dts_subscription_job"
sidebar_current: "docs-alicloud-resource-dts-subscription-job"
description: |-
  Provides a Alicloud DTS Subscription Job resource.
---

# alicloud\_dts\_subscription\_job

Provides a DTS Subscription Job resource.

For information about DTS Subscription Job and how to use it, see [What is Subscription Job](https://help.aliyun.com/).

-> **NOTE:** Available in v1.135.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_dts_subscription_job" "example" {}

```

## Argument Reference

The following arguments are supported:

* `destination_region` - (Required, ForceNew) The destination region.
* `destination_endpoint_engine_name` - (Required, ForceNew) The destination endpoint engine name. Valid values: `ADS`, `DB2`, `DRDS`, `DataHub`, `Greenplum`, `MSSQL`, `MySQL`, `PolarDB`, `PostgreSQL`, `Redis`, `Tablestore`, `as400`, `clickhouse`, `kafka`, `mongodb`, `odps`, `oracle`, `polardb_o`, `polardb_pg`, `tidb`.
* `payment_type` - (Required, ForceNew) The payment type of the resource.
* `source_endpoint_region` - (Required) The source instance region.
* `source_endpoint_engine_name` - (Required) The source database type value is MySQL or Oracle. Valid values: `MySQL`, `Oracle`.
* `auto_pay` - (Optional) The auto pay. Valid values: `false`, `true`. Default value: `false`.
* `auto_start` - (Optional) The auto start. Valid values: `false`, `true`. Default value: `false`.
* `checkpoint` - (Optional) Subscription start time in Unix timestamp format.
* `compute_unit` - (Optional) The compute unit.
* `database_count` - (Optional) The database count.
* `db_list` - (Optional) Subscription object, in the format of JSON strings.
* `delay_notice` - (Optional) The delay notice.
* `delay_phone` - (Optional) The delay phone.
* `delay_rule_time` - (Optional) The delay rule time.
* `dts_instance_id` - (Optional, ForceNew) Subscription instance ID.
* `dts_job_name` - (Optional) Subscription task name.
* `error_notice` - (Optional) The error notice.
* `error_phone` - (Optional) The error phone.
* `instance_class` - (Optional) The instance class. Valid values: `large`, `medium`, `micro`, `small`, `xlarge`, `xxlarge`.
* `job_id` - (Optional) The job id.
* `period` - (Optional) The period. Valid values: `Month`, `Year`.
* `quantity` - (Optional) The quantity.
* `reserve` - (Optional) The reserve.
* `source_endpoint_database_name` - (Optional) To subscribe to the name of the database.
* `source_endpoint_instance_id` - (Optional) SOURCE instance ID, source the instance type of a database, which is RDS MySQL, PolarDB-X 1.0, PolarDB MySQL, this parameter is only available and must be set.
* `source_endpoint_instance_type` - (Optional) Source data type values include: RDS:RDS MySQL. PolarDB:PolarDB MySQL. And writes data into DRDS by running: PolarDB-X 1.0. LocalInstance: have a public IP ADDRESSES of Alibaba. ECS INSTANCE: log on to ECS and the Alibaba. Express: through special lines shall be operated for access of the user-created database is backed. CEN: through cloud enterprise network CEN access Alibaba Cloud. dg: Database Gateway of Alibaba Cloud. Valid values: `CEN`, `DRDS`, `ECS`, `Express`, `LocalInstance`, `PolarDB`, `RDS`, `dg`.
* `source_endpoint_ip` - (Optional) The source endpoint ip.
* `source_endpoint_oracle_sid` - (Optional) Oracle Database SID information, Oracle Database SID information.
* `source_endpoint_owner_id` - (Optional) Both of the Alibaba Cloud account ID of the peer, only in the configure cross-Alibaba Cloud account and remain in the data when the subscription is this parameter is only available and must be set.
* `source_endpoint_password` - (Optional) SOURCE instance database account password.
* `source_endpoint_port` - (Optional) SOURCE database port of the source database of service port.
* `source_endpoint_role` - (Optional) Both the authorization roles. When the source instance and configure subscriptions task of the Alibaba Cloud account is not the same as the need to pass the parameter, to specify the source of the authorization roles, to allow configuration subscription task of the Alibaba Cloud account to access the source of the source instance information.
* `source_endpoint_user_name` - (Optional) SOURCE instance of a database account create a database account.
* `status` - (Optional, Computed) The status of the task. Valid values: `Abnormal`, `Downgrade`, `Locked`, `Normal`, `NotStarted`, `NotStarted`, `PreCheckPass`, `PrecheckFailed`, `Prechecking`, `Retrying`, `Starting`, `Upgrade`.
* `subscription_data_type_ddl` - (Optional, Computed) Is subscribed to the DDL types of data value: true: Yes, is the default values. false: no.
* `subscription_data_type_dml` - (Optional, Computed) Whether the subscription DML type of data value: true: Yes, is the default values. false: no.
* `subscription_instance_network_type` - (Optional) Subscription task type of network value: classic: classic Network. Virtual Private Cloud (vpc): a vpc. Valid values: `classic`, `vpc`.
* `subscription_instance_vpc_id` - (Optional) The subscription instance vpc id.
* `subscription_instance_vswitch_id` - (Optional) The subscription instance vswitch id.
* `sync_architecture` - (Optional) The sync architecture. Valid values: `bidirectional`, `oneway`.
* `synchronization_direction` - (Optional) The synchronization direction.
* `used_time` - (Optional) The used time.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of Subscription Job.

## Import

DTS Subscription Job can be imported using the id, e.g.

```
$ terraform import alicloud_dts_subscription_job.example <id>
```