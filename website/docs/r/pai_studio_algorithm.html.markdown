---
subcategory: "Pai Studio"
layout: "alicloud"
page_title: "Alicloud: alicloud_pai_studio_algorithm"
description: |-
  Provides a Alicloud Pai Studio Algorithm resource.
---

# alicloud_pai_studio_algorithm

Provides a Pai Studio Algorithm resource.



For information about Pai Studio Algorithm and how to use it, see [What is Algorithm](https://next.api.alibabacloud.com/document/PaiStudio/2022-01-12/CreateAlgorithm).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_pai_workspace_workspace" "workspace" {
  workspace_name = "example_work_space_1770122482"
  env_types      = ["prod"]
  description    = "Algorithm资源example用例使用"
  display_name   = "algorithm_example_case"
}


resource "alicloud_pai_studio_algorithm" "default" {
  workspace_id          = alicloud_pai_workspace_workspace.workspace.id
  display_name          = "example"
  algorithm_description = "算法example用例使用"
  algorithm_name        = "example_1770122483"
}
```

## Argument Reference

The following arguments are supported:
* `algorithm_description` - (Optional) Algorithm description.
* `algorithm_name` - (Optional, ForceNew) The algorithm name.
* `display_name` - (Optional) Algorithm display name.
* `workspace_id` - (Optional, ForceNew) The ID of the workspace.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Algorithm creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Algorithm.
* `delete` - (Defaults to 5 mins) Used when delete the Algorithm.
* `update` - (Defaults to 5 mins) Used when update the Algorithm.

## Import

Pai Studio Algorithm can be imported using the id, e.g.

```shell
$ terraform import alicloud_pai_studio_algorithm.example <algorithm_id>
```