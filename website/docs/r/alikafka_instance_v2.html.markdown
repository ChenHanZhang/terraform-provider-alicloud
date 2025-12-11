---
subcategory: "Ali Kafka"
layout: "alicloud"
page_title: "Alicloud: alicloud_alikafka_instance_v2"
description: |-
  Provides a Alicloud Ali Kafka Instance resource.
---

# alicloud_alikafka_instance_v2

Provides a Ali Kafka Instance resource.

Kafka instance information.

For information about Ali Kafka Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/api-alikafka-2019-09-16-startinstance).

-> **NOTE:** Available since v1.59.0.

## Example Usage

Basic Usage

```terraform
variable "instance_name" {
  default = "terraform-example"
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "default" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default" {
  vpc_id     = alicloud_vpc.default.id
  cidr_block = "172.16.0.0/24"
  zone_id    = data.alicloud_zones.default.zones[0].id
}

resource "alicloud_security_group" "default" {
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_alikafka_instance" "default" {
  name           = "${var.instance_name}-${random_integer.default.result}"
  partition_num  = 50
  disk_type      = 1
  disk_size      = 500
  deploy_type    = 5
  io_max         = 20
  vswitch_id     = alicloud_vswitch.default.id
  security_group = alicloud_security_group.default.id
}
```

### Deleting `alicloud_alikafka_instance_v2` or removing it from your configuration

