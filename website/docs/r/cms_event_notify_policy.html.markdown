---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_event_notify_policy"
description: |-
  Provides a Alicloud Cms Event Notify Policy resource.
---

# alicloud_cms_event_notify_policy

Provides a Cms Event Notify Policy resource.



For information about Cms Event Notify Policy and how to use it, see [What is Event Notify Policy](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateNotifyPolicy).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_cms_event_notify_policy" "default" {
  description = "CloudSpec list operation example for EventNotifyPolicy"
  response_plan {
    repeat_notify_setting {
      end_incident_state = "resolved"
      repeat_interval    = "30"
    }
    auto_recover_seconds = "600"
  }
  notify_strategy {
    description                  = "Notify strategy for list example"
    ignore_restored_notification = false
    grouping_setting {
      grouping_keys = ["severity"]
      period_min    = "5"
      times         = "1"
      silence_sec   = "300"
    }
    routes {
      effect_time_range {
        time_zone            = "Asia/Shanghai"
        start_time_in_minute = "0"
        end_time_in_minute   = "1439"
        day_in_week          = ["1", "2", "3", "4", "5"]
      }
      channels {
        receivers            = ["cspec-example-group"]
        channel_type         = "DING"
        enabled_sub_channels = []
      }
    }
  }
  subscription {
    subscribe_legacy_event = true
    filter_setting {
      relation = "AND"
      conditions {
        field = "severity"
        op    = "EQ"
        value = "CRITICAL"
      }
    }
  }
  workspace = "default-workspace-cn-hangzhou"
  name      = "cspec-list-enp-0716a"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) The description of the policy.
