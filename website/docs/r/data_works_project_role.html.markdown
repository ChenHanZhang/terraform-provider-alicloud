---
subcategory: "Data Works"
layout: "alicloud"
page_title: "Alicloud: alicloud_data_works_project_role"
description: |-
  Provides a Alicloud Data Works Project Role resource.
---

# alicloud_data_works_project_role

Provides a Data Works Project Role resource.



For information about Data Works Project Role and how to use it, see [What is Project Role](https://next.api.alibabacloud.com/document/dataworks-public/2024-05-18/CreateProjectRole).

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `module_permissions` - (Optional, List) The list of DataWorks module permissions. See [`module_permissions`](#module_permissions) below.
* `project_id` - (Required, ForceNew, Int) The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console), navigate to the workspace configuration page, and obtain the workspace ID there.
* `project_role_name` - (Required, ForceNew) The role name in the workspace.

### `module_permissions`

The module_permissions supports the following:
* `module_id` - (Optional, Int) DataWorks module ID:
  - 2: HoloStudio
  - 3: StreamStudio
  - 4: Deployment Management
  - 6: Data Protection Umbrella
  - 7: Data Map
  - 8: Data Service
  - 9: Data Integration
  - 10: Data Modeling (DataBlau DDM)
  - 11: Data Development
  - 12: Data Quality
  - 13: Data Governance
  - 14: Operation Center
  - 15: Resource Optimization
  - 16: Migration Assistant
  - 17: Data Analysis
  - 18: Approval Center
  - 19: Security Center
  - 20: Intelligent Data Modeling
* `permission_type` - (Optional) Permission type:
  - Write (write permission: 1)
  - Read (read permission: 2)
  - NotSet (not authorized: 0)

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<project_id>:<code>`.
* `code` - Workspace role code.
* `module_permissions` - The list of DataWorks module permissions.
  * `module_name` - DataWorks module names:.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Project Role.
* `delete` - (Defaults to 5 mins) Used when delete the Project Role.
* `update` - (Defaults to 5 mins) Used when update the Project Role.

## Import

Data Works Project Role can be imported using the id, e.g.

```shell
$ terraform import alicloud_data_works_project_role.example <project_id>:<code>
```