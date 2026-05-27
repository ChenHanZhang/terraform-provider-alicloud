---
subcategory: "Cloud Monitor Service"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_monitor_service_metric_alarm_rule"
description: |-
  Provides a Alicloud Cloud Monitor Service Metric Alarm Rule resource.
---

# alicloud_cloud_monitor_service_metric_alarm_rule

Provides a Cloud Monitor Service Metric Alarm Rule resource.

Describes the time series indicator alarm rules set by the user.

For information about Cloud Monitor Service Metric Alarm Rule and how to use it, see [What is Metric Alarm Rule](https://next.api.alibabacloud.com/document/Cms/2019-01-01/PutResourceMetricRule).

-> **NOTE:** Available since v1.280.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_cloud_monitor_service_metric_alarm_rule" "default" {
  contact_groups        = "Sun"
  metric_name           = "cpu_total"
  metric_alarm_rule_id  = "henghai-example1666077368000"
  period                = "15"
  rule_name             = "henghai-example"
  namespace             = "acs_ecs_dashboard"
  no_effective_interval = "00:00-23:59"
  escalations {
    info {
      times               = "20"
      threshold           = "10000"
      statistics          = "Average"
      comparison_operator = "GreaterThanOrEqualToThreshold"
    }
    warn {
    }
    critical {
    }
  }
}
```

## Argument Reference

The following arguments are supported:
* `composite_expression` - (Optional, ForceNew, Set) Multi-indicator alarm conditions.

-> **NOTE:**  single indicator and multiple indicators are mutually exclusive and cannot be set at the same time.
 See [`composite_expression`](#composite_expression) below.
* `contact_groups` - (Required, ForceNew) Alarm contact group. Alarm notifications are sent to the alarm contacts in the alarm contact group.

-> **NOTE:**  An alarm Contact Group is a group of alarm contacts that can contain one or more alarm contacts. For more information about how to create alarm contacts and alarm contact groups, see [PutContact](~~ 114923 ~~) and [PutContactGroup](~~ 114929 ~~).

* `effective_interval` - (Optional) The effective time period of the alarm rule.
* `email_subject` - (Optional) Alarm email subject.
* `escalations` - (Optional, ForceNew, Set) Alarm Level trigger condition. See [`escalations`](#escalations) below.
* `interval` - (Optional) The trigger period of the alarm rule. Unit: seconds.

-> **NOTE:**  For more information about how to query the statistical cycle of metrics, see [Cloud product metrics](~~ 163515 ~~).


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `labels` - (Optional, List) When the monitoring item reaches the alarm condition and the alarm is performed, the tag is written into the monitoring item at the same time and displayed in the alarm notification.

-> **NOTE:**  This function is the same as the Label in the Prometheus alarm.
 See [`labels`](#labels) below.
* `metric_alarm_rule_id` - (Required, ForceNew) The unique identifier of the alarm rule.
* `metric_name` - (Required, ForceNew) Monitoring item name.
* `namespace` - (Required, ForceNew) The data namespace of the product is used to distinguish different products.
* `no_data_policy` - (Optional, Computed) The processing method of alarm when there is no monitoring data. Value:
  - KEEP\_last \_State (default): no processing.
  - INSUFFICIENT\_data: The alarm content is no data.
  - OK: normal.
* `no_effective_interval` - (Optional) The time period when the alarm rule does not take effect, for example, 00:00-23:59.
* `period` - (Optional) Statistical cycle.
* `prometheus` - (Optional, Set) Prometheus alarm.

-> **NOTE:**  only if you create a Prometheus alarm rule for enterprise cloud monitoring, you need to set this parameter.
 See [`prometheus`](#prometheus) below.
* `resources` - (Required) Resource information, such as: '[{"instanceId":"i-uf6j91r34rnwawoo ****"}]', '[{"userId":"100931896542 ****"}]'.
For more information about Dimensions supported by resource information, see [Cloud product metrics](~~ 163515 ~~).
* `rule_name` - (Required, ForceNew) The name of the alarm rule.
* `silence_time` - (Optional) Channel silence period, in seconds. The default value is 86400 seconds (1 day) and the shortest is 3600 seconds (1 hour). When the monitoring data continuously exceeds the alarm rule threshold, only one alarm notification is sent in each silence period.
* `status` - (Optional, Computed) Enabled status: true means on, fase means disabled.
* `targets` - (Optional, List) The alarm triggers the target. See [`targets`](#targets) below.
* `webhook` - (Optional) The URL address specified when an alarm callback occurs. A POST request is sent to the URL.

### `composite_expression`

The composite_expression supports the following:
* `expression_list` - (Optional, List) List of alarm conditions created by the standard. See [`expression_list`](#composite_expression-expression_list) below.
* `expression_list_join` - (Optional) The relationship between multi-indicator alarm conditions. Value:
  -'& &': When all indicators meet the alarm conditions, the alarm is triggered. An alarm is triggered only when each expression under the ExpressionList is' true.
  -'| |': One of the indicators meets the alarm condition, that is, an alarm is triggered.
* `expression_raw` - (Optional) The alarm condition created by the expression. Including but not limited to the following:
  - Set an alarm blacklist for some resources, for example: '$instanceId! = 'i-io8kfvcpp7x5 ****' ''& &'' $Average > 50 ', which means that when the'average' of the instanti' i-io8kfvcpp7x5 ****' in the alarm rule is greater than 50, an alarm will not be generated.
  - Set a special alarm threshold for the specified instance in the rule, for example: '$Average > ($instanceId = =' i-io8kfvcpp7x5 * * * * '? 80: 50)', which means that the alarm will be given only when the average' of the instance' i-io8kfvcpp7x5 * * * * 'in the alarm rule is greater than 80, and the alarm will be given when the average' of other instances' Average' is greater than 50.
  - Limit the number of instances that exceed the threshold in the rule, for example: 'count($Average > 20) > 3', which means that the alarm is only performed when the number of instances of' Average 'in the alarm rule is greater than 20.
* `level` - (Optional) Alarm level. Value:
  - Critical: serious.
  - Warn: WARNING.
  - Info: information.
* `times` - (Optional, Int) The number of times the alarm condition is required to be reached when an alarm notification is issued.

### `composite_expression-expression_list`

The composite_expression-expression_list supports the following:
* `comparison_operator` - (Optional) Threshold comparator. Value:
  - GreaterThanOrEqualToThreshold: greater than or equal.
  - Greatthanthreshold: greater.
  - LessThanOrEqualToThreshold: less than or equal.
  - Lesthanthreshold: less.
  - NotEqualToThreshold: not equal.
  - GreaterThanYesterday: up from yesterday's time.
  - Lesthanyesterday: down from yesterday's time.
  - GreaterThanLastWeek: up from the same time last week.
  - Lesthanlastweek: down from the same time last week.
  - GreaterThanLastPeriod: Month-on-month increase in the previous cycle.
  - Lesthanlastperiod: Ring-on-ring decline.
* `metric_name` - (Optional) The name of the monitoring item of the cloud product.
* `period` - (Optional, Int) Aggregation cycle of monitoring items.
Unit: seconds.
* `statistics` - (Optional) Statistical methods for monitoring items. Value:
  -$Maximum: Maximum.
  -$Minimum: Minimum.
  -$Average: Average.
  -$Availability: Availability (usually used for site monitoring).

-> **NOTE:**  '$' is the uniform prefix symbol for monitoring items. For more information about supported cloud products, see [Cloud product monitoring items](~~ 163515 ~~).

* `threshold` - (Optional) Alarm threshold.

### `escalations`

The escalations supports the following:
* `critical` - (Optional, ForceNew, Set) Critical level alarm trigger condition. See [`critical`](#escalations-critical) below.
* `info` - (Optional, ForceNew, Set) The trigger condition of the Info level alarm. See [`info`](#escalations-info) below.
* `warn` - (Optional, ForceNew, Set) The trigger condition of warning level. See [`warn`](#escalations-warn) below.

### `escalations-critical`

The escalations-critical supports the following:
* `comparison_operator` - (Optional) Critical level threshold comparator. Value:
  - GreaterThanOrEqualToThreshold: greater than or equal.
  - Greatthanthreshold: greater.
  - LessThanOrEqualToThreshold: less than or equal.
  - Lesthanthreshold: less.
  - NotEqualToThreshold: not equal.
  - GreaterThanYesterday: up from yesterday's time.
  - Lesthanyesterday: down from yesterday's time.
  - GreaterThanLastWeek: up from the same time last week.
  - Lesthanlastweek: down from the same time last week.
  - GreaterThanLastPeriod: Month-on-month increase in the previous cycle.
  - Lesthanlastperiod: Ring-on-ring decline.
* `statistics` - (Optional) Critical Level Alarm statistics method.
* `threshold` - (Optional) Critical level alarm threshold.
* `times` - (Optional, Int) The number of consecutive occurrences of Critical level alarms. The alarm occurs continuously for this number of times, and the alarm will only be triggered if the threshold is exceeded.

### `escalations-info`

The escalations-info supports the following:
* `comparison_operator` - (Optional) Info level threshold comparator. Value:
  - GreaterThanOrEqualToThreshold: greater than or equal.
  - Greatthanthreshold: greater.
  - LessThanOrEqualToThreshold: less than or equal.
  - Lesthanthreshold: less.
  - NotEqualToThreshold: not equal.
  - GreaterThanYesterday: up from yesterday's time.
  - Lesthanyesterday: down from yesterday's time.
  - GreaterThanLastWeek: up from the same time last week.
  - Lesthanlastweek: down from the same time last week.
  - GreaterThanLastPeriod: Month-on-month increase in the previous cycle.
  - Lesthanlastperiod: Ring-on-ring decline.
* `statistics` - (Optional) Info level alarm Statistics method.
* `threshold` - (Optional) The alarm threshold of the Info level.
* `times` - (Optional, Int) The number of consecutive occurrences of the Info level alarm. The alarm occurs continuously for this number of times, and the alarm will only be triggered if the threshold is exceeded.

### `escalations-warn`

The escalations-warn supports the following:
* `comparison_operator` - (Optional) Warn-level threshold comparator. Value:
  - GreaterThanOrEqualToThreshold: greater than or equal.
  - Greatthanthreshold: greater.
  - LessThanOrEqualToThreshold: less than or equal.
  - Lesthanthreshold: less.
  - NotEqualToThreshold: not equal.
  - GreaterThanYesterday: up from yesterday's time.
  - Lesthanyesterday: down from yesterday's time.
  - GreaterThanLastWeek: up from the same time last week.
  - Lesthanlastweek: down from the same time last week.
  - GreaterThanLastPeriod: Month-on-month increase in the previous cycle.
  - Lesthanlastperiod: Ring-on-ring decline.
* `statistics` - (Optional) Warning statistical method.
* `threshold` - (Optional) Warning threshold.
* `times` - (Optional, Int) The number of consecutive occurrences of warning level alarms. The alarm occurs continuously for this number of times, and the alarm will only be triggered if the threshold is exceeded.

### `labels`

The labels supports the following:
* `key` - (Optional) Label key.
* `value` - (Optional) Label value.

-> **NOTE:**  The label value supports template parameters and replaces the template parameters with the actual label value.


### `prometheus`

The prometheus supports the following:
* `annotations` - (Optional, List) When Prometheus alarms, The Annotated keys and values are rendered to facilitate your understanding of monitoring items or alarm rules.

-> **NOTE:**  This function is equivalent to Prometheus's Annotation.
 See [`annotations`](#prometheus-annotations) below.
* `level` - (Optional) Alarm level. Value:
  - Critical: serious.
  - Warn: WARNING.
  - Info: information.
* `prom_ql` - (Optional) The PromQL query statement.

-> **NOTE:**  The data obtained through the PromQL query statement is the alarm data. Please carry the alarm threshold in this statement.

* `times` - (Optional, Int) The number of times the alarm condition is required to be reached when an alarm notification is issued.

### `prometheus-annotations`

The prometheus-annotations supports the following:
* `key` - (Optional) The key of the comment.
* `value` - (Optional) The value of the comment.

### `targets`

The targets supports the following:
* `arn` - (Required) Resource ARN. The format is 'acs:{cloud product abbreviation }:{ regionId }:{ userId}:/{resource type}/{resource name}/message '. For example: 'acs:mns:cn-hangzhou:120886317861 ****:/queues/test123/message '. The parameters are described as follows:
  -{Cloud product abbreviation}: Currently, only MNS is supported.
  -{userId}: the ID of the Alibaba cloud account.
  -{regionId}: the region where the message queue or topic is located.
  -{Resource type} ': the resource type that receives the alarm. Value:
  - `queues`: queue.
  - `topics`: topic.
  -{Resource name}: resource name.
  - If the resource type is `queues`, the resource name is the queue name.
  - If the resource type is `topics`, the resource name is the subject name.


* `json_params` - (Optional) The JSON format parameter of the alarm callback.
* `level` - (Optional) Alarm level. Value:
  - INFO: information.
  - WARN: WARNING.
  - CRITICAL: urgent.
* `target_id` - (Required) The ID of the alarm trigger target.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `dimensions` - The extended resource associated with the Monitor.
* `escalations` - Alarm Level trigger condition.
  * `critical` - Critical level alarm trigger condition.
    * `pre_condition` - The precondition of Critical level alarm.
  * `info` - The trigger condition of the Info level alarm.
    * `pre_condition` - Preconditions for Info level alarms.
  * `warn` - The trigger condition of warning level.
    * `pre_condition` - The precondition of warning level alarm.
* `source_type` - The alarm rule type.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Metric Alarm Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Metric Alarm Rule.
* `update` - (Defaults to 5 mins) Used when update the Metric Alarm Rule.

## Import

Cloud Monitor Service Metric Alarm Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_monitor_service_metric_alarm_rule.example <metric_alarm_rule_id>
```