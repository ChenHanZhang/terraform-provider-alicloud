---
subcategory: "Cloud SSO"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_sso_group"
description: |-
  Provides a Alicloud Cloud Sso Group resource.
---

# alicloud_cloud_sso_group

Provides a Cloud Sso Group resource.



For information about Cloud Sso Group and how to use it, see [What is Group](https://www.alibabacloud.com/help/en/cloudsso/latest/api-cloudsso-2021-05-15-creategroup).

-> **NOTE:** Available since v1.138.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}
provider "alicloud" {
  region = "cn-shanghai"
}
data "alicloud_cloud_sso_directories" "default" {}

resource "alicloud_cloud_sso_directory" "default" {
  count          = length(data.alicloud_cloud_sso_directories.default.ids) > 0 ? 0 : 1
  directory_name = var.name
}

locals {
  directory_id = length(data.alicloud_cloud_sso_directories.default.ids) > 0 ? data.alicloud_cloud_sso_directories.default.ids[0] : concat(alicloud_cloud_sso_directory.default.*.id, [""])[0]
}

resource "alicloud_cloud_sso_group" "default" {
  directory_id = local.directory_id
  group_name   = var.name
  description  = var.name
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Description
* `directory_id` - (Required, ForceNew) DirectoryId
* `group_name` - (Required) GroupName

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<directory_id>:<group_id>`.
* `create_time` - CreateTime.
* `group_id` - GroupId.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Group.
* `delete` - (Defaults to 5 mins) Used when delete the Group.
* `update` - (Defaults to 5 mins) Used when update the Group.

## Import

Cloud Sso Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_sso_group.example <directory_id>:<group_id>
```