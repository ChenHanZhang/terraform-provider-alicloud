---
subcategory: "S T A R Ops"
layout: "alicloud"
page_title: "Alicloud: alicloud_s_t_a_r_ops_digital_employee"
description: |-
  Provides a Alicloud S T A R Ops Digital Employee resource.
---

# alicloud_s_t_a_r_ops_digital_employee

Provides a S T A R Ops Digital Employee resource.



For information about S T A R Ops Digital Employee and how to use it, see [What is Digital Employee](https://next.api.alibabacloud.com/document/STAROps/2026-04-28/CreateDigitalEmployee).

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `attributes` - (Optional, Map) The attributes of the digital employee.
* `default_rule` - (Optional) Default rule for digital employees.  
* `description` - (Optional) Description of the digital employee.  
* `display_name` - (Optional) Digital employee display name.
* `knowledges` - (Optional, Set) Knowledge base for digital employees.   See [`knowledges`](#knowledges) below.
* `name` - (Required, ForceNew) Digital employee name.
* `resource_group_id` - (Optional, Computed) Resource attribute field representing the resource group.  
* `role_arn` - (Required) RAM role ARN.
* `tags` - (Optional, Map) Resource property field representing resource tags.

### `knowledges`

The knowledges supports the following:
* `bailian` - (Optional, List) List of Bailian knowledge bases.   See [`bailian`](#knowledges-bailian) below.

### `knowledges-bailian`

The knowledges-bailian supports the following:
* `attributes` - (Optional) Knowledge base attributes.
* `index_id` - (Optional) Bailian index ID.  
* `region` - (Optional) Knowledge base region.
* `workspace_id` - (Optional) Bailian workspace ID.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `region_id` - Resource property field representing the region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Digital Employee.
* `delete` - (Defaults to 5 mins) Used when delete the Digital Employee.
* `update` - (Defaults to 5 mins) Used when update the Digital Employee.

## Import

S T A R Ops Digital Employee can be imported using the id, e.g.

```shell
$ terraform import alicloud_s_t_a_r_ops_digital_employee.example <name>
```