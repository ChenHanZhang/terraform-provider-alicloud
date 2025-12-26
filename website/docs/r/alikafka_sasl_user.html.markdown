---
subcategory: "AliKafka"
layout: "alicloud"
page_title: "Alicloud: alicloud_alikafka_sasl_user"
description: |-
  Provides a Alicloud Ali Kafka Sasl User resource.
---

# alicloud_alikafka_sasl_user

Provides a Ali Kafka Sasl User resource.



For information about Ali Kafka Sasl User and how to use it, see [What is Sasl User](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/api-alikafka-2019-09-16-createsasluser).

-> **NOTE:** Available since v1.66.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.4.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "10.4.0.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_zones.default.zones.0.id
}

resource "alicloud_security_group" "default" {
  vpc_id = alicloud_vpc.default.id
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

resource "alicloud_alikafka_instance" "default" {
  name            = "${var.name}-${random_integer.default.result}"
  partition_num   = 50
  disk_type       = "1"
  disk_size       = "500"
  deploy_type     = "5"
  io_max          = "20"
  spec_type       = "professional"
  service_version = "2.2.0"
  vswitch_id      = alicloud_vswitch.default.id
  security_group  = alicloud_security_group.default.id
  config          = <<EOF
  {
    "enable.acl": "true"
  }
  EOF
}

resource "alicloud_alikafka_sasl_user" "default" {
  instance_id = alicloud_alikafka_instance.default.id
  username    = var.name
  password    = "tf_example123"
}
```

## Argument Reference

The following arguments are supported:
* `instance_id` - (Required, ForceNew) The instance ID.
* `mechanism` - (Optional, ForceNew, Computed, Available since v1.266.0) The encryption method. Valid values:

*   SCRAM-SHA-512. This is the default value.
*   SCRAM-SHA-256

-> **NOTE:**   This parameter is available only for serverless ApsaraMQ for Kafka instances.

* `password` - (Required) The password.
* `type` - (Optional, ForceNew, Computed, Available since v1.159.0) Type. Value:
  - `plain`: A simple user name and password verification mechanism. Kafka optimizes the PLAIN mechanism and supports the dynamic addition of SASL users without restarting the instance.
  - `scram`: A user name and password verification mechanism with higher security than PLAIN. Message Queue for Kafka uses SCRAM-SHA-256.
  - `LDAP`: This user type is displayed only for Confluent instances.

The default value is **plain * *.
* `username` - (Required, ForceNew) The user name.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<instance_id>:<username>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Sasl User.
* `delete` - (Defaults to 5 mins) Used when delete the Sasl User.
* `update` - (Defaults to 5 mins) Used when update the Sasl User.

## Import

Ali Kafka Sasl User can be imported using the id, e.g.

```shell
$ terraform import alicloud_alikafka_sasl_user.example <instance_id>:<username>
```