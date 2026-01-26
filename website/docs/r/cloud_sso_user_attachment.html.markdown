---
subcategory: "Cloud SSO"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_sso_user_attachment"
description: |-
  Provides a Alicloud Cloud Sso User Attachment resource.
---

# alicloud_cloud_sso_user_attachment

Provides a Cloud Sso User Attachment resource.

Add users to user groups.

For information about Cloud Sso User Attachment and how to use it, see [What is User Attachment](https://www.alibabacloud.com/help/en/cloudsso/latest/api-cloudsso-2021-05-15-addusertogroup).

-> **NOTE:** Available since v1.141.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}

data "alicloud_cloud_sso_directories" "default" {
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

resource "alicloud_cloud_sso_directory" "default" {
  count          = length(data.alicloud_cloud_sso_directories.default.ids) > 0 ? 0 : 1
  directory_name = var.name
}

resource "alicloud_cloud_sso_user" "default" {
  directory_id = local.directory_id
  user_name    = "${var.name}-${random_integer.default.result}"
}

resource "alicloud_cloud_sso_group" "default" {
  directory_id = local.directory_id
  group_name   = var.name
  description  = var.name
}

resource "alicloud_cloud_sso_user_attachment" "default" {
  directory_id = local.directory_id
  user_id      = alicloud_cloud_sso_user.default.user_id
  group_id     = alicloud_cloud_sso_group.default.group_id
}

locals {
  directory_id = length(data.alicloud_cloud_sso_directories.default.ids) > 0 ? data.alicloud_cloud_sso_directories.default.ids[0] : concat(alicloud_cloud_sso_directory.default.*.id, [""])[0]
}
```

## Argument Reference

The following arguments are supported:
* `directory_id` - (Required, ForceNew) Directory ID
* `group_id` - (Required, ForceNew) The resource attribute field that represents the resource name.
* `user_id` - (Required, ForceNew) User ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<directory_id>:<group_id>:<user_id>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the User Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the User Attachment.

## Import

Cloud Sso User Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_sso_user_attachment.example <directory_id>:<group_id>:<user_id>
```