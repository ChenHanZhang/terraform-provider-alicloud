---
subcategory: "P A I Workspace"
layout: "alicloud"
page_title: "Alicloud: alicloud_p_a_i_workspace_model"
description: |-
  Provides a Alicloud P A I Workspace Model resource.
---

# alicloud_p_a_i_workspace_model

Provides a P A I Workspace Model resource. 

For information about P A I Workspace Model and how to use it, see [What is Model](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.212.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_p_a_i_workspace_workspace" "defaultDI9fsL" {
  description    = 368
  workspace_name = var.name

  env_types    = ["prod"]
  display_name = test_pop_925
}


resource "alicloud_p_a_i_workspace_model" "default" {
  model_name = var.name

  workspace_id      = alicloud_p_a_i_workspace_workspace.defaultDI9fsL.id
  origin            = "Civitai"
  task              = "text-to-image-synthesis"
  accessibility     = "PRIVATE"
  model_type        = "Checkpoint"
  order_number      = "1"
  model_description = "ModelDescription."
  model_doc         = "https://eas-***.oss-cn-hangzhou.aliyuncs.com/s**.safetensors"
  domain            = "aigc"
  labels {
    key   = "base_model"
    value = "SD 1.5"
  }
  extra_info {
  }
}
```

## Argument Reference

The following arguments are supported:
* `accessibility` - (Optional) Workspace visibility, possible values are:
  - PRIVATE: In this workspace, it is only visible to you and the administrator.
  - PUBLIC: In this workspace, it is visible to everyone.
* `domain` - (Optional) Field. Describe the domain in which the model solves the problem. For example: nlp (natural language processing), cv (computer vision), etc.
* `extra_info` - (Optional, Map) Other Information.
* `labels` - (Optional, ForceNew) A list of tags. See [`labels`](#labels) below.
* `model_description` - (Optional) Model description.
* `model_doc` - (Optional) Model documentation.
* `model_name` - (Required) The model name.
* `model_type` - (Optional) Model Type.
* `order_number` - (Optional) Model serial number. Can be used for custom sorting.
* `origin` - (Optional) Model source. Describes the community or organization to which the source model belongs, for example, ModelScope and HuggingFace.
* `task` - (Optional) Tasks. Describe the specific problems that the model solves, such as text-classification.
* `workspace_id` - (Optional, ForceNew) The ID of the workspace.

### `labels`

The labels supports the following:
* `key` - (Optional, ForceNew) label key.
* `value` - (Optional, ForceNew) label value.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Model.
* `delete` - (Defaults to 5 mins) Used when delete the Model.
* `update` - (Defaults to 5 mins) Used when update the Model.

## Import

P A I Workspace Model can be imported using the id, e.g.

```shell
$ terraform import alicloud_p_a_i_workspace_model.example <id>
```