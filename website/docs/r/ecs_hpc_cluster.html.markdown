---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_hpc_cluster"
description: |-
  Provides a Alicloud ECS Hpc Cluster resource.
---

# alicloud_ecs_hpc_cluster

Provides a ECS Hpc Cluster resource.

High Performance Computing.

For information about ECS Hpc Cluster and how to use it, see [What is Hpc Cluster](https://www.alibabacloud.com/help/en/doc-detail/109138.htm).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_ecs_hpc_cluster" "example" {
  name        = "tf-testAcc"
  description = "For Terraform Test"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Description
* `hpc_cluster_name` - (Required) Name

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Hpc Cluster.
* `delete` - (Defaults to 5 mins) Used when delete the Hpc Cluster.
* `update` - (Defaults to 5 mins) Used when update the Hpc Cluster.

## Import

ECS Hpc Cluster can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_hpc_cluster.example <hpc_cluster_id>
```