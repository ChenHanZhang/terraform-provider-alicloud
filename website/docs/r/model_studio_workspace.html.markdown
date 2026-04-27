---
subcategory: "Model Studio"
layout: "alicloud"
page_title: "Alicloud: alicloud_model_studio_workspace"
description: |-
  Provides a Alicloud Model Studio Workspace resource.
---

# alicloud_model_studio_workspace

Provides a Model Studio Workspace resource.



For information about Model Studio Workspace and how to use it, see [What is Workspace](https://next.api.alibabacloud.com/document/ModelStudio/2026-02-10/CreateWorkspace).

-> **NOTE:** Available since v1.277.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_model_studio_workspace" "default" {
  service_site   = "global"
  workspace_name = "Resource_Workspace"
}
```

### Deleting `alicloud_model_studio_workspace` or removing it from your configuration

Terraform cannot destroy resource `alicloud_model_studio_workspace`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `service_site` - (Optional, ForceNew) The service deployment region. For more information, [see the documentation](https://www.alibabacloud.com/help/zh/model-studio/regions/).
* `workspace_name` - (Required) The workspace name.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time.
* `region_id` - The region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Workspace.
* `update` - (Defaults to 5 mins) Used when update the Workspace.

## Import

Model Studio Workspace can be imported using the id, e.g.

```shell
$ terraform import alicloud_model_studio_workspace.example <workspace_id>
```