---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_pipeline_job"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Pipeline Job resource.
---

# alicloud_adb_pipeline_job

Provides a AnalyticDB for MySQL (ADB) Pipeline Job resource.

ADB Pipeline Job resource definition.

For information about AnalyticDB for MySQL (ADB) Pipeline Job and how to use it, see [What is Pipeline Job](https://next.api.alibabacloud.com/document/adb/2021-12-01/CreatePipelineJob).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_vpc" "createVpc_sls_iceberg_external" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-iceberg-external-vpc"
}

resource "alicloud_log_project" "create_sls_project_iceberg_external" {
  description = "SLS project for Iceberg external lake example"
  name        = "sls-iceberg-external-project"
}

resource "alicloud_vswitch" "createVswitch_sls_iceberg_external" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_iceberg_external.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.5.0/24"
  vswitch_name = "sls-iceberg-external-vsw"
}

resource "alicloud_log_store" "create_sls_logstore_iceberg_external" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_iceberg_external.id
  name             = "sls-iceberg-external-logstore"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_iceberg_external" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_iceberg_external.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_iceberg_external.id
  db_cluster_description        = "sls_iceberg_external_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


resource "alicloud_adb_pipeline_job" "default" {
  description = "SLS to ADB Iceberg External Lake (OSS) Pipeline Job"
  sink {
    type = "OSS"
    lake {
      table_name = "example_table_sls_iceberg_external"
      partition_specs {
        source_column = "project"
        strategy      = "Identity"
        target_column = "project"
      }
      table_creation_mode = "CREATE_NEW"
      iceberg {
        write_distribution = "auto"
        format_version     = "2"
      }
      db_name        = "example_db_sls_iceberg_external"
      storage_target = "oss://adb-example-bucket/sls-iceberg-external/"
      table_format   = "APACHE_ICEBERG"
    }
  }
  db_cluster_id = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_iceberg_external.id
  transform {
    column_map {
      type     = "STRING"
      map_name = "unique_id"
      map_type = "STRING"
      name     = "unique_id"
    }
    column_map {
      type     = "STRING"
      map_name = "consumer_group"
      map_type = "STRING"
      name     = "consumer_group"
    }
    column_map {
      type     = "STRING"
      map_name = "project"
      map_type = "STRING"
      name     = "project"
    }
    column_map {
      type     = "STRING"
      map_name = "fallbehind"
      map_type = "STRING"
      name     = "fallbehind"
    }
    column_map {
      type     = "STRING"
      map_name = "shard"
      map_type = "STRING"
      name     = "shard"
    }
    column_map {
      type     = "STRING"
      map_name = "logstore"
      map_type = "STRING"
      name     = "logstore"
    }
    dirty_data_handle_mode = "ERROR"
  }
  source {
    type = "SLS"
    sls {
      start_offset_mode    = "earliest"
      project              = alicloud_log_project.create_sls_project_iceberg_external.id
      enable_cross_account = false
      log_store            = alicloud_log_store.create_sls_logstore_iceberg_external.name
    }
  }
  job_config {
    checkpoint_interval = "10000"
    enable_serverless   = true
    max_compute_unit    = "20"
    min_compute_unit    = "2"
    resource_group      = "serverless"
  }
}
```

## Argument Reference

The following arguments are supported:
* `advanced_config` - (Optional, ForceNew) KV key-value pairs, joined by '=', multiple pairs separated by newline
* `db_cluster_id` - (Required, ForceNew) DBCluster ID
* `description` - (Optional, ForceNew) Pipeline Job Description
* `job_config` - (Optional, ForceNew, Set) Task Configuration See [`job_config`](#job_config) below.
* `sink` - (Optional, ForceNew, Set) Sink Configuration See [`sink`](#sink) below.
* `source` - (Optional, ForceNew, Set) Source Configuration See [`source`](#source) below.
* `status` - (Optional, Computed) The status of the resource. Valid values: Created, Submitting, Running, Scaling, Stopping, Stopped, Failed
* `transform` - (Optional, ForceNew, Set) Transform Configuration See [`transform`](#transform) below.

### `job_config`

The job_config supports the following:
* `checkpoint_interval` - (Optional, ForceNew, Int) Checkpoint interval in milliseconds
* `enable_serverless` - (Optional, ForceNew) Enable serverless
* `max_compute_unit` - (Optional, ForceNew, Int) Maximum ACU count for the job. MaxComputeUnit must be greater than or equal to MinComputeUnit
* `min_compute_unit` - (Optional, ForceNew, Int) Minimum ACU count for the job. For lake jobs the minimum is 2; for warehouse jobs the minimum is 1
* `resource_group` - (Optional, ForceNew) Resource group

### `sink`

The sink supports the following:
* `lake` - (Optional, ForceNew, Set) Lake configuration See [`lake`](#sink-lake) below.
* `type` - (Required, ForceNew) Sink type. Valid values: ADB_WAREHOUSE, ADB_LAKE, OSS. ADB_WAREHOUSE indicates ADB warehouse storage; ADB_LAKE indicates ADB lake storage; OSS indicates user OSS lake storage.
* `warehouse` - (Optional, ForceNew, Set) Warehouse configuration See [`warehouse`](#sink-warehouse) below.

### `sink-lake`

The sink-lake supports the following:
* `db_name` - (Optional, ForceNew) Database name
* `iceberg` - (Optional, ForceNew, Set) Iceberg configuration See [`iceberg`](#sink-lake-iceberg) below.
* `paimon` - (Optional, ForceNew, Set) Paimon configuration See [`paimon`](#sink-lake-paimon) below.
* `partition_specs` - (Optional, ForceNew, List) Partition specifications See [`partition_specs`](#sink-lake-partition_specs) below.
* `storage_target` - (Optional, ForceNew) Storage target
* `table_creation_mode` - (Optional, ForceNew) Table creation mode
* `table_format` - (Optional, ForceNew) Table format
* `table_name` - (Optional, ForceNew) Table name

### `sink-warehouse`

The sink-warehouse supports the following:
* `db_name` - (Optional, ForceNew) Database name
* `enable_ip_white_list` - (Optional, ForceNew) Whether to add the IP whitelist of the Kafka switch network segment to ADB in the Kafka-to-warehouse scenario
* `enable_unix_timestamp_convert` - (Optional, ForceNew) Whether to enable unix timestamp conversion. Default is false.
* `password` - (Optional, ForceNew) Database password
* `table_creation_mode` - (Optional, ForceNew) Table creation mode
* `table_name` - (Optional, ForceNew) Table name
* `unix_timestamp_convert_format` - (Optional, ForceNew) Unix timestamp format. Required when EnableUnixTimestampConvert is true. Valid values: APSLiteralTimestampMilliSecond (microsecond), APSLiteralTimestampMicroSecond (millisecond), APSLiteralTimestampSecond (second)
* `user_name` - (Optional, ForceNew) Database username

### `sink-lake-iceberg`

The sink-lake-iceberg supports the following:
* `format_version` - (Optional, ForceNew) Iceberg storage format version. Valid values: V2, V3. Default: V2.
* `primary_key` - (Optional, ForceNew) Primary key columns, separated by ','. If primary key is configured, data will be written via upsert by default.
* `write_distribution` - (Optional, ForceNew) Iceberg write distribution mode. Valid values: AUTO, HASH, RANGE. Default: AUTO.

### `sink-lake-paimon`

The sink-lake-paimon supports the following:
* `primary_key` - (Optional, ForceNew) Primary key columns. Effective only when TableCreationMode is CREATE_NEW and TableType is TABLE_WITH_PK. Multiple columns are separated by ','. If primary key is configured, data will be written via upsert by default.
* `table_type` - (Optional, ForceNew) The type of Paimon table. Valid values: TABLE_WITHOUT_PK (table without primary key), TABLE_WITH_PK (table with primary key). This parameter is effective only when TableCreationMode is CREATE_NEW.

### `sink-lake-partition_specs`

The sink-lake-partition_specs supports the following:
* `bucket_num` - (Optional, ForceNew, Int) Number of buckets when strategy=Bucket (Iceberg hash partitioning)
* `source_column` - (Required, ForceNew) Name of the column used for partitioning. Supported partition strategies depend on the column type.
* `source_type_format` - (Optional, ForceNew) Source type format (required when strategy=ParseAsTimeAndFormat). Valid values: APSLiteralTimestampMilliSecond, APSLiteralTimestampMicroSecond, APSLiteralTimestampSecond
* `strategy` - (Required, ForceNew) Partition strategy. Valid values: Identity, Year/Month/Day/Hour, Bucket, Truncate, ParseAsTimeAndFormat.
* `target_column` - (Optional, ForceNew) Target partition column name. Required for custom time partition strategy to specify the directory field name.
* `target_type_format` - (Optional, ForceNew) Target type format (required when strategy=ParseAsTimeAndFormat), e.g., yyyy-MM-dd, yyyyMMdd
* `truncate_width` - (Optional, ForceNew, Int) Truncate width when strategy=Truncate (Iceberg truncate partitioning).

### `source`

The source supports the following:
* `kafka` - (Optional, ForceNew, Set) Kafka source configuration See [`kafka`](#source-kafka) below.
* `sls` - (Optional, ForceNew, Set) SLS source configuration See [`sls`](#source-sls) below.
* `type` - (Required, ForceNew) Source type

### `source-kafka`

The source-kafka supports the following:
* `cdc_config` - (Optional, ForceNew, Set) CDC configuration See [`cdc_config`](#source-kafka-cdc_config) below.
* `cloud_managed` - (Optional, ForceNew, Set) Cloud managed configuration See [`cloud_managed`](#source-kafka-cloud_managed) below.
* `instance_type` - (Optional, ForceNew) Kafka instance type. Valid values: CLOUD_MANAGED (Alibaba Cloud Kafka instance), NETWORK_GATEWAY (instance connected via gateway, including on-prem self-built and other clouds connected via NAT gateway)
* `kafka_cluster_id` - (Optional, ForceNew) Kafka cluster ID
* `kafka_topic` - (Optional, ForceNew) Kafka topic name
* `message_format` - (Optional, ForceNew) Kafka message format. Supported formats: json, general_canal_json, mongo_canal_json, dataworks_json, shareplex_json
* `network_gateway` - (Optional, ForceNew, Set) Network gateway configuration See [`network_gateway`](#source-kafka-network_gateway) below.
* `start_offset_mode` - (Optional, ForceNew) Start offset mode. When Sink.Type is ADB_WAREHOUSE, only timestamp is supported
* `start_offset_value` - (Optional, ForceNew) Start offset value

### `source-sls`

The source-sls supports the following:
* `across_role` - (Optional, ForceNew) Cross account role name
* `across_uid` - (Optional, ForceNew) Cross account UID
* `enable_cross_account` - (Optional, ForceNew) Enable cross account
* `log_store` - (Optional, ForceNew) SLS LogStore name
* `project` - (Optional, ForceNew) SLS Project name
* `start_offset_mode` - (Optional, ForceNew) Start offset mode. When Sink.Type is ADB_WAREHOUSE, only timestamp is supported
* `start_offset_value` - (Optional, ForceNew) Start offset value, required when startOffsetMode=timestamp

### `source-kafka-cdc_config`

The source-kafka-cdc_config supports the following:
* `sql_types` - (Optional, ForceNew) SQL types
* `sync_tables` - (Optional, ForceNew) Tables to sync

### `source-kafka-cloud_managed`

The source-kafka-cloud_managed supports the following:
* `across_role` - (Optional, ForceNew) Cross account role name
* `across_uid` - (Optional, ForceNew) Cross account UID
* `enable_cross_account` - (Optional, ForceNew) Enable cross account

### `source-kafka-network_gateway`

The source-kafka-network_gateway supports the following:
* `bootstrap_servers` - (Optional, ForceNew) Kafka bootstrap servers
* `sasl_password` - (Optional, ForceNew) SASL password
* `sasl_username` - (Optional, ForceNew) SASL username
* `security_protocol` - (Optional, ForceNew) Security protocol
* `vpc_id` - (Optional, ForceNew) VPC ID
* `vswitch_id` - (Optional, ForceNew) VSwitch ID

### `transform`

The transform supports the following:
* `column_map` - (Optional, ForceNew, List) Column mapping list See [`column_map`](#transform-column_map) below.
* `dirty_data_handle_mode` - (Optional, ForceNew) Strategy for handling records that do not match the schema. Valid values: ERROR, STORE, SKIP, TREAT_AS_NULL. Current support: for SLS->warehouse and Kafka->warehouse, only ERROR (stop sync) and TREAT_AS_NULL are supported; for SLS->lake and Kafka->lake, ERROR/STORE/SKIP are supported.

### `transform-column_map`

The transform-column_map supports the following:
* `map_name` - (Optional, ForceNew) Target column name
* `map_type` - (Optional, ForceNew) Target column type, e.g., STRING, INT, STRUCT, ARRAY, DECIMAL
* `name` - (Optional, ForceNew) Source column name, supports JSONPath expression (e.g., user.name, user.profile)
* `type` - (Optional, ForceNew) Source column type

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<pipeline_job_id>`.
* `create_time` - The time when the pipeline job was created.
* `pipeline_job_id` - Pipeline Job ID.
* `region_id` - Region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Pipeline Job.
* `delete` - (Defaults to 5 mins) Used when delete the Pipeline Job.
* `update` - (Defaults to 5 mins) Used when update the Pipeline Job.

## Import

AnalyticDB for MySQL (ADB) Pipeline Job can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_pipeline_job.example <db_cluster_id>:<pipeline_job_id>
```