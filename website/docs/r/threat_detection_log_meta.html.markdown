---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_log_meta"
description: |-
  Provides a Alicloud Threat Detection Log Meta resource.
---

# alicloud_threat_detection_log_meta

Provides a Threat Detection Log Meta resource.

Log Analysis Delivery Status.

For information about Threat Detection Log Meta and how to use it, see [What is Log Meta](https://next.api.alibabacloud.com/document/Sas/2018-12-03/ModifyLogMetaStatus).

-> **NOTE:** Available since v1.245.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}


resource "alicloud_threat_detection_log_meta" "default" {
  status        = "disabled"
  log_meta_name = "aegis-log-client"
}
```

### Deleting `alicloud_threat_detection_log_meta` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_log_meta`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `log_meta_name` - (Required, ForceNew) The name of the dedicated Logstore where logs are stored.

-> **NOTE:**  You can call the [DescribeLogMeta](~~DescribeLogMeta~~) operation to obtain this parameter.

* `status` - (Required) Specifies the status of the log to be modified. Valid values:
  - `enabled`: Enabled
  - `disabled`: Disabled

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Log Meta.
* `update` - (Defaults to 5 mins) Used when update the Log Meta.

## Import

Threat Detection Log Meta can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_log_meta.example <log_meta_name>
```