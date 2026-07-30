---
subcategory: "Message Service"
layout: "alicloud"
page_title: "Alicloud: alicloud_message_service_account_logging"
description: |-
  Provides a Alicloud Message Service Account Logging resource.
---

# alicloud_message_service_account_logging

Provides a Message Service Account Logging resource.

Account Logging Configuration.

For information about Message Service Account Logging and how to use it, see [What is Account Logging](https://next.api.alibabacloud.com/document/Mns-open/2022-01-19/SetAccountAttributes).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_log_project" "project" {
  description = "cloudspec example project for AccountLogging"
  name        = "cloudspec-example-project"
}

resource "alicloud_log_store" "logstore" {
  append_meta      = false
  retention_period = "30"
  shard_count      = "2"
  project          = "cloudspec-example-project"
  name             = "cloudspec-example-logstore"
}


resource "alicloud_message_service_account_logging" "default" {
  log_store_name        = "cloudspec-example-logstore"
  project_name          = "cloudspec-example-project"
  log_enabled           = true
  message_trace_enabled = false
}
```

### Deleting `alicloud_message_service_account_logging` or removing it from your configuration

Terraform cannot destroy resource `alicloud_message_service_account_logging`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `log_enabled` - (Optional, ForceNew) Whether to enable log delivery to SLS
* `log_store_name` - (Optional, ForceNew) The name of the SLS logstore
* `message_trace_enabled` - (Optional, ForceNew) Whether to enable message trace
* `project_name` - (Optional, ForceNew) The name of the SLS project

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<Alibaba Cloud Account ID>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Account Logging.
* `update` - (Defaults to 5 mins) Used when update the Account Logging.

## Import

Message Service Account Logging can be imported using the id, e.g.

```shell
$ terraform import alicloud_message_service_account_logging.example <Alibaba Cloud Account ID>
```