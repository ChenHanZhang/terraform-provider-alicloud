---
subcategory: "AliKafka"
layout: "alicloud"
page_title: "Alicloud: alicloud_alikafka_topic"
description: |-
  Provides a Alicloud Ali Kafka Topic resource.
---

# alicloud_alikafka_topic

Provides a Ali Kafka Topic resource.

Topic in kafka.

For information about Ali Kafka Topic and how to use it, see [What is Topic](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/api-alikafka-2019-09-16-createtopic).

-> **NOTE:** Available since v1.56.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_zones.default.zones.0.id
}

resource "alicloud_security_group" "default" {
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_alikafka_instance" "default" {
  name            = var.name
  partition_num   = 50
  disk_type       = "1"
  disk_size       = "500"
  deploy_type     = "5"
  io_max          = "20"
  spec_type       = "professional"
  service_version = "2.2.0"
  vswitch_id      = alicloud_vswitch.default.id
  security_group  = alicloud_security_group.default.id
  config = jsonencode(
    {
      "enable.acl" : "true"
    }
  )
}

resource "alicloud_alikafka_topic" "default" {
  instance_id   = alicloud_alikafka_instance.default.id
  topic         = var.name
  remark        = var.name
  local_topic   = "true"
  compact_topic = "true"
  partition_num = "18"
  configs = jsonencode(
    {
      "message.format.version" : "2.2.0",
      "max.message.bytes" : "10485760",
      "min.insync.replicas" : "1",
      "replication-factor" : "2",
      "retention.ms" : "3600000"
    }
  )
  tags = {
    Created = "TF",
    For     = "example",
  }
}
```

## Argument Reference

The following arguments are supported:
* `compact_topic` - (Optional, ForceNew) When the storage engine of the Topic is configured as Local storage, a log cleanup policy is configured. Value:
  - false:delete the cleanup policy.
  - true:compact cleanup policy.
* `config` - (Optional, Available since v1.267.0) The key of the Topic configuration word.
  - Currently, Reserved Instances only support the Topic of the Local storage engine, and the Topic of the cloud storage engine cannot be modified.
  - Supports Serverless instances to modify Topic configurations.
  - The 'local topic' supported keys for Reserved Instances are retention.ms, max.message.bytes, replications, message.timestamp.type, message.timestamp.difference.max.ms.
  - A Serverless instance supports retention.hours, max.message.bytes, message.timestamp.type, and message.timestamp.difference.max.ms.
* `configs` - (Optional, Computed, JsonString, Available since v1.262.1) Supplementary configuration.
Must be in JSON format.
Currently, only the Key is replications. Indicates the number of Topic replicas. The value type is Integer, and the value limit is 1 to 3.
This parameter can be specified only if the value of LocalTopic is true or the specification type is open source version (local disk).
Reserved Instances support the following configurations:
The retention.ms (message retention duration) range is 3600000 to 31536000000 milliseconds.
max.message.bytes (indicating the maximum size of the message sent) is in the range of 1048576 to 10485760 bytes.
message.timestamp.type Specifies the type of message timestamp. CreateTime indicates the specified timestamp when the producer sends the message. If it is not specified, it is the creation time of the message on the client. LogAppendTime indicates the drop time of the message on the server. Optional values are: CreateTime or LogAppendTime.
Serverless instances support the following configurations:
retention.hours (message retention duration). The value type is String, and the value limit is 24 to 8760.
max.message.bytes (the maximum size of the message to be sent). The value type is String, and the value is limited to 1048576 to 10485760.
message.timestamp.type (the type of the message timestamp). CreateTime indicates the specified timestamp when the producer sends the message. If it is not specified, it is the creation time of the message on the client. LogAppendTime indicates the drop time of the message on the server. Optional values are: CreateTime or LogAppendTime.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `instance_id` - (Required, ForceNew) The instance ID.
* `local_topic` - (Optional, ForceNew) The storage engine of the Topic. Value:
  - false: Cloud storage.
  - true:Local storage.
* `min_insync_replicas` - (Optional, Int, Available since v1.267.0) Minimum number of ISR synchronization replicas.
  - This parameter can be specified only when the value of `LocalTopic` is `true` or The `type` is **open source version (local disk).
  - The value must be less than the number of Topic copies.
  - The number of synchronized copies is limited to 1~3.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `partition_num` - (Optional, Computed, Int) Number of partitions
* `remark` - (Required) Remarks
* `replication_factor` - (Optional, Int, Available since v1.267.0) The number of Topic replicas.
  - This parameter can be specified only when the value of `LocalTopic` is `true` or The `type` is **open source version (local disk).
  - The number of copies is limited to 1~3.

-> **NOTE:**  When the number of copies is `1`, there is a risk of data loss. Please set it carefully.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `tags` - (Optional, Map, Available since v1.63.0) The tag of the kafka console, which is used to group instance,topic, and consumption.
* `topic` - (Required, ForceNew) The Topic name.
* `value` - (Optional, Available since v1.267.0) The value of the Topic configuration word.
  - Serverless instances support the following configurations:
  -'retention.hours' indicates the message retention duration. The value type is String, and the value limit is 24 to 8760.
  -'max.message.bytes' indicates the maximum size of the sent message. The value type is String and the value limit is 1048576~10485760.
  -'message.timestamp.type' Specifies the type of message timestamp. CreateTime indicates the specified timestamp when the producer sends the message. If it is not specified, it is the creation time of the message on the client. LogAppendTime indicates the drop time of the message on the server. Optional values are: CreateTime or LogAppendTime.
  -'message.timestamp.difference.max.ms' indicates the maximum allowed difference between the timestamp when the server receives the message and the timestamp specified in the message. When message.timestamp.type is set to CreateTime, **messages are rejected** if the timestamp difference exceeds this threshold * *. This configuration does not take effect when message.timestamp.type is LogAppendTime.
  - Reserved Instances support the following configurations:
  -'retention.ms' indicates the message retention duration. The value type is String and the value limit is 3600000~31536000000.
  -'max.message.bytes' indicates the maximum size of the sent message. The value type is String and the value limit is 1048576~10485760.
  -'replications' indicates the number of replicas. The value type is String, and the value is limited to 1 to 3.
  -'message.timestamp.type' Specifies the type of message timestamp. CreateTime indicates the specified timestamp when the producer sends the message. If it is not specified, it is the creation time of the message on the client. LogAppendTime indicates the drop time of the message on the server. Optional values are: CreateTime or LogAppendTime.
  -'message.timestamp.difference.max.ms' indicates the maximum allowed difference between the timestamp when the server receives the message and the timestamp specified in the message. When message.timestamp.type is set to CreateTime, **messages are rejected** if the timestamp difference exceeds this threshold * *. This configuration does not take effect when message.timestamp.type is LogAppendTime.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<instance_id>:<topic>`.
* `create_time` - The creation timestamp. Unit: millisecond
* `region_id` - The region ID of the Topic instance.
* `status` - Service status. Value:

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Topic.
* `delete` - (Defaults to 16 mins) Used when delete the Topic.
* `update` - (Defaults to 5 mins) Used when update the Topic.

## Import

Ali Kafka Topic can be imported using the id, e.g.

```shell
$ terraform import alicloud_alikafka_topic.example <instance_id>:<topic>
```