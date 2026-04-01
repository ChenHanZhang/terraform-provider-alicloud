---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_user_alarm_config"
description: |-
  Provides a Alicloud Cloud Firewall User Alarm Config resource.
---

# alicloud_cloud_firewall_user_alarm_config

Provides a Cloud Firewall User Alarm Config resource.

Configure alert notifications and contacts.

For information about Cloud Firewall User Alarm Config and how to use it, see [What is User Alarm Config](https://next.api.alibabacloud.com/document/Cloudfw/2017-12-07/DescribeUserAlarmConfig).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_cloud_firewall_user_alarm_config" "default" {
  alarm_config {
    alarm_value    = "on"
    alarm_type     = "bandwidth"
    alarm_period   = "1"
    alarm_hour     = "0"
    alarm_notify   = "0"
    alarm_week_day = "0"
  }
  use_default_contact = "1"
  notify_config {
    notify_value = "13000000000"
    notify_type  = "sms"
  }
  alarm_lang = "zh"
  lang       = "zh"
}
```

### Deleting `alicloud_cloud_firewall_user_alarm_config` or removing it from your configuration

Terraform cannot destroy resource `alicloud_cloud_firewall_user_alarm_config`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `alarm_config` - (Required, List) Alert configuration.   See [`alarm_config`](#alarm_config) below.
* `alarm_lang` - (Optional) Language for message notifications.  
* `contact_config` - (Optional, Computed, List) Contact information.   See [`contact_config`](#contact_config) below.
* `lang` - (Optional) The language used for sending and receiving messages.  

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `use_default_contact` - (Optional) Use default contact method.  

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.


### `alarm_config`

The alarm_config supports the following:
* `alarm_hour` - (Optional, Computed) The hour for sending alert notifications.  
* `alarm_notify` - (Optional, Computed) Notification method.
* `alarm_period` - (Optional, Computed) The alert period.  
* `alarm_type` - (Optional, Computed) The alert type.
* `alarm_value` - (Optional, Computed) The alert notification message.  
* `alarm_week_day` - (Optional, Computed) Days of the week for alert notifications.

### `contact_config`

The contact_config supports the following:
* `email` - (Optional, Computed) Email address.  
* `mobile_phone` - (Optional, Computed) Mobile phone number.  
* `name` - (Optional, Computed) Recipient of alert notifications.  
* `status` - (Optional, Computed) Alert status.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<Alibaba Cloud Account ID>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `update` - (Defaults to 5 mins) Used when update the User Alarm Config.

## Import

Cloud Firewall User Alarm Config can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_user_alarm_config.example <Alibaba Cloud Account ID>
```