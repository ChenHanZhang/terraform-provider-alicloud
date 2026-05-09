---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_agg_task_group"
description: |-
  Provides a Alicloud Cms Agg Task Group resource.
---

# alicloud_cms_agg_task_group

Provides a Cms Agg Task Group resource.

Aggregation Task Group.

For information about Cms Agg Task Group and how to use it, see [What is Agg Task Group](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateAggTaskGroup).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_cms_agg_task_group" "default" {
  target_prometheus_id       = "rw-ac05cec04ad891d38d3fe1c5c7c9"
  source_prometheus_id       = "rw-ac05cec04ad891d38d3fe1c5c7c9"
  agg_task_group_name        = "example-group-model-example-278"
  agg_task_group_config      = <<EOF
---
groups:
- name: example-group
  interval: 1m
  rules:
  - record: cpu_user_util:1m
    expr: cpu_user_util

EOF
  override_if_exists         = true
  max_retries                = "20"
  status                     = "Running"
  schedule_time_expr         = "@m"
  description                = "描述example"
  agg_task_group_config_type = "RecordingRuleYaml"
  schedule_mode              = "FixedRate"
  precheck_string            = "{\"policy\":\"skip\",\"prometheusId\":\"rw-ac05cec04ad891d38d3fe1c5c7c9\",\"query\":\"scalar(sum(count_over_time(up{job=\\\"_arms/kubelet/cadvisor\\\"}[15s])) / 21)\",\"threshold\":1.0,\"timeout\":13,\"type\":\"none\"}"
  from_time                  = "1727409939"
  to_time                    = "0"
  max_run_time_in_seconds    = "200"
  delay                      = "31"
}
```

## Argument Reference

The following arguments are supported:
* `agg_task_group_config` - (Required) Aggregation Task Group configuration
* `agg_task_group_config_type` - (Optional) Aggregation Task Group configuration type

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `agg_task_group_name` - (Required) The name of the resource
* `cron_expr` - (Optional) The Cron expression executed when the scheduling mode is Cron
* `delay` - (Optional, Int) Aggregation Task Group Delay
* `description` - (Optional) Aggregation Task Group Description
* `from_time` - (Optional, Int) Schedule Start Timestamp
* `max_retries` - (Optional, Int) Aggregation task group maximum retries
* `max_run_time_in_seconds` - (Optional, Int) Aggregation task group maximum run time
* `override_if_exists` - (Optional) When creating an aggtask group, whether the resource with the same name is overwritten and updated

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `precheck_string` - (Optional) Precheck configuration string
* `schedule_mode` - (Optional) Scheduling Mode for Aggregation Task Groups
* `schedule_time_expr` - (Optional) Schedule expressions for rules that aggregate task groups
* `source_prometheus_id` - (Required, ForceNew) The ID of the Prometheus instance where the aggregation task group read from.
* `status` - (Optional, Computed) The status of the resource
* `tags` - (Optional, Map) The tag of the resource
* `target_prometheus_id` - (Required, ForceNew) The ID of the Prometheus instance where the aggregation task group writtes to
* `to_time` - (Optional, Int) Schedule End Timestamp

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<source_prometheus_id>:<agg_task_group_id>`.
* `agg_task_group_id` - The Id of the aggregation task group.
* `region_id` - The region ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Agg Task Group.
* `delete` - (Defaults to 5 mins) Used when delete the Agg Task Group.
* `update` - (Defaults to 5 mins) Used when update the Agg Task Group.

## Import

Cms Agg Task Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_agg_task_group.example <source_prometheus_id>:<agg_task_group_id>
```