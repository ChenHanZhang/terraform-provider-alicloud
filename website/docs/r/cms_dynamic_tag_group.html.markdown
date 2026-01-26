---
subcategory: "Cloud Monitor Service"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_dynamic_tag_group"
description: |-
  Provides a Alicloud Cloud Monitor Service Dynamic Tag Group resource.
---

# alicloud_cms_dynamic_tag_group

Provides a Cloud Monitor Service Dynamic Tag Group resource.



For information about Cloud Monitor Service Dynamic Tag Group and how to use it, see [What is Dynamic Tag Group](https://www.alibabacloud.com/help/en/cloudmonitor/latest/createdynamictaggroup).

-> **NOTE:** Available since v1.142.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_cms_alarm_contact_group" "default" {
  alarm_contact_group_name = var.name
}

resource "alicloud_cms_dynamic_tag_group" "default" {
  tag_key            = var.name
  contact_group_list = [alicloud_cms_alarm_contact_group.default.id]
  match_express {
    tag_value                = var.name
    tag_value_match_function = "all"
  }
}
```

## Argument Reference

The following arguments are supported:
* `contact_group_list` - (Required, ForceNew, List) Alarm contact group. The value range of N is 1~100. The alarm notification of the application group is sent to the alarm contact in the alarm contact group.

-> **NOTE:**  An alarm Contact Group is a group of alarm contacts that can contain one or more alarm contacts. For more information about how to create alarm contacts and alarm contact groups, see [PutContact](~~ 114923 ~~) and [PutContactGroup](~~ 114929 ~~).

* `enable_install_agent` - (Optional, Available since v1.269.0) Whether the application Group is enabled to automatically install the cloud monitoring plug-in. CloudMonitor automatically installs the CloudMonitor plug-in for hosts in the application Group. Value:
  - true: on.
  - false (default): Off.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `enable_subscribe_event` - (Optional, Available since v1.269.0) Whether the application group turns on automatic subscription event notification. CloudMonitor sends alarm notifications when serious and warning-level events occur in resources in the application Group. Value:
  - true: on.
  - false (default): Off.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `match_express` - (Required, ForceNew, List) The label generates a matching expression that applies the grouping. See [`match_express`](#match_express) below.
* `match_express_filter_relation` - (Optional, ForceNew, Computed) The relationship between conditional expressions. The value is:
  -'and': Relationship
  -'or': or relationship

-> **NOTE:**  currently, only one combination condition is supported, and Alibaba Cloud will support multiple combination conditions in the future.

* `tag_key` - (Required, ForceNew) Label key.
* `tag_region_id` - (Optional, ForceNew, Available since v1.269.0) The ID of the region to which the tag belongs.
* `template_id_list` - (Optional, ForceNew, List) Alarm template ID list

### `match_express`

The match_express supports the following:
* `tag_name` - (Optional, ForceNew, Available since v1.269.0) The Key of the Tag used to create the group. If there are multiple resources containing this Key, the matching resources will be added to the same group according to the filtering conditions according to the same Key-Value.
* `tag_value` - (Optional, ForceNew) Resource label value. The value of N is 1.

-> **NOTE:**  parameters' matchexpress. N.TagValueMatchFunction' and'matchexpress. N.TagValue' must be set at the same time.

* `tag_value_match_function` - (Optional, ForceNew) The matching method of the resource label value. The value of N is 1. Value:
  - contains: contains.
  - startWith: prefix.
  - endWith: suffix.
  - notContains: not included.
  - equals: equal.
  - all: all.

-> **NOTE:**  parameters' matchexpress. N.TagValueMatchFunction' and'matchexpress. N.TagValue' must be set at the same time.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Dynamic Tag Group.
* `delete` - (Defaults to 5 mins) Used when delete the Dynamic Tag Group.

## Import

Cloud Monitor Service Dynamic Tag Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_dynamic_tag_group.example <dynamic_tag_rule_id>
```