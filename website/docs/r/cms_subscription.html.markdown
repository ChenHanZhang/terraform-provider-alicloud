---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_subscription"
description: |-
  Provides a Alicloud Cms Subscription resource.
---

# alicloud_cms_subscription

Provides a Cms Subscription resource.



For information about Cms Subscription and how to use it, see [What is Subscription](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateSubscription).

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `agent_config` - (Optional, ForceNew, Set) Agent configuration   See [`agent_config`](#agent_config) below.
* `description` - (Optional) Description.
* `filter_setting` - (Optional, Set) Filter settings. See [`filter_setting`](#filter_setting) below.
* `notify_strategy_id` - (Optional) Notification strategy ID.  
* `pushing_setting` - (Optional, Set) Push settings.   See [`pushing_setting`](#pushing_setting) below.
* `subscription_name` - (Optional) Subscription name.  
* `workspace` - (Optional) Workspace name.

### `agent_config`

The agent_config supports the following:
* `agent_uuid` - (Optional) Agent UUID  
* `routes` - (Optional, List) Notification configuration   See [`routes`](#agent_config-routes) below.

### `agent_config-routes`

The agent_config-routes supports the following:
* `channels` - (Optional, List) Notification channels See [`channels`](#agent_config-routes-channels) below.
* `effect_time_range` - (Optional, Set) Effective time range See [`effect_time_range`](#agent_config-routes-effect_time_range) below.

### `agent_config-routes-channels`

The agent_config-routes-channels supports the following:
* `channel_type` - (Optional) Channel type  
* `enabled_sub_channels` - (Optional, List) Sub-channel types  
* `receivers` - (Optional, List) Notification recipients  

### `agent_config-routes-effect_time_range`

The agent_config-routes-effect_time_range supports the following:
* `day_in_week` - (Optional, List) Day of the week  
* `time_zone` - (Optional) Time zone  
* `end_time_in_minute` - (Optional, Int) End time in minutes within a day  
* `start_time_in_minute` - (Optional, Int) Start time in minutes within a day  

### `filter_setting`

The filter_setting supports the following:
* `conditions` - (Optional, List) List of filter conditions. See [`conditions`](#filter_setting-conditions) below.
* `expression` - (Optional) Expression.
* `relation` - (Optional) Relationship between conditions.

### `filter_setting-conditions`

The filter_setting-conditions supports the following:
* `field` - (Required) Filter field.
* `op` - (Required) Filter operator.  
* `value` - (Required) Filter value.  

### `pushing_setting`

The pushing_setting supports the following:
* `alert_action_ids` - (Optional, List) List of action plan IDs for alert push notifications.  
* `response_plan_id` - (Optional) Response plan ID.  
* `restore_action_ids` - (Optional, List) List of action plan IDs for restore push notifications.  
* `template_uuid` - (Optional) Template ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Subscription.
* `delete` - (Defaults to 5 mins) Used when delete the Subscription.
* `update` - (Defaults to 5 mins) Used when update the Subscription.

## Import

Cms Subscription can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_subscription.example <subscription_id>
```