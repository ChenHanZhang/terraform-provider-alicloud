---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_deployment_set"
description: |-
  Provides a Alicloud ECS Deployment Set resource.
---

# alicloud_ecs_deployment_set

Provides a ECS Deployment Set resource.

Instance deployment set.

For information about ECS Deployment Set and how to use it, see [What is Deployment Set](https://www.alibabacloud.com/help/en/doc-detail/91269.htm).

-> **NOTE:** Available since v1.140.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_ecs_deployment_set" "default" {
  strategy            = "Availability"
  deployment_set_name = var.name
  description         = var.name
}
```

## Argument Reference

The following arguments are supported:
* `deployment_set_name` - (Optional) The name of the deployment set. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
* `description` - (Optional) The description of the deployment set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
* `domain` - (Optional, ForceNew, Computed) >This parameter is deprecated.
* `granularity` - (Optional, ForceNew, Computed) >This parameter is deprecated.
* `group_count` - (Optional, ForceNew, Computed, Int, Available since v1.287.0) The number of groups for the high availability strategy of the deployment set group. Valid values: 1 to 7.

Default value: 3.

-> **NOTE:**  This parameter takes effect only when `Strategy=AvailabilityGroup`.

* `on_unable_to_redeploy_failed_instance` - (Optional) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `strategy` - (Optional, ForceNew, Computed) The deployment strategy. Valid values:
  - Availability: High Availability Strategy
  - AvailabilityGroup: High Availability Strategy for Deployment Set Group
  - LowLatency: Network Low Latency Strategy

Default value: Availability.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the deployment set.
* `instance_amount` - The number of instances in the deployment set.
* `instance_ids` - The instance ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Deployment Set.
* `delete` - (Defaults to 5 mins) Used when delete the Deployment Set.
* `update` - (Defaults to 5 mins) Used when update the Deployment Set.

## Import

ECS Deployment Set can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_deployment_set.example <deployment_set_id>
```