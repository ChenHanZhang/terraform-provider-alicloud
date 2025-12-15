---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_check_structure"
description: |-
  Provides a Alicloud Threat Detection Check Structure resource.
---

# alicloud_threat_detection_check_structure

Provides a Threat Detection Check Structure resource.

Cloud platform configuration check item structure information.

For information about Threat Detection Check Structure and how to use it, see [What is Check Structure](https://next.api.alibabacloud.com/document/Sas/2018-12-03/GetCheckStructure).

-> **NOTE:** Available since v1.266.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_check_structure" "default" {
}
```

### Deleting `alicloud_threat_detection_check_structure` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_check_structure`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as ``.


## Import

Threat Detection Check Structure can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_check_structure.example 
```