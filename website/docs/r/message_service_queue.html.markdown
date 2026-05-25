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
* `delay_seconds` - (Optional, Computed, Int) The delay duration, in seconds, applied to all messages sent to this queue. Messages become available for consumption only after this specified delay period has elapsed.
Valid values: 0 to 604 800 seconds.
Default value: 0.  
* `dlq_policy` - (Optional, Computed, Set) Dead-letter queue (DLQ) policy for the queue. See [`dlq_policy`](#dlq_policy) below.
* `logging_enabled` - (Optional) Indicates whether log management is enabled. Valid values:
  - true: Enabled.
  - false: Disabled.

Default value: false.  
* `maximum_message_size` - (Optional, Computed, Int) The maximum size of a message body that can be sent to the queue.  
Valid values: 1024 to 65536 bytes.  
Default value: 65536.
* `message_retention_period` - (Optional, Computed, Int) The maximum duration, in seconds, that a message can remain in the queue before it is automatically deleted, regardless of whether it has been consumed. This duration starts from the time the message is sent to the queue.
Valid values: 60 to 604 800 seconds.
Default value: 345 600.  
* `polling_wait_seconds` - (Optional, Computed, Int) The maximum wait time for a ReceiveMessage request when the queue contains no messages.  
Valid values: 0 to 30 seconds.  
Default value: 0.
* `queue_name` - (Required, ForceNew) The name of the queue.  
* `queue_type` - (Optional, ForceNew, Available since v1.280.0) The type of the queue. Valid values:

  - normal: Standard queue

  - fifo: FIFO queue  
* `tags` - (Optional, Map) A list of resource tags.
* `visibility_timeout` - (Optional, Computed, Int) The duration for which a message remains invisible (Inactive) after being retrieved from the queue.  
Valid values: 1 to 43200 seconds.  
Default value: 30.

### `dlq_policy`

The dlq_policy supports the following:
* `dead_letter_target_queue` - (Optional) Target queue to which dead-letter messages are delivered.
* `enabled` - (Optional) Indicates whether delivery of dead-letter messages is enabled.
* `max_receive_count` - (Optional, Int) Maximum number of times a message can be delivered.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the queue.

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