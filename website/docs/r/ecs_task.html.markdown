---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_task"
description: |-
  Provides a Alicloud ECS Task resource.
---

# alicloud_ecs_task

Provides a ECS Task resource.



For information about ECS Task and how to use it, see [What is Task](https://next.api.alibabacloud.com/document/Ecs/2014-05-26/DescribeTaskAttribute).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_ecs_task" "default" {
}
```

### Deleting `alicloud_ecs_task` or removing it from your configuration

Terraform cannot destroy resource `alicloud_ecs_task`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `error_code` - Error code.
* `error_msg` - Error message.
* `region_id` - The region ID of the resource.
* `related_item_set` - Resource information type.
  * `name` - Related item name.
  * `value` - Related item value.
* `status` - Operating status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `update` - (Defaults to 5 mins) Used when update the Task.

## Import

ECS Task can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_task.example <task_id>
```