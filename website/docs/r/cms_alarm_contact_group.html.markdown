---
subcategory: "Cloud Monitor Service"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_alarm_contact_group"
description: |-
  Provides a Alicloud Cloud Monitor Service Alarm Contact Group resource.
---

# alicloud_cms_alarm_contact_group

Provides a Cloud Monitor Service Alarm Contact Group resource.

Describes the alarm Contact Group set by the user.

For information about Cloud Monitor Service Alarm Contact Group and how to use it, see [What is Alarm Contact Group](https://www.alibabacloud.com/help/en/cloudmonitor/latest/putcontactgroup).

-> **NOTE:** Available since v1.101.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_cms_alarm_contact_group" "example" {
  alarm_contact_group_name = "tf-example"
}
```

## Argument Reference

The following arguments are supported:
* `alarm_contact_group_name` - (Required, ForceNew) AlarmContactGroupName
* `contact_names` - (Optional, List, Available since v1.269.0) The name of the alarm contact. The value range of N: 1~100.
* `describe` - (Optional) Description of the alarm Contact Group
* `enable_subscribed` - (Optional, Computed) Whether to enable the weekly subscription function. Value:
  - true: Yes.
  - false: No.

-> **NOTE:**  Currently, the weekly subscription function can be enabled only when the ECS instance is equal to or greater than 5.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Alarm Contact Group.
* `delete` - (Defaults to 5 mins) Used when delete the Alarm Contact Group.
* `update` - (Defaults to 5 mins) Used when update the Alarm Contact Group.

## Import

Cloud Monitor Service Alarm Contact Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_alarm_contact_group.example <alarm_contact_group_name>
```