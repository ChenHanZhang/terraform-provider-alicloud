---
subcategory: "AliKafka"
layout: "alicloud"
page_title: "Alicloud: alicloud_ali_kafka_allowed_ip_attachment"
description: |-
  Provides a Alicloud Ali Kafka Allowed Ip Attachment resource.
---

# alicloud_ali_kafka_allowed_ip_attachment

Provides a Ali Kafka Allowed Ip Attachment resource.



For information about Ali Kafka Allowed Ip Attachment and how to use it, see [What is Allowed Ip Attachment](https://next.api.alibabacloud.com/document/alikafka/2019-09-16/UpdateAllowedIp).

-> **NOTE:** Available since v1.267.0.

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

variable "inner_ip" {
  default = "192.168.1.810"
}

variable "desc" {
  default = "example-531"
}

resource "alicloud_vpc" "defaultbqyuKu" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultjvEBtT" {
  vpc_id     = alicloud_vpc.defaultbqyuKu.id
  zone_id    = "cn-beijing-a"
  cidr_block = "172.18.0.0/23"
}

resource "alicloud_alikafka_instance_v2" "defaultBzMsgy" {
  deploy_type = "5"
  spec_type   = "normal"
  config      = "{\"auto.create.topics.enable\":\"true\",\"enable.acl\":\"true\",\"enable.vpc_sasl_ssl\":\"false\",\"kafka.ssl.bit\":\"4096\",\"log.retention.hours\":\"72\",\"message.max.bytes\":\"1048576\",\"num.partitions\":\"3\",\"offsets.retention.minutes\":\"10080\"}"
  zone_id     = "cn-beijing-a"
  vswitch_id  = alicloud_vswitch.defaultjvEBtT.id
  vpc_id      = alicloud_vpc.defaultbqyuKu.id
  paid_type   = "3"
  serverless_config {
    reserved_publish_capacity   = "60"
    reserved_subscribe_capacity = "60"
  }
}


resource "alicloud_ali_kafka_allowed_ip_attachment" "default" {
  description       = var.desc
  instance_id       = alicloud_alikafka_instance_v2.defaultBzMsgy.id
  port_range        = "9092/9092"
  allowed_list_ip   = var.inner_ip
  allowed_list_type = "vpc"
}
```

## Argument Reference

The following arguments are supported:
* `allowed_list_ip` - (Required, ForceNew) The IP list. It can be a network segment, for example, **192.168.0.0/16**.
  - When `UpdateType` is `add`, you can enter multiple entries separated by commas (,).
  - When `UpdateType` is set to `delete`, only one item can be entered at a time.
  - Delete with caution.
* `allowed_list_type` - (Required, ForceNew) The type of whitelist. Value:
  - `vpc`: VPC.
  - `internet`: public network.
* `description` - (Optional) Whitelist description.

-> **NOTE:** This parameter only applies during resource creation, deletion. If modified in isolation without other property changes, Terraform will not trigger any action.

* `instance_id` - (Required, ForceNew) The first ID of the resource
* `port_range` - (Required, ForceNew) Port range. Value:
  - **9092/9092**: Proprietary network VPC-PLAINTEXT protocol.
  - **9093/9093**: public network-SASL_SSL protocol.
  - **9094/9094**: Proprietary network VPC-SASL_PLAINTEXT protocol.
  - **9095/9095**: Proprietary network VPC-SASL_SSL protocol.

This parameter must correspond to `AllowdedListType`.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<instance_id>:<port_range>:<allowed_list_type>:<allowed_list_ip>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Allowed Ip Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Allowed Ip Attachment.

## Import

Ali Kafka Allowed Ip Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_ali_kafka_allowed_ip_attachment.example <instance_id>:<port_range>:<allowed_list_type>:<allowed_list_ip>
```