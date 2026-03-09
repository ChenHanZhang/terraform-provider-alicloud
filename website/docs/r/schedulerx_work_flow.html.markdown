---
subcategory: "Schedulerx"
layout: "alicloud"
page_title: "Alicloud: alicloud_schedulerx_work_flow"
description: |-
  Provides a Alicloud Schedulerx Work Flow resource.
---

# alicloud_schedulerx_work_flow

Provides a Schedulerx Work Flow resource.



For information about Schedulerx Work Flow and how to use it, see [What is Work Flow](https://next.api.alibabacloud.com/document/schedulerx2/2019-04-30/CreateWorkflow).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_schedulerx_namespace" "CreateNameSpace" {
  namespace_name = "example-namespace-pop-autoexample"
  description    = "由workflow 资源example用例前置步骤创建"
}

resource "alicloud_schedulerx_app_group" "CreateAppGroup" {
  description    = "由workflow 资源example用例前置步骤创建"
  enable_log     = false
  namespace      = alicloud_schedulerx_namespace.CreateNameSpace.id
  group_id       = "example-appgroup-pop-autoexample"
  app_name       = "example-appgroup-pop-autoexample"
  app_version    = "2"
  namespace_name = alicloud_schedulerx_namespace.CreateNameSpace.namespace_name
  app_type       = "2"
  max_jobs       = "100"
}


resource "alicloud_schedulerx_work_flow" "default" {
  timezone         = "GTM+7"
  description      = "workflow资源用例自动生成的任务"
  workflow_name    = "example-workflow-pop-resource-autoexample"
  max_concurrency  = "10"
  time_expression  = "0 0 18 1 */1 ?"
  namespace        = alicloud_schedulerx_namespace.CreateNameSpace.id
  group_id         = alicloud_schedulerx_app_group.CreateAppGroup.group_id
  time_type        = "1"
  status           = "Disable"
  namespace_source = "schedulerx"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Workflow description
* `group_id` - (Required, ForceNew) The application Group ID, which is obtained from the application management page of the console.
* `max_concurrency` - (Optional, ForceNew, Int) The maximum number of workflow instances running at the same time. The default value is 1, that is, the last trigger is not completed, and the next trigger will not be performed even when the running time is reached.
* `namespace` - (Required, ForceNew) The namespace ID, which is obtained from the namespace page in the console.
* `namespace_source` - (Optional) Special third parties are required.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `status` - (Optional, Computed) Workflow Status
* `time_expression` - (Optional) Time expression, which is set based on the selected time type.
  - cron: Fill in standard cron expressions to support online verification.
  - api: No time expression.
* `time_type` - (Required, Int) Time type. Currently, the following time types are supported:
  - cron:1
  - api:100
* `timezone` - (Optional) Time Zone

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `workflow_name` - (Required) Workflow Name

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<namespace>:<group_id>:<work_flow_id>`.
* `work_flow_id` - workflow id.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Work Flow.
* `delete` - (Defaults to 5 mins) Used when delete the Work Flow.
* `update` - (Defaults to 5 mins) Used when update the Work Flow.

## Import

Schedulerx Work Flow can be imported using the id, e.g.

```shell
$ terraform import alicloud_schedulerx_work_flow.example <namespace>:<group_id>:<work_flow_id>
```