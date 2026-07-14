---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_asset"
description: |-
  Provides a Alicloud Threat Detection Asset resource.
---

# alicloud_threat_detection_asset

Provides a Threat Detection Asset resource.



For information about Threat Detection Asset and how to use it, see [What is Asset](https://next.api.alibabacloud.com/document/Sas/2018-12-03/DescribeAssetDetailByUuid).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_asset" "default" {
}
```

### Deleting `alicloud_threat_detection_asset` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_asset`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `resource_group_id` - (Optional, Computed) The resource property field that represents the resource group.

-> **NOTE:** This parameter is only evaluated during resource operations. Modifying it in isolation will not trigger any action.

* `uuid` - (Optional, ForceNew, Computed) The UUID of the instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID of the asset.


## Import

Threat Detection Asset can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_asset.example <uuid>
```