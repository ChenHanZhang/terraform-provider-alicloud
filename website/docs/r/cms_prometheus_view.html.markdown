---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_prometheus_view"
description: |-
  Provides a Alicloud Cms Prometheus View resource.
---

# alicloud_cms_prometheus_view

Provides a Cms Prometheus View resource.



For information about Cms Prometheus View and how to use it, see [What is Prometheus View](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreatePrometheusView).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

data "alicloud_resource_manager_resource_groups" "default" {}


resource "alicloud_cms_prometheus_view" "default" {
  status               = "Running"
  prometheus_view_name = "view-ctb-1"
  version              = "V2"
  prometheus_instances {
    prometheus_instance_id = "c1e94e934622f4e42a7f88e1744e858b1"
    region_id              = "cn-hangzhou"
    user_id                = "1707387247477962"
  }
  workspace = "prometheus-1511928242963727"
}
```

## Argument Reference

The following arguments are supported:
* `auth_free_read_policy` - (Optional) Read password-free address whitelist policy.
* `enable_auth_free_read` - (Optional, Computed) Whether to turn on the read-free password.
* `enable_auth_token` - (Optional, Computed) Whether to enable authToken.
* `prometheus_instances` - (Required, List) List of prometheus instances contained in the view See [`prometheus_instances`](#prometheus_instances) below.
* `prometheus_view_name` - (Required) The name of the resource
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `status` - (Optional, Computed) The status of the resource
* `tags` - (Optional, Map) The tag of the resource
* `version` - (Required, ForceNew) V1: old version (global-view)
v2: new version (prom-view)
* `workspace` - (Required, ForceNew) Environment-owned workspace

### `prometheus_instances`

The prometheus_instances supports the following:
* `prometheus_instance_id` - (Required) Prometheus instance Id
* `region_id` - (Required) The region Id of the prometheus instance.
* `user_id` - (Required) The user ID of the prometheus instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Instance creation time, using UTC +0 time, in the format of yyyy-MM-ddTHH:mmZ.
* `region_id` - Instance Region.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Prometheus View.
* `delete` - (Defaults to 5 mins) Used when delete the Prometheus View.
* `update` - (Defaults to 5 mins) Used when update the Prometheus View.

## Import

Cms Prometheus View can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_prometheus_view.example <prometheus_view_id>
```