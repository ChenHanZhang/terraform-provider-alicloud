---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_security_center_group"
description: |-
  Provides a Alicloud Threat Detection Group resource.
---

# alicloud_security_center_group

Provides a Threat Detection Group resource.

Asset grouping in Security Center.

For information about Threat Detection Group and how to use it, see [What is Group](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createorupdateassetgroup).

-> **NOTE:** Available since v1.133.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf_example"
}
resource "alicloud_security_center_group" "example" {
  group_name = var.name
}
```

## Argument Reference

The following arguments are supported:
* `group_id` - (Optional, ForceNew, Computed, Int) The ID of the server group for which you want to add to or remove servers.

-> **NOTE:**   To modify the mapping between an asset and an asset group, you must provide the ID of the asset group. You can call the [DescribeAllGroups](~~DescribeAllGroups~~) to query the IDs of asset groups. If you do not configure this parameter when you call this operation, an asset group is created.

* `group_name` - (Optional) The name of the server group that you want to create or the server group for which you want to add or remove a server.

-> **NOTE:**   To modify the mapping between a server and a server group, you must provide the name of the server group. You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the names of server groups. If you do not configure GroupID when you call this operation, a server group is created. In this case, you must configure GroupName.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `group_flag` - The type of the server group.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Group.
* `delete` - (Defaults to 5 mins) Used when delete the Group.
* `update` - (Defaults to 5 mins) Used when update the Group.

## Import

Threat Detection Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_security_center_group.example <group_id>
```