---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_diagnostic_metric_set"
description: |-
  Provides a Alicloud ECS Diagnostic Metric Set resource.
---

# alicloud_ecs_diagnostic_metric_set

Provides a ECS Diagnostic Metric Set resource.

Diagnostic Metric Set.

For information about ECS Diagnostic Metric Set and how to use it, see [What is Diagnostic Metric Set](https://next.api.alibabacloud.com/document/Ecs/2014-05-26/CreateDiagnosticMetricSet).

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


resource "alicloud_ecs_diagnostic_metric_set" "default" {
  description     = "example pop metric set"
  metric_ids      = ["Instance.NetworkSessionError"]
  resource_type   = "instance"
  metric_set_name = "pop example metric set"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) The description of the diagnostic indicator set.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `metric_ids` - (Required, ForceNew, List) List of diagnostic indicators. It supports up to 100.
* `metric_set_name` - (Optional) The name of the diagnostic metric collection.
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `resource_type` - (Required, ForceNew) The diagnostic resource type.
Default value: instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Diagnostic Metric Set.
* `delete` - (Defaults to 5 mins) Used when delete the Diagnostic Metric Set.
* `update` - (Defaults to 5 mins) Used when update the Diagnostic Metric Set.

## Import

ECS Diagnostic Metric Set can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_diagnostic_metric_set.example <metric_set_id>
```