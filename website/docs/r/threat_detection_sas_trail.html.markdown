---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_sas_trail"
description: |-
  Provides a Alicloud Threat Detection Sas Trail resource.
---

# alicloud_threat_detection_sas_trail

Provides a Threat Detection Sas Trail resource. 

For information about Threat Detection Sas Trail and how to use it, see [What is Sas Trail](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.212.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_sas_trail" "default" {
}
```

## Argument Reference

The following arguments are supported:

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as ``.
* `create_time` - The service trace creation timestamp, in milliseconds.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Sas Trail.
* `delete` - (Defaults to 5 mins) Used when delete the Sas Trail.

## Import

Threat Detection Sas Trail can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_sas_trail.example 
```