---
subcategory: "Resource Manager"
layout: "alicloud"
page_title: "Alicloud: alicloud_resource_manager_resource_share"
description: |-
  Provides a Alicloud Resource Manager Resource Share resource.
---

# alicloud_resource_manager_resource_share

Provides a Resource Manager Resource Share resource.

RS Resource Sharing.

For information about Resource Manager Resource Share and how to use it, see [What is Resource Share](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).

-> **NOTE:** Available since v1.111.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_resource_manager_resource_share" "example" {
  resource_share_name = var.name
}
```

## Argument Reference

The following arguments are supported:
* `allow_external_targets` - (Optional, Available since v1.261.0) Specifies whether sharing with accounts outside the resource directory is allowed. Valid values:
  - false (default): Sharing is allowed only within the resource directory.
  - true: Sharing with any account is allowed.
* `permission_names` - (Optional, List, Available since v1.261.0) The names of sharing permissions. If left empty, the system automatically binds the default permissions associated with the resource type. For more information, see [Permission Library](https://help.aliyun.com/document_detail/465474.html).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_arns` - (Optional, List, Available since v1.268.0) A list of ARNs of shared resources.
Valid values for N: 1 to 5. You can add up to five shared resources at a time.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed, Available since v1.261.0) The resource group ID.
* `resource_properties` - (Optional, List, Available since v1.274.0) A list of resource properties. See [`resource_properties`](#resource_properties) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_share_name` - (Required) The name of the new resource share.
Length: 1 to 50 characters.
Format: English letters, digits, Chinese characters, periods (.), underscores (_), or hyphens (-) are allowed.
* `resources` - (Optional, List, Available since v1.261.0) The list of shared resources. See [`resources`](#resources) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `tags` - (Optional, Map, Available since v1.261.0) A list of tags. You can attach up to 20 tags.
* `targets` - (Optional, List) Resource consumers.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


### `resource_properties`

The resource_properties supports the following:
* `property` - (Optional, Available since v1.274.0) The property of the resource.
* `resource_arn` - (Optional, Available since v1.274.0) The ARN of the resource.

### `resources`

The resources supports the following:
* `resource_id` - (Optional, Available since v1.261.0) The ID of the shared resource.

Valid values for N: 1 to 5. You can add up to five shared resources at a time.

-> **NOTE:**  `Resources.N.ResourceId` and `Resources.N.ResourceType` must be specified together as a pair.

* `resource_type` - (Optional) The type of the shared resource.

Valid values for N: 1 to 5. You can add up to five shared resources at a time.

For information about supported resource types, see [Cloud services that support resource sharing](https://help.aliyun.com/document_detail/450526.html).

-> **NOTE:**  `Resources.N.ResourceId` and `Resources.N.ResourceType` must be specified together as a pair.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The time when the resource share was created.
* `resource_share_owner` - The owner of the resource share.
* `status` - The status of the resource share.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Resource Share.
* `delete` - (Defaults to 10 mins) Used when delete the Resource Share.
* `update` - (Defaults to 5 mins) Used when update the Resource Share.

## Import

Resource Manager Resource Share can be imported using the id, e.g.

```shell
$ terraform import alicloud_resource_manager_resource_share.example <resource_share_id>
```