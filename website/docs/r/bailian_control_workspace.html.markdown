---
subcategory: "Bailian Control"
layout: "alicloud"
page_title: "Alicloud: alicloud_bailian_control_workspace"
description: |-
  Provides a Alicloud Bailian Control Workspace resource.
---

# alicloud_bailian_control_workspace

Provides a Bailian Control Workspace resource.

An isolated environment for data and applications.

For information about Bailian Control Workspace and how to use it, see [What is Workspace](https://next.api.alibabacloud.com/document/BailianControl/2024-08-16/ListTagResources).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_bailian_control_workspace` or removing it from your configuration

Terraform cannot destroy resource `alicloud_bailian_control_workspace`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `tags` - (Optional, Map) A resource attribute field representing resource tags.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `update` - (Defaults to 5 mins) Used when update the Workspace.

## Import

Bailian Control Workspace can be imported using the id, e.g.

```shell
$ terraform import alicloud_bailian_control_workspace.example <workspace_id>
```