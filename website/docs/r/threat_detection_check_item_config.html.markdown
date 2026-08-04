---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_check_item_config"
description: |-
  Provides a Alicloud Threat Detection Check Item Config resource.
---

# alicloud_threat_detection_check_item_config

Provides a Threat Detection Check Item Config resource.

Cloud platform configuration check items.

For information about Threat Detection Check Item Config and how to use it, see [What is Check Item Config](https://next.api.alibabacloud.com/document/Sas/2018-12-03/ListCheckItem).

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


resource "alicloud_threat_detection_check_item_config" "default" {
}
```

### Deleting `alicloud_threat_detection_check_item_config` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_check_item_config`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as ``.
* `check_show_name` - The name of the check item.
* `check_type` - The source type of the Security Center check item:.
* `risk_level` - The risk level of the check item.
* `section_ids` - The list of sections associated with the check item.


## Import

Threat Detection Check Item Config can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_check_item_config.example 
```