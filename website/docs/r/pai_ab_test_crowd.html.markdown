---
subcategory: "Pai Ab Test"
layout: "alicloud"
page_title: "Alicloud: alicloud_pai_ab_test_crowd"
description: |-
  Provides a Alicloud Pai Ab Test Crowd resource.
---

# alicloud_pai_ab_test_crowd

Provides a Pai Ab Test Crowd resource.



For information about Pai Ab Test Crowd and how to use it, see [What is Crowd](https://next.api.alibabacloud.com/document/PAIABTest/2024-01-19/CreateCrowd).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_pai_ab_example_crowd" "default" {
  crowd_name   = "crowd-1"
  description  = "this is a example crowd"
  workspace_id = "45699"
  users        = "uid1,uid2"
  label        = "example"
}
```

## Argument Reference

The following arguments are supported:
* `crowd_name` - (Required) Crowd Name
* `description` - (Optional) Description
* `label` - (Optional) Crowd Label
* `users` - (Required) Users included in the crowd
* `workspace_id` - (Required, ForceNew) Workspace ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Crowd.
* `delete` - (Defaults to 5 mins) Used when delete the Crowd.
* `update` - (Defaults to 5 mins) Used when update the Crowd.

## Import

Pai Ab Test Crowd can be imported using the id, e.g.

```shell
$ terraform import alicloud_pai_ab_test_crowd.example <crowd_id>
```