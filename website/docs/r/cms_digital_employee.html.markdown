---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_digital_employee"
description: |-
  Provides a Alicloud Cms Digital Employee resource.
---

# alicloud_cms_digital_employee

Provides a Cms Digital Employee resource.



For information about Cms Digital Employee and how to use it, see [What is Digital Employee](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateDigitalEmployee).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `default_rule` - (Optional) Default rule for digital employees.  
* `description` - (Optional) Description of the digital employee.  
* `display_name` - (Optional) The display name of the digital employee.  
* `knowledges` - (Optional, Set) Knowledge base for digital employees.   See [`knowledges`](#knowledges) below.
* `name` - (Required, ForceNew) The name of the digital employee.  
* `resource_group_id` - (Optional, Computed) Resource attribute field representing the resource group.  
* `role_arn` - (Required) The ARN of the RAM role.  
* `tags` - (Optional, Map) The resource attribute field that represents resource tags.  

### `knowledges`

The knowledges supports the following:
* `bailian` - (Optional, List) List of Bailian knowledge bases.   See [`bailian`](#knowledges-bailian) below.

### `knowledges-bailian`

The knowledges-bailian supports the following:
* `attributes` - (Optional) Knowledge base attributes  
* `index_id` - (Optional) Bailian index ID.  
* `region` - (Optional) Knowledge base region  
* `workspace_id` - (Optional) Bailian workspace ID.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `region_id` - The resource attribute field that represents the region.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Digital Employee.
* `delete` - (Defaults to 5 mins) Used when delete the Digital Employee.
* `update` - (Defaults to 5 mins) Used when update the Digital Employee.

## Import

Cms Digital Employee can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_digital_employee.example <name>
```