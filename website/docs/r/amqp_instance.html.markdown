---
subcategory: "RabbitMQ (AMQP)"
layout: "alicloud"
page_title: "Alicloud: alicloud_amqp_instance"
description: |-
  Provides a Alicloud RabbitMQ (AMQP) Instance resource.
---

# alicloud_amqp_instance

Provides a RabbitMQ (AMQP) Instance resource.

The instance of Amqp.

For information about RabbitMQ (AMQP) Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/createinstance).

-> **NOTE:** Available since v1.128.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}

resource "alicloud_amqp_instance" "default" {
  instance_name  = var.name
  instance_type  = "enterprise"
  max_tps        = "1000"
  queue_capacity = "50"
  period_cycle   = "Year"
  support_eip    = "false"
  period         = "1"
  auto_renew     = "true"
  payment_type   = "Subscription"
}
```

### Deleting `alicloud_amqp_instance` or removing it from your configuration

The `alicloud_amqp_instance` resource allows you to manage  `payment_type = "Subscription"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `auto_renew` - (Optional, Available since v1.129.0) The renewal method. Valid values:  
  - true: Auto-renewal  
  - false: Manual renewal

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `edition` - (Optional, Computed, Available since v1.266.0) The deployment architecture for Serverless instances. Valid values:  
  - shared: Shared architecture, applicable to reserved plus elastic (shared) and pay-as-you-go editions.  
  - dedicated: Dedicated architecture, applicable to reserved plus elastic (dedicated) edition.  

-> **NOTE:**  Modifying the Edition parameter triggers instance cluster migration. Before making this change, submit a ticket to the cloud service team. [Submit a Ticket](https://smartservice.console.aliyun.com/service/create-ticket?entrance=100&product=rabbitmq)

* `instance_name` - (Optional, Computed) The name of the instance to be updated. There are no restrictions on the value, but it is recommended that the length does not exceed 64 characters.
* `instance_type` - (Optional, Computed) Instance type.  
Valid values for subscription instances are as follows:  
  - professional: Professional Edition  
  - enterprise: Enterprise Edition  
  - vip: Platinum Edition  

For Serverless instances, this parameter is not required.  
* `max_connections` - (Optional, Computed, Int, Available since v1.129.0) Maximum number of connections.  
Configure this parameter according to the values provided on the [Message Queue for RabbitMQ purchase page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).  
* `max_eip_tps` - (Optional, Computed) Peak public network TPS throughput of the instance, measured in messages per second.  
* `max_tps` - (Optional, Computed) Private network TPS throughput, measured in requests per second.  
Configure this parameter according to the values provided on the [Message Queue for RabbitMQ purchase page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).  
* `modify_type` - (Optional) Specification change type. Valid values:
  - UPGRADE: Upgrade
  - DOWNGRADE: Downgrade

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `payment_type` - (Required, ForceNew) The billing method. Valid values:  
  - Subscription: Prepaid series  
  - PayAsYouGo: Serverless series
* `period` - (Optional, Int) The subscription period, in periodCycle units.  

-> **NOTE:**  This parameter is valid only when PaymentType is set to Subscription. The default value is 1.  


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `period_cycle` - (Optional, Available since v1.129.0) Billing cycle unit for subscription instances. Valid values are as follows:  
  - Month: Month  
  - Year: Year  

This parameter is valid only when PaymentType is set to Subscription. The default value is Month.  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `provisioned_capacity` - (Optional, Int, Available since v1.266.0) The provisioned TPS capacity for reserved plus elastic instances.  
* `queue_capacity` - (Optional, Computed) The maximum number of queues for the instance. Unit: count.
* `renewal_duration` - (Optional, Computed, Int) The auto-renewal duration, in the unit specified by RenewalDurationUnit. The default value is month. Valid values:
  - 1: 1 month  
  - 2: 2 months  
  - 3: 3 months  
  - 6: 6 months  
  - 12: 1 year  
  - 24: 2 years  
  - 36: 3 years  

-> **NOTE:**  This parameter takes effect only when AutoRenew is set to true. The default value is 1.

* `renewal_duration_unit` - (Optional, Computed) The unit of the auto-renewal duration. Valid values:  
  - M: Month  
  - Y: Year  

-> **NOTE:**  This parameter is required when RenewalStatus is set to AutoRenewal.

* `renewal_status` - (Optional, Computed) The renewal status, which is equivalent to autoRenew. You only need to configure one of them. Valid values:  
  - AutoRenewal: Automatic renewal  

-> **NOTE:**  Both renewStatus and this parameter specify the renewal method. If both are specified, renewStatus takes precedence.  

* `serverless_charge_type` - (Optional, Available since v1.129.0) Billing type for pay-as-you-go (Serverless) instances. Valid values:
  - onDemand: Pay-as-you-go billing.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `storage_size` - (Optional, Computed) The message storage space, in GB. Valid values:  
  - Professional Edition and Enterprise Edition instances: Fixed at 0.  

-> **NOTE:**  A value of 0 indicates that storage fees are not charged for Professional Edition and Enterprise Edition instances, not that there is no storage space available.  
  - Platinum Edition instances: m × 100, where m ranges from [7, 28].  
* `support_eip` - (Optional) Indicates whether public network access is supported. Valid values:  
  - true: Public network access is supported.  
  - false: Public network access is not supported.  
* `support_tracing` - (Optional) Specifies whether to enable the message tracing feature. Valid values:  
  - true: Enable message tracing  
  - false: Disable message tracing  

-> **NOTE:**  - Platinum Edition instances provide message tracing for 15 days free of charge. For these instances, message tracing must be enabled and the retention period must be set to 15 days.  
  - For instances of other editions, you can choose to enable or disable message tracing.
* `tracing_storage_time` - (Optional, Computed, Int) The retention period for message tracing data. Unit: days. Valid values:  
  - 3: 3 days  
  - 7: 7 days  
  - 15: 15 days  

This parameter takes effect only when SupportTracing is set to true.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time.  
* `status` - Instance status. Valid values are as follows:  

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 6 mins) Used when create the Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Instance.
* `update` - (Defaults to 200 mins) Used when update the Instance.

## Import

RabbitMQ (AMQP) Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_amqp_instance.example <instance_id>
```