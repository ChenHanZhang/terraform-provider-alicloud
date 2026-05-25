---
subcategory: "Message Service"
layout: "alicloud"
page_title: "Alicloud: alicloud_message_service_queue"
description: |-
  Provides a Alicloud Message Service Queue resource.
---

# alicloud_message_service_queue

Provides a Message Service Queue resource.



For information about Message Service Queue and how to use it, see [What is Queue](https://www.alibabacloud.com/help/en/message-service/latest/createqueue).

-> **NOTE:** Available since v1.188.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_message_service_queue" "default" {
  queue_name               = var.name
  delay_seconds            = "2"
  polling_wait_seconds     = "2"
  message_retention_period = "566"
  maximum_message_size     = "1126"
  visibility_timeout       = "30"
}
```

## Argument Reference

The following arguments are supported:
* `delay_seconds` - (Optional, Computed, Int) The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
* `dlq_policy` - (Optional, Computed, Set) Queue dead letter policy See [`dlq_policy`](#dlq_policy) below.
* `logging_enabled` - (Optional) Specifies whether to enable the logging feature. Valid values:

  - True
  - False (default)
* `maximum_message_size` - (Optional, Computed, Int) The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
* `message_retention_period` - (Optional, Computed, Int) The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
* `polling_wait_seconds` - (Optional, Computed, Int) The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0
* `queue_name` - (Required, ForceNew) The name of the queue.
* `tags` - (Optional, Map) Resource tags, allowing character letters, spaces and numbers, as well as the following special characters 【 -.! @#$%? /^ & *)(+ ={}[]",'~ ·'? :;____]:;_]
* `visibility_timeout` - (Optional, Computed, Int) The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.

### `dlq_policy`

The dlq_policy supports the following:
* `dead_letter_target_queue` - (Optional) Target Dead Letter Queue Name
* `enabled` - (Optional) Dead letter policy activation status
* `max_receive_count` - (Optional, Int) After the number of received messages reaches MaxReceiveCount, the message enters the dead-letter queue.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Represents the time when the Queue was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Queue.
* `delete` - (Defaults to 5 mins) Used when delete the Queue.
* `update` - (Defaults to 5 mins) Used when update the Queue.

## Import

Message Service Queue can be imported using the id, e.g.

```shell
$ terraform import alicloud_message_service_queue.example <queue_name>
```