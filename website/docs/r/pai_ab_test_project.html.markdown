---
subcategory: "Pai Ab Test"
layout: "alicloud"
page_title: "Alicloud: alicloud_pai_ab_test_project"
description: |-
  Provides a Alicloud Pai Ab Test Project resource.
---

# alicloud_pai_ab_test_project

Provides a Pai Ab Test Project resource.



For information about Pai Ab Test Project and how to use it, see [What is Project](https://next.api.alibabacloud.com/document/PAIABTest/2024-01-19/CreateProject).

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


resource "alicloud_pai_ab_example_project" "default" {
  description  = "this is a example project"
  project_name = "project_1"
  workspace_id = "45699"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Description
* `project_name` - (Required) Project Name
* `workspace_id` - (Required, ForceNew) Workspace ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Project.
* `delete` - (Defaults to 5 mins) Used when delete the Project.
* `update` - (Defaults to 5 mins) Used when update the Project.

## Import

Pai Ab Test Project can be imported using the id, e.g.

```shell
$ terraform import alicloud_pai_ab_test_project.example <project_id>
```