The `alicloud_alikafka_instance_v2` resource allows you to manage  `paid_type = "0,4"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `config` - (Optional, Computed, Available since v1.112.0) The initial configuration of the deployed message queue Kafka version. The configuration information must be a valid JSON string.

This parameter is empty by default when it is not filled in.

`Config` currently supports the following parameters:
  - **enable. VPC_sasl_ssl**: indicates whether to enable VPC encryption. The values are described as follows:
  - `true`: Enable VPC encryption. If it is turned on, ACL must be turned on at the same time.
  - `false`: The default value. VPC encryption is not enabled.
  - **enable.acl**: whether to enable ACL. The values are described as follows:
  - `true`: Enable ACL.
  - `false`: The default value. ACL is not enabled.
  - **kafka.log.retention.hours**: The maximum retention time of messages when the disk capacity is sufficient. Unit: Hours. Value range \[24, 480], default value **72 * *. When the disk capacity is insufficient (that is, the disk water level reaches 85%), the old message will be deleted in advance to ensure service availability.
  - **kafka.message.max.bytes**: the maximum value of messages that Kafka can send and receive in bytes. Valid value: [1048576, 10485760]. Default value: **1048576 * *. Before modifying the configuration, check whether the modified value matches the corresponding configuration of the production and consumer clients.
* `confluent_config` - (Optional, List, Available since v1.266.0) Confluent component related configuration.


-> **NOTE:**  Must be passed when creating a Confluent series instance.
 See [`confluent_config`](#confluent_config) below.
* `cross_zone` - (Optional, Available since v1.266.0) Whether to cross Availability Zones.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `default_topic_partition_num` - (Optional, Computed, Int, Available since v1.241.0) Adjust the number of automatically created Topic partitions by default.
This value is passed only when the operatevalue is updatePartition.
* `deploy_module` - (Required, Available since v1.266.0) Deployment mode. Value:
  - `vpc`:VPC instance
  - `eip`: public network/VPC instance

The deployment mode of the instance must be consistent with its type. The VPC instance. The deployment mode is **vpc * *. Public network or VPC instance. The deployment mode is **eip * *.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `deploy_type` - (Required, ForceNew, Int) Deployment type. Value:
4: public network/VPC
5:VPC
* `disk_size` - (Optional, Computed, Int) Disk size
* `disk_type` - (Optional, ForceNew, Computed, Int) Disk type
* `duration` - (Optional, Int, Available since v1.266.0) Duration of purchase. The unit is a month. The default value is 1. Value:
  - **1 to 12 * *

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `eip_max` - (Optional, Computed, Int) Public network traffic.
If the DeployType is 4, enter the value.
For the value range, see Billing Description.
If you create a Serverless instance or a Confluent instance, you do not need to pass this parameter.
* `eip_model` - (Optional, Available since v1.266.0) Whether the public network is required. The values are as follows:
  - true: The public network is required.
  - false: No public network is required.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `enable_auto_group` - (Optional, Available since v1.241.0) Whether to enable automatic Group creation
* `enable_auto_topic` - (Optional, Computed, Available since v1.241.0) Currently only these three inputs are supported:
  - enable: enables automatic Topic creation.
  - disable: disables automatic Topic creation.
* `instance_id` - (Optional, ForceNew, Computed, Available since v1.266.0) The ID of the instance.
* `instance_name` - (Optional, Available since v1.266.0) Note name
* `io_max_spec` - (Optional, Computed, Available since v1.201.0) Flow specifications (recommend).
  - Flow peak and flow specification must be optional. When filling in at the same time, the flow specification shall prevail. It is recommended that you only fill in the traffic specifications.
  - Value range. For details, see [Billing Description](~~ 84737 ~~).

-> **NOTE:**  If you create a Serverless instance, you do not need to pass this parameter.

* `is_eip_inner` - (Optional, Available since v1.266.0) Whether to support EIP. Value:
  - `true`: public network/VPC instance
  - `false`:VPC instance

The value of this parameter must be the same as the instance type. For example, if the instance type is VPC, the value must be set to **false * *.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `is_force_selected_zones` - (Optional, Available since v1.266.0) Whether to force deployment in the selected zone

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `is_set_user_and_password` - (Optional, Available since v1.266.0) Whether to set a new user name and password. Value:
  - `true`: Set a new user name and password.
  - `false`: Do not set a new username and password.

Only public/VPC instances are supported.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `kms_key_id` - (Optional, ForceNew, Available since v1.180.0) The key ID of the disk encryption in the same region. You can use the [Key Management Service Console](https://kms.console.aliyun.com/? spm = a2c4g.11186623.2.5.336745 b8hfiU21) View the key ID and create a new key. For more information, see [Management Key](~~ 181610 ~~).
This parameter indicates that instance encryption is enabled (it cannot be turned off after it is enabled). When this interface is called, the system checks whether to create a service Associated role. If it is not created, the service Associated role is automatically created. For more information, see [Service Associated Roles](~~ 190460 ~~).
* `notifier` - (Optional, Available since v1.266.0) Alert contacts

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `order_id` - (Optional, ForceNew, Available since v1.266.0) The order ID. You can click [order order](https://usercenter2.aliyun.com/order/list? pageIndex = 1 & pageSize = 20). 
[Order Management](https://usercenter2-intl.aliyun.com/order/list? pageIndex = 1 & pageSize = 20 & spm = 5176.12818093.top-nav.ditem-ord. 36f016d0oqfmja). 

-> **NOTE:** Serverless instances do not support querying by order ID.> 

* `paid_type` - (Optional, Int) Type of payment. Value:
  - 0: Prepaid
  - 1: Pay-as-you-go for Reserved Instances.
  - 3: pay after reserving Serverless instance specifications + pay after Serverless auto scaling.
  - 4:Confluent Series Prepaid

* `partition_num_of_buy` - (Optional, Int, Available since v1.266.0) PartitionNumOfBuy
* `password` - (Optional, Available since v1.266.0) User password.
Only public/VPC instances are supported.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `resource_group_id` - (Optional, Computed, Available since v1.224.0) The ID of the resource group
* `security_group` - (Optional, ForceNew, Computed, Available since v1.93.0) The security group of the instance.
If you do not specify this parameter, Kafka automatically configures a security group for your instance. To complete this parameter, you must first create a security group for the instance. For more information, see [Create a security group](~~ 25468 ~~).
* `selected_zones` - (Optional) Select the candidate set of the deployed primary zone and the two-dimensional array of the candidate set of the secondary zone. Custom code('zone {zone} ') and standard code('zone-cn-RegionID-{zone}') are supported.
  - If you want to deploy across zones (isCrossZone = true), and the candidate set of the primary zone is zone H or zone F, and the candidate set of the secondary zone is zone K, the input parameters are '[["zoneh", "zonef"],["zonek"]]'.

-> **NOTE:**  When multiple zones are filled in the primary or secondary zone, it indicates that one of the zones has no priority relationship. For example, '[["zoneh", "zonef"],["zonek"]]' is used to deploy the primary zone H or F, and the secondary zone K.
  - When you do not want to deploy across zones (isCrossZone = false), the zone is deployed in zone K. For example, if the custom code is used, pass in the parameter '[["zonek"],[]]'. Note that you still need to pass in two arrays, and the second array indicating the candidate set of the spare zone is empty [].

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `serverless_config` - (Optional, List, Available since v1.266.0) The settings of the Serverless instance. Must be passed when creating a Serverless instance. See [`serverless_config`](#serverless_config) below.
* `service_version` - (Optional, Computed, Available since v1.112.0) The version of Cloud message queue for Kafka.
The reserved instance value is 2.2.0 or 2.6.2.
The value of the Serverless instance is 3.3.1.
The optional value for the Confluent family is 7.4.0.
Default:
Reserved Instance Family: 2.2.0
Serverless instance series: 3.3.1
Confluent Series: 7.4.0
* `spec_type` - (Optional) Specification type.
When the PaidType parameter is set to 0 (for a prepaid instance), the value is as follows:
normal: normal Edition (High Writing Edition)
professional: professional Edition (High Writing Edition)
Professional for Highread: Professional Edition (High Reading Edition)

When the PaidType parameter is 1 (Pay-As-You-Go Reserved Instance), the value is as follows:
normal: Standard Edition (High Writing Edition)
professional: professional Edition (High Writing Edition)
Professional for Highread: Professional Edition (High Reading Edition)
When the PaidType parameter is set to 3 (pay after Serverless instance reservation + pay after Serverless auto scaling):
normal:Serverless Standard Edition
When the PaidType parameter is set to 4 (Confluent instance subscription), the value is as follows:
professional: professional Edition
enterprise: enterprise Edition
For more information about the specifications, see [Billing Description](~~ 84737 ~~).
* `status` - (Optional, Computed) The status of the resource
* `tags` - (Optional, Map, Available since v1.63.0) The tag of the kafka console, which is used to group instance,topic, and consumption.
* `update_default_topic_partition_num` - (Optional, Available since v1.266.0) Whether to modify the default value of the automatically created Topic. Note that this parameter and the operation parameter cannot be passed at the same time.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `user_phone_num` - (Optional, Available since v1.266.0) Early warning contact mobile phone number

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `username` - (Optional, Available since v1.266.0) User name.
Only public/VPC instances are supported.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `vswitch_id` - (Required, ForceNew) Switch id
* `vswitch_ids` - (Optional, ForceNew, Computed, List, Available since v1.241.0) The vSwitch ID of the instance deployment. This parameter is required for Reserved Instances and Serverless instances. Confluent instances support this parameter, and at least one item is required for VSwitchIds and VSwitchIds. If both items are specified, VSwitchIds is preferred.
* `vpc_id` - (Optional, ForceNew, Computed, Available since v1.185.0) VpcId
* `zone_id` - (Optional, ForceNew, Computed, Available since v1.185.0) The ID of the zone in which the instance is deployed.
  - Must be the zone ID of vSwitch.
  - The format can be zoneX or RegionId-X. For example, it can be set to zonea or cn-hangzhou-k.

### `confluent_config`

The confluent_config supports the following:
* `connect_cu` - (Optional, Int, Available since v1.266.0) The number of CPU cores of the Connect component.
* `connect_replica` - (Optional, Int, Available since v1.266.0) The number of copies of the Connect component.
* `control_center_cu` - (Optional, Int, Available since v1.266.0) The number of CPU cores of the ControlCenter component.
* `control_center_replica` - (Optional, Int, Available since v1.266.0) The number of copies of the ControlCenter component.
* `control_center_storage` - (Optional, Int, Available since v1.266.0) The disk capacity of the ControlCenter component. Unit of GB.
* `kafka_cu` - (Optional, Int, Available since v1.266.0) The number of Kafka Broker CPU cores.
* `kafka_replica` - (Optional, Int, Available since v1.266.0) Number of copies of Kafka Broker.
* `kafka_rest_proxy_cu` - (Optional, Int, Available since v1.266.0) The number of CPU cores of the KafkaRestProxy component.
* `kafka_rest_proxy_replica` - (Optional, Int, Available since v1.266.0) Number of copies of the KafkaRestProxy component.
* `kafka_storage` - (Optional, Int, Available since v1.266.0) The disk capacity of the Kafka Broker. Unit of GB.
* `ksql_cu` - (Optional, Int, Available since v1.266.0) KsqlDB component
The number of CPU cores.
* `ksql_replica` - (Optional, Int, Available since v1.266.0) Number of copies of the KsqlDB component.
* `ksql_storage` - (Optional, Int, Available since v1.266.0) KsqlDB component disk capacity. Unit of GB.
* `schema_registry_cu` - (Optional, Int, Available since v1.266.0) The number of CPU cores of the schemaregery component.
* `schema_registry_replica` - (Optional, Int, Available since v1.266.0) The number of copies of the schemaregery component.
* `zoo_keeper_cu` - (Optional, Int, Available since v1.266.0) The number of CPU cores of the ZooKeeper component.
* `zoo_keeper_replica` - (Optional, Int, Available since v1.266.0) Number of copies of the ZooKeeper component.
* `zoo_keeper_storage` - (Optional, Int, Available since v1.266.0) ZooKeeper component disk capacity. Unit of GB.

### `serverless_config`

The serverless_config supports the following:
* `reserved_publish_capacity` - (Optional, Int, Available since v1.266.0) The specification of the reserved sending traffic. Only the incoming integer is supported. The minimum value is 60.


-> **NOTE:**  The actual upper limit is affected by the current regional inventory. Please refer to the optional range on the sales page.

* `reserved_subscribe_capacity` - (Optional, Int, Available since v1.266.0) The reserved consumption traffic specification. Only an integer can be passed in. The minimum value is 20.

-> **NOTE:**  The actual upper limit is affected by the current regional inventory. Please refer to the optional range on the sales page.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - Creation time
* `domain_endpoint` - The domain name of the default access point. Cloud message queue Kafka version instances support domain name access points and IP access points., in the format of '{instance domain name }:{ port number}'., in the format of '{IP address of the Broker }:{ port number}'.
* `end_point` - Access point
* `group_left` - GroupLeft
* `group_used` - GroupUsed
* `is_partition_buy` - isPartitionBuy
* `partition_left` - PartitionLeft
* `partition_used` - PartitionUsed
* `region_id` - RegionId
* `sasl_domain_endpoint` - The domain name access point of the SASL access point. The Kafka instance supports domain name access points and IP access points.'{instance domain name }:{ port number}'.{Broker IP address }:{ port number} '.
* `ssl_domain_endpoint` - The domain name of the SSL access point. The Kafka instance supports domain name access points and IP access points.'{instance domain name }:{ port number}'.{Broker IP address }:{ port number} '.
* `topic_left` - topic left
* `topic_num_of_buy` - TopicNumOfBuy
* `topic_used` - topic used

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 33 mins) Used when create the Instance.
* `delete` - (Defaults to 55 mins) Used when delete the Instance.
* `update` - (Defaults to 55 mins) Used when update the Instance.

## Import

Ali Kafka Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_alikafka_instance_v2.example <id>
```