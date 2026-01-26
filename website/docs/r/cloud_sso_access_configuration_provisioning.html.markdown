---
subcategory: "Cloud SSO"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_sso_access_configuration_provisioning"
description: |-
  Provides a Alicloud Cloud Sso Access Configuration Provisioning resource.
---

# alicloud_cloud_sso_access_configuration_provisioning

Provides a Cloud Sso Access Configuration Provisioning resource.

deploy a special access configuration to the target account.

For information about Cloud Sso Access Configuration Provisioning and how to use it, see [What is Access Configuration Provisioning](https://www.alibabacloud.com/help/en/cloudsso/latest/api-cloudsso-2021-05-15-addpermissionpolicytoaccessconfiguration).

-> **NOTE:** Available since v1.148.0.

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
data "alicloud_resource_manager_resource_directories" "default" {}

resource "alicloud_cloud_sso_directory" "default" {
  count          = length(data.alicloud_cloud_sso_directories.default.ids) > 0 ? 0 : 1
  directory_name = var.name
}

locals {
  directory_id = length(data.alicloud_cloud_sso_directories.default.ids) > 0 ? data.alicloud_cloud_sso_directories.default.ids[0] : concat(alicloud_cloud_sso_directory.default.*.id, [""])[0]
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

resource "alicloud_cloud_sso_user" "default" {
  directory_id = local.directory_id
  user_name    = "${var.name}-${random_integer.default.result}"
}

resource "alicloud_cloud_sso_access_configuration" "default" {
  access_configuration_name = "${var.name}-${random_integer.default.result}"
  directory_id              = local.directory_id
}

resource "alicloud_cloud_sso_access_configuration_provisioning" "default" {
  directory_id            = local.directory_id
  access_configuration_id = alicloud_cloud_sso_access_configuration.default.access_configuration_id
  target_type             = "RD-Account"
  target_id               = data.alicloud_resource_manager_resource_directories.default.directories.0.master_account_id
}
```

## Argument Reference

The following arguments are supported:
* `access_configuration_id` - (Required, ForceNew) Access configuration ID
* `directory_id` - (Required, ForceNew) Directory ID
* `target_id` - (Required, ForceNew) The ID of the target to create the resource range.
* `target_type` - (Required, ForceNew) The type of the resource range target to be accessed. Only a single RD primary account or member account can be specified in the first phase.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<directory_id>:<access_configuration_id>:<target_type>:<target_id>`.
* `create_time` - The creation time of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Access Configuration Provisioning.
* `delete` - (Defaults to 5 mins) Used when delete the Access Configuration Provisioning.

## Import

Cloud Sso Access Configuration Provisioning can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_sso_access_configuration_provisioning.example <directory_id>:<access_configuration_id>:<target_type>:<target_id>
```