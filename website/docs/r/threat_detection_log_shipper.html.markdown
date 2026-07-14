---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_log_shipper"
description: |-
  Provides a Alicloud Threat Detection Log Shipper resource.
---

# alicloud_threat_detection_log_shipper

Provides a Threat Detection Log Shipper resource.



For information about Threat Detection Log Shipper and how to use it, see [What is Log Shipper](https://next.api.alibabacloud.com/document/Sas/2018-12-03/ModifyOpenLogShipper).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-beijing"
}


resource "alicloud_threat_detection_log_shipper" "default" {
  from = "sas"
}
```

### Deleting `alicloud_threat_detection_log_shipper` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_log_shipper`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `from` - (Optional) The ID of the request source. Default value: `aegis`. Valid value:
  - `aegis`: Server Knight version.
  - `sas`: Security Center version.

-> **NOTE:**  Server Guard users use `aegis`, and Security Center users use **sas * *.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<Alibaba Cloud Account ID>`.
* `auth_status` - Log Analysis Service authorization status.
* `buy_status` - Cloud Security Center purchase status.
* `sls_project_status` - Log analysis project status.
* `sls_service_status` - Log Analysis Service is activated.
* `status` - Log analysis shipping activation status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Log Shipper.

## Import

Threat Detection Log Shipper can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_log_shipper.example <Alibaba Cloud Account ID>
```