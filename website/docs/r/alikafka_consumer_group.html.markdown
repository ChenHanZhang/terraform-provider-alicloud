---
subcategory: "Ali Kafka"
layout: "alicloud"
page_title: "Alicloud: alicloud_alikafka_consumer_group"
description: |-
  Provides a Alicloud Ali Kafka Consumer Group resource.
---

# alicloud_alikafka_consumer_group

Provides a Ali Kafka Consumer Group resource.

Group in kafka.

For information about Ali Kafka Consumer Group and how to use it, see [What is Consumer Group](https://next.api.alibabacloud.com/document/alikafka/2019-09-16/CreateConsumerGroup).

-> **NOTE:** Available since v1.56.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "default" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default" {
  vpc_id     = alicloud_vpc.default.id
  cidr_block = "172.16.0.0/24"
  zone_id    = data.alicloud_zones.default.zones[0].id
}

resource "alicloud_security_group" "default" {
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_alikafka_instance" "default" {
  name           = "${var.name}-${random_integer.default.result}"
  partition_num  = "50"
  disk_type      = "1"
  disk_size      = "500"
  deploy_type    = "5"
  io_max         = "20"
  vswitch_id     = alicloud_vswitch.default.id
  security_group = alicloud_security_group.default.id
}

resource "alicloud_alikafka_consumer_group" "default" {
  consumer_id = var.name
  instance_id = alicloud_alikafka_instance.default.id
}
```

## Argument Reference

The following arguments are supported:
* `consumer_id` - (Required, ForceNew) The Group name.
  - Reserved Instances:
It supports uppercase and lowercase letters, digits, underscores (_), dashes (-), and periods (.), and is limited to 3 to 64 characters in length.
  - Serverless instance:
Can only contain letters, numbers and special characters "@._* $#^! &-", limited to 1~249 characters.
* `instance_id` - (Required, ForceNew) The ID of the instance.
* `offset` - (Optional, Int, Available since v1.266.0) Partition site.
* `partition` - (Optional, Int, Available since v1.266.0) The partition ID.
* `remark` - (Optional, ForceNew, Available since v1.266.0) Remarks.
* `reset_type` - (Optional, Available since v1.266.0) The type of the consumer group's consumption point. The following two types are supported:
  - `timestamp` (default)
  - `offset`

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `tags` - (Optional, Map) A list of tags.
* `time` - (Optional, Available since v1.266.0) The input time parameter, in Unix timestamp format, in milliseconds.
The parameter range must be **less than 0** or **within the retention time range of the consumption site**. This parameter takes effect only when the resetType is set to timestamp.
  - To reset to the latest consumption point, pass in - 1.
  - To reset to the earliest consumption point, pass in - 2.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `topic` - (Optional, ForceNew, Available since v1.266.0) The Topic to which the Consumer subscribes

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<instance_id>:<consumer_id>`.
* `create_time` - The creation timestamp. Unit: millisecond
* `region_id` - The region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Consumer Group.
* `delete` - (Defaults to 5 mins) Used when delete the Consumer Group.
* `update` - (Defaults to 5 mins) Used when update the Consumer Group.

## Import

Ali Kafka Consumer Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_alikafka_consumer_group.example <instance_id>:<consumer_id>
```