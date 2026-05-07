---
subcategory: "S T A R Ops"
layout: "alicloud"
page_title: "Alicloud: alicloud_s_t_a_r_ops_mission"
description: |-
  Provides a Alicloud S T A R Ops Mission resource.
---

# alicloud_s_t_a_r_ops_mission

Provides a S T A R Ops Mission resource.

A Mission is an asynchronous session that includes multiple Blueprint scheduling tasks.  .

For information about S T A R Ops Mission and how to use it, see [What is Mission](https://next.api.alibabacloud.com/document/STAROps/2026-04-28/CreateMission).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `blueprint` - (Optional, Set) Blueprint configuration (cron/calendar/event). See [`blueprint`](#blueprint) below.
* `configuration` - (Optional, Set) Runtime configuration. See [`configuration`](#configuration) below.
* `description` - (Optional) Description.
* `digital_employee_name` - (Required) Name of the associated digital employee.
* `display_name` - (Optional) Display name.  
* `enabled` - (Optional, Computed) Whether it is enabled.
* `name` - (Required, ForceNew) Unique identifier of the mission.
* `notification` - (Optional, Set) Notification configuration (an array of strings by channel: dingtalk/feishu/slack/wechat/call/sms/email/webhook). See [`notification`](#notification) below.
* `variables` - (Optional, Map) Mission variables (including workspace or project).

### `blueprint`

The blueprint supports the following:
* `calendar` - (Optional, List) List of iCalendar RRULE-based scheduling Blueprints. See [`calendar`](#blueprint-calendar) below.
* `cron` - (Optional, List) A list of Blueprint resources scheduled by Cron. See [`cron`](#blueprint-cron) below.
* `event` - (Optional, List) List of event-triggered Blueprints. See [`event`](#blueprint-event) below.

### `blueprint-calendar`

The blueprint-calendar supports the following:
* `concurrency_policy` - (Optional, ForceNew, Computed) Concurrency policy.  
* `delay` - (Optional, ForceNew, Computed, Int) Delay in seconds.
* `description` - (Optional, ForceNew) Description.  
* `display_name` - (Optional, ForceNew) Display name.  
* `id` - (Optional, ForceNew) Blueprint ID.
* `priority` - (Optional, ForceNew, Computed, Int) Priority.  
* `prompt` - (Optional, ForceNew) Prompt instruction.
* `rrule` - (Optional, ForceNew) iCalendar RRULE expression.  
* `run_immediately` - (Optional, ForceNew) Run immediately after creation.  
* `time_zone` - (Optional, ForceNew, Computed) Time zone.  
* `timeout_seconds` - (Optional, ForceNew, Int) Timeout in seconds.
* `variables` - (Optional, ForceNew, Map) Variables.

### `blueprint-cron`

The blueprint-cron supports the following:
* `concurrency_policy` - (Optional, ForceNew, Computed) Concurrency policy.
* `cron_expression` - (Optional, ForceNew) Cron expression.  
* `delay` - (Optional, ForceNew, Computed, Int) Delay in seconds.  
* `description` - (Optional, ForceNew) Description.
* `display_name` - (Optional, ForceNew) Display name.  
* `id` - (Optional, ForceNew) Blueprint ID.  
* `priority` - (Optional, ForceNew, Computed, Int) Priority.  
* `prompt` - (Optional, ForceNew) Prompt instruction.  
* `run_immediately` - (Optional, ForceNew) Run immediately after creation.
* `time_zone` - (Optional, ForceNew, Computed) Time zone.
* `timeout_seconds` - (Optional, ForceNew, Int) Timeout in seconds.
* `variables` - (Optional, ForceNew, Map) Variables.

### `blueprint-event`

The blueprint-event supports the following:
* `concurrency_policy` - (Optional, ForceNew, Computed) Concurrency policy.
* `debounce_seconds` - (Optional, ForceNew, Computed, Int) Debounce duration in seconds.
* `description` - (Optional, ForceNew) Description.
* `display_name` - (Optional, ForceNew) Display name.
* `id` - (Optional, ForceNew) Blueprint ID.
* `max_concurrency` - (Optional, ForceNew, Computed, Int) Maximum concurrency.
* `priority` - (Optional, ForceNew, Computed, Int) Priority.
* `prompt` - (Optional, ForceNew) Prompt instruction.
* `timeout_seconds` - (Optional, ForceNew, Int) Timeout in seconds.
* `variables` - (Optional, ForceNew, Map) Variables.
* `workspace` - (Optional, ForceNew) Name of the workspace to which the consumed event belongs.

### `configuration`

The configuration supports the following:
* `credits` - (Optional, Computed, Int) Mission credit limit.

### `notification`

The notification supports the following:
* `call` - (Optional, List) List of mobile phone numbers for voice notifications.
* `dingtalk` - (Optional, List) DingTalk bot ID.
* `email` - (Optional, List) List of email addresses for email notifications.
* `feishu` - (Optional, List) Feishu ID.
* `slack` - (Optional, List) Slack ID.
* `sms` - (Optional, List) List of mobile phone numbers for SMS notifications.
* `webhook` - (Optional, List) List of generic webhook URLs.
* `wechat` - (Optional, List) WeCom ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Mission.
* `delete` - (Defaults to 5 mins) Used when delete the Mission.
* `update` - (Defaults to 5 mins) Used when update the Mission.

## Import

S T A R Ops Mission can be imported using the id, e.g.

```shell
$ terraform import alicloud_s_t_a_r_ops_mission.example <name>
```