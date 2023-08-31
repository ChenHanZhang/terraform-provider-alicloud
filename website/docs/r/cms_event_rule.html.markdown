---
subcategory: "Cloud Monitor Service"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_event_rule"
description: |-
  Provides a Alicloud Cloud Monitor Service Event Rule resource.
---

# alicloud_cms_event_rule

Provides a Cloud Monitor Service Event Rule resource. 

For information about Cloud Monitor Service Event Rule and how to use it, see [What is Event Rule](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.210.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_cms_monitor_group" "defaultMs3hsi" {
  monitor_group_name = var.name
}

resource "alicloud_cms_monitor_group" "defaultGp1" {
  monitor_group_name = "${var.name}1"
}


resource "alicloud_cms_event_rule" "default" {
  status       = "ENABLED"
  silence_time = "86400"
  event_pattern {
    sql_filter      = "11223ANDxx"
    product         = "ECS"
    level_list      = ["CRITICAL"]
    event_type_list = ["Exception"]
    name_list       = ["Disk:ErrorDetected:Executing"]
    status_list     = ["Normal"]
  }
  description     = "Description1"
  event_rule_name = var.name
  group_id        = alicloud_cms_monitor_group.defaultMs3hsi.group_id
  contact_parameters {
    contact_group_name    = "ContactGroupName111"
    level                 = "3"
    contact_parameters_id = "111"
  }
  open_api_parameters {
    action                 = "PutLogs1"
    arn                    = "acs:openapi:cn-hangzhou:null:log1/PutLogs1/2019-01-01/MyRole1"
    product                = "log1"
    region                 = "cn-hangzhou"
    role                   = "MyRole1"
    version                = "2019-01-01"
    open_api_parameters_id = "22"
  }
  sls_parameters {
    project           = "testproject3"
    log_store         = "testlogstore3"
    arn               = "acs:log:cn-hangzhou::project/testproject3/logstore/testlogstore3"
    region            = "cn-hangzhou"
    sls_parameters_id = "33"
  }
  webhook_parameters {
    url                   = "http://www.aliyun.com1"
    method                = "POST"
    protocol              = "http"
    webhook_parameters_id = "44"
  }
  fc_parameters {
    service_name     = "ServiceName15"
    function_name    = "FunctionNamee5"
    arn              = "acs:log:cn-hangzhou::services/ServiceName15/functions/FunctionNamee5"
    region           = "cn-hangzhou"
    fc_parameters_id = "55"
  }
  mns_parameters {
    queue             = "queue16"
    arn               = "acs:mns:cn-hangzhou::/queues/queue16/messages"
    region            = "cn-hangzhou"
    mns_parameters_id = "66"
  }
}
```

## Argument Reference

The following arguments are supported:
* `contact_parameters` - (Optional) Alarm contact group notification method. See [`contact_parameters`](#contact_parameters) below.
* `description` - (Optional, Available since v1.182.0) The description information of the rule.
* `event_pattern` - (Optional, Available since v1.182.0) Event mode, used to describe the trigger conditions for this event. See [`event_pattern`](#event_pattern) below.
* `event_rule_name` - (Required, ForceNew) The name of the alarm rule.
* `fc_parameters` - (Optional) Function calculation notification method. See [`fc_parameters`](#fc_parameters) below.
* `group_id` - (Optional, Available since v1.182.0) The ID of the application Group.
* `mns_parameters` - (Optional) Message service notification. See [`mns_parameters`](#mns_parameters) below.
* `open_api_parameters` - (Optional) The meaning of the callback API parameter. See [`open_api_parameters`](#open_api_parameters) below.
* `silence_time` - (Optional, Available since v1.182.0) Channel silence time. Unit: Seconds.
* `sls_parameters` - (Optional) Log service (SLS) notification method. See [`sls_parameters`](#sls_parameters) below.
* `status` - (Optional, Computed, Available since v1.182.0) The status of the event alarm rule. Value:ENABLED: ENABLED.DISABLED: DISABLED.
* `webhook_parameters` - (Optional) URL callback notification method. See [`webhook_parameters`](#webhook_parameters) below.

### `contact_parameters`

The contact_parameters supports the following:
* `contact_group_name` - (Optional) The name of the alarm contact group.
* `contact_parameters_id` - (Optional) The unique identifier of the rule sending destination. The value range of N: 1~5.
* `level` - (Optional) Alarm notification level. The value range of N: 1~5. Value:2: telephone, SMS, DingTalk, email3: SMS, DingTalk, email4: DingTalk, mailbox.

### `event_pattern`

The event_pattern supports the following:
* `event_type_list` - (Optional) The list of Event Alarm rule types.
* `level_list` - (Optional) The level of the alarm rule, CRITICAL, warning, and INFO.
* `name_list` - (Optional) The name of the event.
* `product` - (Required) Product name.
* `sql_filter` - (Optional) SQL event filtering. When the event content meets the SQL conditions, an alarm is automatically triggered.DescriptionThe syntax of SQL event filtering is consistent with the query syntax of log service SLS.
* `status_list` - (Optional) The status of the event.

### `fc_parameters`

The fc_parameters supports the following:
* `fc_parameters_id` - (Optional) The unique identifier of the rule sending destination. The value range of N: 1~5.
* `function_name` - (Optional) The name of the function.
* `region` - (Optional) The region of the function Compute Service.
* `service_name` - (Optional) Function calculation service name.

### `mns_parameters`

The mns_parameters supports the following:
* `mns_parameters_id` - (Optional) The unique identifier of the rule sending destination. The value range of N: 1~5.
* `queue` - (Optional) The name of the message queue.
* `region` - (Optional) The region of the message service.
* `topic` - (Optional) The topic of the message service.

### `open_api_parameters`

The open_api_parameters supports the following:
* `action` - (Optional) API name.
* `open_api_parameters_id` - (Optional) The unique identifier of the API callback notification method.
* `product` - (Optional) The ID of the API.
* `region` - (Optional) The region of the resource.
* `role` - (Optional) Role name.
* `version` - (Optional) The version of the API.

### `sls_parameters`

The sls_parameters supports the following:
* `log_store` - (Optional) LogStore name.
* `project` - (Optional) The name of the Project.
* `region` - (Optional) The region of the Project.
* `sls_parameters_id` - (Optional) The unique identifier of the rule sending destination. The value range of N: 1~5.

### `webhook_parameters`

The webhook_parameters supports the following:
* `method` - (Optional) The request method of the HTTP callback. The GET and POST methods are currently supported.
* `protocol` - (Optional) Protocol type.
* `url` - (Optional) The URL address of the callback.
* `webhook_parameters_id` - (Optional) The unique identifier of the rule sending destination. The value range of N: 1~5.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `fc_parameters` - Function calculation notification method.
  * `arn` - ARN resource description.

Format: 'ARN: acs :${ Service }:${ Region }:${ Account }:${ ResourceType}/${ResourceId}'. The meaning of each field is as follows:
  - Service: Cloud product code.
  - Region: Region ID.
  - Account: the ID of the Alibaba cloud Account.
  - ResourceType: resource type.
  - ResourceId: Resource ID.
* `mns_parameters` - Message service notification.
  * `arn` - ARN resource description.

Format: 'ARN: acs :${ Service }:${ Region }:${ Account }:${ ResourceType}/${ResourceId}'. The meaning of each field is as follows:
  - Service: cloud Service Code.
  - Region: Region ID.
  - Account: the ID of the Alibaba cloud Account.
  - ResourceType: resource type.
  - ResourceId: Resource ID.
* `open_api_parameters` - The meaning of the callback API parameter.
  * `arn` - ARN resource description.
Format: 'ARN: acs :${ Service }:${ Region }:${ Account }:${ ResourceType}/${ResourceId}'. The meaning of each field is as follows:
  - Service: Alibaba cloud products.
  - Region: Region ID.
  - Account: the ID of the Alibaba cloud Account.
  - ResourceType: resource type.
  - ResourceId: Resource ID. ARN resource description.
Format: 'ARN: acs :${ Service }:${ Region }:${ Account }:${ ResourceType}/${ResourceId}'. The meaning of each field is as follows:
  - Service: Alibaba cloud products.
  - Region: Region ID.
  - Account: the ID of the Alibaba cloud Account.
  - ResourceType: resource type.
  - ResourceId: Resource ID.
* `sls_parameters` - Log service (SLS) notification method.
  * `arn` - ARN resource description.

Format: 'ARN: acs :${ Service }:${ Region }:${ Account }:${ ResourceType}/${ResourceId}'. The meaning of each field is as follows:
  - Service: cloud Service Code.
  - Region: Region ID.
  - Account: the ID of the Alibaba cloud Account.
  - ResourceType: resource type.
  - ResourceId: Resource ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Event Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Event Rule.
* `update` - (Defaults to 5 mins) Used when update the Event Rule.

## Import

Cloud Monitor Service Event Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_event_rule.example <id>
```