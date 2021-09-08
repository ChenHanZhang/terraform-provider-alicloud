---
subcategory: "DTS"
layout: "alicloud"
page_title: "Alicloud: alicloud_dts_subscription_jobs"
sidebar_current: "docs-alicloud-datasource-dts-subscription-jobs"
description: |-
  Provides a list of Dts Subscription Jobs to the user.
---

# alicloud\_dts\_subscription\_jobs

This data source provides the Dts Subscription Jobs of the current Alibaba Cloud user.

-> **NOTE:** Available in v1.135.0+.

## Example Usage

Basic Usage

```terraform
data "alicloud_dts_subscription_jobs" "ids" {}
output "dts_subscription_job_id_1" {
  value = data.alicloud_dts_subscription_jobs.ids.jobs.0.id
}
            
```

## Argument Reference

The following arguments are supported:

* `enable_details` - (Optional) Default to `false`. Set it to `true` can output more details about resource attributes.
* `group_id` - (Optional, ForceNew) The group id.
* `ids` - (Optional, ForceNew, Computed)  A list of Subscription Job IDs.
* `order_column` - (Optional, ForceNew) The order column.
* `order_direction` - (Optional, ForceNew) The order direction.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).
* `params` - (Optional, ForceNew) The params.
* `region` - (Optional, ForceNew) The region.
* `status` - (Optional, ForceNew) The status of the task. Valid values: `Abnormal`, `Downgrade`, `Locked`, `Normal`, `NotStarted`, `NotStarted`, `PreCheckPass`, `PrecheckFailed`, `Prechecking`, `Retrying`, `Starting`, `Upgrade`.

## Argument Reference

The following attributes are exported in addition to the arguments listed above:

* `jobs` - A list of Dts Subscription Jobs. Each element contains the following attributes:
	* `checkpoint` - Subscription start time in Unix timestamp format.
	* `create_time` - Creation time.
	* `db_list` - Subscription object, in the format of JSON strings.
	* `dts_instance_id` - Subscription instance ID.
	* `dts_job_id` - Subscription ID.
	* `dts_job_name` - Subscription task name.
	* `expire_time` - Examples of the Expiration Time yyyy-MM-ddTHH:mm:ssZ(UTC time).
	* `id` - The ID of the Subscription Job.
	* `payment_type` - The paymen type of the resource.
	* `source_endpoint_database_name` - To subscribe to the name of the database.
	* `source_endpoint_engine_name` - The source database type value is MySQL or Oracle.
	* `source_endpoint_instance_id` - SOURCE instance ID, source the instance type of a database, which is RDS MySQL, PolarDB-X 1.0, PolarDB MySQL, this parameter is only available and must be set.
	* `source_endpoint_instance_type` - Source data type values include: RDS:RDS MySQL. PolarDB:PolarDB MySQL. And writes data into DRDS by running: PolarDB-X 1.0. LocalInstance: have a public IP ADDRESSES of Alibaba. ECS INSTANCE: log on to ECS and the Alibaba. Express: through special lines shall be operated for access of the user-created database is backed. CEN: through cloud enterprise network CEN access Alibaba Cloud. dg: Database Gateway of Alibaba Cloud.
	* `` - The source endpoint ip.
	* `source_endpoint_oracle_sid` - Oracle Database SID information, Oracle Database SID information.
	* `source_endpoint_owner_id` - Both of the Alibaba Cloud account ID of the peer, only in the configure cross-Alibaba Cloud account and remain in the data when the subscription is this parameter is only available and must be set.
	* `source_endpoint_port` - SOURCE database port of the source database of service port.
	* `source_endpoint_region` - SOURCE instance region where the source instance region where.
	* `source_endpoint_role` - Both the authorization roles. When the source instance and configure subscriptions task of the Alibaba Cloud account is not the same as the need to pass the parameter, to specify the source of the authorization roles, to allow configuration subscription task of the Alibaba Cloud account to access the source of the source instance information.
	* `source_endpoint_user_name` - SOURCE instance of a database account create a database account.
	* `status` - The status of the task.
	* `subscription_data_type_ddl` - Is subscribed to the DDL types of data value: true: Yes, is the default values. false: no.
	* `subscription_data_type_dml` - Whether the subscription DML type of data value: true: Yes, is the default values. false: no.
	* `subscription_host` - Network information.
		* `private_host` - The classic network address is.
		* `public_host` - Public network address.
		* `vpc_host` - Virtual Private Cloud (VPC network address.
	* `subscription_instance_network_type` - Subscription task type of network value: classic: classic Network. Virtual Private Cloud (vpc): a vpc.
	* `` - The subscription instance vpc id.
	* `` - The subscription instance vswitch id.
	* `tags` - The tag of the resource.
		* `tag_key` - The key of the tags.
		* `tag_value` - The value of the tags.