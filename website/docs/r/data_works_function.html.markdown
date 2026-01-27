---
subcategory: "Data Works"
layout: "alicloud"
page_title: "Alicloud: alicloud_data_works_function"
description: |-
  Provides a Alicloud Data Works Function resource.
---

# alicloud_data_works_function

Provides a Data Works Function resource.

The UDF of the computing engine.

For information about Data Works Function and how to use it, see [What is Function](https://next.api.alibabacloud.com/document/dataworks-public/2024-05-18/CreateFunction).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_data_works_project" "CreateProject" {
  description      = "新版数据开发exampleterraform接入的工作空间"
  project_name     = "tf_example_datastudio5_4"
  pai_task_enabled = false
  display_name     = "tf_example_datastudio5"
}


resource "alicloud_data_works_function" "default" {
  function_name = "OpenAPIFunc"
  project_id    = alicloud_data_works_project.CreateProject.id
  spec = jsonencode({
    "version" : "1.1.0",
    "kind" : "Function",
    "spec" : {
      "functions" : [
        {
          "name" : "OpenAPIFunc",
          "id" : "5806679648885952157",
          "script" : {
            "content" : "{  \"uuid\": \"5806679648885952157\",  \"name\": \"OpenAPIFunc\"}",
            "path" : "莫泣/OpenAPIexample/函数example/OpenAPIFunc",
            "runtime" : {
              "command" : "ODPS_FUNCTION"
            }
          }
        }
      ]
    }
  })
}
```

## Argument Reference

The following arguments are supported:
* `function_name` - (Optional) Resource Name
* `project_id` - (Required, ForceNew) The id of the workspace to which the udf function belongs.
* `spec` - (Required, JsonString) The Spec definition of a function

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<project_id>:<function_id>`.
* `create_time` - Udf function creation timestamp.
* `function_id` - Unique identifier of the udf function.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Function.
* `delete` - (Defaults to 5 mins) Used when delete the Function.
* `update` - (Defaults to 5 mins) Used when update the Function.

## Import

Data Works Function can be imported using the id, e.g.

```shell
$ terraform import alicloud_data_works_function.example <project_id>:<function_id>
```