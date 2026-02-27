---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_monitor_account"
description: |-
  Provides a Alicloud Threat Detection Monitor Account resource.
---

# alicloud_threat_detection_monitor_account

Provides a Threat Detection Monitor Account resource.

Multi-account management account.

For information about Threat Detection Monitor Account and how to use it, see [What is Monitor Account](https://next.api.alibabacloud.com/document/Sas/2018-12-03/CreateMonitorAccount).

-> **NOTE:** Available since v1.272.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_resource_manager_account" "defaulttSNZEf" {
}


resource "alicloud_threat_detection_monitor_account" "default" {
  account_ids = alicloud_resource_manager_account.defaulttSNZEf.id
}
```

### Deleting `alicloud_threat_detection_monitor_account` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_monitor_account`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `account_ids` - (Optional) The list of member account IDs in the resource directory.

-> **NOTE:**  You can use [to](~~ ~~) to obtain member account IDs. Multiple member account IDs must be separated by half-angle commas. The monitoring account list will be replaced according to the incoming list. If this parameter is not passed, the existing monitoring account list will be cleared.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<Alibaba Cloud Account ID>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Monitor Account.
* `update` - (Defaults to 5 mins) Used when update the Monitor Account.

## Import

Threat Detection Monitor Account can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_monitor_account.example <Alibaba Cloud Account ID>
```