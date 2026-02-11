---
subcategory: "Log Service (SLS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_log_store"
description: |-
  Provides a Alicloud Log Service (SLS) Log Store resource.
---

# alicloud_log_store

Provides a Log Service (SLS) Log Store resource.



For information about Log Service (SLS) Log Store and how to use it, see [What is Log Store](https://www.alibabacloud.com/help/doc-detail/48874.htm).

-> **NOTE:** Available since v1.0.0.

## Example Usage

Basic Usage

```terraform
resource "random_integer" "default" {
  max = 99999
  min = 10000
}

resource "alicloud_log_project" "example" {
  name        = "terraform-example-${random_integer.default.result}"
  description = "terraform-example"
}

resource "alicloud_log_store" "example" {
  project               = alicloud_log_project.example.name
  name                  = "example-store"
  shard_count           = 3
  auto_split            = true
  max_split_shard_count = 60
  append_meta           = true
}
```

## Argument Reference

The following arguments are supported:
* `append_meta` - (Optional, Computed) Specifies whether to record the **public IP address** and **log ingestion time**. Default value: false.
  - true: Enables recording of the public IP address and log ingestion time. After this feature is enabled, Log Service automatically adds the public IP address of the log source device and the time when logs arrive at the server to the log tag fields.
  - false: Disables recording of the public IP address and log ingestion time.
* `auto_split` - (Optional) Specifies whether to automatically split shards.
  - true: Automatically splits shards.
  - false: Does not automatically split shards.
* `enable_web_tracking` - (Optional) Specifies whether to enable WebTracking. Default value: false.
  - true: Enables WebTracking.
  - false: Disables WebTracking.
* `encrypt_conf` - (Optional, Computed, Set) Encryption configuration. See [`encrypt_conf`](#encrypt_conf) below.
* `hot_ttl` - (Optional, Int) The retention period for data in the hot storage tier of the Logstore. Unit: days. The minimum value is 7, and the maximum value cannot exceed the TTL value. A value of - 1 indicates that all data is stored in hot storage for the entire TTL duration.
When data exceeds the configured hot storage retention period, it is automatically transitioned to infrequent access storage. For more information, see [Intelligent Hot and Cold Tiered Storage](https://help.aliyun.com/document_detail/308645.html).
* `infrequent_access_ttl` - (Optional, Int, Available since v1.229.0) Infrequent access storage. There is no minimum storage duration requirement, but data must be retained for at least 30 days before it can be transitioned to archive storage.
* `logstore_name` - (Optional, ForceNew, Available since v1.215.0) The Logstore name. The naming rules are as follows:
  - Logstore names must be unique within the same project.
  - The name can only contain lowercase letters, digits, hyphens (-), and underscores (_).
  - The name must start and end with a lowercase letter or a digit.
  - The length must be between 2 and 63 characters.
* `max_split_shard_count` - (Optional, Int) The maximum number of shards when auto-splitting is enabled. The minimum value is 1 and the maximum value is 256.

-> **NOTE:**  This parameter is required when autoSplit is set to true.

* `metering_mode` - (Optional, Computed, Available since v1.215.0) Metering mode. Valid values: ChargeByFunction (default metering mode) and ChargeByDataIngest (ingestion-based metering mode).
* `mode` - (Optional, Computed) Log Service provides two types of Logstores: Standard and Query.
  - `standard`: Supports the all-in-one data analytics capabilities of Log Service. It is suitable for scenarios such as real-time monitoring, interactive analysis, and building comprehensive observability systems.
  - `query`: Supports high-performance queries and reduces index traffic costs by approximately half compared to Standard mode, but does not support SQL analysis. It is suitable for scenarios involving large data volumes, long storage periods (weeks or months or longer), and no log analysis requirements.
* `project_name` - (Optional, ForceNew, Available since v1.215.0) The project name.
* `retention_period` - (Optional, Computed, Int) The data retention period, in days. Valid values range from 1 to 3650. If set to 3650, data is retained permanently.
* `shard_count` - (Optional, ForceNew, Computed, Int) The number of shards.

-> **NOTE:**  This API does not support updating the shard count. You can only modify the shard count by using the SplitShard or MergeShards API.

* `tags` - (Optional, Map, Available since v1.271.0) The tag of the resource
* `telemetry_type` - (Optional, ForceNew) The observability data type. Valid values:
  - `None`: Log data. This is the default value.
  - `Metrics`: Time-series data.

The following arguments will be discarded. Please use new fields as soon as possible:
* `project` - (Deprecated since v1.271.0). Field 'project' has been deprecated from provider version 1.271.0. New field 'project_name' instead.
* `name` - (Deprecated since v1.271.0). Field 'name' has been deprecated from provider version 1.271.0. New field 'logstore_name' instead.

### `encrypt_conf`

The encrypt_conf supports the following:
* `enable` - (Optional, Computed) Specifies whether to enable log encryption.  
After the encryption configuration is successfully updated, subsequent update requests can only modify whether log encryption is enabled. You cannot modify the encryptType or userCmkInfo parameters.
* `encrypt_type` - (Optional, ForceNew, Computed) The encryption algorithm type. Supported values include default, m4, sm4_ecb, sm4_cbc, sm4_gcm, aes_ecb, aes_cbc, aes_cfb, aes_ofb, and aes_gcm.
* `user_cmk_info` - (Optional, ForceNew, Computed, Set) User encryption configuration. See [`user_cmk_info`](#encrypt_conf-user_cmk_info) below.

### `encrypt_conf-user_cmk_info`

The encrypt_conf-user_cmk_info supports the following:
* `arn` - (Optional, ForceNew, Computed) The Amazon Resource Name (ARN) of the RAM role, in the format acs:ram::12344***:role/xxxxx. To use BYOK encryption, you must first create a RAM role and grant it the AliyunKMSReadOnlyAccess and AliyunKMSCryptoUserAccess permissions. Additionally, you must grant the PassRole permission on this RAM role to the API caller.
* `cmk_key_id` - (Optional, ForceNew, Computed) The Customer Master Key (CMK) ID for Bring Your Own Key (BYOK). You can create a CMK in Alibaba Cloud Key Management Service (KMS). The region of the CMK must match the region of the Log Service endpoint.
* `region_id` - (Optional, ForceNew, Computed) The region, such as cn-hangzhou.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<project_name>:<logstore_name>`.
* `create_time` - The time when the Logstore was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Log Store.
* `delete` - (Defaults to 5 mins) Used when delete the Log Store.
* `update` - (Defaults to 5 mins) Used when update the Log Store.

## Import

Log Service (SLS) Log Store can be imported using the id, e.g.

```shell
$ terraform import alicloud_log_store.example <project_name>:<logstore_name>
```