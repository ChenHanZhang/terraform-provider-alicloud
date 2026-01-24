---
subcategory: "AliKafka"
layout: "alicloud"
page_title: "Alicloud: alicloud_ali_kafka_scheduled_scaling_attachment"
description: |-
  Provides a Alicloud Ali Kafka Scheduled Scaling Attachment resource.
---

# alicloud_ali_kafka_scheduled_scaling_attachment

Provides a Ali Kafka Scheduled Scaling Attachment resource.

Elastic timing strategy.

For information about Ali Kafka Scheduled Scaling Attachment and how to use it, see [What is Scheduled Scaling Attachment](https://next.api.alibabacloud.com/document/alikafka/2019-09-16/CreateScheduledScalingRule).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-beijing"
}

variable "weekly_types" {
  default = "Monday"
}

variable "rulename" {
  default = "example"
}

variable "region" {
  default = "cn-beijing"
}

resource "alicloud_vpc" "defaultqSIhPN" {
  cidr_block = "195.0.0.0/24"
}

resource "alicloud_vswitch" "defaulttqV6cH" {
  vpc_id     = alicloud_vpc.defaultqSIhPN.id
  cidr_block = "195.0.0.0/25"
  zone_id    = "cn-beijing-a"
}

resource "alicloud_alikafka_instance_v2" "default1wRjcq" {
  deploy_type   = "5"
  spec_type     = "normal"
  deploy_module = "vpc"
  vswitch_id    = alicloud_vswitch.defaulttqV6cH.id
  vpc_id        = alicloud_vpc.defaultqSIhPN.id
  paid_type     = "3"
  serverless_config {
    reserved_publish_capacity   = "60"
    reserved_subscribe_capacity = "60"
  }
}


resource "alicloud_ali_kafka_scheduled_scaling_attachment" "default" {
  schedule_type        = "repeat"
  reserved_sub_flow    = "200"
  reserved_pub_flow    = "200"
  time_zone            = "GMT+8"
  duration_minutes     = "100"
  first_scheduled_time = "1769256652"
  enable               = true
  repeat_type          = "Weekly"
  weekly_types         = ["${var.weekly_types}"]
  rule_name            = var.rulename
  instance_id          = alicloud_alikafka_instance_v2.default1wRjcq.id
}
```

## Argument Reference

The following arguments are supported:
* `duration_minutes` - (Required, ForceNew, Int) Duration (in minutes), which is the duration of a scheduled elastic task.

-> **NOTE:**  No less than 15 minutes.

* `enable` - (Optional) Enable or disable the scheduled task policy. The value range is as follows:
  - `true`: Enabled.
  - `false`: Disabled.
* `first_scheduled_time` - (Required, ForceNew, Int) Start timing policy execution time.

If the policy type is a single-time scheduling, ensure that the execution start time is more than 30 minutes later than the current time.



-> **NOTE:**  In order to prevent the server from continuously performing up-and-down tasks, the time interval between different timing tasks should be at least greater than 60 minutes.> 

* `instance_id` - (Required, ForceNew) The instance ID.
* `repeat_type` - (Optional, ForceNew) When ScheduleType is repeat, the parameter is required. The enumerated value is
 Daily: scheduled every day.
 Weekly: scheduled Weekly.
* `reserved_pub_flow` - (Required, ForceNew, Int) Timing elastic reserve production specifications (unit: MB/s).

-> **NOTE:**  At least one of the ReservedPubFlow and ReservedSubFlow input parameters must be higher than the current specification.

* `reserved_sub_flow` - (Required, ForceNew, Int) Timing elastic reservation consumption specification (unit: MB/s).

-> **NOTE:**  At least one of the ReservedSubFlow and ReservedPubFlow input parameters must be higher than the current specification.

* `rule_name` - (Required, ForceNew) The scheduled policy rule name.

-> **NOTE:**  cannot be the same as other rule names in the same instance.

* `schedule_type` - (Required, ForceNew) Timing type. The values are as follows:
  - at: Dispatch only once.
  - repeat: repeat the schedule.
* `time_zone` - (Required, ForceNew) Time zone (Coordinated Universal Time).
* `weekly_types` - (Optional, ForceNew, List) Weekly type. Support multi-day repeat execution.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Scheduled Scaling Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Scheduled Scaling Attachment.
* `update` - (Defaults to 5 mins) Used when update the Scheduled Scaling Attachment.

## Import

Ali Kafka Scheduled Scaling Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_ali_kafka_scheduled_scaling_attachment.example <instance_id>
```