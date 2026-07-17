---
subcategory: "Resource Manager"
layout: "alicloud"
page_title: "Alicloud: alicloud_resource_manager_resource_group"
description: |-
  Provides a Alicloud Resource Manager Resource Group resource.
---

# alicloud_resource_manager_resource_group

Provides a Resource Manager Resource Group resource.

The method of logically grouping resources in an Alibaba Cloud account.

For information about Resource Manager Resource Group and how to use it, see [What is Resource Group](https://www.alibabacloud.com/help/en/resource-management/developer-reference/api-createresourcegroup).

-> **NOTE:** Available since v1.82.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tfexample"
}

resource "alicloud_resource_manager_resource_group" "example" {
  resource_group_name = var.name
  display_name        = var.name
}
```

## Argument Reference

The following arguments are supported:
* `display_name` - (Required) The display name of the resource group.
* `resource_group_name` - (Required, ForceNew, Available since v1.114.0) The unique identifier of the resource group.
* `tags` - (Optional, Map, Available since v1.220.0) The tag of the resource

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `account_id` - The ID of the Alibaba Cloud account to which the resource group belongs.
* `create_time` - The time when the resource group was created.
* `region_statuses` - The status of the resource group in all regions.
  * `region_id` - The region ID.
  * `status` - The status of the regional resource group.
* `resource_group_id` - The ID of the resource group.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Resource Group.
* `delete` - (Defaults to 5 mins) Used when delete the Resource Group.
* `update` - (Defaults to 5 mins) Used when update the Resource Group.

## Import

Resource Manager Resource Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_resource_manager_resource_group.example <resource_group_id>
```