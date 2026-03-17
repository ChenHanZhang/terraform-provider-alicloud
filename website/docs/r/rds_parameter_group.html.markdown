---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_parameter_group"
description: |-
  Provides a Alicloud RDS Parameter Group resource.
---

# alicloud_rds_parameter_group

Provides a RDS Parameter Group resource.

Used for batch management of database parameters  .

For information about RDS Parameter Group and how to use it, see [What is Parameter Group](https://next.api.alibabacloud.com/document/Rds/2014-08-15/CreateParameterGroup).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf_example"
}

resource "alicloud_rds_parameter_group" "default" {
  engine         = "mysql"
  engine_version = "5.7"
  param_detail {
    param_name  = "back_log"
    param_value = "4000"
  }
  param_detail {
    param_name  = "wait_timeout"
    param_value = "86460"
  }
  parameter_group_desc = var.name
  parameter_group_name = var.name
}
```

## Argument Reference

The following arguments are supported:
* `engine` - (Required, ForceNew) Database engine. Valid values:  
  - `mysql`  
  - `PostgreSQL`.
* `engine_version` - (Required, ForceNew) The database version. Valid values:
  - MySQL:
    * **5.6**
    * **5.7**
    * **8.0**
  - PostgreSQL:
    * **10.0**
    * **11.0**
    * **12.0**
    * **13.0**
    * **14.0**
    * **15.0**.
* `modify_mode` - (Optional) The modification mode of the parameter template. Valid values:  
* `Collectivity` (default): Add or update.  

-> **NOTE:**  Adds the parameters specified in the `Parameters` property to the current parameter template, or updates existing parameters in the template. Other parameters in the current template remain unaffected.  

* `Individual`: Overwrite.  

-> **NOTE:**  Replaces all parameters in the current parameter template with the parameters specified in the `Parameters` property.  


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `parameter_detail` - (Required, List) A list of parameters. See [`parameter_detail`](#parameter_detail) below.
* `parameter_group_desc` - (Optional) The description of the parameter template. It can be 0 to 200 characters in length.
* `parameter_group_name` - (Required) The name of the parameter template.
* It must start with a letter and can contain letters, digits, periods (.), or underscores (_).
* It must be 8 to 64 characters in length.

-> **NOTE:** If you do not specify this parameter, the original parameter template name is retained.

* `resource_group_id` - (Optional, Computed) The resource group ID. You can obtain it by calling the DescribeDBInstanceAttribute operation.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.


### `parameter_detail`

The parameter_detail supports the following:
* `param_name` - (Optional) The parameter name.
* `param_value` - (Optional) The parameter value.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Parameter Group.
* `delete` - (Defaults to 5 mins) Used when delete the Parameter Group.
* `update` - (Defaults to 5 mins) Used when update the Parameter Group.

## Import

RDS Parameter Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_parameter_group.example <id>
```