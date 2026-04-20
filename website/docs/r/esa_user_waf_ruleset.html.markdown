---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_user_waf_ruleset"
description: |-
  Provides a Alicloud ESA User Waf Ruleset resource.
---

# alicloud_esa_user_waf_ruleset

Provides a ESA User Waf Ruleset resource.

Global configuration - Global WAF  .

For information about ESA User Waf Ruleset and how to use it, see [What is User Waf Ruleset](https://next.api.alibabacloud.com/document/ESA/2024-09-10/CreateUserWafRuleset).

-> **NOTE:** Available since v1.277.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `description` - (Optional) The description of the WAF ruleset.
* `expression` - (Required) WAF ruleset expression.  
* `instance_id` - (Required, ForceNew) The instance ID.
* `name` - (Required) WAF rule set name.
* `phase` - (Required, ForceNew) The execution phase of WAF rules.  
  - http_whitelist: Allowlist rules  
  - http_custom: Custom rules  
  - http_managed: Managed rules  
  - http_anti_scan: Scan protection rules  
  - http_ratelimit: Rate limiting rules  
  - ip_access_rule: IP access rules  
  - http_bot: Advanced bot management rules  
  - http_security_level_rule: Security level rules
* `rules` - (Optional, List) A list of rule configurations under the WAF ruleset. See [`rules`](#rules) below.
* `shared` - (Optional, ForceNew, Set) The sharing configuration of the WAF ruleset. See [`shared`](#shared) below.
* `status` - (Required) The status of the WAF ruleset.

### `rules`

The rules supports the following:
* `config` - (Optional, Set) WAF rule configuration. See [`config`](#rules-config) below.
* `position` - (Optional, Int) WAF rule position.  

### `rules-config`

The rules-config supports the following:
* `action` - (Optional) Action to perform.
* `actions` - (Optional, ForceNew, Set) Extended actions. See [`actions`](#rules-config-actions) below.
* `app_package` - (Optional, ForceNew, Set) Repackaging detection. See [`app_package`](#rules-config-app_package) below.
* `app_sdk` - (Optional, ForceNew, Set) App SDK. See [`app_sdk`](#rules-config-app_sdk) below.
* `expression` - (Optional, ForceNew) Expression.  
* `id` - (Optional, Int) Rule ID.  
* `managed_group_id` - (Optional, ForceNew, Int) Managed rule group ID.
* `managed_list` - (Optional, ForceNew) Name of the list.
* `managed_rulesets` - (Optional, ForceNew, List) List of managed rulesets. See [`managed_rulesets`](#rules-config-managed_rulesets) below.
* `name` - (Optional) Rule name.
* `notes` - (Optional, ForceNew) Notes.
* `rate_limit` - (Optional, ForceNew, Set) Rate limiting.   See [`rate_limit`](#rules-config-rate_limit) below.
* `security_level` - (Optional, ForceNew, Set) Security level.   See [`security_level`](#rules-config-security_level) below.
* `sigchl` - (Optional, ForceNew, List) Token verification.
* `status` - (Optional) Rule status.
* `timer` - (Optional, Set) Timer. See [`timer`](#rules-config-timer) below.
* `type` - (Optional) Rule type.
* `value` - (Optional, ForceNew) IP access control value.  

### `rules-config-actions`

The rules-config-actions supports the following:
* `bypass` - (Optional, ForceNew, Set) Skip modules. See [`bypass`](#rules-config-actions-bypass) below.
* `response` - (Optional, ForceNew, Set) Custom response page. See [`response`](#rules-config-actions-response) below.

### `rules-config-app_package`

The rules-config-app_package supports the following:
* `package_signs` - (Optional, ForceNew, List) Repackaging detection. See [`package_signs`](#rules-config-app_package-package_signs) below.

### `rules-config-app_sdk`

The rules-config-app_sdk supports the following:
* `custom_sign` - (Optional, ForceNew, Set) Custom signing field. See [`custom_sign`](#rules-config-app_sdk-custom_sign) below.
* `custom_sign_status` - (Optional, ForceNew) Custom signing field switch.
* `feature_abnormal` - (Optional, ForceNew, List) Feature anomaly.  

### `rules-config-managed_rulesets`

The rules-config-managed_rulesets supports the following:
* `action` - (Optional, ForceNew) Action to perform.
* `attack_type` - (Optional, ForceNew, Int) Attack type.
* `managed_rules` - (Optional, ForceNew, List) List of managed rules. See [`managed_rules`](#rules-config-managed_rulesets-managed_rules) below.
* `number_enabled` - (Optional, ForceNew, Int) Number of enabled rules.
* `number_total` - (Optional, ForceNew, Int) Total number of rules.
* `protection_level` - (Optional, ForceNew, Int) Protection level.

### `rules-config-rate_limit`

The rules-config-rate_limit supports the following:
* `characteristics` - (Optional, ForceNew, Set) Specify the object for rate statistics.   See [`characteristics`](#rules-config-rate_limit-characteristics) below.
* `interval` - (Optional, ForceNew, Int) Statistical duration.
* `on_hit` - (Optional, ForceNew) Apply to requests that hit the cache.
* `ttl` - (Optional, ForceNew, Int) Timeout period.
* `threshold` - (Optional, ForceNew, Set) Threshold. See [`threshold`](#rules-config-rate_limit-threshold) below.

### `rules-config-security_level`

The rules-config-security_level supports the following:
* `value` - (Optional, ForceNew) Value of the security level.  

### `rules-config-timer`

The rules-config-timer supports the following:
* `periods` - (Optional, ForceNew, List) Effective time periods. See [`periods`](#rules-config-timer-periods) below.
* `scopes` - (Optional, ForceNew) Schedule type: permanently effective (permanent or empty), effective during specified periods (periods), or recurring weekly (weekly).  
* `weekly_periods` - (Optional, ForceNew, List) Weekly effective time periods. See [`weekly_periods`](#rules-config-timer-weekly_periods) below.
* `zone` - (Optional, ForceNew, Int) Time zone. Defaults to UTC+00:00 if not specified.Example: 8 indicates UTC+8, and - 8 indicates UTC-8.Range: - 12 to +14.  

### `rules-config-timer-periods`

The rules-config-timer-periods supports the following:
* `end` - (Optional, ForceNew) End time, in RFC3339-formatted UTC time.  
* `start` - (Optional, ForceNew) Start time, in RFC3339-formatted UTC time.  

### `rules-config-timer-weekly_periods`

The rules-config-timer-weekly_periods supports the following:
* `daily_periods` - (Optional, ForceNew, List) Effective time periods within the specified days. See [`daily_periods`](#rules-config-timer-weekly_periods-daily_periods) below.
* `days` - (Optional, ForceNew) Days of the week, separated by commas. Values 1–7 represent Monday through Sunday, respectively.Example: Monday and Wednesday are represented as "1,3".

### `rules-config-timer-weekly_periods-daily_periods`

The rules-config-timer-weekly_periods-daily_periods supports the following:
* `end` - (Optional, ForceNew) End time in HH:mm:ss format.  
* `start` - (Optional, ForceNew) Start time in HH:mm:ss format.

### `rules-config-rate_limit-characteristics`

The rules-config-rate_limit-characteristics supports the following:
* `criteria` - (Optional, ForceNew, List) Logic list. See [`criteria`](#rules-config-rate_limit-characteristics-criteria) below.
* `logic` - (Optional, ForceNew) Logical relationship.
* `match_type` - (Optional, ForceNew) Matching field.  

### `rules-config-rate_limit-threshold`

The rules-config-rate_limit-threshold supports the following:
* `distinct_managed_rules` - (Optional, ForceNew, Int) Threshold for the number of distinct managed rules.
* `managed_rules_blocked` - (Optional, ForceNew, Int) Threshold for hits on managed rules.
* `request` - (Optional, ForceNew, Int) Request threshold.
* `response_status` - (Optional, ForceNew, Set) Response status code threshold. See [`response_status`](#rules-config-rate_limit-threshold-response_status) below.
* `traffic` - (Optional, ForceNew) Traffic threshold.

### `rules-config-rate_limit-threshold-response_status`

The rules-config-rate_limit-threshold-response_status supports the following:
* `code` - (Optional, ForceNew, Int) HTTP response status code.
* `count` - (Optional, ForceNew, Int) Threshold for the number of occurrences of a response status code.
* `ratio` - (Optional, ForceNew, Int) Response status code percentage.

### `rules-config-rate_limit-characteristics-criteria`

The rules-config-rate_limit-characteristics-criteria supports the following:
* `criteria` - (Optional, ForceNew, List) Logic list. See [`criteria`](#rules-config-rate_limit-characteristics-criteria-criteria) below.
* `logic` - (Optional, ForceNew) Logical relationship.
* `match_type` - (Optional, ForceNew) Match field.

### `rules-config-rate_limit-characteristics-criteria-criteria`

The rules-config-rate_limit-characteristics-criteria-criteria supports the following:
* `criteria` - (Optional, ForceNew, List) Logic list.   See [`criteria`](#rules-config-rate_limit-characteristics-criteria-criteria-criteria) below.
* `logic` - (Optional, ForceNew) Logical relationship.
* `match_type` - (Optional, ForceNew) Match field.

### `rules-config-rate_limit-characteristics-criteria-criteria-criteria`

The rules-config-rate_limit-characteristics-criteria-criteria-criteria supports the following:
* `match_type` - (Optional, ForceNew) Match field.  

### `rules-config-managed_rulesets-managed_rules`

The rules-config-managed_rulesets-managed_rules supports the following:
* `action` - (Optional, ForceNew) Action of the managed rule.
* `id` - (Optional, ForceNew, Int) ID of the managed rule.
* `status` - (Optional, ForceNew) Status of the managed rule.

### `rules-config-app_sdk-custom_sign`

The rules-config-app_sdk-custom_sign supports the following:
* `key` - (Optional, ForceNew) Field name.
* `value` - (Optional, ForceNew) Field value.

### `rules-config-app_package-package_signs`

The rules-config-app_package-package_signs supports the following:
* `name` - (Optional, ForceNew) Specify a valid package name.
* `sign` - (Optional, ForceNew) Package signature.

### `rules-config-actions-bypass`

The rules-config-actions-bypass supports the following:
* `custom_rules` - (Optional, ForceNew, List) A list of custom rule IDs.
* `regular_rules` - (Optional, ForceNew, List) A list of managed rule IDs.
* `regular_types` - (Optional, ForceNew, List) List of managed rule types.
* `skip` - (Optional, ForceNew) The type of modules to skip.
* `tags` - (Optional, List) List of modules to bypass.

### `rules-config-actions-response`

The rules-config-actions-response supports the following:
* `code` - (Optional, ForceNew, Int) Custom response code.
* `id` - (Optional, ForceNew, Int) Custom response page ID.

### `shared`

The shared supports the following:
* `action` - (Optional, ForceNew) Action.
* `actions` - (Optional, ForceNew, Set) Action extensions.   See [`actions`](#shared-actions) below.
* `cross_site_id` - (Optional, ForceNew, Int) Specifies the cross-site ID.  
* `expression` - (Optional, ForceNew) Expression.  
* `match` - (Optional, ForceNew, Set) Matching engine.   See [`match`](#shared-match) below.
* `mode` - (Optional) Web SDK integration mode: automatic (automatic) or manual (manual).  
* `name` - (Optional, ForceNew) Ruleset name.  
* `target` - (Optional) Protected target type: web or app.

### `shared-actions`

The shared-actions supports the following:
* `response` - (Optional, ForceNew, Set) Custom response.   See [`response`](#shared-actions-response) below.

### `shared-match`

The shared-match supports the following:
* `convert_to_lower` - (Optional, ForceNew) The value is case-insensitive.
* `criteria` - (Optional, ForceNew, List) Logic list.   See [`criteria`](#shared-match-criteria) below.
* `logic` - (Optional, ForceNew) Logical relationship.
* `match_operator` - (Optional, ForceNew) Match operator.
* `match_type` - (Optional, ForceNew) Matching field.  
* `match_value` - (Optional, ForceNew) Match value.
* `negate` - (Optional, ForceNew) Invert the match result.

### `shared-match-criteria`

The shared-match-criteria supports the following:
* `convert_to_lower` - (Optional, ForceNew) The value is case-insensitive.
* `criteria` - (Optional, ForceNew, List) Logic list.   See [`criteria`](#shared-match-criteria-criteria) below.
* `logic` - (Optional, ForceNew) Logical relationship.
* `match_operator` - (Optional, ForceNew) Match operator.  
* `match_type` - (Optional, ForceNew) Matching field.  
* `match_value` - (Optional, ForceNew) Match value.  
* `negate` - (Optional, ForceNew) Invert the match result.  

### `shared-match-criteria-criteria`

The shared-match-criteria-criteria supports the following:
* `convert_to_lower` - (Optional, ForceNew) Value is case-insensitive.  
* `criteria` - (Optional, ForceNew, List) A logical list.   See [`criteria`](#shared-match-criteria-criteria-criteria) below.
* `logic` - (Optional, ForceNew) Logical relationship.  
* `match_operator` - (Optional, ForceNew) Match operator.  
* `match_type` - (Optional, ForceNew) Matching field.
* `match_value` - (Optional, ForceNew) Match value.  
* `negate` - (Optional, ForceNew) Inverts the match result.

### `shared-match-criteria-criteria-criteria`

The shared-match-criteria-criteria-criteria supports the following:
* `convert_to_lower` - (Optional, ForceNew) The value is case-insensitive.  
* `match_operator` - (Optional, ForceNew) Match operator.  
* `match_type` - (Optional, ForceNew) Match field.  
* `match_value` - (Optional, ForceNew) Match value.
* `negate` - (Optional, ForceNew) Inverts the match result.  

### `shared-actions-response`

The shared-actions-response supports the following:
* `code` - (Optional, ForceNew, Int) Custom response code.  
* `id` - (Optional, ForceNew, Int) Custom response page ID.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<id>:<instance_id>`.
* `id` - WAF ruleset ID.
* `rules` - A list of rule configurations under the WAF ruleset.
  * `characteristics_fields` - List of WAF rule statistical fields.
  * `fields` - List of WAF rule matching fields.
  * `id` - WAF rule ID.
  * `name` - WAF rule name.
  * `phase` - WAF phase.
  * `ruleset_id` - WAF ruleset ID.
  * `skip` - Method for skipping WAF rules.
  * `status` - WAF rule status.
  * `tags` - Stage at which WAF rules are skipped.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the User Waf Ruleset.
* `delete` - (Defaults to 5 mins) Used when delete the User Waf Ruleset.
* `update` - (Defaults to 5 mins) Used when update the User Waf Ruleset.

## Import

ESA User Waf Ruleset can be imported using the id, e.g.

```shell
$ terraform import alicloud_esa_user_waf_ruleset.example <id>:<instance_id>
```