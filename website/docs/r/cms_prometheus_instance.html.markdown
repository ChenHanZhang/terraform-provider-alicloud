---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_prometheus_instance"
description: |-
  Provides a Alicloud Cms Prometheus Instance resource.
---

# alicloud_cms_prometheus_instance

Provides a Cms Prometheus Instance resource.



For information about Cms Prometheus Instance and how to use it, see [What is Prometheus Instance](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreatePrometheusInstance).

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


resource "alicloud_cms_prometheus_instance" "default" {
  status                   = "Running"
  archive_duration         = "60"
  prometheus_instance_name = "model-example-hz-2"
  auth_free_read_policy    = "1.1.1.1"
  auth_free_write_policy   = "2.2.2.2"
  storage_duration         = "30"
  enable_auth_free_read    = true
  enable_auth_free_write   = true
  workspace                = "prometheus-1511928242963727"
}
```

## Argument Reference

The following arguments are supported:
* `archive_duration` - (Optional, Int) The number of days after the storage expires that the automatic archive is saved (optional values: 60, 90, 180, 365). 0 means save without archiving.
* `auth_free_read_policy` - (Optional) Read password-free address whitelist policy.
* `auth_free_write_policy` - (Optional) Write a secret-free address whitelist policy.
* `enable_auth_free_read` - (Optional, Computed) Whether to turn on the read-free password.
* `enable_auth_free_write` - (Optional, Computed) Whether to turn on the write-free password.
* `enable_auth_token` - (Optional, Computed) Whether to enable access Token authentication
* `prometheus_instance_name` - (Required) Prometheus instance name
* `status` - (Optional, Computed) Back-end data storage status
* `storage_duration` - (Optional, Int) Data store duration (days).
* `tags` - (Optional, Map) The tag of the resource
* `workspace` - (Required, ForceNew) Work space

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Instance creation time, using UTC +0 time, in the format of yyyy-MM-ddTHH:mmZ.
* `payment_type` - Payment Type.
* `region_id` - The region ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Prometheus Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Prometheus Instance.
* `update` - (Defaults to 5 mins) Used when update the Prometheus Instance.

## Import

Cms Prometheus Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_prometheus_instance.example <prometheus_instance_id>
```