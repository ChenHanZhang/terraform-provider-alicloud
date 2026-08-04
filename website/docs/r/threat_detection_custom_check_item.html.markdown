---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_custom_check_item"
description: |-
  Provides a Alicloud Threat Detection Custom Check Item resource.
---

# alicloud_threat_detection_custom_check_item

Provides a Threat Detection Custom Check Item resource.

Custom Check Item.

For information about Threat Detection Custom Check Item and how to use it, see [What is Custom Check Item](https://next.api.alibabacloud.com/document/Sas/2018-12-03/CreateCheckItem).

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


resource "alicloud_threat_detection_custom_check_item" "default" {
  status            = "EDIT"
  instance_sub_type = "DISK"
  description {
    type  = "text"
    value = "description"
  }
  check_show_name = "example检查项1"
  check_rule      = "{\"AssociatedData\":null,\"MatchProperty\":{\"Operator\":\"AND\",\"MatchProperties\":[{\"DataName\":\"ACS_ECS_Disk\",\"PropertyPath\":\"DiskId\",\"MatchOperator\":\"EQ\",\"MatchPropertyValue\":\"dada\"}]}}"
  vendor          = "ALIYUN"
  assist_info {
    type  = "text"
    value = "AssistInfo"
  }
  instance_type = "ECS"
  risk_level    = "HIGH"
  solution {
    type  = "text"
    value = "Solution"
  }
  section_ids = ["110", "256", "273"]
  remark      = "76de"
}
```

## Argument Reference

The following arguments are supported:
* `assist_info` - (Optional, Set) The help information for the check item description. See [`assist_info`](#assist_info) below.
* `check_rule` - (Required) The definition rule of the custom check item.
* `check_show_name` - (Required) The name of the custom check item.
* `description` - (Optional, Set) The description of the check item. See [`description`](#description) below.
* `instance_sub_type` - (Required) The asset subtype of the cloud service.

-> **NOTE:**  You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) operation to obtain this parameter.

* `instance_type` - (Required) The asset type of the cloud service.

-> **NOTE:**  You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) operation to obtain this parameter.

* `remark` - (Optional) The remark information.
* `risk_level` - (Required) The risk level of the check item. Valid values:
  - `HIGH`: high
  - `MEDIUM`: medium
  - `LOW`: low
* `section_ids` - (Required, List) The section IDs of the check item.
* `solution` - (Optional, Set) The solution for the check item. See [`solution`](#solution) below.
* `status` - (Required) The status of the check item. Valid values:
  - `EDIT`: Editing
  - `RELEASE`: Published

-> **NOTE:**  - Changing the status from `Published` to `Editing` clears all historical records.

-> **NOTE:**  - Only check items in the `Published` state can be used for checks.

* `vendor` - (Required) The cloud asset vendor.

-> **NOTE:**  You can call the [ListCloudAssetSchemas](~~ListCloudAssetSchemas~~) operation to obtain valid vendors.


### `assist_info`

The assist_info supports the following:
* `type` - (Optional) The type of help information for the check item risk. Valid values:
  - `text`: Text
* `value` - (Optional) The content of the help information for the check item risk.

### `description`

The description supports the following:
* `type` - (Optional) The type of the check item description. Valid values:
  - `text`: Text.
* `value` - (Optional) The specific content of the description.

### `solution`

The solution supports the following:
* `type` - (Optional) The information type of the solution for the check item. Valid values:
  - `text`: text
* `value` - (Optional) The content of the solution for the check item risk.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Custom Check Item.
* `delete` - (Defaults to 5 mins) Used when delete the Custom Check Item.
* `update` - (Defaults to 5 mins) Used when update the Custom Check Item.

## Import

Threat Detection Custom Check Item can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_custom_check_item.example <check_id>
```