* `enabled` - (Optional) Indicates whether the policy is enabled.
* `name` - (Optional) Filters results by fuzzy matching on the policy name.
* `notify_strategy` - (Optional, Set) The notification strategy sub-entity, which includes grouping and merging settings, notification routing, channels, and custom templates. See [`notify_strategy`](#notify_strategy) below.
* `response_plan` - (Optional, Set) Response plan sub-entities: escalation, repeated notification, automatic recovery, and action integration. See [`response_plan`](#response_plan) below.
* `subscription` - (Optional, Set) Subscription sub-entities: event filtering, cross-workspace routing, and the switch for legacy product event subscription. See [`subscription`](#subscription) below.
* `uuid` - (Optional, ForceNew, Computed) The unique identifier of the notification policy, which is returned by the creation API.
* `workspace` - (Required) The workspace ID, which is used to isolate notification policy resources for different business workspaces. Example: `default-cms-xxxx-cn-hangzhou`.

### `notify_strategy`

The notify_strategy supports the following:
* `custom_template_entries` - (Optional, List) The list of custom notification templates. See [`custom_template_entries`](#notify_strategy-custom_template_entries) below.
* `description` - (Optional) The description.
* `grouping_setting` - (Optional, Set) The grouping and merging settings. See [`grouping_setting`](#notify_strategy-grouping_setting) below.
* `ignore_restored_notification` - (Optional) Indicates whether to ignore notifications for recovered alerts.
* `routes` - (Optional, List) The notification channel routing settings. See [`routes`](#notify_strategy-routes) below.

### `notify_strategy-custom_template_entries`

The notify_strategy-custom_template_entries supports the following:
* `template_uuid` - (Optional) The UUID of the template.

### `notify_strategy-grouping_setting`

The notify_strategy-grouping_setting supports the following:
* `grouping_keys` - (Optional, List) The list of grouping fields.
* `period_min` - (Optional, Int) The check period in minutes.
* `silence_sec` - (Optional, Int) The silence duration in seconds.
* `times` - (Optional, Int) The number of triggers.

### `notify_strategy-routes`

The notify_strategy-routes supports the following:
* `channels` - (Optional, List) The notification channels. See [`channels`](#notify_strategy-routes-channels) below.
* `digital_employee_name` - (Optional) The name of the digital employee. This parameter is required when enableRca is set to true.
* `effect_time_range` - (Optional, Set) The effective time range. See [`effect_time_range`](#notify_strategy-routes-effect_time_range) below.
* `enable_rca` - (Optional) Specifies whether to enable Root Cause Analysis (RCA).
* `filter_setting` - (Optional, Set) Route-level event filtering. See [`filter_setting`](#notify_strategy-routes-filter_setting) below.

### `notify_strategy-routes-channels`

The notify_strategy-routes-channels supports the following:
* `channel_type` - (Optional) The channel type.
* `enabled_sub_channels` - (Optional, List) The enabled notification methods.
* `receivers` - (Optional, List) The list of recipients for the channel.

### `notify_strategy-routes-effect_time_range`

The notify_strategy-routes-effect_time_range supports the following:
* `day_in_week` - (Optional, List) The effective days of the week. Valid values: 0 to 6 (0 indicates Sunday and 6 indicates Saturday).
* `end_time_in_minute` - (Optional, Int) The end time in minutes. Valid values: 0 to 1439.
* `start_time_in_minute` - (Optional, Int) The start time in minutes. Valid values: 0 to 1438.
* `time_zone` - (Optional) The time zone, such as Asia/Shanghai.

### `notify_strategy-routes-filter_setting`

The notify_strategy-routes-filter_setting supports the following:
* `conditions` - (Optional, List) Filter conditions. See [`conditions`](#notify_strategy-routes-filter_setting-conditions) below.
* `expression` - (Optional) The relational expression.
* `relation` - (Optional) The logical relationship between conditions.

### `notify_strategy-routes-filter_setting-conditions`

The notify_strategy-routes-filter_setting-conditions supports the following:
* `field` - (Optional) The filter field.
* `op` - (Optional) The filter operator.
* `value` - (Optional) The filter value.

### `response_plan`

The response_plan supports the following:
* `auto_recover_seconds` - (Optional, Int) The auto-recovery duration in seconds (when no events persist).
* `escalation_id` - (Optional, List) The list of escalation plan IDs.
* `pushing_setting` - (Optional, Set) Action integration push settings. See [`pushing_setting`](#response_plan-pushing_setting) below.
* `repeat_notify_setting` - (Optional, Set) Repeated notification configuration. See [`repeat_notify_setting`](#response_plan-repeat_notify_setting) below.

### `response_plan-pushing_setting`

The response_plan-pushing_setting supports the following:
* `alert_action_ids` - (Optional, List) The list of alert action integration IDs triggered by alerts.
* `restore_action_ids` - (Optional, List) The list of action integration IDs triggered upon recovery.

### `response_plan-repeat_notify_setting`

The response_plan-repeat_notify_setting supports the following:
* `end_incident_state` - (Optional) The state when the incident ends.
* `repeat_interval` - (Optional, Int) The repetition interval in minutes.

### `subscription`

The subscription supports the following:
* `filter_setting` - (Optional, Set) Event content filtering. See [`filter_setting`](#subscription-filter_setting) below.
* `subscribe_legacy_event` - (Optional) Specifies whether to subscribe to legacy product events (events with an empty workspace, such as CMS 1.0, ARMS, or SLS events).
* `workspace_filter_setting` - (Optional, Set) Cross-workspace event routing (global subscription). See [`workspace_filter_setting`](#subscription-workspace_filter_setting) below.

### `subscription-filter_setting`

The subscription-filter_setting supports the following:
* `conditions` - (Optional, List) Filter conditions. See [`conditions`](#subscription-filter_setting-conditions) below.
* `expression` - (Optional) Relational expression.
* `relation` - (Optional) Condition relationship.

### `subscription-workspace_filter_setting`

The subscription-workspace_filter_setting supports the following:
* `tag_selector` - (Optional, Set) The tag selector. See [`tag_selector`](#subscription-workspace_filter_setting-tag_selector) below.
* `workspace_uuids` - (Optional, List) The list of workspace UUIDs.

### `subscription-workspace_filter_setting-tag_selector`

The subscription-workspace_filter_setting-tag_selector supports the following:
* `conditions` - (Optional, List) The filter conditions. See [`conditions`](#subscription-workspace_filter_setting-tag_selector-conditions) below.
* `expression` - (Optional) The relational expression.
* `relation` - (Optional) The condition relation.

### `subscription-workspace_filter_setting-tag_selector-conditions`

The subscription-workspace_filter_setting-tag_selector-conditions supports the following:
* `field` - (Optional) The filter field.
* `op` - (Optional) The filter operator.
* `value` - (Optional) The filter value.

### `subscription-filter_setting-conditions`

The subscription-filter_setting-conditions supports the following:
* `field` - (Optional) Filter field.
* `op` - (Optional) Filter operator.
* `value` - (Optional) The filter value.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<uuid>:<workspace>`.
* `create_time` - The creation time.
* `update_time` - The update time.
* `user_id` - The user ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Event Notify Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Event Notify Policy.
* `update` - (Defaults to 5 mins) Used when update the Event Notify Policy.

## Import

Cms Event Notify Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_event_notify_policy.example <uuid>:<workspace>
```