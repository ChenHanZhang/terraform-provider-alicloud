---
subcategory: "Quotas"
layout: "alicloud"
page_title: "Alicloud: alicloud_quotas_quota_alarm"
sidebar_current: "docs-alicloud-resource-quotas-quota-alarm"
description: |-
  Provides a Alicloud Quotas Quota Alarm resource.
---

# alicloud_quotas_quota_alarm

Provides a Quotas Quota Alarm resource. .

For information about Quotas Quota Alarm and how to use it, see [What is Quota Alarm](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.206.0+.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-testacc-example"
}


resource "alicloud_quotas_quota_alarm" "default" {
  quota_action_code = "q_desktop-count"
  quota_dimensions {
    key   = "regionId"
    value = "cn-hangzhou"
  }
  threshold_percent = 80
  product_code      = "gws"
  quota_alarm_name  = var.name
  web_hook          = "https://oapi.dingtalk.com/robot/send?access_token=0a09bd617f43d07e8607b258c6cdffbacf0e023f1bbe46cfeb0265127802bf43"
  threshold_type    = "used"
}
```

## Argument Reference

The following arguments are supported:
* `product_code` - (Required, ForceNew, Available in v1.116.0+) ProductCode.
* `quota_action_code` - (Required, ForceNew, Available in v1.116.0+) QuotaActionCode.
* `quota_alarm_name` - (Required, Available in v1.116.0+) QuotaAlarmName.
* `quota_dimensions` - (Optional, ForceNew, Available in v1.116.0+) QuotaDimensions.See the following `Block QuotaDimensions`.
* `threshold` - (Optional, Available in v1.116.0+) Threshold.
* `threshold_percent` - (Optional, Available in v1.116.0+) ThresholdPercent.
* `threshold_type` - (Optional, Computed) Quota alarm type. Value:-used: Quota used alarm.-usable: alarm for the remaining available quota.
* `web_hook` - (Optional, Available in v1.116.0+) WebHook.


#### Block QuotaDimensions

The QuotaDimensions supports the following:
* `key` - (Optional, ForceNew) Key.
* `value` - (Optional, ForceNew) Value.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `alarm_id` - The first ID of the resource.
* `create_time` - The creation time of the resource.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Quota Alarm.
* `delete` - (Defaults to 5 mins) Used when delete the Quota Alarm.
* `update` - (Defaults to 5 mins) Used when update the Quota Alarm.

## Import

Quotas Quota Alarm can be imported using the id, e.g.

```shell
$ terraform import alicloud_quotas_quota_alarm.example <id>
```