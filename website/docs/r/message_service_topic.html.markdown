---
subcategory: "Message Service"
layout: "alicloud"
page_title: "Alicloud: alicloud_message_service_topic"
description: |-
  Provides a Alicloud Message Service Topic resource.
---

# alicloud_message_service_topic

Provides a Message Service Topic resource.



For information about Message Service Topic and how to use it, see [What is Topic](https://www.alibabacloud.com/help/en/message-service/latest/createtopic).

-> **NOTE:** Available since v1.188.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_message_service_topic" "default" {
  topic_name       = var.name
  max_message_size = 16888
  enable_logging   = true
}
```

## Argument Reference

The following arguments are supported:
* `enable_logging` - (Optional, Computed, Available since v1.241.0) Represents whether the log management function is enabled
* `max_message_size` - (Optional, Computed, Int) Represents the maximum length of the message body sent to the topic
* `tags` - (Optional, Map, Available since v1.241.0) Resource tags, allowing character letters, spaces and numbers, as well as the following special characters 【 -.! @#$%? /^ & *)(+ ={}[]",'~ ·'? :;____]:;_]
* `topic_name` - (Required, ForceNew) The name of the resource
* `topic_type` - (Optional, ForceNew, Computed, Available since v1.283.0) This property does not have a description in the spec, please add it before generating code.

The following arguments will be discarded. Please use new fields as soon as possible:
* `logging_enabled` - (Deprecated since v1.283.0). Field 'logging_enabled' has been deprecated from provider version 1.283.0. New field 'enable_logging' instead.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `last_modify_time` - Represents the time when the topic attribute information was last modified.
* `message_count` - Represents the number of messages in the current topic.
* `message_retention_period` - Represents the longest life time of the message in the topic.
* `topic_inner_url` - TopicInnerUrl.
* `topic_url` - TopicUrl.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Topic.
* `delete` - (Defaults to 5 mins) Used when delete the Topic.
* `update` - (Defaults to 5 mins) Used when update the Topic.

## Import

Message Service Topic can be imported using the id, e.g.

```shell
$ terraform import alicloud_message_service_topic.example <topic_name>
```