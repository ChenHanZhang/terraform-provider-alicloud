---
subcategory: "Ali Kafka"
layout: "alicloud"
page_title: "Alicloud: alicloud_ali_kafka_acl"
description: |-
  Provides a Alicloud Ali Kafka Acl resource.
---

# alicloud_ali_kafka_acl

Provides a Ali Kafka Acl resource.

Kafka access control.

For information about Ali Kafka Acl and how to use it, see [What is Acl](https://next.api.alibabacloud.com/document/alikafka/2019-09-16/CreateAcl).

-> **NOTE:** Available since v1.266.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-beijing"
}

variable "region" {
  default = "cn-beijing"
}

resource "alicloud_vpc" "default0ppNRd" {
  ipv4_cidr_mask = "24"
  is_default     = false
  cidr_block     = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultn6iyAD" {
  vpc_id     = alicloud_vpc.default0ppNRd.id
  cidr_block = "172.18.0.0/23"
  zone_id    = "cn-beijing-a"
}

resource "alicloud_alikafka_instance_v2" "defaultZgBJfR" {
  deploy_type = "5"
  spec_type   = "normal"
  config      = "{\"auto.create.topics.enable\":\"true\",\"enable.acl\":\"true\",\"enable.vpc_sasl_ssl\":\"false\",\"kafka.ssl.bit\":\"4096\",\"log.retention.hours\":\"72\",\"message.max.bytes\":\"1048576\",\"num.partitions\":\"3\",\"offsets.retention.minutes\":\"10080\"}"
  zone_id     = "cn-beijing-a"
  vswitch_id  = alicloud_vswitch.defaultn6iyAD.id
  paid_type   = "3"
  serverless_config {
    reserved_publish_capacity   = "60"
    reserved_subscribe_capacity = "60"
  }
  vpc_id = alicloud_vpc.default0ppNRd.id
}

resource "alicloud_alikafka_sasl_user" "defaulty2f9ds" {
  type        = "scram"
  username    = "qwoeiuqwoieurandom203"
  password    = "123123"
  instance_id = alicloud_alikafka_instance_v2.defaultZgBJfR.id
}


resource "alicloud_ali_kafka_acl" "default" {
  username                  = alicloud_alikafka_sasl_user.defaulty2f9ds.username
  acl_operation_types       = "IDEMPOTENT_WRITE"
  acl_operation_type        = "IDEMPOTENT_WRITE"
  acl_permission_type       = "ALLOW"
  host                      = "*"
  acl_resource_pattern_type = "LITERAL"
  acl_resource_name         = "*"
  instance_id               = alicloud_alikafka_instance_v2.defaultZgBJfR.id
  acl_resource_type         = "CLUSTER"
}
```

## Argument Reference

The following arguments are supported:
* `acl_operation_type` - (Required, ForceNew) The operation type.
Package/Pay by Hour Instance Value:
  - `Write`: Write
  - `Read`: Read
  - `Describe`:: Read TransactionalId
  - `IdempotentWrite`:: idempotent Cluster

Serverless instance value:
  - `WRITE`: WRITE
  - `READ`: READ
  - `DESCRIBE`: read TransactionalId
  - `IDEMPOTENT_WRITE`: idempotent Cluster
  - `DESCRIBE_CONFIGS`: queries the configuration
* `acl_operation_types` - (Optional) The type of the volume authorization operation. Multiple operations to, split.

The operation type. Return value:
  - `WRITE`: WRITE
  - `READ`: READ
  - `DESCRIBE`: read TransactionalId
  - `IDEMPOTENT_WRITE`: idempotent Cluster
  - `DESCRIBE_CONFIGS`: queries the configuration

-> **NOTE:**  This parameter only supports Serverless instances.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `acl_permission_type` - (Optional, ForceNew) Authorization method. Value:
  - DENY: refuse
  - ALLOW: ALLOW

-> **NOTE:**  This field only supports Serverless instances.

* `acl_resource_name` - (Required, ForceNew) The resource name.
  - Name of the Topic or Consumer Group.
  - Supports the use of an asterisk (*) to indicate the names of all topics or Conusmer groups.
* `acl_resource_pattern_type` - (Required, ForceNew) Match the pattern. Value:
  - `LITERAL`: Full match
  - `PREFIXED`: prefix matching
* `acl_resource_type` - (Required, ForceNew) The resource type. Value:
  - `TOPIC`: the message TOPIC.
  - `GROUP`: consumer GROUP.
  - `CLUSTER`: the instance.
  - `TRANSACTIONAL_ID`: The transaction ID.
* `host` - (Optional, ForceNew) Host.
* `instance_id` - (Required, ForceNew) The instance ID.
* `username` - (Required, ForceNew) The user name.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<instance_id>:<acl_resource_type>:<acl_resource_pattern_type>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Acl.
* `delete` - (Defaults to 5 mins) Used when delete the Acl.

## Import

Ali Kafka Acl can be imported using the id, e.g.

```shell
$ terraform import alicloud_ali_kafka_acl.example <instance_id>:<acl_resource_type>:<acl_resource_pattern_type>
```