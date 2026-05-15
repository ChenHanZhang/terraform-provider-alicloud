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
* `description` - (Optional) Description
* `filter_setting` - (Optional, Set) Filter Setting See [`filter_setting`](#filter_setting) below.
* `notify_strategy_id` - (Optional) Notify strategy ID
* `pushing_setting` - (Optional, Set) Pushing Setting See [`pushing_setting`](#pushing_setting) below.
* `subscription_name` - (Optional) Name
* `workspace` - (Optional) The private property for list operation

### `filter_setting`

The filter_setting supports the following:
* `conditions` - (Optional, List) Filter condition list See [`conditions`](#filter_setting-conditions) below.
* `expression` - (Optional) Expression
* `relation` - (Optional) Relation of conditions

### `filter_setting-conditions`

The filter_setting-conditions supports the following:
* `field` - (Required) Filter Field
* `op` - (Required) Filter operator
* `value` - (Required) Filter Value

### `pushing_setting`

The pushing_setting supports the following:
* `alert_action_ids` - (Optional, List) Alert pushing action plan list
* `response_plan_id` - (Optional) response plan Id
* `restore_action_ids` - (Optional, List) Restore pushing action plan list
* `template_uuid` - (Optional) Template ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Create Time.

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