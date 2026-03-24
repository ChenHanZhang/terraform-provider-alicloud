---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_alert_robot"
description: |-
  Provides a Alicloud Cms Alert Robot resource.
---

# alicloud_cms_alert_robot

Provides a Cms Alert Robot resource.



For information about Cms Alert Robot and how to use it, see [What is Alert Robot](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateAlertRobot).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_cms_alert_robot" "default" {
  type             = "DING"
  alert_robot_name = "resourceTest"
  alert_robot_id   = "resourceTest"
  lang             = "zh_CN"
  url              = "www.aaa.com"
}
```

## Argument Reference

The following arguments are supported:
* `alert_robot_id` - (Optional, ForceNew, Computed) The unique ID of the robot.
* `alert_robot_name` - (Optional) The name of the robot.
* `digital_employee_name` - (Optional) Associated digital employee.
* `lang` - (Optional) Language.
* `type` - (Required, ForceNew) The type of robot.
* `url` - (Optional) The robot's webhook URL.
* `workspace` - (Optional, ForceNew) Workspace.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Alert Robot.
* `delete` - (Defaults to 5 mins) Used when delete the Alert Robot.
* `update` - (Defaults to 5 mins) Used when update the Alert Robot.

## Import

Cms Alert Robot can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_alert_robot.example <alert_robot_id>
```