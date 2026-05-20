---
subcategory: "RabbitMQ (AMQP)"
layout: "alicloud"
page_title: "Alicloud: alicloud_amqp_open_source_permission"
description: |-
  Provides a Alicloud RabbitMQ (AMQP) Open Source Permission resource.
---

# alicloud_amqp_open_source_permission

Provides a RabbitMQ (AMQP) Open Source Permission resource.

Permissions in the open-source authentication and permission management system.

For information about RabbitMQ (AMQP) Open Source Permission and how to use it, see [What is Open Source Permission](https://next.api.alibabacloud.com/document/amqp-open/2019-12-12/CreateOpenSourcePermission).

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "instance_name" {
  default = "example开源鉴权实例"
}

variable "vhost" {
  default = "/"
}

variable "user_name" {
  default = "Suhao123_WithPer"
}

resource "alicloud_amqp_instance" "CreateInstance" {
  period_cycle  = "Month"
  instance_name = var.instance_name
}


resource "alicloud_amqp_open_source_permission" "default" {
  write       = ".*"
  read        = ".*"
  vhost       = var.vhost
  user_name   = var.user_name
  instance_id = alicloud_amqp_instance.CreateInstance.id
  configure   = ".*"
}
```

## Argument Reference

The following arguments are supported:
* `configure` - (Optional) Permission configuration, such as .*
* `instance_id` - (Required, ForceNew) Instance ID
* `read` - (Optional) Read permission, such as .*
* `user_name` - (Required, ForceNew) Username
* `vhost` - (Optional, ForceNew, Computed) Vhost of the instance
* `write` - (Optional) Write permission, such as .*

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<user_name>:<vhost>:<instance_id>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Open Source Permission.
* `delete` - (Defaults to 5 mins) Used when delete the Open Source Permission.
* `update` - (Defaults to 5 mins) Used when update the Open Source Permission.

## Import

RabbitMQ (AMQP) Open Source Permission can be imported using the id, e.g.

```shell
$ terraform import alicloud_amqp_open_source_permission.example <user_name>:<vhost>:<instance_id>
```