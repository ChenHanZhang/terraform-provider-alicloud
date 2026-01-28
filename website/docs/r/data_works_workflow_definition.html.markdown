---
subcategory: "Data Works"
layout: "alicloud"
page_title: "Alicloud: alicloud_data_works_workflow_definition"
description: |-
  Provides a Alicloud Data Works Workflow Definition resource.
---

# alicloud_data_works_workflow_definition

Provides a Data Works Workflow Definition resource.

Several nodes are assembled into a directed acyclic graph based on dependencies to complete a series of computations or operational processes with pre-and post-dependencies.

For information about Data Works Workflow Definition and how to use it, see [What is Workflow Definition](https://next.api.alibabacloud.com/document/dataworks-public/2024-05-18/CreateWorkflowDefinition).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}

resource "alicloud_data_works_project" "CreateProject" {
  description      = "新版数据开发exampleterraform接入的工作空间"
  project_name     = "tf_example_datastudio5_2"
  pai_task_enabled = false
  display_name     = "tf_example_datastudio5"
}


resource "alicloud_data_works_workflow_definition" "default" {
  workflow_definition_name = "OpenAPIexample工作流Demo"
  project_id               = alicloud_data_works_project.CreateProject.id
  spec = jsonencode({
    "metadata" : {
      "uuid" : "4634978808809548226"
    },
    "kind" : "CycleWorkflow",
    "version" : "1.1.0",
    "spec" : {
      "name" : "OpenAPIexample工作流Demo",
      "id" : "4634978808809548226",
      "type" : "CycleWorkflow",
      "workflows" : [
        {
          "script" : {
            "path" : "莫泣/OpenAPIexample/工作流example/OpenAPIexample工作流Demo",
            "runtime" : {
              "command" : "WORKFLOW"
            },
            "id" : "6980027813686443486"
          },
          "id" : "4634978808809548226",
          "trigger" : {
            "type" : "Scheduler",
            "id" : "6525678244703541090",
            "cron" : "00 02 00 * * ?",
            "startTime" : "1970-01-01 00:00:00",
            "endTime" : "9999-01-01 00:00:00",
            "timezone" : "Asia/Shanghai",
            "delaySeconds" : 0
          },
          "strategy" : {
            "timeout" : 0,
            "instanceMode" : "T+1",
            "rerunMode" : "Allowed",
            "rerunTimes" : 3,
            "rerunInterval" : 180000,
            "failureStrategy" : "Break"
          },
          "name" : "OpenAPIexample工作流Demo",
          "inputs" : {},
          "outputs" : {
            "nodeOutputs" : [
              {
                "data" : "4634978808809548226",
                "artifactType" : "NodeOutput",
                "refTableName" : "OpenAPIexample工作流Demo",
                "isDefault" : true
              }
            ]
          },
          "nodes" : [],
          "dependencies" : []
        }
      ]
    }
  })
}
```

## Argument Reference

The following arguments are supported:
* `project_id` - (Required, ForceNew) Workspace of the workflow definition
* `spec` - (Required, JsonString) FlowSpec definition of workflow
* `workflow_definition_name` - (Optional) Name of the workflow

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<project_id>:<workflow_definition_id>`.
* `create_time` - Creation time of the workflow.
* `workflow_definition_id` - Unique identification of the workflow definition.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Workflow Definition.
* `delete` - (Defaults to 5 mins) Used when delete the Workflow Definition.
* `update` - (Defaults to 5 mins) Used when update the Workflow Definition.

## Import

Data Works Workflow Definition can be imported using the id, e.g.

```shell
$ terraform import alicloud_data_works_workflow_definition.example <project_id>:<workflow_definition_id>
```