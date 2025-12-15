---
subcategory: "Actiontrail"
layout: "alicloud"
page_title: "Alicloud: alicloud_actiontrail_history_delivery_job"
description: |-
  Provides a Alicloud Action Trail History Delivery Job resource.
---

# alicloud_actiontrail_history_delivery_job

Provides a Action Trail History Delivery Job resource.

Delivery History Tasks.

For information about Action Trail History Delivery Job and how to use it, see [What is History Delivery Job](https://www.alibabacloud.com/help/en/actiontrail/latest/api-actiontrail-2020-07-06-createdeliveryhistoryjob).

-> **NOTE:** Available since v1.139.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}
resource "random_integer" "default" {
  min = 10000
  max = 99999
}
data "alicloud_regions" "example" {
  current = true
}
data "alicloud_account" "example" {}

resource "alicloud_log_project" "example" {
  project_name = "${var.name}-${random_integer.default.result}"
  description  = "tf actiontrail example"
}

resource "alicloud_actiontrail_trail" "example" {
  trail_name      = "${var.name}-${random_integer.default.result}"
  sls_project_arn = "acs:log:${data.alicloud_regions.example.regions.0.id}:${data.alicloud_account.example.id}:project/${alicloud_log_project.example.name}"
}

resource "alicloud_actiontrail_history_delivery_job" "example" {
  trail_name = alicloud_actiontrail_trail.example.id
}
```

## Argument Reference

The following arguments are supported:
* `trail_name` - (Required, ForceNew) The Track Name.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The creation time of the resource
* `status` - The status of the resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the History Delivery Job.
* `delete` - (Defaults to 5 mins) Used when delete the History Delivery Job.

## Import

Action Trail History Delivery Job can be imported using the id, e.g.

```shell
$ terraform import alicloud_actiontrail_history_delivery_job.example <id>
```