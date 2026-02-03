---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_user_alarm_config"
description: |-
  Provides a Alicloud Cloud Firewall User Alarm Config resource.
---

# alicloud_cloud_firewall_user_alarm_config

Provides a Cloud Firewall User Alarm Config resource.

Configure alarm notifications and contacts.

For information about Cloud Firewall User Alarm Config and how to use it, see [What is User Alarm Config](https://next.api.alibabacloud.com/document/Cloudfw/2017-12-07/DescribeUserAlarmConfig).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


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
    notify_value = "18858120562"
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
* `alarm_config` - (Required, List) This property does not have a description in the spec, please add it before generating code. See [`alarm_config`](#alarm_config) below.
* `alarm_lang` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `contact_config` - (Optional, List) This property does not have a description in the spec, please add it before generating code. See [`contact_config`](#contact_config) below.
* `lang` - (Optional) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `notify_config` - (Optional, List) This property does not have a description in the spec, please add it before generating code. See [`notify_config`](#notify_config) below.
* `use_default_contact` - (Optional) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.


### `alarm_config`

The alarm_config supports the following:
* `alarm_hour` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `alarm_notify` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `alarm_period` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `alarm_type` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `alarm_value` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `alarm_week_day` - (Optional) This property does not have a description in the spec, please add it before generating code.

### `contact_config`

The contact_config supports the following:
* `email` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `mobile_phone` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `name` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `status` - (Optional) This property does not have a description in the spec, please add it before generating code.

### `notify_config`

The notify_config supports the following:
* `notify_type` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `notify_value` - (Optional) This property does not have a description in the spec, please add it before generating code.

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