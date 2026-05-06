---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_transit_router_policy_table"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Transit Router Policy Table resource.
---

# alicloud_cen_transit_router_policy_table

Provides a Cloud Enterprise Network (CEN) Transit Router Policy Table resource.

The Policy Table of Transitr Router.

For information about Cloud Enterprise Network (CEN) Transit Router Policy Table and how to use it, see [What is Transit Router Policy Table](https://next.api.alibabacloud.com/document/Cbn/2017-09-12/CreateTransitRouterPolicyTable).

-> **NOTE:** Available since v1.278.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-wulanchabu-example-6"
}

variable "region_id" {
  default = "cn-wulanchabu-example-6"
}

resource "alicloud_cen_instance" "defaultiEuUyd" {
  cen_instance_name = "镇元自动example用例"
}

resource "alicloud_cen_transit_router" "defaultu3Rgvn" {
  cen_id = alicloud_cen_instance.defaultiEuUyd.id
}


resource "alicloud_cen_transit_router_policy_table" "default" {
  name              = "自动化资源用例名称"
  transit_router_id = alicloud_cen_transit_router.defaultu3Rgvn.transit_router_id
  description       = "自动化资源用例描述"
  dry_run           = false
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `dry_run` - (Optional) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `name` - (Optional) The name of the resource
* `transit_router_id` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Transit Router Policy Table.
* `delete` - (Defaults to 5 mins) Used when delete the Transit Router Policy Table.
* `update` - (Defaults to 5 mins) Used when update the Transit Router Policy Table.

## Import

Cloud Enterprise Network (CEN) Transit Router Policy Table can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_transit_router_policy_table.example <policy_table_id>
```