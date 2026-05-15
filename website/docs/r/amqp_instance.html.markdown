---
subcategory: "RabbitMQ (AMQP)"
layout: "alicloud"
page_title: "Alicloud: alicloud_amqp_instance"
description: |-
  Provides a Alicloud RabbitMQ (AMQP) Instance resource.
---

# alicloud_amqp_instance

Provides a RabbitMQ (AMQP) Instance resource.

AMQP instance  .

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
* `auth_model` - (Optional, Available since v1.279.0) Authentication mode. Valid values: ram (RAM mode) and openSource (open-source mode).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_renew` - (Optional, Available since v1.129.0) The renewal method. Valid values:  
  - true: Auto-renewal  
  - false: Manual renewal

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `edition` - (Optional, Computed, Available since v1.266.0) The deployment architecture of the Serverless instance. Valid values:
  - shared: Shared architecture, applicable to the reserved + elastic (shared) and pay-as-you-go editions.
  - dedicated: Dedicated architecture, applicable to the reserved + elastic (dedicated) edition.

-> **NOTE:**  Modifying the Edition requires migrating the instance cluster. Before making changes, please submit a ticket to the cloud service team. [Submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket?entrance=100&product=rabbitmq)

* `instance_name` - (Optional, Computed) The name of the instance to be updated. There are no restrictions on the value, but it is recommended to keep the length under 64 characters.
* `instance_type` - (Optional, Computed) Instance type.  
For the prepaid series, valid values are:  
  - professional: Professional Edition  
  - enterprise: Enterprise Edition  
  - vip: Platinum Edition  

For Serverless series instances, this parameter is not required.  
* `listener_mode` - (Optional, ForceNew, Computed, Available since v1.274.0) ListenerMode specifies the listener mode:  
SSL_ONLY: Only port 5671 is enabled for the instance.  
TCP_AND_SSL: Both ports 5671 and 5672 are enabled for the instance.
* `max_connections` - (Optional, Computed, Int, Available since v1.129.0) Maximum number of connections.  
Set this parameter according to the values provided on the [Message Queue for RabbitMQ - Purchase Page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).  
* `max_eip_tps` - (Optional, Computed) Peak public network TPS traffic of the instance, in messages per second.  
* `max_tps` - (Optional, Computed) Private network TPS traffic, in requests per second.  
Set this parameter according to the values provided on the [Message Queue for RabbitMQ - Purchase Page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).  
* `modify_type` - (Optional) The configuration change type. Valid values:
  - UPGRADE: Upgrade
  - DOWNGRADE: Downgrade.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `payment_type` - (Required, ForceNew) Payment type. Valid values:  
  - Subscription: Prepaid series  
  - PayAsYouGo: Serverless series  
* `period` - (Optional, Int) The prepaid period, in periodCycle units.

-> **NOTE:**  This parameter is valid only when PaymentType is Subscription. The default value is 1.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `period_cycle` - (Optional, Available since v1.129.0) Prepaid billing cycle unit. Valid values:  
  - Month  
  - Year  

This parameter is valid only when PaymentType is set to Subscription. Default value: Month.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `provisioned_capacity` - (Optional, Int, Available since v1.266.0) The provisioned TPS capacity for reserved plus elastic instances.
* `queue_capacity` - (Optional, Computed) The maximum number of queues allowed for the instance. Unit: queues.
* `renewal_duration` - (Optional, Computed, Int) The auto-renewal duration, in the unit specified by RenewalDurationUnit. The default unit is months. Valid values are as follows:
  - 1: 1 month  
  - 2: 2 months  
  - 3: 3 months  
  - 6: 6 months  
  - 12: 1 year  
  - 24: 2 years  
  - 36: 3 years  

-> **NOTE:**  This parameter takes effect only when AutoRenew is set to true. The default value is 1.

* `renewal_duration_unit` - (Optional, Computed) Auto-renewal period unit. Valid values:  
  - M: Month  
  - Y: Year  

-> **NOTE:**  This parameter is required when RenewalStatus is set to AutoRenewal.  

* `renewal_status` - (Optional, Computed) The renewal status, equivalent to autoRenew. You only need to configure one of these two parameters. Valid values:
  - AutoRenewal: Automatic renewal.

-> **NOTE:**  Both renewalStatus and autoRenew indicate the renewal method. If both are specified, renewalStatus takes precedence.

* `resource_group_id` - (Optional, Computed, Available since v1.279.0) The ID of the resource group to which the resource belongs.
* `security_group_id` - (Required, ForceNew) You must specify a SecurityGroupId when creating an instance to create a PrivateLink endpoint.
* `serverless_charge_type` - (Optional, Available since v1.129.0) The billing type for postpaid (Serverless) instances. Valid values:
  - onDemand: Pay-as-you-go.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `serverless_switch` - (Optional, Computed, Available since v1.279.0) The toggle status for the elasticity capability of the Serverless instance.
* `storage_size` - (Optional, Computed) The message storage space, in GB. Valid values:
  - Professional Edition and Enterprise Edition instances: Fixed at 0.

-> **NOTE:**  A value of 0 indicates that storage fees are not charged for Professional Edition and Enterprise Edition instances, not that storage space is unavailable.
  - Platinum Edition instances: m × 100, where m ranges from 7 to 28 (inclusive).
* `support_eip` - (Optional) Indicates whether public network access is supported. Valid values:
  - true: Public network access is supported.
  - false: Public network access is not supported.
* `support_tracing` - (Optional) Specifies whether to enable the message tracing feature. Valid values:  
  - true: Enable message tracing  
  - false: Disable message tracing  

-> **NOTE:**  - Platinum Edition instances provide message tracing for 15 days free of charge. For these instances, message tracing must be enabled, and the retention period must be set to 15 days.  
  - For instances of other editions, you can choose to enable or disable message tracing.
* `tags` - (Optional, ForceNew, Map, Available since v1.279.0) List of tags  
* `tracing_storage_time` - (Optional, Computed, Int) The retention period for message tracing data, in days. Valid values are as follows:  
  - 3: 3 days  
  - 7: 7 days  
  - 15: 15 days  

This parameter takes effect only when SupportTracing is set to true.
* `vpc_id` - (Required, ForceNew, Available since v1.274.0) When creating an instance, you must specify a VpcId to create a PrivateLink endpoint.
* `vswitch_ids` - (Required, ForceNew, List, Available since v1.274.0) When creating an instance, you must specify VSwitchIds to create a PrivateLink endpoint.  
Except for regions with only one availability zone, you can specify only two VSwitchIds.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time.
* `status` - Instance status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 9 mins) Used when create the Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Instance.
* `update` - (Defaults to 9990 mins) Used when update the Instance.

## Import

RabbitMQ (AMQP) Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_amqp_instance.example <instance_id>
```