---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_addon_release"
description: |-
  Provides a Alicloud Cms Addon Release resource.
---

# alicloud_cms_addon_release

Provides a Cms Addon Release resource.

Release package of observability addon.

For information about Cms Addon Release and how to use it, see [What is Addon Release](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateAddonRelease).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_cms_addon_release" "default" {
  addon_version         = "0.0.2"
  aliyun_lang           = "zh"
  addon_name            = "cs-gpu"
  integration_policy_id = "env-79520003404088b43c8aabaf4bb8"
  workspace             = "sls-mall"
  addon_release_name    = "gpu-inte-example-another-four"
  dry_run               = true
  values                = jsonencode({ "install" : { "mode" : "auto-install", "listenPort" : "9400" }, "discoverMode" : "instances", "discover" : { "instances" : "worker-k8s-for-cs-c126d87c76218487e83ab322017f11b44" }, "scrapeInterval" : "15", "enableSecuritecs-nodeyGroupInjection" : "true", "metricTags" : "" })
  env_type              = "CS"
}
```

## Argument Reference

The following arguments are supported:
* `addon_name` - (Required, ForceNew) Addon Name
* `addon_release_name` - (Optional, ForceNew, Computed) The name of the resource
* `addon_version` - (Required) Version number of Addon. Addon information can be obtained through ListAddons.
* `aliyun_lang` - (Optional, ForceNew) The installed locale.
* `dry_run` - (Optional) Test run, default to false.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `env_type` - (Optional, ForceNew) Environment type. If the Policy type is CS and ECS, use them accordingly, and all other types are uniformly Cloud.
* `force` - (Optional) Whether to forcibly delete, default to false.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `integration_policy_id` - (Required, ForceNew) Environment id
* `parent_addon_release_id` - (Optional, ForceNew) Parent level AddonRelease ID.
* `values` - (Optional, JsonString) Configuration information for installing Addon. Obtain the configuration template from ListAddonSchema, for example, {"host":"mysql-service.default","port":3306,"username":"root","password":"roots"}

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `workspace` - (Optional, ForceNew) The workspace name for installing component resources.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<integration_policy_id>:<addon_release_name>`.
* `create_time` - Creation time.
* `region_id` - region id.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Addon Release.
* `delete` - (Defaults to 5 mins) Used when delete the Addon Release.
* `update` - (Defaults to 5 mins) Used when update the Addon Release.

## Import

Cms Addon Release can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_addon_release.example <integration_policy_id>:<addon_release_name>
```