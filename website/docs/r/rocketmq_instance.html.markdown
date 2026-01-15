---
subcategory: "RocketMQ"
layout: "alicloud"
page_title: "Alicloud: alicloud_rocketmq_instance"
description: |-
  Provides a Alicloud Rocketmq Instance resource.
---

# alicloud_rocketmq_instance

Provides a Rocketmq Instance resource.

RocketMQ instance resources.

For information about Rocketmq Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/developer-reference/api-rocketmq-2022-08-01-createinstance).

-> **NOTE:** Available since v1.212.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

data "alicloud_resource_manager_resource_groups" "default" {
  status = "OK"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "createVPC" {
  description = "example"
  cidr_block  = "172.16.0.0/12"
  vpc_name    = var.name
}

resource "alicloud_vswitch" "createVSwitch" {
  description  = "example"
  vpc_id       = alicloud_vpc.createVPC.id
  cidr_block   = "172.16.0.0/24"
  vswitch_name = var.name
  zone_id      = data.alicloud_zones.default.zones.0.id
}

resource "alicloud_rocketmq_instance" "default" {
  product_info {
    msg_process_spec       = "rmq.u2.10xlarge"
    send_receive_ratio     = "0.3"
    message_retention_time = "70"
  }
  service_code      = "rmq"
  payment_type      = "PayAsYouGo"
  instance_name     = var.name
  sub_series_code   = "cluster_ha"
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
  remark            = "example"
  ip_whitelists     = ["192.168.0.0/16", "10.10.0.0/16", "172.168.0.0/16"]
  software {
    maintain_time = "02:00-06:00"
  }
  tags = {
    Created = "TF"
    For     = "example"
  }
  series_code = "ultimate"
  network_info {
    vpc_info {
      vpc_id = alicloud_vpc.createVPC.id
      vswitches {
        vswitch_id = alicloud_vswitch.createVSwitch.id
      }
    }
    internet_info {
      internet_spec      = "enable"
      flow_out_type      = "payByBandwidth"
      flow_out_bandwidth = "30"
    }
  }
}
```

### Deleting `alicloud_rocketmq_instance` or removing it from your configuration

The `alicloud_rocketmq_instance` resource allows you to manage  `payment_type = "PayAsYouGo"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `acl_info` - (Optional, Computed, List, Available since v1.245.0) access control info. See [`acl_info`](#acl_info) below.
* `auto_renew` - (Optional, ForceNew) Whether to enable auto-renewal. This parameter is only applicable when the payment type for the instance is Subscription (prepaid).
  - true: Enable auto-renewal
  - false: Disable auto-renewal
* `auto_renew_period` - (Optional, ForceNew, Int) Auto-renewal period. This parameter is only valid when auto-renewal is enabled. Unit: months.

The values can be as follows:
  - Monthly renewal: 1, 2, 3, 6, 12
* `commodity_code` - (Optional, ForceNew, Computed, Available since v1.231.0) Commodity code (cn station):
ons_rmqsub_public_cn: Subscription instance
ons_rmqpost_public_cn: Pay-As-You-Go instance
ons_rmqsrvlesspost_public_cn: Serverless instance
Commodity code (International station):
ons_rmqsub_public_intl: Subscription instance
ons_rmqpost_public_intl: Pay-As-You-Go instance
ons_rmqsrvlesspost_public_intl: Serverless instance
serverless instance This parameter is required
* `instance_name` - (Optional) The name of instance
* `ip_whitelists` - (Optional, Computed, List, Available since v1.245.0) Ip whitelist array
* `network_info` - (Required, List) Instance network configuration information See [`network_info`](#network_info) below.
* `payment_type` - (Required, ForceNew) The payment type for the instance. Alibaba Cloud Message Queue RocketMQ version supports two types of payment:

The parameter values are as follows:
  - PayAsYouGo: Pay-as-you-go, a post-payment model where you pay after usage.
  - Subscription: Subscription-based, a pre-payment model where you pay before usage. 

For more information, please refer to [Billing Methods](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/overview-2).
* `period` - (Optional, Int) Duration of purchase. This parameter is only valid when the payment type for the instance is Subscription (prepaid).

The values can be as follows:
  - Monthly purchase: 1, 2, 3, 4, 5, 6
  - Annual purchase: 1, 2, 3

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `period_unit` - (Optional, Computed) The minimum periodic unit for the duration of purchase.

The parameter values are as follows:
  - Month: Purchase on a monthly basis
  - Year: Purchase on an annual basis

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `product_info` - (Optional, ForceNew, List) product info See [`product_info`](#product_info) below.
* `remark` - (Optional) Custom description
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `series_code` - (Required, ForceNew) The primary series encoding for the instance. For specific differences between the primary series, please refer to [Product Selection](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/instance-selection).

The parameter values are as follows:
  - standard: Standard Edition
  - ultimate: Platinum Edition
  - professional: Professional Edition
* `service_code` - (Required, ForceNew) The code of the service code instance. The code of the RocketMQ is rmq.
* `software` - (Optional, Computed, List) Instance software information. See [`software`](#software) below.
* `sub_series_code` - (Required, ForceNew) The sub-series encoding for the instance. For specific differences between the sub-series, please refer to [Product Selection](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/instance-selection).

The parameter values are as follows:
  - cluster_ha: Cluster High Availability Edition
  - serverless：Serverless instance

When selecting the primary series as ultimate (Platinum Edition), the sub-series can only be chosen as cluster_ha (Cluster High Availability Edition).
* `tags` - (Optional, Map) The resource label.

### `acl_info`

The acl_info supports the following:
* `acl_types` - (Optional, Computed, List) Supported Access Control Types
* `default_vpc_auth_free` - (Optional, Computed) Whether the VPC is accessed without secret access under the intelligent identification type
  - true to enable secret-free access
  - false to turn off secret-free access

### `network_info`

The network_info supports the following:
* `internet_info` - (Required, ForceNew, List) instance internet info. See [`internet_info`](#network_info-internet_info) below.
* `vpc_info` - (Required, ForceNew, List) Proprietary network information See [`vpc_info`](#network_info-vpc_info) below.

### `network_info-internet_info`

The network_info-internet_info supports the following:
* `flow_out_bandwidth` - (Optional, ForceNew, Int) Public network bandwidth specification. Unit: Mb/s.
This field should only be filled when the public network billing type is set to payByBandwidth.
The value range is [1 - 1000].
* `flow_out_type` - (Required, ForceNew) Public network billing type.

Parameter values are as follows:
  - payByBandwidth: Fixed bandwidth billing. 
  - uninvolved: Not involved. This parameter must be set to the value when public network access is disabled.
  - payByTraffic: charges are based on traffic volume.
When the public network is open, serverless instances are billed using payByTraffic, while non-serverless instances are billed using payByBandwidth
* `internet_spec` - (Required, ForceNew) Whether to enable public network access.

The parameter values are as follows:
  - enable: Enable public network access
  - disable: Disable public network access 

Instances by default support VPC access. If public network access is enabled, Alibaba Cloud Message Queue RocketMQ version will incur charges for public network outbound bandwidth. For specific billing information, please refer to [Public Network Access Fees](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/internet-access-fee).
* `ip_whitelist` - (Optional, List, Deprecated since v1.269.0) internet ip whitelist.

### `network_info-vpc_info`

The network_info-vpc_info supports the following:
* `security_group_ids` - (Optional, ForceNew, Available since v1.231.0) Security group id
* `vswitch_id` - (Optional, ForceNew, Computed, Deprecated since v1.269.0) VPC network switch
* `vswitches` - (Optional, ForceNew, Computed, List, Available since v1.231.0) Multiple VSwitches. At least two VSwitches are required for a serverless instance. See [`vswitches`](#network_info-vpc_info-vswitches) below.
* `vpc_id` - (Required, ForceNew) Proprietary Network

### `network_info-vpc_info-vswitches`

The network_info-vpc_info-vswitches supports the following:
* `vswitch_id` - (Optional, ForceNew, Computed) The id of the vSwitch. Only Serverless instances support multiple instances. Other types of instances support only one

### `product_info`

The product_info supports the following:
* `auto_scaling` - (Optional, Computed) is open auto scaling
* `message_retention_time` - (Optional, Int) Duration of message retention. Unit: hours.
For the range of values, please refer to [Usage Limits](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/usage-limits)>Resource Quotas>Limitations on Message Retention.
The message storage in AlibabaCloud RocketMQ is fully implemented in a serverless and elastic manner, with charges based on the actual storage space. You can control the storage capacity of messages by adjusting the duration of message retention. For more information, please see [Storage Fees](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/storage-fees).
* `msg_process_spec` - (Optional, ForceNew) Message sending and receiving calculation specifications. For details about the upper limit for sending and receiving messages, see [Instance Specifications](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/product-overview/instance-specifications).
* `send_receive_ratio` - (Optional, Computed, Float) message send receive ratio.
Value range: [0.2, 0.5]
* `storage_encryption` - (Optional, ForceNew, Available since v1.245.0) Enable storage encryption.
* `storage_secret_key` - (Optional, ForceNew, Available since v1.245.0) Storage secret key.
* `trace_on` - (Optional, Computed) Whether to enable the message trace function.
true: Enable message trace function
false: Disable message trace function
This parameter takes effect only for Serverless instances. If the message trace function is enabled, related message trace charges will be incurred. For specific billing information, see Serverless billing instructions. For Pay-As-You-Go and Subscription instances, this parameter does not take effect, that is, regardless of whether the parameter is enabled, the message trace function is supported by default.

### `software`

The software supports the following:
* `maintain_time` - (Optional, Computed) Upgrade time period.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource
* `network_info` - Instance network configuration information
  * `endpoints` - Access point list
    * `endpoint_type` - Access point type
    * `endpoint_url` - Access point address
    * `ip_white_list` - White list of access addresses
* `product_info` - product info
  * `support_auto_scaling` - is support auto scaling.
* `region_id` - The region ID of the resource
* `software` - Instance software information.
  * `software_version` - Software version.
  * `upgrade_method` - Upgrade method.
* `status` - The status of the instance

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 35 mins) Used when create the Instance.
* `delete` - (Defaults to 35 mins) Used when delete the Instance.
* `update` - (Defaults to 35 mins) Used when update the Instance.

## Import

Rocketmq Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_rocketmq_instance.example <instance_id>
```